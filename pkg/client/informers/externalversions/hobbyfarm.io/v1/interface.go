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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/hobbyfarm/gargantua/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AccessCodes returns a AccessCodeInformer.
	AccessCodes() AccessCodeInformer
	// ContentRepositories returns a ContentRepositoryInformer.
	ContentRepositories() ContentRepositoryInformer
	// Courses returns a CourseInformer.
	Courses() CourseInformer
	// DynamicBindConfigurations returns a DynamicBindConfigurationInformer.
	DynamicBindConfigurations() DynamicBindConfigurationInformer
	// DynamicBindRequests returns a DynamicBindRequestInformer.
	DynamicBindRequests() DynamicBindRequestInformer
	// Environments returns a EnvironmentInformer.
	Environments() EnvironmentInformer
	// Scenarios returns a ScenarioInformer.
	Scenarios() ScenarioInformer
	// ScheduledEvents returns a ScheduledEventInformer.
	ScheduledEvents() ScheduledEventInformer
	// Sessions returns a SessionInformer.
	Sessions() SessionInformer
	// Users returns a UserInformer.
	Users() UserInformer
	// VirtualMachines returns a VirtualMachineInformer.
	VirtualMachines() VirtualMachineInformer
	// VirtualMachineClaims returns a VirtualMachineClaimInformer.
	VirtualMachineClaims() VirtualMachineClaimInformer
	// VirtualMachineSets returns a VirtualMachineSetInformer.
	VirtualMachineSets() VirtualMachineSetInformer
	// VirtualMachineTemplates returns a VirtualMachineTemplateInformer.
	VirtualMachineTemplates() VirtualMachineTemplateInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AccessCodes returns a AccessCodeInformer.
func (v *version) AccessCodes() AccessCodeInformer {
	return &accessCodeInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ContentRepositories returns a ContentRepositoryInformer.
func (v *version) ContentRepositories() ContentRepositoryInformer {
	return &contentRepositoryInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Courses returns a CourseInformer.
func (v *version) Courses() CourseInformer {
	return &courseInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// DynamicBindConfigurations returns a DynamicBindConfigurationInformer.
func (v *version) DynamicBindConfigurations() DynamicBindConfigurationInformer {
	return &dynamicBindConfigurationInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// DynamicBindRequests returns a DynamicBindRequestInformer.
func (v *version) DynamicBindRequests() DynamicBindRequestInformer {
	return &dynamicBindRequestInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Environments returns a EnvironmentInformer.
func (v *version) Environments() EnvironmentInformer {
	return &environmentInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Scenarios returns a ScenarioInformer.
func (v *version) Scenarios() ScenarioInformer {
	return &scenarioInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ScheduledEvents returns a ScheduledEventInformer.
func (v *version) ScheduledEvents() ScheduledEventInformer {
	return &scheduledEventInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Sessions returns a SessionInformer.
func (v *version) Sessions() SessionInformer {
	return &sessionInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Users returns a UserInformer.
func (v *version) Users() UserInformer {
	return &userInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// VirtualMachines returns a VirtualMachineInformer.
func (v *version) VirtualMachines() VirtualMachineInformer {
	return &virtualMachineInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// VirtualMachineClaims returns a VirtualMachineClaimInformer.
func (v *version) VirtualMachineClaims() VirtualMachineClaimInformer {
	return &virtualMachineClaimInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// VirtualMachineSets returns a VirtualMachineSetInformer.
func (v *version) VirtualMachineSets() VirtualMachineSetInformer {
	return &virtualMachineSetInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// VirtualMachineTemplates returns a VirtualMachineTemplateInformer.
func (v *version) VirtualMachineTemplates() VirtualMachineTemplateInformer {
	return &virtualMachineTemplateInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
