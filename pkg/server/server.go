package server

import (
	restconfig "github.com/acorn-io/baaah/pkg/restconfig"
	"github.com/acorn-io/mink/pkg/server"
	"github.com/hobbyfarm/gargantua/v3/pkg/openapi"
	"github.com/hobbyfarm/gargantua/v3/pkg/scheme"
	"github.com/hobbyfarm/gargantua/v3/pkg/server/authz"
	"github.com/hobbyfarm/gargantua/v3/pkg/storage/providers/kubernetes"
	"github.com/hobbyfarm/gargantua/v3/pkg/storage/providers/mysql"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/apiserver/pkg/server/mux"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Config struct {
	HTTPPort       int
	HTTPSPort      int
	DSN            string
	KubeconfigPath string
	KubeContext    string
}

func New(cfg Config) (*server.Server, error) {
	var apiGroups *genericapiserver.APIGroupInfo
	var err error

	if cfg.DSN != "" {
		apiGroups, err = mysql.APIGroups(cfg.DSN)
		if err != nil {
			return nil, err
		}
	} else {
		var restConfig *rest.Config
		var err error

		if cfg.KubeconfigPath != "" {
			restConfig, err = restconfig.FromFile(cfg.KubeconfigPath, cfg.KubeContext)
			if err != nil {
				return nil, err
			}
		} else {
			restConfig, err = restconfig.New(scheme.Scheme)
			if err != nil {
				return nil, err
			}
		}

		kclient, err := client.NewWithWatch(restConfig, client.Options{
			Scheme: scheme.Scheme,
		})
		if err != nil {
			return nil, err
		}

		apiGroups, err = kubernetes.NewKubernetesStorage(kclient)
		if err != nil {
			return nil, err
		}
	}

	svr, err := server.New(&server.Config{
		Name:                  "hobbyfarm",
		Authenticator:         nil,
		Authorization:         authz.AuthZ{},
		HTTPListenPort:        cfg.HTTPPort,
		HTTPSListenPort:       cfg.HTTPSPort,
		LongRunningVerbs:      []string{"watch"},
		LongRunningResources:  nil,
		OpenAPIConfig:         openapi.GetOpenAPIDefinitions,
		Scheme:                scheme.Scheme,
		CodecFactory:          &scheme.Codecs,
		APIGroups:             []*genericapiserver.APIGroupInfo{apiGroups},
		Middleware:            nil,
		PostStartFunc:         nil,
		SupportAPIAggregation: false,
	})
	if err != nil {
		return nil, err
	}

	return svr, nil
}

func unregisterPaths(mux *mux.PathRecorderMux, paths ...string) {
	for _, p := range paths {
		mux.Unregister(p)
	}
}
