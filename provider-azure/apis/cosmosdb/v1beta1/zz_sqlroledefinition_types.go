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

type PermissionsObservation struct {
}

type PermissionsParameters struct {

	// A list of data actions that are allowed for the Cosmos DB SQL Role Definition.
	// +kubebuilder:validation:Required
	DataActions []*string `json:"dataActions" tf:"data_actions,omitempty"`
}

type SQLRoleDefinitionObservation struct {

	// The ID of the Cosmos DB SQL Role Definition.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SQLRoleDefinitionParameters struct {

	// The name of the Cosmos DB Account. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/cosmosdb/v1beta1.Account
	// +kubebuilder:validation:Optional
	AccountName *string `json:"accountName,omitempty" tf:"account_name,omitempty"`

	// +kubebuilder:validation:Optional
	AccountNameRef *v1.Reference `json:"accountNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AccountNameSelector *v1.Selector `json:"accountNameSelector,omitempty" tf:"-"`

	// A list of fully qualified scopes at or below which Role Assignments may be created using this Cosmos DB SQL Role Definition. It will allow application of this Cosmos DB SQL Role Definition on the entire Database Account or any underlying Database/Collection. Scopes higher than Database Account are not enforceable as assignable scopes.
	// +kubebuilder:validation:Required
	AssignableScopes []*string `json:"assignableScopes" tf:"assignable_scopes,omitempty"`

	// A permissions block as defined below.
	// +kubebuilder:validation:Required
	Permissions []PermissionsParameters `json:"permissions" tf:"permissions,omitempty"`

	// The name of the Resource Group in which the Cosmos DB SQL Role Definition is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The GUID as the name of the Cosmos DB SQL Role Definition - one will be generated if not specified. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty" tf:"role_definition_id,omitempty"`

	// The type of the Cosmos DB SQL Role Definition. Possible values are BuiltInRole and CustomRole. Defaults to CustomRole. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// SQLRoleDefinitionSpec defines the desired state of SQLRoleDefinition
type SQLRoleDefinitionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLRoleDefinitionParameters `json:"forProvider"`
}

// SQLRoleDefinitionStatus defines the observed state of SQLRoleDefinition.
type SQLRoleDefinitionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLRoleDefinitionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SQLRoleDefinition is the Schema for the SQLRoleDefinitions API. Manages a Cosmos DB SQL Role Definition.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLRoleDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SQLRoleDefinitionSpec   `json:"spec"`
	Status            SQLRoleDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLRoleDefinitionList contains a list of SQLRoleDefinitions
type SQLRoleDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLRoleDefinition `json:"items"`
}

// Repository type metadata.
var (
	SQLRoleDefinition_Kind             = "SQLRoleDefinition"
	SQLRoleDefinition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLRoleDefinition_Kind}.String()
	SQLRoleDefinition_KindAPIVersion   = SQLRoleDefinition_Kind + "." + CRDGroupVersion.String()
	SQLRoleDefinition_GroupVersionKind = CRDGroupVersion.WithKind(SQLRoleDefinition_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLRoleDefinition{}, &SQLRoleDefinitionList{})
}
