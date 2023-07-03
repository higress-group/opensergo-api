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

// RateLimitingLister helps list RateLimitings.
// All objects returned here must be treated as read-only.
type RateLimitingLister interface {
	// List lists all RateLimitings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RateLimiting, err error)
	// RateLimitings returns an object that can list and get RateLimitings.
	RateLimitings(namespace string) RateLimitingNamespaceLister
	RateLimitingListerExpansion
}

// rateLimitingLister implements the RateLimitingLister interface.
type rateLimitingLister struct {
	indexer cache.Indexer
}

// NewRateLimitingLister returns a new RateLimitingLister.
func NewRateLimitingLister(indexer cache.Indexer) RateLimitingLister {
	return &rateLimitingLister{indexer: indexer}
}

// List lists all RateLimitings in the indexer.
func (s *rateLimitingLister) List(selector labels.Selector) (ret []*v1alpha1.RateLimiting, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RateLimiting))
	})
	return ret, err
}

// RateLimitings returns an object that can list and get RateLimitings.
func (s *rateLimitingLister) RateLimitings(namespace string) RateLimitingNamespaceLister {
	return rateLimitingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RateLimitingNamespaceLister helps list and get RateLimitings.
// All objects returned here must be treated as read-only.
type RateLimitingNamespaceLister interface {
	// List lists all RateLimitings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RateLimiting, err error)
	// Get retrieves the RateLimiting from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RateLimiting, error)
	RateLimitingNamespaceListerExpansion
}

// rateLimitingNamespaceLister implements the RateLimitingNamespaceLister
// interface.
type rateLimitingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RateLimitings in the indexer for a given namespace.
func (s rateLimitingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RateLimiting, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RateLimiting))
	})
	return ret, err
}

// Get retrieves the RateLimiting from the indexer for a given namespace and name.
func (s rateLimitingNamespaceLister) Get(name string) (*v1alpha1.RateLimiting, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ratelimiting"), name)
	}
	return obj.(*v1alpha1.RateLimiting), nil
}
