package kubernetes

import (
	"github.com/acorn-io/mink/pkg/serializer"
	v1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	hobbyfarmScheme "github.com/hobbyfarm/gargantua/v3/pkg/scheme"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewKubernetesStorage(client client.WithWatch) (*genericapiserver.APIGroupInfo, error) {
	scheme := hobbyfarmScheme.Scheme

	if err := v1.AddToScheme(scheme); err != nil {
		return nil, err
	}

	apiGroups, err := APIGroups(client)
	if err != nil {
		return nil, err
	}

	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(
		v1.SchemeGroupVersion.Group,
		scheme,
		hobbyfarmScheme.ParameterCodec,
		hobbyfarmScheme.Codecs)
	apiGroupInfo.VersionedResourcesStorageMap["v1"] = apiGroups
	apiGroupInfo.NegotiatedSerializer = serializer.NewNoProtobufSerializer(apiGroupInfo.NegotiatedSerializer)

	return &apiGroupInfo, nil
}
