package courses

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/hobbyfarm/gargantua/pkg/accesscode"
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	apiserverUtil "github.com/hobbyfarm/gargantua/pkg/apiserver/handlers/util"
	"github.com/hobbyfarm/gargantua/pkg/apiserver/stubs"
	hfClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	util "github.com/hobbyfarm/gargantua/pkg/util"
	"github.com/labstack/echo/v4"
	"k8s.io/client-go/tools/cache"
)

const (
	idIndex = "courseserver.hobbyfarm.io/id-index"
)

type Server struct {
	client hfClientset.Clientset
	courseIndexer cache.Indexer
	acClient *accesscode.AccessCodeClient
}

func (s *Server) GetCourses(ctx echo.Context) error {
	user, err := apiserverUtil.GetUserFromContext(ctx)
	if err != nil {
		return err
	}

	var courseIds []string
	for _, ac := range user.Spec.AccessCodes {
		tempCourseIds, err := s.acClient.GetCourseIds(ac)
		if err != nil {
			glog.Errorf("error retrieving course ids for access code %s: %v", ac, err)
		} else {
			courseIds = append(courseIds, tempCourseIds...)
		}
	}

	courseIds = util.UniqueStringSlice(courseIds)

	var courses []stubs.Course
	for _, courseId := range courseIds {
		course, err := s.getCourseById(courseId)
		if err != nil {
			glog.Errorf("error retrieving course with id %s: %v", courseId, err)
		} else {
			courses = append(courses, course)
		}
	}
}

func (s *Server) getCourseById(id string) (hfv1.Course, error) {
	if len(id) == 0 {
		return hfv1.Course{}, fmt.Errorf("course id passed in was blank")
	}
	obj, err := s.courseIndexer.ByIndex(idIndex, id)

	if err != nil {
		return hfv1.Course{}, fmt.Errorf("error while retrieving course by ID %s %v", id, err)
	}

	if len(obj) < 1 {
		return hfv1.Course{}, fmt.Errorf("error while retrieving course by ID %s", id)
	}

	course, ok := obj[0].(*hfv1.Course)

	if !ok {
		return hfv1.Course{}, fmt.Errorf("error while retrieving course by ID %s %v", id, ok)
	}

	return *course, nil
}

func FromStub(stub stubs.Course) (hfv1.Course, error) {
	course := hfv1.Course{
		Spec: hfv1.CourseSpec{
			Description: *stub.Spec.Description,
			Id: *stub.Spec.Id,
			KeepAliveDuration: *stub.Spec.KeepaliveDuration,
			Name: *stub.Spec.Name,
			PauseDuration: *stub.Spec.PauseDuration,
			Pauseable: *stub.Spec.Pauseable,
			Scenarios: *stub.Spec.Scenarios,
			VirtualMachines: *stub.Spec.Virtualmachines,
		}
	}
}

// how do we handle the conversion from an openapi stub type to an hfv1 type?
// the maps look troublesome, the spec has them as *[]struct { AdditionalProperties map[string]string }}
// which does not easily translate.
// Can we do this automatically somehow?
// Is this better off just being done manually on each method or something?

func ToStub() {

}