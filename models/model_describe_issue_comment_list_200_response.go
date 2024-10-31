/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeIssueCommentList200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeIssueCommentList200Response{}

// DescribeIssueCommentList200Response struct for DescribeIssueCommentList200Response
type DescribeIssueCommentList200Response struct {
	Response *DescribeIssueCommentList200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeIssueCommentList200Response instantiates a new DescribeIssueCommentList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeIssueCommentList200Response() *DescribeIssueCommentList200Response {
	this := DescribeIssueCommentList200Response{}
	return &this
}

// NewDescribeIssueCommentList200ResponseWithDefaults instantiates a new DescribeIssueCommentList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeIssueCommentList200ResponseWithDefaults() *DescribeIssueCommentList200Response {
	this := DescribeIssueCommentList200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeIssueCommentList200Response) GetResponse() DescribeIssueCommentList200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeIssueCommentList200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeIssueCommentList200Response) GetResponseOk() (*DescribeIssueCommentList200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeIssueCommentList200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeIssueCommentList200ResponseResponse and assigns it to the Response field.
func (o *DescribeIssueCommentList200Response) SetResponse(v DescribeIssueCommentList200ResponseResponse) {
	o.Response = &v
}

func (o DescribeIssueCommentList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeIssueCommentList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeIssueCommentList200Response struct {
	value *DescribeIssueCommentList200Response
	isSet bool
}

func (v NullableDescribeIssueCommentList200Response) Get() *DescribeIssueCommentList200Response {
	return v.value
}

func (v *NullableDescribeIssueCommentList200Response) Set(val *DescribeIssueCommentList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeIssueCommentList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeIssueCommentList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeIssueCommentList200Response(val *DescribeIssueCommentList200Response) *NullableDescribeIssueCommentList200Response {
	return &NullableDescribeIssueCommentList200Response{value: val, isSet: true}
}

func (v NullableDescribeIssueCommentList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeIssueCommentList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

