package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hobbyfarm/gargantua/pkg/protobuf"
	"google.golang.org/grpc"
	"net/http"
	"sync"
)

func Serve(ctx context.Context, wg *sync.WaitGroup, grpcEndpoint string, httpEndpoint string) error {
	defer wg.Done()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	registerFuncs := []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error {
		protobuf.RegisterAccessCodeServiceHandlerFromEndpoint,
		protobuf.RegisterCourseServiceHandlerFromEndpoint,
		protobuf.RegisterDynamicBindConfigurationServiceHandlerFromEndpoint,
		protobuf.RegisterDynamicBindRequestServiceHandlerFromEndpoint,
		protobuf.RegisterEnvironmentServiceHandlerFromEndpoint,
		protobuf.RegisterScenarioServiceHandlerFromEndpoint,
		protobuf.RegisterScheduledEventServiceHandlerFromEndpoint,
		protobuf.RegisterSessionServiceHandlerFromEndpoint,
		protobuf.RegisterUserServiceHandlerFromEndpoint,
		protobuf.RegisterVirtualMachineServiceHandlerFromEndpoint,
		protobuf.RegisterVirtualMachineClaimServiceHandlerFromEndpoint,
		protobuf.RegisterVirtualMachineSetServiceHandlerFromEndpoint,
		protobuf.RegisterVirtualMachineTemplateServiceHandlerFromEndpoint,
	}

	for _, f := range registerFuncs {
		err := f(ctx, mux, grpcEndpoint, opts)
		if err != nil {
			return err
		}
	}

	return http.ListenAndServe(httpEndpoint, mux)
}
