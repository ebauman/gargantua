package main

import (
	v1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/codegen/generate"
	"log"
)

func main() {
	config := generate.Config{
		OutputPath: "github.com/hobbyfarm/gargantua/pkg/apiserver",
		Groups: map[string]generate.Group{
			"hobbyfarm.io": generate.Group{
				ClientSetPackage: "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned/typed/hobbyfarm.io/v1",
				AuthN: "default",
				Types: []generate.ObjectConfig{
					{
						Type: v1.VirtualMachine{},
						ListConfig: generate.MethodConfig{
							Generate: true,
						},
						GetConfig: generate.MethodConfig{
							Generate: true,
						},
						CreateConfig: generate.MethodConfig{
							Generate: true,
							AuthZ: "admin",
						},
						UpdateConfig: generate.MethodConfig{
							Generate: true,
							AuthZ: "admin",
						},
						DeleteConfig: generate.MethodConfig{
							Generate: true,
							AuthZ: "admin",
						},
					},
					{
						Type: v1.VirtualMachineSet{},
						ListConfig: generate.MethodConfig{
							Generate: true,
						},
					},
					{
						Type: v1.AccessCode{},
						ListConfig: generate.MethodConfig{
							Generate: true,
						},
					},
				},
			},
		},
	}

	err := generate.Generate(config)
	if err != nil {
		log.Fatal(err)
	}
}