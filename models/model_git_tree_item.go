/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the GitTreeItem type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GitTreeItem{}

// GitTreeItem 仓库文件信息
type GitTreeItem struct {
	// 文件类型
	Mode *string `json:"Mode,omitempty"`
	// 文件名
	Name *string `json:"Name,omitempty"`
	// 文件路径
	Path *string `json:"Path,omitempty"`
	// commitID
	Sha *string `json:"Sha,omitempty"`
}

// NewGitTreeItem instantiates a new GitTreeItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitTreeItem() *GitTreeItem {
	this := GitTreeItem{}
	var mode string = ""
	this.Mode = &mode
	var name string = ""
	this.Name = &name
	var path string = ""
	this.Path = &path
	var sha string = ""
	this.Sha = &sha
	return &this
}

// NewGitTreeItemWithDefaults instantiates a new GitTreeItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitTreeItemWithDefaults() *GitTreeItem {
	this := GitTreeItem{}
	var mode string = ""
	this.Mode = &mode
	var name string = ""
	this.Name = &name
	var path string = ""
	this.Path = &path
	var sha string = ""
	this.Sha = &sha
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *GitTreeItem) GetMode() string {
	if o == nil || utils.IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitTreeItem) GetModeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *GitTreeItem) HasMode() bool {
	if o != nil && !utils.IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *GitTreeItem) SetMode(v string) {
	o.Mode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GitTreeItem) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitTreeItem) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GitTreeItem) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GitTreeItem) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *GitTreeItem) GetPath() string {
	if o == nil || utils.IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitTreeItem) GetPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *GitTreeItem) HasPath() bool {
	if o != nil && !utils.IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *GitTreeItem) SetPath(v string) {
	o.Path = &v
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *GitTreeItem) GetSha() string {
	if o == nil || utils.IsNil(o.Sha) {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitTreeItem) GetShaOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Sha) {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *GitTreeItem) HasSha() bool {
	if o != nil && !utils.IsNil(o.Sha) {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *GitTreeItem) SetSha(v string) {
	o.Sha = &v
}

func (o GitTreeItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitTreeItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Mode) {
		toSerialize["Mode"] = o.Mode
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !utils.IsNil(o.Path) {
		toSerialize["Path"] = o.Path
	}
	if !utils.IsNil(o.Sha) {
		toSerialize["Sha"] = o.Sha
	}
	return toSerialize, nil
}

type NullableGitTreeItem struct {
	value *GitTreeItem
	isSet bool
}

func (v NullableGitTreeItem) Get() *GitTreeItem {
	return v.value
}

func (v *NullableGitTreeItem) Set(val *GitTreeItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGitTreeItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGitTreeItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitTreeItem(val *GitTreeItem) *NullableGitTreeItem {
	return &NullableGitTreeItem{value: val, isSet: true}
}

func (v NullableGitTreeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitTreeItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

