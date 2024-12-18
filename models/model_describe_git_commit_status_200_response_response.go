/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeGitCommitStatus200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitCommitStatus200ResponseResponse{}

// DescribeGitCommitStatus200ResponseResponse 公共返回体
type DescribeGitCommitStatus200ResponseResponse struct {
	//  提交流水线状态列表
	StatusCheckResults []StatusCheckResult `json:"StatusCheckResults,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeGitCommitStatus200ResponseResponse instantiates a new DescribeGitCommitStatus200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitCommitStatus200ResponseResponse() *DescribeGitCommitStatus200ResponseResponse {
	this := DescribeGitCommitStatus200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeGitCommitStatus200ResponseResponseWithDefaults instantiates a new DescribeGitCommitStatus200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitCommitStatus200ResponseResponseWithDefaults() *DescribeGitCommitStatus200ResponseResponse {
	this := DescribeGitCommitStatus200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetStatusCheckResults returns the StatusCheckResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DescribeGitCommitStatus200ResponseResponse) GetStatusCheckResults() []StatusCheckResult {
	if o == nil {
		var ret []StatusCheckResult
		return ret
	}
	return o.StatusCheckResults
}

// GetStatusCheckResultsOk returns a tuple with the StatusCheckResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DescribeGitCommitStatus200ResponseResponse) GetStatusCheckResultsOk() ([]StatusCheckResult, bool) {
	if o == nil || utils.IsNil(o.StatusCheckResults) {
		return nil, false
	}
	return o.StatusCheckResults, true
}

// HasStatusCheckResults returns a boolean if a field has been set.
func (o *DescribeGitCommitStatus200ResponseResponse) HasStatusCheckResults() bool {
	if o != nil && !utils.IsNil(o.StatusCheckResults) {
		return true
	}

	return false
}

// SetStatusCheckResults gets a reference to the given []StatusCheckResult and assigns it to the StatusCheckResults field.
func (o *DescribeGitCommitStatus200ResponseResponse) SetStatusCheckResults(v []StatusCheckResult) {
	o.StatusCheckResults = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeGitCommitStatus200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitCommitStatus200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeGitCommitStatus200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeGitCommitStatus200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeGitCommitStatus200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitCommitStatus200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.StatusCheckResults != nil {
		toSerialize["StatusCheckResults"] = o.StatusCheckResults
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeGitCommitStatus200ResponseResponse struct {
	value *DescribeGitCommitStatus200ResponseResponse
	isSet bool
}

func (v NullableDescribeGitCommitStatus200ResponseResponse) Get() *DescribeGitCommitStatus200ResponseResponse {
	return v.value
}

func (v *NullableDescribeGitCommitStatus200ResponseResponse) Set(val *DescribeGitCommitStatus200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitCommitStatus200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitCommitStatus200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitCommitStatus200ResponseResponse(val *DescribeGitCommitStatus200ResponseResponse) *NullableDescribeGitCommitStatus200ResponseResponse {
	return &NullableDescribeGitCommitStatus200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeGitCommitStatus200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitCommitStatus200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


