/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CreateUploadToken200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateUploadToken200ResponseResponse{}

// CreateUploadToken200ResponseResponse 公共返回体
type CreateUploadToken200ResponseResponse struct {
	Token *Token `json:"Token,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewCreateUploadToken200ResponseResponse instantiates a new CreateUploadToken200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUploadToken200ResponseResponse() *CreateUploadToken200ResponseResponse {
	this := CreateUploadToken200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewCreateUploadToken200ResponseResponseWithDefaults instantiates a new CreateUploadToken200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUploadToken200ResponseResponseWithDefaults() *CreateUploadToken200ResponseResponse {
	this := CreateUploadToken200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateUploadToken200ResponseResponse) GetToken() Token {
	if o == nil || utils.IsNil(o.Token) {
		var ret Token
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUploadToken200ResponseResponse) GetTokenOk() (*Token, bool) {
	if o == nil || utils.IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateUploadToken200ResponseResponse) HasToken() bool {
	if o != nil && !utils.IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given Token and assigns it to the Token field.
func (o *CreateUploadToken200ResponseResponse) SetToken(v Token) {
	o.Token = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CreateUploadToken200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUploadToken200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CreateUploadToken200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CreateUploadToken200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o CreateUploadToken200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUploadToken200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Token) {
		toSerialize["Token"] = o.Token
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableCreateUploadToken200ResponseResponse struct {
	value *CreateUploadToken200ResponseResponse
	isSet bool
}

func (v NullableCreateUploadToken200ResponseResponse) Get() *CreateUploadToken200ResponseResponse {
	return v.value
}

func (v *NullableCreateUploadToken200ResponseResponse) Set(val *CreateUploadToken200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUploadToken200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUploadToken200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUploadToken200ResponseResponse(val *CreateUploadToken200ResponseResponse) *NullableCreateUploadToken200ResponseResponse {
	return &NullableCreateUploadToken200ResponseResponse{value: val, isSet: true}
}

func (v NullableCreateUploadToken200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUploadToken200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


