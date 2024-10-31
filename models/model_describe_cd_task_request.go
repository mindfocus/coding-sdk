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

// checks if the DescribeCdTaskRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCdTaskRequest{}

// DescribeCdTaskRequest struct for DescribeCdTaskRequest
type DescribeCdTaskRequest struct {
	// 任务执行记录 ID
	TaskExecutionId string `json:"TaskExecutionId"`
}

type _DescribeCdTaskRequest DescribeCdTaskRequest

// NewDescribeCdTaskRequest instantiates a new DescribeCdTaskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCdTaskRequest(taskExecutionId string) *DescribeCdTaskRequest {
	this := DescribeCdTaskRequest{}
	this.TaskExecutionId = taskExecutionId
	return &this
}

// NewDescribeCdTaskRequestWithDefaults instantiates a new DescribeCdTaskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCdTaskRequestWithDefaults() *DescribeCdTaskRequest {
	this := DescribeCdTaskRequest{}
	return &this
}

// GetTaskExecutionId returns the TaskExecutionId field value
func (o *DescribeCdTaskRequest) GetTaskExecutionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskExecutionId
}

// GetTaskExecutionIdOk returns a tuple with the TaskExecutionId field value
// and a boolean to check if the value has been set.
func (o *DescribeCdTaskRequest) GetTaskExecutionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskExecutionId, true
}

// SetTaskExecutionId sets field value
func (o *DescribeCdTaskRequest) SetTaskExecutionId(v string) {
	o.TaskExecutionId = v
}

func (o DescribeCdTaskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCdTaskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["TaskExecutionId"] = o.TaskExecutionId
	return toSerialize, nil
}

func (o *DescribeCdTaskRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"TaskExecutionId",
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

	varDescribeCdTaskRequest := _DescribeCdTaskRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeCdTaskRequest)

	if err != nil {
		return err
	}

	*o = DescribeCdTaskRequest(varDescribeCdTaskRequest)

	return err
}

type NullableDescribeCdTaskRequest struct {
	value *DescribeCdTaskRequest
	isSet bool
}

func (v NullableDescribeCdTaskRequest) Get() *DescribeCdTaskRequest {
	return v.value
}

func (v *NullableDescribeCdTaskRequest) Set(val *DescribeCdTaskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCdTaskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCdTaskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCdTaskRequest(val *DescribeCdTaskRequest) *NullableDescribeCdTaskRequest {
	return &NullableDescribeCdTaskRequest{value: val, isSet: true}
}

func (v NullableDescribeCdTaskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCdTaskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

