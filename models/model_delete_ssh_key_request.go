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

// checks if the DeleteSshKeyRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DeleteSshKeyRequest{}

// DeleteSshKeyRequest struct for DeleteSshKeyRequest
type DeleteSshKeyRequest struct {
	// ssh id
	SshKeyId int32 `json:"SshKeyId"`
}

type _DeleteSshKeyRequest DeleteSshKeyRequest

// NewDeleteSshKeyRequest instantiates a new DeleteSshKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSshKeyRequest(sshKeyId int32) *DeleteSshKeyRequest {
	this := DeleteSshKeyRequest{}
	this.SshKeyId = sshKeyId
	return &this
}

// NewDeleteSshKeyRequestWithDefaults instantiates a new DeleteSshKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSshKeyRequestWithDefaults() *DeleteSshKeyRequest {
	this := DeleteSshKeyRequest{}
	return &this
}

// GetSshKeyId returns the SshKeyId field value
func (o *DeleteSshKeyRequest) GetSshKeyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SshKeyId
}

// GetSshKeyIdOk returns a tuple with the SshKeyId field value
// and a boolean to check if the value has been set.
func (o *DeleteSshKeyRequest) GetSshKeyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SshKeyId, true
}

// SetSshKeyId sets field value
func (o *DeleteSshKeyRequest) SetSshKeyId(v int32) {
	o.SshKeyId = v
}

func (o DeleteSshKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteSshKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["SshKeyId"] = o.SshKeyId
	return toSerialize, nil
}

func (o *DeleteSshKeyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"SshKeyId",
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

	varDeleteSshKeyRequest := _DeleteSshKeyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteSshKeyRequest)

	if err != nil {
		return err
	}

	*o = DeleteSshKeyRequest(varDeleteSshKeyRequest)

	return err
}

type NullableDeleteSshKeyRequest struct {
	value *DeleteSshKeyRequest
	isSet bool
}

func (v NullableDeleteSshKeyRequest) Get() *DeleteSshKeyRequest {
	return v.value
}

func (v *NullableDeleteSshKeyRequest) Set(val *DeleteSshKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSshKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSshKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSshKeyRequest(val *DeleteSshKeyRequest) *NullableDeleteSshKeyRequest {
	return &NullableDeleteSshKeyRequest{value: val, isSet: true}
}

func (v NullableDeleteSshKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSshKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


