/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeTest200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeTest200ResponseResponse{}

// DescribeTest200ResponseResponse 公共返回体
type DescribeTest200ResponseResponse struct {
	Data *TestData `json:"Data,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeTest200ResponseResponse instantiates a new DescribeTest200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeTest200ResponseResponse() *DescribeTest200ResponseResponse {
	this := DescribeTest200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeTest200ResponseResponseWithDefaults instantiates a new DescribeTest200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeTest200ResponseResponseWithDefaults() *DescribeTest200ResponseResponse {
	this := DescribeTest200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DescribeTest200ResponseResponse) GetData() TestData {
	if o == nil || utils.IsNil(o.Data) {
		var ret TestData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTest200ResponseResponse) GetDataOk() (*TestData, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DescribeTest200ResponseResponse) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given TestData and assigns it to the Data field.
func (o *DescribeTest200ResponseResponse) SetData(v TestData) {
	o.Data = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeTest200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTest200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeTest200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeTest200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeTest200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeTest200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Data) {
		toSerialize["Data"] = o.Data
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeTest200ResponseResponse struct {
	value *DescribeTest200ResponseResponse
	isSet bool
}

func (v NullableDescribeTest200ResponseResponse) Get() *DescribeTest200ResponseResponse {
	return v.value
}

func (v *NullableDescribeTest200ResponseResponse) Set(val *DescribeTest200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeTest200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeTest200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeTest200ResponseResponse(val *DescribeTest200ResponseResponse) *NullableDescribeTest200ResponseResponse {
	return &NullableDescribeTest200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeTest200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeTest200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


