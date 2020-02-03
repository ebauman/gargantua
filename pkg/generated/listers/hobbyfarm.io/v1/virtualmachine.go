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

// Code generated by main. DO NOT EDIT.

package v1

import (
	v1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VirtualMachineLister helps list VirtualMachines.
type VirtualMachineLister interface {
	// List lists all VirtualMachines in the indexer.
	List(selector labels.Selector) (ret []*v1.VirtualMachine, err error)
	// Get retrieves the VirtualMachine from the index for a given name.
	Get(name string) (*v1.VirtualMachine, error)
	VirtualMachineListerExpansion
}

// virtualMachineLister implements the VirtualMachineLister interface.
type virtualMachineLister struct {
	indexer cache.Indexer
}

// NewVirtualMachineLister returns a new VirtualMachineLister.
func NewVirtualMachineLister(indexer cache.Indexer) VirtualMachineLister {
	return &virtualMachineLister{indexer: indexer}
}

// List lists all VirtualMachines in the indexer.
func (s *virtualMachineLister) List(selector labels.Selector) (ret []*v1.VirtualMachine, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VirtualMachine))
	})
	return ret, err
}

// Get retrieves the VirtualMachine from the index for a given name.
func (s *virtualMachineLister) Get(name string) (*v1.VirtualMachine, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("virtualmachine"), name)
	}
	return obj.(*v1.VirtualMachine), nil
}
