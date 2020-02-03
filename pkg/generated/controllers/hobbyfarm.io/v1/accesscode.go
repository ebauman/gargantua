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
	"context"
	"time"

	v1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	clientset "github.com/hobbyfarm/gargantua/pkg/generated/clientset/versioned/typed/hobbyfarm.io/v1"
	informers "github.com/hobbyfarm/gargantua/pkg/generated/informers/externalversions/hobbyfarm.io/v1"
	listers "github.com/hobbyfarm/gargantua/pkg/generated/listers/hobbyfarm.io/v1"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type AccessCodeHandler func(string, *v1.AccessCode) (*v1.AccessCode, error)

type AccessCodeController interface {
	generic.ControllerMeta
	AccessCodeClient

	OnChange(ctx context.Context, name string, sync AccessCodeHandler)
	OnRemove(ctx context.Context, name string, sync AccessCodeHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() AccessCodeCache
}

type AccessCodeClient interface {
	Create(*v1.AccessCode) (*v1.AccessCode, error)
	Update(*v1.AccessCode) (*v1.AccessCode, error)

	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v1.AccessCode, error)
	List(opts metav1.ListOptions) (*v1.AccessCodeList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AccessCode, err error)
}

type AccessCodeCache interface {
	Get(name string) (*v1.AccessCode, error)
	List(selector labels.Selector) ([]*v1.AccessCode, error)

	AddIndexer(indexName string, indexer AccessCodeIndexer)
	GetByIndex(indexName, key string) ([]*v1.AccessCode, error)
}

type AccessCodeIndexer func(obj *v1.AccessCode) ([]string, error)

type accessCodeController struct {
	controllerManager *generic.ControllerManager
	clientGetter      clientset.AccessCodesGetter
	informer          informers.AccessCodeInformer
	gvk               schema.GroupVersionKind
}

func NewAccessCodeController(gvk schema.GroupVersionKind, controllerManager *generic.ControllerManager, clientGetter clientset.AccessCodesGetter, informer informers.AccessCodeInformer) AccessCodeController {
	return &accessCodeController{
		controllerManager: controllerManager,
		clientGetter:      clientGetter,
		informer:          informer,
		gvk:               gvk,
	}
}

func FromAccessCodeHandlerToHandler(sync AccessCodeHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.AccessCode
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.AccessCode))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *accessCodeController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.AccessCode))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateAccessCodeDeepCopyOnChange(client AccessCodeClient, obj *v1.AccessCode, handler func(obj *v1.AccessCode) (*v1.AccessCode, error)) (*v1.AccessCode, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *accessCodeController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controllerManager.AddHandler(ctx, c.gvk, c.informer.Informer(), name, handler)
}

func (c *accessCodeController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	removeHandler := generic.NewRemoveHandler(name, c.Updater(), handler)
	c.controllerManager.AddHandler(ctx, c.gvk, c.informer.Informer(), name, removeHandler)
}

func (c *accessCodeController) OnChange(ctx context.Context, name string, sync AccessCodeHandler) {
	c.AddGenericHandler(ctx, name, FromAccessCodeHandlerToHandler(sync))
}

func (c *accessCodeController) OnRemove(ctx context.Context, name string, sync AccessCodeHandler) {
	removeHandler := generic.NewRemoveHandler(name, c.Updater(), FromAccessCodeHandlerToHandler(sync))
	c.AddGenericHandler(ctx, name, removeHandler)
}

func (c *accessCodeController) Enqueue(name string) {
	c.controllerManager.Enqueue(c.gvk, c.informer.Informer(), "", name)
}

func (c *accessCodeController) EnqueueAfter(name string, duration time.Duration) {
	c.controllerManager.EnqueueAfter(c.gvk, c.informer.Informer(), "", name, duration)
}

func (c *accessCodeController) Informer() cache.SharedIndexInformer {
	return c.informer.Informer()
}

func (c *accessCodeController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *accessCodeController) Cache() AccessCodeCache {
	return &accessCodeCache{
		lister:  c.informer.Lister(),
		indexer: c.informer.Informer().GetIndexer(),
	}
}

func (c *accessCodeController) Create(obj *v1.AccessCode) (*v1.AccessCode, error) {
	return c.clientGetter.AccessCodes().Create(obj)
}

func (c *accessCodeController) Update(obj *v1.AccessCode) (*v1.AccessCode, error) {
	return c.clientGetter.AccessCodes().Update(obj)
}

func (c *accessCodeController) Delete(name string, options *metav1.DeleteOptions) error {
	return c.clientGetter.AccessCodes().Delete(name, options)
}

func (c *accessCodeController) Get(name string, options metav1.GetOptions) (*v1.AccessCode, error) {
	return c.clientGetter.AccessCodes().Get(name, options)
}

func (c *accessCodeController) List(opts metav1.ListOptions) (*v1.AccessCodeList, error) {
	return c.clientGetter.AccessCodes().List(opts)
}

func (c *accessCodeController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientGetter.AccessCodes().Watch(opts)
}

func (c *accessCodeController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.AccessCode, err error) {
	return c.clientGetter.AccessCodes().Patch(name, pt, data, subresources...)
}

type accessCodeCache struct {
	lister  listers.AccessCodeLister
	indexer cache.Indexer
}

func (c *accessCodeCache) Get(name string) (*v1.AccessCode, error) {
	return c.lister.Get(name)
}

func (c *accessCodeCache) List(selector labels.Selector) ([]*v1.AccessCode, error) {
	return c.lister.List(selector)
}

func (c *accessCodeCache) AddIndexer(indexName string, indexer AccessCodeIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.AccessCode))
		},
	}))
}

func (c *accessCodeCache) GetByIndex(indexName, key string) (result []*v1.AccessCode, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		result = append(result, obj.(*v1.AccessCode))
	}
	return result, nil
}
