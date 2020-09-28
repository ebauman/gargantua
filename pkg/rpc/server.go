package rpc

import (
	hfClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	hfInformers "github.com/hobbyfarm/gargantua/pkg/client/informers/externalversions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Setup(rpc *grpc.Server, setup chan bool, clientset *hfClientset.Clientset, factory hfInformers.SharedInformerFactory) error {
	setupAccessCodeServer(rpc, clientset, factory)
	setupCourseServer(rpc, clientset, factory)
	setupDynamicBindConfigurationServer(rpc, clientset, factory)
	setupDynamicBindRequestServer(rpc, clientset, factory)
	setupEnvironmentServer(rpc, clientset)
	setupScenarioServer(rpc, clientset, factory)
	setupScheduledEventServer(rpc, clientset)
	setupVirtualMachineServer(rpc, clientset, factory)
	setupVirtualMachineClaimServer(rpc, clientset, factory)
	setupVirtualMachineSetServer(rpc, clientset, factory)
	setupVirtualMachineTemplateServer(rpc, clientset, factory)

	reflection.Register(rpc) // TODO - should we have this? it is good for dev using grpcurl though

	setup <- true // notify that setup is complete

	select{}
}