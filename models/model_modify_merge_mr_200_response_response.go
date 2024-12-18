/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ModifyMergeMR200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyMergeMR200ResponseResponse{}

// ModifyMergeMR200ResponseResponse 公共返回体
type ModifyMergeMR200ResponseResponse struct {
	MergeRequestInfo *MergeRequestInfo `json:"MergeRequestInfo,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewModifyMergeMR200ResponseResponse instantiates a new ModifyMergeMR200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyMergeMR200ResponseResponse() *ModifyMergeMR200ResponseResponse {
	this := ModifyMergeMR200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewModifyMergeMR200ResponseResponseWithDefaults instantiates a new ModifyMergeMR200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyMergeMR200ResponseResponseWithDefaults() *ModifyMergeMR200ResponseResponse {
	this := ModifyMergeMR200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetMergeRequestInfo returns the MergeRequestInfo field value if set, zero value otherwise.
func (o *ModifyMergeMR200ResponseResponse) GetMergeRequestInfo() MergeRequestInfo {
	if o == nil || utils.IsNil(o.MergeRequestInfo) {
		var ret MergeRequestInfo
		return ret
	}
	return *o.MergeRequestInfo
}

// GetMergeRequestInfoOk returns a tuple with the MergeRequestInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyMergeMR200ResponseResponse) GetMergeRequestInfoOk() (*MergeRequestInfo, bool) {
	if o == nil || utils.IsNil(o.MergeRequestInfo) {
		return nil, false
	}
	return o.MergeRequestInfo, true
}

// HasMergeRequestInfo returns a boolean if a field has been set.
func (o *ModifyMergeMR200ResponseResponse) HasMergeRequestInfo() bool {
	if o != nil && !utils.IsNil(o.MergeRequestInfo) {
		return true
	}

	return false
}

// SetMergeRequestInfo gets a reference to the given MergeRequestInfo and assigns it to the MergeRequestInfo field.
func (o *ModifyMergeMR200ResponseResponse) SetMergeRequestInfo(v MergeRequestInfo) {
	o.MergeRequestInfo = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ModifyMergeMR200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyMergeMR200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ModifyMergeMR200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ModifyMergeMR200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o ModifyMergeMR200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyMergeMR200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.MergeRequestInfo) {
		toSerialize["MergeRequestInfo"] = o.MergeRequestInfo
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableModifyMergeMR200ResponseResponse struct {
	value *ModifyMergeMR200ResponseResponse
	isSet bool
}

func (v NullableModifyMergeMR200ResponseResponse) Get() *ModifyMergeMR200ResponseResponse {
	return v.value
}

func (v *NullableModifyMergeMR200ResponseResponse) Set(val *ModifyMergeMR200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyMergeMR200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyMergeMR200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyMergeMR200ResponseResponse(val *ModifyMergeMR200ResponseResponse) *NullableModifyMergeMR200ResponseResponse {
	return &NullableModifyMergeMR200ResponseResponse{value: val, isSet: true}
}

func (v NullableModifyMergeMR200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyMergeMR200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


