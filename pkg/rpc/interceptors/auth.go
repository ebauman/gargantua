package interceptors

import (
	"context"
	"errors"
	"fmt"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func UnaryAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	isAdmin, err := retrieveIsAdmin(ctx)
	if err != nil {
		return nil, err
	}



	return handler(ctx, req)
}

func retrieveIsAdmin(ctx context.Context) (bool, error) {
	sd := ctx.Value("service-descriptor")
	methodDescriptor, ok := sd.(*desc.MethodDescriptor)
	if !ok {
		return false, errors.New("error pulling service descriptor from rpc context")
	}

	opts := methodDescriptor.GetMethodOptions()

	proto.HasExtension(opts, protobuf.E_RequiresAdmin)
	val := proto.GetExtension(opts, protobuf.E_RequiresAdmin)
	switch v := val.(type) {
	case bool:
		return v, nil
	default:
		return false, errors.New("retrieved non-bool value from method extension")
	}
}