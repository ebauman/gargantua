package generate

var mapperTemplate = `package {{.Kind | ToLower}}

import objPkg "{{.KindPackage}}"

type Flat{{.Kind}} struct {
	Name string ` + "`json:\"name\" validate:\"required\"`" + `
	{{ range .FieldMap }}{{ .Name }} {{ .Type }} ` + "`{{ .Tag }}`" + `
	{{ end }}
}

func FromInput(f *Flat{{.Kind}}) *objPkg.{{.Kind}} {
	obj := &objPkg.{{.Kind}}{}
	obj.Name = f.Name
	{{ range .FieldMap }}obj.Spec.{{.Name}} = f.{{.Name}}
	{{ end }}
	return obj
}

func ToOutput(i *objPkg.{{.Kind}}) *Flat{{.Kind}} {
	obj := &Flat{{.Kind}}{}

	obj.Name = i.Name
	{{ range .FieldMap }}obj.{{.Name}} = i.Spec.{{.Name}}
	{{ end }}

	return obj
}

func ToOutputList(i *[]objPkg.{{.Kind}}) *[]Flat{{.Kind}} {
	out := make([]Flat{{.Kind}}, len(*i))
	for _, o := range *i {
		obj := ToOutput(&o)
		out = append(out, *obj)
	}

	return &out
}
`