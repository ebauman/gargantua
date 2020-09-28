package interceptors

import (
	"context"
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/grpc"
)

func MethodInterceptor(methods *map[string]*desc.MethodDescriptor)  func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		mm := *methods
		md := mm[info.FullMethod]

		newCtx := context.WithValue(ctx, "service-descriptor", md) // yes?
		return handler(newCtx, req)
	}
}
