package virtualmachinetemplate

import (
	"github.com/gofiber/fiber"
	"github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned/typed/hobbyfarm.io/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"strings"
)

type VirtualMachineTemplateServer struct {
	client v1.VirtualMachineTemplateInterface
}

const (
	Group = "hobbyfarm.io"
	Version = "v1"
	Kind = "virtualmachinetemplates"
)

func Register(app *fiber.App, templateInterface *v1.VirtualMachineTemplateInterface) {

	var gvk = []string{Group, Version, Kind}
	var path = strings.Join(gvk, "/")


	// app.Get(path, ListVirtualMachineTemplate)
}

func (vmts *VirtualMachineTemplateServer) ListVirtualMachineTemplate(c *fiber.Ctx) {
	data, err := vmts.client.List(v12.ListOptions{})
	if err != nil {
		c.Next(err)
	}


}

func GetVirtualMachineTemplate() {}

func CreateVirtualMachineTemplate() {}

func UpdateVirtualMachineTemplate() {}