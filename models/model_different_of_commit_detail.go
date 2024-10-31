/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DifferentOfCommitDetail type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DifferentOfCommitDetail{}

// DifferentOfCommitDetail 请求列表及文件差别列表
type DifferentOfCommitDetail struct {
	// 请求列表
	Commits []GitCommit `json:"Commits,omitempty"`
	// 总删除行数
	Deletions *int64 `json:"Deletions,omitempty"`
	// 差异文件列表
	DifferentOfCommits []DifferentOfCommit `json:"DifferentOfCommits,omitempty"`
	// 总新增行数
	Insertions *int64 `json:"Insertions,omitempty"`
	// 总文件修改数
	UpdateFileNum *int64 `json:"UpdateFileNum,omitempty"`
}

// NewDifferentOfCommitDetail instantiates a new DifferentOfCommitDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDifferentOfCommitDetail() *DifferentOfCommitDetail {
	this := DifferentOfCommitDetail{}
	return &this
}

// NewDifferentOfCommitDetailWithDefaults instantiates a new DifferentOfCommitDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDifferentOfCommitDetailWithDefaults() *DifferentOfCommitDetail {
	this := DifferentOfCommitDetail{}
	return &this
}

// GetCommits returns the Commits field value if set, zero value otherwise.
func (o *DifferentOfCommitDetail) GetCommits() []GitCommit {
	if o == nil || utils.IsNil(o.Commits) {
		var ret []GitCommit
		return ret
	}
	return o.Commits
}

// GetCommitsOk returns a tuple with the Commits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DifferentOfCommitDetail) GetCommitsOk() ([]GitCommit, bool) {
	if o == nil || utils.IsNil(o.Commits) {
		return nil, false
	}
	return o.Commits, true
}

// HasCommits returns a boolean if a field has been set.
func (o *DifferentOfCommitDetail) HasCommits() bool {
	if o != nil && !utils.IsNil(o.Commits) {
		return true
	}

	return false
}

// SetCommits gets a reference to the given []GitCommit and assigns it to the Commits field.
func (o *DifferentOfCommitDetail) SetCommits(v []GitCommit) {
	o.Commits = v
}

// GetDeletions returns the Deletions field value if set, zero value otherwise.
func (o *DifferentOfCommitDetail) GetDeletions() int64 {
	if o == nil || utils.IsNil(o.Deletions) {
		var ret int64
		return ret
	}
	return *o.Deletions
}

// GetDeletionsOk returns a tuple with the Deletions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DifferentOfCommitDetail) GetDeletionsOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Deletions) {
		return nil, false
	}
	return o.Deletions, true
}

// HasDeletions returns a boolean if a field has been set.
func (o *DifferentOfCommitDetail) HasDeletions() bool {
	if o != nil && !utils.IsNil(o.Deletions) {
		return true
	}

	return false
}

// SetDeletions gets a reference to the given int64 and assigns it to the Deletions field.
func (o *DifferentOfCommitDetail) SetDeletions(v int64) {
	o.Deletions = &v
}

// GetDifferentOfCommits returns the DifferentOfCommits field value if set, zero value otherwise.
func (o *DifferentOfCommitDetail) GetDifferentOfCommits() []DifferentOfCommit {
	if o == nil || utils.IsNil(o.DifferentOfCommits) {
		var ret []DifferentOfCommit
		return ret
	}
	return o.DifferentOfCommits
}

// GetDifferentOfCommitsOk returns a tuple with the DifferentOfCommits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DifferentOfCommitDetail) GetDifferentOfCommitsOk() ([]DifferentOfCommit, bool) {
	if o == nil || utils.IsNil(o.DifferentOfCommits) {
		return nil, false
	}
	return o.DifferentOfCommits, true
}

// HasDifferentOfCommits returns a boolean if a field has been set.
func (o *DifferentOfCommitDetail) HasDifferentOfCommits() bool {
	if o != nil && !utils.IsNil(o.DifferentOfCommits) {
		return true
	}

	return false
}

// SetDifferentOfCommits gets a reference to the given []DifferentOfCommit and assigns it to the DifferentOfCommits field.
func (o *DifferentOfCommitDetail) SetDifferentOfCommits(v []DifferentOfCommit) {
	o.DifferentOfCommits = v
}

// GetInsertions returns the Insertions field value if set, zero value otherwise.
func (o *DifferentOfCommitDetail) GetInsertions() int64 {
	if o == nil || utils.IsNil(o.Insertions) {
		var ret int64
		return ret
	}
	return *o.Insertions
}

// GetInsertionsOk returns a tuple with the Insertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DifferentOfCommitDetail) GetInsertionsOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Insertions) {
		return nil, false
	}
	return o.Insertions, true
}

// HasInsertions returns a boolean if a field has been set.
func (o *DifferentOfCommitDetail) HasInsertions() bool {
	if o != nil && !utils.IsNil(o.Insertions) {
		return true
	}

	return false
}

// SetInsertions gets a reference to the given int64 and assigns it to the Insertions field.
func (o *DifferentOfCommitDetail) SetInsertions(v int64) {
	o.Insertions = &v
}

// GetUpdateFileNum returns the UpdateFileNum field value if set, zero value otherwise.
func (o *DifferentOfCommitDetail) GetUpdateFileNum() int64 {
	if o == nil || utils.IsNil(o.UpdateFileNum) {
		var ret int64
		return ret
	}
	return *o.UpdateFileNum
}

// GetUpdateFileNumOk returns a tuple with the UpdateFileNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DifferentOfCommitDetail) GetUpdateFileNumOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.UpdateFileNum) {
		return nil, false
	}
	return o.UpdateFileNum, true
}

// HasUpdateFileNum returns a boolean if a field has been set.
func (o *DifferentOfCommitDetail) HasUpdateFileNum() bool {
	if o != nil && !utils.IsNil(o.UpdateFileNum) {
		return true
	}

	return false
}

// SetUpdateFileNum gets a reference to the given int64 and assigns it to the UpdateFileNum field.
func (o *DifferentOfCommitDetail) SetUpdateFileNum(v int64) {
	o.UpdateFileNum = &v
}

func (o DifferentOfCommitDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DifferentOfCommitDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Commits) {
		toSerialize["Commits"] = o.Commits
	}
	if !utils.IsNil(o.Deletions) {
		toSerialize["Deletions"] = o.Deletions
	}
	if !utils.IsNil(o.DifferentOfCommits) {
		toSerialize["DifferentOfCommits"] = o.DifferentOfCommits
	}
	if !utils.IsNil(o.Insertions) {
		toSerialize["Insertions"] = o.Insertions
	}
	if !utils.IsNil(o.UpdateFileNum) {
		toSerialize["UpdateFileNum"] = o.UpdateFileNum
	}
	return toSerialize, nil
}

type NullableDifferentOfCommitDetail struct {
	value *DifferentOfCommitDetail
	isSet bool
}

func (v NullableDifferentOfCommitDetail) Get() *DifferentOfCommitDetail {
	return v.value
}

func (v *NullableDifferentOfCommitDetail) Set(val *DifferentOfCommitDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableDifferentOfCommitDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableDifferentOfCommitDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDifferentOfCommitDetail(val *DifferentOfCommitDetail) *NullableDifferentOfCommitDetail {
	return &NullableDifferentOfCommitDetail{value: val, isSet: true}
}

func (v NullableDifferentOfCommitDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDifferentOfCommitDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

