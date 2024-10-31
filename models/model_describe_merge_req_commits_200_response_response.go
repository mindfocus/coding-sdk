/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeMergeReqCommits200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeMergeReqCommits200ResponseResponse{}

// DescribeMergeReqCommits200ResponseResponse 公共返回体
type DescribeMergeReqCommits200ResponseResponse struct {
	// commit集合
	Commits []GitCommit `json:"Commits,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeMergeReqCommits200ResponseResponse instantiates a new DescribeMergeReqCommits200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeMergeReqCommits200ResponseResponse() *DescribeMergeReqCommits200ResponseResponse {
	this := DescribeMergeReqCommits200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeMergeReqCommits200ResponseResponseWithDefaults instantiates a new DescribeMergeReqCommits200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeMergeReqCommits200ResponseResponseWithDefaults() *DescribeMergeReqCommits200ResponseResponse {
	this := DescribeMergeReqCommits200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetCommits returns the Commits field value if set, zero value otherwise.
func (o *DescribeMergeReqCommits200ResponseResponse) GetCommits() []GitCommit {
	if o == nil || utils.IsNil(o.Commits) {
		var ret []GitCommit
		return ret
	}
	return o.Commits
}

// GetCommitsOk returns a tuple with the Commits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeMergeReqCommits200ResponseResponse) GetCommitsOk() ([]GitCommit, bool) {
	if o == nil || utils.IsNil(o.Commits) {
		return nil, false
	}
	return o.Commits, true
}

// HasCommits returns a boolean if a field has been set.
func (o *DescribeMergeReqCommits200ResponseResponse) HasCommits() bool {
	if o != nil && !utils.IsNil(o.Commits) {
		return true
	}

	return false
}

// SetCommits gets a reference to the given []GitCommit and assigns it to the Commits field.
func (o *DescribeMergeReqCommits200ResponseResponse) SetCommits(v []GitCommit) {
	o.Commits = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeMergeReqCommits200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeMergeReqCommits200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeMergeReqCommits200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeMergeReqCommits200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeMergeReqCommits200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeMergeReqCommits200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Commits) {
		toSerialize["Commits"] = o.Commits
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeMergeReqCommits200ResponseResponse struct {
	value *DescribeMergeReqCommits200ResponseResponse
	isSet bool
}

func (v NullableDescribeMergeReqCommits200ResponseResponse) Get() *DescribeMergeReqCommits200ResponseResponse {
	return v.value
}

func (v *NullableDescribeMergeReqCommits200ResponseResponse) Set(val *DescribeMergeReqCommits200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeMergeReqCommits200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeMergeReqCommits200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeMergeReqCommits200ResponseResponse(val *DescribeMergeReqCommits200ResponseResponse) *NullableDescribeMergeReqCommits200ResponseResponse {
	return &NullableDescribeMergeReqCommits200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeMergeReqCommits200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeMergeReqCommits200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


