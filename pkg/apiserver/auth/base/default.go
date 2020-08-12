package base

import (
	"github.com/gofiber/fiber"
	"github.com/hobbyfarm/gargantua/pkg/apiserver/auth"
)

const (
	NilAuthN auth.AuthNProvider = "nil"
	NilAuthZ auth.AuthZProvider = "nil"
)

type DefaultAuthServer struct {
}

func New() (*DefaultAuthServer, error) {
	d := &DefaultAuthServer{}
	return d, nil
}

func (d DefaultAuthServer) AuthNProviders() map[auth.AuthNProvider]func() fiber.Handler {
	providers := make(map[auth.AuthNProvider]func() fiber.Handler)

	providers[NilAuthN] = func() fiber.Handler {
		return func(ctx *fiber.Ctx) {
			ctx.Next()
		}
	}

	return providers
}

func (d DefaultAuthServer) AuthZProviders() map[auth.AuthZProvider]func() fiber.Handler {
	providers := make(map[auth.AuthZProvider]func() fiber.Handler)

	providers[NilAuthZ] = func() fiber.Handler {
		return func(ctx *fiber.Ctx) {
			ctx.Next()
		}
	}

	return providers
}