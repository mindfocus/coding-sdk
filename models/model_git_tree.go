/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the GitTree type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GitTree{}

// GitTree git的树
type GitTree struct {
	// 文件或目录权限
	Mode utils.NullableString `json:"Mode,omitempty"`
	// 路径
	Path utils.NullableString `json:"Path,omitempty"`
	// 哈希值
	Sha utils.NullableString `json:"Sha,omitempty"`
	// 类型
	Type utils.NullableString `json:"Type,omitempty"`
}

// NewGitTree instantiates a new GitTree object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitTree() *GitTree {
	this := GitTree{}
	var mode string = ""
	this.Mode = *utils.NewNullableString(&mode)
	var path string = ""
	this.Path = *utils.NewNullableString(&path)
	var sha string = ""
	this.Sha = *utils.NewNullableString(&sha)
	var type_ string = ""
	this.Type = *utils.NewNullableString(&type_)
	return &this
}

// NewGitTreeWithDefaults instantiates a new GitTree object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitTreeWithDefaults() *GitTree {
	this := GitTree{}
	var mode string = ""
	this.Mode = *utils.NewNullableString(&mode)
	var path string = ""
	this.Path = *utils.NewNullableString(&path)
	var sha string = ""
	this.Sha = *utils.NewNullableString(&sha)
	var type_ string = ""
	this.Type = *utils.NewNullableString(&type_)
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitTree) GetMode() string {
	if o == nil || utils.IsNil(o.Mode.Get()) {
		var ret string
		return ret
	}
	return *o.Mode.Get()
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitTree) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode.Get(), o.Mode.IsSet()
}

// HasMode returns a boolean if a field has been set.
func (o *GitTree) HasMode() bool {
	if o != nil && o.Mode.IsSet() {
		return true
	}

	return false
}

// SetMode gets a reference to the given utils.NullableString and assigns it to the Mode field.
func (o *GitTree) SetMode(v string) {
	o.Mode.Set(&v)
}
// SetModeNil sets the value for Mode to be an explicit nil
func (o *GitTree) SetModeNil() {
	o.Mode.Set(nil)
}

// UnsetMode ensures that no value is present for Mode, not even an explicit nil
func (o *GitTree) UnsetMode() {
	o.Mode.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitTree) GetPath() string {
	if o == nil || utils.IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitTree) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *GitTree) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given utils.NullableString and assigns it to the Path field.
func (o *GitTree) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *GitTree) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *GitTree) UnsetPath() {
	o.Path.Unset()
}

// GetSha returns the Sha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitTree) GetSha() string {
	if o == nil || utils.IsNil(o.Sha.Get()) {
		var ret string
		return ret
	}
	return *o.Sha.Get()
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitTree) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sha.Get(), o.Sha.IsSet()
}

// HasSha returns a boolean if a field has been set.
func (o *GitTree) HasSha() bool {
	if o != nil && o.Sha.IsSet() {
		return true
	}

	return false
}

// SetSha gets a reference to the given utils.NullableString and assigns it to the Sha field.
func (o *GitTree) SetSha(v string) {
	o.Sha.Set(&v)
}
// SetShaNil sets the value for Sha to be an explicit nil
func (o *GitTree) SetShaNil() {
	o.Sha.Set(nil)
}

// UnsetSha ensures that no value is present for Sha, not even an explicit nil
func (o *GitTree) UnsetSha() {
	o.Sha.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitTree) GetType() string {
	if o == nil || utils.IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitTree) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *GitTree) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given utils.NullableString and assigns it to the Type field.
func (o *GitTree) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *GitTree) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *GitTree) UnsetType() {
	o.Type.Unset()
}

func (o GitTree) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitTree) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Mode.IsSet() {
		toSerialize["Mode"] = o.Mode.Get()
	}
	if o.Path.IsSet() {
		toSerialize["Path"] = o.Path.Get()
	}
	if o.Sha.IsSet() {
		toSerialize["Sha"] = o.Sha.Get()
	}
	if o.Type.IsSet() {
		toSerialize["Type"] = o.Type.Get()
	}
	return toSerialize, nil
}

type NullableGitTree struct {
	value *GitTree
	isSet bool
}

func (v NullableGitTree) Get() *GitTree {
	return v.value
}

func (v *NullableGitTree) Set(val *GitTree) {
	v.value = val
	v.isSet = true
}

func (v NullableGitTree) IsSet() bool {
	return v.isSet
}

func (v *NullableGitTree) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitTree(val *GitTree) *NullableGitTree {
	return &NullableGitTree{value: val, isSet: true}
}

func (v NullableGitTree) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitTree) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


