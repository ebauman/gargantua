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
	dbcIdIndex = "dynamicbindconfiguration.hobbyfarm.io/id-index"
)

type DynamicBindConfigurationServer struct {
	hfClientSet *hfClientset.Clientset
	dbcIndexer cache.Indexer
}

func setupDynamicBindConfigurationServer(g *grpc.Server, clientset *hfClientset.Clientset, factory hfInformers.SharedInformerFactory) {
	dbcs := DynamicBindConfigurationServer{}

	dbcs.hfClientSet = clientset
	inf := factory.Hobbyfarm().V1().DynamicBindConfigurations().Informer()
	indexers := map[string]cache.IndexFunc{dbcIdIndex: dynamicBindConfigurationIdIndexer}

	inf.AddIndexers(indexers)
	dbcs.dbcIndexer = inf.GetIndexer()

	protobuf.RegisterDynamicBindConfigurationServiceServer(g, dbcs)
}

func (d DynamicBindConfigurationServer) List(ctx context.Context, empty *protobuf.Empty) (*protobuf.DynamicBindConfigurationList, error) {
	list, err := d.hfClientSet.HobbyfarmV1().DynamicBindConfigurations().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	out := make([]*protobuf.DynamicBindConfiguration, len(list.Items))

	for i, v := range list.Items {
		out[i] = converters.DynamicBindConfigurationToRPC(v)
	}

	return &protobuf.DynamicBindConfigurationList{DynamicBindConfigurations: out}, nil
}

func (d DynamicBindConfigurationServer) Get(ctx context.Context, id *protobuf.ID) (*protobuf.DynamicBindConfiguration, error) {
	obj, err := d.dbcIndexer.ByIndex(dbcIdIndex, id.ID)
	if err != nil {
		return nil, err
	}

	if len(obj) < 1 {
		return nil, status.Errorf(codes.NotFound, "cannot locate dynamicbindconfiguration with id %s", id.ID)
	}

	dbc, ok := obj[0].(*hfv1.DynamicBindConfiguration)
	if !ok {
		return nil, status.Error(codes.Internal, "error asserting dynamicbindconfiguration into hfv1.DynamicBindConfiguration")
	}

	return converters.DynamicBindConfigurationToRPC(*dbc), nil
}

func (d DynamicBindConfigurationServer) Create(ctx context.Context, configuration *protobuf.DynamicBindConfiguration) (*protobuf.DynamicBindConfiguration, error) {
	toCreate := converters.DynamicBindConfigurationFromRPC(configuration)
	res, err := d.hfClientSet.HobbyfarmV1().DynamicBindConfigurations().Create(&toCreate)
	if err != nil {
		return nil, err
	}

	return converters.DynamicBindConfigurationToRPC(*res), nil
}

func (d DynamicBindConfigurationServer) Update(ctx context.Context, configuration *protobuf.DynamicBindConfiguration) (*protobuf.DynamicBindConfiguration, error) {
	toUpdate := converters.DynamicBindConfigurationFromRPC(configuration)
	res, err := d.hfClientSet.HobbyfarmV1().DynamicBindConfigurations().Update(&toUpdate)
	if err != nil {
		return nil, err
	}

	return converters.DynamicBindConfigurationToRPC(*res), nil
}

func (d DynamicBindConfigurationServer) Delete(ctx context.Context, id *protobuf.ID) (*protobuf.Empty, error) {
	err := d.hfClientSet.HobbyfarmV1().DynamicBindConfigurations().Delete(id.ID, &metav1.DeleteOptions{})

	if err != nil {
		return nil, err
	}

	return &protobuf.Empty{}, nil
}

func dynamicBindConfigurationIdIndexer(obj interface{}) ([]string, error) {
	dbc, ok := obj.(*hfv1.DynamicBindConfiguration)
	if !ok {
		return []string{}, nil
	}

	return []string{dbc.Spec.Id}, nil
}