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

// checks if the ModifyCloseMRRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyCloseMRRequest{}

// ModifyCloseMRRequest struct for ModifyCloseMRRequest
type ModifyCloseMRRequest struct {
	// 仓库id
	DepotId int64 `json:"DepotId"`
	// 仓库路径，DepotId与DepotPath二选一即可
	DepotPath *string `json:"DepotPath,omitempty"`
	// 合并请求id
	MergeId int64 `json:"MergeId"`
}

type _ModifyCloseMRRequest ModifyCloseMRRequest

// NewModifyCloseMRRequest instantiates a new ModifyCloseMRRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyCloseMRRequest(depotId int64, mergeId int64) *ModifyCloseMRRequest {
	this := ModifyCloseMRRequest{}
	this.DepotId = depotId
	this.MergeId = mergeId
	return &this
}

// NewModifyCloseMRRequestWithDefaults instantiates a new ModifyCloseMRRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyCloseMRRequestWithDefaults() *ModifyCloseMRRequest {
	this := ModifyCloseMRRequest{}
	return &this
}

// GetDepotId returns the DepotId field value
func (o *ModifyCloseMRRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *ModifyCloseMRRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *ModifyCloseMRRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *ModifyCloseMRRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyCloseMRRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *ModifyCloseMRRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *ModifyCloseMRRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

// GetMergeId returns the MergeId field value
func (o *ModifyCloseMRRequest) GetMergeId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MergeId
}

// GetMergeIdOk returns a tuple with the MergeId field value
// and a boolean to check if the value has been set.
func (o *ModifyCloseMRRequest) GetMergeIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MergeId, true
}

// SetMergeId sets field value
func (o *ModifyCloseMRRequest) SetMergeId(v int64) {
	o.MergeId = v
}

func (o ModifyCloseMRRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyCloseMRRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	toSerialize["MergeId"] = o.MergeId
	return toSerialize, nil
}

func (o *ModifyCloseMRRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepotId",
		"MergeId",
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

	varModifyCloseMRRequest := _ModifyCloseMRRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyCloseMRRequest)

	if err != nil {
		return err
	}

	*o = ModifyCloseMRRequest(varModifyCloseMRRequest)

	return err
}

type NullableModifyCloseMRRequest struct {
	value *ModifyCloseMRRequest
	isSet bool
}

func (v NullableModifyCloseMRRequest) Get() *ModifyCloseMRRequest {
	return v.value
}

func (v *NullableModifyCloseMRRequest) Set(val *ModifyCloseMRRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyCloseMRRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyCloseMRRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyCloseMRRequest(val *ModifyCloseMRRequest) *NullableModifyCloseMRRequest {
	return &NullableModifyCloseMRRequest{value: val, isSet: true}
}

func (v NullableModifyCloseMRRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyCloseMRRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


