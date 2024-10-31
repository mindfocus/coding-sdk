/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ProjectIssueField type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ProjectIssueField{}

// ProjectIssueField 项目的事项属性
type ProjectIssueField struct {
	// 创建时间戳
	CreatedAt *int64 `json:"CreatedAt,omitempty"`
	IssueField *IssueField `json:"IssueField,omitempty"`
	// 关联属性 ID
	IssueFieldId *int64 `json:"IssueFieldId,omitempty"`
	// 事项类型
	IssueType *string `json:"IssueType,omitempty"`
	// 是否有默认值
	NeedDefault *bool `json:"NeedDefault,omitempty"`
	// 是否必填
	Required *bool `json:"Required,omitempty"`
	// 修改时间戳
	UpdatedAt *int64 `json:"UpdatedAt,omitempty"`
	// 默认值， JSON 字符串。例如：{\"type\":\"VARIABLE\",\"value\":\"CREATOR\"}  type：默认值类型，VARIABLE（变量）、CONSTANT（常量）；  value：默认值，根据 IssueField.ComponentType，可能为不同类型的值（数值、字符串、数组）；常量值 CREATOR 表示创建者
	ValueString utils.NullableString `json:"ValueString,omitempty"`
}

// NewProjectIssueField instantiates a new ProjectIssueField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectIssueField() *ProjectIssueField {
	this := ProjectIssueField{}
	var issueType string = ""
	this.IssueType = &issueType
	var needDefault bool = false
	this.NeedDefault = &needDefault
	var required bool = false
	this.Required = &required
	var valueString string = ""
	this.ValueString = *utils.NewNullableString(&valueString)
	return &this
}

// NewProjectIssueFieldWithDefaults instantiates a new ProjectIssueField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectIssueFieldWithDefaults() *ProjectIssueField {
	this := ProjectIssueField{}
	var issueType string = ""
	this.IssueType = &issueType
	var needDefault bool = false
	this.NeedDefault = &needDefault
	var required bool = false
	this.Required = &required
	var valueString string = ""
	this.ValueString = *utils.NewNullableString(&valueString)
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProjectIssueField) GetCreatedAt() int64 {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectIssueField) GetCreatedAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProjectIssueField) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *ProjectIssueField) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetIssueField returns the IssueField field value if set, zero value otherwise.
func (o *ProjectIssueField) GetIssueField() IssueField {
	if o == nil || utils.IsNil(o.IssueField) {
		var ret IssueField
		return ret
	}
	return *o.IssueField
}

// GetIssueFieldOk returns a tuple with the IssueField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectIssueField) GetIssueFieldOk() (*IssueField, bool) {
	if o == nil || utils.IsNil(o.IssueField) {
		return nil, false
	}
	return o.IssueField, true
}

// HasIssueField returns a boolean if a field has been set.
func (o *ProjectIssueField) HasIssueField() bool {
	if o != nil && !utils.IsNil(o.IssueField) {
		return true
	}

	return false
}

// SetIssueField gets a reference to the given IssueField and assigns it to the IssueField field.
func (o *ProjectIssueField) SetIssueField(v IssueField) {
	o.IssueField = &v
}

// GetIssueFieldId returns the IssueFieldId field value if set, zero value otherwise.
func (o *ProjectIssueField) GetIssueFieldId() int64 {
	if o == nil || utils.IsNil(o.IssueFieldId) {
		var ret int64
		return ret
	}
	return *o.IssueFieldId
}

// GetIssueFieldIdOk returns a tuple with the IssueFieldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectIssueField) GetIssueFieldIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.IssueFieldId) {
		return nil, false
	}
	return o.IssueFieldId, true
}

// HasIssueFieldId returns a boolean if a field has been set.
func (o *ProjectIssueField) HasIssueFieldId() bool {
	if o != nil && !utils.IsNil(o.IssueFieldId) {
		return true
	}

	return false
}

// SetIssueFieldId gets a reference to the given int64 and assigns it to the IssueFieldId field.
func (o *ProjectIssueField) SetIssueFieldId(v int64) {
	o.IssueFieldId = &v
}

// GetIssueType returns the IssueType field value if set, zero value otherwise.
func (o *ProjectIssueField) GetIssueType() string {
	if o == nil || utils.IsNil(o.IssueType) {
		var ret string
		return ret
	}
	return *o.IssueType
}

// GetIssueTypeOk returns a tuple with the IssueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectIssueField) GetIssueTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IssueType) {
		return nil, false
	}
	return o.IssueType, true
}

// HasIssueType returns a boolean if a field has been set.
func (o *ProjectIssueField) HasIssueType() bool {
	if o != nil && !utils.IsNil(o.IssueType) {
		return true
	}

	return false
}

// SetIssueType gets a reference to the given string and assigns it to the IssueType field.
func (o *ProjectIssueField) SetIssueType(v string) {
	o.IssueType = &v
}

// GetNeedDefault returns the NeedDefault field value if set, zero value otherwise.
func (o *ProjectIssueField) GetNeedDefault() bool {
	if o == nil || utils.IsNil(o.NeedDefault) {
		var ret bool
		return ret
	}
	return *o.NeedDefault
}

// GetNeedDefaultOk returns a tuple with the NeedDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectIssueField) GetNeedDefaultOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NeedDefault) {
		return nil, false
	}
	return o.NeedDefault, true
}

// HasNeedDefault returns a boolean if a field has been set.
func (o *ProjectIssueField) HasNeedDefault() bool {
	if o != nil && !utils.IsNil(o.NeedDefault) {
		return true
	}

	return false
}

// SetNeedDefault gets a reference to the given bool and assigns it to the NeedDefault field.
func (o *ProjectIssueField) SetNeedDefault(v bool) {
	o.NeedDefault = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ProjectIssueField) GetRequired() bool {
	if o == nil || utils.IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectIssueField) GetRequiredOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ProjectIssueField) HasRequired() bool {
	if o != nil && !utils.IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *ProjectIssueField) SetRequired(v bool) {
	o.Required = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProjectIssueField) GetUpdatedAt() int64 {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectIssueField) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProjectIssueField) HasUpdatedAt() bool {
	if o != nil && !utils.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *ProjectIssueField) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetValueString returns the ValueString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProjectIssueField) GetValueString() string {
	if o == nil || utils.IsNil(o.ValueString.Get()) {
		var ret string
		return ret
	}
	return *o.ValueString.Get()
}

// GetValueStringOk returns a tuple with the ValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectIssueField) GetValueStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValueString.Get(), o.ValueString.IsSet()
}

// HasValueString returns a boolean if a field has been set.
func (o *ProjectIssueField) HasValueString() bool {
	if o != nil && o.ValueString.IsSet() {
		return true
	}

	return false
}

// SetValueString gets a reference to the given utils.NullableString and assigns it to the ValueString field.
func (o *ProjectIssueField) SetValueString(v string) {
	o.ValueString.Set(&v)
}
// SetValueStringNil sets the value for ValueString to be an explicit nil
func (o *ProjectIssueField) SetValueStringNil() {
	o.ValueString.Set(nil)
}

// UnsetValueString ensures that no value is present for ValueString, not even an explicit nil
func (o *ProjectIssueField) UnsetValueString() {
	o.ValueString.Unset()
}

func (o ProjectIssueField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectIssueField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if !utils.IsNil(o.IssueField) {
		toSerialize["IssueField"] = o.IssueField
	}
	if !utils.IsNil(o.IssueFieldId) {
		toSerialize["IssueFieldId"] = o.IssueFieldId
	}
	if !utils.IsNil(o.IssueType) {
		toSerialize["IssueType"] = o.IssueType
	}
	if !utils.IsNil(o.NeedDefault) {
		toSerialize["NeedDefault"] = o.NeedDefault
	}
	if !utils.IsNil(o.Required) {
		toSerialize["Required"] = o.Required
	}
	if !utils.IsNil(o.UpdatedAt) {
		toSerialize["UpdatedAt"] = o.UpdatedAt
	}
	if o.ValueString.IsSet() {
		toSerialize["ValueString"] = o.ValueString.Get()
	}
	return toSerialize, nil
}

type NullableProjectIssueField struct {
	value *ProjectIssueField
	isSet bool
}

func (v NullableProjectIssueField) Get() *ProjectIssueField {
	return v.value
}

func (v *NullableProjectIssueField) Set(val *ProjectIssueField) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectIssueField) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectIssueField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectIssueField(val *ProjectIssueField) *NullableProjectIssueField {
	return &NullableProjectIssueField{value: val, isSet: true}
}

func (v NullableProjectIssueField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectIssueField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


