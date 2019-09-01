// Copyright (c) 2019 Tigera, Inc. All rights reserved.

// Code generated by informer-gen. DO NOT EDIT.

package v3

import (
	internalinterfaces "github.com/tigera/api/pkg/client/informers_generated/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// GlobalReportTypes returns a GlobalReportTypeInformer.
	GlobalReportTypes() GlobalReportTypeInformer
	// LicenseKeys returns a LicenseKeyInformer.
	LicenseKeys() LicenseKeyInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// GlobalReportTypes returns a GlobalReportTypeInformer.
func (v *version) GlobalReportTypes() GlobalReportTypeInformer {
	return &globalReportTypeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LicenseKeys returns a LicenseKeyInformer.
func (v *version) LicenseKeys() LicenseKeyInformer {
	return &licenseKeyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
