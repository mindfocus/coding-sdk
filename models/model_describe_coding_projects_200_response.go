/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCodingProjects200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCodingProjects200Response{}

// DescribeCodingProjects200Response struct for DescribeCodingProjects200Response
type DescribeCodingProjects200Response struct {
	Response *DescribeCodingProjects200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeCodingProjects200Response instantiates a new DescribeCodingProjects200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCodingProjects200Response() *DescribeCodingProjects200Response {
	this := DescribeCodingProjects200Response{}
	return &this
}

// NewDescribeCodingProjects200ResponseWithDefaults instantiates a new DescribeCodingProjects200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCodingProjects200ResponseWithDefaults() *DescribeCodingProjects200Response {
	this := DescribeCodingProjects200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeCodingProjects200Response) GetResponse() DescribeCodingProjects200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeCodingProjects200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCodingProjects200Response) GetResponseOk() (*DescribeCodingProjects200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeCodingProjects200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeCodingProjects200ResponseResponse and assigns it to the Response field.
func (o *DescribeCodingProjects200Response) SetResponse(v DescribeCodingProjects200ResponseResponse) {
	o.Response = &v
}

func (o DescribeCodingProjects200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCodingProjects200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeCodingProjects200Response struct {
	value *DescribeCodingProjects200Response
	isSet bool
}

func (v NullableDescribeCodingProjects200Response) Get() *DescribeCodingProjects200Response {
	return v.value
}

func (v *NullableDescribeCodingProjects200Response) Set(val *DescribeCodingProjects200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCodingProjects200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCodingProjects200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCodingProjects200Response(val *DescribeCodingProjects200Response) *NullableDescribeCodingProjects200Response {
	return &NullableDescribeCodingProjects200Response{value: val, isSet: true}
}

func (v NullableDescribeCodingProjects200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCodingProjects200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


