/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DepotSpecDepotLevelParam type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DepotSpecDepotLevelParam{}

// DepotSpecDepotLevelParam 仓库规范 - 新增、修改仓库级别规范的参数
type DepotSpecDepotLevelParam struct {
	// 允许创建规定分支类型以外的分支
	AllowPushWildRef *bool `json:"AllowPushWildRef,omitempty"`
	// 分支类型列表
	DepotBranchTypeList []DepotBranchTypeParam `json:"DepotBranchTypeList,omitempty"`
	// 合并方向规则列表
	DepotMergeRequestRuleList []DepotMergeRequestRuleParam `json:"DepotMergeRequestRuleList,omitempty"`
}

// NewDepotSpecDepotLevelParam instantiates a new DepotSpecDepotLevelParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepotSpecDepotLevelParam() *DepotSpecDepotLevelParam {
	this := DepotSpecDepotLevelParam{}
	var allowPushWildRef bool = false
	this.AllowPushWildRef = &allowPushWildRef
	return &this
}

// NewDepotSpecDepotLevelParamWithDefaults instantiates a new DepotSpecDepotLevelParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepotSpecDepotLevelParamWithDefaults() *DepotSpecDepotLevelParam {
	this := DepotSpecDepotLevelParam{}
	var allowPushWildRef bool = false
	this.AllowPushWildRef = &allowPushWildRef
	return &this
}

// GetAllowPushWildRef returns the AllowPushWildRef field value if set, zero value otherwise.
func (o *DepotSpecDepotLevelParam) GetAllowPushWildRef() bool {
	if o == nil || utils.IsNil(o.AllowPushWildRef) {
		var ret bool
		return ret
	}
	return *o.AllowPushWildRef
}

// GetAllowPushWildRefOk returns a tuple with the AllowPushWildRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotSpecDepotLevelParam) GetAllowPushWildRefOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AllowPushWildRef) {
		return nil, false
	}
	return o.AllowPushWildRef, true
}

// HasAllowPushWildRef returns a boolean if a field has been set.
func (o *DepotSpecDepotLevelParam) HasAllowPushWildRef() bool {
	if o != nil && !utils.IsNil(o.AllowPushWildRef) {
		return true
	}

	return false
}

// SetAllowPushWildRef gets a reference to the given bool and assigns it to the AllowPushWildRef field.
func (o *DepotSpecDepotLevelParam) SetAllowPushWildRef(v bool) {
	o.AllowPushWildRef = &v
}

// GetDepotBranchTypeList returns the DepotBranchTypeList field value if set, zero value otherwise.
func (o *DepotSpecDepotLevelParam) GetDepotBranchTypeList() []DepotBranchTypeParam {
	if o == nil || utils.IsNil(o.DepotBranchTypeList) {
		var ret []DepotBranchTypeParam
		return ret
	}
	return o.DepotBranchTypeList
}

// GetDepotBranchTypeListOk returns a tuple with the DepotBranchTypeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotSpecDepotLevelParam) GetDepotBranchTypeListOk() ([]DepotBranchTypeParam, bool) {
	if o == nil || utils.IsNil(o.DepotBranchTypeList) {
		return nil, false
	}
	return o.DepotBranchTypeList, true
}

// HasDepotBranchTypeList returns a boolean if a field has been set.
func (o *DepotSpecDepotLevelParam) HasDepotBranchTypeList() bool {
	if o != nil && !utils.IsNil(o.DepotBranchTypeList) {
		return true
	}

	return false
}

// SetDepotBranchTypeList gets a reference to the given []DepotBranchTypeParam and assigns it to the DepotBranchTypeList field.
func (o *DepotSpecDepotLevelParam) SetDepotBranchTypeList(v []DepotBranchTypeParam) {
	o.DepotBranchTypeList = v
}

// GetDepotMergeRequestRuleList returns the DepotMergeRequestRuleList field value if set, zero value otherwise.
func (o *DepotSpecDepotLevelParam) GetDepotMergeRequestRuleList() []DepotMergeRequestRuleParam {
	if o == nil || utils.IsNil(o.DepotMergeRequestRuleList) {
		var ret []DepotMergeRequestRuleParam
		return ret
	}
	return o.DepotMergeRequestRuleList
}

// GetDepotMergeRequestRuleListOk returns a tuple with the DepotMergeRequestRuleList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotSpecDepotLevelParam) GetDepotMergeRequestRuleListOk() ([]DepotMergeRequestRuleParam, bool) {
	if o == nil || utils.IsNil(o.DepotMergeRequestRuleList) {
		return nil, false
	}
	return o.DepotMergeRequestRuleList, true
}

// HasDepotMergeRequestRuleList returns a boolean if a field has been set.
func (o *DepotSpecDepotLevelParam) HasDepotMergeRequestRuleList() bool {
	if o != nil && !utils.IsNil(o.DepotMergeRequestRuleList) {
		return true
	}

	return false
}

// SetDepotMergeRequestRuleList gets a reference to the given []DepotMergeRequestRuleParam and assigns it to the DepotMergeRequestRuleList field.
func (o *DepotSpecDepotLevelParam) SetDepotMergeRequestRuleList(v []DepotMergeRequestRuleParam) {
	o.DepotMergeRequestRuleList = v
}

func (o DepotSpecDepotLevelParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepotSpecDepotLevelParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AllowPushWildRef) {
		toSerialize["AllowPushWildRef"] = o.AllowPushWildRef
	}
	if !utils.IsNil(o.DepotBranchTypeList) {
		toSerialize["DepotBranchTypeList"] = o.DepotBranchTypeList
	}
	if !utils.IsNil(o.DepotMergeRequestRuleList) {
		toSerialize["DepotMergeRequestRuleList"] = o.DepotMergeRequestRuleList
	}
	return toSerialize, nil
}

type NullableDepotSpecDepotLevelParam struct {
	value *DepotSpecDepotLevelParam
	isSet bool
}

func (v NullableDepotSpecDepotLevelParam) Get() *DepotSpecDepotLevelParam {
	return v.value
}

func (v *NullableDepotSpecDepotLevelParam) Set(val *DepotSpecDepotLevelParam) {
	v.value = val
	v.isSet = true
}

func (v NullableDepotSpecDepotLevelParam) IsSet() bool {
	return v.isSet
}

func (v *NullableDepotSpecDepotLevelParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepotSpecDepotLevelParam(val *DepotSpecDepotLevelParam) *NullableDepotSpecDepotLevelParam {
	return &NullableDepotSpecDepotLevelParam{value: val, isSet: true}
}

func (v NullableDepotSpecDepotLevelParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepotSpecDepotLevelParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


