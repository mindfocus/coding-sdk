/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the PolicyResourceScopeInfo type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PolicyResourceScopeInfo{}

// PolicyResourceScopeInfo 权限组的可用资源范围信息
type PolicyResourceScopeInfo struct {
	// 创建时间
	CreatedAt *int32 `json:"CreatedAt,omitempty"`
	// 权限组 ID
	PolicyId *int64 `json:"PolicyId,omitempty"`
	Resource *ResourceInfoOfPolicyScope `json:"Resource,omitempty"`
}

// NewPolicyResourceScopeInfo instantiates a new PolicyResourceScopeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyResourceScopeInfo() *PolicyResourceScopeInfo {
	this := PolicyResourceScopeInfo{}
	return &this
}

// NewPolicyResourceScopeInfoWithDefaults instantiates a new PolicyResourceScopeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyResourceScopeInfoWithDefaults() *PolicyResourceScopeInfo {
	this := PolicyResourceScopeInfo{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PolicyResourceScopeInfo) GetCreatedAt() int32 {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResourceScopeInfo) GetCreatedAtOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PolicyResourceScopeInfo) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *PolicyResourceScopeInfo) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *PolicyResourceScopeInfo) GetPolicyId() int64 {
	if o == nil || utils.IsNil(o.PolicyId) {
		var ret int64
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResourceScopeInfo) GetPolicyIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *PolicyResourceScopeInfo) HasPolicyId() bool {
	if o != nil && !utils.IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given int64 and assigns it to the PolicyId field.
func (o *PolicyResourceScopeInfo) SetPolicyId(v int64) {
	o.PolicyId = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *PolicyResourceScopeInfo) GetResource() ResourceInfoOfPolicyScope {
	if o == nil || utils.IsNil(o.Resource) {
		var ret ResourceInfoOfPolicyScope
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyResourceScopeInfo) GetResourceOk() (*ResourceInfoOfPolicyScope, bool) {
	if o == nil || utils.IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *PolicyResourceScopeInfo) HasResource() bool {
	if o != nil && !utils.IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given ResourceInfoOfPolicyScope and assigns it to the Resource field.
func (o *PolicyResourceScopeInfo) SetResource(v ResourceInfoOfPolicyScope) {
	o.Resource = &v
}

func (o PolicyResourceScopeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyResourceScopeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if !utils.IsNil(o.PolicyId) {
		toSerialize["PolicyId"] = o.PolicyId
	}
	if !utils.IsNil(o.Resource) {
		toSerialize["Resource"] = o.Resource
	}
	return toSerialize, nil
}

type NullablePolicyResourceScopeInfo struct {
	value *PolicyResourceScopeInfo
	isSet bool
}

func (v NullablePolicyResourceScopeInfo) Get() *PolicyResourceScopeInfo {
	return v.value
}

func (v *NullablePolicyResourceScopeInfo) Set(val *PolicyResourceScopeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyResourceScopeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyResourceScopeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyResourceScopeInfo(val *PolicyResourceScopeInfo) *NullablePolicyResourceScopeInfo {
	return &NullablePolicyResourceScopeInfo{value: val, isSet: true}
}

func (v NullablePolicyResourceScopeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyResourceScopeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


