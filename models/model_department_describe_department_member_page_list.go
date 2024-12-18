/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"bytes"
	"fmt"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DepartmentDescribeDepartmentMemberPageList type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DepartmentDescribeDepartmentMemberPageList{}

// DepartmentDescribeDepartmentMemberPageList 部门成员分页数据
type DepartmentDescribeDepartmentMemberPageList struct {
	// 部门成员
	DepartmentMembers []DepartmentDepartmentMembersData `json:"DepartmentMembers"`
	// 页数
	PageNumber int64 `json:"PageNumber"`
	// 每页数量
	PageSize int64 `json:"PageSize"`
	// 总数
	TotalCount int64 `json:"TotalCount"`
}

type _DepartmentDescribeDepartmentMemberPageList DepartmentDescribeDepartmentMemberPageList

// NewDepartmentDescribeDepartmentMemberPageList instantiates a new DepartmentDescribeDepartmentMemberPageList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepartmentDescribeDepartmentMemberPageList(departmentMembers []DepartmentDepartmentMembersData, pageNumber int64, pageSize int64, totalCount int64) *DepartmentDescribeDepartmentMemberPageList {
	this := DepartmentDescribeDepartmentMemberPageList{}
	this.DepartmentMembers = departmentMembers
	this.PageNumber = pageNumber
	this.PageSize = pageSize
	this.TotalCount = totalCount
	return &this
}

// NewDepartmentDescribeDepartmentMemberPageListWithDefaults instantiates a new DepartmentDescribeDepartmentMemberPageList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepartmentDescribeDepartmentMemberPageListWithDefaults() *DepartmentDescribeDepartmentMemberPageList {
	this := DepartmentDescribeDepartmentMemberPageList{}
	var pageNumber int64 = 0
	this.PageNumber = pageNumber
	var pageSize int64 = 0
	this.PageSize = pageSize
	var totalCount int64 = 0
	this.TotalCount = totalCount
	return &this
}

// GetDepartmentMembers returns the DepartmentMembers field value
func (o *DepartmentDescribeDepartmentMemberPageList) GetDepartmentMembers() []DepartmentDepartmentMembersData {
	if o == nil {
		var ret []DepartmentDepartmentMembersData
		return ret
	}

	return o.DepartmentMembers
}

// GetDepartmentMembersOk returns a tuple with the DepartmentMembers field value
// and a boolean to check if the value has been set.
func (o *DepartmentDescribeDepartmentMemberPageList) GetDepartmentMembersOk() ([]DepartmentDepartmentMembersData, bool) {
	if o == nil {
		return nil, false
	}
	return o.DepartmentMembers, true
}

// SetDepartmentMembers sets field value
func (o *DepartmentDescribeDepartmentMemberPageList) SetDepartmentMembers(v []DepartmentDepartmentMembersData) {
	o.DepartmentMembers = v
}

// GetPageNumber returns the PageNumber field value
func (o *DepartmentDescribeDepartmentMemberPageList) GetPageNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *DepartmentDescribeDepartmentMemberPageList) GetPageNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value
func (o *DepartmentDescribeDepartmentMemberPageList) SetPageNumber(v int64) {
	o.PageNumber = v
}

// GetPageSize returns the PageSize field value
func (o *DepartmentDescribeDepartmentMemberPageList) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *DepartmentDescribeDepartmentMemberPageList) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *DepartmentDescribeDepartmentMemberPageList) SetPageSize(v int64) {
	o.PageSize = v
}

// GetTotalCount returns the TotalCount field value
func (o *DepartmentDescribeDepartmentMemberPageList) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *DepartmentDescribeDepartmentMemberPageList) GetTotalCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *DepartmentDescribeDepartmentMemberPageList) SetTotalCount(v int64) {
	o.TotalCount = v
}

func (o DepartmentDescribeDepartmentMemberPageList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepartmentDescribeDepartmentMemberPageList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DepartmentMembers"] = o.DepartmentMembers
	toSerialize["PageNumber"] = o.PageNumber
	toSerialize["PageSize"] = o.PageSize
	toSerialize["TotalCount"] = o.TotalCount
	return toSerialize, nil
}

func (o *DepartmentDescribeDepartmentMemberPageList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepartmentMembers",
		"PageNumber",
		"PageSize",
		"TotalCount",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDepartmentDescribeDepartmentMemberPageList := _DepartmentDescribeDepartmentMemberPageList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDepartmentDescribeDepartmentMemberPageList)

	if err != nil {
		return err
	}

	*o = DepartmentDescribeDepartmentMemberPageList(varDepartmentDescribeDepartmentMemberPageList)

	return err
}

type NullableDepartmentDescribeDepartmentMemberPageList struct {
	value *DepartmentDescribeDepartmentMemberPageList
	isSet bool
}

func (v NullableDepartmentDescribeDepartmentMemberPageList) Get() *DepartmentDescribeDepartmentMemberPageList {
	return v.value
}

func (v *NullableDepartmentDescribeDepartmentMemberPageList) Set(val *DepartmentDescribeDepartmentMemberPageList) {
	v.value = val
	v.isSet = true
}

func (v NullableDepartmentDescribeDepartmentMemberPageList) IsSet() bool {
	return v.isSet
}

func (v *NullableDepartmentDescribeDepartmentMemberPageList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepartmentDescribeDepartmentMemberPageList(val *DepartmentDescribeDepartmentMemberPageList) *NullableDepartmentDescribeDepartmentMemberPageList {
	return &NullableDepartmentDescribeDepartmentMemberPageList{value: val, isSet: true}
}

func (v NullableDepartmentDescribeDepartmentMemberPageList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepartmentDescribeDepartmentMemberPageList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


