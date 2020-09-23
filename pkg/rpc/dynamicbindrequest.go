package rpc

import (
	"context"
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	hfClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	hfInformers "github.com/hobbyfarm/gargantua/pkg/client/informers/externalversions"
	"github.com/hobbyfarm/gargantua/pkg/converters"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

const (
	dbrIdIndex = "dynamicbindrequest.hobbyfarm.io/id-index"
)

type DynamicBindRequestServer struct {
	hfClientSet *hfClientset.Clientset
	dbrIndexer cache.Indexer
}

func setupDynamicBindRequestServer(g *grpc.Server, clientset *hfClientset.Clientset, factory hfInformers.SharedInformerFactory) {
	dbrs := DynamicBindRequestServer{}

	dbrs.hfClientSet = clientset

	inf := factory.Hobbyfarm().V1().DynamicBindRequests().Informer()
	indexers := map[string]cache.IndexFunc{dbrIdIndex: dbrIdIndexer}
	inf.AddIndexers(indexers)

	dbrs.dbrIndexer = inf.GetIndexer()

	protobuf.RegisterDynamicBindRequestServiceServer(g, dbrs)
}

func (d DynamicBindRequestServer) List(ctx context.Context, empty *protobuf.Empty) (*protobuf.DynamicBindRequestList, error) {
	list, err := d.hfClientSet.HobbyfarmV1().DynamicBindRequests().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	out := make([]*protobuf.DynamicBindRequest, len(list.Items))

	for i, v := range list.Items {
		out[i] = converters.DynamicBindRequestToRPC(v)
	}

	return &protobuf.DynamicBindRequestList{DynamicBindRequests: out}, nil
}

func (d DynamicBindRequestServer) Get(ctx context.Context, id *protobuf.ID) (*protobuf.DynamicBindRequest, error) {
	obj, err := d.dbrIndexer.ByIndex(dbrIdIndex, id.ID)
	if err != nil {
		return nil, err
	}

	if len(obj) < 1 {
		return nil, status.Errorf(codes.NotFound, "cannot locate dynamicbindrequest with id %s", id.ID)
	}

	dbr, ok := obj[0].(*hfv1.DynamicBindRequest)
	if !ok {
		return nil, status.Error(codes.Internal, "error asserting dynamicbindrequest into hfv1.DynamicBindRequest")
	}

	return converters.DynamicBindRequestToRPC(*dbr), nil
}

func (d DynamicBindRequestServer) Create(ctx context.Context, request *protobuf.DynamicBindRequest) (*protobuf.DynamicBindRequest, error) {
	toCreate := converters.DynamicBindRequestFromRPC(request)
	res, err := d.hfClientSet.HobbyfarmV1().DynamicBindRequests().Create(&toCreate)
	if err != nil {
		return nil, err
	}

	return converters.DynamicBindRequestToRPC(*res), nil
}

func (d DynamicBindRequestServer) Update(ctx context.Context, request *protobuf.DynamicBindRequest) (*protobuf.DynamicBindRequest, error) {
	toUpdate := converters.DynamicBindRequestFromRPC(request)
	res, err := d.hfClientSet.HobbyfarmV1().DynamicBindRequests().Update(&toUpdate)
	if err != nil {
		return nil, err
	}

	return converters.DynamicBindRequestToRPC(*res), nil
}

func (d DynamicBindRequestServer) Delete(ctx context.Context, id *protobuf.ID) (*protobuf.Empty, error) {
	err := d.hfClientSet.HobbyfarmV1().DynamicBindRequests().Delete(id.ID, &metav1.DeleteOptions{})
	if err != nil {
		return nil, err
	}

	return &protobuf.Empty{}, nil
}

func dbrIdIndexer(obj interface{}) ([]string, error) {
	dbr, ok := obj.(*hfv1.DynamicBindRequest)
	if !ok {
		return []string{}, nil
	}
	return []string{dbr.Spec.Id}, nil
}