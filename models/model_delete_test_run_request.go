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

// checks if the DeleteTestRunRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DeleteTestRunRequest{}

// DeleteTestRunRequest struct for DeleteTestRunRequest
type DeleteTestRunRequest struct {
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// 计划 ID
	RunId int32 `json:"RunId"`
}

type _DeleteTestRunRequest DeleteTestRunRequest

// NewDeleteTestRunRequest instantiates a new DeleteTestRunRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTestRunRequest(projectName string, runId int32) *DeleteTestRunRequest {
	this := DeleteTestRunRequest{}
	this.ProjectName = projectName
	this.RunId = runId
	return &this
}

// NewDeleteTestRunRequestWithDefaults instantiates a new DeleteTestRunRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTestRunRequestWithDefaults() *DeleteTestRunRequest {
	this := DeleteTestRunRequest{}
	return &this
}

// GetProjectName returns the ProjectName field value
func (o *DeleteTestRunRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *DeleteTestRunRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *DeleteTestRunRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetRunId returns the RunId field value
func (o *DeleteTestRunRequest) GetRunId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *DeleteTestRunRequest) GetRunIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *DeleteTestRunRequest) SetRunId(v int32) {
	o.RunId = v
}

func (o DeleteTestRunRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteTestRunRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ProjectName"] = o.ProjectName
	toSerialize["RunId"] = o.RunId
	return toSerialize, nil
}

func (o *DeleteTestRunRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ProjectName",
		"RunId",
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

	varDeleteTestRunRequest := _DeleteTestRunRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteTestRunRequest)

	if err != nil {
		return err
	}

	*o = DeleteTestRunRequest(varDeleteTestRunRequest)

	return err
}

type NullableDeleteTestRunRequest struct {
	value *DeleteTestRunRequest
	isSet bool
}

func (v NullableDeleteTestRunRequest) Get() *DeleteTestRunRequest {
	return v.value
}

func (v *NullableDeleteTestRunRequest) Set(val *DeleteTestRunRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTestRunRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTestRunRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTestRunRequest(val *DeleteTestRunRequest) *NullableDeleteTestRunRequest {
	return &NullableDeleteTestRunRequest{value: val, isSet: true}
}

func (v NullableDeleteTestRunRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTestRunRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


