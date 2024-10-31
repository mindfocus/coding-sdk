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

// checks if the ServiceHookResendServiceHookResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ServiceHookResendServiceHookResult{}

// ServiceHookResendServiceHookResult struct for ServiceHookResendServiceHookResult
type ServiceHookResendServiceHookResult struct {
	// 是否成功
	Success bool `json:"Success"`
}

type _ServiceHookResendServiceHookResult ServiceHookResendServiceHookResult

// NewServiceHookResendServiceHookResult instantiates a new ServiceHookResendServiceHookResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceHookResendServiceHookResult(success bool) *ServiceHookResendServiceHookResult {
	this := ServiceHookResendServiceHookResult{}
	this.Success = success
	return &this
}

// NewServiceHookResendServiceHookResultWithDefaults instantiates a new ServiceHookResendServiceHookResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceHookResendServiceHookResultWithDefaults() *ServiceHookResendServiceHookResult {
	this := ServiceHookResendServiceHookResult{}
	var success bool = false
	this.Success = success
	return &this
}

// GetSuccess returns the Success field value
func (o *ServiceHookResendServiceHookResult) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ServiceHookResendServiceHookResult) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *ServiceHookResendServiceHookResult) SetSuccess(v bool) {
	o.Success = v
}

func (o ServiceHookResendServiceHookResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceHookResendServiceHookResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Success"] = o.Success
	return toSerialize, nil
}

func (o *ServiceHookResendServiceHookResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Success",
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

	varServiceHookResendServiceHookResult := _ServiceHookResendServiceHookResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceHookResendServiceHookResult)

	if err != nil {
		return err
	}

	*o = ServiceHookResendServiceHookResult(varServiceHookResendServiceHookResult)

	return err
}

type NullableServiceHookResendServiceHookResult struct {
	value *ServiceHookResendServiceHookResult
	isSet bool
}

func (v NullableServiceHookResendServiceHookResult) Get() *ServiceHookResendServiceHookResult {
	return v.value
}

func (v *NullableServiceHookResendServiceHookResult) Set(val *ServiceHookResendServiceHookResult) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceHookResendServiceHookResult) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceHookResendServiceHookResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceHookResendServiceHookResult(val *ServiceHookResendServiceHookResult) *NullableServiceHookResendServiceHookResult {
	return &NullableServiceHookResendServiceHookResult{value: val, isSet: true}
}

func (v NullableServiceHookResendServiceHookResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceHookResendServiceHookResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


