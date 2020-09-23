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
	vmtIdIndex = "vmts.hobbyfarm.io/id-index"
	vmtNameIndex = "vmts.hobbyfarm.io/name-index"
)

type VirtualMachineTemplateServer struct {
	hfClientSet *hfClientset.Clientset

	vmTemplateIndexer cache.Indexer
}

func setupVirtualMachineTemplateServer(g *grpc.Server, clientset *hfClientset.Clientset, factory hfInformers.SharedInformerFactory) {
	vmts := VirtualMachineTemplateServer{}

	vmts.hfClientSet = clientset

	inf := factory.Hobbyfarm().V1().VirtualMachineTemplates().Informer()
	indexers := map[string]cache.IndexFunc{vmtIdIndex: vmtIdIndexer, vmtNameIndex: vmtNameIndexer}
	inf.AddIndexers(indexers)
	vmts.vmTemplateIndexer = inf.GetIndexer()

	protobuf.RegisterVirtualMachineTemplateServiceServer(g, vmts)
}


func (v VirtualMachineTemplateServer) List(ctx context.Context, empty *protobuf.Empty) (*protobuf.VirtualMachineTemplateList, error) {
	list, err := v.hfClientSet.HobbyfarmV1().VirtualMachineTemplates().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	out := make([]*protobuf.VirtualMachineTemplate, len(list.Items))

	for i, v := range list.Items {
		out[i] = converters.VirtualMachineTemplateToRPC(v)
	}

	return &protobuf.VirtualMachineTemplateList{VirtualMachineTemplates: out}, nil
}

func (v VirtualMachineTemplateServer) Get(ctx context.Context, id *protobuf.ID) (*protobuf.VirtualMachineTemplate, error) {
	obj, err := v.hfClientSet.HobbyfarmV1().VirtualMachineTemplates().Get(id.ID, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return converters.VirtualMachineTemplateToRPC(*obj), nil
}

func (v VirtualMachineTemplateServer) Create(ctx context.Context, template *protobuf.VirtualMachineTemplate) (*protobuf.VirtualMachineTemplate, error) {
	toCreate := converters.VirtualMachineTemplateFromRPC(template)
	res, err := v.hfClientSet.HobbyfarmV1().VirtualMachineTemplates().Create(&toCreate)
	if err != nil {
		return nil, err
	}

	return converters.VirtualMachineTemplateToRPC(*res), nil
}

func (v VirtualMachineTemplateServer) Update(ctx context.Context, template *protobuf.VirtualMachineTemplate) (*protobuf.VirtualMachineTemplate, error) {
	toUpdate := converters.VirtualMachineTemplateFromRPC(template)
	res, err := v.hfClientSet.HobbyfarmV1().VirtualMachineTemplates().Update(&toUpdate)
	if err != nil {
		return nil, err
	}

	return converters.VirtualMachineTemplateToRPC(*res), nil
}

func (v VirtualMachineTemplateServer) Delete(ctx context.Context, id *protobuf.ID) (*protobuf.Empty, error) {
	err := v.hfClientSet.HobbyfarmV1().VirtualMachineTemplates().Delete(id.ID, &metav1.DeleteOptions{})
	if err != nil {
		return nil, err
	}

	return &protobuf.Empty{}, nil
}

func vmtIdIndexer(obj interface{}) ([]string, error) {
	vmt, ok := obj.(*hfv1.VirtualMachineTemplate)
	if !ok {
		return []string{}, nil
	}
	return []string{vmt.Spec.Id}, nil
}

func vmtNameIndexer(obj interface{}) ([]string, error) {
	vmt, ok := obj.(*hfv1.VirtualMachineTemplate)
	if !ok {
		return []string{}, nil
	}
	return []string{vmt.Spec.Name}, nil
}
