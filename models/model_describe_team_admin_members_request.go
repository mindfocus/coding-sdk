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

// checks if the DescribeTeamAdminMembersRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeTeamAdminMembersRequest{}

// DescribeTeamAdminMembersRequest struct for DescribeTeamAdminMembersRequest
type DescribeTeamAdminMembersRequest struct {
	// 第几页
	PageNumber int32 `json:"PageNumber"`
	// 每页条数
	PageSize int32 `json:"PageSize"`
}

type _DescribeTeamAdminMembersRequest DescribeTeamAdminMembersRequest

// NewDescribeTeamAdminMembersRequest instantiates a new DescribeTeamAdminMembersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeTeamAdminMembersRequest(pageNumber int32, pageSize int32) *DescribeTeamAdminMembersRequest {
	this := DescribeTeamAdminMembersRequest{}
	this.PageNumber = pageNumber
	this.PageSize = pageSize
	return &this
}

// NewDescribeTeamAdminMembersRequestWithDefaults instantiates a new DescribeTeamAdminMembersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeTeamAdminMembersRequestWithDefaults() *DescribeTeamAdminMembersRequest {
	this := DescribeTeamAdminMembersRequest{}
	return &this
}

// GetPageNumber returns the PageNumber field value
func (o *DescribeTeamAdminMembersRequest) GetPageNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *DescribeTeamAdminMembersRequest) GetPageNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value
func (o *DescribeTeamAdminMembersRequest) SetPageNumber(v int32) {
	o.PageNumber = v
}

// GetPageSize returns the PageSize field value
func (o *DescribeTeamAdminMembersRequest) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *DescribeTeamAdminMembersRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *DescribeTeamAdminMembersRequest) SetPageSize(v int32) {
	o.PageSize = v
}

func (o DescribeTeamAdminMembersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeTeamAdminMembersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["PageNumber"] = o.PageNumber
	toSerialize["PageSize"] = o.PageSize
	return toSerialize, nil
}

func (o *DescribeTeamAdminMembersRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDescribeTeamAdminMembersRequest := _DescribeTeamAdminMembersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeTeamAdminMembersRequest)

	if err != nil {
		return err
	}

	*o = DescribeTeamAdminMembersRequest(varDescribeTeamAdminMembersRequest)

	return err
}

type NullableDescribeTeamAdminMembersRequest struct {
	value *DescribeTeamAdminMembersRequest
	isSet bool
}

func (v NullableDescribeTeamAdminMembersRequest) Get() *DescribeTeamAdminMembersRequest {
	return v.value
}

func (v *NullableDescribeTeamAdminMembersRequest) Set(val *DescribeTeamAdminMembersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeTeamAdminMembersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeTeamAdminMembersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeTeamAdminMembersRequest(val *DescribeTeamAdminMembersRequest) *NullableDescribeTeamAdminMembersRequest {
	return &NullableDescribeTeamAdminMembersRequest{value: val, isSet: true}
}

func (v NullableDescribeTeamAdminMembersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeTeamAdminMembersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


