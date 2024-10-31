/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeGitMergeBase200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitMergeBase200Response{}

// DescribeGitMergeBase200Response struct for DescribeGitMergeBase200Response
type DescribeGitMergeBase200Response struct {
	Response *DescribeGitMergeBase200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeGitMergeBase200Response instantiates a new DescribeGitMergeBase200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitMergeBase200Response() *DescribeGitMergeBase200Response {
	this := DescribeGitMergeBase200Response{}
	return &this
}

// NewDescribeGitMergeBase200ResponseWithDefaults instantiates a new DescribeGitMergeBase200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitMergeBase200ResponseWithDefaults() *DescribeGitMergeBase200Response {
	this := DescribeGitMergeBase200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeGitMergeBase200Response) GetResponse() DescribeGitMergeBase200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeGitMergeBase200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitMergeBase200Response) GetResponseOk() (*DescribeGitMergeBase200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeGitMergeBase200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeGitMergeBase200ResponseResponse and assigns it to the Response field.
func (o *DescribeGitMergeBase200Response) SetResponse(v DescribeGitMergeBase200ResponseResponse) {
	o.Response = &v
}

func (o DescribeGitMergeBase200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitMergeBase200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeGitMergeBase200Response struct {
	value *DescribeGitMergeBase200Response
	isSet bool
}

func (v NullableDescribeGitMergeBase200Response) Get() *DescribeGitMergeBase200Response {
	return v.value
}

func (v *NullableDescribeGitMergeBase200Response) Set(val *DescribeGitMergeBase200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitMergeBase200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitMergeBase200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitMergeBase200Response(val *DescribeGitMergeBase200Response) *NullableDescribeGitMergeBase200Response {
	return &NullableDescribeGitMergeBase200Response{value: val, isSet: true}
}

func (v NullableDescribeGitMergeBase200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitMergeBase200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


