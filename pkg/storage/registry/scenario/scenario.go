package scenario

import (
	"github.com/acorn-io/mink/pkg/stores"
	"github.com/acorn-io/mink/pkg/strategy"
	v1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	"k8s.io/apiserver/pkg/registry/rest"
)

func NewScenarioStorage(scenarioStrategy strategy.CompleteStrategy) (rest.Storage, error) {
	return stores.NewBuilder(scenarioStrategy.Scheme(), &v1.Scenario{}).
		WithCompleteCRUD(scenarioStrategy).
		Build(), nil
}
