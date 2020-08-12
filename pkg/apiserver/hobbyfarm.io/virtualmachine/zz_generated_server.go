package virtualmachine

import (
	objPkg "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/apiserver"
	objClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned/typed/hobbyfarm.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type Server struct {
	client objClientset.VirtualMachineInterface
	dependencies map[string]interface{}
}

const (
	Group = "hobbyfarm.io"
	Kind = "virtualmachine"
)

func Register(server *apiserver.APIServer, templateInterface objClientset.VirtualMachineInterface) {
	var gvk = []string{Group, Kind}
	var path = strings.Join(gvk, "/")

	objServer := &Server{}
	objServer.client = templateInterface
	
	server.Get(path, objServer.do_list, "default", "")
	server.Post(path, objServer.do_create, "default", "admin")
	server.Get(path + "/:name", objServer.do_get, "default", "")
	server.Put(path + "/:name", objServer.do_update, "default", "admin")
	server.Delete(path + "/:name", objServer.do_delete, "default", "admin")
}
func (objServer *Server) do_list(ctx *apiserver.Context) (interface{}, error) {
	actionData, err := objServer.ListVirtualMachine(ctx)
	if err != nil {
		return nil, err
	}

	return ToOutputList(actionData), nil
}

func (objServer *Server) ListVirtualMachine(ctx *apiserver.Context) (*[]objPkg.VirtualMachine, error) {
	data, err := objServer.client.List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return &data.Items, nil
}
func (objServer *Server) do_create(ctx *apiserver.Context) (interface{}, error) {
	preData, err := objServer.pre_CreateVirtualMachine(ctx)
	if err != nil {
		return nil, err
	}

	actionData, err := objServer.CreateVirtualMachine(preData)
	if err != nil {
		return nil, err
	}

	return ToOutput(actionData), nil
}

func (objServer *Server) pre_CreateVirtualMachine(ctx *apiserver.Context) (*objPkg.VirtualMachine, error) {
	// marshal the incoming object into a VirtualMachine
	obj, err := objServer.genericObjectMarshal(ctx)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (objServer *Server) CreateVirtualMachine(obj *objPkg.VirtualMachine) (*objPkg.VirtualMachine, error) {
	data, err := objServer.client.Create(obj)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (objServer *Server) do_get(ctx *apiserver.Context) (interface{}, error) {
	data, err := objServer.GetVirtualMachine(ctx.Fiber.Params("name"))
	if err != nil {
		return nil, err
	}

	return ToOutput(data), nil
}

func (objServer *Server) GetVirtualMachine(name string) (*objPkg.VirtualMachine, error) {
	data, err := objServer.client.Get(name, metav1.GetOptions{})

	if err != nil {
		return nil, err
	}

	return data, nil
}
func (objServer *Server) do_update(ctx *apiserver.Context) (interface{}, error) {
	preObj, err := objServer.pre_UpdateVirtualMachine(ctx)
	if err != nil {
		return nil, err
	}

	data, err := objServer.UpdateVirtualMachine(preObj)
	if err != nil {
		return nil, err
	}

	return ToOutput(data), nil
}

func (objServer *Server) pre_UpdateVirtualMachine(ctx *apiserver.Context) (*objPkg.VirtualMachine, error) {
	obj, err := objServer.genericObjectMarshal(ctx)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (objServer *Server) UpdateVirtualMachine(obj *objPkg.VirtualMachine) (*objPkg.VirtualMachine, error) {
	data, err := objServer.client.Update(obj)

	if err != nil {
		return nil, err
	}

	return data, nil
}
func (objServer *Server) do_delete(ctx *apiserver.Context) (interface{}, error) {
	err := objServer.DeleteVirtualMachine(ctx.Fiber.Params("name"))

	return nil, err
}

func (objServer *Server) DeleteVirtualMachine(name string) (error) {
	err := objServer.client.Delete(name, &metav1.DeleteOptions{})

	return err
}

func (objServer *Server) genericObjectMarshal(ctx *apiserver.Context) (*objPkg.VirtualMachine, error) {
	flatObj := &FlatVirtualMachine{}
	
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
