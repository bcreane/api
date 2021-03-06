// Copyright (c) 2020 Tigera, Inc. All rights reserved.

// Code generated by informer-gen. DO NOT EDIT.

package v3

import (
	internalinterfaces "github.com/tigera/api/pkg/client/informers_generated/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// GlobalAlerts returns a GlobalAlertInformer.
	GlobalAlerts() GlobalAlertInformer
	// GlobalAlertTemplates returns a GlobalAlertTemplateInformer.
	GlobalAlertTemplates() GlobalAlertTemplateInformer
	// GlobalReportTypes returns a GlobalReportTypeInformer.
	GlobalReportTypes() GlobalReportTypeInformer
	// LicenseKeys returns a LicenseKeyInformer.
	LicenseKeys() LicenseKeyInformer
	// ManagedClusters returns a ManagedClusterInformer.
	ManagedClusters() ManagedClusterInformer
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

// GlobalAlerts returns a GlobalAlertInformer.
func (v *version) GlobalAlerts() GlobalAlertInformer {
	return &globalAlertInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// GlobalAlertTemplates returns a GlobalAlertTemplateInformer.
func (v *version) GlobalAlertTemplates() GlobalAlertTemplateInformer {
	return &globalAlertTemplateInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// GlobalReportTypes returns a GlobalReportTypeInformer.
func (v *version) GlobalReportTypes() GlobalReportTypeInformer {
	return &globalReportTypeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LicenseKeys returns a LicenseKeyInformer.
func (v *version) LicenseKeys() LicenseKeyInformer {
	return &licenseKeyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagedClusters returns a ManagedClusterInformer.
func (v *version) ManagedClusters() ManagedClusterInformer {
	return &managedClusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
