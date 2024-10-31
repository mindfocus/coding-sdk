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

// checks if the CreateUserGroupRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateUserGroupRequest{}

// CreateUserGroupRequest struct for CreateUserGroupRequest
type CreateUserGroupRequest struct {
	// 用户组描述
	Description string `json:"Description"`
	// 用户组名称
	Name string `json:"Name"`
}

type _CreateUserGroupRequest CreateUserGroupRequest

// NewCreateUserGroupRequest instantiates a new CreateUserGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserGroupRequest(description string, name string) *CreateUserGroupRequest {
	this := CreateUserGroupRequest{}
	this.Description = description
	this.Name = name
	return &this
}

// NewCreateUserGroupRequestWithDefaults instantiates a new CreateUserGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserGroupRequestWithDefaults() *CreateUserGroupRequest {
	this := CreateUserGroupRequest{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateUserGroupRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateUserGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateUserGroupRequest) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *CreateUserGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateUserGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateUserGroupRequest) SetName(v string) {
	o.Name = v
}

func (o CreateUserGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Description"] = o.Description
	toSerialize["Name"] = o.Name
	return toSerialize, nil
}

func (o *CreateUserGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Description",
		"Name",
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

	varCreateUserGroupRequest := _CreateUserGroupRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateUserGroupRequest)

	if err != nil {
		return err
	}

	*o = CreateUserGroupRequest(varCreateUserGroupRequest)

	return err
}

type NullableCreateUserGroupRequest struct {
	value *CreateUserGroupRequest
	isSet bool
}

func (v NullableCreateUserGroupRequest) Get() *CreateUserGroupRequest {
	return v.value
}

func (v *NullableCreateUserGroupRequest) Set(val *CreateUserGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserGroupRequest(val *CreateUserGroupRequest) *NullableCreateUserGroupRequest {
	return &NullableCreateUserGroupRequest{value: val, isSet: true}
}

func (v NullableCreateUserGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

