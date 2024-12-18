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

// checks if the ModifyGitCommitStatusRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyGitCommitStatusRequest{}

// ModifyGitCommitStatusRequest struct for ModifyGitCommitStatusRequest
type ModifyGitCommitStatusRequest struct {
	// 提交id
	CommitSha string `json:"CommitSha"`
	// 流水线文本内容
	Context string `json:"Context"`
	// 仓库id
	DepotId *int64 `json:"DepotId,omitempty"`
	// 仓库路径，DepotId与DepotPath二选一即可
	DepotPath *string `json:"DepotPath,omitempty"`
	// 流水线状态描述
	Description string `json:"Description"`
	// 提交状态
	State string `json:"State"`
	// 流水线链接地址
	TargetURL string `json:"TargetURL"`
}

type _ModifyGitCommitStatusRequest ModifyGitCommitStatusRequest

// NewModifyGitCommitStatusRequest instantiates a new ModifyGitCommitStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyGitCommitStatusRequest(commitSha string, context string, description string, state string, targetURL string) *ModifyGitCommitStatusRequest {
	this := ModifyGitCommitStatusRequest{}
	this.CommitSha = commitSha
	this.Context = context
	this.Description = description
	this.State = state
	this.TargetURL = targetURL
	return &this
}

// NewModifyGitCommitStatusRequestWithDefaults instantiates a new ModifyGitCommitStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyGitCommitStatusRequestWithDefaults() *ModifyGitCommitStatusRequest {
	this := ModifyGitCommitStatusRequest{}
	return &this
}

// GetCommitSha returns the CommitSha field value
func (o *ModifyGitCommitStatusRequest) GetCommitSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value
// and a boolean to check if the value has been set.
func (o *ModifyGitCommitStatusRequest) GetCommitShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitSha, true
}

// SetCommitSha sets field value
func (o *ModifyGitCommitStatusRequest) SetCommitSha(v string) {
	o.CommitSha = v
}

// GetContext returns the Context field value
func (o *ModifyGitCommitStatusRequest) GetContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *ModifyGitCommitStatusRequest) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *ModifyGitCommitStatusRequest) SetContext(v string) {
	o.Context = v
}

// GetDepotId returns the DepotId field value if set, zero value otherwise.
func (o *ModifyGitCommitStatusRequest) GetDepotId() int64 {
	if o == nil || utils.IsNil(o.DepotId) {
		var ret int64
		return ret
	}
	return *o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyGitCommitStatusRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.DepotId) {
		return nil, false
	}
	return o.DepotId, true
}

// HasDepotId returns a boolean if a field has been set.
func (o *ModifyGitCommitStatusRequest) HasDepotId() bool {
	if o != nil && !utils.IsNil(o.DepotId) {
		return true
	}

	return false
}

// SetDepotId gets a reference to the given int64 and assigns it to the DepotId field.
func (o *ModifyGitCommitStatusRequest) SetDepotId(v int64) {
	o.DepotId = &v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *ModifyGitCommitStatusRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyGitCommitStatusRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *ModifyGitCommitStatusRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *ModifyGitCommitStatusRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

// GetDescription returns the Description field value
func (o *ModifyGitCommitStatusRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ModifyGitCommitStatusRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ModifyGitCommitStatusRequest) SetDescription(v string) {
	o.Description = v
}

// GetState returns the State field value
func (o *ModifyGitCommitStatusRequest) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ModifyGitCommitStatusRequest) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ModifyGitCommitStatusRequest) SetState(v string) {
	o.State = v
}

// GetTargetURL returns the TargetURL field value
func (o *ModifyGitCommitStatusRequest) GetTargetURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetURL
}

// GetTargetURLOk returns a tuple with the TargetURL field value
// and a boolean to check if the value has been set.
func (o *ModifyGitCommitStatusRequest) GetTargetURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetURL, true
}

// SetTargetURL sets field value
func (o *ModifyGitCommitStatusRequest) SetTargetURL(v string) {
	o.TargetURL = v
}

func (o ModifyGitCommitStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyGitCommitStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CommitSha"] = o.CommitSha
	toSerialize["Context"] = o.Context
	if !utils.IsNil(o.DepotId) {
		toSerialize["DepotId"] = o.DepotId
	}
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	toSerialize["Description"] = o.Description
	toSerialize["State"] = o.State
	toSerialize["TargetURL"] = o.TargetURL
	return toSerialize, nil
}

func (o *ModifyGitCommitStatusRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"CommitSha",
		"Context",
		"Description",
		"State",
		"TargetURL",
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

	varModifyGitCommitStatusRequest := _ModifyGitCommitStatusRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyGitCommitStatusRequest)

	if err != nil {
		return err
	}

	*o = ModifyGitCommitStatusRequest(varModifyGitCommitStatusRequest)

	return err
}

type NullableModifyGitCommitStatusRequest struct {
	value *ModifyGitCommitStatusRequest
	isSet bool
}

func (v NullableModifyGitCommitStatusRequest) Get() *ModifyGitCommitStatusRequest {
	return v.value
}

func (v *NullableModifyGitCommitStatusRequest) Set(val *ModifyGitCommitStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyGitCommitStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyGitCommitStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyGitCommitStatusRequest(val *ModifyGitCommitStatusRequest) *NullableModifyGitCommitStatusRequest {
	return &NullableModifyGitCommitStatusRequest{value: val, isSet: true}
}

func (v NullableModifyGitCommitStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyGitCommitStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


