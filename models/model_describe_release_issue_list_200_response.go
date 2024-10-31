/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeReleaseIssueList200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeReleaseIssueList200Response{}

// DescribeReleaseIssueList200Response struct for DescribeReleaseIssueList200Response
type DescribeReleaseIssueList200Response struct {
	Response *DescribeReleaseIssueList200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeReleaseIssueList200Response instantiates a new DescribeReleaseIssueList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeReleaseIssueList200Response() *DescribeReleaseIssueList200Response {
	this := DescribeReleaseIssueList200Response{}
	return &this
}

// NewDescribeReleaseIssueList200ResponseWithDefaults instantiates a new DescribeReleaseIssueList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeReleaseIssueList200ResponseWithDefaults() *DescribeReleaseIssueList200Response {
	this := DescribeReleaseIssueList200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeReleaseIssueList200Response) GetResponse() DescribeReleaseIssueList200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeReleaseIssueList200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseIssueList200Response) GetResponseOk() (*DescribeReleaseIssueList200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeReleaseIssueList200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeReleaseIssueList200ResponseResponse and assigns it to the Response field.
func (o *DescribeReleaseIssueList200Response) SetResponse(v DescribeReleaseIssueList200ResponseResponse) {
	o.Response = &v
}

func (o DescribeReleaseIssueList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeReleaseIssueList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeReleaseIssueList200Response struct {
	value *DescribeReleaseIssueList200Response
	isSet bool
}

func (v NullableDescribeReleaseIssueList200Response) Get() *DescribeReleaseIssueList200Response {
	return v.value
}

func (v *NullableDescribeReleaseIssueList200Response) Set(val *DescribeReleaseIssueList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeReleaseIssueList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeReleaseIssueList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeReleaseIssueList200Response(val *DescribeReleaseIssueList200Response) *NullableDescribeReleaseIssueList200Response {
	return &NullableDescribeReleaseIssueList200Response{value: val, isSet: true}
}

func (v NullableDescribeReleaseIssueList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeReleaseIssueList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


