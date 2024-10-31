/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeGitBlobRaw200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitBlobRaw200ResponseResponse{}

// DescribeGitBlobRaw200ResponseResponse 公共返回体
type DescribeGitBlobRaw200ResponseResponse struct {
	// blob 文本内容
	Content utils.NullableString `json:"Content,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeGitBlobRaw200ResponseResponse instantiates a new DescribeGitBlobRaw200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitBlobRaw200ResponseResponse() *DescribeGitBlobRaw200ResponseResponse {
	this := DescribeGitBlobRaw200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeGitBlobRaw200ResponseResponseWithDefaults instantiates a new DescribeGitBlobRaw200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitBlobRaw200ResponseResponseWithDefaults() *DescribeGitBlobRaw200ResponseResponse {
	this := DescribeGitBlobRaw200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DescribeGitBlobRaw200ResponseResponse) GetContent() string {
	if o == nil || utils.IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DescribeGitBlobRaw200ResponseResponse) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *DescribeGitBlobRaw200ResponseResponse) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given utils.NullableString and assigns it to the Content field.
func (o *DescribeGitBlobRaw200ResponseResponse) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *DescribeGitBlobRaw200ResponseResponse) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *DescribeGitBlobRaw200ResponseResponse) UnsetContent() {
	o.Content.Unset()
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeGitBlobRaw200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitBlobRaw200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeGitBlobRaw200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeGitBlobRaw200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeGitBlobRaw200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitBlobRaw200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Content.IsSet() {
		toSerialize["Content"] = o.Content.Get()
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeGitBlobRaw200ResponseResponse struct {
	value *DescribeGitBlobRaw200ResponseResponse
	isSet bool
}

func (v NullableDescribeGitBlobRaw200ResponseResponse) Get() *DescribeGitBlobRaw200ResponseResponse {
	return v.value
}

func (v *NullableDescribeGitBlobRaw200ResponseResponse) Set(val *DescribeGitBlobRaw200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitBlobRaw200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitBlobRaw200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitBlobRaw200ResponseResponse(val *DescribeGitBlobRaw200ResponseResponse) *NullableDescribeGitBlobRaw200ResponseResponse {
	return &NullableDescribeGitBlobRaw200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeGitBlobRaw200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitBlobRaw200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

