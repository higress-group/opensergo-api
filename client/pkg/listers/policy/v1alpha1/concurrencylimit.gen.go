// Copyright 2023, OpenSergo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "opensergo.io/api/client/pkg/apis/policy/v1alpha1"
)

// ConcurrencyLimitLister helps list ConcurrencyLimits.
// All objects returned here must be treated as read-only.
type ConcurrencyLimitLister interface {
	// List lists all ConcurrencyLimits in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConcurrencyLimit, err error)
	// ConcurrencyLimits returns an object that can list and get ConcurrencyLimits.
	ConcurrencyLimits(namespace string) ConcurrencyLimitNamespaceLister
	ConcurrencyLimitListerExpansion
}

// concurrencyLimitLister implements the ConcurrencyLimitLister interface.
type concurrencyLimitLister struct {
	indexer cache.Indexer
}

// NewConcurrencyLimitLister returns a new ConcurrencyLimitLister.
func NewConcurrencyLimitLister(indexer cache.Indexer) ConcurrencyLimitLister {
	return &concurrencyLimitLister{indexer: indexer}
}

// List lists all ConcurrencyLimits in the indexer.
func (s *concurrencyLimitLister) List(selector labels.Selector) (ret []*v1alpha1.ConcurrencyLimit, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConcurrencyLimit))
	})
	return ret, err
}

// ConcurrencyLimits returns an object that can list and get ConcurrencyLimits.
func (s *concurrencyLimitLister) ConcurrencyLimits(namespace string) ConcurrencyLimitNamespaceLister {
	return concurrencyLimitNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConcurrencyLimitNamespaceLister helps list and get ConcurrencyLimits.
// All objects returned here must be treated as read-only.
type ConcurrencyLimitNamespaceLister interface {
	// List lists all ConcurrencyLimits in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConcurrencyLimit, err error)
	// Get retrieves the ConcurrencyLimit from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ConcurrencyLimit, error)
	ConcurrencyLimitNamespaceListerExpansion
}

// concurrencyLimitNamespaceLister implements the ConcurrencyLimitNamespaceLister
// interface.
type concurrencyLimitNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConcurrencyLimits in the indexer for a given namespace.
func (s concurrencyLimitNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ConcurrencyLimit, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConcurrencyLimit))
	})
	return ret, err
}

// Get retrieves the ConcurrencyLimit from the indexer for a given namespace and name.
func (s concurrencyLimitNamespaceLister) Get(name string) (*v1alpha1.ConcurrencyLimit, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("concurrencylimit"), name)
	}
	return obj.(*v1alpha1.ConcurrencyLimit), nil
}
