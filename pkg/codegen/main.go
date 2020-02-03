package main

import (
	hfv1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	tfv1 "github.com/hobbyfarm/gargantua/pkg/apis/terraformcontroller.cattle.io/v1"
	"github.com/rancher/wrangler/pkg/controller-gen"
	"github.com/rancher/wrangler/pkg/controller-gen/args"
	"os"
)

func main() {
	os.Unsetenv("GOPATH")
	controllergen.Run(args.Options{
		OutputPackage: "github.com/hobbyfarm/gargantua/pkg/generated",
		Boilerplate: "hack/boilerplate.go.txt",
		Groups: map[string]args.Group {
			"hobbyfarm.io": {
				Types: []interface{} {
					hfv1.Environment{},
					hfv1.VirtualMachineClaim{},
					hfv1.VirtualMachineSet{},
					hfv1.ScenarioSession{},
					hfv1.Scenario{},
					hfv1.AccessCode{},
					hfv1.User{},
					hfv1.DynamicBindConfiguration{},
					hfv1.DynamicBindRequest{},
					hfv1.ScheduledEvent{},
					hfv1.VirtualMachine{},
					hfv1.VirtualMachineTemplate{},
				},
				GenerateClients: true,
				GenerateTypes: true,
			},
			"terraformcontroller.cattle.io": {
				Types: []interface{} {
					tfv1.State{},
					tfv1.Execution{},
				},
				GenerateTypes: true,
				GenerateClients: true,
			},
		},
	})
}