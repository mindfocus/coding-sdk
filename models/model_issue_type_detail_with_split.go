/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the IssueTypeDetailWithSplit type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &IssueTypeDetailWithSplit{}

// IssueTypeDetailWithSplit 事项类型详情(带分解类型)
type IssueTypeDetailWithSplit struct {
	// 描述
	Description *string `json:"Description,omitempty"`
	// 事项类型 ID
	Id *int64 `json:"Id,omitempty"`
	// 是否是系统类型
	IsSystem *bool `json:"IsSystem,omitempty"`
	// 事项类型大类
	IssueType *string `json:"IssueType,omitempty"`
	// 事项类型名称
	Name *string `json:"Name,omitempty"`
	// 可分解类型 ID，SplitType = SPECIFIC_TYPE 时需指定
	SplitTargetIssueTypeId []int64 `json:"SplitTargetIssueTypeId,omitempty"`
	// 需求分解类型，SPECIFIC_TYPE - 可分解为制定需求类型，UNSPLITTABLE - 不可分解需求，ALL_REQUIREMENT - 可分解为全部需求类型
	SplitType *string `json:"SplitType,omitempty"`
}

// NewIssueTypeDetailWithSplit instantiates a new IssueTypeDetailWithSplit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueTypeDetailWithSplit() *IssueTypeDetailWithSplit {
	this := IssueTypeDetailWithSplit{}
	var description string = ""
	this.Description = &description
	var isSystem bool = false
	this.IsSystem = &isSystem
	var issueType string = ""
	this.IssueType = &issueType
	var name string = ""
	this.Name = &name
	var splitType string = ""
	this.SplitType = &splitType
	return &this
}

// NewIssueTypeDetailWithSplitWithDefaults instantiates a new IssueTypeDetailWithSplit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueTypeDetailWithSplitWithDefaults() *IssueTypeDetailWithSplit {
	this := IssueTypeDetailWithSplit{}
	var description string = ""
	this.Description = &description
	var isSystem bool = false
	this.IsSystem = &isSystem
	var issueType string = ""
	this.IssueType = &issueType
	var name string = ""
	this.Name = &name
	var splitType string = ""
	this.SplitType = &splitType
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IssueTypeDetailWithSplit) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueTypeDetailWithSplit) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IssueTypeDetailWithSplit) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IssueTypeDetailWithSplit) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IssueTypeDetailWithSplit) GetId() int64 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueTypeDetailWithSplit) GetIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IssueTypeDetailWithSplit) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *IssueTypeDetailWithSplit) SetId(v int64) {
	o.Id = &v
}

// GetIsSystem returns the IsSystem field value if set, zero value otherwise.
func (o *IssueTypeDetailWithSplit) GetIsSystem() bool {
	if o == nil || utils.IsNil(o.IsSystem) {
		var ret bool
		return ret
	}
	return *o.IsSystem
}

// GetIsSystemOk returns a tuple with the IsSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueTypeDetailWithSplit) GetIsSystemOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsSystem) {
		return nil, false
	}
	return o.IsSystem, true
}

// HasIsSystem returns a boolean if a field has been set.
func (o *IssueTypeDetailWithSplit) HasIsSystem() bool {
	if o != nil && !utils.IsNil(o.IsSystem) {
		return true
	}

	return false
}

// SetIsSystem gets a reference to the given bool and assigns it to the IsSystem field.
func (o *IssueTypeDetailWithSplit) SetIsSystem(v bool) {
	o.IsSystem = &v
}

// GetIssueType returns the IssueType field value if set, zero value otherwise.
func (o *IssueTypeDetailWithSplit) GetIssueType() string {
	if o == nil || utils.IsNil(o.IssueType) {
		var ret string
		return ret
	}
	return *o.IssueType
}

// GetIssueTypeOk returns a tuple with the IssueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueTypeDetailWithSplit) GetIssueTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IssueType) {
		return nil, false
	}
	return o.IssueType, true
}

// HasIssueType returns a boolean if a field has been set.
func (o *IssueTypeDetailWithSplit) HasIssueType() bool {
	if o != nil && !utils.IsNil(o.IssueType) {
		return true
	}

	return false
}

// SetIssueType gets a reference to the given string and assigns it to the IssueType field.
func (o *IssueTypeDetailWithSplit) SetIssueType(v string) {
	o.IssueType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IssueTypeDetailWithSplit) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueTypeDetailWithSplit) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IssueTypeDetailWithSplit) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IssueTypeDetailWithSplit) SetName(v string) {
	o.Name = &v
}

// GetSplitTargetIssueTypeId returns the SplitTargetIssueTypeId field value if set, zero value otherwise.
func (o *IssueTypeDetailWithSplit) GetSplitTargetIssueTypeId() []int64 {
	if o == nil || utils.IsNil(o.SplitTargetIssueTypeId) {
		var ret []int64
		return ret
	}
	return o.SplitTargetIssueTypeId
}

// GetSplitTargetIssueTypeIdOk returns a tuple with the SplitTargetIssueTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueTypeDetailWithSplit) GetSplitTargetIssueTypeIdOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.SplitTargetIssueTypeId) {
		return nil, false
	}
	return o.SplitTargetIssueTypeId, true
}

// HasSplitTargetIssueTypeId returns a boolean if a field has been set.
func (o *IssueTypeDetailWithSplit) HasSplitTargetIssueTypeId() bool {
	if o != nil && !utils.IsNil(o.SplitTargetIssueTypeId) {
		return true
	}

	return false
}

// SetSplitTargetIssueTypeId gets a reference to the given []int64 and assigns it to the SplitTargetIssueTypeId field.
func (o *IssueTypeDetailWithSplit) SetSplitTargetIssueTypeId(v []int64) {
	o.SplitTargetIssueTypeId = v
}

// GetSplitType returns the SplitType field value if set, zero value otherwise.
func (o *IssueTypeDetailWithSplit) GetSplitType() string {
	if o == nil || utils.IsNil(o.SplitType) {
		var ret string
		return ret
	}
	return *o.SplitType
}

// GetSplitTypeOk returns a tuple with the SplitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueTypeDetailWithSplit) GetSplitTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SplitType) {
		return nil, false
	}
	return o.SplitType, true
}

// HasSplitType returns a boolean if a field has been set.
func (o *IssueTypeDetailWithSplit) HasSplitType() bool {
	if o != nil && !utils.IsNil(o.SplitType) {
		return true
	}

	return false
}

// SetSplitType gets a reference to the given string and assigns it to the SplitType field.
func (o *IssueTypeDetailWithSplit) SetSplitType(v string) {
	o.SplitType = &v
}

func (o IssueTypeDetailWithSplit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueTypeDetailWithSplit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !utils.IsNil(o.IsSystem) {
		toSerialize["IsSystem"] = o.IsSystem
	}
	if !utils.IsNil(o.IssueType) {
		toSerialize["IssueType"] = o.IssueType
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !utils.IsNil(o.SplitTargetIssueTypeId) {
		toSerialize["SplitTargetIssueTypeId"] = o.SplitTargetIssueTypeId
	}
	if !utils.IsNil(o.SplitType) {
		toSerialize["SplitType"] = o.SplitType
	}
	return toSerialize, nil
}

type NullableIssueTypeDetailWithSplit struct {
	value *IssueTypeDetailWithSplit
	isSet bool
}

func (v NullableIssueTypeDetailWithSplit) Get() *IssueTypeDetailWithSplit {
	return v.value
}

func (v *NullableIssueTypeDetailWithSplit) Set(val *IssueTypeDetailWithSplit) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueTypeDetailWithSplit) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueTypeDetailWithSplit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueTypeDetailWithSplit(val *IssueTypeDetailWithSplit) *NullableIssueTypeDetailWithSplit {
	return &NullableIssueTypeDetailWithSplit{value: val, isSet: true}
}

func (v NullableIssueTypeDetailWithSplit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueTypeDetailWithSplit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


