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

// checks if the DescribeBranchProtectionMembersRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeBranchProtectionMembersRequest{}

// DescribeBranchProtectionMembersRequest struct for DescribeBranchProtectionMembersRequest
type DescribeBranchProtectionMembersRequest struct {
	// 保护分支规则id
	BranchProtectionId int64 `json:"BranchProtectionId"`
	// 仓库id
	DepotId int64 `json:"DepotId"`
	// 仓库路径，DepotId与DepptPath二选一即可
	DepotPath *string `json:"DepotPath,omitempty"`
}

type _DescribeBranchProtectionMembersRequest DescribeBranchProtectionMembersRequest

// NewDescribeBranchProtectionMembersRequest instantiates a new DescribeBranchProtectionMembersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeBranchProtectionMembersRequest(branchProtectionId int64, depotId int64) *DescribeBranchProtectionMembersRequest {
	this := DescribeBranchProtectionMembersRequest{}
	this.BranchProtectionId = branchProtectionId
	this.DepotId = depotId
	return &this
}

// NewDescribeBranchProtectionMembersRequestWithDefaults instantiates a new DescribeBranchProtectionMembersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeBranchProtectionMembersRequestWithDefaults() *DescribeBranchProtectionMembersRequest {
	this := DescribeBranchProtectionMembersRequest{}
	return &this
}

// GetBranchProtectionId returns the BranchProtectionId field value
func (o *DescribeBranchProtectionMembersRequest) GetBranchProtectionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BranchProtectionId
}

// GetBranchProtectionIdOk returns a tuple with the BranchProtectionId field value
// and a boolean to check if the value has been set.
func (o *DescribeBranchProtectionMembersRequest) GetBranchProtectionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchProtectionId, true
}

// SetBranchProtectionId sets field value
func (o *DescribeBranchProtectionMembersRequest) SetBranchProtectionId(v int64) {
	o.BranchProtectionId = v
}

// GetDepotId returns the DepotId field value
func (o *DescribeBranchProtectionMembersRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *DescribeBranchProtectionMembersRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *DescribeBranchProtectionMembersRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *DescribeBranchProtectionMembersRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeBranchProtectionMembersRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *DescribeBranchProtectionMembersRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *DescribeBranchProtectionMembersRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

func (o DescribeBranchProtectionMembersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeBranchProtectionMembersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["BranchProtectionId"] = o.BranchProtectionId
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	return toSerialize, nil
}

func (o *DescribeBranchProtectionMembersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"BranchProtectionId",
		"DepotId",
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

	varDescribeBranchProtectionMembersRequest := _DescribeBranchProtectionMembersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeBranchProtectionMembersRequest)

	if err != nil {
		return err
	}

	*o = DescribeBranchProtectionMembersRequest(varDescribeBranchProtectionMembersRequest)

	return err
}

type NullableDescribeBranchProtectionMembersRequest struct {
	value *DescribeBranchProtectionMembersRequest
	isSet bool
}

func (v NullableDescribeBranchProtectionMembersRequest) Get() *DescribeBranchProtectionMembersRequest {
	return v.value
}

func (v *NullableDescribeBranchProtectionMembersRequest) Set(val *DescribeBranchProtectionMembersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeBranchProtectionMembersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeBranchProtectionMembersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeBranchProtectionMembersRequest(val *DescribeBranchProtectionMembersRequest) *NullableDescribeBranchProtectionMembersRequest {
	return &NullableDescribeBranchProtectionMembersRequest{value: val, isSet: true}
}

func (v NullableDescribeBranchProtectionMembersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeBranchProtectionMembersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

