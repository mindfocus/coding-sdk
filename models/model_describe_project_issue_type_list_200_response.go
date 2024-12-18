/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeProjectIssueTypeList200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeProjectIssueTypeList200Response{}

// DescribeProjectIssueTypeList200Response struct for DescribeProjectIssueTypeList200Response
type DescribeProjectIssueTypeList200Response struct {
	Response *DescribeProjectIssueTypeList200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeProjectIssueTypeList200Response instantiates a new DescribeProjectIssueTypeList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeProjectIssueTypeList200Response() *DescribeProjectIssueTypeList200Response {
	this := DescribeProjectIssueTypeList200Response{}
	return &this
}

// NewDescribeProjectIssueTypeList200ResponseWithDefaults instantiates a new DescribeProjectIssueTypeList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeProjectIssueTypeList200ResponseWithDefaults() *DescribeProjectIssueTypeList200Response {
	this := DescribeProjectIssueTypeList200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeProjectIssueTypeList200Response) GetResponse() DescribeProjectIssueTypeList200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeProjectIssueTypeList200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectIssueTypeList200Response) GetResponseOk() (*DescribeProjectIssueTypeList200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeProjectIssueTypeList200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeProjectIssueTypeList200ResponseResponse and assigns it to the Response field.
func (o *DescribeProjectIssueTypeList200Response) SetResponse(v DescribeProjectIssueTypeList200ResponseResponse) {
	o.Response = &v
}

func (o DescribeProjectIssueTypeList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeProjectIssueTypeList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeProjectIssueTypeList200Response struct {
	value *DescribeProjectIssueTypeList200Response
	isSet bool
}

func (v NullableDescribeProjectIssueTypeList200Response) Get() *DescribeProjectIssueTypeList200Response {
	return v.value
}

func (v *NullableDescribeProjectIssueTypeList200Response) Set(val *DescribeProjectIssueTypeList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeProjectIssueTypeList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeProjectIssueTypeList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeProjectIssueTypeList200Response(val *DescribeProjectIssueTypeList200Response) *NullableDescribeProjectIssueTypeList200Response {
	return &NullableDescribeProjectIssueTypeList200Response{value: val, isSet: true}
}

func (v NullableDescribeProjectIssueTypeList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeProjectIssueTypeList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


