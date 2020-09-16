package rpc

import (
	"context"
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	hfClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	hfInformers "github.com/hobbyfarm/gargantua/pkg/client/informers/externalversions"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	"google.golang.org/grpc"
	"k8s.io/client-go/tools/cache"
)

const (
	idIndex = "vms.hobbyfarm.io/id-index"
)

type VirtualMachineServer struct {
	hfClientSet *hfClientset.Clientset
	vmIndexer cache.Indexer
}

func NewVirtualMachineServer(g *grpc.Server, hfClientset *hfClientset.Clientset, hfInformerFactory hfInformers.SharedInformerFactory) (*VirtualMachineServer, error) {
	vms := &VirtualMachineServer{}

	vms.hfClientSet = hfClientset

	inf := hfInformerFactory.Hobbyfarm().V1().VirtualMachines().Informer()
	indexers := map[string]cache.IndexFunc{idIndex: vmIdIndexer}
	inf.AddIndexers(indexers)
	vms.vmIndexer = inf.GetIndexer()

	protobuf.RegisterVirtualMachineServiceServer(g, vms)
}

func (vms *VirtualMachineServer) List(ctx context.Context, empty *protobuf.Empty) (*protobuf.VirtualMachineList, error) {

	vmlist := &protobuf.VirtualMachineList{

	}
}

func (vms *VirtualMachineServer) Get(ctx context.Context, id *protobuf.ID) (*protobuf.VirtualMachine, error) {
	panic("implement me")
}

func (vms *VirtualMachineServer) Create(ctx context.Context, machine *protobuf.VirtualMachine) (*protobuf.VirtualMachine, error) {
	panic("implement me")
}

func (vms *VirtualMachineServer) Update(ctx context.Context, machine *protobuf.VirtualMachine) (*protobuf.VirtualMachine, error) {
	panic("implement me")
}

func (vms *VirtualMachineServer) Delete(ctx context.Context, id *protobuf.ID) (*protobuf.Empty, error) {
	panic("implement me")
}

func vmIdIndexer(obj interface{}) ([]string, error) {
	vm, ok := obj.(*hfv1.VirtualMachine)
	if !ok {
		return []string{}, nil
	}
	return []string{vm.Spec.Id}, nil
}