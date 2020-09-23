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

}

func (s ScenarioServer) Create(ctx context.Context, scenario *protobuf.Scenario) (*protobuf.Scenario, error) {
	panic("implement me")
}

func (s ScenarioServer) Update(ctx context.Context, scenario *protobuf.Scenario) (*protobuf.Scenario, error) {
	panic("implement me")
}

func (s ScenarioServer) Delete(ctx context.Context, id *protobuf.ID) (*protobuf.Empty, error) {
	panic("implement me")
}

func scenarioIdIndexer(obj interface{}) ([]string, error) {
	scenario, ok := obj.(*hfv1.Scenario)
	if !ok {
		return []string{}, nil
	}
	return []string{scenario.Spec.Id}, nil
}