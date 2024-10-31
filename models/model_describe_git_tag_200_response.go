/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeGitTag200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitTag200Response{}

// DescribeGitTag200Response struct for DescribeGitTag200Response
type DescribeGitTag200Response struct {
	Response *CreateGitTag200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeGitTag200Response instantiates a new DescribeGitTag200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitTag200Response() *DescribeGitTag200Response {
	this := DescribeGitTag200Response{}
	return &this
}

// NewDescribeGitTag200ResponseWithDefaults instantiates a new DescribeGitTag200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitTag200ResponseWithDefaults() *DescribeGitTag200Response {
	this := DescribeGitTag200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeGitTag200Response) GetResponse() CreateGitTag200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret CreateGitTag200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitTag200Response) GetResponseOk() (*CreateGitTag200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeGitTag200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given CreateGitTag200ResponseResponse and assigns it to the Response field.
func (o *DescribeGitTag200Response) SetResponse(v CreateGitTag200ResponseResponse) {
	o.Response = &v
}

func (o DescribeGitTag200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitTag200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeGitTag200Response struct {
	value *DescribeGitTag200Response
	isSet bool
}

func (v NullableDescribeGitTag200Response) Get() *DescribeGitTag200Response {
	return v.value
}

func (v *NullableDescribeGitTag200Response) Set(val *DescribeGitTag200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitTag200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitTag200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitTag200Response(val *DescribeGitTag200Response) *NullableDescribeGitTag200Response {
	return &NullableDescribeGitTag200Response{value: val, isSet: true}
}

func (v NullableDescribeGitTag200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitTag200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

