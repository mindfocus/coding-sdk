/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse{}

// ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse 公共返回体
type ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse struct {
	// git 仓库文件推送规则列表
	GitFilePushRule []GitFilePushRule `json:"GitFilePushRule,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewModifyDepotFilePushRuleDenyPrivilege200ResponseResponse instantiates a new ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyDepotFilePushRuleDenyPrivilege200ResponseResponse() *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse {
	this := ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewModifyDepotFilePushRuleDenyPrivilege200ResponseResponseWithDefaults instantiates a new ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyDepotFilePushRuleDenyPrivilege200ResponseResponseWithDefaults() *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse {
	this := ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetGitFilePushRule returns the GitFilePushRule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) GetGitFilePushRule() []GitFilePushRule {
	if o == nil {
		var ret []GitFilePushRule
		return ret
	}
	return o.GitFilePushRule
}

// GetGitFilePushRuleOk returns a tuple with the GitFilePushRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) GetGitFilePushRuleOk() ([]GitFilePushRule, bool) {
	if o == nil || utils.IsNil(o.GitFilePushRule) {
		return nil, false
	}
	return o.GitFilePushRule, true
}

// HasGitFilePushRule returns a boolean if a field has been set.
func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) HasGitFilePushRule() bool {
	if o != nil && !utils.IsNil(o.GitFilePushRule) {
		return true
	}

	return false
}

// SetGitFilePushRule gets a reference to the given []GitFilePushRule and assigns it to the GitFilePushRule field.
func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) SetGitFilePushRule(v []GitFilePushRule) {
	o.GitFilePushRule = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GitFilePushRule != nil {
		toSerialize["GitFilePushRule"] = o.GitFilePushRule
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableModifyDepotFilePushRuleDenyPrivilege200ResponseResponse struct {
	value *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse
	isSet bool
}

func (v NullableModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) Get() *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse {
	return v.value
}

func (v *NullableModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) Set(val *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyDepotFilePushRuleDenyPrivilege200ResponseResponse(val *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) *NullableModifyDepotFilePushRuleDenyPrivilege200ResponseResponse {
	return &NullableModifyDepotFilePushRuleDenyPrivilege200ResponseResponse{value: val, isSet: true}
}

func (v NullableModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

