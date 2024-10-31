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

// checks if the DescribePoliciesOnResourceTypeRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribePoliciesOnResourceTypeRequest{}

// DescribePoliciesOnResourceTypeRequest struct for DescribePoliciesOnResourceTypeRequest
type DescribePoliciesOnResourceTypeRequest struct {
	Filter DescribePoliciesOnResourceTypeRequestFilter `json:"Filter"`
	// 请求页数
	PageNumber int32 `json:"PageNumber"`
	// 请求条数
	PageSize int32 `json:"PageSize"`
}

type _DescribePoliciesOnResourceTypeRequest DescribePoliciesOnResourceTypeRequest

// NewDescribePoliciesOnResourceTypeRequest instantiates a new DescribePoliciesOnResourceTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribePoliciesOnResourceTypeRequest(filter DescribePoliciesOnResourceTypeRequestFilter, pageNumber int32, pageSize int32) *DescribePoliciesOnResourceTypeRequest {
	this := DescribePoliciesOnResourceTypeRequest{}
	this.Filter = filter
	this.PageNumber = pageNumber
	this.PageSize = pageSize
	return &this
}

// NewDescribePoliciesOnResourceTypeRequestWithDefaults instantiates a new DescribePoliciesOnResourceTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribePoliciesOnResourceTypeRequestWithDefaults() *DescribePoliciesOnResourceTypeRequest {
	this := DescribePoliciesOnResourceTypeRequest{}
	return &this
}

// GetFilter returns the Filter field value
func (o *DescribePoliciesOnResourceTypeRequest) GetFilter() DescribePoliciesOnResourceTypeRequestFilter {
	if o == nil {
		var ret DescribePoliciesOnResourceTypeRequestFilter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *DescribePoliciesOnResourceTypeRequest) GetFilterOk() (*DescribePoliciesOnResourceTypeRequestFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *DescribePoliciesOnResourceTypeRequest) SetFilter(v DescribePoliciesOnResourceTypeRequestFilter) {
	o.Filter = v
}

// GetPageNumber returns the PageNumber field value
func (o *DescribePoliciesOnResourceTypeRequest) GetPageNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *DescribePoliciesOnResourceTypeRequest) GetPageNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value
func (o *DescribePoliciesOnResourceTypeRequest) SetPageNumber(v int32) {
	o.PageNumber = v
}

// GetPageSize returns the PageSize field value
func (o *DescribePoliciesOnResourceTypeRequest) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *DescribePoliciesOnResourceTypeRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *DescribePoliciesOnResourceTypeRequest) SetPageSize(v int32) {
	o.PageSize = v
}

func (o DescribePoliciesOnResourceTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribePoliciesOnResourceTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Filter"] = o.Filter
	toSerialize["PageNumber"] = o.PageNumber
	toSerialize["PageSize"] = o.PageSize
	return toSerialize, nil
}

func (o *DescribePoliciesOnResourceTypeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Filter",
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

	varDescribePoliciesOnResourceTypeRequest := _DescribePoliciesOnResourceTypeRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribePoliciesOnResourceTypeRequest)

	if err != nil {
		return err
	}

	*o = DescribePoliciesOnResourceTypeRequest(varDescribePoliciesOnResourceTypeRequest)

	return err
}

type NullableDescribePoliciesOnResourceTypeRequest struct {
	value *DescribePoliciesOnResourceTypeRequest
	isSet bool
}

func (v NullableDescribePoliciesOnResourceTypeRequest) Get() *DescribePoliciesOnResourceTypeRequest {
	return v.value
}

func (v *NullableDescribePoliciesOnResourceTypeRequest) Set(val *DescribePoliciesOnResourceTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribePoliciesOnResourceTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribePoliciesOnResourceTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribePoliciesOnResourceTypeRequest(val *DescribePoliciesOnResourceTypeRequest) *NullableDescribePoliciesOnResourceTypeRequest {
	return &NullableDescribePoliciesOnResourceTypeRequest{value: val, isSet: true}
}

func (v NullableDescribePoliciesOnResourceTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribePoliciesOnResourceTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

