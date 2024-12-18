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

// checks if the JobCreateJobData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &JobCreateJobData{}

// JobCreateJobData struct for JobCreateJobData
type JobCreateJobData struct {
	// 构建计划 ID
	Id int64 `json:"Id"`
}

type _JobCreateJobData JobCreateJobData

// NewJobCreateJobData instantiates a new JobCreateJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobCreateJobData(id int64) *JobCreateJobData {
	this := JobCreateJobData{}
	this.Id = id
	return &this
}

// NewJobCreateJobDataWithDefaults instantiates a new JobCreateJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobCreateJobDataWithDefaults() *JobCreateJobData {
	this := JobCreateJobData{}
	var id int64 = 0
	this.Id = id
	return &this
}

// GetId returns the Id field value
func (o *JobCreateJobData) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobCreateJobData) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JobCreateJobData) SetId(v int64) {
	o.Id = v
}

func (o JobCreateJobData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobCreateJobData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	return toSerialize, nil
}

func (o *JobCreateJobData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Id",
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

	varJobCreateJobData := _JobCreateJobData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJobCreateJobData)

	if err != nil {
		return err
	}

	*o = JobCreateJobData(varJobCreateJobData)

	return err
}

type NullableJobCreateJobData struct {
	value *JobCreateJobData
	isSet bool
}

func (v NullableJobCreateJobData) Get() *JobCreateJobData {
	return v.value
}

func (v *NullableJobCreateJobData) Set(val *JobCreateJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableJobCreateJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableJobCreateJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobCreateJobData(val *JobCreateJobData) *NullableJobCreateJobData {
	return &NullableJobCreateJobData{value: val, isSet: true}
}

func (v NullableJobCreateJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobCreateJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


