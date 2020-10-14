package rpc

import (
	"context"
	"fmt"
)

func getUser(ctx context.Context) (string, error) {
	user := fmt.Sprintf("%s", ctx.Value("user"))
	if len(user) < 1 {
		return "", fmt.Errorf("error retrieving user from rpc context")
	}

	return user, nil
}