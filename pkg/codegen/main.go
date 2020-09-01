//go:generate go run github.com/hobbyfarm/gargantua/pkg/codegen
package main

import (
	"bytes"
	"fmt"
	"github.com/ebauman/oapi-codegen/pkg/codegen"
	"github.com/ebauman/oapi-codegen/pkg/util"
	"github.com/getkin/kin-openapi/openapi3"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

const securityLayerTmpl = `
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

{{range .}}
{{- if .SecurityDefinitions -}}
// ({{.Method}} {{.Path}})
func (s *secureServer) {{.OperationId}}(ctx echo.Context{{genParamArgs .PathParams}}{{if .RequiresParamObject}}, params {{.OperationId}}Params{{end}}) error {
	{{range .SecurityDefinitions -}}
	if err := s.secure(ctx, "{{.ProviderName}}", nil); err != nil {
		return err
	}
	{{end -}}

	return s.ServerInterface.{{.OperationId}}(ctx{{genParamNames .PathParams}}{{if .RequiresParamObject}}, params{{end}})
}
{{end}}
{{end}}
`

func main() {
	opts := codegen.Options{
		GenerateEchoServer: true,
		GenerateTypes: true,
		EmbedSpec: true,
	}

	swagger, err := util.LoadSwagger("openapi.yaml")
	if err != nil {
		log.Fatalf("error loading openapi spec: %s", err)
	}

	code, err := codegen.Generate(swagger, "stubs", opts)
	if err != nil {
		log.Fatalf("error generating openapi code: %s", err)
	}

	rawCode := []byte(code)

	if len(swagger.Components.SecuritySchemes) > 0 {
		rawCode = generateSecurityLayer(code, swagger, err, rawCode)
	}

	err = ioutil.WriteFile(os.Getenv("GOPATH") + "/src/github.com/hobbyfarm/gargantua/pkg/apiserver/stubs/server.gen.go", rawCode, 0644)
	if err != nil {
		log.Fatalf("error writing code to file: %s", err)
	}
}


func generateSecurityLayer(code string, swagger *openapi3.Swagger, err error, currentCode []byte) []byte {
	writer := bytes.Buffer{}
	writer.WriteString(code)

	// err can not occur as this function is also used by codegen.Generate before (and would already fail there)
	ops, _ := codegen.OperationDefinitions(swagger)

	tmpl := template.Must(template.New("").Funcs(codegen.TemplateFunctions).Parse(securityLayerTmpl))
	err = tmpl.Execute(&writer, ops)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error generating security layer: %s", err)
		os.Exit(1)
	}

	bytes, err := format.Source(writer.Bytes())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error formating generated security layer: %s", err)
		os.Exit(1)
	}
	return bytes
}