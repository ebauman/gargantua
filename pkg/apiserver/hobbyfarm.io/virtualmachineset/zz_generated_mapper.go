package virtualmachineset

import objPkg "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"

type FlatVirtualMachineSet struct {
	Name string `json:"name" validate:"required"`
	Count int `json:"count"`
	Environment string `json:"environment"`
	VMTemplate string `json:"vm_template"`
	BaseName string `json:"base_name"`
	RestrictedBind bool `json:"restricted_bind"`
	RestrictedBindValue string `json:"restricted_bind_value"`
	
}

func FromInput(f *FlatVirtualMachineSet) *objPkg.VirtualMachineSet {
	obj := &objPkg.VirtualMachineSet{}
	obj.Name = f.Name
	obj.Spec.Count = f.Count
	obj.Spec.Environment = f.Environment
	obj.Spec.VMTemplate = f.VMTemplate
	obj.Spec.BaseName = f.BaseName
	obj.Spec.RestrictedBind = f.RestrictedBind
	obj.Spec.RestrictedBindValue = f.RestrictedBindValue
	
	return obj
}

func ToOutput(i *objPkg.VirtualMachineSet) *FlatVirtualMachineSet {
	obj := &FlatVirtualMachineSet{}

	obj.Name = i.Name
	obj.Count = i.Spec.Count
	obj.Environment = i.Spec.Environment
	obj.VMTemplate = i.Spec.VMTemplate
	obj.BaseName = i.Spec.BaseName
	obj.RestrictedBind = i.Spec.RestrictedBind
	obj.RestrictedBindValue = i.Spec.RestrictedBindValue
	

	return obj
}

func ToOutputList(i *[]objPkg.VirtualMachineSet) *[]FlatVirtualMachineSet {
	out := make([]FlatVirtualMachineSet, len(*i))
	for _, o := range *i {
		obj := ToOutput(&o)
		out = append(out, *obj)
	}

	return &out
}
