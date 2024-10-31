/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeProjectDepotInfoList200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeProjectDepotInfoList200ResponseResponse{}

// DescribeProjectDepotInfoList200ResponseResponse 公共返回体
type DescribeProjectDepotInfoList200ResponseResponse struct {
	DepotData *DepotData `json:"DepotData,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeProjectDepotInfoList200ResponseResponse instantiates a new DescribeProjectDepotInfoList200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeProjectDepotInfoList200ResponseResponse() *DescribeProjectDepotInfoList200ResponseResponse {
	this := DescribeProjectDepotInfoList200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeProjectDepotInfoList200ResponseResponseWithDefaults instantiates a new DescribeProjectDepotInfoList200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeProjectDepotInfoList200ResponseResponseWithDefaults() *DescribeProjectDepotInfoList200ResponseResponse {
	this := DescribeProjectDepotInfoList200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetDepotData returns the DepotData field value if set, zero value otherwise.
func (o *DescribeProjectDepotInfoList200ResponseResponse) GetDepotData() DepotData {
	if o == nil || utils.IsNil(o.DepotData) {
		var ret DepotData
		return ret
	}
	return *o.DepotData
}

// GetDepotDataOk returns a tuple with the DepotData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectDepotInfoList200ResponseResponse) GetDepotDataOk() (*DepotData, bool) {
	if o == nil || utils.IsNil(o.DepotData) {
		return nil, false
	}
	return o.DepotData, true
}

// HasDepotData returns a boolean if a field has been set.
func (o *DescribeProjectDepotInfoList200ResponseResponse) HasDepotData() bool {
	if o != nil && !utils.IsNil(o.DepotData) {
		return true
	}

	return false
}

// SetDepotData gets a reference to the given DepotData and assigns it to the DepotData field.
func (o *DescribeProjectDepotInfoList200ResponseResponse) SetDepotData(v DepotData) {
	o.DepotData = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeProjectDepotInfoList200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectDepotInfoList200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeProjectDepotInfoList200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeProjectDepotInfoList200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeProjectDepotInfoList200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeProjectDepotInfoList200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.DepotData) {
		toSerialize["DepotData"] = o.DepotData
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeProjectDepotInfoList200ResponseResponse struct {
	value *DescribeProjectDepotInfoList200ResponseResponse
	isSet bool
}

func (v NullableDescribeProjectDepotInfoList200ResponseResponse) Get() *DescribeProjectDepotInfoList200ResponseResponse {
	return v.value
}

func (v *NullableDescribeProjectDepotInfoList200ResponseResponse) Set(val *DescribeProjectDepotInfoList200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeProjectDepotInfoList200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeProjectDepotInfoList200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeProjectDepotInfoList200ResponseResponse(val *DescribeProjectDepotInfoList200ResponseResponse) *NullableDescribeProjectDepotInfoList200ResponseResponse {
	return &NullableDescribeProjectDepotInfoList200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeProjectDepotInfoList200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeProjectDepotInfoList200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


