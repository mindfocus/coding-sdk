/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the Epic type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Epic{}

// Epic 史诗信息
type Epic struct {
	Assignee *User `json:"Assignee,omitempty"`
	// 史诗 Code
	Code *int64 `json:"Code,omitempty"`
	// 事项状态 Id
	IssueStatusId utils.NullableInt64 `json:"IssueStatusId,omitempty"`
	// 事项状态名称
	IssueStatusName utils.NullableString `json:"IssueStatusName,omitempty"`
	// 名称
	Name *string `json:"Name,omitempty"`
	// 优先级：  \"0\" - 低，  \"1\" - 中，  \"2\" - 高，  \"3\" - 紧急，  \"\" - 未指定
	Priority utils.NullableString `json:"Priority,omitempty"`
	// 史诗 Type
	Type *string `json:"Type,omitempty"`
}

// NewEpic instantiates a new Epic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEpic() *Epic {
	this := Epic{}
	var issueStatusName string = ""
	this.IssueStatusName = *utils.NewNullableString(&issueStatusName)
	var name string = ""
	this.Name = &name
	var priority string = ""
	this.Priority = *utils.NewNullableString(&priority)
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewEpicWithDefaults instantiates a new Epic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEpicWithDefaults() *Epic {
	this := Epic{}
	var issueStatusName string = ""
	this.IssueStatusName = *utils.NewNullableString(&issueStatusName)
	var name string = ""
	this.Name = &name
	var priority string = ""
	this.Priority = *utils.NewNullableString(&priority)
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *Epic) GetAssignee() User {
	if o == nil || utils.IsNil(o.Assignee) {
		var ret User
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Epic) GetAssigneeOk() (*User, bool) {
	if o == nil || utils.IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *Epic) HasAssignee() bool {
	if o != nil && !utils.IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given User and assigns it to the Assignee field.
func (o *Epic) SetAssignee(v User) {
	o.Assignee = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Epic) GetCode() int64 {
	if o == nil || utils.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Epic) GetCodeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Epic) HasCode() bool {
	if o != nil && !utils.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *Epic) SetCode(v int64) {
	o.Code = &v
}

// GetIssueStatusId returns the IssueStatusId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Epic) GetIssueStatusId() int64 {
	if o == nil || utils.IsNil(o.IssueStatusId.Get()) {
		var ret int64
		return ret
	}
	return *o.IssueStatusId.Get()
}

// GetIssueStatusIdOk returns a tuple with the IssueStatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Epic) GetIssueStatusIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssueStatusId.Get(), o.IssueStatusId.IsSet()
}

// HasIssueStatusId returns a boolean if a field has been set.
func (o *Epic) HasIssueStatusId() bool {
	if o != nil && o.IssueStatusId.IsSet() {
		return true
	}

	return false
}

// SetIssueStatusId gets a reference to the given utils.NullableInt64 and assigns it to the IssueStatusId field.
func (o *Epic) SetIssueStatusId(v int64) {
	o.IssueStatusId.Set(&v)
}
// SetIssueStatusIdNil sets the value for IssueStatusId to be an explicit nil
func (o *Epic) SetIssueStatusIdNil() {
	o.IssueStatusId.Set(nil)
}

// UnsetIssueStatusId ensures that no value is present for IssueStatusId, not even an explicit nil
func (o *Epic) UnsetIssueStatusId() {
	o.IssueStatusId.Unset()
}

// GetIssueStatusName returns the IssueStatusName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Epic) GetIssueStatusName() string {
	if o == nil || utils.IsNil(o.IssueStatusName.Get()) {
		var ret string
		return ret
	}
	return *o.IssueStatusName.Get()
}

// GetIssueStatusNameOk returns a tuple with the IssueStatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Epic) GetIssueStatusNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssueStatusName.Get(), o.IssueStatusName.IsSet()
}

// HasIssueStatusName returns a boolean if a field has been set.
func (o *Epic) HasIssueStatusName() bool {
	if o != nil && o.IssueStatusName.IsSet() {
		return true
	}

	return false
}

// SetIssueStatusName gets a reference to the given utils.NullableString and assigns it to the IssueStatusName field.
func (o *Epic) SetIssueStatusName(v string) {
	o.IssueStatusName.Set(&v)
}
// SetIssueStatusNameNil sets the value for IssueStatusName to be an explicit nil
func (o *Epic) SetIssueStatusNameNil() {
	o.IssueStatusName.Set(nil)
}

// UnsetIssueStatusName ensures that no value is present for IssueStatusName, not even an explicit nil
func (o *Epic) UnsetIssueStatusName() {
	o.IssueStatusName.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Epic) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Epic) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Epic) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Epic) SetName(v string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Epic) GetPriority() string {
	if o == nil || utils.IsNil(o.Priority.Get()) {
		var ret string
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Epic) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *Epic) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given utils.NullableString and assigns it to the Priority field.
func (o *Epic) SetPriority(v string) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *Epic) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *Epic) UnsetPriority() {
	o.Priority.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Epic) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Epic) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Epic) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Epic) SetType(v string) {
	o.Type = &v
}

func (o Epic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Epic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Assignee) {
		toSerialize["Assignee"] = o.Assignee
	}
	if !utils.IsNil(o.Code) {
		toSerialize["Code"] = o.Code
	}
	if o.IssueStatusId.IsSet() {
		toSerialize["IssueStatusId"] = o.IssueStatusId.Get()
	}
	if o.IssueStatusName.IsSet() {
		toSerialize["IssueStatusName"] = o.IssueStatusName.Get()
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.Priority.IsSet() {
		toSerialize["Priority"] = o.Priority.Get()
	}
	if !utils.IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	return toSerialize, nil
}

type NullableEpic struct {
	value *Epic
	isSet bool
}

func (v NullableEpic) Get() *Epic {
	return v.value
}

func (v *NullableEpic) Set(val *Epic) {
	v.value = val
	v.isSet = true
}

func (v NullableEpic) IsSet() bool {
	return v.isSet
}

func (v *NullableEpic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpic(val *Epic) *NullableEpic {
	return &NullableEpic{value: val, isSet: true}
}

func (v NullableEpic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


