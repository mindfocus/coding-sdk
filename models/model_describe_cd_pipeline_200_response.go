/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCdPipeline200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCdPipeline200Response{}

// DescribeCdPipeline200Response struct for DescribeCdPipeline200Response
type DescribeCdPipeline200Response struct {
	Response *DescribeCdPipeline200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeCdPipeline200Response instantiates a new DescribeCdPipeline200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCdPipeline200Response() *DescribeCdPipeline200Response {
	this := DescribeCdPipeline200Response{}
	return &this
}

// NewDescribeCdPipeline200ResponseWithDefaults instantiates a new DescribeCdPipeline200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCdPipeline200ResponseWithDefaults() *DescribeCdPipeline200Response {
	this := DescribeCdPipeline200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeCdPipeline200Response) GetResponse() DescribeCdPipeline200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeCdPipeline200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdPipeline200Response) GetResponseOk() (*DescribeCdPipeline200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeCdPipeline200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeCdPipeline200ResponseResponse and assigns it to the Response field.
func (o *DescribeCdPipeline200Response) SetResponse(v DescribeCdPipeline200ResponseResponse) {
	o.Response = &v
}

func (o DescribeCdPipeline200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCdPipeline200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeCdPipeline200Response struct {
	value *DescribeCdPipeline200Response
	isSet bool
}

func (v NullableDescribeCdPipeline200Response) Get() *DescribeCdPipeline200Response {
	return v.value
}

func (v *NullableDescribeCdPipeline200Response) Set(val *DescribeCdPipeline200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCdPipeline200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCdPipeline200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCdPipeline200Response(val *DescribeCdPipeline200Response) *NullableDescribeCdPipeline200Response {
	return &NullableDescribeCdPipeline200Response{value: val, isSet: true}
}

func (v NullableDescribeCdPipeline200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCdPipeline200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


