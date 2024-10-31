/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the MergeReqInfo type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MergeReqInfo{}

// MergeReqInfo 合并请求详情
type MergeReqInfo struct {
	// 描述,为 markdown 格式
	Describe *string `json:"Describe,omitempty"`
	// 源分支
	SourceBranch *string `json:"SourceBranch,omitempty"`
	// 合并状态,状态值如下:  CANMERGE:状态可自动合并;  ACCEPTED:状态已接受;  CANNOTMERGE:状态不可自动合并;  REFUSED:状态已拒绝(关闭);  CANCEL: 取消;  MERGING:正在合并中;  ABNORMAL:状态异常;
	Status *string `json:"Status,omitempty"`
	// 目标分支
	TargetBranch *string `json:"TargetBranch,omitempty"`
	// 合并标题
	Title *string `json:"Title,omitempty"`
}

// NewMergeReqInfo instantiates a new MergeReqInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeReqInfo() *MergeReqInfo {
	this := MergeReqInfo{}
	var describe string = ""
	this.Describe = &describe
	var sourceBranch string = ""
	this.SourceBranch = &sourceBranch
	var status string = ""
	this.Status = &status
	var targetBranch string = ""
	this.TargetBranch = &targetBranch
	var title string = ""
	this.Title = &title
	return &this
}

// NewMergeReqInfoWithDefaults instantiates a new MergeReqInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeReqInfoWithDefaults() *MergeReqInfo {
	this := MergeReqInfo{}
	var describe string = ""
	this.Describe = &describe
	var sourceBranch string = ""
	this.SourceBranch = &sourceBranch
	var status string = ""
	this.Status = &status
	var targetBranch string = ""
	this.TargetBranch = &targetBranch
	var title string = ""
	this.Title = &title
	return &this
}

// GetDescribe returns the Describe field value if set, zero value otherwise.
func (o *MergeReqInfo) GetDescribe() string {
	if o == nil || utils.IsNil(o.Describe) {
		var ret string
		return ret
	}
	return *o.Describe
}

// GetDescribeOk returns a tuple with the Describe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeReqInfo) GetDescribeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Describe) {
		return nil, false
	}
	return o.Describe, true
}

// HasDescribe returns a boolean if a field has been set.
func (o *MergeReqInfo) HasDescribe() bool {
	if o != nil && !utils.IsNil(o.Describe) {
		return true
	}

	return false
}

// SetDescribe gets a reference to the given string and assigns it to the Describe field.
func (o *MergeReqInfo) SetDescribe(v string) {
	o.Describe = &v
}

// GetSourceBranch returns the SourceBranch field value if set, zero value otherwise.
func (o *MergeReqInfo) GetSourceBranch() string {
	if o == nil || utils.IsNil(o.SourceBranch) {
		var ret string
		return ret
	}
	return *o.SourceBranch
}

// GetSourceBranchOk returns a tuple with the SourceBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeReqInfo) GetSourceBranchOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SourceBranch) {
		return nil, false
	}
	return o.SourceBranch, true
}

// HasSourceBranch returns a boolean if a field has been set.
func (o *MergeReqInfo) HasSourceBranch() bool {
	if o != nil && !utils.IsNil(o.SourceBranch) {
		return true
	}

	return false
}

// SetSourceBranch gets a reference to the given string and assigns it to the SourceBranch field.
func (o *MergeReqInfo) SetSourceBranch(v string) {
	o.SourceBranch = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MergeReqInfo) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeReqInfo) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MergeReqInfo) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MergeReqInfo) SetStatus(v string) {
	o.Status = &v
}

// GetTargetBranch returns the TargetBranch field value if set, zero value otherwise.
func (o *MergeReqInfo) GetTargetBranch() string {
	if o == nil || utils.IsNil(o.TargetBranch) {
		var ret string
		return ret
	}
	return *o.TargetBranch
}

// GetTargetBranchOk returns a tuple with the TargetBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeReqInfo) GetTargetBranchOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TargetBranch) {
		return nil, false
	}
	return o.TargetBranch, true
}

// HasTargetBranch returns a boolean if a field has been set.
func (o *MergeReqInfo) HasTargetBranch() bool {
	if o != nil && !utils.IsNil(o.TargetBranch) {
		return true
	}

	return false
}

// SetTargetBranch gets a reference to the given string and assigns it to the TargetBranch field.
func (o *MergeReqInfo) SetTargetBranch(v string) {
	o.TargetBranch = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *MergeReqInfo) GetTitle() string {
	if o == nil || utils.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeReqInfo) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *MergeReqInfo) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *MergeReqInfo) SetTitle(v string) {
	o.Title = &v
}

func (o MergeReqInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MergeReqInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Describe) {
		toSerialize["Describe"] = o.Describe
	}
	if !utils.IsNil(o.SourceBranch) {
		toSerialize["SourceBranch"] = o.SourceBranch
	}
	if !utils.IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !utils.IsNil(o.TargetBranch) {
		toSerialize["TargetBranch"] = o.TargetBranch
	}
	if !utils.IsNil(o.Title) {
		toSerialize["Title"] = o.Title
	}
	return toSerialize, nil
}

type NullableMergeReqInfo struct {
	value *MergeReqInfo
	isSet bool
}

func (v NullableMergeReqInfo) Get() *MergeReqInfo {
	return v.value
}

func (v *NullableMergeReqInfo) Set(val *MergeReqInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeReqInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeReqInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeReqInfo(val *MergeReqInfo) *NullableMergeReqInfo {
	return &NullableMergeReqInfo{value: val, isSet: true}
}

func (v NullableMergeReqInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeReqInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

