/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeGitFiles200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitFiles200Response{}

// DescribeGitFiles200Response struct for DescribeGitFiles200Response
type DescribeGitFiles200Response struct {
	Response *DescribeGitFiles200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeGitFiles200Response instantiates a new DescribeGitFiles200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitFiles200Response() *DescribeGitFiles200Response {
	this := DescribeGitFiles200Response{}
	return &this
}

// NewDescribeGitFiles200ResponseWithDefaults instantiates a new DescribeGitFiles200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitFiles200ResponseWithDefaults() *DescribeGitFiles200Response {
	this := DescribeGitFiles200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeGitFiles200Response) GetResponse() DescribeGitFiles200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeGitFiles200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitFiles200Response) GetResponseOk() (*DescribeGitFiles200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeGitFiles200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeGitFiles200ResponseResponse and assigns it to the Response field.
func (o *DescribeGitFiles200Response) SetResponse(v DescribeGitFiles200ResponseResponse) {
	o.Response = &v
}

func (o DescribeGitFiles200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitFiles200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeGitFiles200Response struct {
	value *DescribeGitFiles200Response
	isSet bool
}

func (v NullableDescribeGitFiles200Response) Get() *DescribeGitFiles200Response {
	return v.value
}

func (v *NullableDescribeGitFiles200Response) Set(val *DescribeGitFiles200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitFiles200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitFiles200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitFiles200Response(val *DescribeGitFiles200Response) *NullableDescribeGitFiles200Response {
	return &NullableDescribeGitFiles200Response{value: val, isSet: true}
}

func (v NullableDescribeGitFiles200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitFiles200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


