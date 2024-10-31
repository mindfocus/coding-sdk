/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeIssueModuleList200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeIssueModuleList200Response{}

// DescribeIssueModuleList200Response struct for DescribeIssueModuleList200Response
type DescribeIssueModuleList200Response struct {
	Response *DescribeIssueModuleList200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeIssueModuleList200Response instantiates a new DescribeIssueModuleList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeIssueModuleList200Response() *DescribeIssueModuleList200Response {
	this := DescribeIssueModuleList200Response{}
	return &this
}

// NewDescribeIssueModuleList200ResponseWithDefaults instantiates a new DescribeIssueModuleList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeIssueModuleList200ResponseWithDefaults() *DescribeIssueModuleList200Response {
	this := DescribeIssueModuleList200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeIssueModuleList200Response) GetResponse() DescribeIssueModuleList200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeIssueModuleList200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeIssueModuleList200Response) GetResponseOk() (*DescribeIssueModuleList200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeIssueModuleList200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeIssueModuleList200ResponseResponse and assigns it to the Response field.
func (o *DescribeIssueModuleList200Response) SetResponse(v DescribeIssueModuleList200ResponseResponse) {
	o.Response = &v
}

func (o DescribeIssueModuleList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeIssueModuleList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeIssueModuleList200Response struct {
	value *DescribeIssueModuleList200Response
	isSet bool
}

func (v NullableDescribeIssueModuleList200Response) Get() *DescribeIssueModuleList200Response {
	return v.value
}

func (v *NullableDescribeIssueModuleList200Response) Set(val *DescribeIssueModuleList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeIssueModuleList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeIssueModuleList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeIssueModuleList200Response(val *DescribeIssueModuleList200Response) *NullableDescribeIssueModuleList200Response {
	return &NullableDescribeIssueModuleList200Response{value: val, isSet: true}
}

func (v NullableDescribeIssueModuleList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeIssueModuleList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


