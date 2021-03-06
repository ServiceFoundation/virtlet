/*
Copyright 2019 Mirantis

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
	time "time"

	virtlet_k8s_v1 "github.com/Mirantis/virtlet/pkg/api/virtlet.k8s/v1"
	versioned "github.com/Mirantis/virtlet/pkg/client/clientset/versioned"
	internalinterfaces "github.com/Mirantis/virtlet/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/Mirantis/virtlet/pkg/client/listers/virtlet.k8s/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VirtletConfigMappingInformer provides access to a shared informer and lister for
// VirtletConfigMappings.
type VirtletConfigMappingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.VirtletConfigMappingLister
}

type virtletConfigMappingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVirtletConfigMappingInformer constructs a new informer for VirtletConfigMapping type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVirtletConfigMappingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVirtletConfigMappingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVirtletConfigMappingInformer constructs a new informer for VirtletConfigMapping type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVirtletConfigMappingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VirtletV1().VirtletConfigMappings(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VirtletV1().VirtletConfigMappings(namespace).Watch(options)
			},
		},
		&virtlet_k8s_v1.VirtletConfigMapping{},
		resyncPeriod,
		indexers,
	)
}

func (f *virtletConfigMappingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVirtletConfigMappingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *virtletConfigMappingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&virtlet_k8s_v1.VirtletConfigMapping{}, f.defaultInformer)
}

func (f *virtletConfigMappingInformer) Lister() v1.VirtletConfigMappingLister {
	return v1.NewVirtletConfigMappingLister(f.Informer().GetIndexer())
}
