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

// checks if the DeleteTeamMemberRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DeleteTeamMemberRequest{}

// DeleteTeamMemberRequest struct for DeleteTeamMemberRequest
type DeleteTeamMemberRequest struct {
	// 用户Id
	UserId int32 `json:"UserId"`
}

type _DeleteTeamMemberRequest DeleteTeamMemberRequest

// NewDeleteTeamMemberRequest instantiates a new DeleteTeamMemberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTeamMemberRequest(userId int32) *DeleteTeamMemberRequest {
	this := DeleteTeamMemberRequest{}
	this.UserId = userId
	return &this
}

// NewDeleteTeamMemberRequestWithDefaults instantiates a new DeleteTeamMemberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTeamMemberRequestWithDefaults() *DeleteTeamMemberRequest {
	this := DeleteTeamMemberRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *DeleteTeamMemberRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *DeleteTeamMemberRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *DeleteTeamMemberRequest) SetUserId(v int32) {
	o.UserId = v
}

func (o DeleteTeamMemberRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteTeamMemberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["UserId"] = o.UserId
	return toSerialize, nil
}

func (o *DeleteTeamMemberRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varDeleteTeamMemberRequest := _DeleteTeamMemberRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteTeamMemberRequest)

	if err != nil {
		return err
	}

	*o = DeleteTeamMemberRequest(varDeleteTeamMemberRequest)

	return err
}

type NullableDeleteTeamMemberRequest struct {
	value *DeleteTeamMemberRequest
	isSet bool
}

func (v NullableDeleteTeamMemberRequest) Get() *DeleteTeamMemberRequest {
	return v.value
}

func (v *NullableDeleteTeamMemberRequest) Set(val *DeleteTeamMemberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTeamMemberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTeamMemberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTeamMemberRequest(val *DeleteTeamMemberRequest) *NullableDeleteTeamMemberRequest {
	return &NullableDeleteTeamMemberRequest{value: val, isSet: true}
}

func (v NullableDeleteTeamMemberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTeamMemberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


