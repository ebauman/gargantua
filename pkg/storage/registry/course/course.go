package course

import (
	"github.com/acorn-io/mink/pkg/stores"
	"github.com/acorn-io/mink/pkg/strategy"
	v1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	"k8s.io/apiserver/pkg/registry/rest"
)

func NewCourseStorage(courseStrategy strategy.CompleteStrategy) (rest.Storage, error) {
	return stores.NewBuilder(courseStrategy.Scheme(), &v1.Course{}).
		WithCompleteCRUD(courseStrategy).
		Build(), nil
}
