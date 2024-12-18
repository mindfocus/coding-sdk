/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the Program type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Program{}

// Program 项目/项目集数据
type Program struct {
	// 结束时间
	CreatedAt *int64 `json:"CreatedAt,omitempty"`
	// 描述信息
	Description *string `json:"Description,omitempty"`
	// 展示名称
	DisplayName *string `json:"DisplayName,omitempty"`
	// 结束时间
	EndDate *int64 `json:"EndDate,omitempty"`
	// 图标
	Icon *string `json:"Icon,omitempty"`
	// ID
	Id *int32 `json:"Id,omitempty"`
	// 标识
	Name *string `json:"Name,omitempty"`
	// 名称拼音
	NamePinyin *string `json:"NamePinyin,omitempty"`
	// 开始时间
	StartDate *int64 `json:"StartDate,omitempty"`
	// 修改时间
	UpdatedAt *int64 `json:"UpdatedAt,omitempty"`
}

// NewProgram instantiates a new Program object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgram() *Program {
	this := Program{}
	return &this
}

// NewProgramWithDefaults instantiates a new Program object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramWithDefaults() *Program {
	this := Program{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Program) GetCreatedAt() int64 {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Program) GetCreatedAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Program) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *Program) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Program) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Program) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Program) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Program) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Program) GetDisplayName() string {
	if o == nil || utils.IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Program) GetDisplayNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Program) HasDisplayName() bool {
	if o != nil && !utils.IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Program) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *Program) GetEndDate() int64 {
	if o == nil || utils.IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Program) GetEndDateOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *Program) HasEndDate() bool {
	if o != nil && !utils.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *Program) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *Program) GetIcon() string {
	if o == nil || utils.IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Program) GetIconOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *Program) HasIcon() bool {
	if o != nil && !utils.IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *Program) SetIcon(v string) {
	o.Icon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Program) GetId() int32 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Program) GetIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Program) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Program) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Program) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Program) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Program) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Program) SetName(v string) {
	o.Name = &v
}

// GetNamePinyin returns the NamePinyin field value if set, zero value otherwise.
func (o *Program) GetNamePinyin() string {
	if o == nil || utils.IsNil(o.NamePinyin) {
		var ret string
		return ret
	}
	return *o.NamePinyin
}

// GetNamePinyinOk returns a tuple with the NamePinyin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Program) GetNamePinyinOk() (*string, bool) {
	if o == nil || utils.IsNil(o.NamePinyin) {
		return nil, false
	}
	return o.NamePinyin, true
}

// HasNamePinyin returns a boolean if a field has been set.
func (o *Program) HasNamePinyin() bool {
	if o != nil && !utils.IsNil(o.NamePinyin) {
		return true
	}

	return false
}

// SetNamePinyin gets a reference to the given string and assigns it to the NamePinyin field.
func (o *Program) SetNamePinyin(v string) {
	o.NamePinyin = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Program) GetStartDate() int64 {
	if o == nil || utils.IsNil(o.StartDate) {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Program) GetStartDateOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Program) HasStartDate() bool {
	if o != nil && !utils.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *Program) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Program) GetUpdatedAt() int64 {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Program) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Program) HasUpdatedAt() bool {
	if o != nil && !utils.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *Program) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

func (o Program) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Program) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if !utils.IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !utils.IsNil(o.DisplayName) {
		toSerialize["DisplayName"] = o.DisplayName
	}
	if !utils.IsNil(o.EndDate) {
		toSerialize["EndDate"] = o.EndDate
	}
	if !utils.IsNil(o.Icon) {
		toSerialize["Icon"] = o.Icon
	}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !utils.IsNil(o.NamePinyin) {
		toSerialize["NamePinyin"] = o.NamePinyin
	}
	if !utils.IsNil(o.StartDate) {
		toSerialize["StartDate"] = o.StartDate
	}
	if !utils.IsNil(o.UpdatedAt) {
		toSerialize["UpdatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableProgram struct {
	value *Program
	isSet bool
}

func (v NullableProgram) Get() *Program {
	return v.value
}

func (v *NullableProgram) Set(val *Program) {
	v.value = val
	v.isSet = true
}

func (v NullableProgram) IsSet() bool {
	return v.isSet
}

func (v *NullableProgram) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgram(val *Program) *NullableProgram {
	return &NullableProgram{value: val, isSet: true}
}

func (v NullableProgram) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgram) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


