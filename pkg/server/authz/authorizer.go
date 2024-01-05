package authz

import (
	"context"
	"k8s.io/apiserver/pkg/authorization/authorizer"
)

type AuthZ struct{}

func (az AuthZ) Authorize(ctx context.Context, a authorizer.Attributes) (authorized authorizer.Decision, reason string, err error) {
	//if !onlyApis(a) {
	//	return authorizer.DecisionDeny, "invalid path", nil
	//}

	if startsWithAny(a.GetPath(), "/healthz", "/livez", "/metrics") {
		return authorizer.DecisionDeny, "path not supported", nil
	}

	return authorizer.DecisionAllow, "", nil
}

func onlyApis(a authorizer.Attributes) bool {
	if len(a.GetPath()) < 5 {
		return false
	}

	if a.GetPath()[0:5] != "/apis" {
		return false
	}

	return true
}

func startsWithAny(inspect string, values ...string) bool {
	for _, v := range values {
		if startsWith(inspect, v) {
			return true
		}
	}

	return false
}

func startsWith(inspect string, value string) bool {
	if len(inspect) < len(value) {
		return false
	}

	if inspect[0:len(value)] != value {
		return false
	}

	return true
}
