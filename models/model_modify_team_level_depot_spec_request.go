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

// checks if the ModifyTeamLevelDepotSpecRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyTeamLevelDepotSpecRequest{}

// ModifyTeamLevelDepotSpecRequest struct for ModifyTeamLevelDepotSpecRequest
type ModifyTeamLevelDepotSpecRequest struct {
	Param DepotSpecTeamLevelParam `json:"Param"`
}

type _ModifyTeamLevelDepotSpecRequest ModifyTeamLevelDepotSpecRequest

// NewModifyTeamLevelDepotSpecRequest instantiates a new ModifyTeamLevelDepotSpecRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyTeamLevelDepotSpecRequest(param DepotSpecTeamLevelParam) *ModifyTeamLevelDepotSpecRequest {
	this := ModifyTeamLevelDepotSpecRequest{}
	this.Param = param
	return &this
}

// NewModifyTeamLevelDepotSpecRequestWithDefaults instantiates a new ModifyTeamLevelDepotSpecRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyTeamLevelDepotSpecRequestWithDefaults() *ModifyTeamLevelDepotSpecRequest {
	this := ModifyTeamLevelDepotSpecRequest{}
	return &this
}

// GetParam returns the Param field value
func (o *ModifyTeamLevelDepotSpecRequest) GetParam() DepotSpecTeamLevelParam {
	if o == nil {
		var ret DepotSpecTeamLevelParam
		return ret
	}

	return o.Param
}

// GetParamOk returns a tuple with the Param field value
// and a boolean to check if the value has been set.
func (o *ModifyTeamLevelDepotSpecRequest) GetParamOk() (*DepotSpecTeamLevelParam, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Param, true
}

// SetParam sets field value
func (o *ModifyTeamLevelDepotSpecRequest) SetParam(v DepotSpecTeamLevelParam) {
	o.Param = v
}

func (o ModifyTeamLevelDepotSpecRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyTeamLevelDepotSpecRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Param"] = o.Param
	return toSerialize, nil
}

func (o *ModifyTeamLevelDepotSpecRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Param",
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

	varModifyTeamLevelDepotSpecRequest := _ModifyTeamLevelDepotSpecRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyTeamLevelDepotSpecRequest)

	if err != nil {
		return err
	}

	*o = ModifyTeamLevelDepotSpecRequest(varModifyTeamLevelDepotSpecRequest)

	return err
}

type NullableModifyTeamLevelDepotSpecRequest struct {
	value *ModifyTeamLevelDepotSpecRequest
	isSet bool
}

func (v NullableModifyTeamLevelDepotSpecRequest) Get() *ModifyTeamLevelDepotSpecRequest {
	return v.value
}

func (v *NullableModifyTeamLevelDepotSpecRequest) Set(val *ModifyTeamLevelDepotSpecRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyTeamLevelDepotSpecRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyTeamLevelDepotSpecRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyTeamLevelDepotSpecRequest(val *ModifyTeamLevelDepotSpecRequest) *NullableModifyTeamLevelDepotSpecRequest {
	return &NullableModifyTeamLevelDepotSpecRequest{value: val, isSet: true}
}

func (v NullableModifyTeamLevelDepotSpecRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyTeamLevelDepotSpecRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


