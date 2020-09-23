package converters

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ScheduledEventFromRPC(req *protobuf.ScheduledEvent) hfv1.ScheduledEvent {
	return hfv1.ScheduledEvent{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec:       ScheduledEventSpecFromRPC(req.Spec),
		Status:     ScheduledEventStatusFromRPC(req.Status),
	}
}

func ScheduledEventToRPC(req hfv1.ScheduledEvent) *protobuf.ScheduledEvent {
	return &protobuf.ScheduledEvent{
		Name:   req.Name,
		Spec:   ScheduledEventSpecToRPC(req.Spec),
		Status: ScheduledEventStatusToRPC(req.Status),
	}
}

func ScheduledEventSpecFromRPC(req *protobuf.ScheduledEventSpec) hfv1.ScheduledEventSpec {
	convertedVms := map[string]map[string]int{}

	for k, v := range req.RequiredVirtualMachines {
		convertedVms[k] = IntMapFromRPC(v)
	}
	return hfv1.ScheduledEventSpec{
		Creator:                 req.Creator,
		Name:                    req.Name,
		Description:             req.Description,
		StartTime:               req.StartTime,
		EndTime:                 req.EndTime,
		RequiredVirtualMachines: convertedVms,
		AccessCode:              req.AccessCode,
		RestrictedBind:          req.RestrictedBind,
		RestrictedBindValue:     req.RestrictedBindValue,
		Scenarios:               req.Scenarios,
		Courses:                 req.Courses,
	}
}

func ScheduledEventSpecToRPC(req hfv1.ScheduledEventSpec) *protobuf.ScheduledEventSpec {
	convertedVms := map[string]*protobuf.Int32Map{}
	for k, v := range req.RequiredVirtualMachines {
		convertedVms[k] = IntMapToRPC(v)
	}
	
	return &protobuf.ScheduledEventSpec{
		Creator:                 req.Creator,
		Name:                    req.Name,
		Description:             req.Description,
		StartTime:               req.StartTime,
		EndTime:                 req.EndTime,
		RequiredVirtualMachines: convertedVms,
		AccessCode:              req.AccessCode,
		RestrictedBind:          req.RestrictedBind,
		RestrictedBindValue:     req.RestrictedBindValue,
		Scenarios:               req.Scenarios,
		Courses:                 req.Courses,
	}
}

func ScheduledEventStatusFromRPC(req *protobuf.ScheduledEventStatus) hfv1.ScheduledEventStatus {
	return hfv1.ScheduledEventStatus{
		AccessCodeId:       req.AccessCodeId,
		VirtualMachineSets: req.VirtualMachineSets,
		Active:             req.Active,
		Provisioned:        req.Provisioned,
		Ready:              req.Ready,
		Finished:           req.Finished,
	}
}

func ScheduledEventStatusToRPC(req hfv1.ScheduledEventStatus) *protobuf.ScheduledEventStatus {
	return &protobuf.ScheduledEventStatus{
		AccessCodeId:       req.AccessCodeId,
		VirtualMachineSets: req.VirtualMachineSets,
		Active:             req.Active,
		Provisioned:        req.Provisioned,
		Ready:              req.Ready,
		Finished:           req.Finished,
	}
}
