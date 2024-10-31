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

// checks if the DescribeCodingCIBuildStageRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCodingCIBuildStageRequest{}

// DescribeCodingCIBuildStageRequest struct for DescribeCodingCIBuildStageRequest
type DescribeCodingCIBuildStageRequest struct {
	// 构建 ID
	BuildId int32 `json:"BuildId"`
}

type _DescribeCodingCIBuildStageRequest DescribeCodingCIBuildStageRequest

// NewDescribeCodingCIBuildStageRequest instantiates a new DescribeCodingCIBuildStageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCodingCIBuildStageRequest(buildId int32) *DescribeCodingCIBuildStageRequest {
	this := DescribeCodingCIBuildStageRequest{}
	this.BuildId = buildId
	return &this
}

// NewDescribeCodingCIBuildStageRequestWithDefaults instantiates a new DescribeCodingCIBuildStageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCodingCIBuildStageRequestWithDefaults() *DescribeCodingCIBuildStageRequest {
	this := DescribeCodingCIBuildStageRequest{}
	return &this
}

// GetBuildId returns the BuildId field value
func (o *DescribeCodingCIBuildStageRequest) GetBuildId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BuildId
}

// GetBuildIdOk returns a tuple with the BuildId field value
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildStageRequest) GetBuildIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildId, true
}

// SetBuildId sets field value
func (o *DescribeCodingCIBuildStageRequest) SetBuildId(v int32) {
	o.BuildId = v
}

func (o DescribeCodingCIBuildStageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCodingCIBuildStageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["BuildId"] = o.BuildId
	return toSerialize, nil
}

func (o *DescribeCodingCIBuildStageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"BuildId",
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

	varDescribeCodingCIBuildStageRequest := _DescribeCodingCIBuildStageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeCodingCIBuildStageRequest)

	if err != nil {
		return err
	}

	*o = DescribeCodingCIBuildStageRequest(varDescribeCodingCIBuildStageRequest)

	return err
}

type NullableDescribeCodingCIBuildStageRequest struct {
	value *DescribeCodingCIBuildStageRequest
	isSet bool
}

func (v NullableDescribeCodingCIBuildStageRequest) Get() *DescribeCodingCIBuildStageRequest {
	return v.value
}

func (v *NullableDescribeCodingCIBuildStageRequest) Set(val *DescribeCodingCIBuildStageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCodingCIBuildStageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCodingCIBuildStageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCodingCIBuildStageRequest(val *DescribeCodingCIBuildStageRequest) *NullableDescribeCodingCIBuildStageRequest {
	return &NullableDescribeCodingCIBuildStageRequest{value: val, isSet: true}
}

func (v NullableDescribeCodingCIBuildStageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCodingCIBuildStageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


