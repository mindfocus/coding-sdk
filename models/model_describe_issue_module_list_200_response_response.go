/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeIssueModuleList200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeIssueModuleList200ResponseResponse{}

// DescribeIssueModuleList200ResponseResponse 公共返回体
type DescribeIssueModuleList200ResponseResponse struct {
	// 模块列表
	IssueModules []IssueModule `json:"IssueModules,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeIssueModuleList200ResponseResponse instantiates a new DescribeIssueModuleList200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeIssueModuleList200ResponseResponse() *DescribeIssueModuleList200ResponseResponse {
	this := DescribeIssueModuleList200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeIssueModuleList200ResponseResponseWithDefaults instantiates a new DescribeIssueModuleList200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeIssueModuleList200ResponseResponseWithDefaults() *DescribeIssueModuleList200ResponseResponse {
	this := DescribeIssueModuleList200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetIssueModules returns the IssueModules field value if set, zero value otherwise.
func (o *DescribeIssueModuleList200ResponseResponse) GetIssueModules() []IssueModule {
	if o == nil || utils.IsNil(o.IssueModules) {
		var ret []IssueModule
		return ret
	}
	return o.IssueModules
}

// GetIssueModulesOk returns a tuple with the IssueModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeIssueModuleList200ResponseResponse) GetIssueModulesOk() ([]IssueModule, bool) {
	if o == nil || utils.IsNil(o.IssueModules) {
		return nil, false
	}
	return o.IssueModules, true
}

// HasIssueModules returns a boolean if a field has been set.
func (o *DescribeIssueModuleList200ResponseResponse) HasIssueModules() bool {
	if o != nil && !utils.IsNil(o.IssueModules) {
		return true
	}

	return false
}

// SetIssueModules gets a reference to the given []IssueModule and assigns it to the IssueModules field.
func (o *DescribeIssueModuleList200ResponseResponse) SetIssueModules(v []IssueModule) {
	o.IssueModules = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeIssueModuleList200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeIssueModuleList200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeIssueModuleList200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeIssueModuleList200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeIssueModuleList200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeIssueModuleList200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.IssueModules) {
		toSerialize["IssueModules"] = o.IssueModules
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeIssueModuleList200ResponseResponse struct {
	value *DescribeIssueModuleList200ResponseResponse
	isSet bool
}

func (v NullableDescribeIssueModuleList200ResponseResponse) Get() *DescribeIssueModuleList200ResponseResponse {
	return v.value
}

func (v *NullableDescribeIssueModuleList200ResponseResponse) Set(val *DescribeIssueModuleList200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeIssueModuleList200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeIssueModuleList200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeIssueModuleList200ResponseResponse(val *DescribeIssueModuleList200ResponseResponse) *NullableDescribeIssueModuleList200ResponseResponse {
	return &NullableDescribeIssueModuleList200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeIssueModuleList200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeIssueModuleList200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


