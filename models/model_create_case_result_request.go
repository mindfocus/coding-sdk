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

// checks if the CreateCaseResultRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateCaseResultRequest{}

// CreateCaseResultRequest struct for CreateCaseResultRequest
type CreateCaseResultRequest struct {
	// 测试任务 ID
	CaseId int32 `json:"CaseId"`
	// 每一步的测试结果（步骤用例时需要本参数）
	CustomStepStatus []string `json:"CustomStepStatus,omitempty"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// 测试计划 ID
	RunId int32 `json:"RunId"`
	// 该任务的测试结果，可选值：UNTESTED:未测试,PASSED:通过,BLOCKED:阻塞,RETEST:重测,FAILED:失败
	Status string `json:"Status"`
}

type _CreateCaseResultRequest CreateCaseResultRequest

// NewCreateCaseResultRequest instantiates a new CreateCaseResultRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCaseResultRequest(caseId int32, projectName string, runId int32, status string) *CreateCaseResultRequest {
	this := CreateCaseResultRequest{}
	this.CaseId = caseId
	this.ProjectName = projectName
	this.RunId = runId
	this.Status = status
	return &this
}

// NewCreateCaseResultRequestWithDefaults instantiates a new CreateCaseResultRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCaseResultRequestWithDefaults() *CreateCaseResultRequest {
	this := CreateCaseResultRequest{}
	return &this
}

// GetCaseId returns the CaseId field value
func (o *CreateCaseResultRequest) GetCaseId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value
// and a boolean to check if the value has been set.
func (o *CreateCaseResultRequest) GetCaseIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaseId, true
}

// SetCaseId sets field value
func (o *CreateCaseResultRequest) SetCaseId(v int32) {
	o.CaseId = v
}

// GetCustomStepStatus returns the CustomStepStatus field value if set, zero value otherwise.
func (o *CreateCaseResultRequest) GetCustomStepStatus() []string {
	if o == nil || utils.IsNil(o.CustomStepStatus) {
		var ret []string
		return ret
	}
	return o.CustomStepStatus
}

// GetCustomStepStatusOk returns a tuple with the CustomStepStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCaseResultRequest) GetCustomStepStatusOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.CustomStepStatus) {
		return nil, false
	}
	return o.CustomStepStatus, true
}

// HasCustomStepStatus returns a boolean if a field has been set.
func (o *CreateCaseResultRequest) HasCustomStepStatus() bool {
	if o != nil && !utils.IsNil(o.CustomStepStatus) {
		return true
	}

	return false
}

// SetCustomStepStatus gets a reference to the given []string and assigns it to the CustomStepStatus field.
func (o *CreateCaseResultRequest) SetCustomStepStatus(v []string) {
	o.CustomStepStatus = v
}

// GetProjectName returns the ProjectName field value
func (o *CreateCaseResultRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *CreateCaseResultRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *CreateCaseResultRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetRunId returns the RunId field value
func (o *CreateCaseResultRequest) GetRunId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *CreateCaseResultRequest) GetRunIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value
func (o *CreateCaseResultRequest) SetRunId(v int32) {
	o.RunId = v
}

// GetStatus returns the Status field value
func (o *CreateCaseResultRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateCaseResultRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateCaseResultRequest) SetStatus(v string) {
	o.Status = v
}

func (o CreateCaseResultRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCaseResultRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CaseId"] = o.CaseId
	if !utils.IsNil(o.CustomStepStatus) {
		toSerialize["CustomStepStatus"] = o.CustomStepStatus
	}
	toSerialize["ProjectName"] = o.ProjectName
	toSerialize["RunId"] = o.RunId
	toSerialize["Status"] = o.Status
	return toSerialize, nil
}

func (o *CreateCaseResultRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"CaseId",
		"ProjectName",
		"RunId",
		"Status",
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

	varCreateCaseResultRequest := _CreateCaseResultRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateCaseResultRequest)

	if err != nil {
		return err
	}

	*o = CreateCaseResultRequest(varCreateCaseResultRequest)

	return err
}

type NullableCreateCaseResultRequest struct {
	value *CreateCaseResultRequest
	isSet bool
}

func (v NullableCreateCaseResultRequest) Get() *CreateCaseResultRequest {
	return v.value
}

func (v *NullableCreateCaseResultRequest) Set(val *CreateCaseResultRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCaseResultRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCaseResultRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCaseResultRequest(val *CreateCaseResultRequest) *NullableCreateCaseResultRequest {
	return &NullableCreateCaseResultRequest{value: val, isSet: true}
}

func (v NullableCreateCaseResultRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCaseResultRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


