/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCdHostServerGroupsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCdHostServerGroupsRequest{}

// DescribeCdHostServerGroupsRequest struct for DescribeCdHostServerGroupsRequest
type DescribeCdHostServerGroupsRequest struct {
	// 检索关键字
	Keyword *string `json:"Keyword,omitempty"`
	// 页码（默认为 1）
	PageNumber *int64 `json:"PageNumber,omitempty"`
	// 每页条数（默认为 10，最大值为 500）
	PageSize *int64 `json:"PageSize,omitempty"`
}

// NewDescribeCdHostServerGroupsRequest instantiates a new DescribeCdHostServerGroupsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCdHostServerGroupsRequest() *DescribeCdHostServerGroupsRequest {
	this := DescribeCdHostServerGroupsRequest{}
	return &this
}

// NewDescribeCdHostServerGroupsRequestWithDefaults instantiates a new DescribeCdHostServerGroupsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCdHostServerGroupsRequestWithDefaults() *DescribeCdHostServerGroupsRequest {
	this := DescribeCdHostServerGroupsRequest{}
	return &this
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *DescribeCdHostServerGroupsRequest) GetKeyword() string {
	if o == nil || utils.IsNil(o.Keyword) {
		var ret string
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdHostServerGroupsRequest) GetKeywordOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *DescribeCdHostServerGroupsRequest) HasKeyword() bool {
	if o != nil && !utils.IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *DescribeCdHostServerGroupsRequest) SetKeyword(v string) {
	o.Keyword = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *DescribeCdHostServerGroupsRequest) GetPageNumber() int64 {
	if o == nil || utils.IsNil(o.PageNumber) {
		var ret int64
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdHostServerGroupsRequest) GetPageNumberOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *DescribeCdHostServerGroupsRequest) HasPageNumber() bool {
	if o != nil && !utils.IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int64 and assigns it to the PageNumber field.
func (o *DescribeCdHostServerGroupsRequest) SetPageNumber(v int64) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *DescribeCdHostServerGroupsRequest) GetPageSize() int64 {
	if o == nil || utils.IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdHostServerGroupsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *DescribeCdHostServerGroupsRequest) HasPageSize() bool {
	if o != nil && !utils.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *DescribeCdHostServerGroupsRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

func (o DescribeCdHostServerGroupsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCdHostServerGroupsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Keyword) {
		toSerialize["Keyword"] = o.Keyword
	}
	if !utils.IsNil(o.PageNumber) {
		toSerialize["PageNumber"] = o.PageNumber
	}
	if !utils.IsNil(o.PageSize) {
		toSerialize["PageSize"] = o.PageSize
	}
	return toSerialize, nil
}

type NullableDescribeCdHostServerGroupsRequest struct {
	value *DescribeCdHostServerGroupsRequest
	isSet bool
}

func (v NullableDescribeCdHostServerGroupsRequest) Get() *DescribeCdHostServerGroupsRequest {
	return v.value
}

func (v *NullableDescribeCdHostServerGroupsRequest) Set(val *DescribeCdHostServerGroupsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCdHostServerGroupsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCdHostServerGroupsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCdHostServerGroupsRequest(val *DescribeCdHostServerGroupsRequest) *NullableDescribeCdHostServerGroupsRequest {
	return &NullableDescribeCdHostServerGroupsRequest{value: val, isSet: true}
}

func (v NullableDescribeCdHostServerGroupsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCdHostServerGroupsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

