/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeReleaseList200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeReleaseList200Response{}

// DescribeReleaseList200Response struct for DescribeReleaseList200Response
type DescribeReleaseList200Response struct {
	Response *DescribeReleaseList200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeReleaseList200Response instantiates a new DescribeReleaseList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeReleaseList200Response() *DescribeReleaseList200Response {
	this := DescribeReleaseList200Response{}
	return &this
}

// NewDescribeReleaseList200ResponseWithDefaults instantiates a new DescribeReleaseList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeReleaseList200ResponseWithDefaults() *DescribeReleaseList200Response {
	this := DescribeReleaseList200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeReleaseList200Response) GetResponse() DescribeReleaseList200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeReleaseList200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseList200Response) GetResponseOk() (*DescribeReleaseList200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeReleaseList200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeReleaseList200ResponseResponse and assigns it to the Response field.
func (o *DescribeReleaseList200Response) SetResponse(v DescribeReleaseList200ResponseResponse) {
	o.Response = &v
}

func (o DescribeReleaseList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeReleaseList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeReleaseList200Response struct {
	value *DescribeReleaseList200Response
	isSet bool
}

func (v NullableDescribeReleaseList200Response) Get() *DescribeReleaseList200Response {
	return v.value
}

func (v *NullableDescribeReleaseList200Response) Set(val *DescribeReleaseList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeReleaseList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeReleaseList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeReleaseList200Response(val *DescribeReleaseList200Response) *NullableDescribeReleaseList200Response {
	return &NullableDescribeReleaseList200Response{value: val, isSet: true}
}

func (v NullableDescribeReleaseList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeReleaseList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


