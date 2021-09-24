/*
Copyright 2021 TriggerMesh Inc.

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

package v1alpha1

import (
	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/transformation/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TransformationLister helps list Transformations.
// All objects returned here must be treated as read-only.
type TransformationLister interface {
	// List lists all Transformations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Transformation, err error)
	// Transformations returns an object that can list and get Transformations.
	Transformations(namespace string) TransformationNamespaceLister
	TransformationListerExpansion
}

// transformationLister implements the TransformationLister interface.
type transformationLister struct {
	indexer cache.Indexer
}

// NewTransformationLister returns a new TransformationLister.
func NewTransformationLister(indexer cache.Indexer) TransformationLister {
	return &transformationLister{indexer: indexer}
}

// List lists all Transformations in the indexer.
func (s *transformationLister) List(selector labels.Selector) (ret []*v1alpha1.Transformation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Transformation))
	})
	return ret, err
}

// Transformations returns an object that can list and get Transformations.
func (s *transformationLister) Transformations(namespace string) TransformationNamespaceLister {
	return transformationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TransformationNamespaceLister helps list and get Transformations.
// All objects returned here must be treated as read-only.
type TransformationNamespaceLister interface {
	// List lists all Transformations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Transformation, err error)
	// Get retrieves the Transformation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Transformation, error)
	TransformationNamespaceListerExpansion
}

// transformationNamespaceLister implements the TransformationNamespaceLister
// interface.
type transformationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Transformations in the indexer for a given namespace.
func (s transformationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Transformation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Transformation))
	})
	return ret, err
}

// Get retrieves the Transformation from the indexer for a given namespace and name.
func (s transformationNamespaceLister) Get(name string) (*v1alpha1.Transformation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("transformation"), name)
	}
	return obj.(*v1alpha1.Transformation), nil
}