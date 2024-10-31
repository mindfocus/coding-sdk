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

// checks if the DescribeCodingCIBuildStepRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCodingCIBuildStepRequest{}

// DescribeCodingCIBuildStepRequest struct for DescribeCodingCIBuildStepRequest
type DescribeCodingCIBuildStepRequest struct {
	// 构建 ID
	BuildId int32 `json:"BuildId"`
	// 阶段 ID
	StageId int32 `json:"StageId"`
}

type _DescribeCodingCIBuildStepRequest DescribeCodingCIBuildStepRequest

// NewDescribeCodingCIBuildStepRequest instantiates a new DescribeCodingCIBuildStepRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCodingCIBuildStepRequest(buildId int32, stageId int32) *DescribeCodingCIBuildStepRequest {
	this := DescribeCodingCIBuildStepRequest{}
	this.BuildId = buildId
	this.StageId = stageId
	return &this
}

// NewDescribeCodingCIBuildStepRequestWithDefaults instantiates a new DescribeCodingCIBuildStepRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCodingCIBuildStepRequestWithDefaults() *DescribeCodingCIBuildStepRequest {
	this := DescribeCodingCIBuildStepRequest{}
	return &this
}

// GetBuildId returns the BuildId field value
func (o *DescribeCodingCIBuildStepRequest) GetBuildId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BuildId
}

// GetBuildIdOk returns a tuple with the BuildId field value
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildStepRequest) GetBuildIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildId, true
}

// SetBuildId sets field value
func (o *DescribeCodingCIBuildStepRequest) SetBuildId(v int32) {
	o.BuildId = v
}

// GetStageId returns the StageId field value
func (o *DescribeCodingCIBuildStepRequest) GetStageId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StageId
}

// GetStageIdOk returns a tuple with the StageId field value
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildStepRequest) GetStageIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StageId, true
}

// SetStageId sets field value
func (o *DescribeCodingCIBuildStepRequest) SetStageId(v int32) {
	o.StageId = v
}

func (o DescribeCodingCIBuildStepRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCodingCIBuildStepRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["BuildId"] = o.BuildId
	toSerialize["StageId"] = o.StageId
	return toSerialize, nil
}

func (o *DescribeCodingCIBuildStepRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"BuildId",
		"StageId",
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

	varDescribeCodingCIBuildStepRequest := _DescribeCodingCIBuildStepRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeCodingCIBuildStepRequest)

	if err != nil {
		return err
	}

	*o = DescribeCodingCIBuildStepRequest(varDescribeCodingCIBuildStepRequest)

	return err
}

type NullableDescribeCodingCIBuildStepRequest struct {
	value *DescribeCodingCIBuildStepRequest
	isSet bool
}

func (v NullableDescribeCodingCIBuildStepRequest) Get() *DescribeCodingCIBuildStepRequest {
	return v.value
}

func (v *NullableDescribeCodingCIBuildStepRequest) Set(val *DescribeCodingCIBuildStepRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCodingCIBuildStepRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCodingCIBuildStepRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCodingCIBuildStepRequest(val *DescribeCodingCIBuildStepRequest) *NullableDescribeCodingCIBuildStepRequest {
	return &NullableDescribeCodingCIBuildStepRequest{value: val, isSet: true}
}

func (v NullableDescribeCodingCIBuildStepRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCodingCIBuildStepRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


