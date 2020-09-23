package converters

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func VirtualMachineSetFromRPC(req *protobuf.VirtualMachineSet) hfv1.VirtualMachineSet {
	return hfv1.VirtualMachineSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec:       VirtualMachineSetSpecFromRPC(req.Spec),
		Status:     VirtualMachineSetStatusFromRPC(req.Status),
	}
}

func VirtualMachineSetToRPC(req hfv1.VirtualMachineSet) *protobuf.VirtualMachineSet {
	return &protobuf.VirtualMachineSet{
		Name:   req.Name,
		Spec:   VirtualMachineSetSpecToRPC(req.Spec),
		Status: VirtualMachineSetStatusToRPC(req.Status),
	}
}

func VirtualMachineSetSpecFromRPC(req *protobuf.VirtualMachineSetSpec) hfv1.VirtualMachineSetSpec {
	return hfv1.VirtualMachineSetSpec{
		Count:               int(req.Count),
		Environment:         req.Environment,
		VMTemplate:          req.VMTemplate,
		BaseName:            req.BaseName,
		RestrictedBind:      req.RestrictedBind,
		RestrictedBindValue: req.RestrictedBindValue,
	}
}

func VirtualMachineSetSpecToRPC(req hfv1.VirtualMachineSetSpec) *protobuf.VirtualMachineSetSpec {
	return &protobuf.VirtualMachineSetSpec{
		Count:               int32(req.Count),
		Environment:         req.Environment,
		VMTemplate:          req.VMTemplate,
		BaseName:            req.BaseName,
		RestrictedBind:      req.RestrictedBind,
		RestrictedBindValue: req.RestrictedBindValue,
	}
}

func VirtualMachineSetStatusFromRPC(req *protobuf.VirtualMachineSetStatus) hfv1.VirtualMachineSetStatus {
	convertedMachines := make([]hfv1.VirtualMachineProvision, len(req.Machines))
	for i, v := range req.Machines {
		newV := VirtualMachineProvisionFromRPC(v)
		convertedMachines[i] = newV
	}
	return hfv1.VirtualMachineSetStatus{
		Machines:         convertedMachines,
		AvailableCount:   int(req.AvailableCount),
		ProvisionedCount: int(req.ProvisionedCount),
	}
}

func VirtualMachineSetStatusToRPC(req hfv1.VirtualMachineSetStatus) *protobuf.VirtualMachineSetStatus {
	convertedMachines := make([]*protobuf.VirtualMachineProvision, len(req.Machines))
	for i, v := range req.Machines {
		newV := VirtualMachineProvisionToRPC(v)
		convertedMachines[i] = newV
	}
	return &protobuf.VirtualMachineSetStatus{
		Machines:         convertedMachines,
		AvailableCount:   int32(req.AvailableCount),
		ProvisionedCount: int32(req.ProvisionedCount),
	}
}

func VirtualMachineProvisionFromRPC(req *protobuf.VirtualMachineProvision) hfv1.VirtualMachineProvision {
	return hfv1.VirtualMachineProvision{
		VirtualMachineName: req.VirtualMachineName,
		TFControllerState:  req.TFControllerState,
		TFControllerCM:     req.TFControllerCM,
	}
}

func VirtualMachineProvisionToRPC(req hfv1.VirtualMachineProvision) *protobuf.VirtualMachineProvision {
	return &protobuf.VirtualMachineProvision{
		VirtualMachineName: req.VirtualMachineName,
		TFControllerState:  req.TFControllerState,
		TFControllerCM:     req.VirtualMachineName,
	}
}