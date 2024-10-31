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

// checks if the ModifyGitMergeRequestRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyGitMergeRequestRequest{}

// ModifyGitMergeRequestRequest struct for ModifyGitMergeRequestRequest
type ModifyGitMergeRequestRequest struct {
	// 待修改的合并请求描述
	Content *string `json:"Content,omitempty"`
	// 仓库 ID
	DepotId int64 `json:"DepotId"`
	// 仓库路径，与仓库ID二选一
	DepotPath *string `json:"DepotPath,omitempty"`
	// 合并请求 IID
	MergeId int64 `json:"MergeId"`
	// 待修改的合并请求标题
	Title *string `json:"Title,omitempty"`
}

type _ModifyGitMergeRequestRequest ModifyGitMergeRequestRequest

// NewModifyGitMergeRequestRequest instantiates a new ModifyGitMergeRequestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyGitMergeRequestRequest(depotId int64, mergeId int64) *ModifyGitMergeRequestRequest {
	this := ModifyGitMergeRequestRequest{}
	this.DepotId = depotId
	this.MergeId = mergeId
	return &this
}

// NewModifyGitMergeRequestRequestWithDefaults instantiates a new ModifyGitMergeRequestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyGitMergeRequestRequestWithDefaults() *ModifyGitMergeRequestRequest {
	this := ModifyGitMergeRequestRequest{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ModifyGitMergeRequestRequest) GetContent() string {
	if o == nil || utils.IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyGitMergeRequestRequest) GetContentOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ModifyGitMergeRequestRequest) HasContent() bool {
	if o != nil && !utils.IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ModifyGitMergeRequestRequest) SetContent(v string) {
	o.Content = &v
}

// GetDepotId returns the DepotId field value
func (o *ModifyGitMergeRequestRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *ModifyGitMergeRequestRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *ModifyGitMergeRequestRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *ModifyGitMergeRequestRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyGitMergeRequestRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *ModifyGitMergeRequestRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *ModifyGitMergeRequestRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

// GetMergeId returns the MergeId field value
func (o *ModifyGitMergeRequestRequest) GetMergeId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MergeId
}

// GetMergeIdOk returns a tuple with the MergeId field value
// and a boolean to check if the value has been set.
func (o *ModifyGitMergeRequestRequest) GetMergeIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MergeId, true
}

// SetMergeId sets field value
func (o *ModifyGitMergeRequestRequest) SetMergeId(v int64) {
	o.MergeId = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ModifyGitMergeRequestRequest) GetTitle() string {
	if o == nil || utils.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyGitMergeRequestRequest) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ModifyGitMergeRequestRequest) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ModifyGitMergeRequestRequest) SetTitle(v string) {
	o.Title = &v
}

func (o ModifyGitMergeRequestRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyGitMergeRequestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Content) {
		toSerialize["Content"] = o.Content
	}
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	toSerialize["MergeId"] = o.MergeId
	if !utils.IsNil(o.Title) {
		toSerialize["Title"] = o.Title
	}
	return toSerialize, nil
}

func (o *ModifyGitMergeRequestRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepotId",
		"MergeId",
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

	varModifyGitMergeRequestRequest := _ModifyGitMergeRequestRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyGitMergeRequestRequest)

	if err != nil {
		return err
	}

	*o = ModifyGitMergeRequestRequest(varModifyGitMergeRequestRequest)

	return err
}

type NullableModifyGitMergeRequestRequest struct {
	value *ModifyGitMergeRequestRequest
	isSet bool
}

func (v NullableModifyGitMergeRequestRequest) Get() *ModifyGitMergeRequestRequest {
	return v.value
}

func (v *NullableModifyGitMergeRequestRequest) Set(val *ModifyGitMergeRequestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyGitMergeRequestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyGitMergeRequestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyGitMergeRequestRequest(val *ModifyGitMergeRequestRequest) *NullableModifyGitMergeRequestRequest {
	return &NullableModifyGitMergeRequestRequest{value: val, isSet: true}
}

func (v NullableModifyGitMergeRequestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyGitMergeRequestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


