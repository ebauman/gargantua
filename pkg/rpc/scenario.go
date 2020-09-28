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
	scenarioIdIndex = "scenarioserver.hobbyfarm.io/id-index"
)

type ScenarioServer struct {
	hfClientSet *hfClientset.Clientset
	scenarioIndexer cache.Indexer
}

func setupScenarioServer(g *grpc.Server, clientset *hfClientset.Clientset, factory hfInformers.SharedInformerFactory) {
	ss := ScenarioServer{}

	ss.hfClientSet = clientset
	inf := factory.Hobbyfarm().V1().Scenarios().Informer()
	indexers := map[string]cache.IndexFunc{scenarioIdIndex: scenarioIdIndexer}
	inf.AddIndexers(indexers)

	ss.scenarioIndexer = inf.GetIndexer()

	protobuf.RegisterScenarioServiceServer(g, ss)
}

func (s ScenarioServer) List(ctx context.Context, empty *protobuf.Empty) (*protobuf.ScenarioList, error) {
	list, err := s.hfClientSet.HobbyfarmV1().Scenarios().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	out := make([]*protobuf.Scenario, len(list.Items))

	for i, v := range list.Items {
		out[i] = converters.ScenarioToRPC(v)
	}

	return &protobuf.ScenarioList{Scenarios: out}, nil
}

func (s ScenarioServer) Get(ctx context.Context, id *protobuf.ID) (*protobuf.Scenario, error) {
	obj, err := s.scenarioIndexer.ByIndex(scenarioIdIndex, id.ID)
	if err != nil {
		return nil, err
	}

	if len(obj) < 1 {
		return nil, status.Errorf(codes.NotFound, "cannot locate scenario with id %s", id.ID)
	}

	scenario, ok := obj[0].(*hfv1.Scenario)
	if !ok {
		return nil, status.Error(codes.Internal, "error asserting scenario into hfv1.Scenario")
	}

	return converters.ScenarioToRPC(*scenario), nil
}

func (s ScenarioServer) Create(ctx context.Context, scenario *protobuf.Scenario) (*protobuf.Scenario, error) {
	toCreate := converters.ScenarioFromRPC(scenario)
	res, err := s.hfClientSet.HobbyfarmV1().Scenarios().Create(&toCreate)
	if err != nil {
		return nil, err
	}

	return converters.ScenarioToRPC(*res), nil
}

func (s ScenarioServer) Update(ctx context.Context, scenario *protobuf.Scenario) (*protobuf.Scenario, error) {
	toUpdate := converters.ScenarioFromRPC(scenario)
	res, err := s.hfClientSet.HobbyfarmV1().Scenarios().Update(&toUpdate)
	if err != nil {
		return nil, err
	}
	return converters.ScenarioToRPC(*res), nil
}

func (s ScenarioServer) Delete(ctx context.Context, id *protobuf.ID) (*protobuf.Empty, error) {
	err := s.hfClientSet.HobbyfarmV1().Scenarios().Delete(id.ID, &metav1.DeleteOptions{})
	if err != nil {
		return nil, err
	}

	return &protobuf.Empty{}, nil
}

func scenarioIdIndexer(obj interface{}) ([]string, error) {
	scenario, ok := obj.(*hfv1.Scenario)
	if !ok {
		return []string{}, nil
	}
	return []string{scenario.Spec.Id}, nil
}