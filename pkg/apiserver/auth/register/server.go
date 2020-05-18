package register

import (
	"github.com/gofiber/fiber"
	"github.com/golang/glog"
	"github.com/hobbyfarm/gargantua/pkg/util"
	"strings"
)

func NewServer(app *fiber.App) {

}

func RegisterWithAccessCode(c *fiber.Ctx) {
	email := c.FormValue("email")
	accessCode := strings.ToLower(c.FormValue("access_code"))
	password := c.FormValue("password")

	if len(email) == 0 || len(accessCode) == 0 || len(password) == 0 {

		util.ReturnHTTPMessage(w, r, 400, "error", "invalid input. required fields: email, access_code, password")
		return
	}

	userId, err := a.NewUser(email, password)

	if err != nil {
		glog.Errorf("error creating user %s %v", email, err)
		util.ReturnHTTPMessage(w, r, 400, "error", "error creating user")
		return
	}

	err = a.AddAccessCode(userId, accessCode)

	if err != nil {
		glog.Errorf("error creating user %s %v", email, err)
		util.ReturnHTTPMessage(w, r, 400, "error", "error creating user")
		return
	}

	glog.V(2).Infof("created user %s", email)
	util.ReturnHTTPMessage(w, r, 201, "info", "created user")
}