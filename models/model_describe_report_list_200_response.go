/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeReportList200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeReportList200Response{}

// DescribeReportList200Response struct for DescribeReportList200Response
type DescribeReportList200Response struct {
	Response *DescribeReportList200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeReportList200Response instantiates a new DescribeReportList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeReportList200Response() *DescribeReportList200Response {
	this := DescribeReportList200Response{}
	return &this
}

// NewDescribeReportList200ResponseWithDefaults instantiates a new DescribeReportList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeReportList200ResponseWithDefaults() *DescribeReportList200Response {
	this := DescribeReportList200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeReportList200Response) GetResponse() DescribeReportList200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeReportList200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReportList200Response) GetResponseOk() (*DescribeReportList200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeReportList200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeReportList200ResponseResponse and assigns it to the Response field.
func (o *DescribeReportList200Response) SetResponse(v DescribeReportList200ResponseResponse) {
	o.Response = &v
}

func (o DescribeReportList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeReportList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeReportList200Response struct {
	value *DescribeReportList200Response
	isSet bool
}

func (v NullableDescribeReportList200Response) Get() *DescribeReportList200Response {
	return v.value
}

func (v *NullableDescribeReportList200Response) Set(val *DescribeReportList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeReportList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeReportList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeReportList200Response(val *DescribeReportList200Response) *NullableDescribeReportList200Response {
	return &NullableDescribeReportList200Response{value: val, isSet: true}
}

func (v NullableDescribeReportList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeReportList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

