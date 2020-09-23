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
	accessCodeIdIndex = "accesscodes.hobbyfarm.io/id-index"
)

type AccessCodeServer struct {
	hfClientset *hfClientset.Clientset
	accessCodeIndexer cache.Indexer
}

func setupAccessCodeServer(g *grpc.Server, clientset *hfClientset.Clientset, factory hfInformers.SharedInformerFactory) {
	acs := AccessCodeServer{}

	acs.hfClientset = clientset

	inf := factory.Hobbyfarm().V1().AccessCodes().Informer()
	indexers := map[string]cache.IndexFunc{accessCodeIdIndex: accessCodeIdIndexer}

	inf.AddIndexers(indexers)
	acs.accessCodeIndexer = inf.GetIndexer()

	protobuf.RegisterAccessCodeServiceServer(g, acs)
}

func (a AccessCodeServer) List(ctx context.Context, empty *protobuf.Empty) (*protobuf.AccessCodeList, error) {
	list, err := a.hfClientset.HobbyfarmV1().AccessCodes().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	out := make([]*protobuf.AccessCode, len(list.Items))

	for i, v := range list.Items {
		out[i] = converters.AccessCodeToRPC(v)
	}

	return &protobuf.AccessCodeList{AccessCodes: out}, nil
}

func (a AccessCodeServer) Get(ctx context.Context, id *protobuf.ID) (*protobuf.AccessCode, error) {
	obj, err := a.accessCodeIndexer.ByIndex(accessCodeIdIndex, id.ID)
	if err != nil {
		return nil, err
	}

	if len(obj) < 1 {
		return nil, status.Errorf(codes.NotFound, "cannot locate access code with id %s", id.ID)
	}

	accessCode, ok := obj[0].(*hfv1.AccessCode)
	if !ok {
		return nil, status.Errorf(codes.Internal, "error asserting access code into hfv1.AccessCode: %v", err)
	}

	return converters.AccessCodeToRPC(*accessCode), nil
}

func (a AccessCodeServer) Create(ctx context.Context, code *protobuf.AccessCode) (*protobuf.AccessCode, error) {
	toCreate := converters.AccessCodeFromRPC(code)
	res, err := a.hfClientset.HobbyfarmV1().AccessCodes().Create(&toCreate)
	if err != nil {
		return nil, err
	}

	return converters.AccessCodeToRPC(*res), nil
}

func (a AccessCodeServer) Update(ctx context.Context, code *protobuf.AccessCode) (*protobuf.AccessCode, error) {
	toUpdate := converters.AccessCodeFromRPC(code)
	res, err := a.hfClientset.HobbyfarmV1().AccessCodes().Update(&toUpdate)
	if err != nil {
		return nil, err
	}

	return converters.AccessCodeToRPC(*res), nil
}

func (a AccessCodeServer) Delete(ctx context.Context, id *protobuf.ID) (*protobuf.Empty, error) {
	err := a.hfClientset.HobbyfarmV1().AccessCodes().Delete(id.ID, &metav1.DeleteOptions{})
	if err != nil {
		return nil, err
	}

	return &protobuf.Empty{}, nil
}

func accessCodeIdIndexer(obj interface{}) ([]string, error) {
	accessCode, ok := obj.(*hfv1.Course)
	if !ok {
		return []string{}, nil
	}

	return []string{accessCode.Spec.Id}, nil
}