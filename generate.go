//go:generate go run k8s.io/kube-openapi/cmd/openapi-gen -i github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1,github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v2,k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/apimachinery/pkg/runtime,k8s.io/apimachinery/pkg/version,k8s.io/apimachinery/pkg/api/resource,k8s.io/api/core/v1,k8s.io/api/rbac/v1,k8s.io/apimachinery/pkg/util/intstr -o ./  -p /pkg/openapi -h hack/boilerplate.go.txt
package main

import (
	_ "k8s.io/kube-openapi/cmd/openapi-gen/args"
)
