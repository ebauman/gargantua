package main

import (
	"context"
	"crypto/tls"
	"flag"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/ebauman/crder"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	hfClientset "github.com/hobbyfarm/gargantua/v3/pkg/client/clientset/versioned"
	"github.com/hobbyfarm/gargantua/v3/pkg/crd"
	"github.com/hobbyfarm/gargantua/v3/pkg/microservices"
	"github.com/hobbyfarm/gargantua/v3/pkg/signals"
	tls2 "github.com/hobbyfarm/gargantua/v3/pkg/tls"
	"github.com/hobbyfarm/gargantua/v3/pkg/util"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/golang/glog"
	userservice "github.com/hobbyfarm/gargantua/services/usersvc/v3/internal"
	hfInformers "github.com/hobbyfarm/gargantua/v3/pkg/client/informers/externalversions"

	"github.com/hobbyfarm/gargantua/v3/protos/authn"
	"github.com/hobbyfarm/gargantua/v3/protos/authr"
	"github.com/hobbyfarm/gargantua/v3/protos/rbac"
	"github.com/hobbyfarm/gargantua/v3/protos/user"
)

var (
	userTLSCert      string
	userTLSKey       string
	userTLSCA        string
	localMasterUrl   string
	localKubeconfig  string
	webhookTLSCA     string
	enableReflection bool
)

func init() {
	flag.StringVar(&localKubeconfig, "kubeconfig", "", "Path to kubeconfig of local cluster. Only required if out-of-cluster.")
	flag.StringVar(&localMasterUrl, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&userTLSCert, "user-tls-cert", "/etc/ssl/certs/tls.crt", "Path to TLS certificate for user server")
	flag.StringVar(&userTLSKey, "user-tls-key", "/etc/ssl/certs/tls.key", "Path to TLS key for user server")
	flag.StringVar(&userTLSCA, "user-tls-ca", "/etc/ssl/certs/ca.crt", "Path to CA cert for user server")
	flag.StringVar(&webhookTLSCA, "webhook-tls-ca", "/webhook-secret/ca.crt", "Path to CA cert for webhook server")
	flag.BoolVar(&enableReflection, "enableReflection", true, "Enable reflection")
}

func main() {
	stopCh := signals.SetupSignalHandler()
	flag.Parse()

	const (
		ClientGoQPS   = 100
		ClientGoBurst = 100
	)
	cfg, err := rest.InClusterConfig()
	if err != nil {
		cfg, err = clientcmd.BuildConfigFromFlags(localMasterUrl, localKubeconfig)
		if err != nil {
			glog.Fatalf("Error building kubeconfig: %s", err.Error())
		}
	}
	cfg.QPS = ClientGoQPS
	cfg.Burst = ClientGoBurst

	namespace := util.GetReleaseNamespace()

	hfClient, err := hfClientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatal(err)
	}

	hfInformerFactory := hfInformers.NewSharedInformerFactoryWithOptions(hfClient, time.Second*30, hfInformers.WithNamespace(namespace))

	ca, err := os.ReadFile(webhookTLSCA)
	if err != nil {
		glog.Fatalf("error reading ca certificate: %s", err.Error())
	}

	crds := userservice.GenerateUserCRD(string(ca), crd.ServiceReference{
		Namespace: util.GetReleaseNamespace(),
		Name:      "hobbyfarm-webhook",
	})

	glog.Info("installing/updating user CRD")
	err = crder.InstallUpdateCRDs(cfg, crds...)
	if err != nil {
		glog.Fatalf("failed installing/updating user crd: %s", err.Error())
	}
	glog.Info("finished installing/updating user CRD")

	ctx := context.Background()

	cert, err := tls2.ReadKeyPair(userTLSCert, userTLSKey)
	if err != nil {
		glog.Fatalf("error generating x509keypair from conversion cert and key: %s", err)
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{*cert},
	})

	rbacConn, err := microservices.EstablishConnection(microservices.Rbac, userTLSCA)
	if err != nil {
		glog.Fatalf("failed connecting to service rbac-service: %v", err)
	}
	defer rbacConn.Close()

	rbacClient := rbac.NewRbacSvcClient(rbacConn)

	authnConn, err := microservices.EstablishConnection(microservices.AuthN, userTLSCA)
	if err != nil {
		glog.Fatalf("failed connecting to service authn-service: %v", err)
	}
	defer authnConn.Close()

	authnClient := authn.NewAuthNClient(authnConn)

	authrConn, err := microservices.EstablishConnection(microservices.AuthR, userTLSCA)
	if err != nil {
		glog.Fatalf("failed connecting to service authr-service: %v", err)
	}
	defer authrConn.Close()

	authrClient := authr.NewAuthRClient(authrConn)

	gs := microservices.CreateGRPCServer(creds)
	us, err := userservice.NewGrpcUserServer(hfClient, hfInformerFactory, ctx)

	if err != nil {
		glog.Fatalf("starting grpc user server failed: %v", err)
	}

	user.RegisterUserSvcServer(gs, us)
	if enableReflection {
		reflection.Register(gs)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		grpcPort := os.Getenv("GRPC_PORT")
		if grpcPort == "" {
			grpcPort = "8080"
		}
		l, errr := net.Listen("tcp", ":"+grpcPort)
		if errr != nil {
			glog.Fatalf("Can not serve grpc: %v", errr)
		}
		glog.Info("grpc user server listening on " + grpcPort)
		glog.Fatal(gs.Serve(l))
	}()

	go func() {
		defer wg.Done()

		r := mux.NewRouter()
		userServer, err := userservice.NewUserServer(authnClient, authrClient, rbacClient, us)
		if err != nil {
			glog.Fatal(err)
		}

		userServer.SetupRoutes(r)
		http.Handle("/", r)

		apiPort := os.Getenv("PORT")
		if apiPort == "" {
			apiPort = "80"
		}

		glog.Info("http user server listening on " + apiPort)
		glog.Fatal(http.ListenAndServe(":"+apiPort, handlers.CORS(microservices.CORS_HANDLER_ALLOWED_HEADERS, microservices.CORS_HANDLER_ALLOWED_METHODS, microservices.CORS_HANDLER_ALLOWED_ORIGINS)(r)))
	}()

	hfInformerFactory.Start(stopCh)
	wg.Wait()
}
