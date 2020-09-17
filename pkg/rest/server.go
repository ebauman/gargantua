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
	err := protobuf.RegisterVirtualMachineServiceHandlerFromEndpoint(ctx, mux, httpEndpoint, opts)
	if err != nil{
		return err
	}

	return http.ListenAndServe(httpEndpoint, mux)
}
