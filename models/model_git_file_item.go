/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the GitFileItem type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GitFileItem{}

// GitFileItem 文件详情
type GitFileItem struct {
	// 加密后文件的内容
	Content utils.NullableString `json:"Content,omitempty"`
	// 文件内容的hash结果
	ContentSha256 utils.NullableString `json:"ContentSha256,omitempty"`
	// 加密形式
	Encoding *string `json:"Encoding,omitempty"`
	// 文件名
	FileName *string `json:"FileName,omitempty"`
	// 文件路径
	FilePath *string `json:"FilePath,omitempty"`
	// commitID
	Sha *string `json:"Sha,omitempty"`
	// 文件大小
	Size *int64 `json:"Size,omitempty"`
}

// NewGitFileItem instantiates a new GitFileItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitFileItem() *GitFileItem {
	this := GitFileItem{}
	var content string = ""
	this.Content = *utils.NewNullableString(&content)
	var contentSha256 string = ""
	this.ContentSha256 = *utils.NewNullableString(&contentSha256)
	var encoding string = ""
	this.Encoding = &encoding
	var fileName string = ""
	this.FileName = &fileName
	var filePath string = ""
	this.FilePath = &filePath
	var sha string = ""
	this.Sha = &sha
	return &this
}

// NewGitFileItemWithDefaults instantiates a new GitFileItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitFileItemWithDefaults() *GitFileItem {
	this := GitFileItem{}
	var content string = ""
	this.Content = *utils.NewNullableString(&content)
	var contentSha256 string = ""
	this.ContentSha256 = *utils.NewNullableString(&contentSha256)
	var encoding string = ""
	this.Encoding = &encoding
	var fileName string = ""
	this.FileName = &fileName
	var filePath string = ""
	this.FilePath = &filePath
	var sha string = ""
	this.Sha = &sha
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitFileItem) GetContent() string {
	if o == nil || utils.IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitFileItem) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *GitFileItem) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given utils.NullableString and assigns it to the Content field.
func (o *GitFileItem) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *GitFileItem) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *GitFileItem) UnsetContent() {
	o.Content.Unset()
}

// GetContentSha256 returns the ContentSha256 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitFileItem) GetContentSha256() string {
	if o == nil || utils.IsNil(o.ContentSha256.Get()) {
		var ret string
		return ret
	}
	return *o.ContentSha256.Get()
}

// GetContentSha256Ok returns a tuple with the ContentSha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitFileItem) GetContentSha256Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentSha256.Get(), o.ContentSha256.IsSet()
}

// HasContentSha256 returns a boolean if a field has been set.
func (o *GitFileItem) HasContentSha256() bool {
	if o != nil && o.ContentSha256.IsSet() {
		return true
	}

	return false
}

// SetContentSha256 gets a reference to the given utils.NullableString and assigns it to the ContentSha256 field.
func (o *GitFileItem) SetContentSha256(v string) {
	o.ContentSha256.Set(&v)
}
// SetContentSha256Nil sets the value for ContentSha256 to be an explicit nil
func (o *GitFileItem) SetContentSha256Nil() {
	o.ContentSha256.Set(nil)
}

// UnsetContentSha256 ensures that no value is present for ContentSha256, not even an explicit nil
func (o *GitFileItem) UnsetContentSha256() {
	o.ContentSha256.Unset()
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *GitFileItem) GetEncoding() string {
	if o == nil || utils.IsNil(o.Encoding) {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitFileItem) GetEncodingOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Encoding) {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *GitFileItem) HasEncoding() bool {
	if o != nil && !utils.IsNil(o.Encoding) {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *GitFileItem) SetEncoding(v string) {
	o.Encoding = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *GitFileItem) GetFileName() string {
	if o == nil || utils.IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitFileItem) GetFileNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *GitFileItem) HasFileName() bool {
	if o != nil && !utils.IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *GitFileItem) SetFileName(v string) {
	o.FileName = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *GitFileItem) GetFilePath() string {
	if o == nil || utils.IsNil(o.FilePath) {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitFileItem) GetFilePathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FilePath) {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *GitFileItem) HasFilePath() bool {
	if o != nil && !utils.IsNil(o.FilePath) {
		return true
	}

	return false
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *GitFileItem) SetFilePath(v string) {
	o.FilePath = &v
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *GitFileItem) GetSha() string {
	if o == nil || utils.IsNil(o.Sha) {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitFileItem) GetShaOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Sha) {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *GitFileItem) HasSha() bool {
	if o != nil && !utils.IsNil(o.Sha) {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *GitFileItem) SetSha(v string) {
	o.Sha = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *GitFileItem) GetSize() int64 {
	if o == nil || utils.IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitFileItem) GetSizeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *GitFileItem) HasSize() bool {
	if o != nil && !utils.IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *GitFileItem) SetSize(v int64) {
	o.Size = &v
}

func (o GitFileItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitFileItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Content.IsSet() {
		toSerialize["Content"] = o.Content.Get()
	}
	if o.ContentSha256.IsSet() {
		toSerialize["ContentSha256"] = o.ContentSha256.Get()
	}
	if !utils.IsNil(o.Encoding) {
		toSerialize["Encoding"] = o.Encoding
	}
	if !utils.IsNil(o.FileName) {
		toSerialize["FileName"] = o.FileName
	}
	if !utils.IsNil(o.FilePath) {
		toSerialize["FilePath"] = o.FilePath
	}
	if !utils.IsNil(o.Sha) {
		toSerialize["Sha"] = o.Sha
	}
	if !utils.IsNil(o.Size) {
		toSerialize["Size"] = o.Size
	}
	return toSerialize, nil
}

type NullableGitFileItem struct {
	value *GitFileItem
	isSet bool
}

func (v NullableGitFileItem) Get() *GitFileItem {
	return v.value
}

func (v *NullableGitFileItem) Set(val *GitFileItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGitFileItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGitFileItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitFileItem(val *GitFileItem) *NullableGitFileItem {
	return &NullableGitFileItem{value: val, isSet: true}
}

func (v NullableGitFileItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitFileItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


