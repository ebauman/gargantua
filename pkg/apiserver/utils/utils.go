package utils

import (
	"github.com/gofiber/fiber"
	"github.com/hobbyfarm/gargantua/pkg/apiserver/auth"
)

func MergeAuthNMaps(a map[auth.AuthNProvider]func() fiber.Handler, b map[auth.AuthNProvider]func() fiber.Handler) {
	for k, v := range b {
		a[k] = v
	}
}

func MergeAuthZMaps(a map[auth.AuthZProvider]func() fiber.Handler, b map[auth.AuthZProvider]func() fiber.Handler) {
	for k, v := range b {
		a[k] = v
	}
}