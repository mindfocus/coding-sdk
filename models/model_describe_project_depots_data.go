/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeProjectDepotsData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeProjectDepotsData{}

// DescribeProjectDepotsData DescribeProjectDepots 返回数据结构
type DescribeProjectDepotsData struct {
	// 仓库信息列表
	DepotList []CodingCIProjectDepot `json:"DepotList,omitempty"`
	// 仓库类型是否被授权，如 Github 是否被授权
	IsBound *bool `json:"IsBound,omitempty"`
}

// NewDescribeProjectDepotsData instantiates a new DescribeProjectDepotsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeProjectDepotsData() *DescribeProjectDepotsData {
	this := DescribeProjectDepotsData{}
	var isBound bool = false
	this.IsBound = &isBound
	return &this
}

// NewDescribeProjectDepotsDataWithDefaults instantiates a new DescribeProjectDepotsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeProjectDepotsDataWithDefaults() *DescribeProjectDepotsData {
	this := DescribeProjectDepotsData{}
	var isBound bool = false
	this.IsBound = &isBound
	return &this
}

// GetDepotList returns the DepotList field value if set, zero value otherwise.
func (o *DescribeProjectDepotsData) GetDepotList() []CodingCIProjectDepot {
	if o == nil || utils.IsNil(o.DepotList) {
		var ret []CodingCIProjectDepot
		return ret
	}
	return o.DepotList
}

// GetDepotListOk returns a tuple with the DepotList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectDepotsData) GetDepotListOk() ([]CodingCIProjectDepot, bool) {
	if o == nil || utils.IsNil(o.DepotList) {
		return nil, false
	}
	return o.DepotList, true
}

// HasDepotList returns a boolean if a field has been set.
func (o *DescribeProjectDepotsData) HasDepotList() bool {
	if o != nil && !utils.IsNil(o.DepotList) {
		return true
	}

	return false
}

// SetDepotList gets a reference to the given []CodingCIProjectDepot and assigns it to the DepotList field.
func (o *DescribeProjectDepotsData) SetDepotList(v []CodingCIProjectDepot) {
	o.DepotList = v
}

// GetIsBound returns the IsBound field value if set, zero value otherwise.
func (o *DescribeProjectDepotsData) GetIsBound() bool {
	if o == nil || utils.IsNil(o.IsBound) {
		var ret bool
		return ret
	}
	return *o.IsBound
}

// GetIsBoundOk returns a tuple with the IsBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectDepotsData) GetIsBoundOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsBound) {
		return nil, false
	}
	return o.IsBound, true
}

// HasIsBound returns a boolean if a field has been set.
func (o *DescribeProjectDepotsData) HasIsBound() bool {
	if o != nil && !utils.IsNil(o.IsBound) {
		return true
	}

	return false
}

// SetIsBound gets a reference to the given bool and assigns it to the IsBound field.
func (o *DescribeProjectDepotsData) SetIsBound(v bool) {
	o.IsBound = &v
}

func (o DescribeProjectDepotsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeProjectDepotsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.DepotList) {
		toSerialize["DepotList"] = o.DepotList
	}
	if !utils.IsNil(o.IsBound) {
		toSerialize["IsBound"] = o.IsBound
	}
	return toSerialize, nil
}

type NullableDescribeProjectDepotsData struct {
	value *DescribeProjectDepotsData
	isSet bool
}

func (v NullableDescribeProjectDepotsData) Get() *DescribeProjectDepotsData {
	return v.value
}

func (v *NullableDescribeProjectDepotsData) Set(val *DescribeProjectDepotsData) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeProjectDepotsData) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeProjectDepotsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeProjectDepotsData(val *DescribeProjectDepotsData) *NullableDescribeProjectDepotsData {
	return &NullableDescribeProjectDepotsData{value: val, isSet: true}
}

func (v NullableDescribeProjectDepotsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeProjectDepotsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

