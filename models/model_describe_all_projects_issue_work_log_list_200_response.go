/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeAllProjectsIssueWorkLogList200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeAllProjectsIssueWorkLogList200Response{}

// DescribeAllProjectsIssueWorkLogList200Response struct for DescribeAllProjectsIssueWorkLogList200Response
type DescribeAllProjectsIssueWorkLogList200Response struct {
	Response *DescribeIssueWorkLogList200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeAllProjectsIssueWorkLogList200Response instantiates a new DescribeAllProjectsIssueWorkLogList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeAllProjectsIssueWorkLogList200Response() *DescribeAllProjectsIssueWorkLogList200Response {
	this := DescribeAllProjectsIssueWorkLogList200Response{}
	return &this
}

// NewDescribeAllProjectsIssueWorkLogList200ResponseWithDefaults instantiates a new DescribeAllProjectsIssueWorkLogList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeAllProjectsIssueWorkLogList200ResponseWithDefaults() *DescribeAllProjectsIssueWorkLogList200Response {
	this := DescribeAllProjectsIssueWorkLogList200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeAllProjectsIssueWorkLogList200Response) GetResponse() DescribeIssueWorkLogList200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeIssueWorkLogList200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAllProjectsIssueWorkLogList200Response) GetResponseOk() (*DescribeIssueWorkLogList200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeAllProjectsIssueWorkLogList200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeIssueWorkLogList200ResponseResponse and assigns it to the Response field.
func (o *DescribeAllProjectsIssueWorkLogList200Response) SetResponse(v DescribeIssueWorkLogList200ResponseResponse) {
	o.Response = &v
}

func (o DescribeAllProjectsIssueWorkLogList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeAllProjectsIssueWorkLogList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeAllProjectsIssueWorkLogList200Response struct {
	value *DescribeAllProjectsIssueWorkLogList200Response
	isSet bool
}

func (v NullableDescribeAllProjectsIssueWorkLogList200Response) Get() *DescribeAllProjectsIssueWorkLogList200Response {
	return v.value
}

func (v *NullableDescribeAllProjectsIssueWorkLogList200Response) Set(val *DescribeAllProjectsIssueWorkLogList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeAllProjectsIssueWorkLogList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeAllProjectsIssueWorkLogList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeAllProjectsIssueWorkLogList200Response(val *DescribeAllProjectsIssueWorkLogList200Response) *NullableDescribeAllProjectsIssueWorkLogList200Response {
	return &NullableDescribeAllProjectsIssueWorkLogList200Response{value: val, isSet: true}
}

func (v NullableDescribeAllProjectsIssueWorkLogList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeAllProjectsIssueWorkLogList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


