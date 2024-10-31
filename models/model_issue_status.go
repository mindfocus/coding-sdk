/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the IssueStatus type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &IssueStatus{}

// IssueStatus 事项状态
type IssueStatus struct {
	// 创建时间戳
	CreatedAt *int64 `json:"CreatedAt,omitempty"`
	// 描述
	Description utils.NullableString `json:"Description,omitempty"`
	// 事项状态 ID
	Id *int64 `json:"Id,omitempty"`
	// 状态序号
	Index *int64 `json:"Index,omitempty"`
	// 是否是系统内置
	IsSystem *bool `json:"IsSystem,omitempty"`
	// 名称
	Name *string `json:"Name,omitempty"`
	// 类型：TODO、PROCESSING、COMPLETED
	Type *string `json:"Type,omitempty"`
	// 修改时间戳
	UpdatedAt *int64 `json:"UpdatedAt,omitempty"`
}

// NewIssueStatus instantiates a new IssueStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueStatus() *IssueStatus {
	this := IssueStatus{}
	var description string = ""
	this.Description = *utils.NewNullableString(&description)
	var isSystem bool = false
	this.IsSystem = &isSystem
	var name string = ""
	this.Name = &name
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewIssueStatusWithDefaults instantiates a new IssueStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueStatusWithDefaults() *IssueStatus {
	this := IssueStatus{}
	var description string = ""
	this.Description = *utils.NewNullableString(&description)
	var isSystem bool = false
	this.IsSystem = &isSystem
	var name string = ""
	this.Name = &name
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IssueStatus) GetCreatedAt() int64 {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueStatus) GetCreatedAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IssueStatus) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *IssueStatus) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueStatus) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueStatus) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *IssueStatus) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given utils.NullableString and assigns it to the Description field.
func (o *IssueStatus) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *IssueStatus) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *IssueStatus) UnsetDescription() {
	o.Description.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IssueStatus) GetId() int64 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueStatus) GetIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IssueStatus) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *IssueStatus) SetId(v int64) {
	o.Id = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *IssueStatus) GetIndex() int64 {
	if o == nil || utils.IsNil(o.Index) {
		var ret int64
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueStatus) GetIndexOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *IssueStatus) HasIndex() bool {
	if o != nil && !utils.IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int64 and assigns it to the Index field.
func (o *IssueStatus) SetIndex(v int64) {
	o.Index = &v
}

// GetIsSystem returns the IsSystem field value if set, zero value otherwise.
func (o *IssueStatus) GetIsSystem() bool {
	if o == nil || utils.IsNil(o.IsSystem) {
		var ret bool
		return ret
	}
	return *o.IsSystem
}

// GetIsSystemOk returns a tuple with the IsSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueStatus) GetIsSystemOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsSystem) {
		return nil, false
	}
	return o.IsSystem, true
}

// HasIsSystem returns a boolean if a field has been set.
func (o *IssueStatus) HasIsSystem() bool {
	if o != nil && !utils.IsNil(o.IsSystem) {
		return true
	}

	return false
}

// SetIsSystem gets a reference to the given bool and assigns it to the IsSystem field.
func (o *IssueStatus) SetIsSystem(v bool) {
	o.IsSystem = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IssueStatus) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueStatus) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IssueStatus) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IssueStatus) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IssueStatus) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueStatus) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IssueStatus) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IssueStatus) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IssueStatus) GetUpdatedAt() int64 {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueStatus) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IssueStatus) HasUpdatedAt() bool {
	if o != nil && !utils.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *IssueStatus) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

func (o IssueStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !utils.IsNil(o.Index) {
		toSerialize["Index"] = o.Index
	}
	if !utils.IsNil(o.IsSystem) {
		toSerialize["IsSystem"] = o.IsSystem
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !utils.IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !utils.IsNil(o.UpdatedAt) {
		toSerialize["UpdatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableIssueStatus struct {
	value *IssueStatus
	isSet bool
}

func (v NullableIssueStatus) Get() *IssueStatus {
	return v.value
}

func (v *NullableIssueStatus) Set(val *IssueStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueStatus(val *IssueStatus) *NullableIssueStatus {
	return &NullableIssueStatus{value: val, isSet: true}
}

func (v NullableIssueStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


