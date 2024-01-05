package kubernetes

import (
	"github.com/acorn-io/mink/pkg/strategy/remote"
	v1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/v3/pkg/storage/registry/accesscode"
	"github.com/hobbyfarm/gargantua/v3/pkg/storage/registry/course"
	"github.com/hobbyfarm/gargantua/v3/pkg/storage/registry/scenario"
	"k8s.io/apiserver/pkg/registry/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func APIGroups(client client.WithWatch) (map[string]rest.Storage, error) {
	scenarioRemote := remote.NewRemote(&v1.Scenario{}, client)
	courseRemote := remote.NewRemote(&v1.Course{}, client)
	accessCodeRemote := remote.NewRemote(&v1.AccessCode{}, client)

	scenarioStorage, err := scenario.NewScenarioStorage(scenarioRemote)
	if err != nil {
		return nil, err
	}
	courseStorage, err := course.NewCourseStorage(courseRemote)
	if err != nil {
		return nil, err
	}
	accessCodeStorage, err := accesscode.NewAccessCodeStorage(accessCodeRemote,
		scenarioRemote, courseRemote)
	if err != nil {
		return nil, err
	}

	return map[string]rest.Storage{
		"accesscodes": accessCodeStorage,
		"scenarios":   scenarioStorage,
		"courses":     courseStorage,
	}, nil
}
