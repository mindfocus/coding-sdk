/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the IssueComment type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &IssueComment{}

// IssueComment 事项评论
type IssueComment struct {
	// 评论 ID
	CommentId *int64 `json:"CommentId,omitempty"`
	// 解析后的内容
	Content *string `json:"Content,omitempty"`
	// 创建时间戳
	CreatedAt *int64 `json:"CreatedAt,omitempty"`
	// 创建人 ID
	CreatorId *int64 `json:"CreatorId,omitempty"`
	// 父评论 ID
	ParentId *int64 `json:"ParentId,omitempty"`
	// 内容
	RawContent *string `json:"RawContent,omitempty"`
	// 更新时间戳
	UpdatedAt *int64 `json:"UpdatedAt,omitempty"`
}

// NewIssueComment instantiates a new IssueComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueComment() *IssueComment {
	this := IssueComment{}
	var content string = ""
	this.Content = &content
	var rawContent string = ""
	this.RawContent = &rawContent
	return &this
}

// NewIssueCommentWithDefaults instantiates a new IssueComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueCommentWithDefaults() *IssueComment {
	this := IssueComment{}
	var content string = ""
	this.Content = &content
	var rawContent string = ""
	this.RawContent = &rawContent
	return &this
}

// GetCommentId returns the CommentId field value if set, zero value otherwise.
func (o *IssueComment) GetCommentId() int64 {
	if o == nil || utils.IsNil(o.CommentId) {
		var ret int64
		return ret
	}
	return *o.CommentId
}

// GetCommentIdOk returns a tuple with the CommentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueComment) GetCommentIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.CommentId) {
		return nil, false
	}
	return o.CommentId, true
}

// HasCommentId returns a boolean if a field has been set.
func (o *IssueComment) HasCommentId() bool {
	if o != nil && !utils.IsNil(o.CommentId) {
		return true
	}

	return false
}

// SetCommentId gets a reference to the given int64 and assigns it to the CommentId field.
func (o *IssueComment) SetCommentId(v int64) {
	o.CommentId = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *IssueComment) GetContent() string {
	if o == nil || utils.IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueComment) GetContentOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *IssueComment) HasContent() bool {
	if o != nil && !utils.IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *IssueComment) SetContent(v string) {
	o.Content = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IssueComment) GetCreatedAt() int64 {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueComment) GetCreatedAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IssueComment) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *IssueComment) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *IssueComment) GetCreatorId() int64 {
	if o == nil || utils.IsNil(o.CreatorId) {
		var ret int64
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueComment) GetCreatorIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *IssueComment) HasCreatorId() bool {
	if o != nil && !utils.IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given int64 and assigns it to the CreatorId field.
func (o *IssueComment) SetCreatorId(v int64) {
	o.CreatorId = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *IssueComment) GetParentId() int64 {
	if o == nil || utils.IsNil(o.ParentId) {
		var ret int64
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueComment) GetParentIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *IssueComment) HasParentId() bool {
	if o != nil && !utils.IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given int64 and assigns it to the ParentId field.
func (o *IssueComment) SetParentId(v int64) {
	o.ParentId = &v
}

// GetRawContent returns the RawContent field value if set, zero value otherwise.
func (o *IssueComment) GetRawContent() string {
	if o == nil || utils.IsNil(o.RawContent) {
		var ret string
		return ret
	}
	return *o.RawContent
}

// GetRawContentOk returns a tuple with the RawContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueComment) GetRawContentOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RawContent) {
		return nil, false
	}
	return o.RawContent, true
}

// HasRawContent returns a boolean if a field has been set.
func (o *IssueComment) HasRawContent() bool {
	if o != nil && !utils.IsNil(o.RawContent) {
		return true
	}

	return false
}

// SetRawContent gets a reference to the given string and assigns it to the RawContent field.
func (o *IssueComment) SetRawContent(v string) {
	o.RawContent = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IssueComment) GetUpdatedAt() int64 {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueComment) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IssueComment) HasUpdatedAt() bool {
	if o != nil && !utils.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *IssueComment) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

func (o IssueComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CommentId) {
		toSerialize["CommentId"] = o.CommentId
	}
	if !utils.IsNil(o.Content) {
		toSerialize["Content"] = o.Content
	}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if !utils.IsNil(o.CreatorId) {
		toSerialize["CreatorId"] = o.CreatorId
	}
	if !utils.IsNil(o.ParentId) {
		toSerialize["ParentId"] = o.ParentId
	}
	if !utils.IsNil(o.RawContent) {
		toSerialize["RawContent"] = o.RawContent
	}
	if !utils.IsNil(o.UpdatedAt) {
		toSerialize["UpdatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableIssueComment struct {
	value *IssueComment
	isSet bool
}

func (v NullableIssueComment) Get() *IssueComment {
	return v.value
}

func (v *NullableIssueComment) Set(val *IssueComment) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueComment) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueComment(val *IssueComment) *NullableIssueComment {
	return &NullableIssueComment{value: val, isSet: true}
}

func (v NullableIssueComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

