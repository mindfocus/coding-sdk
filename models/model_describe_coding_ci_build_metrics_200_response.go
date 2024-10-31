/*
CODING OPEN API

  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCodingCIBuildMetrics200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCodingCIBuildMetrics200Response{}

// DescribeCodingCIBuildMetrics200Response struct for DescribeCodingCIBuildMetrics200Response
type DescribeCodingCIBuildMetrics200Response struct {
	Response *DescribeCodingCIBuildMetrics200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeCodingCIBuildMetrics200Response instantiates a new DescribeCodingCIBuildMetrics200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCodingCIBuildMetrics200Response() *DescribeCodingCIBuildMetrics200Response {
	this := DescribeCodingCIBuildMetrics200Response{}
	return &this
}

// NewDescribeCodingCIBuildMetrics200ResponseWithDefaults instantiates a new DescribeCodingCIBuildMetrics200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCodingCIBuildMetrics200ResponseWithDefaults() *DescribeCodingCIBuildMetrics200Response {
	this := DescribeCodingCIBuildMetrics200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeCodingCIBuildMetrics200Response) GetResponse() DescribeCodingCIBuildMetrics200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeCodingCIBuildMetrics200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIBuildMetrics200Response) GetResponseOk() (*DescribeCodingCIBuildMetrics200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeCodingCIBuildMetrics200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeCodingCIBuildMetrics200ResponseResponse and assigns it to the Response field.
func (o *DescribeCodingCIBuildMetrics200Response) SetResponse(v DescribeCodingCIBuildMetrics200ResponseResponse) {
	o.Response = &v
}

func (o DescribeCodingCIBuildMetrics200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCodingCIBuildMetrics200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeCodingCIBuildMetrics200Response struct {
	value *DescribeCodingCIBuildMetrics200Response
	isSet bool
}

func (v NullableDescribeCodingCIBuildMetrics200Response) Get() *DescribeCodingCIBuildMetrics200Response {
	return v.value
}

func (v *NullableDescribeCodingCIBuildMetrics200Response) Set(val *DescribeCodingCIBuildMetrics200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCodingCIBuildMetrics200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCodingCIBuildMetrics200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCodingCIBuildMetrics200Response(val *DescribeCodingCIBuildMetrics200Response) *NullableDescribeCodingCIBuildMetrics200Response {
	return &NullableDescribeCodingCIBuildMetrics200Response{value: val, isSet: true}
}

func (v NullableDescribeCodingCIBuildMetrics200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCodingCIBuildMetrics200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

