/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CreateGitMergeReq200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateGitMergeReq200ResponseResponse{}

// CreateGitMergeReq200ResponseResponse 公共返回体
type CreateGitMergeReq200ResponseResponse struct {
	MergeInfo *MergeInfo `json:"MergeInfo,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewCreateGitMergeReq200ResponseResponse instantiates a new CreateGitMergeReq200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGitMergeReq200ResponseResponse() *CreateGitMergeReq200ResponseResponse {
	this := CreateGitMergeReq200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewCreateGitMergeReq200ResponseResponseWithDefaults instantiates a new CreateGitMergeReq200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGitMergeReq200ResponseResponseWithDefaults() *CreateGitMergeReq200ResponseResponse {
	this := CreateGitMergeReq200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetMergeInfo returns the MergeInfo field value if set, zero value otherwise.
func (o *CreateGitMergeReq200ResponseResponse) GetMergeInfo() MergeInfo {
	if o == nil || utils.IsNil(o.MergeInfo) {
		var ret MergeInfo
		return ret
	}
	return *o.MergeInfo
}

// GetMergeInfoOk returns a tuple with the MergeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGitMergeReq200ResponseResponse) GetMergeInfoOk() (*MergeInfo, bool) {
	if o == nil || utils.IsNil(o.MergeInfo) {
		return nil, false
	}
	return o.MergeInfo, true
}

// HasMergeInfo returns a boolean if a field has been set.
func (o *CreateGitMergeReq200ResponseResponse) HasMergeInfo() bool {
	if o != nil && !utils.IsNil(o.MergeInfo) {
		return true
	}

	return false
}

// SetMergeInfo gets a reference to the given MergeInfo and assigns it to the MergeInfo field.
func (o *CreateGitMergeReq200ResponseResponse) SetMergeInfo(v MergeInfo) {
	o.MergeInfo = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CreateGitMergeReq200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGitMergeReq200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CreateGitMergeReq200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CreateGitMergeReq200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o CreateGitMergeReq200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGitMergeReq200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.MergeInfo) {
		toSerialize["MergeInfo"] = o.MergeInfo
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableCreateGitMergeReq200ResponseResponse struct {
	value *CreateGitMergeReq200ResponseResponse
	isSet bool
}

func (v NullableCreateGitMergeReq200ResponseResponse) Get() *CreateGitMergeReq200ResponseResponse {
	return v.value
}

func (v *NullableCreateGitMergeReq200ResponseResponse) Set(val *CreateGitMergeReq200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGitMergeReq200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGitMergeReq200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGitMergeReq200ResponseResponse(val *CreateGitMergeReq200ResponseResponse) *NullableCreateGitMergeReq200ResponseResponse {
	return &NullableCreateGitMergeReq200ResponseResponse{value: val, isSet: true}
}

func (v NullableCreateGitMergeReq200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGitMergeReq200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

