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

// checks if the CodingCIBuildStepData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CodingCIBuildStepData{}

// CodingCIBuildStepData 获取构建任务指定阶段的步骤
type CodingCIBuildStepData struct {
	// 步骤
	StepJsonString string `json:"StepJsonString"`
}

type _CodingCIBuildStepData CodingCIBuildStepData

// NewCodingCIBuildStepData instantiates a new CodingCIBuildStepData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodingCIBuildStepData(stepJsonString string) *CodingCIBuildStepData {
	this := CodingCIBuildStepData{}
	this.StepJsonString = stepJsonString
	return &this
}

// NewCodingCIBuildStepDataWithDefaults instantiates a new CodingCIBuildStepData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodingCIBuildStepDataWithDefaults() *CodingCIBuildStepData {
	this := CodingCIBuildStepData{}
	var stepJsonString string = ""
	this.StepJsonString = stepJsonString
	return &this
}

// GetStepJsonString returns the StepJsonString field value
func (o *CodingCIBuildStepData) GetStepJsonString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StepJsonString
}

// GetStepJsonStringOk returns a tuple with the StepJsonString field value
// and a boolean to check if the value has been set.
func (o *CodingCIBuildStepData) GetStepJsonStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StepJsonString, true
}

// SetStepJsonString sets field value
func (o *CodingCIBuildStepData) SetStepJsonString(v string) {
	o.StepJsonString = v
}

func (o CodingCIBuildStepData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CodingCIBuildStepData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["StepJsonString"] = o.StepJsonString
	return toSerialize, nil
}

func (o *CodingCIBuildStepData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"StepJsonString",
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

	varCodingCIBuildStepData := _CodingCIBuildStepData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCodingCIBuildStepData)

	if err != nil {
		return err
	}

	*o = CodingCIBuildStepData(varCodingCIBuildStepData)

	return err
}

type NullableCodingCIBuildStepData struct {
	value *CodingCIBuildStepData
	isSet bool
}

func (v NullableCodingCIBuildStepData) Get() *CodingCIBuildStepData {
	return v.value
}

func (v *NullableCodingCIBuildStepData) Set(val *CodingCIBuildStepData) {
	v.value = val
	v.isSet = true
}

func (v NullableCodingCIBuildStepData) IsSet() bool {
	return v.isSet
}

func (v *NullableCodingCIBuildStepData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodingCIBuildStepData(val *CodingCIBuildStepData) *NullableCodingCIBuildStepData {
	return &NullableCodingCIBuildStepData{value: val, isSet: true}
}

func (v NullableCodingCIBuildStepData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodingCIBuildStepData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

