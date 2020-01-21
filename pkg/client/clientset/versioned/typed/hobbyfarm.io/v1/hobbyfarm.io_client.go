/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type HobbyfarmV1Interface interface {
	RESTClient() rest.Interface
	AccessCodesGetter
	DynamicBindConfigurationsGetter
	DynamicBindRequestsGetter
	EnvironmentsGetter
	ScenariosGetter
	ScenarioSessionsGetter
	ScheduledEventsGetter
	UsersGetter
	VirtualMachinesGetter
	VirtualMachineClaimsGetter
	VirtualMachineSetsGetter
	VirtualMachineTemplatesGetter
}

// HobbyfarmV1Client is used to interact with features provided by the hobbyfarm.io group.
type HobbyfarmV1Client struct {
	restClient rest.Interface
}

func (c *HobbyfarmV1Client) AccessCodes() AccessCodeInterface {
	return newAccessCodes(c)
}

func (c *HobbyfarmV1Client) DynamicBindConfigurations() DynamicBindConfigurationInterface {
	return newDynamicBindConfigurations(c)
}

func (c *HobbyfarmV1Client) DynamicBindRequests() DynamicBindRequestInterface {
	return newDynamicBindRequests(c)
}

func (c *HobbyfarmV1Client) Environments() EnvironmentInterface {
	return newEnvironments(c)
}

func (c *HobbyfarmV1Client) Scenarios() ScenarioInterface {
	return newScenarios(c)
}

func (c *HobbyfarmV1Client) ScenarioSessions() ScenarioSessionInterface {
	return newScenarioSessions(c)
}

func (c *HobbyfarmV1Client) ScheduledEvents() ScheduledEventInterface {
	return newScheduledEvents(c)
}

func (c *HobbyfarmV1Client) Users() UserInterface {
	return newUsers(c)
}

func (c *HobbyfarmV1Client) VirtualMachines() VirtualMachineInterface {
	return newVirtualMachines(c)
}

func (c *HobbyfarmV1Client) VirtualMachineClaims() VirtualMachineClaimInterface {
	return newVirtualMachineClaims(c)
}

func (c *HobbyfarmV1Client) VirtualMachineSets() VirtualMachineSetInterface {
	return newVirtualMachineSets(c)
}

func (c *HobbyfarmV1Client) VirtualMachineTemplates() VirtualMachineTemplateInterface {
	return newVirtualMachineTemplates(c)
}

// NewForConfig creates a new HobbyfarmV1Client for the given config.
func NewForConfig(c *rest.Config) (*HobbyfarmV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &HobbyfarmV1Client{client}, nil
}

// NewForConfigOrDie creates a new HobbyfarmV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *HobbyfarmV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new HobbyfarmV1Client for the given RESTClient.
func New(c rest.Interface) *HobbyfarmV1Client {
	return &HobbyfarmV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *HobbyfarmV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
