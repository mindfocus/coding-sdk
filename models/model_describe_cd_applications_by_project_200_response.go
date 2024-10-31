/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCdApplicationsByProject200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCdApplicationsByProject200Response{}

// DescribeCdApplicationsByProject200Response struct for DescribeCdApplicationsByProject200Response
type DescribeCdApplicationsByProject200Response struct {
	Response *DescribeCdApplicationsByProject200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeCdApplicationsByProject200Response instantiates a new DescribeCdApplicationsByProject200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCdApplicationsByProject200Response() *DescribeCdApplicationsByProject200Response {
	this := DescribeCdApplicationsByProject200Response{}
	return &this
}

// NewDescribeCdApplicationsByProject200ResponseWithDefaults instantiates a new DescribeCdApplicationsByProject200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCdApplicationsByProject200ResponseWithDefaults() *DescribeCdApplicationsByProject200Response {
	this := DescribeCdApplicationsByProject200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeCdApplicationsByProject200Response) GetResponse() DescribeCdApplicationsByProject200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeCdApplicationsByProject200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdApplicationsByProject200Response) GetResponseOk() (*DescribeCdApplicationsByProject200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeCdApplicationsByProject200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeCdApplicationsByProject200ResponseResponse and assigns it to the Response field.
func (o *DescribeCdApplicationsByProject200Response) SetResponse(v DescribeCdApplicationsByProject200ResponseResponse) {
	o.Response = &v
}

func (o DescribeCdApplicationsByProject200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCdApplicationsByProject200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeCdApplicationsByProject200Response struct {
	value *DescribeCdApplicationsByProject200Response
	isSet bool
}

func (v NullableDescribeCdApplicationsByProject200Response) Get() *DescribeCdApplicationsByProject200Response {
	return v.value
}

func (v *NullableDescribeCdApplicationsByProject200Response) Set(val *DescribeCdApplicationsByProject200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCdApplicationsByProject200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCdApplicationsByProject200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCdApplicationsByProject200Response(val *DescribeCdApplicationsByProject200Response) *NullableDescribeCdApplicationsByProject200Response {
	return &NullableDescribeCdApplicationsByProject200Response{value: val, isSet: true}
}

func (v NullableDescribeCdApplicationsByProject200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCdApplicationsByProject200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


