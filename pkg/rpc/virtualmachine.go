package rpc

import (
	"context"
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	hfClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	hfInformers "github.com/hobbyfarm/gargantua/pkg/client/informers/externalversions"
	"github.com/hobbyfarm/gargantua/pkg/converters"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	"google.golang.org/grpc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

const (
	idIndex = "vms.hobbyfarm.io/id-index"
)

type VirtualMachineServer struct {
	hfClientSet *hfClientset.Clientset
	vmIndexer cache.Indexer
}

func setupVirtualMachineServer(g *grpc.Server, hfClientset *hfClientset.Clientset, hfInformerFactory hfInformers.SharedInformerFactory) (error) {
	vms := &VirtualMachineServer{}

	vms.hfClientSet = hfClientset

	inf := hfInformerFactory.Hobbyfarm().V1().VirtualMachines().Informer()
	indexers := map[string]cache.IndexFunc{idIndex: vmIdIndexer}
	inf.AddIndexers(indexers)
	vms.vmIndexer = inf.GetIndexer()

	protobuf.RegisterVirtualMachineServiceServer(g, vms)

	return nil
}

func (vms *VirtualMachineServer) List(ctx context.Context, empty *protobuf.Empty) (*protobuf.VirtualMachineList, error) {
	list, err := vms.hfClientSet.HobbyfarmV1().VirtualMachines().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	convertedList := converters.ToVirtualMachineListRPC(list.Items)

	return &protobuf.VirtualMachineList{VirtualMachines: convertedList}, nil
}

func (vms *VirtualMachineServer) Get(ctx context.Context, id *protobuf.ID) (*protobuf.VirtualMachine, error) {
	obj, err := vms.hfClientSet.HobbyfarmV1().VirtualMachines().Get(id.ID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	convertedObj := converters.ToVirtualMachineRPC(*obj)

	return convertedObj, nil
}

func (vms *VirtualMachineServer) Create(ctx context.Context, machine *protobuf.VirtualMachine) (*protobuf.VirtualMachine, error) {
	toCreate := converters.FromVirtualMachineRPC(machine)
	res, err := vms.hfClientSet.HobbyfarmV1().VirtualMachines().Create(&toCreate)
	if err != nil {
		return nil, err
	}

	return converters.ToVirtualMachineRPC(*res), nil
}

func (vms *VirtualMachineServer) Update(ctx context.Context, machine *protobuf.VirtualMachine) (*protobuf.VirtualMachine, error) {
	toUpdate := converters.FromVirtualMachineRPC(machine)
	res, err := vms.hfClientSet.HobbyfarmV1().VirtualMachines().Update(&toUpdate)
	if err != nil {
		return nil, err
	}

	return converters.ToVirtualMachineRPC(*res), nil
}

func (vms *VirtualMachineServer) Delete(ctx context.Context, id *protobuf.ID) (*protobuf.Empty, error) {
	err := vms.hfClientSet.HobbyfarmV1().VirtualMachines().Delete(id.ID, &metav1.DeleteOptions{})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func vmIdIndexer(obj interface{}) ([]string, error) {
	vm, ok := obj.(*hfv1.VirtualMachine)
	if !ok {
		return []string{}, nil
	}
	return []string{vm.Spec.Id}, nil
}