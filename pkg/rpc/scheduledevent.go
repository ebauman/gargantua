package rpc

import (
	hfClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	"google.golang.org/grpc"
)

type ScheduledEventServer struct {
	hfClientset *hfClientset.Clientset
}

func setupScheduledEventServer(g *grpc.Server, clientset *hfClientset.Clientset) {

}