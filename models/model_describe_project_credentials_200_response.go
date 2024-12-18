/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeProjectCredentials200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeProjectCredentials200Response{}

// DescribeProjectCredentials200Response struct for DescribeProjectCredentials200Response
type DescribeProjectCredentials200Response struct {
	Response *DescribeProjectCredentials200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeProjectCredentials200Response instantiates a new DescribeProjectCredentials200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeProjectCredentials200Response() *DescribeProjectCredentials200Response {
	this := DescribeProjectCredentials200Response{}
	return &this
}

// NewDescribeProjectCredentials200ResponseWithDefaults instantiates a new DescribeProjectCredentials200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeProjectCredentials200ResponseWithDefaults() *DescribeProjectCredentials200Response {
	this := DescribeProjectCredentials200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeProjectCredentials200Response) GetResponse() DescribeProjectCredentials200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeProjectCredentials200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectCredentials200Response) GetResponseOk() (*DescribeProjectCredentials200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeProjectCredentials200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeProjectCredentials200ResponseResponse and assigns it to the Response field.
func (o *DescribeProjectCredentials200Response) SetResponse(v DescribeProjectCredentials200ResponseResponse) {
	o.Response = &v
}

func (o DescribeProjectCredentials200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeProjectCredentials200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeProjectCredentials200Response struct {
	value *DescribeProjectCredentials200Response
	isSet bool
}

func (v NullableDescribeProjectCredentials200Response) Get() *DescribeProjectCredentials200Response {
	return v.value
}

func (v *NullableDescribeProjectCredentials200Response) Set(val *DescribeProjectCredentials200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeProjectCredentials200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeProjectCredentials200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeProjectCredentials200Response(val *DescribeProjectCredentials200Response) *NullableDescribeProjectCredentials200Response {
	return &NullableDescribeProjectCredentials200Response{value: val, isSet: true}
}

func (v NullableDescribeProjectCredentials200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeProjectCredentials200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


