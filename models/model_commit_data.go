/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CommitData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CommitData{}

// CommitData 代码仓库的提交信息
type CommitData struct {
	// 提交信息
	Commits []Commit `json:"Commits,omitempty"`
	Page *PageInfo `json:"Page,omitempty"`
}

// NewCommitData instantiates a new CommitData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitData() *CommitData {
	this := CommitData{}
	return &this
}

// NewCommitDataWithDefaults instantiates a new CommitData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitDataWithDefaults() *CommitData {
	this := CommitData{}
	return &this
}

// GetCommits returns the Commits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommitData) GetCommits() []Commit {
	if o == nil {
		var ret []Commit
		return ret
	}
	return o.Commits
}

// GetCommitsOk returns a tuple with the Commits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommitData) GetCommitsOk() ([]Commit, bool) {
	if o == nil || utils.IsNil(o.Commits) {
		return nil, false
	}
	return o.Commits, true
}

// HasCommits returns a boolean if a field has been set.
func (o *CommitData) HasCommits() bool {
	if o != nil && !utils.IsNil(o.Commits) {
		return true
	}

	return false
}

// SetCommits gets a reference to the given []Commit and assigns it to the Commits field.
func (o *CommitData) SetCommits(v []Commit) {
	o.Commits = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *CommitData) GetPage() PageInfo {
	if o == nil || utils.IsNil(o.Page) {
		var ret PageInfo
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitData) GetPageOk() (*PageInfo, bool) {
	if o == nil || utils.IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *CommitData) HasPage() bool {
	if o != nil && !utils.IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given PageInfo and assigns it to the Page field.
func (o *CommitData) SetPage(v PageInfo) {
	o.Page = &v
}

func (o CommitData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommitData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Commits != nil {
		toSerialize["Commits"] = o.Commits
	}
	if !utils.IsNil(o.Page) {
		toSerialize["Page"] = o.Page
	}
	return toSerialize, nil
}

type NullableCommitData struct {
	value *CommitData
	isSet bool
}

func (v NullableCommitData) Get() *CommitData {
	return v.value
}

func (v *NullableCommitData) Set(val *CommitData) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitData) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitData(val *CommitData) *NullableCommitData {
	return &NullableCommitData{value: val, isSet: true}
}

func (v NullableCommitData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


