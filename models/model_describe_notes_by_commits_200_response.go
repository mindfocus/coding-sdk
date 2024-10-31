/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeNotesByCommits200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeNotesByCommits200Response{}

// DescribeNotesByCommits200Response struct for DescribeNotesByCommits200Response
type DescribeNotesByCommits200Response struct {
	Response *DescribeNotesByCommits200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeNotesByCommits200Response instantiates a new DescribeNotesByCommits200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeNotesByCommits200Response() *DescribeNotesByCommits200Response {
	this := DescribeNotesByCommits200Response{}
	return &this
}

// NewDescribeNotesByCommits200ResponseWithDefaults instantiates a new DescribeNotesByCommits200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeNotesByCommits200ResponseWithDefaults() *DescribeNotesByCommits200Response {
	this := DescribeNotesByCommits200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeNotesByCommits200Response) GetResponse() DescribeNotesByCommits200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeNotesByCommits200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeNotesByCommits200Response) GetResponseOk() (*DescribeNotesByCommits200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeNotesByCommits200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeNotesByCommits200ResponseResponse and assigns it to the Response field.
func (o *DescribeNotesByCommits200Response) SetResponse(v DescribeNotesByCommits200ResponseResponse) {
	o.Response = &v
}

func (o DescribeNotesByCommits200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeNotesByCommits200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeNotesByCommits200Response struct {
	value *DescribeNotesByCommits200Response
	isSet bool
}

func (v NullableDescribeNotesByCommits200Response) Get() *DescribeNotesByCommits200Response {
	return v.value
}

func (v *NullableDescribeNotesByCommits200Response) Set(val *DescribeNotesByCommits200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeNotesByCommits200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeNotesByCommits200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeNotesByCommits200Response(val *DescribeNotesByCommits200Response) *NullableDescribeNotesByCommits200Response {
	return &NullableDescribeNotesByCommits200Response{value: val, isSet: true}
}

func (v NullableDescribeNotesByCommits200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeNotesByCommits200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

