package interceptors

import (
	"context"
	"errors"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/hobbyfarm/gargantua/pkg/auth"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"strings"
)

type AuthInterceptor struct {
	ac *auth.AuthClient
}

func NewAuthInterceptor(ac *auth.AuthClient) *AuthInterceptor {
	ai := &AuthInterceptor{}

	ai.ac = ac

	return ai
}

func (a *AuthInterceptor) GetInterceptor() func (ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func (ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		bypass, err := bypassAuth(ctx) // defined for certain functions like registration, login, etc.
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error determining bypass status for endpoint: %s", err)
		}
		if bypass {
			return handler(ctx, req) // continue on our merry way!
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.InvalidArgument, "missing metadata")
		}

		token, err := a.parseToken(md["authorization"])
		if err != nil {
			return nil, err
		}

		requiresAdmin, err := retrieveRequiresAdmin(ctx)
		if err != nil {
			return nil, err
		}

		user, err := a.ac.AuthN(token, requiresAdmin)
		if err != nil {
			return nil, status.Error(codes.PermissionDenied, "not authorized")
		}

		// set the user context
		metadata.AppendToOutgoingContext(ctx, "user", user.Name)

		return handler(ctx, req)
	}
}

func (a *AuthInterceptor) parseToken(authorization []string) (string, error) {
	if len(authorization) < 1 {
		return "", status.Error(codes.InvalidArgument, "invalid auth header")
	}

	tokenString := strings.TrimPrefix(authorization[0], "Bearer ")

	return tokenString, nil
}

func getMethodOptions(ctx context.Context) (*descriptor.MethodOptions, error) {
	sd := ctx.Value("service-descriptor")
	methodDescriptor, ok := sd.(*desc.MethodDescriptor)
	if !ok {
		return nil, errors.New("error pulling service descriptor from rpc context")
	}

	return methodDescriptor.GetMethodOptions(), nil
}

func coerceBool(val interface{}) (bool, error) {
	switch v := val.(type) {
	case bool:
		return v, nil
	default:
		return false, errors.New("could not coerce bool from interface{}")
	}
}

func retrieveRequiresAdmin(ctx context.Context) (bool, error) {
	opts, err := getMethodOptions(ctx)
	if err != nil {
		return false, err
	}

	proto.HasExtension(opts, protobuf.E_RequiresAdmin)
	val := proto.GetExtension(opts, protobuf.E_RequiresAdmin)
	return coerceBool(val)
}

func bypassAuth(ctx context.Context) (bool, error) {
	opts, err := getMethodOptions(ctx)
	if err != nil {
		return false, err
	}

	proto.HasExtension(opts, protobuf.E_BypassAuth)
	val := proto.GetExtension(opts, protobuf.E_BypassAuth)
	return coerceBool(val)
}