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

// checks if the ModifyIssueDescriptionRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyIssueDescriptionRequest{}

// ModifyIssueDescriptionRequest struct for ModifyIssueDescriptionRequest
type ModifyIssueDescriptionRequest struct {
	// 事项描述
	Description string `json:"Description"`
	// 事项 Code
	IssueCode int64 `json:"IssueCode"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
}

type _ModifyIssueDescriptionRequest ModifyIssueDescriptionRequest

// NewModifyIssueDescriptionRequest instantiates a new ModifyIssueDescriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyIssueDescriptionRequest(description string, issueCode int64, projectName string) *ModifyIssueDescriptionRequest {
	this := ModifyIssueDescriptionRequest{}
	this.Description = description
	this.IssueCode = issueCode
	this.ProjectName = projectName
	return &this
}

// NewModifyIssueDescriptionRequestWithDefaults instantiates a new ModifyIssueDescriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyIssueDescriptionRequestWithDefaults() *ModifyIssueDescriptionRequest {
	this := ModifyIssueDescriptionRequest{}
	return &this
}

// GetDescription returns the Description field value
func (o *ModifyIssueDescriptionRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ModifyIssueDescriptionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ModifyIssueDescriptionRequest) SetDescription(v string) {
	o.Description = v
}

// GetIssueCode returns the IssueCode field value
func (o *ModifyIssueDescriptionRequest) GetIssueCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IssueCode
}

// GetIssueCodeOk returns a tuple with the IssueCode field value
// and a boolean to check if the value has been set.
func (o *ModifyIssueDescriptionRequest) GetIssueCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueCode, true
}

// SetIssueCode sets field value
func (o *ModifyIssueDescriptionRequest) SetIssueCode(v int64) {
	o.IssueCode = v
}

// GetProjectName returns the ProjectName field value
func (o *ModifyIssueDescriptionRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *ModifyIssueDescriptionRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *ModifyIssueDescriptionRequest) SetProjectName(v string) {
	o.ProjectName = v
}

func (o ModifyIssueDescriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyIssueDescriptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Description"] = o.Description
	toSerialize["IssueCode"] = o.IssueCode
	toSerialize["ProjectName"] = o.ProjectName
	return toSerialize, nil
}

func (o *ModifyIssueDescriptionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Description",
		"IssueCode",
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

	varModifyIssueDescriptionRequest := _ModifyIssueDescriptionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyIssueDescriptionRequest)

	if err != nil {
		return err
	}

	*o = ModifyIssueDescriptionRequest(varModifyIssueDescriptionRequest)

	return err
}

type NullableModifyIssueDescriptionRequest struct {
	value *ModifyIssueDescriptionRequest
	isSet bool
}

func (v NullableModifyIssueDescriptionRequest) Get() *ModifyIssueDescriptionRequest {
	return v.value
}

func (v *NullableModifyIssueDescriptionRequest) Set(val *ModifyIssueDescriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyIssueDescriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyIssueDescriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyIssueDescriptionRequest(val *ModifyIssueDescriptionRequest) *NullableModifyIssueDescriptionRequest {
	return &NullableModifyIssueDescriptionRequest{value: val, isSet: true}
}

func (v NullableModifyIssueDescriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyIssueDescriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


