/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ModifyTestCaseSection200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyTestCaseSection200ResponseResponse{}

// ModifyTestCaseSection200ResponseResponse 公共返回体
type ModifyTestCaseSection200ResponseResponse struct {
	Data *SectionData `json:"Data,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewModifyTestCaseSection200ResponseResponse instantiates a new ModifyTestCaseSection200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyTestCaseSection200ResponseResponse() *ModifyTestCaseSection200ResponseResponse {
	this := ModifyTestCaseSection200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewModifyTestCaseSection200ResponseResponseWithDefaults instantiates a new ModifyTestCaseSection200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyTestCaseSection200ResponseResponseWithDefaults() *ModifyTestCaseSection200ResponseResponse {
	this := ModifyTestCaseSection200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ModifyTestCaseSection200ResponseResponse) GetData() SectionData {
	if o == nil || utils.IsNil(o.Data) {
		var ret SectionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestCaseSection200ResponseResponse) GetDataOk() (*SectionData, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ModifyTestCaseSection200ResponseResponse) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SectionData and assigns it to the Data field.
func (o *ModifyTestCaseSection200ResponseResponse) SetData(v SectionData) {
	o.Data = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ModifyTestCaseSection200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestCaseSection200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ModifyTestCaseSection200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ModifyTestCaseSection200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o ModifyTestCaseSection200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyTestCaseSection200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Data) {
		toSerialize["Data"] = o.Data
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableModifyTestCaseSection200ResponseResponse struct {
	value *ModifyTestCaseSection200ResponseResponse
	isSet bool
}

func (v NullableModifyTestCaseSection200ResponseResponse) Get() *ModifyTestCaseSection200ResponseResponse {
	return v.value
}

func (v *NullableModifyTestCaseSection200ResponseResponse) Set(val *ModifyTestCaseSection200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyTestCaseSection200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyTestCaseSection200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyTestCaseSection200ResponseResponse(val *ModifyTestCaseSection200ResponseResponse) *NullableModifyTestCaseSection200ResponseResponse {
	return &NullableModifyTestCaseSection200ResponseResponse{value: val, isSet: true}
}

func (v NullableModifyTestCaseSection200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyTestCaseSection200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


