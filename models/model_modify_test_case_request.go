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

// checks if the ModifyTestCaseRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyTestCaseRequest{}

// ModifyTestCaseRequest struct for ModifyTestCaseRequest
type ModifyTestCaseRequest struct {
	// 附件 ID 数组：来自“生成附件预上传信息”接口
	AttachmentIds []int32 `json:"AttachmentIds,omitempty"`
	// 用例 ID
	CaseId int32 `json:"CaseId"`
	// 自定义步骤（步骤用例必填）
	CustomSteps []CustomStep `json:"CustomSteps,omitempty"`
	// 预期结果（适用于文本用例）
	Expected *string `json:"Expected,omitempty"`
	// 前置步骤
	Preconds *string `json:"Preconds,omitempty"`
	// 优先级，默认2（中），可选值：0（紧急）,1（高）,2（中）,3（低）
	Priority *int32 `json:"Priority,omitempty"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// 分组ID
	SectionId int32 `json:"SectionId"`
	// 文本描述（适用于文本用例）
	Steps *string `json:"Steps,omitempty"`
	// 用例类型，可选值：STEPS(步骤用例)，TEXT(文本用例)
	TemplateType string `json:"TemplateType"`
	// 用例标题
	Title string `json:"Title"`
}

type _ModifyTestCaseRequest ModifyTestCaseRequest

// NewModifyTestCaseRequest instantiates a new ModifyTestCaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyTestCaseRequest(caseId int32, projectName string, sectionId int32, templateType string, title string) *ModifyTestCaseRequest {
	this := ModifyTestCaseRequest{}
	this.CaseId = caseId
	this.ProjectName = projectName
	this.SectionId = sectionId
	this.TemplateType = templateType
	this.Title = title
	return &this
}

// NewModifyTestCaseRequestWithDefaults instantiates a new ModifyTestCaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyTestCaseRequestWithDefaults() *ModifyTestCaseRequest {
	this := ModifyTestCaseRequest{}
	return &this
}

// GetAttachmentIds returns the AttachmentIds field value if set, zero value otherwise.
func (o *ModifyTestCaseRequest) GetAttachmentIds() []int32 {
	if o == nil || utils.IsNil(o.AttachmentIds) {
		var ret []int32
		return ret
	}
	return o.AttachmentIds
}

// GetAttachmentIdsOk returns a tuple with the AttachmentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestCaseRequest) GetAttachmentIdsOk() ([]int32, bool) {
	if o == nil || utils.IsNil(o.AttachmentIds) {
		return nil, false
	}
	return o.AttachmentIds, true
}

// HasAttachmentIds returns a boolean if a field has been set.
func (o *ModifyTestCaseRequest) HasAttachmentIds() bool {
	if o != nil && !utils.IsNil(o.AttachmentIds) {
		return true
	}

	return false
}

// SetAttachmentIds gets a reference to the given []int32 and assigns it to the AttachmentIds field.
func (o *ModifyTestCaseRequest) SetAttachmentIds(v []int32) {
	o.AttachmentIds = v
}

// GetCaseId returns the CaseId field value
func (o *ModifyTestCaseRequest) GetCaseId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CaseId
}

// GetCaseIdOk returns a tuple with the CaseId field value
// and a boolean to check if the value has been set.
func (o *ModifyTestCaseRequest) GetCaseIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaseId, true
}

// SetCaseId sets field value
func (o *ModifyTestCaseRequest) SetCaseId(v int32) {
	o.CaseId = v
}

// GetCustomSteps returns the CustomSteps field value if set, zero value otherwise.
func (o *ModifyTestCaseRequest) GetCustomSteps() []CustomStep {
	if o == nil || utils.IsNil(o.CustomSteps) {
		var ret []CustomStep
		return ret
	}
	return o.CustomSteps
}

// GetCustomStepsOk returns a tuple with the CustomSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestCaseRequest) GetCustomStepsOk() ([]CustomStep, bool) {
	if o == nil || utils.IsNil(o.CustomSteps) {
		return nil, false
	}
	return o.CustomSteps, true
}

// HasCustomSteps returns a boolean if a field has been set.
func (o *ModifyTestCaseRequest) HasCustomSteps() bool {
	if o != nil && !utils.IsNil(o.CustomSteps) {
		return true
	}

	return false
}

// SetCustomSteps gets a reference to the given []CustomStep and assigns it to the CustomSteps field.
func (o *ModifyTestCaseRequest) SetCustomSteps(v []CustomStep) {
	o.CustomSteps = v
}

// GetExpected returns the Expected field value if set, zero value otherwise.
func (o *ModifyTestCaseRequest) GetExpected() string {
	if o == nil || utils.IsNil(o.Expected) {
		var ret string
		return ret
	}
	return *o.Expected
}

// GetExpectedOk returns a tuple with the Expected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestCaseRequest) GetExpectedOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Expected) {
		return nil, false
	}
	return o.Expected, true
}

// HasExpected returns a boolean if a field has been set.
func (o *ModifyTestCaseRequest) HasExpected() bool {
	if o != nil && !utils.IsNil(o.Expected) {
		return true
	}

	return false
}

// SetExpected gets a reference to the given string and assigns it to the Expected field.
func (o *ModifyTestCaseRequest) SetExpected(v string) {
	o.Expected = &v
}

// GetPreconds returns the Preconds field value if set, zero value otherwise.
func (o *ModifyTestCaseRequest) GetPreconds() string {
	if o == nil || utils.IsNil(o.Preconds) {
		var ret string
		return ret
	}
	return *o.Preconds
}

// GetPrecondsOk returns a tuple with the Preconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestCaseRequest) GetPrecondsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Preconds) {
		return nil, false
	}
	return o.Preconds, true
}

// HasPreconds returns a boolean if a field has been set.
func (o *ModifyTestCaseRequest) HasPreconds() bool {
	if o != nil && !utils.IsNil(o.Preconds) {
		return true
	}

	return false
}

// SetPreconds gets a reference to the given string and assigns it to the Preconds field.
func (o *ModifyTestCaseRequest) SetPreconds(v string) {
	o.Preconds = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ModifyTestCaseRequest) GetPriority() int32 {
	if o == nil || utils.IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestCaseRequest) GetPriorityOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ModifyTestCaseRequest) HasPriority() bool {
	if o != nil && !utils.IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *ModifyTestCaseRequest) SetPriority(v int32) {
	o.Priority = &v
}

// GetProjectName returns the ProjectName field value
func (o *ModifyTestCaseRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *ModifyTestCaseRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *ModifyTestCaseRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetSectionId returns the SectionId field value
func (o *ModifyTestCaseRequest) GetSectionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value
// and a boolean to check if the value has been set.
func (o *ModifyTestCaseRequest) GetSectionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SectionId, true
}

// SetSectionId sets field value
func (o *ModifyTestCaseRequest) SetSectionId(v int32) {
	o.SectionId = v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *ModifyTestCaseRequest) GetSteps() string {
	if o == nil || utils.IsNil(o.Steps) {
		var ret string
		return ret
	}
	return *o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyTestCaseRequest) GetStepsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Steps) {
		return nil, false
	}
	return o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *ModifyTestCaseRequest) HasSteps() bool {
	if o != nil && !utils.IsNil(o.Steps) {
		return true
	}

	return false
}

// SetSteps gets a reference to the given string and assigns it to the Steps field.
func (o *ModifyTestCaseRequest) SetSteps(v string) {
	o.Steps = &v
}

// GetTemplateType returns the TemplateType field value
func (o *ModifyTestCaseRequest) GetTemplateType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value
// and a boolean to check if the value has been set.
func (o *ModifyTestCaseRequest) GetTemplateTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateType, true
}

// SetTemplateType sets field value
func (o *ModifyTestCaseRequest) SetTemplateType(v string) {
	o.TemplateType = v
}

// GetTitle returns the Title field value
func (o *ModifyTestCaseRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ModifyTestCaseRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ModifyTestCaseRequest) SetTitle(v string) {
	o.Title = v
}

func (o ModifyTestCaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyTestCaseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AttachmentIds) {
		toSerialize["AttachmentIds"] = o.AttachmentIds
	}
	toSerialize["CaseId"] = o.CaseId
	if !utils.IsNil(o.CustomSteps) {
		toSerialize["CustomSteps"] = o.CustomSteps
	}
	if !utils.IsNil(o.Expected) {
		toSerialize["Expected"] = o.Expected
	}
	if !utils.IsNil(o.Preconds) {
		toSerialize["Preconds"] = o.Preconds
	}
	if !utils.IsNil(o.Priority) {
		toSerialize["Priority"] = o.Priority
	}
	toSerialize["ProjectName"] = o.ProjectName
	toSerialize["SectionId"] = o.SectionId
	if !utils.IsNil(o.Steps) {
		toSerialize["Steps"] = o.Steps
	}
	toSerialize["TemplateType"] = o.TemplateType
	toSerialize["Title"] = o.Title
	return toSerialize, nil
}

func (o *ModifyTestCaseRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"CaseId",
		"ProjectName",
		"SectionId",
		"TemplateType",
		"Title",
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

	varModifyTestCaseRequest := _ModifyTestCaseRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyTestCaseRequest)

	if err != nil {
		return err
	}

	*o = ModifyTestCaseRequest(varModifyTestCaseRequest)

	return err
}

type NullableModifyTestCaseRequest struct {
	value *ModifyTestCaseRequest
	isSet bool
}

func (v NullableModifyTestCaseRequest) Get() *ModifyTestCaseRequest {
	return v.value
}

func (v *NullableModifyTestCaseRequest) Set(val *ModifyTestCaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyTestCaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyTestCaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyTestCaseRequest(val *ModifyTestCaseRequest) *NullableModifyTestCaseRequest {
	return &NullableModifyTestCaseRequest{value: val, isSet: true}
}

func (v NullableModifyTestCaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyTestCaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

