package converters

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func VirtualMachineTemplateFromRPC(req *protobuf.VirtualMachineTemplate) hfv1.VirtualMachineTemplate {
	return hfv1.VirtualMachineTemplate{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec:       VirtualMachineTemplateSpecFromRPC(req.Spec),
	}
}

func VirtualMachineTemplateToRPC(req hfv1.VirtualMachineTemplate) *protobuf.VirtualMachineTemplate {
	return &protobuf.VirtualMachineTemplate{
		Name: req.Name,
		Spec: VirtualMachineTemplateSpecToRPC(req.Spec),
	}
}

func VirtualMachineTemplateSpecFromRPC(req *protobuf.VirtualMachineTemplateSpec) hfv1.VirtualMachineTemplateSpec {
	return hfv1.VirtualMachineTemplateSpec{
		Id:        req.Id,
		Name:      req.Name,
		Image:     req.Image,
		Resources: CMSStructFromRPC(req.Resources),
		CountMap:  req.CountMap,
	}
}

func VirtualMachineTemplateSpecToRPC(req hfv1.VirtualMachineTemplateSpec) *protobuf.VirtualMachineTemplateSpec {
	return &protobuf.VirtualMachineTemplateSpec{
		Id:        req.Id,
		Name:      req.Name,
		Image:     req.Image,
		Resources: CMSStructToRPC(req.Resources),
		CountMap:  req.CountMap,
	}
}