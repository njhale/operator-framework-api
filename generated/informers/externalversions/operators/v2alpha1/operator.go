/*
Copyright 2019 Red Hat, Inc.

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

package v2alpha1

import (
	time "time"

	versioned "github.com/operator-framework/api/generated/clientset/versioned"
	internalinterfaces "github.com/operator-framework/api/generated/informers/externalversions/internalinterfaces"
	v2alpha1 "github.com/operator-framework/api/generated/listers/operators/v2alpha1"
	operatorsv2alpha1 "github.com/operator-framework/api/operators/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OperatorInformer provides access to a shared informer and lister for
// Operators.
type OperatorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2alpha1.OperatorLister
}

type operatorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewOperatorInformer constructs a new informer for Operator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOperatorInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOperatorInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredOperatorInformer constructs a new informer for Operator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOperatorInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorsV2alpha1().Operators().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorsV2alpha1().Operators().Watch(options)
			},
		},
		&operatorsv2alpha1.Operator{},
		resyncPeriod,
		indexers,
	)
}

func (f *operatorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOperatorInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *operatorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operatorsv2alpha1.Operator{}, f.defaultInformer)
}

func (f *operatorInformer) Lister() v2alpha1.OperatorLister {
	return v2alpha1.NewOperatorLister(f.Informer().GetIndexer())
}
