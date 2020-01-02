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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "apisix-customize-controller/pkg/apis/config/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApisixRouteLister helps list ApisixRoutes.
type ApisixRouteLister interface {
	// List lists all ApisixRoutes in the indexer.
	List(selector labels.Selector) (ret []*v1.ApisixRoute, err error)
	// ApisixRoutes returns an object that can list and get ApisixRoutes.
	ApisixRoutes(namespace string) ApisixRouteNamespaceLister
	ApisixRouteListerExpansion
}

// apisixRouteLister implements the ApisixRouteLister interface.
type apisixRouteLister struct {
	indexer cache.Indexer
}

// NewApisixRouteLister returns a new ApisixRouteLister.
func NewApisixRouteLister(indexer cache.Indexer) ApisixRouteLister {
	return &apisixRouteLister{indexer: indexer}
}

// List lists all ApisixRoutes in the indexer.
func (s *apisixRouteLister) List(selector labels.Selector) (ret []*v1.ApisixRoute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ApisixRoute))
	})
	return ret, err
}

// ApisixRoutes returns an object that can list and get ApisixRoutes.
func (s *apisixRouteLister) ApisixRoutes(namespace string) ApisixRouteNamespaceLister {
	return apisixRouteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApisixRouteNamespaceLister helps list and get ApisixRoutes.
type ApisixRouteNamespaceLister interface {
	// List lists all ApisixRoutes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ApisixRoute, err error)
	// Get retrieves the ApisixRoute from the indexer for a given namespace and name.
	Get(name string) (*v1.ApisixRoute, error)
	ApisixRouteNamespaceListerExpansion
}

// apisixRouteNamespaceLister implements the ApisixRouteNamespaceLister
// interface.
type apisixRouteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApisixRoutes in the indexer for a given namespace.
func (s apisixRouteNamespaceLister) List(selector labels.Selector) (ret []*v1.ApisixRoute, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ApisixRoute))
	})
	return ret, err
}

// Get retrieves the ApisixRoute from the indexer for a given namespace and name.
func (s apisixRouteNamespaceLister) Get(name string) (*v1.ApisixRoute, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("apisixroute"), name)
	}
	return obj.(*v1.ApisixRoute), nil
}
