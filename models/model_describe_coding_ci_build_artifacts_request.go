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

// checks if the DescribeCodingCIBuildArtifactsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCodingCIBuildArtifactsRequest{}

// DescribeCodingCIBuildArtifactsRequest struct for DescribeCodingCIBuildArtifactsRequest
type DescribeCodingCIBuildArtifactsRequest struct {
	// 构建 ID
	BuildId int64 `json:"BuildId"`
}

type _DescribeCodingCIBuildArtifactsRequest DescribeCodingCIBuildArtifactsRequest

// NewDescribeCodingCIBuildArtifactsRequest instantiates a new DescribeCodingCIBuildArtifactsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCodingCIBuildArtifactsRequest(buildId int64) *DescribeCodingCIBuildArtifactsRequest {
	this := DescribeCodingCIBuildArtifactsRequest{}
	this.BuildId = buildId
	return &this
}

// NewDescribeCodingCIBuildArtifactsRequestWithDefaults instantiates a new DescribeCodingCIBuildArtifactsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCodingCIBuildArtifactsRequestWithDefaults() *DescribeCodingCIBuildArtifactsRequest {
	this := DescribeCodingCIBuildArtifactsRequest{}
	var buildId int64 = 0
	this.BuildId = buildId
	return &this
}

// GetBuildId returns the BuildId field value
func (o *DescribeCodingCIBuildArtifactsRequest) GetBuildId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BuildId
}

// GetBuildIdOk returns a tuple with the BuildId field value
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildArtifactsRequest) GetBuildIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildId, true
}

// SetBuildId sets field value
func (o *DescribeCodingCIBuildArtifactsRequest) SetBuildId(v int64) {
	o.BuildId = v
}

func (o DescribeCodingCIBuildArtifactsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCodingCIBuildArtifactsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["BuildId"] = o.BuildId
	return toSerialize, nil
}

func (o *DescribeCodingCIBuildArtifactsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"BuildId",
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

	varDescribeCodingCIBuildArtifactsRequest := _DescribeCodingCIBuildArtifactsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeCodingCIBuildArtifactsRequest)

	if err != nil {
		return err
	}

	*o = DescribeCodingCIBuildArtifactsRequest(varDescribeCodingCIBuildArtifactsRequest)

	return err
}

type NullableDescribeCodingCIBuildArtifactsRequest struct {
	value *DescribeCodingCIBuildArtifactsRequest
	isSet bool
}

func (v NullableDescribeCodingCIBuildArtifactsRequest) Get() *DescribeCodingCIBuildArtifactsRequest {
	return v.value
}

func (v *NullableDescribeCodingCIBuildArtifactsRequest) Set(val *DescribeCodingCIBuildArtifactsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCodingCIBuildArtifactsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCodingCIBuildArtifactsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCodingCIBuildArtifactsRequest(val *DescribeCodingCIBuildArtifactsRequest) *NullableDescribeCodingCIBuildArtifactsRequest {
	return &NullableDescribeCodingCIBuildArtifactsRequest{value: val, isSet: true}
}

func (v NullableDescribeCodingCIBuildArtifactsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCodingCIBuildArtifactsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

