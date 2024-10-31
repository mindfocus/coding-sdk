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

// checks if the DeleteAPIDocRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DeleteAPIDocRequest{}

// DeleteAPIDocRequest struct for DeleteAPIDocRequest
type DeleteAPIDocRequest struct {
	// API 文档编号
	Code int32 `json:"Code"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
}

type _DeleteAPIDocRequest DeleteAPIDocRequest

// NewDeleteAPIDocRequest instantiates a new DeleteAPIDocRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAPIDocRequest(code int32, projectName string) *DeleteAPIDocRequest {
	this := DeleteAPIDocRequest{}
	this.Code = code
	this.ProjectName = projectName
	return &this
}

// NewDeleteAPIDocRequestWithDefaults instantiates a new DeleteAPIDocRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAPIDocRequestWithDefaults() *DeleteAPIDocRequest {
	this := DeleteAPIDocRequest{}
	return &this
}

// GetCode returns the Code field value
func (o *DeleteAPIDocRequest) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *DeleteAPIDocRequest) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *DeleteAPIDocRequest) SetCode(v int32) {
	o.Code = v
}

// GetProjectName returns the ProjectName field value
func (o *DeleteAPIDocRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *DeleteAPIDocRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *DeleteAPIDocRequest) SetProjectName(v string) {
	o.ProjectName = v
}

func (o DeleteAPIDocRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteAPIDocRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Code"] = o.Code
	toSerialize["ProjectName"] = o.ProjectName
	return toSerialize, nil
}

func (o *DeleteAPIDocRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Code",
		"ProjectName",
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

	varDeleteAPIDocRequest := _DeleteAPIDocRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteAPIDocRequest)

	if err != nil {
		return err
	}

	*o = DeleteAPIDocRequest(varDeleteAPIDocRequest)

	return err
}

type NullableDeleteAPIDocRequest struct {
	value *DeleteAPIDocRequest
	isSet bool
}

func (v NullableDeleteAPIDocRequest) Get() *DeleteAPIDocRequest {
	return v.value
}

func (v *NullableDeleteAPIDocRequest) Set(val *DeleteAPIDocRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAPIDocRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAPIDocRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAPIDocRequest(val *DeleteAPIDocRequest) *NullableDeleteAPIDocRequest {
	return &NullableDeleteAPIDocRequest{value: val, isSet: true}
}

func (v NullableDeleteAPIDocRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAPIDocRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


