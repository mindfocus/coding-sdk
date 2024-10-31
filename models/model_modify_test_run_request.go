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

// checks if the ModifyTestRunRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyTestRunRequest{}

// ModifyTestRunRequest struct for ModifyTestRunRequest
type ModifyTestRunRequest struct {
	// 处理人 ID
	AssignedToId *int32 `json:"AssignedToId,omitempty"`
	// 包含的用例 ID 列表，IncludeAll=false 必填
	Cases []int32 `json:"Cases,omitempty"`
	// 环境标识
	ConfigEnvironmentId *int32 `json:"ConfigEnvironmentId,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty"`
	// 执行类型：1-手动执行 2-自动化流水线执行
	ExecuteType *int32 `json:"ExecuteType,omitempty"`
	// 项目代码库 ID
	GitDepotId *int32 `json:"GitDepotId,omitempty"`
	// 发布版本 ID
	GitReleaseId *int32 `json:"GitReleaseId,omitempty"`
	// 是否包含全部用例
	IncludeAll *bool `json:"IncludeAll,omitempty"`
	// 标题
	Name *string `json:"Name,omitempty"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// 计划 ID
	RunId int32 `json:"RunId"`
	// 分组 ID
	SectionId *int64 `json:"SectionId,omitempty"`
}

type _ModifyTestRunRequest ModifyTestRunRequest

// NewModifyTestRunRequest instantiates a new ModifyTestRunRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyTestRunRequest(projectName string, runId int32) *ModifyTestRunRequest {
	this := ModifyTestRunRequest{}
	this.ProjectName = projectName
	this.RunId = runId
	return &this
}

// NewModifyTestRunRequestWithDefaults instantiates a new ModifyTestRunRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyTestRunRequestWithDefaults() *ModifyTestRunRequest {
	this := ModifyTestRunRequest{}
	return &this
}

// GetAssignedToId returns the AssignedToId field value if set, zero value otherwise.
func (o *ModifyTestRunRequest) GetAssignedToId() int32 {
	if o == nil || utils.IsNil(o.AssignedToId) {
		var ret int32
		return ret
	}
	return *o.AssignedToId
}

// GetAssignedToIdOk returns a tuple with the AssignedToId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestRunRequest) GetAssignedToIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.AssignedToId) {
		return nil, false
	}
	return o.AssignedToId, true
}

// HasAssignedToId returns a boolean if a field has been set.
func (o *ModifyTestRunRequest) HasAssignedToId() bool {
	if o != nil && !utils.IsNil(o.AssignedToId) {
		return true
	}

	return false
}

// SetAssignedToId gets a reference to the given int32 and assigns it to the AssignedToId field.
func (o *ModifyTestRunRequest) SetAssignedToId(v int32) {
	o.AssignedToId = &v
}

// GetCases returns the Cases field value if set, zero value otherwise.
func (o *ModifyTestRunRequest) GetCases() []int32 {
	if o == nil || utils.IsNil(o.Cases) {
		var ret []int32
		return ret
	}
	return o.Cases
}

// GetCasesOk returns a tuple with the Cases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestRunRequest) GetCasesOk() ([]int32, bool) {
	if o == nil || utils.IsNil(o.Cases) {
		return nil, false
	}
	return o.Cases, true
}

// HasCases returns a boolean if a field has been set.
func (o *ModifyTestRunRequest) HasCases() bool {
	if o != nil && !utils.IsNil(o.Cases) {
		return true
	}

	return false
}

// SetCases gets a reference to the given []int32 and assigns it to the Cases field.
func (o *ModifyTestRunRequest) SetCases(v []int32) {
	o.Cases = v
}

// GetConfigEnvironmentId returns the ConfigEnvironmentId field value if set, zero value otherwise.
func (o *ModifyTestRunRequest) GetConfigEnvironmentId() int32 {
	if o == nil || utils.IsNil(o.ConfigEnvironmentId) {
		var ret int32
		return ret
	}
	return *o.ConfigEnvironmentId
}

// GetConfigEnvironmentIdOk returns a tuple with the ConfigEnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestRunRequest) GetConfigEnvironmentIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ConfigEnvironmentId) {
		return nil, false
	}
	return o.ConfigEnvironmentId, true
}

// HasConfigEnvironmentId returns a boolean if a field has been set.
func (o *ModifyTestRunRequest) HasConfigEnvironmentId() bool {
	if o != nil && !utils.IsNil(o.ConfigEnvironmentId) {
		return true
	}

	return false
}

// SetConfigEnvironmentId gets a reference to the given int32 and assigns it to the ConfigEnvironmentId field.
func (o *ModifyTestRunRequest) SetConfigEnvironmentId(v int32) {
	o.ConfigEnvironmentId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModifyTestRunRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestRunRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModifyTestRunRequest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModifyTestRunRequest) SetDescription(v string) {
	o.Description = &v
}

// GetExecuteType returns the ExecuteType field value if set, zero value otherwise.
func (o *ModifyTestRunRequest) GetExecuteType() int32 {
	if o == nil || utils.IsNil(o.ExecuteType) {
		var ret int32
		return ret
	}
	return *o.ExecuteType
}

// GetExecuteTypeOk returns a tuple with the ExecuteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestRunRequest) GetExecuteTypeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ExecuteType) {
		return nil, false
	}
	return o.ExecuteType, true
}

// HasExecuteType returns a boolean if a field has been set.
func (o *ModifyTestRunRequest) HasExecuteType() bool {
	if o != nil && !utils.IsNil(o.ExecuteType) {
		return true
	}

	return false
}

// SetExecuteType gets a reference to the given int32 and assigns it to the ExecuteType field.
func (o *ModifyTestRunRequest) SetExecuteType(v int32) {
	o.ExecuteType = &v
}

// GetGitDepotId returns the GitDepotId field value if set, zero value otherwise.
func (o *ModifyTestRunRequest) GetGitDepotId() int32 {
	if o == nil || utils.IsNil(o.GitDepotId) {
		var ret int32
		return ret
	}
	return *o.GitDepotId
}

// GetGitDepotIdOk returns a tuple with the GitDepotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestRunRequest) GetGitDepotIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.GitDepotId) {
		return nil, false
	}
	return o.GitDepotId, true
}

// HasGitDepotId returns a boolean if a field has been set.
func (o *ModifyTestRunRequest) HasGitDepotId() bool {
	if o != nil && !utils.IsNil(o.GitDepotId) {
		return true
	}

	return false
}

// SetGitDepotId gets a reference to the given int32 and assigns it to the GitDepotId field.
func (o *ModifyTestRunRequest) SetGitDepotId(v int32) {
	o.GitDepotId = &v
}

// GetGitReleaseId returns the GitReleaseId field value if set, zero value otherwise.
func (o *ModifyTestRunRequest) GetGitReleaseId() int32 {
	if o == nil || utils.IsNil(o.GitReleaseId) {
		var ret int32
		return ret
	}
	return *o.GitReleaseId
}

// GetGitReleaseIdOk returns a tuple with the GitReleaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestRunRequest) GetGitReleaseIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.GitReleaseId) {
		return nil, false
	}
	return o.GitReleaseId, true
}

// HasGitReleaseId returns a boolean if a field has been set.
func (o *ModifyTestRunRequest) HasGitReleaseId() bool {
	if o != nil && !utils.IsNil(o.GitReleaseId) {
		return true
	}

	return false
}

// SetGitReleaseId gets a reference to the given int32 and assigns it to the GitReleaseId field.
func (o *ModifyTestRunRequest) SetGitReleaseId(v int32) {
	o.GitReleaseId = &v
}

// GetIncludeAll returns the IncludeAll field value if set, zero value otherwise.
func (o *ModifyTestRunRequest) GetIncludeAll() bool {
	if o == nil || utils.IsNil(o.IncludeAll) {
		var ret bool
		return ret
	}
	return *o.IncludeAll
}

// GetIncludeAllOk returns a tuple with the IncludeAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestRunRequest) GetIncludeAllOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IncludeAll) {
		return nil, false
	}
	return o.IncludeAll, true
}

// HasIncludeAll returns a boolean if a field has been set.
func (o *ModifyTestRunRequest) HasIncludeAll() bool {
	if o != nil && !utils.IsNil(o.IncludeAll) {
		return true
	}

	return false
}

// SetIncludeAll gets a reference to the given bool and assigns it to the IncludeAll field.
func (o *ModifyTestRunRequest) SetIncludeAll(v bool) {
	o.IncludeAll = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModifyTestRunRequest) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestRunRequest) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModifyTestRunRequest) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModifyTestRunRequest) SetName(v string) {
	o.Name = &v
}

// GetProjectName returns the ProjectName field value
func (o *ModifyTestRunRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *ModifyTestRunRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *ModifyTestRunRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetRunId returns the RunId field value
func (o *ModifyTestRunRequest) GetRunId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *ModifyTestRunRequest) GetRunIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *ModifyTestRunRequest) SetRunId(v int32) {
	o.RunId = v
}

// GetSectionId returns the SectionId field value if set, zero value otherwise.
func (o *ModifyTestRunRequest) GetSectionId() int64 {
	if o == nil || utils.IsNil(o.SectionId) {
		var ret int64
		return ret
	}
	return *o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestRunRequest) GetSectionIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.SectionId) {
		return nil, false
	}
	return o.SectionId, true
}

// HasSectionId returns a boolean if a field has been set.
func (o *ModifyTestRunRequest) HasSectionId() bool {
	if o != nil && !utils.IsNil(o.SectionId) {
		return true
	}

	return false
}

// SetSectionId gets a reference to the given int64 and assigns it to the SectionId field.
func (o *ModifyTestRunRequest) SetSectionId(v int64) {
	o.SectionId = &v
}

func (o ModifyTestRunRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyTestRunRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AssignedToId) {
		toSerialize["AssignedToId"] = o.AssignedToId
	}
	if !utils.IsNil(o.Cases) {
		toSerialize["Cases"] = o.Cases
	}
	if !utils.IsNil(o.ConfigEnvironmentId) {
		toSerialize["ConfigEnvironmentId"] = o.ConfigEnvironmentId
	}
	if !utils.IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !utils.IsNil(o.ExecuteType) {
		toSerialize["ExecuteType"] = o.ExecuteType
	}
	if !utils.IsNil(o.GitDepotId) {
		toSerialize["GitDepotId"] = o.GitDepotId
	}
	if !utils.IsNil(o.GitReleaseId) {
		toSerialize["GitReleaseId"] = o.GitReleaseId
	}
	if !utils.IsNil(o.IncludeAll) {
		toSerialize["IncludeAll"] = o.IncludeAll
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	toSerialize["ProjectName"] = o.ProjectName
	toSerialize["RunId"] = o.RunId
	if !utils.IsNil(o.SectionId) {
		toSerialize["SectionId"] = o.SectionId
	}
	return toSerialize, nil
}

func (o *ModifyTestRunRequest) UnmarshalJSON(data []byte) (err error) {
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

	varModifyTestRunRequest := _ModifyTestRunRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyTestRunRequest)

	if err != nil {
		return err
	}

	*o = ModifyTestRunRequest(varModifyTestRunRequest)

	return err
}

type NullableModifyTestRunRequest struct {
	value *ModifyTestRunRequest
	isSet bool
}

func (v NullableModifyTestRunRequest) Get() *ModifyTestRunRequest {
	return v.value
}

func (v *NullableModifyTestRunRequest) Set(val *ModifyTestRunRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyTestRunRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyTestRunRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyTestRunRequest(val *ModifyTestRunRequest) *NullableModifyTestRunRequest {
	return &NullableModifyTestRunRequest{value: val, isSet: true}
}

func (v NullableModifyTestRunRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyTestRunRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

