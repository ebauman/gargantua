package generate

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"
	"text/template"
)

func Generate(config Config) error {
	gopath := os.Getenv("GOPATH")

	funcMap := template.FuncMap{
		"ToLower": strings.ToLower,
		"GetAuthN": func(m MethodConfig, o ObjectConfig, g Group) string {
			if m.AuthN != "" {
				return m.AuthN
			} else if o.AuthN != "" {
				return o.AuthN
			} else {
				return g.AuthN
			}
		},
		"GetAuthZ": func(m MethodConfig, o ObjectConfig, g Group) string {
			if m.AuthZ != "" {
				return m.AuthZ
			} else if o.AuthZ != "" {
				return o.AuthZ
			} else {
				return g.AuthZ
			}
		},
	}

	serverTmpl, err := template.New("server.template").
		Funcs(funcMap).
		Parse(serverTemplate)
	mapperTmpl, err := template.New("mapper.template").
		Funcs(funcMap).
		Parse(mapperTemplate)

	if err != nil {
		return err
	}

	for groupName, groupConfig := range config.Groups {
		// for each group
		for _, objectConfig := range groupConfig.Types {
			typeNameLower := strings.ToLower(reflect.TypeOf(objectConfig.Type).Name())
			pkgPath := path.Join(gopath, "src", config.OutputPath, groupName, typeNameLower)

			err := checkOrCreatePath(pkgPath)
			if err != nil {
				return err
			}

			serverPath := path.Join(pkgPath, "zz_generated_server.go")
			mapperPath := path.Join(pkgPath, "zz_generated_mapper.go")
			serverOutput, err := os.Create(serverPath)
			mapperOutput, err := os.Create(mapperPath)

			if err != nil {
				return err
			}
			defer serverOutput.Close()
			defer mapperOutput.Close()

			fieldMap, err := flattenStruct(objectConfig.Type)
			cfg := struct {
				Kind         string
				KindPackage  string
				GroupName    string
				GroupConfig  Group
				ObjectConfig ObjectConfig
				FieldMap     []reflect.StructField
			}{
				reflect.TypeOf(objectConfig.Type).Name(),
				reflect.TypeOf(objectConfig.Type).PkgPath(),
				groupName,
				groupConfig,
				objectConfig,
				fieldMap,
			}

			err = serverTmpl.Execute(serverOutput, cfg)
			if err != nil {
				return err
			}
			err = mapperTmpl.Execute(mapperOutput, cfg)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkOrCreatePath(path string) error {
	_, err := os.Stat(path)

	if !os.IsNotExist(err) {
		return err
	}

	if os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func flattenStruct(obj interface{}) ([]reflect.StructField, error) {
	fields := make([]reflect.StructField, 0)

	v := reflect.TypeOf(obj)

	// get all of the spec fields
	specField, ok := v.FieldByName("Spec")
	if !ok {
		return nil, fmt.Errorf("error retrieving Spec field in reflected type %v", v)
	}

	for i := 0; i < specField.Type.NumField(); i++ {
		field := specField.Type.Field(i)
		fields = append(fields, field)
	}

	return fields, nil
}
