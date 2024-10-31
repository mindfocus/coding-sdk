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

// checks if the DescribeArtifactCreditRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeArtifactCreditRequest{}

// DescribeArtifactCreditRequest struct for DescribeArtifactCreditRequest
type DescribeArtifactCreditRequest struct {
	// 授信清单ID
	ArtifactCreditId int64 `json:"ArtifactCreditId"`
}

type _DescribeArtifactCreditRequest DescribeArtifactCreditRequest

// NewDescribeArtifactCreditRequest instantiates a new DescribeArtifactCreditRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeArtifactCreditRequest(artifactCreditId int64) *DescribeArtifactCreditRequest {
	this := DescribeArtifactCreditRequest{}
	this.ArtifactCreditId = artifactCreditId
	return &this
}

// NewDescribeArtifactCreditRequestWithDefaults instantiates a new DescribeArtifactCreditRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeArtifactCreditRequestWithDefaults() *DescribeArtifactCreditRequest {
	this := DescribeArtifactCreditRequest{}
	var artifactCreditId int64 = 0
	this.ArtifactCreditId = artifactCreditId
	return &this
}

// GetArtifactCreditId returns the ArtifactCreditId field value
func (o *DescribeArtifactCreditRequest) GetArtifactCreditId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ArtifactCreditId
}

// GetArtifactCreditIdOk returns a tuple with the ArtifactCreditId field value
// and a boolean to check if the value has been set.
func (o *DescribeArtifactCreditRequest) GetArtifactCreditIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactCreditId, true
}

// SetArtifactCreditId sets field value
func (o *DescribeArtifactCreditRequest) SetArtifactCreditId(v int64) {
	o.ArtifactCreditId = v
}

func (o DescribeArtifactCreditRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeArtifactCreditRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ArtifactCreditId"] = o.ArtifactCreditId
	return toSerialize, nil
}

func (o *DescribeArtifactCreditRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ArtifactCreditId",
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

	varDescribeArtifactCreditRequest := _DescribeArtifactCreditRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeArtifactCreditRequest)

	if err != nil {
		return err
	}

	*o = DescribeArtifactCreditRequest(varDescribeArtifactCreditRequest)

	return err
}

type NullableDescribeArtifactCreditRequest struct {
	value *DescribeArtifactCreditRequest
	isSet bool
}

func (v NullableDescribeArtifactCreditRequest) Get() *DescribeArtifactCreditRequest {
	return v.value
}

func (v *NullableDescribeArtifactCreditRequest) Set(val *DescribeArtifactCreditRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeArtifactCreditRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeArtifactCreditRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeArtifactCreditRequest(val *DescribeArtifactCreditRequest) *NullableDescribeArtifactCreditRequest {
	return &NullableDescribeArtifactCreditRequest{value: val, isSet: true}
}

func (v NullableDescribeArtifactCreditRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeArtifactCreditRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


