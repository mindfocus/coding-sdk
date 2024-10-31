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

// checks if the CreateIssueCommentRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateIssueCommentRequest{}

// CreateIssueCommentRequest struct for CreateIssueCommentRequest
type CreateIssueCommentRequest struct {
	// 评论内容
	Content string `json:"Content"`
	// 事项 Code
	IssueCode int64 `json:"IssueCode"`
	// 父评论ID
	ParentId int64 `json:"ParentId"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
}

type _CreateIssueCommentRequest CreateIssueCommentRequest

// NewCreateIssueCommentRequest instantiates a new CreateIssueCommentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIssueCommentRequest(content string, issueCode int64, parentId int64, projectName string) *CreateIssueCommentRequest {
	this := CreateIssueCommentRequest{}
	this.Content = content
	this.IssueCode = issueCode
	this.ParentId = parentId
	this.ProjectName = projectName
	return &this
}

// NewCreateIssueCommentRequestWithDefaults instantiates a new CreateIssueCommentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIssueCommentRequestWithDefaults() *CreateIssueCommentRequest {
	this := CreateIssueCommentRequest{}
	return &this
}

// GetContent returns the Content field value
func (o *CreateIssueCommentRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CreateIssueCommentRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CreateIssueCommentRequest) SetContent(v string) {
	o.Content = v
}

// GetIssueCode returns the IssueCode field value
func (o *CreateIssueCommentRequest) GetIssueCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IssueCode
}

// GetIssueCodeOk returns a tuple with the IssueCode field value
// and a boolean to check if the value has been set.
func (o *CreateIssueCommentRequest) GetIssueCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueCode, true
}

// SetIssueCode sets field value
func (o *CreateIssueCommentRequest) SetIssueCode(v int64) {
	o.IssueCode = v
}

// GetParentId returns the ParentId field value
func (o *CreateIssueCommentRequest) GetParentId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *CreateIssueCommentRequest) GetParentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *CreateIssueCommentRequest) SetParentId(v int64) {
	o.ParentId = v
}

// GetProjectName returns the ProjectName field value
func (o *CreateIssueCommentRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *CreateIssueCommentRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *CreateIssueCommentRequest) SetProjectName(v string) {
	o.ProjectName = v
}

func (o CreateIssueCommentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIssueCommentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Content"] = o.Content
	toSerialize["IssueCode"] = o.IssueCode
	toSerialize["ParentId"] = o.ParentId
	toSerialize["ProjectName"] = o.ProjectName
	return toSerialize, nil
}

func (o *CreateIssueCommentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Content",
		"IssueCode",
		"ParentId",
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

	varCreateIssueCommentRequest := _CreateIssueCommentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateIssueCommentRequest)

	if err != nil {
		return err
	}

	*o = CreateIssueCommentRequest(varCreateIssueCommentRequest)

	return err
}

type NullableCreateIssueCommentRequest struct {
	value *CreateIssueCommentRequest
	isSet bool
}

func (v NullableCreateIssueCommentRequest) Get() *CreateIssueCommentRequest {
	return v.value
}

func (v *NullableCreateIssueCommentRequest) Set(val *CreateIssueCommentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIssueCommentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIssueCommentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIssueCommentRequest(val *CreateIssueCommentRequest) *NullableCreateIssueCommentRequest {
	return &NullableCreateIssueCommentRequest{value: val, isSet: true}
}

func (v NullableCreateIssueCommentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIssueCommentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


