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

// RateLimitLister helps list RateLimits.
// All objects returned here must be treated as read-only.
type RateLimitLister interface {
	// List lists all RateLimits in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RateLimit, err error)
	// RateLimits returns an object that can list and get RateLimits.
	RateLimits(namespace string) RateLimitNamespaceLister
	RateLimitListerExpansion
}

// rateLimitLister implements the RateLimitLister interface.
type rateLimitLister struct {
	indexer cache.Indexer
}

// NewRateLimitLister returns a new RateLimitLister.
func NewRateLimitLister(indexer cache.Indexer) RateLimitLister {
	return &rateLimitLister{indexer: indexer}
}

// List lists all RateLimits in the indexer.
func (s *rateLimitLister) List(selector labels.Selector) (ret []*v1alpha1.RateLimit, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RateLimit))
	})
	return ret, err
}

// RateLimits returns an object that can list and get RateLimits.
func (s *rateLimitLister) RateLimits(namespace string) RateLimitNamespaceLister {
	return rateLimitNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RateLimitNamespaceLister helps list and get RateLimits.
// All objects returned here must be treated as read-only.
type RateLimitNamespaceLister interface {
	// List lists all RateLimits in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RateLimit, err error)
	// Get retrieves the RateLimit from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RateLimit, error)
	RateLimitNamespaceListerExpansion
}

// rateLimitNamespaceLister implements the RateLimitNamespaceLister
// interface.
type rateLimitNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RateLimits in the indexer for a given namespace.
func (s rateLimitNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RateLimit, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RateLimit))
	})
	return ret, err
}

// Get retrieves the RateLimit from the indexer for a given namespace and name.
func (s rateLimitNamespaceLister) Get(name string) (*v1alpha1.RateLimit, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ratelimit"), name)
	}
	return obj.(*v1alpha1.RateLimit), nil
}
