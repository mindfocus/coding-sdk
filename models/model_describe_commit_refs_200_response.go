/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCommitRefs200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCommitRefs200Response{}

// DescribeCommitRefs200Response struct for DescribeCommitRefs200Response
type DescribeCommitRefs200Response struct {
	Response *DescribeCommitRefs200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeCommitRefs200Response instantiates a new DescribeCommitRefs200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCommitRefs200Response() *DescribeCommitRefs200Response {
	this := DescribeCommitRefs200Response{}
	return &this
}

// NewDescribeCommitRefs200ResponseWithDefaults instantiates a new DescribeCommitRefs200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCommitRefs200ResponseWithDefaults() *DescribeCommitRefs200Response {
	this := DescribeCommitRefs200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeCommitRefs200Response) GetResponse() DescribeCommitRefs200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeCommitRefs200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCommitRefs200Response) GetResponseOk() (*DescribeCommitRefs200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeCommitRefs200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeCommitRefs200ResponseResponse and assigns it to the Response field.
func (o *DescribeCommitRefs200Response) SetResponse(v DescribeCommitRefs200ResponseResponse) {
	o.Response = &v
}

func (o DescribeCommitRefs200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCommitRefs200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeCommitRefs200Response struct {
	value *DescribeCommitRefs200Response
	isSet bool
}

func (v NullableDescribeCommitRefs200Response) Get() *DescribeCommitRefs200Response {
	return v.value
}

func (v *NullableDescribeCommitRefs200Response) Set(val *DescribeCommitRefs200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCommitRefs200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCommitRefs200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCommitRefs200Response(val *DescribeCommitRefs200Response) *NullableDescribeCommitRefs200Response {
	return &NullableDescribeCommitRefs200Response{value: val, isSet: true}
}

func (v NullableDescribeCommitRefs200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCommitRefs200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


