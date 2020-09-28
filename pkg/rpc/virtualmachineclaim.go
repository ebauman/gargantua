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
	vmcsIdIndex = "vmcs.hobbyfarm.io/id-index"
)

type VirtualMachineClaimServer struct {
	hfClientSet *hfClientset.Clientset
	vmClaimIndexer cache.Indexer
}

func setupVirtualMachineClaimServer(g *grpc.Server, hfClientset *hfClientset.Clientset, hfInformerFactory hfInformers.SharedInformerFactory) {
	vmcs := &VirtualMachineClaimServer{}

	vmcs.hfClientSet = hfClientset

	inf := hfInformerFactory.Hobbyfarm().V1().VirtualMachineClaims().Informer()
	indexers := map[string]cache.IndexFunc{vmcsIdIndex: vmcIdIndexer}
	inf.AddIndexers(indexers)
	vmcs.vmClaimIndexer = inf.GetIndexer()

	protobuf.RegisterVirtualMachineClaimServiceServer(g, vmcs)
}

func (v VirtualMachineClaimServer) List(ctx context.Context, empty *protobuf.Empty) (*protobuf.VirtualMachineClaimList, error) {
	list, err := v.hfClientSet.HobbyfarmV1().VirtualMachineClaims().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	convertedList := make([]*protobuf.VirtualMachineClaim, 0)
	for _, v := range list.Items {
		convertedList = append(convertedList, converters.VirtualMachineClaimToRPC(v))
	}

	return &protobuf.VirtualMachineClaimList{VirtualMachineClaims: convertedList}, nil
}

func (v VirtualMachineClaimServer) Get(ctx context.Context, id *protobuf.ID) (*protobuf.VirtualMachineClaim, error) {
	obj, err := v.hfClientSet.HobbyfarmV1().VirtualMachineClaims().Get(id.ID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return converters.VirtualMachineClaimToRPC(*obj), nil
}

func (v VirtualMachineClaimServer) Create(ctx context.Context, claim *protobuf.VirtualMachineClaim) (*protobuf.VirtualMachineClaim, error) {
	toCreate := converters.VirtualMachineClaimFromRPC(claim)
	res, err := v.hfClientSet.HobbyfarmV1().VirtualMachineClaims().Create(&toCreate)
	if err != nil {
		return nil, err
	}

	return converters.VirtualMachineClaimToRPC(*res), nil
}

func (v VirtualMachineClaimServer) Update(ctx context.Context, claim *protobuf.VirtualMachineClaim) (*protobuf.VirtualMachineClaim, error) {
	toUpdate := converters.VirtualMachineClaimFromRPC(claim)
	res, err := v.hfClientSet.HobbyfarmV1().VirtualMachineClaims().Update(&toUpdate)
	if err != nil {
		return nil, err
	}

	return converters.VirtualMachineClaimToRPC(*res), nil
}

func (v VirtualMachineClaimServer) Delete(ctx context.Context, id *protobuf.ID) (*protobuf.Empty, error) {
	err := v.hfClientSet.HobbyfarmV1().VirtualMachineClaims().Delete(id.ID, &metav1.DeleteOptions{})
	if err != nil {
		return nil, err
	}

	return &protobuf.Empty{}, nil
}

func vmcIdIndexer(obj interface{}) ([]string, error) {
	vmc, ok := obj.(*hfv1.VirtualMachineClaim)
	if !ok {
		return []string{}, nil
	}
	return []string{vmc.Spec.Id}, nil
}