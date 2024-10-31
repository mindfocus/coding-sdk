/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"bytes"
	"fmt"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCodingCIBuildsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCodingCIBuildsRequest{}

// DescribeCodingCIBuildsRequest struct for DescribeCodingCIBuildsRequest
type DescribeCodingCIBuildsRequest struct {
	// 构建计划 ID
	JobId int32 `json:"JobId"`
	// 页码
	PageNumber int32 `json:"PageNumber"`
	// 每页条数
	PageSize int32 `json:"PageSize"`
}

type _DescribeCodingCIBuildsRequest DescribeCodingCIBuildsRequest

// NewDescribeCodingCIBuildsRequest instantiates a new DescribeCodingCIBuildsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCodingCIBuildsRequest(jobId int32, pageNumber int32, pageSize int32) *DescribeCodingCIBuildsRequest {
	this := DescribeCodingCIBuildsRequest{}
	this.JobId = jobId
	this.PageNumber = pageNumber
	this.PageSize = pageSize
	return &this
}

// NewDescribeCodingCIBuildsRequestWithDefaults instantiates a new DescribeCodingCIBuildsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCodingCIBuildsRequestWithDefaults() *DescribeCodingCIBuildsRequest {
	this := DescribeCodingCIBuildsRequest{}
	return &this
}

// GetJobId returns the JobId field value
func (o *DescribeCodingCIBuildsRequest) GetJobId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildsRequest) GetJobIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *DescribeCodingCIBuildsRequest) SetJobId(v int32) {
	o.JobId = v
}

// GetPageNumber returns the PageNumber field value
func (o *DescribeCodingCIBuildsRequest) GetPageNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildsRequest) GetPageNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value
func (o *DescribeCodingCIBuildsRequest) SetPageNumber(v int32) {
	o.PageNumber = v
}

// GetPageSize returns the PageSize field value
func (o *DescribeCodingCIBuildsRequest) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildsRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *DescribeCodingCIBuildsRequest) SetPageSize(v int32) {
	o.PageSize = v
}

func (o DescribeCodingCIBuildsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCodingCIBuildsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["JobId"] = o.JobId
	toSerialize["PageNumber"] = o.PageNumber
	toSerialize["PageSize"] = o.PageSize
	return toSerialize, nil
}

func (o *DescribeCodingCIBuildsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"JobId",
		"PageNumber",
		"PageSize",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDescribeCodingCIBuildsRequest := _DescribeCodingCIBuildsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeCodingCIBuildsRequest)

	if err != nil {
		return err
	}

	*o = DescribeCodingCIBuildsRequest(varDescribeCodingCIBuildsRequest)

	return err
}

type NullableDescribeCodingCIBuildsRequest struct {
	value *DescribeCodingCIBuildsRequest
	isSet bool
}

func (v NullableDescribeCodingCIBuildsRequest) Get() *DescribeCodingCIBuildsRequest {
	return v.value
}

func (v *NullableDescribeCodingCIBuildsRequest) Set(val *DescribeCodingCIBuildsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCodingCIBuildsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCodingCIBuildsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCodingCIBuildsRequest(val *DescribeCodingCIBuildsRequest) *NullableDescribeCodingCIBuildsRequest {
	return &NullableDescribeCodingCIBuildsRequest{value: val, isSet: true}
}

func (v NullableDescribeCodingCIBuildsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCodingCIBuildsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


