/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the GitCommitComment type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GitCommitComment{}

// GitCommitComment 提交评论信息带分页
type GitCommitComment struct {
	// 提交评论详细信息
	CommitComments []CommitComment `json:"CommitComments,omitempty"`
	Page *PageInfo `json:"Page,omitempty"`
}

// NewGitCommitComment instantiates a new GitCommitComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitCommitComment() *GitCommitComment {
	this := GitCommitComment{}
	return &this
}

// NewGitCommitCommentWithDefaults instantiates a new GitCommitComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitCommitCommentWithDefaults() *GitCommitComment {
	this := GitCommitComment{}
	return &this
}

// GetCommitComments returns the CommitComments field value if set, zero value otherwise.
func (o *GitCommitComment) GetCommitComments() []CommitComment {
	if o == nil || utils.IsNil(o.CommitComments) {
		var ret []CommitComment
		return ret
	}
	return o.CommitComments
}

// GetCommitCommentsOk returns a tuple with the CommitComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitCommitComment) GetCommitCommentsOk() ([]CommitComment, bool) {
	if o == nil || utils.IsNil(o.CommitComments) {
		return nil, false
	}
	return o.CommitComments, true
}

// HasCommitComments returns a boolean if a field has been set.
func (o *GitCommitComment) HasCommitComments() bool {
	if o != nil && !utils.IsNil(o.CommitComments) {
		return true
	}

	return false
}

// SetCommitComments gets a reference to the given []CommitComment and assigns it to the CommitComments field.
func (o *GitCommitComment) SetCommitComments(v []CommitComment) {
	o.CommitComments = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GitCommitComment) GetPage() PageInfo {
	if o == nil || utils.IsNil(o.Page) {
		var ret PageInfo
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitCommitComment) GetPageOk() (*PageInfo, bool) {
	if o == nil || utils.IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GitCommitComment) HasPage() bool {
	if o != nil && !utils.IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given PageInfo and assigns it to the Page field.
func (o *GitCommitComment) SetPage(v PageInfo) {
	o.Page = &v
}

func (o GitCommitComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitCommitComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CommitComments) {
		toSerialize["CommitComments"] = o.CommitComments
	}
	if !utils.IsNil(o.Page) {
		toSerialize["Page"] = o.Page
	}
	return toSerialize, nil
}

type NullableGitCommitComment struct {
	value *GitCommitComment
	isSet bool
}

func (v NullableGitCommitComment) Get() *GitCommitComment {
	return v.value
}

func (v *NullableGitCommitComment) Set(val *GitCommitComment) {
	v.value = val
	v.isSet = true
}

func (v NullableGitCommitComment) IsSet() bool {
	return v.isSet
}

func (v *NullableGitCommitComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitCommitComment(val *GitCommitComment) *NullableGitCommitComment {
	return &NullableGitCommitComment{value: val, isSet: true}
}

func (v NullableGitCommitComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitCommitComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


