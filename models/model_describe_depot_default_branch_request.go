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

// checks if the DescribeDepotDefaultBranchRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeDepotDefaultBranchRequest{}

// DescribeDepotDefaultBranchRequest struct for DescribeDepotDefaultBranchRequest
type DescribeDepotDefaultBranchRequest struct {
	// 仓库路径
	DepotPath string `json:"DepotPath"`
}

type _DescribeDepotDefaultBranchRequest DescribeDepotDefaultBranchRequest

// NewDescribeDepotDefaultBranchRequest instantiates a new DescribeDepotDefaultBranchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeDepotDefaultBranchRequest(depotPath string) *DescribeDepotDefaultBranchRequest {
	this := DescribeDepotDefaultBranchRequest{}
	this.DepotPath = depotPath
	return &this
}

// NewDescribeDepotDefaultBranchRequestWithDefaults instantiates a new DescribeDepotDefaultBranchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeDepotDefaultBranchRequestWithDefaults() *DescribeDepotDefaultBranchRequest {
	this := DescribeDepotDefaultBranchRequest{}
	return &this
}

// GetDepotPath returns the DepotPath field value
func (o *DescribeDepotDefaultBranchRequest) GetDepotPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value
// and a boolean to check if the value has been set.
func (o *DescribeDepotDefaultBranchRequest) GetDepotPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotPath, true
}

// SetDepotPath sets field value
func (o *DescribeDepotDefaultBranchRequest) SetDepotPath(v string) {
	o.DepotPath = v
}

func (o DescribeDepotDefaultBranchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeDepotDefaultBranchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DepotPath"] = o.DepotPath
	return toSerialize, nil
}

func (o *DescribeDepotDefaultBranchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepotPath",
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

	varDescribeDepotDefaultBranchRequest := _DescribeDepotDefaultBranchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeDepotDefaultBranchRequest)

	if err != nil {
		return err
	}

	*o = DescribeDepotDefaultBranchRequest(varDescribeDepotDefaultBranchRequest)

	return err
}

type NullableDescribeDepotDefaultBranchRequest struct {
	value *DescribeDepotDefaultBranchRequest
	isSet bool
}

func (v NullableDescribeDepotDefaultBranchRequest) Get() *DescribeDepotDefaultBranchRequest {
	return v.value
}

func (v *NullableDescribeDepotDefaultBranchRequest) Set(val *DescribeDepotDefaultBranchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeDepotDefaultBranchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeDepotDefaultBranchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeDepotDefaultBranchRequest(val *DescribeDepotDefaultBranchRequest) *NullableDescribeDepotDefaultBranchRequest {
	return &NullableDescribeDepotDefaultBranchRequest{value: val, isSet: true}
}

func (v NullableDescribeDepotDefaultBranchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeDepotDefaultBranchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


