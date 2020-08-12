package accesscode

import (
	objPkg "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/apiserver"
	objClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned/typed/hobbyfarm.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type Server struct {
	client objClientset.AccessCodeInterface
	dependencies map[string]interface{}
}

const (
	Group = "hobbyfarm.io"
	Kind = "accesscode"
)

func Register(server *apiserver.APIServer, templateInterface objClientset.AccessCodeInterface) {
	var gvk = []string{Group, Kind}
	var path = strings.Join(gvk, "/")

	objServer := &Server{}
	objServer.client = templateInterface
	
	server.Get(path, objServer.do_list, "default", "")
}
func (objServer *Server) do_list(ctx *apiserver.Context) (interface{}, error) {
	actionData, err := objServer.ListAccessCode(ctx)
	if err != nil {
		return nil, err
	}

	return ToOutputList(actionData), nil
}

func (objServer *Server) ListAccessCode(ctx *apiserver.Context) (*[]objPkg.AccessCode, error) {
	data, err := objServer.client.List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return &data.Items, nil
}

func (objServer *Server) genericObjectMarshal(ctx *apiserver.Context) (*objPkg.AccessCode, error) {
	flatObj := &FlatAccessCode{}
	
	err := ctx.Fiber.BodyParser(&flatObj)
	if err != nil {
		return nil, err
	}

	err = ctx.Validator.Struct(flatObj)
	if err != nil {
		return nil, err
	}

	obj := FromInput(flatObj)
	return obj, nil
}
