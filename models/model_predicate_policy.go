/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the PredicatePolicy type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PredicatePolicy{}

// PredicatePolicy 资源权限判定策略信息
type PredicatePolicy struct {
	// 是否存在admin授权
	HasAdministrator *bool `json:"HasAdministrator,omitempty"`
	//   SELF_PARENT  // 同时使用父级资源+当前资源   SELF_NONE  // 只使用当前资源   NONE_PARENT  // 只使用父级资源
	ResourcePredicatePolicy *string `json:"ResourcePredicatePolicy,omitempty"`
}

// NewPredicatePolicy instantiates a new PredicatePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredicatePolicy() *PredicatePolicy {
	this := PredicatePolicy{}
	var hasAdministrator bool = false
	this.HasAdministrator = &hasAdministrator
	var resourcePredicatePolicy string = ""
	this.ResourcePredicatePolicy = &resourcePredicatePolicy
	return &this
}

// NewPredicatePolicyWithDefaults instantiates a new PredicatePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredicatePolicyWithDefaults() *PredicatePolicy {
	this := PredicatePolicy{}
	var hasAdministrator bool = false
	this.HasAdministrator = &hasAdministrator
	var resourcePredicatePolicy string = ""
	this.ResourcePredicatePolicy = &resourcePredicatePolicy
	return &this
}

// GetHasAdministrator returns the HasAdministrator field value if set, zero value otherwise.
func (o *PredicatePolicy) GetHasAdministrator() bool {
	if o == nil || utils.IsNil(o.HasAdministrator) {
		var ret bool
		return ret
	}
	return *o.HasAdministrator
}

// GetHasAdministratorOk returns a tuple with the HasAdministrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredicatePolicy) GetHasAdministratorOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.HasAdministrator) {
		return nil, false
	}
	return o.HasAdministrator, true
}

// HasHasAdministrator returns a boolean if a field has been set.
func (o *PredicatePolicy) HasHasAdministrator() bool {
	if o != nil && !utils.IsNil(o.HasAdministrator) {
		return true
	}

	return false
}

// SetHasAdministrator gets a reference to the given bool and assigns it to the HasAdministrator field.
func (o *PredicatePolicy) SetHasAdministrator(v bool) {
	o.HasAdministrator = &v
}

// GetResourcePredicatePolicy returns the ResourcePredicatePolicy field value if set, zero value otherwise.
func (o *PredicatePolicy) GetResourcePredicatePolicy() string {
	if o == nil || utils.IsNil(o.ResourcePredicatePolicy) {
		var ret string
		return ret
	}
	return *o.ResourcePredicatePolicy
}

// GetResourcePredicatePolicyOk returns a tuple with the ResourcePredicatePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredicatePolicy) GetResourcePredicatePolicyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ResourcePredicatePolicy) {
		return nil, false
	}
	return o.ResourcePredicatePolicy, true
}

// HasResourcePredicatePolicy returns a boolean if a field has been set.
func (o *PredicatePolicy) HasResourcePredicatePolicy() bool {
	if o != nil && !utils.IsNil(o.ResourcePredicatePolicy) {
		return true
	}

	return false
}

// SetResourcePredicatePolicy gets a reference to the given string and assigns it to the ResourcePredicatePolicy field.
func (o *PredicatePolicy) SetResourcePredicatePolicy(v string) {
	o.ResourcePredicatePolicy = &v
}

func (o PredicatePolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PredicatePolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.HasAdministrator) {
		toSerialize["HasAdministrator"] = o.HasAdministrator
	}
	if !utils.IsNil(o.ResourcePredicatePolicy) {
		toSerialize["ResourcePredicatePolicy"] = o.ResourcePredicatePolicy
	}
	return toSerialize, nil
}

type NullablePredicatePolicy struct {
	value *PredicatePolicy
	isSet bool
}

func (v NullablePredicatePolicy) Get() *PredicatePolicy {
	return v.value
}

func (v *NullablePredicatePolicy) Set(val *PredicatePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullablePredicatePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePredicatePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredicatePolicy(val *PredicatePolicy) *NullablePredicatePolicy {
	return &NullablePredicatePolicy{value: val, isSet: true}
}

func (v NullablePredicatePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredicatePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


