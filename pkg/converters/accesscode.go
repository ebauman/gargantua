package converters

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func AccessCodeFromRPC(req *protobuf.AccessCode) hfv1.AccessCode {
	return hfv1.AccessCode{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec:      	AccessCodeSpecFromRPC(req.Spec),
	}
}

func AccessCodeToRPC(req hfv1.AccessCode) *protobuf.AccessCode {
	return &protobuf.AccessCode{
		Name: req.Name,
		Spec: AccessCodeSpecToRPC(req.Spec),
	}
}

func AccessCodeSpecFromRPC(req *protobuf.AccessCodeSpec) hfv1.AccessCodeSpec {
	return hfv1.AccessCodeSpec{
		Code:                req.Code,
		Description:         req.Description,
		Scenarios:           req.Scenarios,
		Courses:             req.Courses,
		Expiration:          req.Expiration,
		VirtualMachineSets:  req.VirtualMachineSets,
		RestrictedBind:      req.RestrictedBind,
		RestrictedBindValue: req.RestrictedBindValue,
	}
}

func AccessCodeSpecToRPC(req hfv1.AccessCodeSpec) *protobuf.AccessCodeSpec {
	return &protobuf.AccessCodeSpec{
		Code:                req.Code,
		Description:         req.Description,
		Scenarios:           req.Scenarios,
		Courses:             req.Courses,
		Expiration:          req.Expiration,
		VirtualMachineSets:  req.VirtualMachineSets,
		RestrictedBind:      req.RestrictedBind,
		RestrictedBindValue: req.RestrictedBindValue,
	}
}
