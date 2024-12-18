/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCommitRefs200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCommitRefs200ResponseResponse{}

// DescribeCommitRefs200ResponseResponse 公共返回体
type DescribeCommitRefs200ResponseResponse struct {
	// 仓库信息详情列表
	CommitRefs []CommitRef `json:"CommitRefs,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeCommitRefs200ResponseResponse instantiates a new DescribeCommitRefs200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCommitRefs200ResponseResponse() *DescribeCommitRefs200ResponseResponse {
	this := DescribeCommitRefs200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeCommitRefs200ResponseResponseWithDefaults instantiates a new DescribeCommitRefs200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCommitRefs200ResponseResponseWithDefaults() *DescribeCommitRefs200ResponseResponse {
	this := DescribeCommitRefs200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetCommitRefs returns the CommitRefs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DescribeCommitRefs200ResponseResponse) GetCommitRefs() []CommitRef {
	if o == nil {
		var ret []CommitRef
		return ret
	}
	return o.CommitRefs
}

// GetCommitRefsOk returns a tuple with the CommitRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DescribeCommitRefs200ResponseResponse) GetCommitRefsOk() ([]CommitRef, bool) {
	if o == nil || utils.IsNil(o.CommitRefs) {
		return nil, false
	}
	return o.CommitRefs, true
}

// HasCommitRefs returns a boolean if a field has been set.
func (o *DescribeCommitRefs200ResponseResponse) HasCommitRefs() bool {
	if o != nil && !utils.IsNil(o.CommitRefs) {
		return true
	}

	return false
}

// SetCommitRefs gets a reference to the given []CommitRef and assigns it to the CommitRefs field.
func (o *DescribeCommitRefs200ResponseResponse) SetCommitRefs(v []CommitRef) {
	o.CommitRefs = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeCommitRefs200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCommitRefs200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeCommitRefs200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeCommitRefs200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeCommitRefs200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCommitRefs200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CommitRefs != nil {
		toSerialize["CommitRefs"] = o.CommitRefs
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeCommitRefs200ResponseResponse struct {
	value *DescribeCommitRefs200ResponseResponse
	isSet bool
}

func (v NullableDescribeCommitRefs200ResponseResponse) Get() *DescribeCommitRefs200ResponseResponse {
	return v.value
}

func (v *NullableDescribeCommitRefs200ResponseResponse) Set(val *DescribeCommitRefs200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCommitRefs200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCommitRefs200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCommitRefs200ResponseResponse(val *DescribeCommitRefs200ResponseResponse) *NullableDescribeCommitRefs200ResponseResponse {
	return &NullableDescribeCommitRefs200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeCommitRefs200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCommitRefs200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


