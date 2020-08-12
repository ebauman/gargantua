package generate

import (
	"github.com/hobbyfarm/gargantua/pkg/apiserver"
)

type PreFunc func(ctx *apiserver.Context) (interface{}, error)
type ListFunc func() (interface{}, error)
type GetFunc func(name string) (interface{}, error)
type CreateUpdateFunc func(obj interface{}) (interface{}, error)
type DeleteFunc func(name string) error

/**
struct {
				Kind         string
				KindPackage  string
				GroupConfig  Group
				ObjectConfig ObjectConfig
				FieldMap     []reflect.StructField
			}
 */

type MethodConfig struct {
	Generate   bool
	AuthN      string
	AuthZ      string
	PreFunc    *PreFunc
	MethodFunc *interface{}
}

type ObjectConfig struct {
	Type interface{}

	NameOverride string
	PathOverride string

	AuthN string
	AuthZ string

	ListConfig   MethodConfig
	GetConfig    MethodConfig
	CreateConfig MethodConfig
	UpdateConfig MethodConfig
	DeleteConfig MethodConfig
}

type Group struct {
	Types            []ObjectConfig
	ClientSetPackage string
	AuthN            string
	AuthZ            string
}

type Config struct {
	Groups     map[string]Group
	OutputPath string
}
