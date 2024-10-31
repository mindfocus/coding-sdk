/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeIssueList200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeIssueList200Response{}

// DescribeIssueList200Response struct for DescribeIssueList200Response
type DescribeIssueList200Response struct {
	Response *DescribeIssueList200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeIssueList200Response instantiates a new DescribeIssueList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeIssueList200Response() *DescribeIssueList200Response {
	this := DescribeIssueList200Response{}
	return &this
}

// NewDescribeIssueList200ResponseWithDefaults instantiates a new DescribeIssueList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeIssueList200ResponseWithDefaults() *DescribeIssueList200Response {
	this := DescribeIssueList200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeIssueList200Response) GetResponse() DescribeIssueList200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeIssueList200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeIssueList200Response) GetResponseOk() (*DescribeIssueList200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeIssueList200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeIssueList200ResponseResponse and assigns it to the Response field.
func (o *DescribeIssueList200Response) SetResponse(v DescribeIssueList200ResponseResponse) {
	o.Response = &v
}

func (o DescribeIssueList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeIssueList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeIssueList200Response struct {
	value *DescribeIssueList200Response
	isSet bool
}

func (v NullableDescribeIssueList200Response) Get() *DescribeIssueList200Response {
	return v.value
}

func (v *NullableDescribeIssueList200Response) Set(val *DescribeIssueList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeIssueList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeIssueList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeIssueList200Response(val *DescribeIssueList200Response) *NullableDescribeIssueList200Response {
	return &NullableDescribeIssueList200Response{value: val, isSet: true}
}

func (v NullableDescribeIssueList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeIssueList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

