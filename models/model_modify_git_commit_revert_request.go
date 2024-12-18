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

// checks if the ModifyGitCommitRevertRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyGitCommitRevertRequest{}

// ModifyGitCommitRevertRequest struct for ModifyGitCommitRevertRequest
type ModifyGitCommitRevertRequest struct {
	// 分支名
	BranchName string `json:"BranchName"`
	// 仓库 ID
	DepotId int64 `json:"DepotId"`
	// 仓库路径，DepotId与DepotPath二选一即可
	DepotPath *string `json:"DepotPath,omitempty"`
	// 提交描述
	Message string `json:"Message"`
	// 欲还原的提交 ID
	Sha string `json:"Sha"`
}

type _ModifyGitCommitRevertRequest ModifyGitCommitRevertRequest

// NewModifyGitCommitRevertRequest instantiates a new ModifyGitCommitRevertRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyGitCommitRevertRequest(branchName string, depotId int64, message string, sha string) *ModifyGitCommitRevertRequest {
	this := ModifyGitCommitRevertRequest{}
	this.BranchName = branchName
	this.DepotId = depotId
	this.Message = message
	this.Sha = sha
	return &this
}

// NewModifyGitCommitRevertRequestWithDefaults instantiates a new ModifyGitCommitRevertRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyGitCommitRevertRequestWithDefaults() *ModifyGitCommitRevertRequest {
	this := ModifyGitCommitRevertRequest{}
	return &this
}

// GetBranchName returns the BranchName field value
func (o *ModifyGitCommitRevertRequest) GetBranchName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value
// and a boolean to check if the value has been set.
func (o *ModifyGitCommitRevertRequest) GetBranchNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchName, true
}

// SetBranchName sets field value
func (o *ModifyGitCommitRevertRequest) SetBranchName(v string) {
	o.BranchName = v
}

// GetDepotId returns the DepotId field value
func (o *ModifyGitCommitRevertRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *ModifyGitCommitRevertRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *ModifyGitCommitRevertRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *ModifyGitCommitRevertRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyGitCommitRevertRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *ModifyGitCommitRevertRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *ModifyGitCommitRevertRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

// GetMessage returns the Message field value
func (o *ModifyGitCommitRevertRequest) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ModifyGitCommitRevertRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ModifyGitCommitRevertRequest) SetMessage(v string) {
	o.Message = v
}

// GetSha returns the Sha field value
func (o *ModifyGitCommitRevertRequest) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *ModifyGitCommitRevertRequest) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value
func (o *ModifyGitCommitRevertRequest) SetSha(v string) {
	o.Sha = v
}

func (o ModifyGitCommitRevertRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyGitCommitRevertRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["BranchName"] = o.BranchName
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	toSerialize["Message"] = o.Message
	toSerialize["Sha"] = o.Sha
	return toSerialize, nil
}

func (o *ModifyGitCommitRevertRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"BranchName",
		"DepotId",
		"Message",
		"Sha",
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

	varModifyGitCommitRevertRequest := _ModifyGitCommitRevertRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyGitCommitRevertRequest)

	if err != nil {
		return err
	}

	*o = ModifyGitCommitRevertRequest(varModifyGitCommitRevertRequest)

	return err
}

type NullableModifyGitCommitRevertRequest struct {
	value *ModifyGitCommitRevertRequest
	isSet bool
}

func (v NullableModifyGitCommitRevertRequest) Get() *ModifyGitCommitRevertRequest {
	return v.value
}

func (v *NullableModifyGitCommitRevertRequest) Set(val *ModifyGitCommitRevertRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyGitCommitRevertRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyGitCommitRevertRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyGitCommitRevertRequest(val *ModifyGitCommitRevertRequest) *NullableModifyGitCommitRevertRequest {
	return &NullableModifyGitCommitRevertRequest{value: val, isSet: true}
}

func (v NullableModifyGitCommitRevertRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyGitCommitRevertRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


