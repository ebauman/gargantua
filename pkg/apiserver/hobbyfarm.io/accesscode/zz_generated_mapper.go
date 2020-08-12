package accesscode

import objPkg "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"

type FlatAccessCode struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code"`
	Description string `json:"description"`
	Scenarios []string `json:"scenarios"`
	Courses []string `json:"courses"`
	Expiration string `json:"expiration"`
	VirtualMachineSets []string `json:"vmsets"`
	RestrictedBind bool `json:"restricted_bind"`
	RestrictedBindValue string `json:"restricted_bind_value"`
	
}

func FromInput(f *FlatAccessCode) *objPkg.AccessCode {
	obj := &objPkg.AccessCode{}
	obj.Name = f.Name
	obj.Spec.Code = f.Code
	obj.Spec.Description = f.Description
	obj.Spec.Scenarios = f.Scenarios
	obj.Spec.Courses = f.Courses
	obj.Spec.Expiration = f.Expiration
	obj.Spec.VirtualMachineSets = f.VirtualMachineSets
	obj.Spec.RestrictedBind = f.RestrictedBind
	obj.Spec.RestrictedBindValue = f.RestrictedBindValue
	
	return obj
}

func ToOutput(i *objPkg.AccessCode) *FlatAccessCode {
	obj := &FlatAccessCode{}

	obj.Name = i.Name
	obj.Code = i.Spec.Code
	obj.Description = i.Spec.Description
	obj.Scenarios = i.Spec.Scenarios
	obj.Courses = i.Spec.Courses
	obj.Expiration = i.Spec.Expiration
	obj.VirtualMachineSets = i.Spec.VirtualMachineSets
	obj.RestrictedBind = i.Spec.RestrictedBind
	obj.RestrictedBindValue = i.Spec.RestrictedBindValue
	

	return obj
}

func ToOutputList(i *[]objPkg.AccessCode) *[]FlatAccessCode {
	out := make([]FlatAccessCode, len(*i))
	for _, o := range *i {
		obj := ToOutput(&o)
		out = append(out, *obj)
	}

	return &out
}
