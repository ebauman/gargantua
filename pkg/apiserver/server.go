package apiserver

import (
	"github.com/gofiber/fiber"
)

type APIServer struct {
	app *fiber.App
}

type APIServerSettings struct {
	FiberSettings []*fiber.Settings
}

type Context struct {
	Fiber *fiber.Ctx

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
	app := fiber.New(set.FiberSettings...)
	a.app = app

	return a
}

func RunServer(port int, stop chan interface{}) {

}

func BuildContext (ctx *fiber.Ctx) (error, *Context) {
	// given a fiber context, build an api context
	c := &Context{
		Fiber: ctx,
		Locals: map[string]interface{}{},
	}

	return nil, c
}

func (a *APIServer) Get(path string, f func(*Context) (interface{}, error)) {
	a.app.Get(path, func (c *fiber.Ctx) {
		defer c.Next()

		// give me a context, then give that to the inbound f()
		err, ctx := BuildContext(c)
		if err != nil {
			// TODO - Handle error
		}

		data, err := f(ctx)

		if err != nil {
			handleError(ctx, err)
			return
		}

		handleData(ctx, data)
	})
}

func handleError(ctx *Context, err error) {
	// figure out the type of error, write error object and status

}

func handleData(ctx *Context, data interface{}) {
	// do things with the data
}

