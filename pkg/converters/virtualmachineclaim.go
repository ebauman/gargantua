package converters

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func VirtualMachineClaimFromRPC(req *protobuf.VirtualMachineClaim) hfv1.VirtualMachineClaim {
	return hfv1.VirtualMachineClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec:       VirtualMachineClaimSpecFromRPC(req.Spec),
		Status:     VirtualMachineClaimStatusFromRPC(req.Status),
	}
}

func VirtualMachineClaimToRPC(req hfv1.VirtualMachineClaim) *protobuf.VirtualMachineClaim {
	return &protobuf.VirtualMachineClaim{
		Name:   req.Name,
		Spec:   VirtualMachineClaimSpecToRPC(req.Spec),
		Status: VirtualMachineClaimStatusToRPC(req.Status),
	}
}

func VirtualMachineClaimSpecFromRPC(req *protobuf.VirtualMachineClaimSpec) hfv1.VirtualMachineClaimSpec {
	convertedVms := map[string]hfv1.VirtualMachineClaimVM{}
	for k, v := range req.VirtualMachines {
		newV := VirtualMachineClaimVMFromRPC(v)
		convertedVms[k] = newV
	}

	return hfv1.VirtualMachineClaimSpec{
		Id:                  req.Id,
		UserId:              req.UserId,
		RestrictedBind:      req.RestrictedBind,
		RestrictedBindValue: req.RestrictedBindValue,
		VirtualMachines:     convertedVms,
		DynamicCapable:      req.DynamicCapable,
	}
}

func VirtualMachineClaimSpecToRPC(req hfv1.VirtualMachineClaimSpec) *protobuf.VirtualMachineClaimSpec {
	convertedVms := map[string]*protobuf.VirtualMachineClaimVM{}
	for k, v := range req.VirtualMachines {
		newV := VirtualMachineClaimVMToRPC(v)
		convertedVms[k] = newV
	}

	return &protobuf.VirtualMachineClaimSpec{
		Id:                  req.Id,
		UserId:              req.UserId,
		RestrictedBind:      req.RestrictedBind,
		RestrictedBindValue: req.RestrictedBindValue,
		VirtualMachines:     convertedVms,
		DynamicCapable:      req.DynamicCapable,
	}
}

func VirtualMachineClaimStatusFromRPC(req *protobuf.VirtualMachineClaimStatus) hfv1.VirtualMachineClaimStatus {
	return hfv1.VirtualMachineClaimStatus{
		BindMode:             req.BindMode,
		StaticBindAttempts:   int(req.StaticBindAttempts),
		DynamicBindRequestId: req.DynamicBindRequestId,
		Bound:                req.Bound,
		Ready:                req.Ready,
		Tainted:              req.Tainted,
	}
}

func VirtualMachineClaimStatusToRPC(req hfv1.VirtualMachineClaimStatus) *protobuf.VirtualMachineClaimStatus {
	return &protobuf.VirtualMachineClaimStatus{
		BindMode:             req.BindMode,
		StaticBindAttempts:   int32(req.StaticBindAttempts),
		DynamicBindRequestId: req.DynamicBindRequestId,
		Bound:                req.Bound,
		Ready:                req.Ready,
		Tainted:              req.Tainted,
	}
}

func VirtualMachineClaimVMFromRPC(req *protobuf.VirtualMachineClaimVM) hfv1.VirtualMachineClaimVM {
	return hfv1.VirtualMachineClaimVM{
		Template:         req.Template,
		VirtualMachineId: req.VirtualMachineId,
	}
}

func VirtualMachineClaimVMToRPC(req hfv1.VirtualMachineClaimVM) *protobuf.VirtualMachineClaimVM {
	return &protobuf.VirtualMachineClaimVM{
		Template:         req.Template,
		VirtualMachineId: req.VirtualMachineId,
	}
}