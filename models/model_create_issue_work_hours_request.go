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

// checks if the CreateIssueWorkHoursRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateIssueWorkHoursRequest{}

// CreateIssueWorkHoursRequest struct for CreateIssueWorkHoursRequest
type CreateIssueWorkHoursRequest struct {
	// 事项 Code
	IssueCode int64 `json:"IssueCode"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// 剩余工时
	RemainingHour float32 `json:"RemainingHour"`
	// 使用工时
	SpendHour float32 `json:"SpendHour"`
	// 开始时间戳
	StartAt int64 `json:"StartAt"`
	// 工作描述
	WorkingDesc *string `json:"WorkingDesc,omitempty"`
}

type _CreateIssueWorkHoursRequest CreateIssueWorkHoursRequest

// NewCreateIssueWorkHoursRequest instantiates a new CreateIssueWorkHoursRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIssueWorkHoursRequest(issueCode int64, projectName string, remainingHour float32, spendHour float32, startAt int64) *CreateIssueWorkHoursRequest {
	this := CreateIssueWorkHoursRequest{}
	this.IssueCode = issueCode
	this.ProjectName = projectName
	this.RemainingHour = remainingHour
	this.SpendHour = spendHour
	this.StartAt = startAt
	return &this
}

// NewCreateIssueWorkHoursRequestWithDefaults instantiates a new CreateIssueWorkHoursRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIssueWorkHoursRequestWithDefaults() *CreateIssueWorkHoursRequest {
	this := CreateIssueWorkHoursRequest{}
	return &this
}

// GetIssueCode returns the IssueCode field value
func (o *CreateIssueWorkHoursRequest) GetIssueCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IssueCode
}

// GetIssueCodeOk returns a tuple with the IssueCode field value
// and a boolean to check if the value has been set.
func (o *CreateIssueWorkHoursRequest) GetIssueCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueCode, true
}

// SetIssueCode sets field value
func (o *CreateIssueWorkHoursRequest) SetIssueCode(v int64) {
	o.IssueCode = v
}

// GetProjectName returns the ProjectName field value
func (o *CreateIssueWorkHoursRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *CreateIssueWorkHoursRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *CreateIssueWorkHoursRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetRemainingHour returns the RemainingHour field value
func (o *CreateIssueWorkHoursRequest) GetRemainingHour() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.RemainingHour
}

// GetRemainingHourOk returns a tuple with the RemainingHour field value
// and a boolean to check if the value has been set.
func (o *CreateIssueWorkHoursRequest) GetRemainingHourOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingHour, true
}

// SetRemainingHour sets field value
func (o *CreateIssueWorkHoursRequest) SetRemainingHour(v float32) {
	o.RemainingHour = v
}

// GetSpendHour returns the SpendHour field value
func (o *CreateIssueWorkHoursRequest) GetSpendHour() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SpendHour
}

// GetSpendHourOk returns a tuple with the SpendHour field value
// and a boolean to check if the value has been set.
func (o *CreateIssueWorkHoursRequest) GetSpendHourOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpendHour, true
}

// SetSpendHour sets field value
func (o *CreateIssueWorkHoursRequest) SetSpendHour(v float32) {
	o.SpendHour = v
}

// GetStartAt returns the StartAt field value
func (o *CreateIssueWorkHoursRequest) GetStartAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value
// and a boolean to check if the value has been set.
func (o *CreateIssueWorkHoursRequest) GetStartAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartAt, true
}

// SetStartAt sets field value
func (o *CreateIssueWorkHoursRequest) SetStartAt(v int64) {
	o.StartAt = v
}

// GetWorkingDesc returns the WorkingDesc field value if set, zero value otherwise.
func (o *CreateIssueWorkHoursRequest) GetWorkingDesc() string {
	if o == nil || utils.IsNil(o.WorkingDesc) {
		var ret string
		return ret
	}
	return *o.WorkingDesc
}

// GetWorkingDescOk returns a tuple with the WorkingDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueWorkHoursRequest) GetWorkingDescOk() (*string, bool) {
	if o == nil || utils.IsNil(o.WorkingDesc) {
		return nil, false
	}
	return o.WorkingDesc, true
}

// HasWorkingDesc returns a boolean if a field has been set.
func (o *CreateIssueWorkHoursRequest) HasWorkingDesc() bool {
	if o != nil && !utils.IsNil(o.WorkingDesc) {
		return true
	}

	return false
}

// SetWorkingDesc gets a reference to the given string and assigns it to the WorkingDesc field.
func (o *CreateIssueWorkHoursRequest) SetWorkingDesc(v string) {
	o.WorkingDesc = &v
}

func (o CreateIssueWorkHoursRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIssueWorkHoursRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["IssueCode"] = o.IssueCode
	toSerialize["ProjectName"] = o.ProjectName
	toSerialize["RemainingHour"] = o.RemainingHour
	toSerialize["SpendHour"] = o.SpendHour
	toSerialize["StartAt"] = o.StartAt
	if !utils.IsNil(o.WorkingDesc) {
		toSerialize["WorkingDesc"] = o.WorkingDesc
	}
	return toSerialize, nil
}

func (o *CreateIssueWorkHoursRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"IssueCode",
		"ProjectName",
		"RemainingHour",
		"SpendHour",
		"StartAt",
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

	varCreateIssueWorkHoursRequest := _CreateIssueWorkHoursRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateIssueWorkHoursRequest)

	if err != nil {
		return err
	}

	*o = CreateIssueWorkHoursRequest(varCreateIssueWorkHoursRequest)

	return err
}

type NullableCreateIssueWorkHoursRequest struct {
	value *CreateIssueWorkHoursRequest
	isSet bool
}

func (v NullableCreateIssueWorkHoursRequest) Get() *CreateIssueWorkHoursRequest {
	return v.value
}

func (v *NullableCreateIssueWorkHoursRequest) Set(val *CreateIssueWorkHoursRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIssueWorkHoursRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIssueWorkHoursRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIssueWorkHoursRequest(val *CreateIssueWorkHoursRequest) *NullableCreateIssueWorkHoursRequest {
	return &NullableCreateIssueWorkHoursRequest{value: val, isSet: true}
}

func (v NullableCreateIssueWorkHoursRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIssueWorkHoursRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


