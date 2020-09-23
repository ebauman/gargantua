package rpc

import (
	hfClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	hfInformers "github.com/hobbyfarm/gargantua/pkg/client/informers/externalversions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Setup(rpc *grpc.Server, setup chan bool, clientset *hfClientset.Clientset, factory hfInformers.SharedInformerFactory) error {
	err := setupVirtualMachineServer(rpc, clientset, factory)
	if err != nil {
		return err
	}

	reflection.Register(rpc) // TODO - should we have this? it is good for dev using grpcurl though

	setup <- true // notify that setup is complete

	select{}
}