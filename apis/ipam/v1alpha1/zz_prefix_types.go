/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PrefixInitParameters struct {

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean)
	IsPool *bool `json:"isPool,omitempty" tf:"is_pool,omitempty"`

	// (Boolean)
	MarkUtilized *bool `json:"markUtilized,omitempty" tf:"mark_utilized,omitempty"`

	// (String)
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// (Number)
	RoleID *float64 `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// (Number)
	SiteID *float64 `json:"siteId,omitempty" tf:"site_id,omitempty"`

	// (String) Valid values are active, container, reserved and deprecated.
	// Valid values are `active`, `container`, `reserved` and `deprecated`.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (Set of String)
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Number)
	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// (Number)
	VlanID *float64 `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`

	// (Number)
	VrfID *float64 `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`
}

type PrefixObservation struct {

	// (String)
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean)
	IsPool *bool `json:"isPool,omitempty" tf:"is_pool,omitempty"`

	// (Boolean)
	MarkUtilized *bool `json:"markUtilized,omitempty" tf:"mark_utilized,omitempty"`

	// (String)
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// (Number)
	RoleID *float64 `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// (Number)
	SiteID *float64 `json:"siteId,omitempty" tf:"site_id,omitempty"`

	// (String) Valid values are active, container, reserved and deprecated.
	// Valid values are `active`, `container`, `reserved` and `deprecated`.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (Set of String)
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Number)
	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// (Number)
	VlanID *float64 `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`

	// (Number)
	VrfID *float64 `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`
}

type PrefixParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean)
	// +kubebuilder:validation:Optional
	IsPool *bool `json:"isPool,omitempty" tf:"is_pool,omitempty"`

	// (Boolean)
	// +kubebuilder:validation:Optional
	MarkUtilized *bool `json:"markUtilized,omitempty" tf:"mark_utilized,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	RoleID *float64 `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	SiteID *float64 `json:"siteId,omitempty" tf:"site_id,omitempty"`

	// (String) Valid values are active, container, reserved and deprecated.
	// Valid values are `active`, `container`, `reserved` and `deprecated`.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (Set of String)
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	VlanID *float64 `json:"vlanId,omitempty" tf:"vlan_id,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	VrfID *float64 `json:"vrfId,omitempty" tf:"vrf_id,omitempty"`
}

// PrefixSpec defines the desired state of Prefix
type PrefixSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrefixParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider PrefixInitParameters `json:"initProvider,omitempty"`
}

// PrefixStatus defines the observed state of Prefix.
type PrefixStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrefixObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Prefix is the Schema for the Prefixs API. From the official documentation https://docs.netbox.dev/en/stable/features/ipam/#prefixes: A prefix is an IPv4 or IPv6 network and mask expressed in CIDR notation (e.g. 192.0.2.0/24). A prefix entails only the "network portion" of an IP address: All bits in the address not covered by the mask must be zero. (In other words, a prefix cannot be a specific IP address.) Prefixes are automatically organized by their parent aggregates. Additionally, each prefix can be assigned to a particular site and virtual routing and forwarding instance (VRF). Each VRF represents a separate IP space or routing table. All prefixes not assigned to a VRF are considered to be in the "global" table.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netbox}
type Prefix struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.prefix) || has(self.initProvider.prefix)",message="prefix is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.status) || has(self.initProvider.status)",message="status is a required parameter"
	Spec   PrefixSpec   `json:"spec"`
	Status PrefixStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrefixList contains a list of Prefixs
type PrefixList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Prefix `json:"items"`
}

// Repository type metadata.
var (
	Prefix_Kind             = "Prefix"
	Prefix_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Prefix_Kind}.String()
	Prefix_KindAPIVersion   = Prefix_Kind + "." + CRDGroupVersion.String()
	Prefix_GroupVersionKind = CRDGroupVersion.WithKind(Prefix_Kind)
)

func init() {
	SchemeBuilder.Register(&Prefix{}, &PrefixList{})
}
