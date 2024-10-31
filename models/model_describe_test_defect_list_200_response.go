/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeTestDefectList200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeTestDefectList200Response{}

// DescribeTestDefectList200Response struct for DescribeTestDefectList200Response
type DescribeTestDefectList200Response struct {
	Response *DescribeTestDefectList200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeTestDefectList200Response instantiates a new DescribeTestDefectList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeTestDefectList200Response() *DescribeTestDefectList200Response {
	this := DescribeTestDefectList200Response{}
	return &this
}

// NewDescribeTestDefectList200ResponseWithDefaults instantiates a new DescribeTestDefectList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeTestDefectList200ResponseWithDefaults() *DescribeTestDefectList200Response {
	this := DescribeTestDefectList200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeTestDefectList200Response) GetResponse() DescribeTestDefectList200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeTestDefectList200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTestDefectList200Response) GetResponseOk() (*DescribeTestDefectList200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeTestDefectList200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeTestDefectList200ResponseResponse and assigns it to the Response field.
func (o *DescribeTestDefectList200Response) SetResponse(v DescribeTestDefectList200ResponseResponse) {
	o.Response = &v
}

func (o DescribeTestDefectList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeTestDefectList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeTestDefectList200Response struct {
	value *DescribeTestDefectList200Response
	isSet bool
}

func (v NullableDescribeTestDefectList200Response) Get() *DescribeTestDefectList200Response {
	return v.value
}

func (v *NullableDescribeTestDefectList200Response) Set(val *DescribeTestDefectList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeTestDefectList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeTestDefectList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeTestDefectList200Response(val *DescribeTestDefectList200Response) *NullableDescribeTestDefectList200Response {
	return &NullableDescribeTestDefectList200Response{value: val, isSet: true}
}

func (v NullableDescribeTestDefectList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeTestDefectList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

