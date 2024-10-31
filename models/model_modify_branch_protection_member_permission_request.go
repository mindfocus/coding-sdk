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

// checks if the ModifyBranchProtectionMemberPermissionRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyBranchProtectionMemberPermissionRequest{}

// ModifyBranchProtectionMemberPermissionRequest struct for ModifyBranchProtectionMemberPermissionRequest
type ModifyBranchProtectionMemberPermissionRequest struct {
	// 是否直接推送
	AllowPush bool `json:"AllowPush"`
	// 分支规则名称
	BranchRuleName string `json:"BranchRuleName"`
	// 仓库路径
	DepotPath string `json:"DepotPath"`
	// 用户GlobalKey
	UserGlobalKey string `json:"UserGlobalKey"`
}

type _ModifyBranchProtectionMemberPermissionRequest ModifyBranchProtectionMemberPermissionRequest

// NewModifyBranchProtectionMemberPermissionRequest instantiates a new ModifyBranchProtectionMemberPermissionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyBranchProtectionMemberPermissionRequest(allowPush bool, branchRuleName string, depotPath string, userGlobalKey string) *ModifyBranchProtectionMemberPermissionRequest {
	this := ModifyBranchProtectionMemberPermissionRequest{}
	this.AllowPush = allowPush
	this.BranchRuleName = branchRuleName
	this.DepotPath = depotPath
	this.UserGlobalKey = userGlobalKey
	return &this
}

// NewModifyBranchProtectionMemberPermissionRequestWithDefaults instantiates a new ModifyBranchProtectionMemberPermissionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyBranchProtectionMemberPermissionRequestWithDefaults() *ModifyBranchProtectionMemberPermissionRequest {
	this := ModifyBranchProtectionMemberPermissionRequest{}
	return &this
}

// GetAllowPush returns the AllowPush field value
func (o *ModifyBranchProtectionMemberPermissionRequest) GetAllowPush() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowPush
}

// GetAllowPushOk returns a tuple with the AllowPush field value
// and a boolean to check if the value has been set.
func (o *ModifyBranchProtectionMemberPermissionRequest) GetAllowPushOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowPush, true
}

// SetAllowPush sets field value
func (o *ModifyBranchProtectionMemberPermissionRequest) SetAllowPush(v bool) {
	o.AllowPush = v
}

// GetBranchRuleName returns the BranchRuleName field value
func (o *ModifyBranchProtectionMemberPermissionRequest) GetBranchRuleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BranchRuleName
}

// GetBranchRuleNameOk returns a tuple with the BranchRuleName field value
// and a boolean to check if the value has been set.
func (o *ModifyBranchProtectionMemberPermissionRequest) GetBranchRuleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchRuleName, true
}

// SetBranchRuleName sets field value
func (o *ModifyBranchProtectionMemberPermissionRequest) SetBranchRuleName(v string) {
	o.BranchRuleName = v
}

// GetDepotPath returns the DepotPath field value
func (o *ModifyBranchProtectionMemberPermissionRequest) GetDepotPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value
// and a boolean to check if the value has been set.
func (o *ModifyBranchProtectionMemberPermissionRequest) GetDepotPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotPath, true
}

// SetDepotPath sets field value
func (o *ModifyBranchProtectionMemberPermissionRequest) SetDepotPath(v string) {
	o.DepotPath = v
}

// GetUserGlobalKey returns the UserGlobalKey field value
func (o *ModifyBranchProtectionMemberPermissionRequest) GetUserGlobalKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserGlobalKey
}

// GetUserGlobalKeyOk returns a tuple with the UserGlobalKey field value
// and a boolean to check if the value has been set.
func (o *ModifyBranchProtectionMemberPermissionRequest) GetUserGlobalKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserGlobalKey, true
}

// SetUserGlobalKey sets field value
func (o *ModifyBranchProtectionMemberPermissionRequest) SetUserGlobalKey(v string) {
	o.UserGlobalKey = v
}

func (o ModifyBranchProtectionMemberPermissionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyBranchProtectionMemberPermissionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AllowPush"] = o.AllowPush
	toSerialize["BranchRuleName"] = o.BranchRuleName
	toSerialize["DepotPath"] = o.DepotPath
	toSerialize["UserGlobalKey"] = o.UserGlobalKey
	return toSerialize, nil
}

func (o *ModifyBranchProtectionMemberPermissionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"AllowPush",
		"BranchRuleName",
		"DepotPath",
		"UserGlobalKey",
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

	varModifyBranchProtectionMemberPermissionRequest := _ModifyBranchProtectionMemberPermissionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyBranchProtectionMemberPermissionRequest)

	if err != nil {
		return err
	}

	*o = ModifyBranchProtectionMemberPermissionRequest(varModifyBranchProtectionMemberPermissionRequest)

	return err
}

type NullableModifyBranchProtectionMemberPermissionRequest struct {
	value *ModifyBranchProtectionMemberPermissionRequest
	isSet bool
}

func (v NullableModifyBranchProtectionMemberPermissionRequest) Get() *ModifyBranchProtectionMemberPermissionRequest {
	return v.value
}

func (v *NullableModifyBranchProtectionMemberPermissionRequest) Set(val *ModifyBranchProtectionMemberPermissionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyBranchProtectionMemberPermissionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyBranchProtectionMemberPermissionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyBranchProtectionMemberPermissionRequest(val *ModifyBranchProtectionMemberPermissionRequest) *NullableModifyBranchProtectionMemberPermissionRequest {
	return &NullableModifyBranchProtectionMemberPermissionRequest{value: val, isSet: true}
}

func (v NullableModifyBranchProtectionMemberPermissionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyBranchProtectionMemberPermissionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

