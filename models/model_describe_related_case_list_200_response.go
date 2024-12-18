/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeRelatedCaseList200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeRelatedCaseList200Response{}

// DescribeRelatedCaseList200Response struct for DescribeRelatedCaseList200Response
type DescribeRelatedCaseList200Response struct {
	Response *DescribeRelatedCaseList200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeRelatedCaseList200Response instantiates a new DescribeRelatedCaseList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeRelatedCaseList200Response() *DescribeRelatedCaseList200Response {
	this := DescribeRelatedCaseList200Response{}
	return &this
}

// NewDescribeRelatedCaseList200ResponseWithDefaults instantiates a new DescribeRelatedCaseList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeRelatedCaseList200ResponseWithDefaults() *DescribeRelatedCaseList200Response {
	this := DescribeRelatedCaseList200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeRelatedCaseList200Response) GetResponse() DescribeRelatedCaseList200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeRelatedCaseList200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeRelatedCaseList200Response) GetResponseOk() (*DescribeRelatedCaseList200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeRelatedCaseList200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeRelatedCaseList200ResponseResponse and assigns it to the Response field.
func (o *DescribeRelatedCaseList200Response) SetResponse(v DescribeRelatedCaseList200ResponseResponse) {
	o.Response = &v
}

func (o DescribeRelatedCaseList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeRelatedCaseList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeRelatedCaseList200Response struct {
	value *DescribeRelatedCaseList200Response
	isSet bool
}

func (v NullableDescribeRelatedCaseList200Response) Get() *DescribeRelatedCaseList200Response {
	return v.value
}

func (v *NullableDescribeRelatedCaseList200Response) Set(val *DescribeRelatedCaseList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeRelatedCaseList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeRelatedCaseList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeRelatedCaseList200Response(val *DescribeRelatedCaseList200Response) *NullableDescribeRelatedCaseList200Response {
	return &NullableDescribeRelatedCaseList200Response{value: val, isSet: true}
}

func (v NullableDescribeRelatedCaseList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeRelatedCaseList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


