package rpc

import (
	"context"
	hfClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	"github.com/hobbyfarm/gargantua/pkg/converters"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	"google.golang.org/grpc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EnvironmentServer struct {
	hfClientSet *hfClientset.Clientset
}

func setupEnvironmentServer(g *grpc.Server, clientset *hfClientset.Clientset) {
	es := EnvironmentServer{}

	es.hfClientSet = clientset

	protobuf.RegisterEnvironmentServiceServer(g, es)
}

func (e EnvironmentServer) List(ctx context.Context, empty *protobuf.Empty) (*protobuf.EnvironmentList, error) {
	list, err := e.hfClientSet.HobbyfarmV1().Environments().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	out := make([]*protobuf.Environment, len(list.Items))

	for i, v := range list.Items {
		out[i] = converters.EnvironmentToRPC(v)
	}

	return &protobuf.EnvironmentList{Environments: out}, nil
}

func (e EnvironmentServer) Get(ctx context.Context, id *protobuf.ID) (*protobuf.Environment, error) {
	obj, err := e.hfClientSet.HobbyfarmV1().Environments().Get(id.ID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return converters.EnvironmentToRPC(*obj), nil
}

func (e EnvironmentServer) Create(ctx context.Context, environment *protobuf.Environment) (*protobuf.Environment, error) {
	toCreate := converters.EnvironmentFromRPC(environment)
	res, err := e.hfClientSet.HobbyfarmV1().Environments().Create(&toCreate)
	if err != nil {
		return nil, err
	}

	return converters.EnvironmentToRPC(*res), nil
}

func (e EnvironmentServer) Update(ctx context.Context, environment *protobuf.Environment) (*protobuf.Environment, error) {
	toUpdate := converters.EnvironmentFromRPC(environment)
	res, err := e.hfClientSet.HobbyfarmV1().Environments().Update(&toUpdate)
	if err != nil {
		return nil, err
	}

	return converters.EnvironmentToRPC(*res), nil
}

func (e EnvironmentServer) Delete(ctx context.Context, id *protobuf.ID) (*protobuf.Empty, error) {
	err := e.hfClientSet.HobbyfarmV1().Environments().Delete(id.ID, &metav1.DeleteOptions{})
	if err != nil {
		return nil, err
	}

	return &protobuf.Empty{}, nil
}