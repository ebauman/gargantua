package courses

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/hobbyfarm/gargantua/pkg/accesscode"
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	apiserverUtil "github.com/hobbyfarm/gargantua/pkg/apiserver/handlers/util"
	"github.com/hobbyfarm/gargantua/pkg/apiserver/stubs"
	"github.com/hobbyfarm/gargantua/pkg/authclient"
	hfClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	hfInformers "github.com/hobbyfarm/gargantua/pkg/client/informers/externalversions"
	util "github.com/hobbyfarm/gargantua/pkg/util"
	"github.com/labstack/echo/v4"
	"k8s.io/client-go/tools/cache"
)

const (
	idIndex = "courseserver.hobbyfarm.io/id-index"
)

type Server struct {
	client *hfClientset.Clientset
	courseIndexer cache.Indexer
	acClient *accesscode.AccessCodeClient
}

func idIndexer(obj interface{}) ([]string, error) {
	course, ok := obj.(*hfv1.Course)
	if !ok {
		return []string{}, nil
	}
	return []string{course.Spec.Id}, nil
}

func NewServer(acClient *accesscode.AccessCodeClient, hfClientset *hfClientset.Clientset, hfInformerFactory hfInformers.SharedInformerFactory) (*CourseServer, error) {
	course := Server{}

	course.client = hfClientset
	course.acClient = acClient
	inf := hfInformerFactory.Hobbyfarm().V1().Courses().Informer()
	indexers := map[string]cache.IndexFunc{idIndex: idIndexer}

	inf.AddIndexers(indexers)
	course.courseIndexer = inf.GetIndexer()

	return &course, nil
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
			courses = append(courses, ToStub(course))
		}
	}

	return ctx.JSON(200, courses)
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

func FromStub(stub stubs.Course) hfv1.Course {
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
		},
	}

	return course
}

func ToStub(course hfv1.Course) stubs.Course {
	outCourse := stubs.Course{
		Name: &course.Name,
		Spec: &stubs.Coursespec{
			Description:       &course.Spec.Description,
			Id:                &course.Spec.Id,
			KeepaliveDuration: &course.Spec.KeepAliveDuration,
			Name:              &course.Spec.Name,
			PauseDuration:     &course.Spec.PauseDuration,
			Pauseable:         &course.Spec.Pauseable,
			Scenarios:         &course.Spec.Scenarios,
			Virtualmachines:   &course.Spec.VirtualMachines,
		},
	}

	return outCourse
}