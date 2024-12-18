/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the EnabledServiceHook200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EnabledServiceHook200ResponseResponse{}

// EnabledServiceHook200ResponseResponse 公共返回体
type EnabledServiceHook200ResponseResponse struct {
	// 是否操作成功
	Succeed *bool `json:"Succeed,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewEnabledServiceHook200ResponseResponse instantiates a new EnabledServiceHook200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnabledServiceHook200ResponseResponse() *EnabledServiceHook200ResponseResponse {
	this := EnabledServiceHook200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewEnabledServiceHook200ResponseResponseWithDefaults instantiates a new EnabledServiceHook200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnabledServiceHook200ResponseResponseWithDefaults() *EnabledServiceHook200ResponseResponse {
	this := EnabledServiceHook200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetSucceed returns the Succeed field value if set, zero value otherwise.
func (o *EnabledServiceHook200ResponseResponse) GetSucceed() bool {
	if o == nil || utils.IsNil(o.Succeed) {
		var ret bool
		return ret
	}
	return *o.Succeed
}

// GetSucceedOk returns a tuple with the Succeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnabledServiceHook200ResponseResponse) GetSucceedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Succeed) {
		return nil, false
	}
	return o.Succeed, true
}

// HasSucceed returns a boolean if a field has been set.
func (o *EnabledServiceHook200ResponseResponse) HasSucceed() bool {
	if o != nil && !utils.IsNil(o.Succeed) {
		return true
	}

	return false
}

// SetSucceed gets a reference to the given bool and assigns it to the Succeed field.
func (o *EnabledServiceHook200ResponseResponse) SetSucceed(v bool) {
	o.Succeed = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *EnabledServiceHook200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnabledServiceHook200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *EnabledServiceHook200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *EnabledServiceHook200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o EnabledServiceHook200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnabledServiceHook200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Succeed) {
		toSerialize["Succeed"] = o.Succeed
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableEnabledServiceHook200ResponseResponse struct {
	value *EnabledServiceHook200ResponseResponse
	isSet bool
}

func (v NullableEnabledServiceHook200ResponseResponse) Get() *EnabledServiceHook200ResponseResponse {
	return v.value
}

func (v *NullableEnabledServiceHook200ResponseResponse) Set(val *EnabledServiceHook200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnabledServiceHook200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnabledServiceHook200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnabledServiceHook200ResponseResponse(val *EnabledServiceHook200ResponseResponse) *NullableEnabledServiceHook200ResponseResponse {
	return &NullableEnabledServiceHook200ResponseResponse{value: val, isSet: true}
}

func (v NullableEnabledServiceHook200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnabledServiceHook200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


