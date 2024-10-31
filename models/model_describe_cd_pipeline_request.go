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

// checks if the DescribeCdPipelineRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCdPipelineRequest{}

// DescribeCdPipelineRequest struct for DescribeCdPipelineRequest
type DescribeCdPipelineRequest struct {
	// 部署流程执行记录 ID
	PipelineExecutionId string `json:"PipelineExecutionId"`
}

type _DescribeCdPipelineRequest DescribeCdPipelineRequest

// NewDescribeCdPipelineRequest instantiates a new DescribeCdPipelineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCdPipelineRequest(pipelineExecutionId string) *DescribeCdPipelineRequest {
	this := DescribeCdPipelineRequest{}
	this.PipelineExecutionId = pipelineExecutionId
	return &this
}

// NewDescribeCdPipelineRequestWithDefaults instantiates a new DescribeCdPipelineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCdPipelineRequestWithDefaults() *DescribeCdPipelineRequest {
	this := DescribeCdPipelineRequest{}
	return &this
}

// GetPipelineExecutionId returns the PipelineExecutionId field value
func (o *DescribeCdPipelineRequest) GetPipelineExecutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PipelineExecutionId
}

// GetPipelineExecutionIdOk returns a tuple with the PipelineExecutionId field value
// and a boolean to check if the value has been set.
func (o *DescribeCdPipelineRequest) GetPipelineExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PipelineExecutionId, true
}

// SetPipelineExecutionId sets field value
func (o *DescribeCdPipelineRequest) SetPipelineExecutionId(v string) {
	o.PipelineExecutionId = v
}

func (o DescribeCdPipelineRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCdPipelineRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["PipelineExecutionId"] = o.PipelineExecutionId
	return toSerialize, nil
}

func (o *DescribeCdPipelineRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"PipelineExecutionId",
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

	varDescribeCdPipelineRequest := _DescribeCdPipelineRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeCdPipelineRequest)

	if err != nil {
		return err
	}

	*o = DescribeCdPipelineRequest(varDescribeCdPipelineRequest)

	return err
}

type NullableDescribeCdPipelineRequest struct {
	value *DescribeCdPipelineRequest
	isSet bool
}

func (v NullableDescribeCdPipelineRequest) Get() *DescribeCdPipelineRequest {
	return v.value
}

func (v *NullableDescribeCdPipelineRequest) Set(val *DescribeCdPipelineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCdPipelineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCdPipelineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCdPipelineRequest(val *DescribeCdPipelineRequest) *NullableDescribeCdPipelineRequest {
	return &NullableDescribeCdPipelineRequest{value: val, isSet: true}
}

func (v NullableDescribeCdPipelineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCdPipelineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

