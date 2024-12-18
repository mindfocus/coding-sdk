/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCdHostServerGroupsResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCdHostServerGroupsResponseData{}

// DescribeCdHostServerGroupsResponseData DescribeCdHostServerGroupsResponseData 结构
type DescribeCdHostServerGroupsResponseData struct {
	// 主机组列表
	HostServerGroups []HostServerGroup `json:"HostServerGroups,omitempty"`
	// 页码
	PageNumber *int64 `json:"PageNumber,omitempty"`
	// 每页条数
	PageSize *int64 `json:"PageSize,omitempty"`
	// 总共页数
	TotalPage *int64 `json:"TotalPage,omitempty"`
	// 总共条数
	TotalRow *int64 `json:"TotalRow,omitempty"`
}

// NewDescribeCdHostServerGroupsResponseData instantiates a new DescribeCdHostServerGroupsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCdHostServerGroupsResponseData() *DescribeCdHostServerGroupsResponseData {
	this := DescribeCdHostServerGroupsResponseData{}
	return &this
}

// NewDescribeCdHostServerGroupsResponseDataWithDefaults instantiates a new DescribeCdHostServerGroupsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCdHostServerGroupsResponseDataWithDefaults() *DescribeCdHostServerGroupsResponseData {
	this := DescribeCdHostServerGroupsResponseData{}
	return &this
}

// GetHostServerGroups returns the HostServerGroups field value if set, zero value otherwise.
func (o *DescribeCdHostServerGroupsResponseData) GetHostServerGroups() []HostServerGroup {
	if o == nil || utils.IsNil(o.HostServerGroups) {
		var ret []HostServerGroup
		return ret
	}
	return o.HostServerGroups
}

// GetHostServerGroupsOk returns a tuple with the HostServerGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdHostServerGroupsResponseData) GetHostServerGroupsOk() ([]HostServerGroup, bool) {
	if o == nil || utils.IsNil(o.HostServerGroups) {
		return nil, false
	}
	return o.HostServerGroups, true
}

// HasHostServerGroups returns a boolean if a field has been set.
func (o *DescribeCdHostServerGroupsResponseData) HasHostServerGroups() bool {
	if o != nil && !utils.IsNil(o.HostServerGroups) {
		return true
	}

	return false
}

// SetHostServerGroups gets a reference to the given []HostServerGroup and assigns it to the HostServerGroups field.
func (o *DescribeCdHostServerGroupsResponseData) SetHostServerGroups(v []HostServerGroup) {
	o.HostServerGroups = v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *DescribeCdHostServerGroupsResponseData) GetPageNumber() int64 {
	if o == nil || utils.IsNil(o.PageNumber) {
		var ret int64
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdHostServerGroupsResponseData) GetPageNumberOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *DescribeCdHostServerGroupsResponseData) HasPageNumber() bool {
	if o != nil && !utils.IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int64 and assigns it to the PageNumber field.
func (o *DescribeCdHostServerGroupsResponseData) SetPageNumber(v int64) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *DescribeCdHostServerGroupsResponseData) GetPageSize() int64 {
	if o == nil || utils.IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdHostServerGroupsResponseData) GetPageSizeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *DescribeCdHostServerGroupsResponseData) HasPageSize() bool {
	if o != nil && !utils.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *DescribeCdHostServerGroupsResponseData) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetTotalPage returns the TotalPage field value if set, zero value otherwise.
func (o *DescribeCdHostServerGroupsResponseData) GetTotalPage() int64 {
	if o == nil || utils.IsNil(o.TotalPage) {
		var ret int64
		return ret
	}
	return *o.TotalPage
}

// GetTotalPageOk returns a tuple with the TotalPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdHostServerGroupsResponseData) GetTotalPageOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.TotalPage) {
		return nil, false
	}
	return o.TotalPage, true
}

// HasTotalPage returns a boolean if a field has been set.
func (o *DescribeCdHostServerGroupsResponseData) HasTotalPage() bool {
	if o != nil && !utils.IsNil(o.TotalPage) {
		return true
	}

	return false
}

// SetTotalPage gets a reference to the given int64 and assigns it to the TotalPage field.
func (o *DescribeCdHostServerGroupsResponseData) SetTotalPage(v int64) {
	o.TotalPage = &v
}

// GetTotalRow returns the TotalRow field value if set, zero value otherwise.
func (o *DescribeCdHostServerGroupsResponseData) GetTotalRow() int64 {
	if o == nil || utils.IsNil(o.TotalRow) {
		var ret int64
		return ret
	}
	return *o.TotalRow
}

// GetTotalRowOk returns a tuple with the TotalRow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdHostServerGroupsResponseData) GetTotalRowOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.TotalRow) {
		return nil, false
	}
	return o.TotalRow, true
}

// HasTotalRow returns a boolean if a field has been set.
func (o *DescribeCdHostServerGroupsResponseData) HasTotalRow() bool {
	if o != nil && !utils.IsNil(o.TotalRow) {
		return true
	}

	return false
}

// SetTotalRow gets a reference to the given int64 and assigns it to the TotalRow field.
func (o *DescribeCdHostServerGroupsResponseData) SetTotalRow(v int64) {
	o.TotalRow = &v
}

func (o DescribeCdHostServerGroupsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCdHostServerGroupsResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.HostServerGroups) {
		toSerialize["HostServerGroups"] = o.HostServerGroups
	}
	if !utils.IsNil(o.PageNumber) {
		toSerialize["PageNumber"] = o.PageNumber
	}
	if !utils.IsNil(o.PageSize) {
		toSerialize["PageSize"] = o.PageSize
	}
	if !utils.IsNil(o.TotalPage) {
		toSerialize["TotalPage"] = o.TotalPage
	}
	if !utils.IsNil(o.TotalRow) {
		toSerialize["TotalRow"] = o.TotalRow
	}
	return toSerialize, nil
}

type NullableDescribeCdHostServerGroupsResponseData struct {
	value *DescribeCdHostServerGroupsResponseData
	isSet bool
}

func (v NullableDescribeCdHostServerGroupsResponseData) Get() *DescribeCdHostServerGroupsResponseData {
	return v.value
}

func (v *NullableDescribeCdHostServerGroupsResponseData) Set(val *DescribeCdHostServerGroupsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCdHostServerGroupsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCdHostServerGroupsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCdHostServerGroupsResponseData(val *DescribeCdHostServerGroupsResponseData) *NullableDescribeCdHostServerGroupsResponseData {
	return &NullableDescribeCdHostServerGroupsResponseData{value: val, isSet: true}
}

func (v NullableDescribeCdHostServerGroupsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCdHostServerGroupsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


