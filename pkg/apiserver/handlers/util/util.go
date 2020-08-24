package util

import (
	"fmt"
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/labstack/echo/v4"
)

func GetUserFromContext(ctx echo.Context) (hfv1.User, error) {
	user := ctx.Get("user")
	if user == nil {
		return hfv1.User{}, fmt.Errorf("unable to retrieve user from context")
	}

	var usr hfv1.User
	usr, ok := user.(hfv1.User)
	if !ok {
		return hfv1.User{}, fmt.Errorf("unable to assert user from context")
	}

	return usr, nil
}
