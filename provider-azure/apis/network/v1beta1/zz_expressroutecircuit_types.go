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

type ExpressRouteCircuitObservation struct {

	// The ID of the ExpressRoute circuit.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ExpressRoute circuit provisioning state from your chosen service provider. Possible values are NotProvisioned, Provisioning, Provisioned, and Deprovisioning.
	ServiceProviderProvisioningState *string `json:"serviceProviderProvisioningState,omitempty" tf:"service_provider_provisioning_state,omitempty"`
}

type ExpressRouteCircuitParameters struct {

	// Allow the circuit to interact with classic  resources. Defaults to false.
	// +kubebuilder:validation:Optional
	AllowClassicOperations *bool `json:"allowClassicOperations,omitempty" tf:"allow_classic_operations,omitempty"`

	// The bandwidth in Gbps of the circuit being created on the Express Route Port.
	// +kubebuilder:validation:Optional
	BandwidthInGbps *float64 `json:"bandwidthInGbps,omitempty" tf:"bandwidth_in_gbps,omitempty"`

	// The bandwidth in Mbps of the circuit being created on the Service Provider.
	// +kubebuilder:validation:Optional
	BandwidthInMbps *float64 `json:"bandwidthInMbps,omitempty" tf:"bandwidth_in_mbps,omitempty"`

	// The ID of the Express Route Port this Express Route Circuit is based on.
	// +kubebuilder:validation:Optional
	ExpressRoutePortID *string `json:"expressRoutePortId,omitempty" tf:"express_route_port_id,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the peering location and not the Azure resource location. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PeeringLocation *string `json:"peeringLocation,omitempty" tf:"peering_location,omitempty"`

	// The name of the resource group in which to create the ExpressRoute circuit. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the ExpressRoute Service Provider. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ServiceProviderName *string `json:"serviceProviderName,omitempty" tf:"service_provider_name,omitempty"`

	// A sku block for the ExpressRoute circuit as documented below.
	// +kubebuilder:validation:Required
	Sku []ExpressRouteCircuitSkuParameters `json:"sku" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ExpressRouteCircuitSkuObservation struct {
}

type ExpressRouteCircuitSkuParameters struct {

	// The billing mode for bandwidth. Possible values are MeteredData or UnlimitedData.
	// +kubebuilder:validation:Required
	Family *string `json:"family" tf:"family,omitempty"`

	// The service tier. Possible values are Basic, Local, Standard or Premium.
	// +kubebuilder:validation:Required
	Tier *string `json:"tier" tf:"tier,omitempty"`
}

// ExpressRouteCircuitSpec defines the desired state of ExpressRouteCircuit
type ExpressRouteCircuitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExpressRouteCircuitParameters `json:"forProvider"`
}

// ExpressRouteCircuitStatus defines the observed state of ExpressRouteCircuit.
type ExpressRouteCircuitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExpressRouteCircuitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuit is the Schema for the ExpressRouteCircuits API. Manages an ExpressRoute circuit.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ExpressRouteCircuit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExpressRouteCircuitSpec   `json:"spec"`
	Status            ExpressRouteCircuitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpressRouteCircuitList contains a list of ExpressRouteCircuits
type ExpressRouteCircuitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpressRouteCircuit `json:"items"`
}

// Repository type metadata.
var (
	ExpressRouteCircuit_Kind             = "ExpressRouteCircuit"
	ExpressRouteCircuit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExpressRouteCircuit_Kind}.String()
	ExpressRouteCircuit_KindAPIVersion   = ExpressRouteCircuit_Kind + "." + CRDGroupVersion.String()
	ExpressRouteCircuit_GroupVersionKind = CRDGroupVersion.WithKind(ExpressRouteCircuit_Kind)
)

func init() {
	SchemeBuilder.Register(&ExpressRouteCircuit{}, &ExpressRouteCircuitList{})
}
