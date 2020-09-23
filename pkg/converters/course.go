package converters

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CourseFromRPC(req *protobuf.Course) hfv1.Course {
	return hfv1.Course{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec:      	CourseSpecFromRPC(req.Spec),
	}
}

func CourseToRPC(req hfv1.Course) *protobuf.Course {
	return &protobuf.Course{
		Name: req.Name,
		Spec: CourseSpecToRPC(req.Spec),
	}
}

func CourseSpecFromRPC(req *protobuf.CourseSpec) hfv1.CourseSpec {
	convertedVms := make([]map[string]string, len(req.VirtualMachines))
	for i, v := range req.VirtualMachines {
		newV := StringMapFromRPC(v)
		convertedVms[i] = newV
	}
	return hfv1.CourseSpec{
		Id:                req.Id,
		Name:              req.Name,
		Description:       req.Description,
		Scenarios:         req.Scenarios,
		VirtualMachines:   convertedVms,
		KeepAliveDuration: req.KeepAliveDuration,
		PauseDuration:     req.PauseDuration,
		Pauseable:         req.Pauseable,
	}
}

func CourseSpecToRPC(req hfv1.CourseSpec) *protobuf.CourseSpec {
	convertedVms := make([]*protobuf.StringMap, len(req.VirtualMachines))
	for i, v := range req.VirtualMachines {
		newV := StringMapToRPC(v)
		convertedVms[i] = newV
	}
	return &protobuf.CourseSpec{
		Id:                req.Id,
		Name:              req.Name,
		Description:       req.Description,
		Scenarios:         req.Scenarios,
		VirtualMachines:   convertedVms,
		KeepAliveDuration: req.KeepAliveDuration,
		PauseDuration:     req.PauseDuration,
		Pauseable:         req.Pauseable,
	}
}