/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeIssueFileUrl200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeIssueFileUrl200ResponseResponse{}

// DescribeIssueFileUrl200ResponseResponse 公共返回体
type DescribeIssueFileUrl200ResponseResponse struct {
	// 文件下载地址
	Url *string `json:"Url,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeIssueFileUrl200ResponseResponse instantiates a new DescribeIssueFileUrl200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeIssueFileUrl200ResponseResponse() *DescribeIssueFileUrl200ResponseResponse {
	this := DescribeIssueFileUrl200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeIssueFileUrl200ResponseResponseWithDefaults instantiates a new DescribeIssueFileUrl200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeIssueFileUrl200ResponseResponseWithDefaults() *DescribeIssueFileUrl200ResponseResponse {
	this := DescribeIssueFileUrl200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *DescribeIssueFileUrl200ResponseResponse) GetUrl() string {
	if o == nil || utils.IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeIssueFileUrl200ResponseResponse) GetUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *DescribeIssueFileUrl200ResponseResponse) HasUrl() bool {
	if o != nil && !utils.IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *DescribeIssueFileUrl200ResponseResponse) SetUrl(v string) {
	o.Url = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeIssueFileUrl200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeIssueFileUrl200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeIssueFileUrl200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeIssueFileUrl200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeIssueFileUrl200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeIssueFileUrl200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Url) {
		toSerialize["Url"] = o.Url
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeIssueFileUrl200ResponseResponse struct {
	value *DescribeIssueFileUrl200ResponseResponse
	isSet bool
}

func (v NullableDescribeIssueFileUrl200ResponseResponse) Get() *DescribeIssueFileUrl200ResponseResponse {
	return v.value
}

func (v *NullableDescribeIssueFileUrl200ResponseResponse) Set(val *DescribeIssueFileUrl200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeIssueFileUrl200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeIssueFileUrl200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeIssueFileUrl200ResponseResponse(val *DescribeIssueFileUrl200ResponseResponse) *NullableDescribeIssueFileUrl200ResponseResponse {
	return &NullableDescribeIssueFileUrl200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeIssueFileUrl200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeIssueFileUrl200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

