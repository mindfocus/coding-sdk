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

// checks if the CreateTestResultRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateTestResultRequest{}

// CreateTestResultRequest struct for CreateTestResultRequest
type CreateTestResultRequest struct {
	// 每一步的测试结果（步骤用例时需要本参数）
	CustomStepStatus []string `json:"CustomStepStatus,omitempty"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// 该任务的测试结果，可选值：UNTESTED:未测试,PASSED:通过,BLOCKED:阻塞,RETEST:重测,FAILED:失败
	Status string `json:"Status"`
	// 测试任务 ID
	TestId int32 `json:"TestId"`
}

type _CreateTestResultRequest CreateTestResultRequest

// NewCreateTestResultRequest instantiates a new CreateTestResultRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTestResultRequest(projectName string, status string, testId int32) *CreateTestResultRequest {
	this := CreateTestResultRequest{}
	this.ProjectName = projectName
	this.Status = status
	this.TestId = testId
	return &this
}

// NewCreateTestResultRequestWithDefaults instantiates a new CreateTestResultRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTestResultRequestWithDefaults() *CreateTestResultRequest {
	this := CreateTestResultRequest{}
	return &this
}

// GetCustomStepStatus returns the CustomStepStatus field value if set, zero value otherwise.
func (o *CreateTestResultRequest) GetCustomStepStatus() []string {
	if o == nil || utils.IsNil(o.CustomStepStatus) {
		var ret []string
		return ret
	}
	return o.CustomStepStatus
}

// GetCustomStepStatusOk returns a tuple with the CustomStepStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTestResultRequest) GetCustomStepStatusOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.CustomStepStatus) {
		return nil, false
	}
	return o.CustomStepStatus, true
}

// HasCustomStepStatus returns a boolean if a field has been set.
func (o *CreateTestResultRequest) HasCustomStepStatus() bool {
	if o != nil && !utils.IsNil(o.CustomStepStatus) {
		return true
	}

	return false
}

// SetCustomStepStatus gets a reference to the given []string and assigns it to the CustomStepStatus field.
func (o *CreateTestResultRequest) SetCustomStepStatus(v []string) {
	o.CustomStepStatus = v
}

// GetProjectName returns the ProjectName field value
func (o *CreateTestResultRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *CreateTestResultRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *CreateTestResultRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetStatus returns the Status field value
func (o *CreateTestResultRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateTestResultRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateTestResultRequest) SetStatus(v string) {
	o.Status = v
}

// GetTestId returns the TestId field value
func (o *CreateTestResultRequest) GetTestId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value
// and a boolean to check if the value has been set.
func (o *CreateTestResultRequest) GetTestIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestId, true
}

// SetTestId sets field value
func (o *CreateTestResultRequest) SetTestId(v int32) {
	o.TestId = v
}

func (o CreateTestResultRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTestResultRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CustomStepStatus) {
		toSerialize["CustomStepStatus"] = o.CustomStepStatus
	}
	toSerialize["ProjectName"] = o.ProjectName
	toSerialize["Status"] = o.Status
	toSerialize["TestId"] = o.TestId
	return toSerialize, nil
}

func (o *CreateTestResultRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ProjectName",
		"Status",
		"TestId",
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

	varCreateTestResultRequest := _CreateTestResultRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTestResultRequest)

	if err != nil {
		return err
	}

	*o = CreateTestResultRequest(varCreateTestResultRequest)

	return err
}

type NullableCreateTestResultRequest struct {
	value *CreateTestResultRequest
	isSet bool
}

func (v NullableCreateTestResultRequest) Get() *CreateTestResultRequest {
	return v.value
}

func (v *NullableCreateTestResultRequest) Set(val *CreateTestResultRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTestResultRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTestResultRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTestResultRequest(val *CreateTestResultRequest) *NullableCreateTestResultRequest {
	return &NullableCreateTestResultRequest{value: val, isSet: true}
}

func (v NullableCreateTestResultRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTestResultRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


