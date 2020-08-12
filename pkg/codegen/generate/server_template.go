package generate

var serverTemplate = `package {{ .Kind | ToLower }}

import (
	objPkg "{{.KindPackage}}"
	"github.com/hobbyfarm/gargantua/pkg/apiserver"
	objClientset "{{.GroupConfig.ClientSetPackage}}"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

type Server struct {
	client objClientset.{{ .Kind }}Interface
	dependencies map[string]interface{}
}

const (
	Group = "{{.GroupName}}"
	Kind = "{{.Kind | ToLower}}"
)

func Register(server *apiserver.APIServer, templateInterface objClientset.{{.Kind}}Interface) {
	var gvk = []string{Group, Kind}
	{{- if ne .ObjectConfig.PathOverride "" }}
	var path = "{{.ObjectConfig.PathOverride}}"
	{{- else }}
	var path = strings.Join(gvk, "/")
	{{- end }}

	objServer := &Server{}
	objServer.client = templateInterface
	{{ "" }}
	{{- if .ObjectConfig.ListConfig.Generate }}
	server.Get(path, objServer.do_list, "{{GetAuthN .ObjectConfig.ListConfig .ObjectConfig .GroupConfig}}", "{{GetAuthZ .ObjectConfig.ListConfig .ObjectConfig .GroupConfig}}")
	{{- end }}
	{{- if .ObjectConfig.CreateConfig.Generate }}
	server.Post(path, objServer.do_create, "{{GetAuthN .ObjectConfig.CreateConfig .ObjectConfig .GroupConfig}}", "{{GetAuthZ .ObjectConfig.CreateConfig .ObjectConfig .GroupConfig}}")
	{{- end }}
	{{- if .ObjectConfig.GetConfig.Generate }}
	server.Get(path + "/:name", objServer.do_get, "{{GetAuthN .ObjectConfig.GetConfig .ObjectConfig .GroupConfig}}", "{{GetAuthZ .ObjectConfig.GetConfig .ObjectConfig .GroupConfig}}")
	{{- end }}
	{{- if .ObjectConfig.UpdateConfig.Generate }}
	server.Put(path + "/:name", objServer.do_update, "{{GetAuthN .ObjectConfig.UpdateConfig .ObjectConfig .GroupConfig}}", "{{GetAuthZ .ObjectConfig.UpdateConfig .ObjectConfig .GroupConfig}}")
	{{- end }}
	{{- if .ObjectConfig.DeleteConfig.Generate }}
	server.Delete(path + "/:name", objServer.do_delete, "{{GetAuthN .ObjectConfig.DeleteConfig .ObjectConfig .GroupConfig}}", "{{GetAuthZ .ObjectConfig.DeleteConfig .ObjectConfig .GroupConfig}}")
	{{- end }}
}

{{- if .ObjectConfig.ListConfig.Generate }}
func (objServer *Server) do_list(ctx *apiserver.Context) (interface{}, error) {
	actionData, err := objServer.List{{.Kind}}(ctx)
	if err != nil {
		return nil, err
	}

	return ToOutputList(actionData), nil
}

func (objServer *Server) List{{.Kind}}(ctx *apiserver.Context) (*[]objPkg.{{.Kind}}, error) {
	data, err := objServer.client.List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return &data.Items, nil
}{{ end }}

{{- if .ObjectConfig.CreateConfig.Generate}}
func (objServer *Server) do_create(ctx *apiserver.Context) (interface{}, error) {
	preData, err := objServer.pre_Create{{.Kind}}(ctx)
	if err != nil {
		return nil, err
	}

	actionData, err := objServer.Create{{.Kind}}(preData)
	if err != nil {
		return nil, err
	}

	return ToOutput(actionData), nil
}

func (objServer *Server) pre_Create{{.Kind}}(ctx *apiserver.Context) (*objPkg.{{.Kind}}, error) {
	// marshal the incoming object into a {{.Kind}}
	obj, err := objServer.genericObjectMarshal(ctx)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (objServer *Server) Create{{.Kind}}(obj *objPkg.{{.Kind}}) (*objPkg.{{.Kind}}, error) {
	data, err := objServer.client.Create(obj)
	if err != nil {
		return nil, err
	}

	return data, nil
}{{ end }}

{{- if .ObjectConfig.GetConfig.Generate }}
func (objServer *Server) do_get(ctx *apiserver.Context) (interface{}, error) {
	data, err := objServer.Get{{.Kind}}(ctx.Fiber.Params("name"))
	if err != nil {
		return nil, err
	}

	return ToOutput(data), nil
}

func (objServer *Server) Get{{.Kind}}(name string) (*objPkg.{{.Kind}}, error) {
	data, err := objServer.client.Get(name, metav1.GetOptions{})

	if err != nil {
		return nil, err
	}

	return data, nil
}{{ end }}

{{- if .ObjectConfig.UpdateConfig.Generate }}
func (objServer *Server) do_update(ctx *apiserver.Context) (interface{}, error) {
	preObj, err := objServer.pre_Update{{.Kind}}(ctx)
	if err != nil {
		return nil, err
	}

	data, err := objServer.Update{{.Kind}}(preObj)
	if err != nil {
		return nil, err
	}

	return ToOutput(data), nil
}

func (objServer *Server) pre_Update{{.Kind}}(ctx *apiserver.Context) (*objPkg.{{.Kind}}, error) {
	obj, err := objServer.genericObjectMarshal(ctx)
	if err != nil {
		return nil, err
	}

	return obj, nil
}

func (objServer *Server) Update{{.Kind}}(obj *objPkg.{{.Kind}}) (*objPkg.{{.Kind}}, error) {
	data, err := objServer.client.Update(obj)

	if err != nil {
		return nil, err
	}

	return data, nil
}{{ end }}

{{- if .ObjectConfig.DeleteConfig.Generate }}
func (objServer *Server) do_delete(ctx *apiserver.Context) (interface{}, error) {
	err := objServer.Delete{{.Kind}}(ctx.Fiber.Params("name"))

	return nil, err
}

func (objServer *Server) Delete{{.Kind}}(name string) (error) {
	err := objServer.client.Delete(name, &metav1.DeleteOptions{})

	return err
}{{ end }}

func (objServer *Server) genericObjectMarshal(ctx *apiserver.Context) (*objPkg.{{.Kind}}, error) {
	flatObj := &Flat{{.Kind}}{}
	
	err := ctx.Fiber.BodyParser(&flatObj)
	if err != nil {
		return nil, err
	}

	err = ctx.Validator.Struct(flatObj)
	if err != nil {
		return nil, err
	}

	obj := FromInput(flatObj)
	return obj, nil
}
`
