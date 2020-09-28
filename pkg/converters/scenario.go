package converters

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ScenarioFromRPC(req *protobuf.Scenario) hfv1.Scenario {
	return hfv1.Scenario{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec:       ScenarioSpecFromRPC(req.Spec),
	}
}

func ScenarioToRPC(req hfv1.Scenario) *protobuf.Scenario {
	return &protobuf.Scenario{
		Name: req.Name,
		Spec: ScenarioSpecToRPC(req.Spec),
	}
}

func ScenarioSpecFromRPC(req *protobuf.ScenarioSpec) hfv1.ScenarioSpec {
	convertedSteps := make([]hfv1.ScenarioStep, len(req.Steps))
	for i, v := range req.Steps {
		newV := ScenarioStepFromRPC(v)
		convertedSteps[i] = newV
	}

	convertedVms := make([]map[string]string, len(req.VirtualMachines))
	for i, v := range req.VirtualMachines {
		newV := StringMapFromRPC(v)
		convertedVms[i] = newV
	}
	return hfv1.ScenarioSpec{
		Id:                req.Id,
		Name:              req.Name,
		Description:       req.Description,
		Steps:            	convertedSteps,
		VirtualMachines:   convertedVms,
		KeepAliveDuration: req.KeepAliveDuration,
		PauseDuration:     req.PauseDuration,
		Pauseable:         req.Pauseable,
	}
}

func ScenarioSpecToRPC(req hfv1.ScenarioSpec) *protobuf.ScenarioSpec {
	convertedSteps := make([]*protobuf.ScenarioStep, len(req.Steps))
	for i, v := range req.Steps {
		newV := ScenarioStepToRPC(v)
		convertedSteps[i] = newV
	}

	convertedVms := make([]*protobuf.StringMap, len(req.VirtualMachines))
	for i, v := range req.VirtualMachines {
		newV := StringMapToRPC(v)
		convertedVms[i] = newV
	}
	return &protobuf.ScenarioSpec{
		Id:                req.Id,
		Name:              req.Name,
		Description:       req.Description,
		Steps:             convertedSteps,
		VirtualMachines:   convertedVms,
		KeepAliveDuration: req.KeepAliveDuration,
		PauseDuration:     req.PauseDuration,
		Pauseable:         req.Pauseable,
	}
}

func ScenarioStepFromRPC(req *protobuf.ScenarioStep) hfv1.ScenarioStep {
	return hfv1.ScenarioStep{
		Title:   req.Title,
		Content: req.Content,
	}
}

func ScenarioStepToRPC(req hfv1.ScenarioStep) *protobuf.ScenarioStep {
	return &protobuf.ScenarioStep{
		Title:   req.Title,
		Content: req.Content,
	}
}
