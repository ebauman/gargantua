package codegen

import (
	"github.com/hobbyfarm/gargantua/pkg/apiserver"
	"k8s.io/apimachinery/pkg/runtime"
)

type PreFunc func(ctx *apiserver.Context) (interface{}, error)
type ListFunc func() (interface{}, error)
type GetFunc func(name string) (interface{}, error)
type CreateUpdateFunc func(obj interface{}) (interface{}, error)
type DeleteFunc func(name string) error

type Config struct {
	Object runtime.Object

	OutputPath string
	GenerateList bool
	GenerateGet bool
	GenerateCreate bool
	GenerateUpdate bool
	GenerateDelete bool

	PreList  PreFunc
	ListFunc ListFunc

	PreGet  PreFunc
	GetFunc GetFunc

	PreCreate  PreFunc
	CreateFunc CreateUpdateFunc

	PreUpdate  PreFunc
	UpdateFunc CreateUpdateFunc

	PreDelete  PreFunc
	DeleteFunc DeleteFunc
}