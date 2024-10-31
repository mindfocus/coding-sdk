/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ModifyIssue200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyIssue200ResponseResponse{}

// ModifyIssue200ResponseResponse 公共返回体
type ModifyIssue200ResponseResponse struct {
	Issue *IssueDetail `json:"Issue,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewModifyIssue200ResponseResponse instantiates a new ModifyIssue200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyIssue200ResponseResponse() *ModifyIssue200ResponseResponse {
	this := ModifyIssue200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewModifyIssue200ResponseResponseWithDefaults instantiates a new ModifyIssue200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyIssue200ResponseResponseWithDefaults() *ModifyIssue200ResponseResponse {
	this := ModifyIssue200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetIssue returns the Issue field value if set, zero value otherwise.
func (o *ModifyIssue200ResponseResponse) GetIssue() IssueDetail {
	if o == nil || utils.IsNil(o.Issue) {
		var ret IssueDetail
		return ret
	}
	return *o.Issue
}

// GetIssueOk returns a tuple with the Issue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssue200ResponseResponse) GetIssueOk() (*IssueDetail, bool) {
	if o == nil || utils.IsNil(o.Issue) {
		return nil, false
	}
	return o.Issue, true
}

// HasIssue returns a boolean if a field has been set.
func (o *ModifyIssue200ResponseResponse) HasIssue() bool {
	if o != nil && !utils.IsNil(o.Issue) {
		return true
	}

	return false
}

// SetIssue gets a reference to the given IssueDetail and assigns it to the Issue field.
func (o *ModifyIssue200ResponseResponse) SetIssue(v IssueDetail) {
	o.Issue = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ModifyIssue200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssue200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ModifyIssue200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ModifyIssue200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o ModifyIssue200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyIssue200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Issue) {
		toSerialize["Issue"] = o.Issue
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableModifyIssue200ResponseResponse struct {
	value *ModifyIssue200ResponseResponse
	isSet bool
}

func (v NullableModifyIssue200ResponseResponse) Get() *ModifyIssue200ResponseResponse {
	return v.value
}

func (v *NullableModifyIssue200ResponseResponse) Set(val *ModifyIssue200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyIssue200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyIssue200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyIssue200ResponseResponse(val *ModifyIssue200ResponseResponse) *NullableModifyIssue200ResponseResponse {
	return &NullableModifyIssue200ResponseResponse{value: val, isSet: true}
}

func (v NullableModifyIssue200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyIssue200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

