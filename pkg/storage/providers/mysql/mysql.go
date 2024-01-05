package mysql

import (
	"github.com/acorn-io/mink/pkg/db"
	"github.com/acorn-io/mink/pkg/serializer"
	"github.com/acorn-io/mink/pkg/strategy"
	v1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	hobbyfarmScheme "github.com/hobbyfarm/gargantua/v3/pkg/scheme"
	"github.com/hobbyfarm/gargantua/v3/pkg/storage/registry/accesscode"
	"github.com/hobbyfarm/gargantua/v3/pkg/storage/registry/scenario"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
)

func APIGroups(dsn string) (*genericapiserver.APIGroupInfo, error) {
	scheme := hobbyfarmScheme.Scheme

	dbFactory, err := db.NewFactory(scheme, dsn)
	if err != nil {
		return nil, err
	}

	var accessCodeStrategy, scenarioStrategy, courseStrategy strategy.CompleteStrategy
	accessCodeStrategy, err = dbFactory.NewDBStrategy(&v1.AccessCode{})
	scenarioStrategy, err = dbFactory.NewDBStrategy(&v1.Scenario{})
	courseStrategy, err = dbFactory.NewDBStrategy(&v1.Course{})

	accessCodeStorage, err := accesscode.NewAccessCodeStorage(accessCodeStrategy, scenarioStrategy, courseStrategy)
	if err != nil {
		return nil, err
	}

	scenarioStorage, err := scenario.NewScenarioStorage(scenarioStrategy)
	if err != nil {
		return nil, err
	}

	courseStorage, err := scenario.NewScenarioStorage(courseStrategy)
	if err != nil {
		return nil, err
	}

	stores := map[string]rest.Storage{
		"accesscodes": accessCodeStorage,
		"scenarios":   scenarioStorage,
		"courses":     courseStorage,
	}

	if err = v1.AddToScheme(scheme); err != nil {
		return nil, err
	}

	apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(v1.SchemeGroupVersion.Group, scheme,
		hobbyfarmScheme.ParameterCodec, hobbyfarmScheme.Codecs)
	apiGroupInfo.VersionedResourcesStorageMap["v1"] = stores
	apiGroupInfo.NegotiatedSerializer = serializer.NewNoProtobufSerializer(apiGroupInfo.NegotiatedSerializer)

	return &apiGroupInfo, nil
}
