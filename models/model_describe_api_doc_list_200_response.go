/*
CODING OPEN API

  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeAPIDocList200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeAPIDocList200Response{}

// DescribeAPIDocList200Response struct for DescribeAPIDocList200Response
type DescribeAPIDocList200Response struct {
	Response *DescribeAPIDocList200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeAPIDocList200Response instantiates a new DescribeAPIDocList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeAPIDocList200Response() *DescribeAPIDocList200Response {
	this := DescribeAPIDocList200Response{}
	return &this
}

// NewDescribeAPIDocList200ResponseWithDefaults instantiates a new DescribeAPIDocList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeAPIDocList200ResponseWithDefaults() *DescribeAPIDocList200Response {
	this := DescribeAPIDocList200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeAPIDocList200Response) GetResponse() DescribeAPIDocList200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeAPIDocList200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAPIDocList200Response) GetResponseOk() (*DescribeAPIDocList200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeAPIDocList200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeAPIDocList200ResponseResponse and assigns it to the Response field.
func (o *DescribeAPIDocList200Response) SetResponse(v DescribeAPIDocList200ResponseResponse) {
	o.Response = &v
}

func (o DescribeAPIDocList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeAPIDocList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeAPIDocList200Response struct {
	value *DescribeAPIDocList200Response
	isSet bool
}

func (v NullableDescribeAPIDocList200Response) Get() *DescribeAPIDocList200Response {
	return v.value
}

func (v *NullableDescribeAPIDocList200Response) Set(val *DescribeAPIDocList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeAPIDocList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeAPIDocList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeAPIDocList200Response(val *DescribeAPIDocList200Response) *NullableDescribeAPIDocList200Response {
	return &NullableDescribeAPIDocList200Response{value: val, isSet: true}
}

func (v NullableDescribeAPIDocList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeAPIDocList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


