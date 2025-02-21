/*
Copyright 2019, 2021, 2022, 2023 The Multi-Cluster App Dispatcher Authors.

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

package v1beta1

import (
	"context"
	time "time"

	controllerv1beta1 "github.com/project-codeflare/multi-cluster-app-dispatcher/pkg/apis/controller/v1beta1"
	versioned "github.com/project-codeflare/multi-cluster-app-dispatcher/pkg/client/clientset/versioned"
	internalinterfaces "github.com/project-codeflare/multi-cluster-app-dispatcher/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/project-codeflare/multi-cluster-app-dispatcher/pkg/client/listers/controller/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AppWrapperInformer provides access to a shared informer and lister for
// AppWrappers.
type AppWrapperInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.AppWrapperLister
}

type appWrapperInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAppWrapperInformer constructs a new informer for AppWrapper type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAppWrapperInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAppWrapperInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAppWrapperInformer constructs a new informer for AppWrapper type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAppWrapperInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.McadV1beta1().AppWrappers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.McadV1beta1().AppWrappers(namespace).Watch(context.TODO(), options)
			},
		},
		&controllerv1beta1.AppWrapper{},
		resyncPeriod,
		indexers,
	)
}

func (f *appWrapperInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAppWrapperInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *appWrapperInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&controllerv1beta1.AppWrapper{}, f.defaultInformer)
}

func (f *appWrapperInformer) Lister() v1beta1.AppWrapperLister {
	return v1beta1.NewAppWrapperLister(f.Informer().GetIndexer())
}
