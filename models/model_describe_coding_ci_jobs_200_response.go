/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCodingCIJobs200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCodingCIJobs200Response{}

// DescribeCodingCIJobs200Response struct for DescribeCodingCIJobs200Response
type DescribeCodingCIJobs200Response struct {
	Response *DescribeCodingCIJobs200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeCodingCIJobs200Response instantiates a new DescribeCodingCIJobs200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCodingCIJobs200Response() *DescribeCodingCIJobs200Response {
	this := DescribeCodingCIJobs200Response{}
	return &this
}

// NewDescribeCodingCIJobs200ResponseWithDefaults instantiates a new DescribeCodingCIJobs200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCodingCIJobs200ResponseWithDefaults() *DescribeCodingCIJobs200Response {
	this := DescribeCodingCIJobs200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeCodingCIJobs200Response) GetResponse() DescribeCodingCIJobs200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeCodingCIJobs200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIJobs200Response) GetResponseOk() (*DescribeCodingCIJobs200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeCodingCIJobs200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeCodingCIJobs200ResponseResponse and assigns it to the Response field.
func (o *DescribeCodingCIJobs200Response) SetResponse(v DescribeCodingCIJobs200ResponseResponse) {
	o.Response = &v
}

func (o DescribeCodingCIJobs200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCodingCIJobs200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeCodingCIJobs200Response struct {
	value *DescribeCodingCIJobs200Response
	isSet bool
}

func (v NullableDescribeCodingCIJobs200Response) Get() *DescribeCodingCIJobs200Response {
	return v.value
}

func (v *NullableDescribeCodingCIJobs200Response) Set(val *DescribeCodingCIJobs200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCodingCIJobs200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCodingCIJobs200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCodingCIJobs200Response(val *DescribeCodingCIJobs200Response) *NullableDescribeCodingCIJobs200Response {
	return &NullableDescribeCodingCIJobs200Response{value: val, isSet: true}
}

func (v NullableDescribeCodingCIJobs200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCodingCIJobs200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


