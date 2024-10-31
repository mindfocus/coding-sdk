/*
CODING OPEN API

  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeRelatedCaseList200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeRelatedCaseList200ResponseResponse{}

// DescribeRelatedCaseList200ResponseResponse 公共返回体
type DescribeRelatedCaseList200ResponseResponse struct {
	// 事项关联的测试用例
	Cases []RelatedCase `json:"Cases,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeRelatedCaseList200ResponseResponse instantiates a new DescribeRelatedCaseList200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeRelatedCaseList200ResponseResponse() *DescribeRelatedCaseList200ResponseResponse {
	this := DescribeRelatedCaseList200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeRelatedCaseList200ResponseResponseWithDefaults instantiates a new DescribeRelatedCaseList200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeRelatedCaseList200ResponseResponseWithDefaults() *DescribeRelatedCaseList200ResponseResponse {
	this := DescribeRelatedCaseList200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetCases returns the Cases field value if set, zero value otherwise.
func (o *DescribeRelatedCaseList200ResponseResponse) GetCases() []RelatedCase {
	if o == nil || utils.IsNil(o.Cases) {
		var ret []RelatedCase
		return ret
	}
	return o.Cases
}

// GetCasesOk returns a tuple with the Cases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeRelatedCaseList200ResponseResponse) GetCasesOk() ([]RelatedCase, bool) {
	if o == nil || utils.IsNil(o.Cases) {
		return nil, false
	}
	return o.Cases, true
}

// HasCases returns a boolean if a field has been set.
func (o *DescribeRelatedCaseList200ResponseResponse) HasCases() bool {
	if o != nil && !utils.IsNil(o.Cases) {
		return true
	}

	return false
}

// SetCases gets a reference to the given []RelatedCase and assigns it to the Cases field.
func (o *DescribeRelatedCaseList200ResponseResponse) SetCases(v []RelatedCase) {
	o.Cases = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeRelatedCaseList200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeRelatedCaseList200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeRelatedCaseList200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeRelatedCaseList200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeRelatedCaseList200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeRelatedCaseList200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Cases) {
		toSerialize["Cases"] = o.Cases
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeRelatedCaseList200ResponseResponse struct {
	value *DescribeRelatedCaseList200ResponseResponse
	isSet bool
}

func (v NullableDescribeRelatedCaseList200ResponseResponse) Get() *DescribeRelatedCaseList200ResponseResponse {
	return v.value
}

func (v *NullableDescribeRelatedCaseList200ResponseResponse) Set(val *DescribeRelatedCaseList200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeRelatedCaseList200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeRelatedCaseList200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeRelatedCaseList200ResponseResponse(val *DescribeRelatedCaseList200ResponseResponse) *NullableDescribeRelatedCaseList200ResponseResponse {
	return &NullableDescribeRelatedCaseList200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeRelatedCaseList200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeRelatedCaseList200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

