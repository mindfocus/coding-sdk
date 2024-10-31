/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ModifyDepartment200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyDepartment200ResponseResponse{}

// ModifyDepartment200ResponseResponse 公共返回体
type ModifyDepartment200ResponseResponse struct {
	Department *DepartmentDepartmentData `json:"Department,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewModifyDepartment200ResponseResponse instantiates a new ModifyDepartment200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyDepartment200ResponseResponse() *ModifyDepartment200ResponseResponse {
	this := ModifyDepartment200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewModifyDepartment200ResponseResponseWithDefaults instantiates a new ModifyDepartment200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyDepartment200ResponseResponseWithDefaults() *ModifyDepartment200ResponseResponse {
	this := ModifyDepartment200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *ModifyDepartment200ResponseResponse) GetDepartment() DepartmentDepartmentData {
	if o == nil || utils.IsNil(o.Department) {
		var ret DepartmentDepartmentData
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyDepartment200ResponseResponse) GetDepartmentOk() (*DepartmentDepartmentData, bool) {
	if o == nil || utils.IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *ModifyDepartment200ResponseResponse) HasDepartment() bool {
	if o != nil && !utils.IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given DepartmentDepartmentData and assigns it to the Department field.
func (o *ModifyDepartment200ResponseResponse) SetDepartment(v DepartmentDepartmentData) {
	o.Department = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ModifyDepartment200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyDepartment200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ModifyDepartment200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ModifyDepartment200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o ModifyDepartment200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyDepartment200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Department) {
		toSerialize["Department"] = o.Department
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableModifyDepartment200ResponseResponse struct {
	value *ModifyDepartment200ResponseResponse
	isSet bool
}

func (v NullableModifyDepartment200ResponseResponse) Get() *ModifyDepartment200ResponseResponse {
	return v.value
}

func (v *NullableModifyDepartment200ResponseResponse) Set(val *ModifyDepartment200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyDepartment200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyDepartment200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyDepartment200ResponseResponse(val *ModifyDepartment200ResponseResponse) *NullableModifyDepartment200ResponseResponse {
	return &NullableModifyDepartment200ResponseResponse{value: val, isSet: true}
}

func (v NullableModifyDepartment200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyDepartment200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

