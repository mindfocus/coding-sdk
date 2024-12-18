/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CustomFieldChangeLog type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CustomFieldChangeLog{}

// CustomFieldChangeLog 事项自定义属性变更日志
type CustomFieldChangeLog struct {
	// 动作类型 
	ActionType utils.NullableString `json:"ActionType,omitempty"`
	// 创建人ID 
	Creator utils.NullableInt64 `json:"Creator,omitempty"`
	// 自定义属性ID 
	FieldId utils.NullableInt64 `json:"FieldId,omitempty"`
	// 自定义属性名字 
	FieldName utils.NullableString `json:"FieldName,omitempty"`
	// 自定义属性类型 
	FieldType utils.NullableString `json:"FieldType,omitempty"`
	// 自定义属性的值 
	FieldValue utils.NullableString `json:"FieldValue,omitempty"`
	// 事项id（不是页面上的ID，是数据库中的唯一ID）
	IssueId utils.NullableInt64 `json:"IssueId,omitempty"`
	// 更新时间戳 
	UpdatedAt utils.NullableInt64 `json:"UpdatedAt,omitempty"`
}

// NewCustomFieldChangeLog instantiates a new CustomFieldChangeLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldChangeLog() *CustomFieldChangeLog {
	this := CustomFieldChangeLog{}
	var actionType string = ""
	this.ActionType = *utils.NewNullableString(&actionType)
	var fieldName string = ""
	this.FieldName = *utils.NewNullableString(&fieldName)
	var fieldType string = ""
	this.FieldType = *utils.NewNullableString(&fieldType)
	var fieldValue string = ""
	this.FieldValue = *utils.NewNullableString(&fieldValue)
	return &this
}

// NewCustomFieldChangeLogWithDefaults instantiates a new CustomFieldChangeLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldChangeLogWithDefaults() *CustomFieldChangeLog {
	this := CustomFieldChangeLog{}
	var actionType string = ""
	this.ActionType = *utils.NewNullableString(&actionType)
	var fieldName string = ""
	this.FieldName = *utils.NewNullableString(&fieldName)
	var fieldType string = ""
	this.FieldType = *utils.NewNullableString(&fieldType)
	var fieldValue string = ""
	this.FieldValue = *utils.NewNullableString(&fieldValue)
	return &this
}

// GetActionType returns the ActionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFieldChangeLog) GetActionType() string {
	if o == nil || utils.IsNil(o.ActionType.Get()) {
		var ret string
		return ret
	}
	return *o.ActionType.Get()
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFieldChangeLog) GetActionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionType.Get(), o.ActionType.IsSet()
}

// HasActionType returns a boolean if a field has been set.
func (o *CustomFieldChangeLog) HasActionType() bool {
	if o != nil && o.ActionType.IsSet() {
		return true
	}

	return false
}

// SetActionType gets a reference to the given utils.NullableString and assigns it to the ActionType field.
func (o *CustomFieldChangeLog) SetActionType(v string) {
	o.ActionType.Set(&v)
}
// SetActionTypeNil sets the value for ActionType to be an explicit nil
func (o *CustomFieldChangeLog) SetActionTypeNil() {
	o.ActionType.Set(nil)
}

// UnsetActionType ensures that no value is present for ActionType, not even an explicit nil
func (o *CustomFieldChangeLog) UnsetActionType() {
	o.ActionType.Unset()
}

// GetCreator returns the Creator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFieldChangeLog) GetCreator() int64 {
	if o == nil || utils.IsNil(o.Creator.Get()) {
		var ret int64
		return ret
	}
	return *o.Creator.Get()
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFieldChangeLog) GetCreatorOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Creator.Get(), o.Creator.IsSet()
}

// HasCreator returns a boolean if a field has been set.
func (o *CustomFieldChangeLog) HasCreator() bool {
	if o != nil && o.Creator.IsSet() {
		return true
	}

	return false
}

// SetCreator gets a reference to the given utils.NullableInt64 and assigns it to the Creator field.
func (o *CustomFieldChangeLog) SetCreator(v int64) {
	o.Creator.Set(&v)
}
// SetCreatorNil sets the value for Creator to be an explicit nil
func (o *CustomFieldChangeLog) SetCreatorNil() {
	o.Creator.Set(nil)
}

// UnsetCreator ensures that no value is present for Creator, not even an explicit nil
func (o *CustomFieldChangeLog) UnsetCreator() {
	o.Creator.Unset()
}

// GetFieldId returns the FieldId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFieldChangeLog) GetFieldId() int64 {
	if o == nil || utils.IsNil(o.FieldId.Get()) {
		var ret int64
		return ret
	}
	return *o.FieldId.Get()
}

// GetFieldIdOk returns a tuple with the FieldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFieldChangeLog) GetFieldIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FieldId.Get(), o.FieldId.IsSet()
}

// HasFieldId returns a boolean if a field has been set.
func (o *CustomFieldChangeLog) HasFieldId() bool {
	if o != nil && o.FieldId.IsSet() {
		return true
	}

	return false
}

// SetFieldId gets a reference to the given utils.NullableInt64 and assigns it to the FieldId field.
func (o *CustomFieldChangeLog) SetFieldId(v int64) {
	o.FieldId.Set(&v)
}
// SetFieldIdNil sets the value for FieldId to be an explicit nil
func (o *CustomFieldChangeLog) SetFieldIdNil() {
	o.FieldId.Set(nil)
}

// UnsetFieldId ensures that no value is present for FieldId, not even an explicit nil
func (o *CustomFieldChangeLog) UnsetFieldId() {
	o.FieldId.Unset()
}

// GetFieldName returns the FieldName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFieldChangeLog) GetFieldName() string {
	if o == nil || utils.IsNil(o.FieldName.Get()) {
		var ret string
		return ret
	}
	return *o.FieldName.Get()
}

// GetFieldNameOk returns a tuple with the FieldName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFieldChangeLog) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FieldName.Get(), o.FieldName.IsSet()
}

// HasFieldName returns a boolean if a field has been set.
func (o *CustomFieldChangeLog) HasFieldName() bool {
	if o != nil && o.FieldName.IsSet() {
		return true
	}

	return false
}

// SetFieldName gets a reference to the given utils.NullableString and assigns it to the FieldName field.
func (o *CustomFieldChangeLog) SetFieldName(v string) {
	o.FieldName.Set(&v)
}
// SetFieldNameNil sets the value for FieldName to be an explicit nil
func (o *CustomFieldChangeLog) SetFieldNameNil() {
	o.FieldName.Set(nil)
}

// UnsetFieldName ensures that no value is present for FieldName, not even an explicit nil
func (o *CustomFieldChangeLog) UnsetFieldName() {
	o.FieldName.Unset()
}

// GetFieldType returns the FieldType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFieldChangeLog) GetFieldType() string {
	if o == nil || utils.IsNil(o.FieldType.Get()) {
		var ret string
		return ret
	}
	return *o.FieldType.Get()
}

// GetFieldTypeOk returns a tuple with the FieldType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFieldChangeLog) GetFieldTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FieldType.Get(), o.FieldType.IsSet()
}

// HasFieldType returns a boolean if a field has been set.
func (o *CustomFieldChangeLog) HasFieldType() bool {
	if o != nil && o.FieldType.IsSet() {
		return true
	}

	return false
}

// SetFieldType gets a reference to the given utils.NullableString and assigns it to the FieldType field.
func (o *CustomFieldChangeLog) SetFieldType(v string) {
	o.FieldType.Set(&v)
}
// SetFieldTypeNil sets the value for FieldType to be an explicit nil
func (o *CustomFieldChangeLog) SetFieldTypeNil() {
	o.FieldType.Set(nil)
}

// UnsetFieldType ensures that no value is present for FieldType, not even an explicit nil
func (o *CustomFieldChangeLog) UnsetFieldType() {
	o.FieldType.Unset()
}

// GetFieldValue returns the FieldValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFieldChangeLog) GetFieldValue() string {
	if o == nil || utils.IsNil(o.FieldValue.Get()) {
		var ret string
		return ret
	}
	return *o.FieldValue.Get()
}

// GetFieldValueOk returns a tuple with the FieldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFieldChangeLog) GetFieldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FieldValue.Get(), o.FieldValue.IsSet()
}

// HasFieldValue returns a boolean if a field has been set.
func (o *CustomFieldChangeLog) HasFieldValue() bool {
	if o != nil && o.FieldValue.IsSet() {
		return true
	}

	return false
}

// SetFieldValue gets a reference to the given utils.NullableString and assigns it to the FieldValue field.
func (o *CustomFieldChangeLog) SetFieldValue(v string) {
	o.FieldValue.Set(&v)
}
// SetFieldValueNil sets the value for FieldValue to be an explicit nil
func (o *CustomFieldChangeLog) SetFieldValueNil() {
	o.FieldValue.Set(nil)
}

// UnsetFieldValue ensures that no value is present for FieldValue, not even an explicit nil
func (o *CustomFieldChangeLog) UnsetFieldValue() {
	o.FieldValue.Unset()
}

// GetIssueId returns the IssueId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFieldChangeLog) GetIssueId() int64 {
	if o == nil || utils.IsNil(o.IssueId.Get()) {
		var ret int64
		return ret
	}
	return *o.IssueId.Get()
}

// GetIssueIdOk returns a tuple with the IssueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFieldChangeLog) GetIssueIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssueId.Get(), o.IssueId.IsSet()
}

// HasIssueId returns a boolean if a field has been set.
func (o *CustomFieldChangeLog) HasIssueId() bool {
	if o != nil && o.IssueId.IsSet() {
		return true
	}

	return false
}

// SetIssueId gets a reference to the given utils.NullableInt64 and assigns it to the IssueId field.
func (o *CustomFieldChangeLog) SetIssueId(v int64) {
	o.IssueId.Set(&v)
}
// SetIssueIdNil sets the value for IssueId to be an explicit nil
func (o *CustomFieldChangeLog) SetIssueIdNil() {
	o.IssueId.Set(nil)
}

// UnsetIssueId ensures that no value is present for IssueId, not even an explicit nil
func (o *CustomFieldChangeLog) UnsetIssueId() {
	o.IssueId.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFieldChangeLog) GetUpdatedAt() int64 {
	if o == nil || utils.IsNil(o.UpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFieldChangeLog) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomFieldChangeLog) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given utils.NullableInt64 and assigns it to the UpdatedAt field.
func (o *CustomFieldChangeLog) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *CustomFieldChangeLog) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *CustomFieldChangeLog) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

func (o CustomFieldChangeLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldChangeLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ActionType.IsSet() {
		toSerialize["ActionType"] = o.ActionType.Get()
	}
	if o.Creator.IsSet() {
		toSerialize["Creator"] = o.Creator.Get()
	}
	if o.FieldId.IsSet() {
		toSerialize["FieldId"] = o.FieldId.Get()
	}
	if o.FieldName.IsSet() {
		toSerialize["FieldName"] = o.FieldName.Get()
	}
	if o.FieldType.IsSet() {
		toSerialize["FieldType"] = o.FieldType.Get()
	}
	if o.FieldValue.IsSet() {
		toSerialize["FieldValue"] = o.FieldValue.Get()
	}
	if o.IssueId.IsSet() {
		toSerialize["IssueId"] = o.IssueId.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["UpdatedAt"] = o.UpdatedAt.Get()
	}
	return toSerialize, nil
}

type NullableCustomFieldChangeLog struct {
	value *CustomFieldChangeLog
	isSet bool
}

func (v NullableCustomFieldChangeLog) Get() *CustomFieldChangeLog {
	return v.value
}

func (v *NullableCustomFieldChangeLog) Set(val *CustomFieldChangeLog) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldChangeLog) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldChangeLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldChangeLog(val *CustomFieldChangeLog) *NullableCustomFieldChangeLog {
	return &NullableCustomFieldChangeLog{value: val, isSet: true}
}

func (v NullableCustomFieldChangeLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldChangeLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


