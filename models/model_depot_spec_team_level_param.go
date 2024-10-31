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

// checks if the DepotSpecTeamLevelParam type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DepotSpecTeamLevelParam{}

// DepotSpecTeamLevelParam 仓库规范 - 新增团队规范的参数
type DepotSpecTeamLevelParam struct {
	// 允许创建规定分支类型以外的分支
	AllowPushWildRef bool `json:"AllowPushWildRef"`
	// 分支类型列表
	DepotBranchTypeList []DepotBranchTypeParam `json:"DepotBranchTypeList,omitempty"`
	// 合并方向规则列表
	DepotMergeRequestRuleList []DepotMergeRequestRuleParam `json:"DepotMergeRequestRuleList,omitempty"`
	// 仓库规范描述
	Description *string `json:"Description,omitempty"`
	// 规范的名字唯一，当名字是已有规范的名字时，为修改；当名字不是已有规范名字时为新增
	Name string `json:"Name"`
	// 当需要修改已有规范的名字时，需要填写新名字
	ReName *string `json:"ReName,omitempty"`
}

type _DepotSpecTeamLevelParam DepotSpecTeamLevelParam

// NewDepotSpecTeamLevelParam instantiates a new DepotSpecTeamLevelParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepotSpecTeamLevelParam(allowPushWildRef bool, name string) *DepotSpecTeamLevelParam {
	this := DepotSpecTeamLevelParam{}
	this.AllowPushWildRef = allowPushWildRef
	var description string = ""
	this.Description = &description
	this.Name = name
	var reName string = ""
	this.ReName = &reName
	return &this
}

// NewDepotSpecTeamLevelParamWithDefaults instantiates a new DepotSpecTeamLevelParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepotSpecTeamLevelParamWithDefaults() *DepotSpecTeamLevelParam {
	this := DepotSpecTeamLevelParam{}
	var allowPushWildRef bool = false
	this.AllowPushWildRef = allowPushWildRef
	var description string = ""
	this.Description = &description
	var name string = ""
	this.Name = name
	var reName string = ""
	this.ReName = &reName
	return &this
}

// GetAllowPushWildRef returns the AllowPushWildRef field value
func (o *DepotSpecTeamLevelParam) GetAllowPushWildRef() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowPushWildRef
}

// GetAllowPushWildRefOk returns a tuple with the AllowPushWildRef field value
// and a boolean to check if the value has been set.
func (o *DepotSpecTeamLevelParam) GetAllowPushWildRefOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowPushWildRef, true
}

// SetAllowPushWildRef sets field value
func (o *DepotSpecTeamLevelParam) SetAllowPushWildRef(v bool) {
	o.AllowPushWildRef = v
}

// GetDepotBranchTypeList returns the DepotBranchTypeList field value if set, zero value otherwise.
func (o *DepotSpecTeamLevelParam) GetDepotBranchTypeList() []DepotBranchTypeParam {
	if o == nil || utils.IsNil(o.DepotBranchTypeList) {
		var ret []DepotBranchTypeParam
		return ret
	}
	return o.DepotBranchTypeList
}

// GetDepotBranchTypeListOk returns a tuple with the DepotBranchTypeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotSpecTeamLevelParam) GetDepotBranchTypeListOk() ([]DepotBranchTypeParam, bool) {
	if o == nil || utils.IsNil(o.DepotBranchTypeList) {
		return nil, false
	}
	return o.DepotBranchTypeList, true
}

// HasDepotBranchTypeList returns a boolean if a field has been set.
func (o *DepotSpecTeamLevelParam) HasDepotBranchTypeList() bool {
	if o != nil && !utils.IsNil(o.DepotBranchTypeList) {
		return true
	}

	return false
}

// SetDepotBranchTypeList gets a reference to the given []DepotBranchTypeParam and assigns it to the DepotBranchTypeList field.
func (o *DepotSpecTeamLevelParam) SetDepotBranchTypeList(v []DepotBranchTypeParam) {
	o.DepotBranchTypeList = v
}

// GetDepotMergeRequestRuleList returns the DepotMergeRequestRuleList field value if set, zero value otherwise.
func (o *DepotSpecTeamLevelParam) GetDepotMergeRequestRuleList() []DepotMergeRequestRuleParam {
	if o == nil || utils.IsNil(o.DepotMergeRequestRuleList) {
		var ret []DepotMergeRequestRuleParam
		return ret
	}
	return o.DepotMergeRequestRuleList
}

// GetDepotMergeRequestRuleListOk returns a tuple with the DepotMergeRequestRuleList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotSpecTeamLevelParam) GetDepotMergeRequestRuleListOk() ([]DepotMergeRequestRuleParam, bool) {
	if o == nil || utils.IsNil(o.DepotMergeRequestRuleList) {
		return nil, false
	}
	return o.DepotMergeRequestRuleList, true
}

// HasDepotMergeRequestRuleList returns a boolean if a field has been set.
func (o *DepotSpecTeamLevelParam) HasDepotMergeRequestRuleList() bool {
	if o != nil && !utils.IsNil(o.DepotMergeRequestRuleList) {
		return true
	}

	return false
}

// SetDepotMergeRequestRuleList gets a reference to the given []DepotMergeRequestRuleParam and assigns it to the DepotMergeRequestRuleList field.
func (o *DepotSpecTeamLevelParam) SetDepotMergeRequestRuleList(v []DepotMergeRequestRuleParam) {
	o.DepotMergeRequestRuleList = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DepotSpecTeamLevelParam) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotSpecTeamLevelParam) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DepotSpecTeamLevelParam) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DepotSpecTeamLevelParam) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *DepotSpecTeamLevelParam) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DepotSpecTeamLevelParam) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DepotSpecTeamLevelParam) SetName(v string) {
	o.Name = v
}

// GetReName returns the ReName field value if set, zero value otherwise.
func (o *DepotSpecTeamLevelParam) GetReName() string {
	if o == nil || utils.IsNil(o.ReName) {
		var ret string
		return ret
	}
	return *o.ReName
}

// GetReNameOk returns a tuple with the ReName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotSpecTeamLevelParam) GetReNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReName) {
		return nil, false
	}
	return o.ReName, true
}

// HasReName returns a boolean if a field has been set.
func (o *DepotSpecTeamLevelParam) HasReName() bool {
	if o != nil && !utils.IsNil(o.ReName) {
		return true
	}

	return false
}

// SetReName gets a reference to the given string and assigns it to the ReName field.
func (o *DepotSpecTeamLevelParam) SetReName(v string) {
	o.ReName = &v
}

func (o DepotSpecTeamLevelParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepotSpecTeamLevelParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AllowPushWildRef"] = o.AllowPushWildRef
	if !utils.IsNil(o.DepotBranchTypeList) {
		toSerialize["DepotBranchTypeList"] = o.DepotBranchTypeList
	}
	if !utils.IsNil(o.DepotMergeRequestRuleList) {
		toSerialize["DepotMergeRequestRuleList"] = o.DepotMergeRequestRuleList
	}
	if !utils.IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	toSerialize["Name"] = o.Name
	if !utils.IsNil(o.ReName) {
		toSerialize["ReName"] = o.ReName
	}
	return toSerialize, nil
}

func (o *DepotSpecTeamLevelParam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"AllowPushWildRef",
		"Name",
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

	varDepotSpecTeamLevelParam := _DepotSpecTeamLevelParam{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDepotSpecTeamLevelParam)

	if err != nil {
		return err
	}

	*o = DepotSpecTeamLevelParam(varDepotSpecTeamLevelParam)

	return err
}

type NullableDepotSpecTeamLevelParam struct {
	value *DepotSpecTeamLevelParam
	isSet bool
}

func (v NullableDepotSpecTeamLevelParam) Get() *DepotSpecTeamLevelParam {
	return v.value
}

func (v *NullableDepotSpecTeamLevelParam) Set(val *DepotSpecTeamLevelParam) {
	v.value = val
	v.isSet = true
}

func (v NullableDepotSpecTeamLevelParam) IsSet() bool {
	return v.isSet
}

func (v *NullableDepotSpecTeamLevelParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepotSpecTeamLevelParam(val *DepotSpecTeamLevelParam) *NullableDepotSpecTeamLevelParam {
	return &NullableDepotSpecTeamLevelParam{value: val, isSet: true}
}

func (v NullableDepotSpecTeamLevelParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepotSpecTeamLevelParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

