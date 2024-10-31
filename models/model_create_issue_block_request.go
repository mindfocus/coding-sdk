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

// checks if the CreateIssueBlockRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateIssueBlockRequest{}

// CreateIssueBlockRequest struct for CreateIssueBlockRequest
type CreateIssueBlockRequest struct {
	// 前置事项 Code
	BlockIssueCode int64 `json:"BlockIssueCode"`
	// 事项 Code
	IssueCode int64 `json:"IssueCode"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
}

type _CreateIssueBlockRequest CreateIssueBlockRequest

// NewCreateIssueBlockRequest instantiates a new CreateIssueBlockRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIssueBlockRequest(blockIssueCode int64, issueCode int64, projectName string) *CreateIssueBlockRequest {
	this := CreateIssueBlockRequest{}
	this.BlockIssueCode = blockIssueCode
	this.IssueCode = issueCode
	this.ProjectName = projectName
	return &this
}

// NewCreateIssueBlockRequestWithDefaults instantiates a new CreateIssueBlockRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIssueBlockRequestWithDefaults() *CreateIssueBlockRequest {
	this := CreateIssueBlockRequest{}
	return &this
}

// GetBlockIssueCode returns the BlockIssueCode field value
func (o *CreateIssueBlockRequest) GetBlockIssueCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BlockIssueCode
}

// GetBlockIssueCodeOk returns a tuple with the BlockIssueCode field value
// and a boolean to check if the value has been set.
func (o *CreateIssueBlockRequest) GetBlockIssueCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockIssueCode, true
}

// SetBlockIssueCode sets field value
func (o *CreateIssueBlockRequest) SetBlockIssueCode(v int64) {
	o.BlockIssueCode = v
}

// GetIssueCode returns the IssueCode field value
func (o *CreateIssueBlockRequest) GetIssueCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IssueCode
}

// GetIssueCodeOk returns a tuple with the IssueCode field value
// and a boolean to check if the value has been set.
func (o *CreateIssueBlockRequest) GetIssueCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueCode, true
}

// SetIssueCode sets field value
func (o *CreateIssueBlockRequest) SetIssueCode(v int64) {
	o.IssueCode = v
}

// GetProjectName returns the ProjectName field value
func (o *CreateIssueBlockRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *CreateIssueBlockRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *CreateIssueBlockRequest) SetProjectName(v string) {
	o.ProjectName = v
}

func (o CreateIssueBlockRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIssueBlockRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["BlockIssueCode"] = o.BlockIssueCode
	toSerialize["IssueCode"] = o.IssueCode
	toSerialize["ProjectName"] = o.ProjectName
	return toSerialize, nil
}

func (o *CreateIssueBlockRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"BlockIssueCode",
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

	varCreateIssueBlockRequest := _CreateIssueBlockRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateIssueBlockRequest)

	if err != nil {
		return err
	}

	*o = CreateIssueBlockRequest(varCreateIssueBlockRequest)

	return err
}

type NullableCreateIssueBlockRequest struct {
	value *CreateIssueBlockRequest
	isSet bool
}

func (v NullableCreateIssueBlockRequest) Get() *CreateIssueBlockRequest {
	return v.value
}

func (v *NullableCreateIssueBlockRequest) Set(val *CreateIssueBlockRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIssueBlockRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIssueBlockRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIssueBlockRequest(val *CreateIssueBlockRequest) *NullableCreateIssueBlockRequest {
	return &NullableCreateIssueBlockRequest{value: val, isSet: true}
}

func (v NullableCreateIssueBlockRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIssueBlockRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

