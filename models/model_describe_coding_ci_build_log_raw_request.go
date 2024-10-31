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

// checks if the DescribeCodingCIBuildLogRawRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCodingCIBuildLogRawRequest{}

// DescribeCodingCIBuildLogRawRequest struct for DescribeCodingCIBuildLogRawRequest
type DescribeCodingCIBuildLogRawRequest struct {
	// 构建 ID
	BuildId int32 `json:"BuildId"`
	// 日志开始位置
	Start int32 `json:"Start"`
}

type _DescribeCodingCIBuildLogRawRequest DescribeCodingCIBuildLogRawRequest

// NewDescribeCodingCIBuildLogRawRequest instantiates a new DescribeCodingCIBuildLogRawRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCodingCIBuildLogRawRequest(buildId int32, start int32) *DescribeCodingCIBuildLogRawRequest {
	this := DescribeCodingCIBuildLogRawRequest{}
	this.BuildId = buildId
	this.Start = start
	return &this
}

// NewDescribeCodingCIBuildLogRawRequestWithDefaults instantiates a new DescribeCodingCIBuildLogRawRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCodingCIBuildLogRawRequestWithDefaults() *DescribeCodingCIBuildLogRawRequest {
	this := DescribeCodingCIBuildLogRawRequest{}
	return &this
}

// GetBuildId returns the BuildId field value
func (o *DescribeCodingCIBuildLogRawRequest) GetBuildId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BuildId
}

// GetBuildIdOk returns a tuple with the BuildId field value
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildLogRawRequest) GetBuildIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildId, true
}

// SetBuildId sets field value
func (o *DescribeCodingCIBuildLogRawRequest) SetBuildId(v int32) {
	o.BuildId = v
}

// GetStart returns the Start field value
func (o *DescribeCodingCIBuildLogRawRequest) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildLogRawRequest) GetStartOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *DescribeCodingCIBuildLogRawRequest) SetStart(v int32) {
	o.Start = v
}

func (o DescribeCodingCIBuildLogRawRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCodingCIBuildLogRawRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["BuildId"] = o.BuildId
	toSerialize["Start"] = o.Start
	return toSerialize, nil
}

func (o *DescribeCodingCIBuildLogRawRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"BuildId",
		"Start",
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

	varDescribeCodingCIBuildLogRawRequest := _DescribeCodingCIBuildLogRawRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeCodingCIBuildLogRawRequest)

	if err != nil {
		return err
	}

	*o = DescribeCodingCIBuildLogRawRequest(varDescribeCodingCIBuildLogRawRequest)

	return err
}

type NullableDescribeCodingCIBuildLogRawRequest struct {
	value *DescribeCodingCIBuildLogRawRequest
	isSet bool
}

func (v NullableDescribeCodingCIBuildLogRawRequest) Get() *DescribeCodingCIBuildLogRawRequest {
	return v.value
}

func (v *NullableDescribeCodingCIBuildLogRawRequest) Set(val *DescribeCodingCIBuildLogRawRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCodingCIBuildLogRawRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCodingCIBuildLogRawRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCodingCIBuildLogRawRequest(val *DescribeCodingCIBuildLogRawRequest) *NullableDescribeCodingCIBuildLogRawRequest {
	return &NullableDescribeCodingCIBuildLogRawRequest{value: val, isSet: true}
}

func (v NullableDescribeCodingCIBuildLogRawRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCodingCIBuildLogRawRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


