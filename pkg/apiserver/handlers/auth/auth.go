package auth

import "github.com/labstack/echo/v4"

func MiddlwareDispatch(ctx echo.Context, provider string, scopes []string) error {
	switch provider {
	case "BearerAuth":
		return HandleBearerAuth()
	}

	return nil
}

func HandleBearerAuth(ctx echo.Context, scopes []string) error {
	// we're going to abuse the scopes functionality to provide
	// "roles" that a user should have to perform a certain action

	// this normally should be handled by the Echo JWT middleware (https://echo.labstack.com/middleware/jwt)


}