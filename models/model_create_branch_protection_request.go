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

// checks if the CreateBranchProtectionRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateBranchProtectionRequest{}

// CreateBranchProtectionRequest struct for CreateBranchProtectionRequest
type CreateBranchProtectionRequest struct {
	// 仓库id
	DepotId int64 `json:"DepotId"`
	// 仓库路径
	DepotPath *string `json:"DepotPath,omitempty"`
	// 保护规则
	Rule string `json:"Rule"`
}

type _CreateBranchProtectionRequest CreateBranchProtectionRequest

// NewCreateBranchProtectionRequest instantiates a new CreateBranchProtectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBranchProtectionRequest(depotId int64, rule string) *CreateBranchProtectionRequest {
	this := CreateBranchProtectionRequest{}
	this.DepotId = depotId
	this.Rule = rule
	return &this
}

// NewCreateBranchProtectionRequestWithDefaults instantiates a new CreateBranchProtectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBranchProtectionRequestWithDefaults() *CreateBranchProtectionRequest {
	this := CreateBranchProtectionRequest{}
	return &this
}

// GetDepotId returns the DepotId field value
func (o *CreateBranchProtectionRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *CreateBranchProtectionRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *CreateBranchProtectionRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *CreateBranchProtectionRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *CreateBranchProtectionRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

// GetRule returns the Rule field value
func (o *CreateBranchProtectionRequest) GetRule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *CreateBranchProtectionRequest) GetRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value
func (o *CreateBranchProtectionRequest) SetRule(v string) {
	o.Rule = v
}

func (o CreateBranchProtectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBranchProtectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	toSerialize["Rule"] = o.Rule
	return toSerialize, nil
}

func (o *CreateBranchProtectionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepotId",
		"Rule",
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

	varCreateBranchProtectionRequest := _CreateBranchProtectionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateBranchProtectionRequest)

	if err != nil {
		return err
	}

	*o = CreateBranchProtectionRequest(varCreateBranchProtectionRequest)

	return err
}

type NullableCreateBranchProtectionRequest struct {
	value *CreateBranchProtectionRequest
	isSet bool
}

func (v NullableCreateBranchProtectionRequest) Get() *CreateBranchProtectionRequest {
	return v.value
}

func (v *NullableCreateBranchProtectionRequest) Set(val *CreateBranchProtectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBranchProtectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBranchProtectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBranchProtectionRequest(val *CreateBranchProtectionRequest) *NullableCreateBranchProtectionRequest {
	return &NullableCreateBranchProtectionRequest{value: val, isSet: true}
}

func (v NullableCreateBranchProtectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBranchProtectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

