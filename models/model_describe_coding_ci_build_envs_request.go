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

// checks if the DescribeCodingCIBuildEnvsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCodingCIBuildEnvsRequest{}

// DescribeCodingCIBuildEnvsRequest struct for DescribeCodingCIBuildEnvsRequest
type DescribeCodingCIBuildEnvsRequest struct {
	// 构建ID
	BuildId int32 `json:"BuildId"`
	// SYSTEM（系统内置环境变量） Param（触发构建输入环境变量） Env（构建计划填写环境变量）
	Type string `json:"Type"`
}

type _DescribeCodingCIBuildEnvsRequest DescribeCodingCIBuildEnvsRequest

// NewDescribeCodingCIBuildEnvsRequest instantiates a new DescribeCodingCIBuildEnvsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCodingCIBuildEnvsRequest(buildId int32, type_ string) *DescribeCodingCIBuildEnvsRequest {
	this := DescribeCodingCIBuildEnvsRequest{}
	this.BuildId = buildId
	this.Type = type_
	return &this
}

// NewDescribeCodingCIBuildEnvsRequestWithDefaults instantiates a new DescribeCodingCIBuildEnvsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCodingCIBuildEnvsRequestWithDefaults() *DescribeCodingCIBuildEnvsRequest {
	this := DescribeCodingCIBuildEnvsRequest{}
	return &this
}

// GetBuildId returns the BuildId field value
func (o *DescribeCodingCIBuildEnvsRequest) GetBuildId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BuildId
}

// GetBuildIdOk returns a tuple with the BuildId field value
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildEnvsRequest) GetBuildIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildId, true
}

// SetBuildId sets field value
func (o *DescribeCodingCIBuildEnvsRequest) SetBuildId(v int32) {
	o.BuildId = v
}

// GetType returns the Type field value
func (o *DescribeCodingCIBuildEnvsRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildEnvsRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DescribeCodingCIBuildEnvsRequest) SetType(v string) {
	o.Type = v
}

func (o DescribeCodingCIBuildEnvsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCodingCIBuildEnvsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["BuildId"] = o.BuildId
	toSerialize["Type"] = o.Type
	return toSerialize, nil
}

func (o *DescribeCodingCIBuildEnvsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"BuildId",
		"Type",
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

	varDescribeCodingCIBuildEnvsRequest := _DescribeCodingCIBuildEnvsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeCodingCIBuildEnvsRequest)

	if err != nil {
		return err
	}

	*o = DescribeCodingCIBuildEnvsRequest(varDescribeCodingCIBuildEnvsRequest)

	return err
}

type NullableDescribeCodingCIBuildEnvsRequest struct {
	value *DescribeCodingCIBuildEnvsRequest
	isSet bool
}

func (v NullableDescribeCodingCIBuildEnvsRequest) Get() *DescribeCodingCIBuildEnvsRequest {
	return v.value
}

func (v *NullableDescribeCodingCIBuildEnvsRequest) Set(val *DescribeCodingCIBuildEnvsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCodingCIBuildEnvsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCodingCIBuildEnvsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCodingCIBuildEnvsRequest(val *DescribeCodingCIBuildEnvsRequest) *NullableDescribeCodingCIBuildEnvsRequest {
	return &NullableDescribeCodingCIBuildEnvsRequest{value: val, isSet: true}
}

func (v NullableDescribeCodingCIBuildEnvsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCodingCIBuildEnvsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


