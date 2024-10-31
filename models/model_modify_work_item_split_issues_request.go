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

// checks if the ModifyWorkItemSplitIssuesRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyWorkItemSplitIssuesRequest{}

// ModifyWorkItemSplitIssuesRequest struct for ModifyWorkItemSplitIssuesRequest
type ModifyWorkItemSplitIssuesRequest struct {
	// 目标项目中的事项ID
	IssueCode string `json:"IssueCode"`
	// 项目集名称
	ProgramName string `json:"ProgramName"`
	// 目标项目名称 
	ProjectName string `json:"ProjectName"`
	// true 表示分解， false 表示取消分解
	Split string `json:"Split"`
	// 页面上工作项ID
	WorkItemCode int64 `json:"WorkItemCode"`
}

type _ModifyWorkItemSplitIssuesRequest ModifyWorkItemSplitIssuesRequest

// NewModifyWorkItemSplitIssuesRequest instantiates a new ModifyWorkItemSplitIssuesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyWorkItemSplitIssuesRequest(issueCode string, programName string, projectName string, split string, workItemCode int64) *ModifyWorkItemSplitIssuesRequest {
	this := ModifyWorkItemSplitIssuesRequest{}
	this.IssueCode = issueCode
	this.ProgramName = programName
	this.ProjectName = projectName
	this.Split = split
	this.WorkItemCode = workItemCode
	return &this
}

// NewModifyWorkItemSplitIssuesRequestWithDefaults instantiates a new ModifyWorkItemSplitIssuesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyWorkItemSplitIssuesRequestWithDefaults() *ModifyWorkItemSplitIssuesRequest {
	this := ModifyWorkItemSplitIssuesRequest{}
	return &this
}

// GetIssueCode returns the IssueCode field value
func (o *ModifyWorkItemSplitIssuesRequest) GetIssueCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssueCode
}

// GetIssueCodeOk returns a tuple with the IssueCode field value
// and a boolean to check if the value has been set.
func (o *ModifyWorkItemSplitIssuesRequest) GetIssueCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueCode, true
}

// SetIssueCode sets field value
func (o *ModifyWorkItemSplitIssuesRequest) SetIssueCode(v string) {
	o.IssueCode = v
}

// GetProgramName returns the ProgramName field value
func (o *ModifyWorkItemSplitIssuesRequest) GetProgramName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProgramName
}

// GetProgramNameOk returns a tuple with the ProgramName field value
// and a boolean to check if the value has been set.
func (o *ModifyWorkItemSplitIssuesRequest) GetProgramNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProgramName, true
}

// SetProgramName sets field value
func (o *ModifyWorkItemSplitIssuesRequest) SetProgramName(v string) {
	o.ProgramName = v
}

// GetProjectName returns the ProjectName field value
func (o *ModifyWorkItemSplitIssuesRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *ModifyWorkItemSplitIssuesRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *ModifyWorkItemSplitIssuesRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetSplit returns the Split field value
func (o *ModifyWorkItemSplitIssuesRequest) GetSplit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Split
}

// GetSplitOk returns a tuple with the Split field value
// and a boolean to check if the value has been set.
func (o *ModifyWorkItemSplitIssuesRequest) GetSplitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Split, true
}

// SetSplit sets field value
func (o *ModifyWorkItemSplitIssuesRequest) SetSplit(v string) {
	o.Split = v
}

// GetWorkItemCode returns the WorkItemCode field value
func (o *ModifyWorkItemSplitIssuesRequest) GetWorkItemCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.WorkItemCode
}

// GetWorkItemCodeOk returns a tuple with the WorkItemCode field value
// and a boolean to check if the value has been set.
func (o *ModifyWorkItemSplitIssuesRequest) GetWorkItemCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkItemCode, true
}

// SetWorkItemCode sets field value
func (o *ModifyWorkItemSplitIssuesRequest) SetWorkItemCode(v int64) {
	o.WorkItemCode = v
}

func (o ModifyWorkItemSplitIssuesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyWorkItemSplitIssuesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["IssueCode"] = o.IssueCode
	toSerialize["ProgramName"] = o.ProgramName
	toSerialize["ProjectName"] = o.ProjectName
	toSerialize["Split"] = o.Split
	toSerialize["WorkItemCode"] = o.WorkItemCode
	return toSerialize, nil
}

func (o *ModifyWorkItemSplitIssuesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"IssueCode",
		"ProgramName",
		"ProjectName",
		"Split",
		"WorkItemCode",
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

	varModifyWorkItemSplitIssuesRequest := _ModifyWorkItemSplitIssuesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyWorkItemSplitIssuesRequest)

	if err != nil {
		return err
	}

	*o = ModifyWorkItemSplitIssuesRequest(varModifyWorkItemSplitIssuesRequest)

	return err
}

type NullableModifyWorkItemSplitIssuesRequest struct {
	value *ModifyWorkItemSplitIssuesRequest
	isSet bool
}

func (v NullableModifyWorkItemSplitIssuesRequest) Get() *ModifyWorkItemSplitIssuesRequest {
	return v.value
}

func (v *NullableModifyWorkItemSplitIssuesRequest) Set(val *ModifyWorkItemSplitIssuesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyWorkItemSplitIssuesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyWorkItemSplitIssuesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyWorkItemSplitIssuesRequest(val *ModifyWorkItemSplitIssuesRequest) *NullableModifyWorkItemSplitIssuesRequest {
	return &NullableModifyWorkItemSplitIssuesRequest{value: val, isSet: true}
}

func (v NullableModifyWorkItemSplitIssuesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyWorkItemSplitIssuesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


