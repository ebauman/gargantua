package converters

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func SessionFromRPC(req *protobuf.Session) hfv1.Session {
	return hfv1.Session{
		ObjectMeta: metav1.ObjectMeta{
			Name: req.Name,
		},
		Spec:       SessionSpecFromRPC(req.Spec),
		Status:     SessionStatusFromRPC(req.Status),
	}
}

func SessionToRPC(req hfv1.Session) *protobuf.Session {
	return &protobuf.Session{
		Name:   req.Name,
		Spec:   SessionSpecToRPC(req.Spec),
		Status: SessionStatusToRPC(req.Status),
	}
}

func SessionSpecFromRPC(req *protobuf.SessionSpec) hfv1.SessionSpec {
	return hfv1.SessionSpec{
		Id:         req.Id,
		ScenarioId: req.ScenarioId,
		CourseId:   req.CourseId,
		UserId:     req.UserId,
		VmClaimSet: req.VmClaimSet,
		AccessCode: req.AccessCode,
	}
}

func SessionSpecToRPC(req hfv1.SessionSpec) *protobuf.SessionSpec {
	return &protobuf.SessionSpec{
		Id:         req.Id,
		ScenarioId: req.ScenarioId,
		CourseId:   req.CourseId,
		UserId:     req.UserId,
		VmClaimSet: req.VmClaimSet,
		AccessCode: req.AccessCode,
	}
}

func SessionStatusFromRPC(req *protobuf.SessionStatus) hfv1.SessionStatus {
	return hfv1.SessionStatus{
		Paused:         req.Paused,
		PausedTime:     req.PausedTime,
		Active:         req.Active,
		Finished:       req.Finished,
		StartTime:      req.StartTime,
		ExpirationTime: req.ExpirationTime,
	}
}

func SessionStatusToRPC(req hfv1.SessionStatus) *protobuf.SessionStatus {
	return &protobuf.SessionStatus{
		Paused:         req.Paused,
		PausedTime:     req.PausedTime,
		Active:         req.Active,
		Finished:       req.Finished,
		StartTime:      req.StartTime,
		ExpirationTime: req.ExpirationTime,
	}
}

