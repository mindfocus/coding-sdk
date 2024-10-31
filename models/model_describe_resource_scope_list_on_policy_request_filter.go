/*
CODING OPEN API

  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeResourceScopeListOnPolicyRequestFilter type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeResourceScopeListOnPolicyRequestFilter{}

// DescribeResourceScopeListOnPolicyRequestFilter 查询指定权限组的可用资源范围（分页）的查询条件
type DescribeResourceScopeListOnPolicyRequestFilter struct {
	// 资源类型（精确匹配）
	ResourceType utils.NullableString `json:"ResourceType,omitempty"`
}

// NewDescribeResourceScopeListOnPolicyRequestFilter instantiates a new DescribeResourceScopeListOnPolicyRequestFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeResourceScopeListOnPolicyRequestFilter() *DescribeResourceScopeListOnPolicyRequestFilter {
	this := DescribeResourceScopeListOnPolicyRequestFilter{}
	var resourceType string = ""
	this.ResourceType = *utils.NewNullableString(&resourceType)
	return &this
}

// NewDescribeResourceScopeListOnPolicyRequestFilterWithDefaults instantiates a new DescribeResourceScopeListOnPolicyRequestFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeResourceScopeListOnPolicyRequestFilterWithDefaults() *DescribeResourceScopeListOnPolicyRequestFilter {
	this := DescribeResourceScopeListOnPolicyRequestFilter{}
	var resourceType string = ""
	this.ResourceType = *utils.NewNullableString(&resourceType)
	return &this
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DescribeResourceScopeListOnPolicyRequestFilter) GetResourceType() string {
	if o == nil || utils.IsNil(o.ResourceType.Get()) {
		var ret string
		return ret
	}
	return *o.ResourceType.Get()
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DescribeResourceScopeListOnPolicyRequestFilter) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceType.Get(), o.ResourceType.IsSet()
}

// HasResourceType returns a boolean if a field has been set.
func (o *DescribeResourceScopeListOnPolicyRequestFilter) HasResourceType() bool {
	if o != nil && o.ResourceType.IsSet() {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given utils.NullableString and assigns it to the ResourceType field.
func (o *DescribeResourceScopeListOnPolicyRequestFilter) SetResourceType(v string) {
	o.ResourceType.Set(&v)
}
// SetResourceTypeNil sets the value for ResourceType to be an explicit nil
func (o *DescribeResourceScopeListOnPolicyRequestFilter) SetResourceTypeNil() {
	o.ResourceType.Set(nil)
}

// UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
func (o *DescribeResourceScopeListOnPolicyRequestFilter) UnsetResourceType() {
	o.ResourceType.Unset()
}

func (o DescribeResourceScopeListOnPolicyRequestFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeResourceScopeListOnPolicyRequestFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceType.IsSet() {
		toSerialize["ResourceType"] = o.ResourceType.Get()
	}
	return toSerialize, nil
}

type NullableDescribeResourceScopeListOnPolicyRequestFilter struct {
	value *DescribeResourceScopeListOnPolicyRequestFilter
	isSet bool
}

func (v NullableDescribeResourceScopeListOnPolicyRequestFilter) Get() *DescribeResourceScopeListOnPolicyRequestFilter {
	return v.value
}

func (v *NullableDescribeResourceScopeListOnPolicyRequestFilter) Set(val *DescribeResourceScopeListOnPolicyRequestFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeResourceScopeListOnPolicyRequestFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeResourceScopeListOnPolicyRequestFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeResourceScopeListOnPolicyRequestFilter(val *DescribeResourceScopeListOnPolicyRequestFilter) *NullableDescribeResourceScopeListOnPolicyRequestFilter {
	return &NullableDescribeResourceScopeListOnPolicyRequestFilter{value: val, isSet: true}
}

func (v NullableDescribeResourceScopeListOnPolicyRequestFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeResourceScopeListOnPolicyRequestFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

