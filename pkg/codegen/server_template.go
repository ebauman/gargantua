package codegen

var serverTemplate = `package {{ .Type }}

import (
	"github.com/hobbyfarm/gargantua/pkg/apiserver"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type {{ .Type }}Server struct {
	client {{ .TypeClient }}
	dependencies map[string]interface{}
}

const (
	Group = "{{.Group}}"
	Version = "{{.Version}}"
	Kind = "{{.Kind}}"
)

func Register(server *apiserver.APIServer, templateInterface {{.TypeClient}}, dependencies map[string]interface{}) {
	var gvk = []string{Group, Version, Kind}
	var path = strings.Join(gvk, "/")

	objServer := &{{.Type}}Server{}
	objServer.client = templateInterface
	objServer.dependencies = dependencies

{{ if .GenerateList }}
	objServer.Get(path, objServer.do_list)
{{ end }}
{{ if .GenerateGet }}
	objServer.Get(path + "/:name", objServer.do_get)
{{ end }}
{{ if .GenerateCreate }}
	objServer.Post(path, objServer.do_create)
{{ end }}
{{ if .GenerateUpdate }}
	objServer.Put(path + "/:name", objServer.do_update)
{{ end }}
{{ if .GenerateDelete }}
	objServer.Delete(path + "/:name", objServer.do_delete)
{{ end }}
}

{{ if .GenerateList }}
func (objServer *{{.Type}}Server) do_list(ctx *apiserver.Context) (interface{}, error) {
	preData, err := objServer.pre_List{{.Type}}(ctx)
	if err != nil {
		return nil, err
	}
	
	actionData, err = objServer.List{{.Type}}()
	if err != nil {
		return nil, err
	}

	return ToOutputList(actionData), nil
}
{{ end }}

{{ if .GenerateGet }}
func (objServer *{{.Type}}Server) do_get(ctx *apiserver.Context) (interface{}, error) {
	preData, err := objServer.pre_Get{{.Type}}(ctx)
	if err != nil {
		return nil, err
	}

	actionData, err := objServer.Get{{.Type}}(ctx.Fiber.Params("name"))
	if err != nil {
		return nil, err
	}

	return ToOutput(actionData), nil
}
{{ end }}
`
