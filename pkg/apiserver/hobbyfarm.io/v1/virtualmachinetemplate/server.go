package virtualmachinetemplate

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/apiserver"
	"github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned/typed/hobbyfarm.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type VirtualMachineTemplateServer struct {
	client v1.VirtualMachineTemplateInterface

	course hfv1.Course
}

const (
	Group = "hobbyfarm.io"
	Version = "v1"
	Kind = "virtualmachinetemplates"
)

// autogenerated
func Register(server *apiserver.APIServer, templateInterface v1.VirtualMachineTemplateInterface) {

	var gvk = []string{Group, Version, Kind}
	var path = strings.Join(gvk, "/")

	vmts := &VirtualMachineTemplateServer{}
	vmts.client = templateInterface



	// pre_ gets a context. returns interface{}, error
	// action_ gets a context, and ...interface{}. returns (error, [object_or_collection])
	// post_ gets a context, and ...interface{}. returns error

	server.Get(path, vmts.do_list) // Registers the action with upstream HTTP provider
	server.Get(path + "/:name", vmts.do_get)
}

func (vmts *VirtualMachineTemplateServer) do_list(ctx *apiserver.Context) (interface{}, error) {
	return vmts.ListVirtualMachineTemplate()
}

func (vmts *VirtualMachineTemplateServer) do_get(ctx *apiserver.Context) (interface{}, error) {
	return vmts.GetVirtualMachineTemplate(ctx.Fiber.Get("name"))
}

func (vmts *VirtualMachineTemplateServer) do_create(ctx *apiserver.Context) (interface{}, error) {
	// need to figure out how to deserialize an inbound object
	// probably need to look at defining schema or requiring that these objects
	// have an OpenAPI spec associated so that validation can be performed.
}

func (vmts *VirtualMachineTemplateServer) ListVirtualMachineTemplate() (*[]hfv1.VirtualMachineTemplate, error) {
	data, err := vmts.client.List(metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	return &data.Items, nil
}

func (vmts *VirtualMachineTemplateServer) GetVirtualMachineTemplate(name string) (*hfv1.VirtualMachineTemplate, error) {
	data, err := vmts.client.Get(name, metav1.GetOptions{})

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (vmts *VirtualMachineTemplateServer) pre_CreateVirtualMachineTemplate(ctx *apiserver.Context) (*hfv1.VirtualMachineTemplate, error) {

}

func (vmts *VirtualMachineTemplateServer) CreateVirtualMachineTemplate(template hfv1.VirtualMachineTemplate) (*hfv1.VirtualMachineTemplate, error) {
	data, err := vmts.client.Create(&template)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (vmts *VirtualMachineTemplateServer) UpdateVirtualMachineTemplate(template hfv1.VirtualMachineTemplate) (*hfv1.VirtualMachineTemplate, error) {
	data, err := vmts.client.Update(&template)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (vmts *VirtualMachineTemplateServer) DeleteVirtualMachineTemplate(name string) (error)  {
	err := vmts.client.Delete(name, &metav1.DeleteOptions{})

	return err
}