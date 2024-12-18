/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeDepotSpecDetailRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeDepotSpecDetailRequest{}

// DescribeDepotSpecDetailRequest struct for DescribeDepotSpecDetailRequest
type DescribeDepotSpecDetailRequest struct {
	// 仓库路径（和仓库规范名字二选一选填）
	DepotPath *string `json:"DepotPath,omitempty"`
	// 仓库规范名字（和仓库路径二选一选填）
	DepotSpecName *string `json:"DepotSpecName,omitempty"`
}

// NewDescribeDepotSpecDetailRequest instantiates a new DescribeDepotSpecDetailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeDepotSpecDetailRequest() *DescribeDepotSpecDetailRequest {
	this := DescribeDepotSpecDetailRequest{}
	return &this
}

// NewDescribeDepotSpecDetailRequestWithDefaults instantiates a new DescribeDepotSpecDetailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeDepotSpecDetailRequestWithDefaults() *DescribeDepotSpecDetailRequest {
	this := DescribeDepotSpecDetailRequest{}
	return &this
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *DescribeDepotSpecDetailRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotSpecDetailRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *DescribeDepotSpecDetailRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *DescribeDepotSpecDetailRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

// GetDepotSpecName returns the DepotSpecName field value if set, zero value otherwise.
func (o *DescribeDepotSpecDetailRequest) GetDepotSpecName() string {
	if o == nil || utils.IsNil(o.DepotSpecName) {
		var ret string
		return ret
	}
	return *o.DepotSpecName
}

// GetDepotSpecNameOk returns a tuple with the DepotSpecName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotSpecDetailRequest) GetDepotSpecNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotSpecName) {
		return nil, false
	}
	return o.DepotSpecName, true
}

// HasDepotSpecName returns a boolean if a field has been set.
func (o *DescribeDepotSpecDetailRequest) HasDepotSpecName() bool {
	if o != nil && !utils.IsNil(o.DepotSpecName) {
		return true
	}

	return false
}

// SetDepotSpecName gets a reference to the given string and assigns it to the DepotSpecName field.
func (o *DescribeDepotSpecDetailRequest) SetDepotSpecName(v string) {
	o.DepotSpecName = &v
}

func (o DescribeDepotSpecDetailRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeDepotSpecDetailRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	if !utils.IsNil(o.DepotSpecName) {
		toSerialize["DepotSpecName"] = o.DepotSpecName
	}
	return toSerialize, nil
}

type NullableDescribeDepotSpecDetailRequest struct {
	value *DescribeDepotSpecDetailRequest
	isSet bool
}

func (v NullableDescribeDepotSpecDetailRequest) Get() *DescribeDepotSpecDetailRequest {
	return v.value
}

func (v *NullableDescribeDepotSpecDetailRequest) Set(val *DescribeDepotSpecDetailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeDepotSpecDetailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeDepotSpecDetailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeDepotSpecDetailRequest(val *DescribeDepotSpecDetailRequest) *NullableDescribeDepotSpecDetailRequest {
	return &NullableDescribeDepotSpecDetailRequest{value: val, isSet: true}
}

func (v NullableDescribeDepotSpecDetailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeDepotSpecDetailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


