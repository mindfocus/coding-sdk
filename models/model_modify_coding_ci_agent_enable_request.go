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

// checks if the ModifyCodingCIAgentEnableRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyCodingCIAgentEnableRequest{}

// ModifyCodingCIAgentEnableRequest struct for ModifyCodingCIAgentEnableRequest
type ModifyCodingCIAgentEnableRequest struct {
	// 节点池 ID
	PoolId int64 `json:"PoolId"`
	// 是否启用：启用true、禁用false
	Enable bool `json:"Enable"`
	// 节点ID
	Id int64 `json:"Id"`
}

type _ModifyCodingCIAgentEnableRequest ModifyCodingCIAgentEnableRequest

// NewModifyCodingCIAgentEnableRequest instantiates a new ModifyCodingCIAgentEnableRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyCodingCIAgentEnableRequest(poolId int64, enable bool, id int64) *ModifyCodingCIAgentEnableRequest {
	this := ModifyCodingCIAgentEnableRequest{}
	this.PoolId = poolId
	this.Enable = enable
	this.Id = id
	return &this
}

// NewModifyCodingCIAgentEnableRequestWithDefaults instantiates a new ModifyCodingCIAgentEnableRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyCodingCIAgentEnableRequestWithDefaults() *ModifyCodingCIAgentEnableRequest {
	this := ModifyCodingCIAgentEnableRequest{}
	var poolId int64 = 0
	this.PoolId = poolId
	var enable bool = false
	this.Enable = enable
	var id int64 = 0
	this.Id = id
	return &this
}

// GetPoolId returns the PoolId field value
func (o *ModifyCodingCIAgentEnableRequest) GetPoolId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value
// and a boolean to check if the value has been set.
func (o *ModifyCodingCIAgentEnableRequest) GetPoolIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolId, true
}

// SetPoolId sets field value
func (o *ModifyCodingCIAgentEnableRequest) SetPoolId(v int64) {
	o.PoolId = v
}

// GetEnable returns the Enable field value
func (o *ModifyCodingCIAgentEnableRequest) GetEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enable
}

// GetEnableOk returns a tuple with the Enable field value
// and a boolean to check if the value has been set.
func (o *ModifyCodingCIAgentEnableRequest) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enable, true
}

// SetEnable sets field value
func (o *ModifyCodingCIAgentEnableRequest) SetEnable(v bool) {
	o.Enable = v
}

// GetId returns the Id field value
func (o *ModifyCodingCIAgentEnableRequest) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModifyCodingCIAgentEnableRequest) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModifyCodingCIAgentEnableRequest) SetId(v int64) {
	o.Id = v
}

func (o ModifyCodingCIAgentEnableRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyCodingCIAgentEnableRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["PoolId"] = o.PoolId
	toSerialize["Enable"] = o.Enable
	toSerialize["Id"] = o.Id
	return toSerialize, nil
}

func (o *ModifyCodingCIAgentEnableRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"PoolId",
		"Enable",
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

	varModifyCodingCIAgentEnableRequest := _ModifyCodingCIAgentEnableRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyCodingCIAgentEnableRequest)

	if err != nil {
		return err
	}

	*o = ModifyCodingCIAgentEnableRequest(varModifyCodingCIAgentEnableRequest)

	return err
}

type NullableModifyCodingCIAgentEnableRequest struct {
	value *ModifyCodingCIAgentEnableRequest
	isSet bool
}

func (v NullableModifyCodingCIAgentEnableRequest) Get() *ModifyCodingCIAgentEnableRequest {
	return v.value
}

func (v *NullableModifyCodingCIAgentEnableRequest) Set(val *ModifyCodingCIAgentEnableRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyCodingCIAgentEnableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyCodingCIAgentEnableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyCodingCIAgentEnableRequest(val *ModifyCodingCIAgentEnableRequest) *NullableModifyCodingCIAgentEnableRequest {
	return &NullableModifyCodingCIAgentEnableRequest{value: val, isSet: true}
}

func (v NullableModifyCodingCIAgentEnableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyCodingCIAgentEnableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


