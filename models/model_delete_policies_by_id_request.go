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

// checks if the DeletePoliciesByIdRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DeletePoliciesByIdRequest{}

// DeletePoliciesByIdRequest struct for DeletePoliciesByIdRequest
type DeletePoliciesByIdRequest struct {
	// 权限组 ID
	Id []int64 `json:"Id"`
}

type _DeletePoliciesByIdRequest DeletePoliciesByIdRequest

// NewDeletePoliciesByIdRequest instantiates a new DeletePoliciesByIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletePoliciesByIdRequest(id []int64) *DeletePoliciesByIdRequest {
	this := DeletePoliciesByIdRequest{}
	this.Id = id
	return &this
}

// NewDeletePoliciesByIdRequestWithDefaults instantiates a new DeletePoliciesByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletePoliciesByIdRequestWithDefaults() *DeletePoliciesByIdRequest {
	this := DeletePoliciesByIdRequest{}
	return &this
}

// GetId returns the Id field value
func (o *DeletePoliciesByIdRequest) GetId() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeletePoliciesByIdRequest) GetIdOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets field value
func (o *DeletePoliciesByIdRequest) SetId(v []int64) {
	o.Id = v
}

func (o DeletePoliciesByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeletePoliciesByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	return toSerialize, nil
}

func (o *DeletePoliciesByIdRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varDeletePoliciesByIdRequest := _DeletePoliciesByIdRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeletePoliciesByIdRequest)

	if err != nil {
		return err
	}

	*o = DeletePoliciesByIdRequest(varDeletePoliciesByIdRequest)

	return err
}

type NullableDeletePoliciesByIdRequest struct {
	value *DeletePoliciesByIdRequest
	isSet bool
}

func (v NullableDeletePoliciesByIdRequest) Get() *DeletePoliciesByIdRequest {
	return v.value
}

func (v *NullableDeletePoliciesByIdRequest) Set(val *DeletePoliciesByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletePoliciesByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletePoliciesByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletePoliciesByIdRequest(val *DeletePoliciesByIdRequest) *NullableDeletePoliciesByIdRequest {
	return &NullableDeletePoliciesByIdRequest{value: val, isSet: true}
}

func (v NullableDeletePoliciesByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletePoliciesByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


