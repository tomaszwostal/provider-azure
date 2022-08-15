/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SpatialAnchorsAccountObservation struct {

	// The domain of the Spatial Anchors Account.
	AccountDomain *string `json:"accountDomain,omitempty" tf:"account_domain,omitempty"`

	// The account ID of the Spatial Anchors Account.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// The ID of the Spatial Anchors Account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SpatialAnchorsAccountParameters struct {

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the resource group in which to create the Spatial Anchors Account.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SpatialAnchorsAccountSpec defines the desired state of SpatialAnchorsAccount
type SpatialAnchorsAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpatialAnchorsAccountParameters `json:"forProvider"`
}

// SpatialAnchorsAccountStatus defines the observed state of SpatialAnchorsAccount.
type SpatialAnchorsAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpatialAnchorsAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpatialAnchorsAccount is the Schema for the SpatialAnchorsAccounts API. Manages an Azure Spatial Anchors Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpatialAnchorsAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpatialAnchorsAccountSpec   `json:"spec"`
	Status            SpatialAnchorsAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpatialAnchorsAccountList contains a list of SpatialAnchorsAccounts
type SpatialAnchorsAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpatialAnchorsAccount `json:"items"`
}

// Repository type metadata.
var (
	SpatialAnchorsAccount_Kind             = "SpatialAnchorsAccount"
	SpatialAnchorsAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpatialAnchorsAccount_Kind}.String()
	SpatialAnchorsAccount_KindAPIVersion   = SpatialAnchorsAccount_Kind + "." + CRDGroupVersion.String()
	SpatialAnchorsAccount_GroupVersionKind = CRDGroupVersion.WithKind(SpatialAnchorsAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&SpatialAnchorsAccount{}, &SpatialAnchorsAccountList{})
}
