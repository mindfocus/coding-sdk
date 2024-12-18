/*
CODING OPEN API

  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the Policy type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Policy{}

// Policy 权限信息
type Policy struct {
	// 项目成员
	PolicyAlias *string `json:"PolicyAlias,omitempty"`
	// 1
	PolicyId *int64 `json:"PolicyId,omitempty"`
	// ProjectMember类型
	PolicyName *string `json:"PolicyName,omitempty"`
}

// NewPolicy instantiates a new Policy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicy() *Policy {
	this := Policy{}
	var policyAlias string = ""
	this.PolicyAlias = &policyAlias
	var policyName string = ""
	this.PolicyName = &policyName
	return &this
}

// NewPolicyWithDefaults instantiates a new Policy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyWithDefaults() *Policy {
	this := Policy{}
	var policyAlias string = ""
	this.PolicyAlias = &policyAlias
	var policyName string = ""
	this.PolicyName = &policyName
	return &this
}

// GetPolicyAlias returns the PolicyAlias field value if set, zero value otherwise.
func (o *Policy) GetPolicyAlias() string {
	if o == nil || utils.IsNil(o.PolicyAlias) {
		var ret string
		return ret
	}
	return *o.PolicyAlias
}

// GetPolicyAliasOk returns a tuple with the PolicyAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetPolicyAliasOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PolicyAlias) {
		return nil, false
	}
	return o.PolicyAlias, true
}

// HasPolicyAlias returns a boolean if a field has been set.
func (o *Policy) HasPolicyAlias() bool {
	if o != nil && !utils.IsNil(o.PolicyAlias) {
		return true
	}

	return false
}

// SetPolicyAlias gets a reference to the given string and assigns it to the PolicyAlias field.
func (o *Policy) SetPolicyAlias(v string) {
	o.PolicyAlias = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *Policy) GetPolicyId() int64 {
	if o == nil || utils.IsNil(o.PolicyId) {
		var ret int64
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetPolicyIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *Policy) HasPolicyId() bool {
	if o != nil && !utils.IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given int64 and assigns it to the PolicyId field.
func (o *Policy) SetPolicyId(v int64) {
	o.PolicyId = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *Policy) GetPolicyName() string {
	if o == nil || utils.IsNil(o.PolicyName) {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetPolicyNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PolicyName) {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *Policy) HasPolicyName() bool {
	if o != nil && !utils.IsNil(o.PolicyName) {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *Policy) SetPolicyName(v string) {
	o.PolicyName = &v
}

func (o Policy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Policy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.PolicyAlias) {
		toSerialize["PolicyAlias"] = o.PolicyAlias
	}
	if !utils.IsNil(o.PolicyId) {
		toSerialize["PolicyId"] = o.PolicyId
	}
	if !utils.IsNil(o.PolicyName) {
		toSerialize["PolicyName"] = o.PolicyName
	}
	return toSerialize, nil
}

type NullablePolicy struct {
	value *Policy
	isSet bool
}

func (v NullablePolicy) Get() *Policy {
	return v.value
}

func (v *NullablePolicy) Set(val *Policy) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicy(val *Policy) *NullablePolicy {
	return &NullablePolicy{value: val, isSet: true}
}

func (v NullablePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


