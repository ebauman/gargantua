package converters

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DynamicBindRequestFromRPC(req *protobuf.DynamicBindRequest) hfv1.DynamicBindRequest {
	return hfv1.DynamicBindRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec:       DynamicBindRequestSpecFromRPC(req.Spec),
		Status:     DynamicBindRequestStatusFromRPC(req.Status),
	}
}

func DynamicBindRequestToRPC(req hfv1.DynamicBindRequest) *protobuf.DynamicBindRequest {
	return &protobuf.DynamicBindRequest{
		Name:   req.Name,
		Spec:   DynamicBindRequestSpecToRPC(req.Spec),
		Status: DynamicBindRequestStatusToRPC(req.Status),
	}
}

func DynamicBindRequestSpecFromRPC(req *protobuf.DynamicBindRequestSpec) hfv1.DynamicBindRequestSpec {
	return hfv1.DynamicBindRequestSpec{
		Id:                  req.Id,
		VirtualMachineClaim: req.VirtualMachineClaim,
		Attempts:            int(req.Attempts),
	}
}

func DynamicBindRequestSpecToRPC(req hfv1.DynamicBindRequestSpec) *protobuf.DynamicBindRequestSpec {
	return &protobuf.DynamicBindRequestSpec{
		Id:                  req.Id,
		VirtualMachineClaim: req.VirtualMachineClaim,
		Attempts:            int32(req.Attempts),
	}
}

func DynamicBindRequestStatusFromRPC(req *protobuf.DynamicBindRequestStatus) hfv1.DynamicBindRequestStatus {
	return hfv1.DynamicBindRequestStatus{
		CurrentAttempts:            int(req.CurrentAttempts),
		Expired:                    req.Expired,
		Fulfilled:                  req.Fulfilled,
		DynamicBindConfigurationId: req.DynamicBindConfigurationId,
		VirtualMachineIds:          req.VirtualMachineIds,
	}
}

func DynamicBindRequestStatusToRPC(req hfv1.DynamicBindRequestStatus) *protobuf.DynamicBindRequestStatus {
	return &protobuf.DynamicBindRequestStatus{
		CurrentAttempts:            int32(req.CurrentAttempts),
		Expired:                    req.Expired,
		Fulfilled:                  req.Fulfilled,
		DynamicBindConfigurationId: req.DynamicBindConfigurationId,
		VirtualMachineIds:          req.VirtualMachineIds,
	}
}