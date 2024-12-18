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

// checks if the DescribeTestCaseListRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeTestCaseListRequest{}

// DescribeTestCaseListRequest struct for DescribeTestCaseListRequest
type DescribeTestCaseListRequest struct {
	// 关键字搜索
	Keyword *string `json:"Keyword,omitempty"`
	// 优先级，默认2（中），可选值：0（紧急）,1（高）,2（中）,3（低）
	Priority *int32 `json:"Priority,omitempty"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// 分组 ID
	SectionId *int32 `json:"SectionId,omitempty"`
	// 用例类型，可选值：STEPS(步骤用例)，TEXT(文本用例)
	TemplateType *string `json:"TemplateType,omitempty"`
}

type _DescribeTestCaseListRequest DescribeTestCaseListRequest

// NewDescribeTestCaseListRequest instantiates a new DescribeTestCaseListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeTestCaseListRequest(projectName string) *DescribeTestCaseListRequest {
	this := DescribeTestCaseListRequest{}
	this.ProjectName = projectName
	return &this
}

// NewDescribeTestCaseListRequestWithDefaults instantiates a new DescribeTestCaseListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeTestCaseListRequestWithDefaults() *DescribeTestCaseListRequest {
	this := DescribeTestCaseListRequest{}
	return &this
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *DescribeTestCaseListRequest) GetKeyword() string {
	if o == nil || utils.IsNil(o.Keyword) {
		var ret string
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTestCaseListRequest) GetKeywordOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *DescribeTestCaseListRequest) HasKeyword() bool {
	if o != nil && !utils.IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *DescribeTestCaseListRequest) SetKeyword(v string) {
	o.Keyword = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *DescribeTestCaseListRequest) GetPriority() int32 {
	if o == nil || utils.IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTestCaseListRequest) GetPriorityOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *DescribeTestCaseListRequest) HasPriority() bool {
	if o != nil && !utils.IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *DescribeTestCaseListRequest) SetPriority(v int32) {
	o.Priority = &v
}

// GetProjectName returns the ProjectName field value
func (o *DescribeTestCaseListRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *DescribeTestCaseListRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *DescribeTestCaseListRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetSectionId returns the SectionId field value if set, zero value otherwise.
func (o *DescribeTestCaseListRequest) GetSectionId() int32 {
	if o == nil || utils.IsNil(o.SectionId) {
		var ret int32
		return ret
	}
	return *o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTestCaseListRequest) GetSectionIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.SectionId) {
		return nil, false
	}
	return o.SectionId, true
}

// HasSectionId returns a boolean if a field has been set.
func (o *DescribeTestCaseListRequest) HasSectionId() bool {
	if o != nil && !utils.IsNil(o.SectionId) {
		return true
	}

	return false
}

// SetSectionId gets a reference to the given int32 and assigns it to the SectionId field.
func (o *DescribeTestCaseListRequest) SetSectionId(v int32) {
	o.SectionId = &v
}

// GetTemplateType returns the TemplateType field value if set, zero value otherwise.
func (o *DescribeTestCaseListRequest) GetTemplateType() string {
	if o == nil || utils.IsNil(o.TemplateType) {
		var ret string
		return ret
	}
	return *o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTestCaseListRequest) GetTemplateTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TemplateType) {
		return nil, false
	}
	return o.TemplateType, true
}

// HasTemplateType returns a boolean if a field has been set.
func (o *DescribeTestCaseListRequest) HasTemplateType() bool {
	if o != nil && !utils.IsNil(o.TemplateType) {
		return true
	}

	return false
}

// SetTemplateType gets a reference to the given string and assigns it to the TemplateType field.
func (o *DescribeTestCaseListRequest) SetTemplateType(v string) {
	o.TemplateType = &v
}

func (o DescribeTestCaseListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeTestCaseListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Keyword) {
		toSerialize["Keyword"] = o.Keyword
	}
	if !utils.IsNil(o.Priority) {
		toSerialize["Priority"] = o.Priority
	}
	toSerialize["ProjectName"] = o.ProjectName
	if !utils.IsNil(o.SectionId) {
		toSerialize["SectionId"] = o.SectionId
	}
	if !utils.IsNil(o.TemplateType) {
		toSerialize["TemplateType"] = o.TemplateType
	}
	return toSerialize, nil
}

func (o *DescribeTestCaseListRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varDescribeTestCaseListRequest := _DescribeTestCaseListRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeTestCaseListRequest)

	if err != nil {
		return err
	}

	*o = DescribeTestCaseListRequest(varDescribeTestCaseListRequest)

	return err
}

type NullableDescribeTestCaseListRequest struct {
	value *DescribeTestCaseListRequest
	isSet bool
}

func (v NullableDescribeTestCaseListRequest) Get() *DescribeTestCaseListRequest {
	return v.value
}

func (v *NullableDescribeTestCaseListRequest) Set(val *DescribeTestCaseListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeTestCaseListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeTestCaseListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeTestCaseListRequest(val *DescribeTestCaseListRequest) *NullableDescribeTestCaseListRequest {
	return &NullableDescribeTestCaseListRequest{value: val, isSet: true}
}

func (v NullableDescribeTestCaseListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeTestCaseListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


