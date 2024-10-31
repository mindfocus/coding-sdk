/*
CODING OPEN API

  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeGitBranches200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitBranches200Response{}

// DescribeGitBranches200Response struct for DescribeGitBranches200Response
type DescribeGitBranches200Response struct {
	Response *DescribeGitBranches200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeGitBranches200Response instantiates a new DescribeGitBranches200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitBranches200Response() *DescribeGitBranches200Response {
	this := DescribeGitBranches200Response{}
	return &this
}

// NewDescribeGitBranches200ResponseWithDefaults instantiates a new DescribeGitBranches200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitBranches200ResponseWithDefaults() *DescribeGitBranches200Response {
	this := DescribeGitBranches200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeGitBranches200Response) GetResponse() DescribeGitBranches200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeGitBranches200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitBranches200Response) GetResponseOk() (*DescribeGitBranches200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeGitBranches200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeGitBranches200ResponseResponse and assigns it to the Response field.
func (o *DescribeGitBranches200Response) SetResponse(v DescribeGitBranches200ResponseResponse) {
	o.Response = &v
}

func (o DescribeGitBranches200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitBranches200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeGitBranches200Response struct {
	value *DescribeGitBranches200Response
	isSet bool
}

func (v NullableDescribeGitBranches200Response) Get() *DescribeGitBranches200Response {
	return v.value
}

func (v *NullableDescribeGitBranches200Response) Set(val *DescribeGitBranches200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitBranches200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitBranches200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitBranches200Response(val *DescribeGitBranches200Response) *NullableDescribeGitBranches200Response {
	return &NullableDescribeGitBranches200Response{value: val, isSet: true}
}

func (v NullableDescribeGitBranches200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitBranches200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

