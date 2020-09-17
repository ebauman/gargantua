package converters

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func FromVirtualMachineListRPC(req []*protobuf.VirtualMachine) []hfv1.VirtualMachine {
	list := make([]hfv1.VirtualMachine, 0)

	for _, v := range req {
		list = append(list, FromVirtualMachineRPC(v))
	}

	return list
}

func ToVirtualMachineListRPC(req []hfv1.VirtualMachine) []*protobuf.VirtualMachine {
	list := make([]*protobuf.VirtualMachine, 0)
	for _, v := range req {
		list = append(list, ToVirtualMachineRPC(v))
	}

	return list
}

func FromVirtualMachineRPC(req *protobuf.VirtualMachine) hfv1.VirtualMachine {
	return hfv1.VirtualMachine{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec:   FromVirtualMachineSpecRPC(req.Spec),
		Status: FromVirtualMachineStatusRPC(req.Status),
	}
}

func ToVirtualMachineRPC(req hfv1.VirtualMachine) *protobuf.VirtualMachine {
	return &protobuf.VirtualMachine{
		Name:   req.ObjectMeta.Name,
		Spec:   ToVirtualMachineSpecRPC(req.Spec),
		Status: ToVirtualMachineStatusRPC(req.Status),
	}
}

func ToVirtualMachineSpecRPC(req hfv1.VirtualMachineSpec) *protobuf.VirtualMachineSpec {
	return &protobuf.VirtualMachineSpec{
		Id:                       req.Id,
		VirtualMachineTemplateId: req.VirtualMachineTemplateId,
		KeyPair:                  req.KeyPair,
		VirtualMachineClaimId:    req.VirtualMachineClaimId,
		UserId:                   req.UserId,
		Provision:                req.Provision,
		VirtualMachineSetId:      req.VirtualMachineSetId,
	}
}

func ToVirtualMachineStatusRPC(req hfv1.VirtualMachineStatus) *protobuf.VirtualMachineStatus {
	return &protobuf.VirtualMachineStatus{
		Status:        ToVMStatusRPC(req.Status),
		Allocated:     req.Allocated,
		Tainted:       req.Tainted,
		PublicIP:      req.PublicIP,
		PrivateIP:     req.PrivateIP,
		EnvironmentId: req.EnvironmentId,
		Hostname:      req.Hostname,
		TFState:       req.TFState,
		WsEndpoint:    req.WsEndpoint,
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
		Status:        FromVMStatusRPC(req.Status),
		Allocated:     req.Allocated,
		Tainted:       req.Tainted,
		PublicIP:      req.PublicIP,
		PrivateIP:     req.PrivateIP,
		EnvironmentId: req.EnvironmentId,
		Hostname:      req.Hostname,
		TFState:       req.TFState,
		WsEndpoint:    req.WsEndpoint,
	}
}

func FromVMStatusRPC(req protobuf.VMStatus) hfv1.VmStatus {
	switch req {
	case protobuf.VMStatus_Provisioned:
		return hfv1.VmStatusProvisioned
	case protobuf.VMStatus_RFP:
		return hfv1.VmStatusRFP
	case protobuf.VMStatus_Running:
		return hfv1.VmStatusRunning
	case protobuf.VMStatus_Terminating:
	default:
		return hfv1.VmStatusTerminating
	}

	panic("shouldn't get here")
}

func ToVMStatusRPC(req hfv1.VmStatus) protobuf.VMStatus {
	switch req {
	case hfv1.VmStatusRunning:
		return protobuf.VMStatus_Running
	case hfv1.VmStatusRFP:
		return protobuf.VMStatus_RFP
	case hfv1.VmStatusProvisioned:
		return protobuf.VMStatus_Provisioned
	case hfv1.VmStatusTerminating:
	default:
		return protobuf.VMStatus_Terminating
	}

	panic("shouldn't get here")
}
