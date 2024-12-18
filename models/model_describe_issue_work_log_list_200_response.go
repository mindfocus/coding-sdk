/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeIssueWorkLogList200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeIssueWorkLogList200Response{}

// DescribeIssueWorkLogList200Response struct for DescribeIssueWorkLogList200Response
type DescribeIssueWorkLogList200Response struct {
	Response *DescribeIssueWorkLogList200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeIssueWorkLogList200Response instantiates a new DescribeIssueWorkLogList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeIssueWorkLogList200Response() *DescribeIssueWorkLogList200Response {
	this := DescribeIssueWorkLogList200Response{}
	return &this
}

// NewDescribeIssueWorkLogList200ResponseWithDefaults instantiates a new DescribeIssueWorkLogList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeIssueWorkLogList200ResponseWithDefaults() *DescribeIssueWorkLogList200Response {
	this := DescribeIssueWorkLogList200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeIssueWorkLogList200Response) GetResponse() DescribeIssueWorkLogList200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeIssueWorkLogList200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeIssueWorkLogList200Response) GetResponseOk() (*DescribeIssueWorkLogList200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeIssueWorkLogList200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeIssueWorkLogList200ResponseResponse and assigns it to the Response field.
func (o *DescribeIssueWorkLogList200Response) SetResponse(v DescribeIssueWorkLogList200ResponseResponse) {
	o.Response = &v
}

func (o DescribeIssueWorkLogList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeIssueWorkLogList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeIssueWorkLogList200Response struct {
	value *DescribeIssueWorkLogList200Response
	isSet bool
}

func (v NullableDescribeIssueWorkLogList200Response) Get() *DescribeIssueWorkLogList200Response {
	return v.value
}

func (v *NullableDescribeIssueWorkLogList200Response) Set(val *DescribeIssueWorkLogList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeIssueWorkLogList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeIssueWorkLogList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeIssueWorkLogList200Response(val *DescribeIssueWorkLogList200Response) *NullableDescribeIssueWorkLogList200Response {
	return &NullableDescribeIssueWorkLogList200Response{value: val, isSet: true}
}

func (v NullableDescribeIssueWorkLogList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeIssueWorkLogList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


