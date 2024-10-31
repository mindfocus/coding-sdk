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

// checks if the ModifyIssueCommentRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyIssueCommentRequest{}

// ModifyIssueCommentRequest struct for ModifyIssueCommentRequest
type ModifyIssueCommentRequest struct {
	// 评论 ID
	CommentId int64 `json:"CommentId"`
	// 评论内容
	Content string `json:"Content"`
	// 事项 Code
	IssueCode int64 `json:"IssueCode"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
}

type _ModifyIssueCommentRequest ModifyIssueCommentRequest

// NewModifyIssueCommentRequest instantiates a new ModifyIssueCommentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyIssueCommentRequest(commentId int64, content string, issueCode int64, projectName string) *ModifyIssueCommentRequest {
	this := ModifyIssueCommentRequest{}
	this.CommentId = commentId
	this.Content = content
	this.IssueCode = issueCode
	this.ProjectName = projectName
	return &this
}

// NewModifyIssueCommentRequestWithDefaults instantiates a new ModifyIssueCommentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyIssueCommentRequestWithDefaults() *ModifyIssueCommentRequest {
	this := ModifyIssueCommentRequest{}
	return &this
}

// GetCommentId returns the CommentId field value
func (o *ModifyIssueCommentRequest) GetCommentId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CommentId
}

// GetCommentIdOk returns a tuple with the CommentId field value
// and a boolean to check if the value has been set.
func (o *ModifyIssueCommentRequest) GetCommentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommentId, true
}

// SetCommentId sets field value
func (o *ModifyIssueCommentRequest) SetCommentId(v int64) {
	o.CommentId = v
}

// GetContent returns the Content field value
func (o *ModifyIssueCommentRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ModifyIssueCommentRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *ModifyIssueCommentRequest) SetContent(v string) {
	o.Content = v
}

// GetIssueCode returns the IssueCode field value
func (o *ModifyIssueCommentRequest) GetIssueCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IssueCode
}

// GetIssueCodeOk returns a tuple with the IssueCode field value
// and a boolean to check if the value has been set.
func (o *ModifyIssueCommentRequest) GetIssueCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueCode, true
}

// SetIssueCode sets field value
func (o *ModifyIssueCommentRequest) SetIssueCode(v int64) {
	o.IssueCode = v
}

// GetProjectName returns the ProjectName field value
func (o *ModifyIssueCommentRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *ModifyIssueCommentRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *ModifyIssueCommentRequest) SetProjectName(v string) {
	o.ProjectName = v
}

func (o ModifyIssueCommentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyIssueCommentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CommentId"] = o.CommentId
	toSerialize["Content"] = o.Content
	toSerialize["IssueCode"] = o.IssueCode
	toSerialize["ProjectName"] = o.ProjectName
	return toSerialize, nil
}

func (o *ModifyIssueCommentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"CommentId",
		"Content",
		"IssueCode",
		"ProjectName",
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

	varModifyIssueCommentRequest := _ModifyIssueCommentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyIssueCommentRequest)

	if err != nil {
		return err
	}

	*o = ModifyIssueCommentRequest(varModifyIssueCommentRequest)

	return err
}

type NullableModifyIssueCommentRequest struct {
	value *ModifyIssueCommentRequest
	isSet bool
}

func (v NullableModifyIssueCommentRequest) Get() *ModifyIssueCommentRequest {
	return v.value
}

func (v *NullableModifyIssueCommentRequest) Set(val *ModifyIssueCommentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyIssueCommentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyIssueCommentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyIssueCommentRequest(val *ModifyIssueCommentRequest) *NullableModifyIssueCommentRequest {
	return &NullableModifyIssueCommentRequest{value: val, isSet: true}
}

func (v NullableModifyIssueCommentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyIssueCommentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

