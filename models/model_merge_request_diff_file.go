/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the MergeRequestDiffFile type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MergeRequestDiffFile{}

// MergeRequestDiffFile diff 信息详情
type MergeRequestDiffFile struct {
	// 文件对应的 blob Id
	BlobSha utils.NullableString `json:"BlobSha,omitempty"`
	// 文件改变类型
	ChangeType utils.NullableString `json:"ChangeType,omitempty"`
	// 一共删除几行
	Deletions utils.NullableInt64 `json:"Deletions,omitempty"`
	// 一共新增几行
	Insertions utils.NullableInt64 `json:"Insertions,omitempty"`
	// 文件路径
	Path utils.NullableString `json:"Path,omitempty"`
	// 文件大小（字节）
	Size utils.NullableInt64 `json:"Size,omitempty"`
}

// NewMergeRequestDiffFile instantiates a new MergeRequestDiffFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeRequestDiffFile() *MergeRequestDiffFile {
	this := MergeRequestDiffFile{}
	var blobSha string = ""
	this.BlobSha = *utils.NewNullableString(&blobSha)
	var changeType string = ""
	this.ChangeType = *utils.NewNullableString(&changeType)
	var path string = ""
	this.Path = *utils.NewNullableString(&path)
	return &this
}

// NewMergeRequestDiffFileWithDefaults instantiates a new MergeRequestDiffFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeRequestDiffFileWithDefaults() *MergeRequestDiffFile {
	this := MergeRequestDiffFile{}
	var blobSha string = ""
	this.BlobSha = *utils.NewNullableString(&blobSha)
	var changeType string = ""
	this.ChangeType = *utils.NewNullableString(&changeType)
	var path string = ""
	this.Path = *utils.NewNullableString(&path)
	return &this
}

// GetBlobSha returns the BlobSha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDiffFile) GetBlobSha() string {
	if o == nil || utils.IsNil(o.BlobSha.Get()) {
		var ret string
		return ret
	}
	return *o.BlobSha.Get()
}

// GetBlobShaOk returns a tuple with the BlobSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDiffFile) GetBlobShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlobSha.Get(), o.BlobSha.IsSet()
}

// HasBlobSha returns a boolean if a field has been set.
func (o *MergeRequestDiffFile) HasBlobSha() bool {
	if o != nil && o.BlobSha.IsSet() {
		return true
	}

	return false
}

// SetBlobSha gets a reference to the given utils.NullableString and assigns it to the BlobSha field.
func (o *MergeRequestDiffFile) SetBlobSha(v string) {
	o.BlobSha.Set(&v)
}
// SetBlobShaNil sets the value for BlobSha to be an explicit nil
func (o *MergeRequestDiffFile) SetBlobShaNil() {
	o.BlobSha.Set(nil)
}

// UnsetBlobSha ensures that no value is present for BlobSha, not even an explicit nil
func (o *MergeRequestDiffFile) UnsetBlobSha() {
	o.BlobSha.Unset()
}

// GetChangeType returns the ChangeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDiffFile) GetChangeType() string {
	if o == nil || utils.IsNil(o.ChangeType.Get()) {
		var ret string
		return ret
	}
	return *o.ChangeType.Get()
}

// GetChangeTypeOk returns a tuple with the ChangeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDiffFile) GetChangeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChangeType.Get(), o.ChangeType.IsSet()
}

// HasChangeType returns a boolean if a field has been set.
func (o *MergeRequestDiffFile) HasChangeType() bool {
	if o != nil && o.ChangeType.IsSet() {
		return true
	}

	return false
}

// SetChangeType gets a reference to the given utils.NullableString and assigns it to the ChangeType field.
func (o *MergeRequestDiffFile) SetChangeType(v string) {
	o.ChangeType.Set(&v)
}
// SetChangeTypeNil sets the value for ChangeType to be an explicit nil
func (o *MergeRequestDiffFile) SetChangeTypeNil() {
	o.ChangeType.Set(nil)
}

// UnsetChangeType ensures that no value is present for ChangeType, not even an explicit nil
func (o *MergeRequestDiffFile) UnsetChangeType() {
	o.ChangeType.Unset()
}

// GetDeletions returns the Deletions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDiffFile) GetDeletions() int64 {
	if o == nil || utils.IsNil(o.Deletions.Get()) {
		var ret int64
		return ret
	}
	return *o.Deletions.Get()
}

// GetDeletionsOk returns a tuple with the Deletions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDiffFile) GetDeletionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deletions.Get(), o.Deletions.IsSet()
}

// HasDeletions returns a boolean if a field has been set.
func (o *MergeRequestDiffFile) HasDeletions() bool {
	if o != nil && o.Deletions.IsSet() {
		return true
	}

	return false
}

// SetDeletions gets a reference to the given utils.NullableInt64 and assigns it to the Deletions field.
func (o *MergeRequestDiffFile) SetDeletions(v int64) {
	o.Deletions.Set(&v)
}
// SetDeletionsNil sets the value for Deletions to be an explicit nil
func (o *MergeRequestDiffFile) SetDeletionsNil() {
	o.Deletions.Set(nil)
}

// UnsetDeletions ensures that no value is present for Deletions, not even an explicit nil
func (o *MergeRequestDiffFile) UnsetDeletions() {
	o.Deletions.Unset()
}

// GetInsertions returns the Insertions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDiffFile) GetInsertions() int64 {
	if o == nil || utils.IsNil(o.Insertions.Get()) {
		var ret int64
		return ret
	}
	return *o.Insertions.Get()
}

// GetInsertionsOk returns a tuple with the Insertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDiffFile) GetInsertionsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Insertions.Get(), o.Insertions.IsSet()
}

// HasInsertions returns a boolean if a field has been set.
func (o *MergeRequestDiffFile) HasInsertions() bool {
	if o != nil && o.Insertions.IsSet() {
		return true
	}

	return false
}

// SetInsertions gets a reference to the given utils.NullableInt64 and assigns it to the Insertions field.
func (o *MergeRequestDiffFile) SetInsertions(v int64) {
	o.Insertions.Set(&v)
}
// SetInsertionsNil sets the value for Insertions to be an explicit nil
func (o *MergeRequestDiffFile) SetInsertionsNil() {
	o.Insertions.Set(nil)
}

// UnsetInsertions ensures that no value is present for Insertions, not even an explicit nil
func (o *MergeRequestDiffFile) UnsetInsertions() {
	o.Insertions.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDiffFile) GetPath() string {
	if o == nil || utils.IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDiffFile) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *MergeRequestDiffFile) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given utils.NullableString and assigns it to the Path field.
func (o *MergeRequestDiffFile) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *MergeRequestDiffFile) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *MergeRequestDiffFile) UnsetPath() {
	o.Path.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDiffFile) GetSize() int64 {
	if o == nil || utils.IsNil(o.Size.Get()) {
		var ret int64
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDiffFile) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *MergeRequestDiffFile) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given utils.NullableInt64 and assigns it to the Size field.
func (o *MergeRequestDiffFile) SetSize(v int64) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *MergeRequestDiffFile) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *MergeRequestDiffFile) UnsetSize() {
	o.Size.Unset()
}

func (o MergeRequestDiffFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MergeRequestDiffFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BlobSha.IsSet() {
		toSerialize["BlobSha"] = o.BlobSha.Get()
	}
	if o.ChangeType.IsSet() {
		toSerialize["ChangeType"] = o.ChangeType.Get()
	}
	if o.Deletions.IsSet() {
		toSerialize["Deletions"] = o.Deletions.Get()
	}
	if o.Insertions.IsSet() {
		toSerialize["Insertions"] = o.Insertions.Get()
	}
	if o.Path.IsSet() {
		toSerialize["Path"] = o.Path.Get()
	}
	if o.Size.IsSet() {
		toSerialize["Size"] = o.Size.Get()
	}
	return toSerialize, nil
}

type NullableMergeRequestDiffFile struct {
	value *MergeRequestDiffFile
	isSet bool
}

func (v NullableMergeRequestDiffFile) Get() *MergeRequestDiffFile {
	return v.value
}

func (v *NullableMergeRequestDiffFile) Set(val *MergeRequestDiffFile) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeRequestDiffFile) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeRequestDiffFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeRequestDiffFile(val *MergeRequestDiffFile) *NullableMergeRequestDiffFile {
	return &NullableMergeRequestDiffFile{value: val, isSet: true}
}

func (v NullableMergeRequestDiffFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeRequestDiffFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

