/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeBlockIssueList200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeBlockIssueList200ResponseResponse{}

// DescribeBlockIssueList200ResponseResponse 公共返回体
type DescribeBlockIssueList200ResponseResponse struct {
	// 后置事项列表
	Issues []IssueSimpleData `json:"Issues,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeBlockIssueList200ResponseResponse instantiates a new DescribeBlockIssueList200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeBlockIssueList200ResponseResponse() *DescribeBlockIssueList200ResponseResponse {
	this := DescribeBlockIssueList200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeBlockIssueList200ResponseResponseWithDefaults instantiates a new DescribeBlockIssueList200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeBlockIssueList200ResponseResponseWithDefaults() *DescribeBlockIssueList200ResponseResponse {
	this := DescribeBlockIssueList200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *DescribeBlockIssueList200ResponseResponse) GetIssues() []IssueSimpleData {
	if o == nil || utils.IsNil(o.Issues) {
		var ret []IssueSimpleData
		return ret
	}
	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeBlockIssueList200ResponseResponse) GetIssuesOk() ([]IssueSimpleData, bool) {
	if o == nil || utils.IsNil(o.Issues) {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *DescribeBlockIssueList200ResponseResponse) HasIssues() bool {
	if o != nil && !utils.IsNil(o.Issues) {
		return true
	}

	return false
}

// SetIssues gets a reference to the given []IssueSimpleData and assigns it to the Issues field.
func (o *DescribeBlockIssueList200ResponseResponse) SetIssues(v []IssueSimpleData) {
	o.Issues = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeBlockIssueList200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeBlockIssueList200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeBlockIssueList200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeBlockIssueList200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeBlockIssueList200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeBlockIssueList200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Issues) {
		toSerialize["Issues"] = o.Issues
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeBlockIssueList200ResponseResponse struct {
	value *DescribeBlockIssueList200ResponseResponse
	isSet bool
}

func (v NullableDescribeBlockIssueList200ResponseResponse) Get() *DescribeBlockIssueList200ResponseResponse {
	return v.value
}

func (v *NullableDescribeBlockIssueList200ResponseResponse) Set(val *DescribeBlockIssueList200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeBlockIssueList200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeBlockIssueList200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeBlockIssueList200ResponseResponse(val *DescribeBlockIssueList200ResponseResponse) *NullableDescribeBlockIssueList200ResponseResponse {
	return &NullableDescribeBlockIssueList200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeBlockIssueList200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeBlockIssueList200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

