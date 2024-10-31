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

// checks if the ModifyGitReleaseRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyGitReleaseRequest{}

// ModifyGitReleaseRequest struct for ModifyGitReleaseRequest
type ModifyGitReleaseRequest struct {
	// 仓库id
	DepotId int64 `json:"DepotId"`
	// 仓库路径
	DepotPath *string `json:"DepotPath,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty"`
	// 是否预发布
	Pre *bool `json:"Pre,omitempty"`
	// 项目下仓库版本唯一标识符
	ReleaseId int64 `json:"ReleaseId"`
	// 标签名称
	TagName string `json:"TagName"`
	// 标题
	Title *string `json:"Title,omitempty"`
}

type _ModifyGitReleaseRequest ModifyGitReleaseRequest

// NewModifyGitReleaseRequest instantiates a new ModifyGitReleaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyGitReleaseRequest(depotId int64, releaseId int64, tagName string) *ModifyGitReleaseRequest {
	this := ModifyGitReleaseRequest{}
	this.DepotId = depotId
	this.ReleaseId = releaseId
	this.TagName = tagName
	return &this
}

// NewModifyGitReleaseRequestWithDefaults instantiates a new ModifyGitReleaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyGitReleaseRequestWithDefaults() *ModifyGitReleaseRequest {
	this := ModifyGitReleaseRequest{}
	return &this
}

// GetDepotId returns the DepotId field value
func (o *ModifyGitReleaseRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *ModifyGitReleaseRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *ModifyGitReleaseRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *ModifyGitReleaseRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyGitReleaseRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *ModifyGitReleaseRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *ModifyGitReleaseRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModifyGitReleaseRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyGitReleaseRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModifyGitReleaseRequest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModifyGitReleaseRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPre returns the Pre field value if set, zero value otherwise.
func (o *ModifyGitReleaseRequest) GetPre() bool {
	if o == nil || utils.IsNil(o.Pre) {
		var ret bool
		return ret
	}
	return *o.Pre
}

// GetPreOk returns a tuple with the Pre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyGitReleaseRequest) GetPreOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Pre) {
		return nil, false
	}
	return o.Pre, true
}

// HasPre returns a boolean if a field has been set.
func (o *ModifyGitReleaseRequest) HasPre() bool {
	if o != nil && !utils.IsNil(o.Pre) {
		return true
	}

	return false
}

// SetPre gets a reference to the given bool and assigns it to the Pre field.
func (o *ModifyGitReleaseRequest) SetPre(v bool) {
	o.Pre = &v
}

// GetReleaseId returns the ReleaseId field value
func (o *ModifyGitReleaseRequest) GetReleaseId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ReleaseId
}

// GetReleaseIdOk returns a tuple with the ReleaseId field value
// and a boolean to check if the value has been set.
func (o *ModifyGitReleaseRequest) GetReleaseIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseId, true
}

// SetReleaseId sets field value
func (o *ModifyGitReleaseRequest) SetReleaseId(v int64) {
	o.ReleaseId = v
}

// GetTagName returns the TagName field value
func (o *ModifyGitReleaseRequest) GetTagName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TagName
}

// GetTagNameOk returns a tuple with the TagName field value
// and a boolean to check if the value has been set.
func (o *ModifyGitReleaseRequest) GetTagNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagName, true
}

// SetTagName sets field value
func (o *ModifyGitReleaseRequest) SetTagName(v string) {
	o.TagName = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ModifyGitReleaseRequest) GetTitle() string {
	if o == nil || utils.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyGitReleaseRequest) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ModifyGitReleaseRequest) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ModifyGitReleaseRequest) SetTitle(v string) {
	o.Title = &v
}

func (o ModifyGitReleaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyGitReleaseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	if !utils.IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !utils.IsNil(o.Pre) {
		toSerialize["Pre"] = o.Pre
	}
	toSerialize["ReleaseId"] = o.ReleaseId
	toSerialize["TagName"] = o.TagName
	if !utils.IsNil(o.Title) {
		toSerialize["Title"] = o.Title
	}
	return toSerialize, nil
}

func (o *ModifyGitReleaseRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepotId",
		"ReleaseId",
		"TagName",
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

	varModifyGitReleaseRequest := _ModifyGitReleaseRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyGitReleaseRequest)

	if err != nil {
		return err
	}

	*o = ModifyGitReleaseRequest(varModifyGitReleaseRequest)

	return err
}

type NullableModifyGitReleaseRequest struct {
	value *ModifyGitReleaseRequest
	isSet bool
}

func (v NullableModifyGitReleaseRequest) Get() *ModifyGitReleaseRequest {
	return v.value
}

func (v *NullableModifyGitReleaseRequest) Set(val *ModifyGitReleaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyGitReleaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyGitReleaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyGitReleaseRequest(val *ModifyGitReleaseRequest) *NullableModifyGitReleaseRequest {
	return &NullableModifyGitReleaseRequest{value: val, isSet: true}
}

func (v NullableModifyGitReleaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyGitReleaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

