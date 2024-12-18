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

// checks if the DescribeMyDepotsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeMyDepotsRequest{}

// DescribeMyDepotsRequest struct for DescribeMyDepotsRequest
type DescribeMyDepotsRequest struct {
	// 页数
	PageNumber int64 `json:"PageNumber"`
	// 页码
	PageSize int64 `json:"PageSize"`
}

type _DescribeMyDepotsRequest DescribeMyDepotsRequest

// NewDescribeMyDepotsRequest instantiates a new DescribeMyDepotsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeMyDepotsRequest(pageNumber int64, pageSize int64) *DescribeMyDepotsRequest {
	this := DescribeMyDepotsRequest{}
	this.PageNumber = pageNumber
	this.PageSize = pageSize
	return &this
}

// NewDescribeMyDepotsRequestWithDefaults instantiates a new DescribeMyDepotsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeMyDepotsRequestWithDefaults() *DescribeMyDepotsRequest {
	this := DescribeMyDepotsRequest{}
	return &this
}

// GetPageNumber returns the PageNumber field value
func (o *DescribeMyDepotsRequest) GetPageNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *DescribeMyDepotsRequest) GetPageNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value
func (o *DescribeMyDepotsRequest) SetPageNumber(v int64) {
	o.PageNumber = v
}

// GetPageSize returns the PageSize field value
func (o *DescribeMyDepotsRequest) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *DescribeMyDepotsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *DescribeMyDepotsRequest) SetPageSize(v int64) {
	o.PageSize = v
}

func (o DescribeMyDepotsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeMyDepotsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["PageNumber"] = o.PageNumber
	toSerialize["PageSize"] = o.PageSize
	return toSerialize, nil
}

func (o *DescribeMyDepotsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varDescribeMyDepotsRequest := _DescribeMyDepotsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeMyDepotsRequest)

	if err != nil {
		return err
	}

	*o = DescribeMyDepotsRequest(varDescribeMyDepotsRequest)

	return err
}

type NullableDescribeMyDepotsRequest struct {
	value *DescribeMyDepotsRequest
	isSet bool
}

func (v NullableDescribeMyDepotsRequest) Get() *DescribeMyDepotsRequest {
	return v.value
}

func (v *NullableDescribeMyDepotsRequest) Set(val *DescribeMyDepotsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeMyDepotsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeMyDepotsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeMyDepotsRequest(val *DescribeMyDepotsRequest) *NullableDescribeMyDepotsRequest {
	return &NullableDescribeMyDepotsRequest{value: val, isSet: true}
}

func (v NullableDescribeMyDepotsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeMyDepotsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


