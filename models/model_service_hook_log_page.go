/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ServiceHookLogPage type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ServiceHookLogPage{}

// ServiceHookLogPage Service Hook 发送记录分页
type ServiceHookLogPage struct {
	// Service Hook 发送记录列表
	Log []ServiceHookLog `json:"Log,omitempty"`
	// 分页页码
	PageNumber *int64 `json:"PageNumber,omitempty"`
	// 分页大小
	PageSize *int64 `json:"PageSize,omitempty"`
	// 总记录数
	TotalCount *int64 `json:"TotalCount,omitempty"`
}

// NewServiceHookLogPage instantiates a new ServiceHookLogPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceHookLogPage() *ServiceHookLogPage {
	this := ServiceHookLogPage{}
	return &this
}

// NewServiceHookLogPageWithDefaults instantiates a new ServiceHookLogPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceHookLogPageWithDefaults() *ServiceHookLogPage {
	this := ServiceHookLogPage{}
	return &this
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *ServiceHookLogPage) GetLog() []ServiceHookLog {
	if o == nil || utils.IsNil(o.Log) {
		var ret []ServiceHookLog
		return ret
	}
	return o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHookLogPage) GetLogOk() ([]ServiceHookLog, bool) {
	if o == nil || utils.IsNil(o.Log) {
		return nil, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *ServiceHookLogPage) HasLog() bool {
	if o != nil && !utils.IsNil(o.Log) {
		return true
	}

	return false
}

// SetLog gets a reference to the given []ServiceHookLog and assigns it to the Log field.
func (o *ServiceHookLogPage) SetLog(v []ServiceHookLog) {
	o.Log = v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *ServiceHookLogPage) GetPageNumber() int64 {
	if o == nil || utils.IsNil(o.PageNumber) {
		var ret int64
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHookLogPage) GetPageNumberOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *ServiceHookLogPage) HasPageNumber() bool {
	if o != nil && !utils.IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int64 and assigns it to the PageNumber field.
func (o *ServiceHookLogPage) SetPageNumber(v int64) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ServiceHookLogPage) GetPageSize() int64 {
	if o == nil || utils.IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHookLogPage) GetPageSizeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ServiceHookLogPage) HasPageSize() bool {
	if o != nil && !utils.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *ServiceHookLogPage) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ServiceHookLogPage) GetTotalCount() int64 {
	if o == nil || utils.IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHookLogPage) GetTotalCountOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ServiceHookLogPage) HasTotalCount() bool {
	if o != nil && !utils.IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *ServiceHookLogPage) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o ServiceHookLogPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceHookLogPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Log) {
		toSerialize["Log"] = o.Log
	}
	if !utils.IsNil(o.PageNumber) {
		toSerialize["PageNumber"] = o.PageNumber
	}
	if !utils.IsNil(o.PageSize) {
		toSerialize["PageSize"] = o.PageSize
	}
	if !utils.IsNil(o.TotalCount) {
		toSerialize["TotalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableServiceHookLogPage struct {
	value *ServiceHookLogPage
	isSet bool
}

func (v NullableServiceHookLogPage) Get() *ServiceHookLogPage {
	return v.value
}

func (v *NullableServiceHookLogPage) Set(val *ServiceHookLogPage) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceHookLogPage) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceHookLogPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceHookLogPage(val *ServiceHookLogPage) *NullableServiceHookLogPage {
	return &NullableServiceHookLogPage{value: val, isSet: true}
}

func (v NullableServiceHookLogPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceHookLogPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


