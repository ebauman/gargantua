module github.com/hobbyfarm/gargantua

go 1.13

replace k8s.io/client-go => k8s.io/client-go v0.15.8

require (
	github.com/dgrijalva/jwt-go v3.2.1-0.20200107013213-dc14462fd587+incompatible
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/fatih/color v1.9.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/handlers v1.4.0
	github.com/gorilla/mux v1.7.1
	github.com/gorilla/websocket v1.4.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.8
	github.com/iancoleman/strcase v0.1.2 // indirect
	github.com/jhump/protoreflect v1.7.0
	github.com/lyft/protoc-gen-star v0.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rancher/terraform-controller v0.0.10-alpha1
	github.com/rancher/wrangler v0.1.0
	golang.org/x/crypto v0.0.0-20191227163750-53104e6ec876
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.31.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/src-d/go-parse-utils.v1 v1.1.2 // indirect
	gopkg.in/src-d/proteus.v1 v1.3.3 // indirect
	gopkg.in/urfave/cli.v1 v1.20.0 // indirect
	k8s.io/api v0.15.8
	k8s.io/apimachinery v0.15.8
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/code-generator v0.15.8
	k8s.io/klog v1.0.0
)
