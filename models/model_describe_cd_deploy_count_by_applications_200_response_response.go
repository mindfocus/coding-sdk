/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCdDeployCountByApplications200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCdDeployCountByApplications200ResponseResponse{}

// DescribeCdDeployCountByApplications200ResponseResponse 公共返回体
type DescribeCdDeployCountByApplications200ResponseResponse struct {
	Data *DescribeCdDeployCountByApplicationsResponseData `json:"Data,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeCdDeployCountByApplications200ResponseResponse instantiates a new DescribeCdDeployCountByApplications200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCdDeployCountByApplications200ResponseResponse() *DescribeCdDeployCountByApplications200ResponseResponse {
	this := DescribeCdDeployCountByApplications200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeCdDeployCountByApplications200ResponseResponseWithDefaults instantiates a new DescribeCdDeployCountByApplications200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCdDeployCountByApplications200ResponseResponseWithDefaults() *DescribeCdDeployCountByApplications200ResponseResponse {
	this := DescribeCdDeployCountByApplications200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DescribeCdDeployCountByApplications200ResponseResponse) GetData() DescribeCdDeployCountByApplicationsResponseData {
	if o == nil || utils.IsNil(o.Data) {
		var ret DescribeCdDeployCountByApplicationsResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdDeployCountByApplications200ResponseResponse) GetDataOk() (*DescribeCdDeployCountByApplicationsResponseData, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DescribeCdDeployCountByApplications200ResponseResponse) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DescribeCdDeployCountByApplicationsResponseData and assigns it to the Data field.
func (o *DescribeCdDeployCountByApplications200ResponseResponse) SetData(v DescribeCdDeployCountByApplicationsResponseData) {
	o.Data = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeCdDeployCountByApplications200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdDeployCountByApplications200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeCdDeployCountByApplications200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeCdDeployCountByApplications200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeCdDeployCountByApplications200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCdDeployCountByApplications200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Data) {
		toSerialize["Data"] = o.Data
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeCdDeployCountByApplications200ResponseResponse struct {
	value *DescribeCdDeployCountByApplications200ResponseResponse
	isSet bool
}

func (v NullableDescribeCdDeployCountByApplications200ResponseResponse) Get() *DescribeCdDeployCountByApplications200ResponseResponse {
	return v.value
}

func (v *NullableDescribeCdDeployCountByApplications200ResponseResponse) Set(val *DescribeCdDeployCountByApplications200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCdDeployCountByApplications200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCdDeployCountByApplications200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCdDeployCountByApplications200ResponseResponse(val *DescribeCdDeployCountByApplications200ResponseResponse) *NullableDescribeCdDeployCountByApplications200ResponseResponse {
	return &NullableDescribeCdDeployCountByApplications200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeCdDeployCountByApplications200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCdDeployCountByApplications200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

