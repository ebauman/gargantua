package converters

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func EnvironmentFromRPC(req *protobuf.Environment) hfv1.Environment {
	return hfv1.Environment{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec:       EnvironmentSpecFromRPC(req.Spec),
		Status:     EnvironmentStatusFromRPC(req.Status),
	}
}

func EnvironmentToRPC(req hfv1.Environment) *protobuf.Environment {
	return &protobuf.Environment{
		Name:   req.Name,
		Spec:   EnvironmentSpecToRPC(req.Spec),
		Status: EnvironmentStatusToRPC(req.Status),
	}
}

func EnvironmentSpecFromRPC(req *protobuf.EnvironmentSpec) hfv1.EnvironmentSpec {
	convertedTemplates := map[string]map[string]string{} // lol
	for k, v := range req.TemplateMapping {
		newV := StringMapFromRPC(v)
		convertedTemplates[k] = newV
	}



	return hfv1.EnvironmentSpec{
		DisplayName:          req.DisplayName,
		DNSSuffix:            req.DNSSuffix,
		Provider:             req.Provider,
		TemplateMapping:      convertedTemplates,
		EnvironmentSpecifics: req.EnvironmentSpecifics,
		IPTranslationMap:     req.IPTranslationMap,
		WsEndpoint:           req.WsEndpoint,
		CapacityMode:         CapacityModeFromRPC(req.CapacityMode),
		BurstCapable:         req.BurstCapable,
		CountCapacity:        Int32MapFromRPC(req.CountCapacity),
		Capacity:             CMSStructFromRPC(req.Capacity),
		BurstCountCapacity:   Int32MapFromRPC(req.BurstCountCapacity),
		BurstCapacity:        CMSStructFromRPC(req.BurstCapacity),
	}
}

func EnvironmentSpecToRPC(req hfv1.EnvironmentSpec) *protobuf.EnvironmentSpec {
	convertedTemplates := map[string]*protobuf.StringMap{}
	for k, v := range req.TemplateMapping {
		newV := StringMapToRPC(v)
		convertedTemplates[k] = newV
	}
	
	return &protobuf.EnvironmentSpec{
		DisplayName:          req.DisplayName,
		DNSSuffix:            req.DNSSuffix,
		Provider:             req.Provider,
		TemplateMapping:      convertedTemplates,
		EnvironmentSpecifics: req.EnvironmentSpecifics,
		IPTranslationMap:     req.IPTranslationMap,
		WsEndpoint:           req.WsEndpoint,
		CapacityMode:         CapacityModeToRPC(req.CapacityMode),
		BurstCapable:         req.BurstCapable,
		CountCapacity:        Int32MapToRPC(req.CountCapacity),
		Capacity:             CMSStructToRPC(req.Capacity),
		BurstCountCapacity:   Int32MapToRPC(req.BurstCountCapacity),
		BurstCapacity:        CMSStructToRPC(req.BurstCapacity),
	}
}

func EnvironmentStatusFromRPC(req *protobuf.EnvironmentStatus) hfv1.EnvironmentStatus {
	return hfv1.EnvironmentStatus{
		Used:           CMSStructFromRPC(req.Used),
		AvailableCount: Int32MapFromRPC(req.AvailableCount),
	}
}

func EnvironmentStatusToRPC(req hfv1.EnvironmentStatus) *protobuf.EnvironmentStatus {
	return &protobuf.EnvironmentStatus{
		Used:           CMSStructToRPC(req.Used),
		AvailableCount: Int32MapToRPC(req.AvailableCount),
	}
}

func CapacityModeFromRPC(req protobuf.CapacityMode) hfv1.CapacityMode {
	switch req {
	case protobuf.CapacityMode_Count:
		return hfv1.CapacityModeCount
	case protobuf.CapacityMode_Raw:
	default:
		return hfv1.CapacityModeRaw
	}

	panic("shouldn't get here")
}

func CapacityModeToRPC(req hfv1.CapacityMode) protobuf.CapacityMode {
	switch req {
	case hfv1.CapacityModeRaw:
		return protobuf.CapacityMode_Raw
	case hfv1.CapacityModeCount:
	default:
		return protobuf.CapacityMode_Count
	}

	panic("shouldn't get here")
}