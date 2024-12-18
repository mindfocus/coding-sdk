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

// checks if the DescribeAllProjectLabelsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeAllProjectLabelsRequest{}

// DescribeAllProjectLabelsRequest struct for DescribeAllProjectLabelsRequest
type DescribeAllProjectLabelsRequest struct {
	// 项目ID
	ProjectId int64 `json:"ProjectId"`
}

type _DescribeAllProjectLabelsRequest DescribeAllProjectLabelsRequest

// NewDescribeAllProjectLabelsRequest instantiates a new DescribeAllProjectLabelsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeAllProjectLabelsRequest(projectId int64) *DescribeAllProjectLabelsRequest {
	this := DescribeAllProjectLabelsRequest{}
	this.ProjectId = projectId
	return &this
}

// NewDescribeAllProjectLabelsRequestWithDefaults instantiates a new DescribeAllProjectLabelsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeAllProjectLabelsRequestWithDefaults() *DescribeAllProjectLabelsRequest {
	this := DescribeAllProjectLabelsRequest{}
	var projectId int64 = 0
	this.ProjectId = projectId
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *DescribeAllProjectLabelsRequest) GetProjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DescribeAllProjectLabelsRequest) GetProjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DescribeAllProjectLabelsRequest) SetProjectId(v int64) {
	o.ProjectId = v
}

func (o DescribeAllProjectLabelsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeAllProjectLabelsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ProjectId"] = o.ProjectId
	return toSerialize, nil
}

func (o *DescribeAllProjectLabelsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ProjectId",
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

	varDescribeAllProjectLabelsRequest := _DescribeAllProjectLabelsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeAllProjectLabelsRequest)

	if err != nil {
		return err
	}

	*o = DescribeAllProjectLabelsRequest(varDescribeAllProjectLabelsRequest)

	return err
}

type NullableDescribeAllProjectLabelsRequest struct {
	value *DescribeAllProjectLabelsRequest
	isSet bool
}

func (v NullableDescribeAllProjectLabelsRequest) Get() *DescribeAllProjectLabelsRequest {
	return v.value
}

func (v *NullableDescribeAllProjectLabelsRequest) Set(val *DescribeAllProjectLabelsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeAllProjectLabelsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeAllProjectLabelsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeAllProjectLabelsRequest(val *DescribeAllProjectLabelsRequest) *NullableDescribeAllProjectLabelsRequest {
	return &NullableDescribeAllProjectLabelsRequest{value: val, isSet: true}
}

func (v NullableDescribeAllProjectLabelsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeAllProjectLabelsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


