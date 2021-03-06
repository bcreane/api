// Copyright (c) 2020 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LicenseKeyLister helps list LicenseKeys.
type LicenseKeyLister interface {
	// List lists all LicenseKeys in the indexer.
	List(selector labels.Selector) (ret []*v3.LicenseKey, err error)
	// LicenseKeys returns an object that can list and get LicenseKeys.
	LicenseKeys(namespace string) LicenseKeyNamespaceLister
	LicenseKeyListerExpansion
}

// licenseKeyLister implements the LicenseKeyLister interface.
type licenseKeyLister struct {
	indexer cache.Indexer
}

// NewLicenseKeyLister returns a new LicenseKeyLister.
func NewLicenseKeyLister(indexer cache.Indexer) LicenseKeyLister {
	return &licenseKeyLister{indexer: indexer}
}

// List lists all LicenseKeys in the indexer.
func (s *licenseKeyLister) List(selector labels.Selector) (ret []*v3.LicenseKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.LicenseKey))
	})
	return ret, err
}

// LicenseKeys returns an object that can list and get LicenseKeys.
func (s *licenseKeyLister) LicenseKeys(namespace string) LicenseKeyNamespaceLister {
	return licenseKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LicenseKeyNamespaceLister helps list and get LicenseKeys.
type LicenseKeyNamespaceLister interface {
	// List lists all LicenseKeys in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v3.LicenseKey, err error)
	// Get retrieves the LicenseKey from the indexer for a given namespace and name.
	Get(name string) (*v3.LicenseKey, error)
	LicenseKeyNamespaceListerExpansion
}

// licenseKeyNamespaceLister implements the LicenseKeyNamespaceLister
// interface.
type licenseKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LicenseKeys in the indexer for a given namespace.
func (s licenseKeyNamespaceLister) List(selector labels.Selector) (ret []*v3.LicenseKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.LicenseKey))
	})
	return ret, err
}

// Get retrieves the LicenseKey from the indexer for a given namespace and name.
func (s licenseKeyNamespaceLister) Get(name string) (*v3.LicenseKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("licensekey"), name)
	}
	return obj.(*v3.LicenseKey), nil
}
