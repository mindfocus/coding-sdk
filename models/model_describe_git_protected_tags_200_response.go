/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeGitProtectedTags200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitProtectedTags200Response{}

// DescribeGitProtectedTags200Response struct for DescribeGitProtectedTags200Response
type DescribeGitProtectedTags200Response struct {
	Response *DescribeGitProtectedTags200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeGitProtectedTags200Response instantiates a new DescribeGitProtectedTags200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitProtectedTags200Response() *DescribeGitProtectedTags200Response {
	this := DescribeGitProtectedTags200Response{}
	return &this
}

// NewDescribeGitProtectedTags200ResponseWithDefaults instantiates a new DescribeGitProtectedTags200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitProtectedTags200ResponseWithDefaults() *DescribeGitProtectedTags200Response {
	this := DescribeGitProtectedTags200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeGitProtectedTags200Response) GetResponse() DescribeGitProtectedTags200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeGitProtectedTags200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitProtectedTags200Response) GetResponseOk() (*DescribeGitProtectedTags200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeGitProtectedTags200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeGitProtectedTags200ResponseResponse and assigns it to the Response field.
func (o *DescribeGitProtectedTags200Response) SetResponse(v DescribeGitProtectedTags200ResponseResponse) {
	o.Response = &v
}

func (o DescribeGitProtectedTags200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitProtectedTags200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeGitProtectedTags200Response struct {
	value *DescribeGitProtectedTags200Response
	isSet bool
}

func (v NullableDescribeGitProtectedTags200Response) Get() *DescribeGitProtectedTags200Response {
	return v.value
}

func (v *NullableDescribeGitProtectedTags200Response) Set(val *DescribeGitProtectedTags200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitProtectedTags200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitProtectedTags200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitProtectedTags200Response(val *DescribeGitProtectedTags200Response) *NullableDescribeGitProtectedTags200Response {
	return &NullableDescribeGitProtectedTags200Response{value: val, isSet: true}
}

func (v NullableDescribeGitProtectedTags200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitProtectedTags200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


