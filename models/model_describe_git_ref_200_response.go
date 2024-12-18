/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeGitRef200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitRef200Response{}

// DescribeGitRef200Response struct for DescribeGitRef200Response
type DescribeGitRef200Response struct {
	Response *DescribeGitRef200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeGitRef200Response instantiates a new DescribeGitRef200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitRef200Response() *DescribeGitRef200Response {
	this := DescribeGitRef200Response{}
	return &this
}

// NewDescribeGitRef200ResponseWithDefaults instantiates a new DescribeGitRef200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitRef200ResponseWithDefaults() *DescribeGitRef200Response {
	this := DescribeGitRef200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeGitRef200Response) GetResponse() DescribeGitRef200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeGitRef200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitRef200Response) GetResponseOk() (*DescribeGitRef200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeGitRef200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeGitRef200ResponseResponse and assigns it to the Response field.
func (o *DescribeGitRef200Response) SetResponse(v DescribeGitRef200ResponseResponse) {
	o.Response = &v
}

func (o DescribeGitRef200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitRef200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeGitRef200Response struct {
	value *DescribeGitRef200Response
	isSet bool
}

func (v NullableDescribeGitRef200Response) Get() *DescribeGitRef200Response {
	return v.value
}

func (v *NullableDescribeGitRef200Response) Set(val *DescribeGitRef200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitRef200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitRef200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitRef200Response(val *DescribeGitRef200Response) *NullableDescribeGitRef200Response {
	return &NullableDescribeGitRef200Response{value: val, isSet: true}
}

func (v NullableDescribeGitRef200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitRef200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


