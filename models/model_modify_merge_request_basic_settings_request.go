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

// checks if the ModifyMergeRequestBasicSettingsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyMergeRequestBasicSettingsRequest{}

// ModifyMergeRequestBasicSettingsRequest struct for ModifyMergeRequestBasicSettingsRequest
type ModifyMergeRequestBasicSettingsRequest struct {
	// 合并请求源分支有新提交时是否自动取消合并授权
	CancelMrGrantAfterPushSrc bool `json:"CancelMrGrantAfterPushSrc"`
	// 是否默认以 Fast-Forward 模式合并
	DefaultFastForwardMerge bool `json:"DefaultFastForwardMerge"`
	// 合并请求默认分支
	DefaultTargetBranch *string `json:"DefaultTargetBranch,omitempty"`
	// 是否默认删除源分支
	DeleteSrcBranchAfterMerged bool `json:"DeleteSrcBranchAfterMerged"`
	// 仓库路径，格式：<team>/<project>/<depot>
	DepotPath string `json:"DepotPath"`
	// 是否开启状态检查
	DepotStatusCheckRequired bool `json:"DepotStatusCheckRequired"`
	// 合并前是否必须获得所有评审者的允许合并
	MrCheckAllReviewersAllow bool `json:"MrCheckAllReviewersAllow"`
	// 合并请求选择方式no_squash:默认直接合并default_squash:默认Squash合并 force_squash:只能Squash合并
	SquashOnMerge *string `json:"SquashOnMerge,omitempty"`
}

type _ModifyMergeRequestBasicSettingsRequest ModifyMergeRequestBasicSettingsRequest

// NewModifyMergeRequestBasicSettingsRequest instantiates a new ModifyMergeRequestBasicSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyMergeRequestBasicSettingsRequest(cancelMrGrantAfterPushSrc bool, defaultFastForwardMerge bool, deleteSrcBranchAfterMerged bool, depotPath string, depotStatusCheckRequired bool, mrCheckAllReviewersAllow bool) *ModifyMergeRequestBasicSettingsRequest {
	this := ModifyMergeRequestBasicSettingsRequest{}
	this.CancelMrGrantAfterPushSrc = cancelMrGrantAfterPushSrc
	this.DefaultFastForwardMerge = defaultFastForwardMerge
	this.DeleteSrcBranchAfterMerged = deleteSrcBranchAfterMerged
	this.DepotPath = depotPath
	this.DepotStatusCheckRequired = depotStatusCheckRequired
	this.MrCheckAllReviewersAllow = mrCheckAllReviewersAllow
	return &this
}

// NewModifyMergeRequestBasicSettingsRequestWithDefaults instantiates a new ModifyMergeRequestBasicSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyMergeRequestBasicSettingsRequestWithDefaults() *ModifyMergeRequestBasicSettingsRequest {
	this := ModifyMergeRequestBasicSettingsRequest{}
	return &this
}

// GetCancelMrGrantAfterPushSrc returns the CancelMrGrantAfterPushSrc field value
func (o *ModifyMergeRequestBasicSettingsRequest) GetCancelMrGrantAfterPushSrc() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CancelMrGrantAfterPushSrc
}

// GetCancelMrGrantAfterPushSrcOk returns a tuple with the CancelMrGrantAfterPushSrc field value
// and a boolean to check if the value has been set.
func (o *ModifyMergeRequestBasicSettingsRequest) GetCancelMrGrantAfterPushSrcOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CancelMrGrantAfterPushSrc, true
}

// SetCancelMrGrantAfterPushSrc sets field value
func (o *ModifyMergeRequestBasicSettingsRequest) SetCancelMrGrantAfterPushSrc(v bool) {
	o.CancelMrGrantAfterPushSrc = v
}

// GetDefaultFastForwardMerge returns the DefaultFastForwardMerge field value
func (o *ModifyMergeRequestBasicSettingsRequest) GetDefaultFastForwardMerge() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DefaultFastForwardMerge
}

// GetDefaultFastForwardMergeOk returns a tuple with the DefaultFastForwardMerge field value
// and a boolean to check if the value has been set.
func (o *ModifyMergeRequestBasicSettingsRequest) GetDefaultFastForwardMergeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultFastForwardMerge, true
}

// SetDefaultFastForwardMerge sets field value
func (o *ModifyMergeRequestBasicSettingsRequest) SetDefaultFastForwardMerge(v bool) {
	o.DefaultFastForwardMerge = v
}

// GetDefaultTargetBranch returns the DefaultTargetBranch field value if set, zero value otherwise.
func (o *ModifyMergeRequestBasicSettingsRequest) GetDefaultTargetBranch() string {
	if o == nil || utils.IsNil(o.DefaultTargetBranch) {
		var ret string
		return ret
	}
	return *o.DefaultTargetBranch
}

// GetDefaultTargetBranchOk returns a tuple with the DefaultTargetBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyMergeRequestBasicSettingsRequest) GetDefaultTargetBranchOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DefaultTargetBranch) {
		return nil, false
	}
	return o.DefaultTargetBranch, true
}

// HasDefaultTargetBranch returns a boolean if a field has been set.
func (o *ModifyMergeRequestBasicSettingsRequest) HasDefaultTargetBranch() bool {
	if o != nil && !utils.IsNil(o.DefaultTargetBranch) {
		return true
	}

	return false
}

// SetDefaultTargetBranch gets a reference to the given string and assigns it to the DefaultTargetBranch field.
func (o *ModifyMergeRequestBasicSettingsRequest) SetDefaultTargetBranch(v string) {
	o.DefaultTargetBranch = &v
}

// GetDeleteSrcBranchAfterMerged returns the DeleteSrcBranchAfterMerged field value
func (o *ModifyMergeRequestBasicSettingsRequest) GetDeleteSrcBranchAfterMerged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DeleteSrcBranchAfterMerged
}

// GetDeleteSrcBranchAfterMergedOk returns a tuple with the DeleteSrcBranchAfterMerged field value
// and a boolean to check if the value has been set.
func (o *ModifyMergeRequestBasicSettingsRequest) GetDeleteSrcBranchAfterMergedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeleteSrcBranchAfterMerged, true
}

// SetDeleteSrcBranchAfterMerged sets field value
func (o *ModifyMergeRequestBasicSettingsRequest) SetDeleteSrcBranchAfterMerged(v bool) {
	o.DeleteSrcBranchAfterMerged = v
}

// GetDepotPath returns the DepotPath field value
func (o *ModifyMergeRequestBasicSettingsRequest) GetDepotPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value
// and a boolean to check if the value has been set.
func (o *ModifyMergeRequestBasicSettingsRequest) GetDepotPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotPath, true
}

// SetDepotPath sets field value
func (o *ModifyMergeRequestBasicSettingsRequest) SetDepotPath(v string) {
	o.DepotPath = v
}

// GetDepotStatusCheckRequired returns the DepotStatusCheckRequired field value
func (o *ModifyMergeRequestBasicSettingsRequest) GetDepotStatusCheckRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DepotStatusCheckRequired
}

// GetDepotStatusCheckRequiredOk returns a tuple with the DepotStatusCheckRequired field value
// and a boolean to check if the value has been set.
func (o *ModifyMergeRequestBasicSettingsRequest) GetDepotStatusCheckRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotStatusCheckRequired, true
}

// SetDepotStatusCheckRequired sets field value
func (o *ModifyMergeRequestBasicSettingsRequest) SetDepotStatusCheckRequired(v bool) {
	o.DepotStatusCheckRequired = v
}

// GetMrCheckAllReviewersAllow returns the MrCheckAllReviewersAllow field value
func (o *ModifyMergeRequestBasicSettingsRequest) GetMrCheckAllReviewersAllow() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MrCheckAllReviewersAllow
}

// GetMrCheckAllReviewersAllowOk returns a tuple with the MrCheckAllReviewersAllow field value
// and a boolean to check if the value has been set.
func (o *ModifyMergeRequestBasicSettingsRequest) GetMrCheckAllReviewersAllowOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MrCheckAllReviewersAllow, true
}

// SetMrCheckAllReviewersAllow sets field value
func (o *ModifyMergeRequestBasicSettingsRequest) SetMrCheckAllReviewersAllow(v bool) {
	o.MrCheckAllReviewersAllow = v
}

// GetSquashOnMerge returns the SquashOnMerge field value if set, zero value otherwise.
func (o *ModifyMergeRequestBasicSettingsRequest) GetSquashOnMerge() string {
	if o == nil || utils.IsNil(o.SquashOnMerge) {
		var ret string
		return ret
	}
	return *o.SquashOnMerge
}

// GetSquashOnMergeOk returns a tuple with the SquashOnMerge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyMergeRequestBasicSettingsRequest) GetSquashOnMergeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SquashOnMerge) {
		return nil, false
	}
	return o.SquashOnMerge, true
}

// HasSquashOnMerge returns a boolean if a field has been set.
func (o *ModifyMergeRequestBasicSettingsRequest) HasSquashOnMerge() bool {
	if o != nil && !utils.IsNil(o.SquashOnMerge) {
		return true
	}

	return false
}

// SetSquashOnMerge gets a reference to the given string and assigns it to the SquashOnMerge field.
func (o *ModifyMergeRequestBasicSettingsRequest) SetSquashOnMerge(v string) {
	o.SquashOnMerge = &v
}

func (o ModifyMergeRequestBasicSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyMergeRequestBasicSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CancelMrGrantAfterPushSrc"] = o.CancelMrGrantAfterPushSrc
	toSerialize["DefaultFastForwardMerge"] = o.DefaultFastForwardMerge
	if !utils.IsNil(o.DefaultTargetBranch) {
		toSerialize["DefaultTargetBranch"] = o.DefaultTargetBranch
	}
	toSerialize["DeleteSrcBranchAfterMerged"] = o.DeleteSrcBranchAfterMerged
	toSerialize["DepotPath"] = o.DepotPath
	toSerialize["DepotStatusCheckRequired"] = o.DepotStatusCheckRequired
	toSerialize["MrCheckAllReviewersAllow"] = o.MrCheckAllReviewersAllow
	if !utils.IsNil(o.SquashOnMerge) {
		toSerialize["SquashOnMerge"] = o.SquashOnMerge
	}
	return toSerialize, nil
}

func (o *ModifyMergeRequestBasicSettingsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"CancelMrGrantAfterPushSrc",
		"DefaultFastForwardMerge",
		"DeleteSrcBranchAfterMerged",
		"DepotPath",
		"DepotStatusCheckRequired",
		"MrCheckAllReviewersAllow",
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

	varModifyMergeRequestBasicSettingsRequest := _ModifyMergeRequestBasicSettingsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyMergeRequestBasicSettingsRequest)

	if err != nil {
		return err
	}

	*o = ModifyMergeRequestBasicSettingsRequest(varModifyMergeRequestBasicSettingsRequest)

	return err
}

type NullableModifyMergeRequestBasicSettingsRequest struct {
	value *ModifyMergeRequestBasicSettingsRequest
	isSet bool
}

func (v NullableModifyMergeRequestBasicSettingsRequest) Get() *ModifyMergeRequestBasicSettingsRequest {
	return v.value
}

func (v *NullableModifyMergeRequestBasicSettingsRequest) Set(val *ModifyMergeRequestBasicSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyMergeRequestBasicSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyMergeRequestBasicSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyMergeRequestBasicSettingsRequest(val *ModifyMergeRequestBasicSettingsRequest) *NullableModifyMergeRequestBasicSettingsRequest {
	return &NullableModifyMergeRequestBasicSettingsRequest{value: val, isSet: true}
}

func (v NullableModifyMergeRequestBasicSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyMergeRequestBasicSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


