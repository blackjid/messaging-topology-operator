/*
RabbitMQ Messaging Topology Kubernetes Operator
Copyright 2021 VMware, Inc.

This product is licensed to you under the Mozilla Public License 2.0 license (the "License").  You may not use this product except in compliance with the Mozilla 2.0 License.

This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	rabbitmqcomv1beta1 "github.com/rabbitmq/messaging-topology-operator/api/v1beta1"
	versioned "github.com/rabbitmq/messaging-topology-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/rabbitmq/messaging-topology-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/rabbitmq/messaging-topology-operator/pkg/generated/listers/rabbitmq.com/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ShovelInformer provides access to a shared informer and lister for
// Shovels.
type ShovelInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ShovelLister
}

type shovelInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewShovelInformer constructs a new informer for Shovel type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewShovelInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredShovelInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredShovelInformer constructs a new informer for Shovel type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredShovelInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RabbitmqV1beta1().Shovels(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RabbitmqV1beta1().Shovels(namespace).Watch(context.TODO(), options)
			},
		},
		&rabbitmqcomv1beta1.Shovel{},
		resyncPeriod,
		indexers,
	)
}

func (f *shovelInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredShovelInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *shovelInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rabbitmqcomv1beta1.Shovel{}, f.defaultInformer)
}

func (f *shovelInformer) Lister() v1beta1.ShovelLister {
	return v1beta1.NewShovelLister(f.Informer().GetIndexer())
}
