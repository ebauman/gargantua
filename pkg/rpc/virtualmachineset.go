package rpc

import (
	"context"
	hfClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	hfInformers "github.com/hobbyfarm/gargantua/pkg/client/informers/externalversions"
	"github.com/hobbyfarm/gargantua/pkg/converters"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	"google.golang.org/grpc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type VirtualMachineSetServer struct {
	hfClientSet *hfClientset.Clientset
}

func setupVirtualMachineSetServer(g *grpc.Server, clientset *hfClientset.Clientset, hfInformerFactory hfInformers.SharedInformerFactory) {
	vmss := &VirtualMachineSetServer{}

	vmss.hfClientSet = clientset

	protobuf.RegisterVirtualMachineSetServiceServer(g, vmss)
}

func (v VirtualMachineSetServer) List(ctx context.Context, empty *protobuf.Empty) (*protobuf.VirtualMachineSetList, error) {
	list, err := v.hfClientSet.HobbyfarmV1().VirtualMachineSets().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	out := make([]*protobuf.VirtualMachineSet, len(list.Items))

	for i, v := range list.Items {
		out[i] = converters.VirtualMachineSetToRPC(v)
	}

	return &protobuf.VirtualMachineSetList{VirtualMachineSets: out}, nil
}

func (v VirtualMachineSetServer) Get(ctx context.Context, id *protobuf.ID) (*protobuf.VirtualMachineSet, error) {
	obj, err := v.hfClientSet.HobbyfarmV1().VirtualMachineSets().Get(id.ID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return converters.VirtualMachineSetToRPC(*obj), nil
}

func (v VirtualMachineSetServer) Create(ctx context.Context, set *protobuf.VirtualMachineSet) (*protobuf.VirtualMachineSet, error) {
	toCreate := converters.VirtualMachineSetFromRPC(set)
	res, err := v.hfClientSet.HobbyfarmV1().VirtualMachineSets().Create(&toCreate)
	if err != nil {
		return nil, err
	}

	return converters.VirtualMachineSetToRPC(*res), nil
}

func (v VirtualMachineSetServer) Update(ctx context.Context, set *protobuf.VirtualMachineSet) (*protobuf.VirtualMachineSet, error) {
	toUpdate := converters.VirtualMachineSetFromRPC(set)
	res, err := v.hfClientSet.HobbyfarmV1().VirtualMachineSets().Update(&toUpdate)
	if err != nil{
		return nil, err
	}

	return converters.VirtualMachineSetToRPC(*res), nil
}

func (v VirtualMachineSetServer) Delete(ctx context.Context, id *protobuf.ID) (*protobuf.Empty, error) {
	err := v.hfClientSet.HobbyfarmV1().VirtualMachineSets().Delete(id.ID, &metav1.DeleteOptions{})
	if err != nil {
		return nil, err
	}

	return &protobuf.Empty{}, nil
}
