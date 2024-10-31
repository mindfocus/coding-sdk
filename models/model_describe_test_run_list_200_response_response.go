/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeTestRunList200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeTestRunList200ResponseResponse{}

// DescribeTestRunList200ResponseResponse 公共返回体
type DescribeTestRunList200ResponseResponse struct {
	Data *RunsData `json:"Data,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeTestRunList200ResponseResponse instantiates a new DescribeTestRunList200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeTestRunList200ResponseResponse() *DescribeTestRunList200ResponseResponse {
	this := DescribeTestRunList200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeTestRunList200ResponseResponseWithDefaults instantiates a new DescribeTestRunList200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeTestRunList200ResponseResponseWithDefaults() *DescribeTestRunList200ResponseResponse {
	this := DescribeTestRunList200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DescribeTestRunList200ResponseResponse) GetData() RunsData {
	if o == nil || utils.IsNil(o.Data) {
		var ret RunsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTestRunList200ResponseResponse) GetDataOk() (*RunsData, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DescribeTestRunList200ResponseResponse) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given RunsData and assigns it to the Data field.
func (o *DescribeTestRunList200ResponseResponse) SetData(v RunsData) {
	o.Data = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeTestRunList200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTestRunList200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeTestRunList200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeTestRunList200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeTestRunList200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeTestRunList200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Data) {
		toSerialize["Data"] = o.Data
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeTestRunList200ResponseResponse struct {
	value *DescribeTestRunList200ResponseResponse
	isSet bool
}

func (v NullableDescribeTestRunList200ResponseResponse) Get() *DescribeTestRunList200ResponseResponse {
	return v.value
}

func (v *NullableDescribeTestRunList200ResponseResponse) Set(val *DescribeTestRunList200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeTestRunList200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeTestRunList200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeTestRunList200ResponseResponse(val *DescribeTestRunList200ResponseResponse) *NullableDescribeTestRunList200ResponseResponse {
	return &NullableDescribeTestRunList200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeTestRunList200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeTestRunList200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

