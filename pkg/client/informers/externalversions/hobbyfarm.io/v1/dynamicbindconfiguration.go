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
	"context"
	time "time"

	hobbyfarmiov1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	versioned "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	internalinterfaces "github.com/hobbyfarm/gargantua/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/hobbyfarm/gargantua/pkg/client/listers/hobbyfarm.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DynamicBindConfigurationInformer provides access to a shared informer and lister for
// DynamicBindConfigurations.
type DynamicBindConfigurationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DynamicBindConfigurationLister
}

type dynamicBindConfigurationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewDynamicBindConfigurationInformer constructs a new informer for DynamicBindConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDynamicBindConfigurationInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDynamicBindConfigurationInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredDynamicBindConfigurationInformer constructs a new informer for DynamicBindConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDynamicBindConfigurationInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HobbyfarmV1().DynamicBindConfigurations().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HobbyfarmV1().DynamicBindConfigurations().Watch(context.TODO(), options)
			},
		},
		&hobbyfarmiov1.DynamicBindConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *dynamicBindConfigurationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDynamicBindConfigurationInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *dynamicBindConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&hobbyfarmiov1.DynamicBindConfiguration{}, f.defaultInformer)
}

func (f *dynamicBindConfigurationInformer) Lister() v1.DynamicBindConfigurationLister {
	return v1.NewDynamicBindConfigurationLister(f.Informer().GetIndexer())
}
