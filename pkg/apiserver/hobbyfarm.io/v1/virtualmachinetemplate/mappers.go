package virtualmachinetemplate

import v1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"

type FlatVirtualMachineTemplate struct {
	Name 	  string 		    `json:"name" validate:"required"`
	// Namespace string
	Id        string            `json:"id" validate:"omitempty"`
	Image     string            `json:"image" validate:"required"`
	Resources v1.CMSStruct      `json:"resources" validate:"required"`
	CountMap  map[string]string `json:"count_map" validate:"omitempty"`
}

func FromInput(f *FlatVirtualMachineTemplate) *v1.VirtualMachineTemplate {
	obj := &v1.VirtualMachineTemplate{}
	obj.Name = f.Name
	obj.Spec.Id = f.Id
	obj.Spec.Name = f.Name
	obj.Spec.Image = f.Image
	obj.Spec.Resources = f.Resources
	obj.Spec.CountMap = f.CountMap

	return obj
}

func ToOutput(i *v1.VirtualMachineTemplate) *FlatVirtualMachineTemplate {
	obj := &FlatVirtualMachineTemplate{}

	obj.Name = i.Name
	obj.CountMap = i.Spec.CountMap
	obj.Image = i.Spec.Image
	obj.Resources = i.Spec.Resources
	obj.Id = i.Spec.Id

	return obj
}

func ToOutputList(i *[]v1.VirtualMachineTemplate) *[]FlatVirtualMachineTemplate {
	out := make([]FlatVirtualMachineTemplate, len(*i))
	for _, o := range *i {
		obj := ToOutput(&o)
		out = append(out, *obj)
	}

	return &out
}