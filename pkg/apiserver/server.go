package apiserver

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
	"github.com/hobbyfarm/gargantua/pkg/apiserver/errors"
	"strconv"
	"strings"
)

type APIServer struct {
	app *fiber.App
	group *fiber.Group
	validator *validator.Validate
}

type APIServerSettings struct {
	PathPrefix []string
	FiberSettings []*fiber.Settings
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

	a.app = fiber.New()
	a.validator = validator.New()

	if len(set.PathPrefix) > 0 {
		path := strings.Join(set.PathPrefix, "/")
		a.group = a.app.Group(path)
	} else {
		a.group = a.app.Group("")
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

func (a *APIServer) Get(path string, f func(*Context) (interface{}, error)) {
	a.group.Get(path, func (c *fiber.Ctx) {
		a.wireContextAndHandle(c, f)
	})
}

func (a *APIServer) Post(path string, f func(*Context) (interface{}, error)) {
	a.group.Post(path, func (c* fiber.Ctx) {
		a.wireContextAndHandle(c, f)
	})
}

func (a *APIServer) Put(path string, f func(*Context) (interface{}, error)) {
	a.group.Put(path, func (c* fiber.Ctx) {
		a.wireContextAndHandle(c, f)
	})
}

func (a *APIServer) Delete(path string, f func(*Context) (interface{}, error)) {
	a.group.Delete(path, func (c* fiber.Ctx) {
		a.wireContextAndHandle(c, f)
	})
}

func (a *APIServer) wireContextAndHandle(c *fiber.Ctx, f func(*Context) (interface{}, error)) {
	defer c.Next()

	ctx, err := a.BuildContext(c)
	if err != nil {
		handleError(ctx, err)
	}

	data, err := f(ctx)

	if err != nil {
		handleError(ctx, err)
		return
	}

	handleData(ctx, data)
}

func handleError(ctx *Context, err error) {
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

	ctx.Fiber.Status(status).JSON(msg)
}

func handleData(ctx *Context, data interface{}) {
	// turn data into json
	jData, err := json.Marshal(data)
	if err != nil {
		handleError(ctx, errors.Internal("unable to marshal data into json"))
		return
	}

	ct := &HTTPContent{
		Status: strconv.Itoa(200),
		Type: "success",
		Content: jData,
	}

	ctx.Fiber.Status(200).JSON(ct)
}


