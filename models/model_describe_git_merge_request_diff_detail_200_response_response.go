/*
CODING OPEN API

  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeGitMergeRequestDiffDetail200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitMergeRequestDiffDetail200ResponseResponse{}

// DescribeGitMergeRequestDiffDetail200ResponseResponse 公共返回体
type DescribeGitMergeRequestDiffDetail200ResponseResponse struct {
	Detail *GitDiff `json:"Detail,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeGitMergeRequestDiffDetail200ResponseResponse instantiates a new DescribeGitMergeRequestDiffDetail200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitMergeRequestDiffDetail200ResponseResponse() *DescribeGitMergeRequestDiffDetail200ResponseResponse {
	this := DescribeGitMergeRequestDiffDetail200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeGitMergeRequestDiffDetail200ResponseResponseWithDefaults instantiates a new DescribeGitMergeRequestDiffDetail200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitMergeRequestDiffDetail200ResponseResponseWithDefaults() *DescribeGitMergeRequestDiffDetail200ResponseResponse {
	this := DescribeGitMergeRequestDiffDetail200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *DescribeGitMergeRequestDiffDetail200ResponseResponse) GetDetail() GitDiff {
	if o == nil || utils.IsNil(o.Detail) {
		var ret GitDiff
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitMergeRequestDiffDetail200ResponseResponse) GetDetailOk() (*GitDiff, bool) {
	if o == nil || utils.IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *DescribeGitMergeRequestDiffDetail200ResponseResponse) HasDetail() bool {
	if o != nil && !utils.IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given GitDiff and assigns it to the Detail field.
func (o *DescribeGitMergeRequestDiffDetail200ResponseResponse) SetDetail(v GitDiff) {
	o.Detail = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeGitMergeRequestDiffDetail200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitMergeRequestDiffDetail200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeGitMergeRequestDiffDetail200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeGitMergeRequestDiffDetail200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeGitMergeRequestDiffDetail200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitMergeRequestDiffDetail200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Detail) {
		toSerialize["Detail"] = o.Detail
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeGitMergeRequestDiffDetail200ResponseResponse struct {
	value *DescribeGitMergeRequestDiffDetail200ResponseResponse
	isSet bool
}

func (v NullableDescribeGitMergeRequestDiffDetail200ResponseResponse) Get() *DescribeGitMergeRequestDiffDetail200ResponseResponse {
	return v.value
}

func (v *NullableDescribeGitMergeRequestDiffDetail200ResponseResponse) Set(val *DescribeGitMergeRequestDiffDetail200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitMergeRequestDiffDetail200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitMergeRequestDiffDetail200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitMergeRequestDiffDetail200ResponseResponse(val *DescribeGitMergeRequestDiffDetail200ResponseResponse) *NullableDescribeGitMergeRequestDiffDetail200ResponseResponse {
	return &NullableDescribeGitMergeRequestDiffDetail200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeGitMergeRequestDiffDetail200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitMergeRequestDiffDetail200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


