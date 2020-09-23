package converters

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func UserFromRPC(req *protobuf.User) hfv1.User {
	return hfv1.User{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec:       UserSpecFromRPC(req.Spec),
	}
}

func UserToRPC(req hfv1.User) *protobuf.User {
	return &protobuf.User{
		Name: req.Name,
		Spec: UserSpecToRPC(req.Spec),
	}
}

func UserSpecFromRPC(req *protobuf.UserSpec) hfv1.UserSpec {
	return hfv1.UserSpec{
		Id:          req.Id,
		Email:       req.Email,
		Password:    req.Password,
		AccessCodes: req.AccessCodes,
		Admin:       req.Admin,
	}
}

func UserSpecToRPC(req hfv1.UserSpec) *protobuf.UserSpec {
	return &protobuf.UserSpec{
		Id:          req.Id,
		Email:       req.Email,
		Password:    req.Password,
		AccessCodes: req.AccessCodes,
		Admin:       req.Admin,
	}
}