package converters

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
)

func FromVirtualMachineRPC(req *protobuf.VirtualMachine) (*hfv1.VirtualMachine, error) {
	vm := &hfv1.VirtualMachine{
		Spec:      	FromVirtualMachineSpecRPC(req.Spec),
		Status:     hfv1.VirtualMachineStatus{
			Status:        "",
			Allocated:     false,
			Tainted:       false,
			PublicIP:      "",
			PrivateIP:     "",
			EnvironmentId: "",
			Hostname:      "",
			TFState:       "",
			WsEndpoint:    "",
		},
	}
}

func FromVirtualMachineSpecRPC(req *protobuf.VirtualMachineSpec) hfv1.VirtualMachineSpec {
	return hfv1.VirtualMachineSpec{
		Id:                       req.Id,
		VirtualMachineTemplateId: req.VirtualMachineTemplateId,
		KeyPair:                  req.KeyPair,
		VirtualMachineClaimId:    req.VirtualMachineClaimId,
		UserId:                   req.UserId,
		Provision:                req.Provision,
		VirtualMachineSetId:      req.VirtualMachineSetId,
	}
}

func FromVirtualMachineStatusRPC(req *protobuf.VirtualMachineStatus) hfv1.VirtualMachineStatus {
	return hfv1.VirtualMachineStatus{
		Status:        req.Status,
		Allocated:     false,
		Tainted:       false,
		PublicIP:      "",
		PrivateIP:     "",
		EnvironmentId: "",
		Hostname:      "",
		TFState:       "",
		WsEndpoint:    "",
	}
}

func FromVMStatusRPC(req *protobuf.VMStatus) hfv1.VmStatus {
	
}