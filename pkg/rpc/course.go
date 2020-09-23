package rpc

import (
	"context"
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	hfClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	hfInformers "github.com/hobbyfarm/gargantua/pkg/client/informers/externalversions"
	"github.com/hobbyfarm/gargantua/pkg/converters"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

const (
	courseIdIndex = "courseserver.hobbyfarm.io/id-index"
)

type CourseServer struct {
	hfClientSet *hfClientset.Clientset
	courseIndexer cache.Indexer
}

func setupCourseServer(g *grpc.Server, clientset *hfClientset.Clientset, factory hfInformers.SharedInformerFactory) {
	cs := CourseServer{}

	cs.hfClientSet = clientset
	inf := factory.Hobbyfarm().V1().Courses().Informer()
	indexers := map[string]cache.IndexFunc{courseIdIndex: courseIdIndexer}

	inf.AddIndexers(indexers)
	cs.courseIndexer = inf.GetIndexer()

	protobuf.RegisterCourseServiceServer(g, cs)
}

func (c CourseServer) List(ctx context.Context, empty *protobuf.Empty) (*protobuf.CourseList, error) {
	list, err := c.hfClientSet.HobbyfarmV1().Courses().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	out := make([]*protobuf.Course, len(list.Items))

	for i, v := range list.Items {
		out[i] = converters.CourseToRPC(v)
	}

	return &protobuf.CourseList{Courses: out}, nil
}

func (c CourseServer) Get(ctx context.Context, id *protobuf.ID) (*protobuf.Course, error) {
	obj, err := c.courseIndexer.ByIndex(courseIdIndex, id.ID)
	if err != nil {
		return nil, err
	}

	if len(obj) < 1 {
		return nil, status.Errorf(codes.NotFound, "cannot locate course with id %s", id.ID)
	}

	course, ok := obj[0].(*hfv1.Course)
	if !ok {
		return nil, status.Error(codes.Internal, "error asserting course into hfv1.Course")
	}

	return converters.CourseToRPC(*course), nil
}

func (c CourseServer) Create(ctx context.Context, course *protobuf.Course) (*protobuf.Course, error) {
	toCreate := converters.CourseFromRPC(course)
	res, err := c.hfClientSet.HobbyfarmV1().Courses().Create(&toCreate)
	if err != nil {
		return nil, err
	}

	return converters.CourseToRPC(*res), nil
}

func (c CourseServer) Update(ctx context.Context, course *protobuf.Course) (*protobuf.Course, error) {
	toUpdate := converters.CourseFromRPC(course)
	res, err := c.hfClientSet.HobbyfarmV1().Courses().Update(&toUpdate)
	if err != nil {
		return nil, err
	}

	return converters.CourseToRPC(*res), nil
}

func (c CourseServer) Delete(ctx context.Context, id *protobuf.ID) (*protobuf.Empty, error) {
	err := c.hfClientSet.HobbyfarmV1().Courses().Delete(id.ID, &metav1.DeleteOptions{})
	if err != nil {
		return nil, err
	}

	return &protobuf.Empty{}, nil
}

func courseIdIndexer(obj interface{}) ([]string, error) {
	course, ok := obj.(*hfv1.Course)
	if !ok {
		return []string{}, nil
	}
	return []string{course.Spec.Id}, nil
}