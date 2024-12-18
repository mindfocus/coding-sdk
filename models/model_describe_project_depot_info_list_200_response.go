/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeProjectDepotInfoList200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeProjectDepotInfoList200Response{}

// DescribeProjectDepotInfoList200Response struct for DescribeProjectDepotInfoList200Response
type DescribeProjectDepotInfoList200Response struct {
	Response *DescribeProjectDepotInfoList200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeProjectDepotInfoList200Response instantiates a new DescribeProjectDepotInfoList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeProjectDepotInfoList200Response() *DescribeProjectDepotInfoList200Response {
	this := DescribeProjectDepotInfoList200Response{}
	return &this
}

// NewDescribeProjectDepotInfoList200ResponseWithDefaults instantiates a new DescribeProjectDepotInfoList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeProjectDepotInfoList200ResponseWithDefaults() *DescribeProjectDepotInfoList200Response {
	this := DescribeProjectDepotInfoList200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeProjectDepotInfoList200Response) GetResponse() DescribeProjectDepotInfoList200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeProjectDepotInfoList200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectDepotInfoList200Response) GetResponseOk() (*DescribeProjectDepotInfoList200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeProjectDepotInfoList200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeProjectDepotInfoList200ResponseResponse and assigns it to the Response field.
func (o *DescribeProjectDepotInfoList200Response) SetResponse(v DescribeProjectDepotInfoList200ResponseResponse) {
	o.Response = &v
}

func (o DescribeProjectDepotInfoList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeProjectDepotInfoList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeProjectDepotInfoList200Response struct {
	value *DescribeProjectDepotInfoList200Response
	isSet bool
}

func (v NullableDescribeProjectDepotInfoList200Response) Get() *DescribeProjectDepotInfoList200Response {
	return v.value
}

func (v *NullableDescribeProjectDepotInfoList200Response) Set(val *DescribeProjectDepotInfoList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeProjectDepotInfoList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeProjectDepotInfoList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeProjectDepotInfoList200Response(val *DescribeProjectDepotInfoList200Response) *NullableDescribeProjectDepotInfoList200Response {
	return &NullableDescribeProjectDepotInfoList200Response{value: val, isSet: true}
}

func (v NullableDescribeProjectDepotInfoList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeProjectDepotInfoList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


