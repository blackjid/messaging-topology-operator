/*
RabbitMQ Messaging Topology Kubernetes Operator
Copyright 2021 VMware, Inc.

This product is licensed to you under the Mozilla Public License 2.0 license (the "License").  You may not use this product except in compliance with the Mozilla 2.0 License.

This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/rabbitmq/messaging-topology-operator/api/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PolicyLister helps list Policies.
// All objects returned here must be treated as read-only.
type PolicyLister interface {
	// List lists all Policies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha2.Policy, err error)
	// Policies returns an object that can list and get Policies.
	Policies(namespace string) PolicyNamespaceLister
	PolicyListerExpansion
}

// policyLister implements the PolicyLister interface.
type policyLister struct {
	indexer cache.Indexer
}

// NewPolicyLister returns a new PolicyLister.
func NewPolicyLister(indexer cache.Indexer) PolicyLister {
	return &policyLister{indexer: indexer}
}

// List lists all Policies in the indexer.
func (s *policyLister) List(selector labels.Selector) (ret []*v1alpha2.Policy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.Policy))
	})
	return ret, err
}

// Policies returns an object that can list and get Policies.
func (s *policyLister) Policies(namespace string) PolicyNamespaceLister {
	return policyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PolicyNamespaceLister helps list and get Policies.
// All objects returned here must be treated as read-only.
type PolicyNamespaceLister interface {
	// List lists all Policies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha2.Policy, err error)
	// Get retrieves the Policy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha2.Policy, error)
	PolicyNamespaceListerExpansion
}

// policyNamespaceLister implements the PolicyNamespaceLister
// interface.
type policyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Policies in the indexer for a given namespace.
func (s policyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.Policy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.Policy))
	})
	return ret, err
}

// Get retrieves the Policy from the indexer for a given namespace and name.
func (s policyNamespaceLister) Get(name string) (*v1alpha2.Policy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("policy"), name)
	}
	return obj.(*v1alpha2.Policy), nil
}
