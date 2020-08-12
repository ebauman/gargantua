package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"github.com/hobbyfarm/gargantua/pkg/apiserver/auth"
	"github.com/hobbyfarm/gargantua/pkg/apiserver/auth/base"
	"github.com/hobbyfarm/gargantua/pkg/apiserver/errors"
	"github.com/hobbyfarm/gargantua/pkg/apiserver/utils"
	"strconv"
	"strings"
)

type APIServer struct {
	app *fiber.App
	router fiber.Router
	validator *validator.Validate
	authNProviders map[auth.AuthNProvider]func() fiber.Handler
	authZProviders map[auth.AuthZProvider]func() fiber.Handler
}

type APIServerSettings struct {
	PathPrefix []string
	FiberSettings *fiber.Settings
	AuthServers []auth.AuthServer
}

type Context struct {
	Fiber *fiber.Ctx
	Validator *validator.Validate
	Locals map[string]interface{}
}

type HTTPContent struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	Content []byte `json:"content"`
}

type HTTPMessage struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func New(set *APIServerSettings) *APIServer {
	a := &APIServer{}

	if set.FiberSettings.ErrorHandler == nil {
		set.FiberSettings.ErrorHandler = handleError
	}

	a.app = fiber.New(set.FiberSettings)
	a.validator = validator.New()

	if len(set.PathPrefix) > 0 {
		path := strings.Join(set.PathPrefix, "/")
		a.router = a.app.Group(path)
	} else {
		a.router = a.app.Group("")
	}

	d, _ := base.New()
	set.AuthServers = append(set.AuthServers, d)

	a.authNProviders = make(map[auth.AuthNProvider]func() fiber.Handler)
	a.authZProviders = make(map[auth.AuthZProvider]func() fiber.Handler)

	for _, s := range set.AuthServers {
		utils.MergeAuthNMaps(a.authNProviders, s.AuthNProviders())
		utils.MergeAuthZMaps(a.authZProviders, s.AuthZProviders())
	}

	return a
}


func (a *APIServer) Listen(address interface{}) {
	a.app.Listen(address)
}

func (a *APIServer) BuildContext (ctx *fiber.Ctx) (*Context, error) {
	// given a fiber context, build an api context
	c := &Context{
		Fiber: ctx,
		Locals: map[string]interface{}{},
		Validator: a.validator,
	}

	return c, nil
}

func (a *APIServer) auth(path string, n string, z string) {
	var nProv auth.AuthNProvider
	var zProv auth.AuthZProvider

	if n == "" {
		n = "nil"
	}

	if z == "" {
		z = "nil"
	}

	nProv = auth.AuthNProvider(n)
	zProv = auth.AuthZProvider(z)

	if _, ok := a.authNProviders[nProv]; !ok {
		a.app.Use(path, func(ctx *fiber.Ctx) {
			ctx.Next(fmt.Errorf("invalid authn provider: %s", nProv))
		})
		return // don't wire auth
	}

	if _, ok := a.authZProviders[zProv]; !ok {
		a.app.Use(path, func(ctx *fiber.Ctx) {
			ctx.Next(fmt.Errorf("invalid authz provider: %s", zProv))
		})
		return // don't wire auth
	}

	a.app.Use(path, a.authNProviders[nProv]())
	a.app.Use(path, a.authZProviders[zProv]())
}

func (a *APIServer) Get(path string, f func(*Context) (interface{}, error), n string, z string) {
	a.auth(path, n, z)

	a.router.Get(path, func (c *fiber.Ctx) {
		a.wireContextAndHandle(c, f)
	})
}

func (a *APIServer) Post(path string, f func(*Context) (interface{}, error), n string, z string)  {
	a.auth(path, n, z)

	a.router.Post(path, func (c* fiber.Ctx) {
		a.wireContextAndHandle(c, f)
	})
}

func (a *APIServer) Put(path string, f func(*Context) (interface{}, error), n string, z string) {
	a.auth(path, n, z)

	a.router.Put(path, func (c* fiber.Ctx) {
		a.wireContextAndHandle(c, f)
	})
}

func (a *APIServer) Delete(path string, f func(*Context) (interface{}, error), n string, z string) {
	a.auth(path, n, z)

	a.router.Delete(path, func (c* fiber.Ctx) {
		a.wireContextAndHandle(c, f)
	})
}

func (a *APIServer) wireContextAndHandle(c *fiber.Ctx, f func(*Context) (interface{}, error)) {
	ctx, err := a.BuildContext(c)
	if err != nil {
		handleError(c, err)
	}

	data, err := f(ctx)

	if err != nil {
		handleError(c, err)
		return
	}

	handleData(ctx, data)
}

func handleError(ctx *fiber.Ctx, err error) {
	t := errors.GetType(err)
	var status int
	switch t {
	case errors.StatusBadRequest:
		status = 400
		break
	case errors.StatusForbidden:
		status = 401
		break
	case errors.StatusNotFound:
		status = 404
		break
	case errors.StatusUnauthorized:
		status = 403
		break
	case errors.StatusConflict:
		status = 409
		break
	case errors.StatusInternal:
		fallthrough
	default:
		status = 500
		break
	}

	msg := &HTTPMessage{
		Status: strconv.Itoa(status),
		Type: string(errors.GetType(err)),
		Message: err.Error(),
	}

	ctx.Status(status).JSON(msg)
}

func handleData(ctx *Context, data interface{}) {
	// turn data into json
	jData, err := json.Marshal(data)
	if err != nil {
		handleError(ctx.Fiber, errors.Internal("unable to marshal data into json"))
		return
	}

	ct := &HTTPContent{
		Status: strconv.Itoa(200),
		Type: "success",
		Content: jData,
	}

	ctx.Fiber.Status(200).JSON(ct)
}