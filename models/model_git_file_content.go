/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the GitFileContent type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GitFileContent{}

// GitFileContent 文件详情
type GitFileContent struct {
	// 内容
	Content *string `json:"Content,omitempty"`
	// 是否为lfs文件
	IsLargeFileStorage *bool `json:"IsLargeFileStorage,omitempty"`
	// 是否lfs文件
	IsLfs *bool `json:"IsLfs,omitempty"`
	// 是否文本
	IsText *bool `json:"IsText,omitempty"`
}

// NewGitFileContent instantiates a new GitFileContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitFileContent() *GitFileContent {
	this := GitFileContent{}
	var content string = ""
	this.Content = &content
	var isLargeFileStorage bool = false
	this.IsLargeFileStorage = &isLargeFileStorage
	var isLfs bool = false
	this.IsLfs = &isLfs
	var isText bool = false
	this.IsText = &isText
	return &this
}

// NewGitFileContentWithDefaults instantiates a new GitFileContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitFileContentWithDefaults() *GitFileContent {
	this := GitFileContent{}
	var content string = ""
	this.Content = &content
	var isLargeFileStorage bool = false
	this.IsLargeFileStorage = &isLargeFileStorage
	var isLfs bool = false
	this.IsLfs = &isLfs
	var isText bool = false
	this.IsText = &isText
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *GitFileContent) GetContent() string {
	if o == nil || utils.IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitFileContent) GetContentOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *GitFileContent) HasContent() bool {
	if o != nil && !utils.IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *GitFileContent) SetContent(v string) {
	o.Content = &v
}

// GetIsLargeFileStorage returns the IsLargeFileStorage field value if set, zero value otherwise.
func (o *GitFileContent) GetIsLargeFileStorage() bool {
	if o == nil || utils.IsNil(o.IsLargeFileStorage) {
		var ret bool
		return ret
	}
	return *o.IsLargeFileStorage
}

// GetIsLargeFileStorageOk returns a tuple with the IsLargeFileStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitFileContent) GetIsLargeFileStorageOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsLargeFileStorage) {
		return nil, false
	}
	return o.IsLargeFileStorage, true
}

// HasIsLargeFileStorage returns a boolean if a field has been set.
func (o *GitFileContent) HasIsLargeFileStorage() bool {
	if o != nil && !utils.IsNil(o.IsLargeFileStorage) {
		return true
	}

	return false
}

// SetIsLargeFileStorage gets a reference to the given bool and assigns it to the IsLargeFileStorage field.
func (o *GitFileContent) SetIsLargeFileStorage(v bool) {
	o.IsLargeFileStorage = &v
}

// GetIsLfs returns the IsLfs field value if set, zero value otherwise.
func (o *GitFileContent) GetIsLfs() bool {
	if o == nil || utils.IsNil(o.IsLfs) {
		var ret bool
		return ret
	}
	return *o.IsLfs
}

// GetIsLfsOk returns a tuple with the IsLfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitFileContent) GetIsLfsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsLfs) {
		return nil, false
	}
	return o.IsLfs, true
}

// HasIsLfs returns a boolean if a field has been set.
func (o *GitFileContent) HasIsLfs() bool {
	if o != nil && !utils.IsNil(o.IsLfs) {
		return true
	}

	return false
}

// SetIsLfs gets a reference to the given bool and assigns it to the IsLfs field.
func (o *GitFileContent) SetIsLfs(v bool) {
	o.IsLfs = &v
}

// GetIsText returns the IsText field value if set, zero value otherwise.
func (o *GitFileContent) GetIsText() bool {
	if o == nil || utils.IsNil(o.IsText) {
		var ret bool
		return ret
	}
	return *o.IsText
}

// GetIsTextOk returns a tuple with the IsText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitFileContent) GetIsTextOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsText) {
		return nil, false
	}
	return o.IsText, true
}

// HasIsText returns a boolean if a field has been set.
func (o *GitFileContent) HasIsText() bool {
	if o != nil && !utils.IsNil(o.IsText) {
		return true
	}

	return false
}

// SetIsText gets a reference to the given bool and assigns it to the IsText field.
func (o *GitFileContent) SetIsText(v bool) {
	o.IsText = &v
}

func (o GitFileContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitFileContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Content) {
		toSerialize["Content"] = o.Content
	}
	if !utils.IsNil(o.IsLargeFileStorage) {
		toSerialize["IsLargeFileStorage"] = o.IsLargeFileStorage
	}
	if !utils.IsNil(o.IsLfs) {
		toSerialize["IsLfs"] = o.IsLfs
	}
	if !utils.IsNil(o.IsText) {
		toSerialize["IsText"] = o.IsText
	}
	return toSerialize, nil
}

type NullableGitFileContent struct {
	value *GitFileContent
	isSet bool
}

func (v NullableGitFileContent) Get() *GitFileContent {
	return v.value
}

func (v *NullableGitFileContent) Set(val *GitFileContent) {
	v.value = val
	v.isSet = true
}

func (v NullableGitFileContent) IsSet() bool {
	return v.isSet
}

func (v *NullableGitFileContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitFileContent(val *GitFileContent) *NullableGitFileContent {
	return &NullableGitFileContent{value: val, isSet: true}
}

func (v NullableGitFileContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitFileContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

