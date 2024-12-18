/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeResourceReferencesCiting200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeResourceReferencesCiting200ResponseResponse{}

// DescribeResourceReferencesCiting200ResponseResponse 公共返回体
type DescribeResourceReferencesCiting200ResponseResponse struct {
	// 资源引用列表
	Resource []ResourceReferenceResourceInfo `json:"Resource,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeResourceReferencesCiting200ResponseResponse instantiates a new DescribeResourceReferencesCiting200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeResourceReferencesCiting200ResponseResponse() *DescribeResourceReferencesCiting200ResponseResponse {
	this := DescribeResourceReferencesCiting200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeResourceReferencesCiting200ResponseResponseWithDefaults instantiates a new DescribeResourceReferencesCiting200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeResourceReferencesCiting200ResponseResponseWithDefaults() *DescribeResourceReferencesCiting200ResponseResponse {
	this := DescribeResourceReferencesCiting200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *DescribeResourceReferencesCiting200ResponseResponse) GetResource() []ResourceReferenceResourceInfo {
	if o == nil || utils.IsNil(o.Resource) {
		var ret []ResourceReferenceResourceInfo
		return ret
	}
	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceReferencesCiting200ResponseResponse) GetResourceOk() ([]ResourceReferenceResourceInfo, bool) {
	if o == nil || utils.IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *DescribeResourceReferencesCiting200ResponseResponse) HasResource() bool {
	if o != nil && !utils.IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given []ResourceReferenceResourceInfo and assigns it to the Resource field.
func (o *DescribeResourceReferencesCiting200ResponseResponse) SetResource(v []ResourceReferenceResourceInfo) {
	o.Resource = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeResourceReferencesCiting200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceReferencesCiting200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeResourceReferencesCiting200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeResourceReferencesCiting200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeResourceReferencesCiting200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeResourceReferencesCiting200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Resource) {
		toSerialize["Resource"] = o.Resource
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeResourceReferencesCiting200ResponseResponse struct {
	value *DescribeResourceReferencesCiting200ResponseResponse
	isSet bool
}

func (v NullableDescribeResourceReferencesCiting200ResponseResponse) Get() *DescribeResourceReferencesCiting200ResponseResponse {
	return v.value
}

func (v *NullableDescribeResourceReferencesCiting200ResponseResponse) Set(val *DescribeResourceReferencesCiting200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeResourceReferencesCiting200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeResourceReferencesCiting200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeResourceReferencesCiting200ResponseResponse(val *DescribeResourceReferencesCiting200ResponseResponse) *NullableDescribeResourceReferencesCiting200ResponseResponse {
	return &NullableDescribeResourceReferencesCiting200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeResourceReferencesCiting200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeResourceReferencesCiting200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


