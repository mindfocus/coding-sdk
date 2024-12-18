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

// checks if the ModifyDepotDescriptionRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyDepotDescriptionRequest{}

// ModifyDepotDescriptionRequest struct for ModifyDepotDescriptionRequest
type ModifyDepotDescriptionRequest struct {
	// 仓库 ID
	DepotId int64 `json:"DepotId"`
	// 仓库路径，与仓库ID二选一
	DepotPath *string `json:"DepotPath,omitempty"`
	// 仓库描述信息
	Description string `json:"Description"`
	// 用户 ID
	UserId int64 `json:"UserId"`
}

type _ModifyDepotDescriptionRequest ModifyDepotDescriptionRequest

// NewModifyDepotDescriptionRequest instantiates a new ModifyDepotDescriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyDepotDescriptionRequest(depotId int64, description string, userId int64) *ModifyDepotDescriptionRequest {
	this := ModifyDepotDescriptionRequest{}
	this.DepotId = depotId
	this.Description = description
	this.UserId = userId
	return &this
}

// NewModifyDepotDescriptionRequestWithDefaults instantiates a new ModifyDepotDescriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyDepotDescriptionRequestWithDefaults() *ModifyDepotDescriptionRequest {
	this := ModifyDepotDescriptionRequest{}
	return &this
}

// GetDepotId returns the DepotId field value
func (o *ModifyDepotDescriptionRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *ModifyDepotDescriptionRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *ModifyDepotDescriptionRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *ModifyDepotDescriptionRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyDepotDescriptionRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *ModifyDepotDescriptionRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *ModifyDepotDescriptionRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

// GetDescription returns the Description field value
func (o *ModifyDepotDescriptionRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ModifyDepotDescriptionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ModifyDepotDescriptionRequest) SetDescription(v string) {
	o.Description = v
}

// GetUserId returns the UserId field value
func (o *ModifyDepotDescriptionRequest) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ModifyDepotDescriptionRequest) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ModifyDepotDescriptionRequest) SetUserId(v int64) {
	o.UserId = v
}

func (o ModifyDepotDescriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyDepotDescriptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	toSerialize["Description"] = o.Description
	toSerialize["UserId"] = o.UserId
	return toSerialize, nil
}

func (o *ModifyDepotDescriptionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepotId",
		"Description",
		"UserId",
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

	varModifyDepotDescriptionRequest := _ModifyDepotDescriptionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyDepotDescriptionRequest)

	if err != nil {
		return err
	}

	*o = ModifyDepotDescriptionRequest(varModifyDepotDescriptionRequest)

	return err
}

type NullableModifyDepotDescriptionRequest struct {
	value *ModifyDepotDescriptionRequest
	isSet bool
}

func (v NullableModifyDepotDescriptionRequest) Get() *ModifyDepotDescriptionRequest {
	return v.value
}

func (v *NullableModifyDepotDescriptionRequest) Set(val *ModifyDepotDescriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyDepotDescriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyDepotDescriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyDepotDescriptionRequest(val *ModifyDepotDescriptionRequest) *NullableModifyDepotDescriptionRequest {
	return &NullableModifyDepotDescriptionRequest{value: val, isSet: true}
}

func (v NullableModifyDepotDescriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyDepotDescriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


