/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCodingCIBuildStatistics200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCodingCIBuildStatistics200ResponseResponse{}

// DescribeCodingCIBuildStatistics200ResponseResponse 公共返回体
type DescribeCodingCIBuildStatistics200ResponseResponse struct {
	Data *DescribeCodingCIBuildStatisticsResponseData `json:"Data,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeCodingCIBuildStatistics200ResponseResponse instantiates a new DescribeCodingCIBuildStatistics200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCodingCIBuildStatistics200ResponseResponse() *DescribeCodingCIBuildStatistics200ResponseResponse {
	this := DescribeCodingCIBuildStatistics200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeCodingCIBuildStatistics200ResponseResponseWithDefaults instantiates a new DescribeCodingCIBuildStatistics200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCodingCIBuildStatistics200ResponseResponseWithDefaults() *DescribeCodingCIBuildStatistics200ResponseResponse {
	this := DescribeCodingCIBuildStatistics200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DescribeCodingCIBuildStatistics200ResponseResponse) GetData() DescribeCodingCIBuildStatisticsResponseData {
	if o == nil || utils.IsNil(o.Data) {
		var ret DescribeCodingCIBuildStatisticsResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildStatistics200ResponseResponse) GetDataOk() (*DescribeCodingCIBuildStatisticsResponseData, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DescribeCodingCIBuildStatistics200ResponseResponse) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DescribeCodingCIBuildStatisticsResponseData and assigns it to the Data field.
func (o *DescribeCodingCIBuildStatistics200ResponseResponse) SetData(v DescribeCodingCIBuildStatisticsResponseData) {
	o.Data = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeCodingCIBuildStatistics200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildStatistics200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeCodingCIBuildStatistics200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeCodingCIBuildStatistics200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeCodingCIBuildStatistics200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCodingCIBuildStatistics200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Data) {
		toSerialize["Data"] = o.Data
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeCodingCIBuildStatistics200ResponseResponse struct {
	value *DescribeCodingCIBuildStatistics200ResponseResponse
	isSet bool
}

func (v NullableDescribeCodingCIBuildStatistics200ResponseResponse) Get() *DescribeCodingCIBuildStatistics200ResponseResponse {
	return v.value
}

func (v *NullableDescribeCodingCIBuildStatistics200ResponseResponse) Set(val *DescribeCodingCIBuildStatistics200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCodingCIBuildStatistics200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCodingCIBuildStatistics200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCodingCIBuildStatistics200ResponseResponse(val *DescribeCodingCIBuildStatistics200ResponseResponse) *NullableDescribeCodingCIBuildStatistics200ResponseResponse {
	return &NullableDescribeCodingCIBuildStatistics200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeCodingCIBuildStatistics200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCodingCIBuildStatistics200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

