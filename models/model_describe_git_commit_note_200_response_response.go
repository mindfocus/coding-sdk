/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeGitCommitNote200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitCommitNote200ResponseResponse{}

// DescribeGitCommitNote200ResponseResponse 公共返回体
type DescribeGitCommitNote200ResponseResponse struct {
	// 提交注释
	CommitNote *string `json:"CommitNote,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeGitCommitNote200ResponseResponse instantiates a new DescribeGitCommitNote200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitCommitNote200ResponseResponse() *DescribeGitCommitNote200ResponseResponse {
	this := DescribeGitCommitNote200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeGitCommitNote200ResponseResponseWithDefaults instantiates a new DescribeGitCommitNote200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitCommitNote200ResponseResponseWithDefaults() *DescribeGitCommitNote200ResponseResponse {
	this := DescribeGitCommitNote200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetCommitNote returns the CommitNote field value if set, zero value otherwise.
func (o *DescribeGitCommitNote200ResponseResponse) GetCommitNote() string {
	if o == nil || utils.IsNil(o.CommitNote) {
		var ret string
		return ret
	}
	return *o.CommitNote
}

// GetCommitNoteOk returns a tuple with the CommitNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitCommitNote200ResponseResponse) GetCommitNoteOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CommitNote) {
		return nil, false
	}
	return o.CommitNote, true
}

// HasCommitNote returns a boolean if a field has been set.
func (o *DescribeGitCommitNote200ResponseResponse) HasCommitNote() bool {
	if o != nil && !utils.IsNil(o.CommitNote) {
		return true
	}

	return false
}

// SetCommitNote gets a reference to the given string and assigns it to the CommitNote field.
func (o *DescribeGitCommitNote200ResponseResponse) SetCommitNote(v string) {
	o.CommitNote = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeGitCommitNote200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitCommitNote200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeGitCommitNote200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeGitCommitNote200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeGitCommitNote200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitCommitNote200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CommitNote) {
		toSerialize["CommitNote"] = o.CommitNote
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeGitCommitNote200ResponseResponse struct {
	value *DescribeGitCommitNote200ResponseResponse
	isSet bool
}

func (v NullableDescribeGitCommitNote200ResponseResponse) Get() *DescribeGitCommitNote200ResponseResponse {
	return v.value
}

func (v *NullableDescribeGitCommitNote200ResponseResponse) Set(val *DescribeGitCommitNote200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitCommitNote200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitCommitNote200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitCommitNote200ResponseResponse(val *DescribeGitCommitNote200ResponseResponse) *NullableDescribeGitCommitNote200ResponseResponse {
	return &NullableDescribeGitCommitNote200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeGitCommitNote200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitCommitNote200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


