package converters

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DynamicBindConfigurationFromRPC(req *protobuf.DynamicBindConfiguration) hfv1.DynamicBindConfiguration {
	return hfv1.DynamicBindConfiguration{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec:       DynamicBindConfigurationSpecFromRPC(req.Spec),
	}
}

func DynamicBindConfigurationToRPC(req hfv1.DynamicBindConfiguration) *protobuf.DynamicBindConfiguration {
	return &protobuf.DynamicBindConfiguration{
		Name: req.Name,
		Spec: DynamicBindConfigurationSpecToRPC(req.Spec),
	}
}

func DynamicBindConfigurationSpecFromRPC(req *protobuf.DynamicBindConfigurationSpec) hfv1.DynamicBindConfigurationSpec {
	convertedCapacity := map[string]int{}
	for k, v := range req.BurstCountCapacity {
		convertedCapacity[k] = int(v)
	}
	return hfv1.DynamicBindConfigurationSpec{
		Id:                  req.Id,
		Environment:         req.Environment,
		BaseName:            req.BaseName,
		RestrictedBind:      req.RestrictedBind,
		RestrictedBindValue: req.RestrictedBindValue,
		BurstCountCapacity:  convertedCapacity,
		BurstCapacity:       CMSStructFromRPC(req.BurstCapacity),
	}
}

func DynamicBindConfigurationSpecToRPC(req hfv1.DynamicBindConfigurationSpec) *protobuf.DynamicBindConfigurationSpec {
	convertedCapacity := map[string]int32{}
	for k, v := range req.BurstCountCapacity {
		convertedCapacity[k] = int32(v)
	}
	return &protobuf.DynamicBindConfigurationSpec{
		Id:                  req.Id,
		Environment:         req.Environment,
		BaseName:            req.BaseName,
		RestrictedBind:      req.RestrictedBind,
		RestrictedBindValue: req.RestrictedBindValue,
		BurstCountCapacity:  convertedCapacity,
		BurstCapacity:       CMSStructToRPC(req.BurstCapacity),
	}
}