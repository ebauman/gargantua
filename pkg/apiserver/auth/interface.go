package auth

import (
	"github.com/gofiber/fiber"
)

type AuthNProvider string
type AuthZProvider string

type AuthServer interface {
	AuthNProviders() map[AuthNProvider]func() fiber.Handler
	AuthZProviders() map[AuthZProvider]func() fiber.Handler
}