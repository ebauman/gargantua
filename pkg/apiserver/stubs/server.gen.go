// Package stubs provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package stubs

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"strings"
)

// Course defines model for course.
type Course struct {
	Name *string `json:"name,omitempty"`
	Spec *struct {
		Description       *string   `json:"description,omitempty"`
		Id                *string   `json:"id,omitempty"`
		KeepaliveDuration *string   `json:"keepalive_duration,omitempty"`
		Name              *string   `json:"name,omitempty"`
		PauseDuration     *string   `json:"pause_duration,omitempty"`
		Pauseable         *bool     `json:"pauseable,omitempty"`
		Scenarios         *[]string `json:"scenarios,omitempty"`
		Virtualmachines   *[]string `json:"virtualmachines,omitempty"`
	} `json:"spec,omitempty"`
}

// Error defines model for error.
type Error struct {
	Message *string `json:"message,omitempty"`
	Status  *int    `json:"status,omitempty"`
	Type    *string `json:"type,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /courses)
	GetCourses(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetCourses converts echo context to params.
func (w *ServerInterfaceWrapper) GetCourses(ctx echo.Context) error {
	var err error

	ctx.Set("BearerAuth.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCourses(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/courses", wrapper.GetCourses)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/5RTUW/UMAz+K5Xh8bTrgKe+DSTGhARIPE4TclNfm9E6keOcdDr1v6Ok1brjerA9Nfnq",
	"z/5ifz6CcYN3TKwBqiME09GA+WhclEDp5MV5ErWUccYho3rwBBUEFcstjBsInsx5eEPBiPVqHa+ybLMK",
	"/yby2Ns9/Wqi4EX2RTEeY/gPN4dg3T9PUDvXE3J+jiFGsS6/wioNYTXLDKAIHtJ9b0Uj9gOazjK9irwA",
	"rn4ko2vIBkjEyXmfBwoB2wuTUdT4XIFlpZZkkXBGWtUSyESxeviZXDKV/UgoJDfNYPnJP7mRGYenLJ2q",
	"T+Xm+KjdC8JTTcs7d+YjuPlxV+ycFGgMhWC5LZCbYkDGNl2+uLo+fEYZCqHgohgKKbfVNGxY/t6itMga",
	"ETawJwlT8uur8qpMYp0nRm+hgvcZSp7RLr97O61HPrek6ZOmkd1210AFt6Sf5pANCAXveA5/V5bTfrES",
	"ZyZ631uTudvH4JZO4ol93grtoII322Vnt/PCznJWPDVu/mrd968p6kN5/SoR/6o9OXKl1DenRRo1sabM",
	"1Jy4CKr744kf7h/Gh3Ecxz8BAAD//2o4NLCVBAAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

// SecurityMiddleware authenticates incoming requests against the specified provider
type SecurityMiddleware func(ctx echo.Context, provider string, scopes []string) error

// secureServer is a small wrapper around an ServerInterface to ensure security through a centralised security middleware
type secureServer struct {
	ServerInterface

	secure SecurityMiddleware
}

func NewSecureServer(srv ServerInterface, securityLayer SecurityMiddleware) ServerInterface {
	return &secureServer{
		ServerInterface: srv,
		secure:          securityLayer,
	}
}

// (GET /courses)
func (s *secureServer) GetCourses(ctx echo.Context) error {
	if err := s.secure(ctx, "BearerAuth", nil); err != nil {
		return err
	}
	return s.ServerInterface.GetCourses(ctx)
}
