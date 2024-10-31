/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeResourceReferencesCiting200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeResourceReferencesCiting200Response{}

// DescribeResourceReferencesCiting200Response struct for DescribeResourceReferencesCiting200Response
type DescribeResourceReferencesCiting200Response struct {
	Response *DescribeResourceReferencesCiting200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeResourceReferencesCiting200Response instantiates a new DescribeResourceReferencesCiting200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeResourceReferencesCiting200Response() *DescribeResourceReferencesCiting200Response {
	this := DescribeResourceReferencesCiting200Response{}
	return &this
}

// NewDescribeResourceReferencesCiting200ResponseWithDefaults instantiates a new DescribeResourceReferencesCiting200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeResourceReferencesCiting200ResponseWithDefaults() *DescribeResourceReferencesCiting200Response {
	this := DescribeResourceReferencesCiting200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeResourceReferencesCiting200Response) GetResponse() DescribeResourceReferencesCiting200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeResourceReferencesCiting200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceReferencesCiting200Response) GetResponseOk() (*DescribeResourceReferencesCiting200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeResourceReferencesCiting200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeResourceReferencesCiting200ResponseResponse and assigns it to the Response field.
func (o *DescribeResourceReferencesCiting200Response) SetResponse(v DescribeResourceReferencesCiting200ResponseResponse) {
	o.Response = &v
}

func (o DescribeResourceReferencesCiting200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeResourceReferencesCiting200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeResourceReferencesCiting200Response struct {
	value *DescribeResourceReferencesCiting200Response
	isSet bool
}

func (v NullableDescribeResourceReferencesCiting200Response) Get() *DescribeResourceReferencesCiting200Response {
	return v.value
}

func (v *NullableDescribeResourceReferencesCiting200Response) Set(val *DescribeResourceReferencesCiting200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeResourceReferencesCiting200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeResourceReferencesCiting200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeResourceReferencesCiting200Response(val *DescribeResourceReferencesCiting200Response) *NullableDescribeResourceReferencesCiting200Response {
	return &NullableDescribeResourceReferencesCiting200Response{value: val, isSet: true}
}

func (v NullableDescribeResourceReferencesCiting200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeResourceReferencesCiting200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


