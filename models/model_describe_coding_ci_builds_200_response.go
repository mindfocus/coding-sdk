/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCodingCIBuilds200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCodingCIBuilds200Response{}

// DescribeCodingCIBuilds200Response struct for DescribeCodingCIBuilds200Response
type DescribeCodingCIBuilds200Response struct {
	Response *DescribeCodingCIBuilds200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeCodingCIBuilds200Response instantiates a new DescribeCodingCIBuilds200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCodingCIBuilds200Response() *DescribeCodingCIBuilds200Response {
	this := DescribeCodingCIBuilds200Response{}
	return &this
}

// NewDescribeCodingCIBuilds200ResponseWithDefaults instantiates a new DescribeCodingCIBuilds200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCodingCIBuilds200ResponseWithDefaults() *DescribeCodingCIBuilds200Response {
	this := DescribeCodingCIBuilds200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeCodingCIBuilds200Response) GetResponse() DescribeCodingCIBuilds200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeCodingCIBuilds200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuilds200Response) GetResponseOk() (*DescribeCodingCIBuilds200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeCodingCIBuilds200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeCodingCIBuilds200ResponseResponse and assigns it to the Response field.
func (o *DescribeCodingCIBuilds200Response) SetResponse(v DescribeCodingCIBuilds200ResponseResponse) {
	o.Response = &v
}

func (o DescribeCodingCIBuilds200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCodingCIBuilds200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeCodingCIBuilds200Response struct {
	value *DescribeCodingCIBuilds200Response
	isSet bool
}

func (v NullableDescribeCodingCIBuilds200Response) Get() *DescribeCodingCIBuilds200Response {
	return v.value
}

func (v *NullableDescribeCodingCIBuilds200Response) Set(val *DescribeCodingCIBuilds200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCodingCIBuilds200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCodingCIBuilds200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCodingCIBuilds200Response(val *DescribeCodingCIBuilds200Response) *NullableDescribeCodingCIBuilds200Response {
	return &NullableDescribeCodingCIBuilds200Response{value: val, isSet: true}
}

func (v NullableDescribeCodingCIBuilds200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCodingCIBuilds200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


