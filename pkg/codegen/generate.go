package codegen

import (
	"os"
	"path"
	"text/template"
)

func Generate(config Config) error {
	//config := Config{
	//	GenerateGet: true,
	//	GenerateList: true,
	//	OutputPath: "github.com/hobbyfarm/gargantua/pkg/apis",
	//}
	// load template
	tmpl, err := template.New("server.template").
		Parse(serverTemplate)

	if err != nil {
		return err
	}

	// check for and maybe build directory structure
	gvk := config.Object.GetObjectKind().GroupVersionKind()
	pkgPath := path.Join(config.OutputPath, gvk.Group, gvk.Version, gvk.Kind)

	if _, err := os.Stat(pkgPath); os.IsNotExist(err) {
		err = os.MkdirAll(pkgPath, os.ModeDir)
		if err != nil {
			return err
		}
	}

	serverPath := path.Join(pkgPath, "zz_generated_server.go")
	output, err := os.Create(serverPath)
	if err != nil {
		return err
	}
	defer output.Close()

	return tmpl.Execute(output, config)
}