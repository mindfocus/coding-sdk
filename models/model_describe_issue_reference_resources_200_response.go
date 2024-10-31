/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeIssueReferenceResources200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeIssueReferenceResources200Response{}

// DescribeIssueReferenceResources200Response struct for DescribeIssueReferenceResources200Response
type DescribeIssueReferenceResources200Response struct {
	Response *DescribeIssueReferenceResources200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeIssueReferenceResources200Response instantiates a new DescribeIssueReferenceResources200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeIssueReferenceResources200Response() *DescribeIssueReferenceResources200Response {
	this := DescribeIssueReferenceResources200Response{}
	return &this
}

// NewDescribeIssueReferenceResources200ResponseWithDefaults instantiates a new DescribeIssueReferenceResources200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeIssueReferenceResources200ResponseWithDefaults() *DescribeIssueReferenceResources200Response {
	this := DescribeIssueReferenceResources200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeIssueReferenceResources200Response) GetResponse() DescribeIssueReferenceResources200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeIssueReferenceResources200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeIssueReferenceResources200Response) GetResponseOk() (*DescribeIssueReferenceResources200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeIssueReferenceResources200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeIssueReferenceResources200ResponseResponse and assigns it to the Response field.
func (o *DescribeIssueReferenceResources200Response) SetResponse(v DescribeIssueReferenceResources200ResponseResponse) {
	o.Response = &v
}

func (o DescribeIssueReferenceResources200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeIssueReferenceResources200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeIssueReferenceResources200Response struct {
	value *DescribeIssueReferenceResources200Response
	isSet bool
}

func (v NullableDescribeIssueReferenceResources200Response) Get() *DescribeIssueReferenceResources200Response {
	return v.value
}

func (v *NullableDescribeIssueReferenceResources200Response) Set(val *DescribeIssueReferenceResources200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeIssueReferenceResources200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeIssueReferenceResources200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeIssueReferenceResources200Response(val *DescribeIssueReferenceResources200Response) *NullableDescribeIssueReferenceResources200Response {
	return &NullableDescribeIssueReferenceResources200Response{value: val, isSet: true}
}

func (v NullableDescribeIssueReferenceResources200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeIssueReferenceResources200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

