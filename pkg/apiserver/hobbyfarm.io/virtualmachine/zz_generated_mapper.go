package virtualmachine

import objPkg "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"

type FlatVirtualMachine struct {
	Name string `json:"name" validate:"required"`
	Id string `json:"id"`
	VirtualMachineTemplateId string `json:"vm_template_id"`
	KeyPair string `json:"keypair_name"`
	VirtualMachineClaimId string `json:"vm_claim_id"`
	UserId string `json:"user"`
	Provision bool `json:"provision"`
	VirtualMachineSetId string `json:"vm_set_id"`
	
}

func FromInput(f *FlatVirtualMachine) *objPkg.VirtualMachine {
	obj := &objPkg.VirtualMachine{}
	obj.Name = f.Name
	obj.Spec.Id = f.Id
	obj.Spec.VirtualMachineTemplateId = f.VirtualMachineTemplateId
	obj.Spec.KeyPair = f.KeyPair
	obj.Spec.VirtualMachineClaimId = f.VirtualMachineClaimId
	obj.Spec.UserId = f.UserId
	obj.Spec.Provision = f.Provision
	obj.Spec.VirtualMachineSetId = f.VirtualMachineSetId
	
	return obj
}

func ToOutput(i *objPkg.VirtualMachine) *FlatVirtualMachine {
	obj := &FlatVirtualMachine{}

	obj.Name = i.Name
	obj.Id = i.Spec.Id
	obj.VirtualMachineTemplateId = i.Spec.VirtualMachineTemplateId
	obj.KeyPair = i.Spec.KeyPair
	obj.VirtualMachineClaimId = i.Spec.VirtualMachineClaimId
	obj.UserId = i.Spec.UserId
	obj.Provision = i.Spec.Provision
	obj.VirtualMachineSetId = i.Spec.VirtualMachineSetId
	

	return obj
}

func ToOutputList(i *[]objPkg.VirtualMachine) *[]FlatVirtualMachine {
	out := make([]FlatVirtualMachine, len(*i))
	for _, o := range *i {
		obj := ToOutput(&o)
		out = append(out, *obj)
	}

	return &out
}
