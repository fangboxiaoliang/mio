/*
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

import (
	time "time"

	miov1alpha1 "hidevops.io/mio/pkg/apis/mio/v1alpha1"
	versioned "hidevops.io/mio/pkg/client/clientset/versioned"
	internalinterfaces "hidevops.io/mio/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "hidevops.io/mio/pkg/client/listers/mio/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceConfigInformer provides access to a shared informer and lister for
// ServiceConfigs.
type ServiceConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ServiceConfigLister
}

type serviceConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceConfigInformer constructs a new informer for ServiceConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceConfigInformer constructs a new informer for ServiceConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MioV1alpha1().ServiceConfigs(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MioV1alpha1().ServiceConfigs(namespace).Watch(options)
			},
		},
		&miov1alpha1.ServiceConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&miov1alpha1.ServiceConfig{}, f.defaultInformer)
}

func (f *serviceConfigInformer) Lister() v1alpha1.ServiceConfigLister {
	return v1alpha1.NewServiceConfigLister(f.Informer().GetIndexer())
}
