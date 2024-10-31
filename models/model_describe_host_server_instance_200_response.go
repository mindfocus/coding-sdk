/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeHostServerInstance200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeHostServerInstance200Response{}

// DescribeHostServerInstance200Response struct for DescribeHostServerInstance200Response
type DescribeHostServerInstance200Response struct {
	Response *DescribeHostServerInstance200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeHostServerInstance200Response instantiates a new DescribeHostServerInstance200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeHostServerInstance200Response() *DescribeHostServerInstance200Response {
	this := DescribeHostServerInstance200Response{}
	return &this
}

// NewDescribeHostServerInstance200ResponseWithDefaults instantiates a new DescribeHostServerInstance200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeHostServerInstance200ResponseWithDefaults() *DescribeHostServerInstance200Response {
	this := DescribeHostServerInstance200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeHostServerInstance200Response) GetResponse() DescribeHostServerInstance200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeHostServerInstance200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeHostServerInstance200Response) GetResponseOk() (*DescribeHostServerInstance200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeHostServerInstance200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeHostServerInstance200ResponseResponse and assigns it to the Response field.
func (o *DescribeHostServerInstance200Response) SetResponse(v DescribeHostServerInstance200ResponseResponse) {
	o.Response = &v
}

func (o DescribeHostServerInstance200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeHostServerInstance200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeHostServerInstance200Response struct {
	value *DescribeHostServerInstance200Response
	isSet bool
}

func (v NullableDescribeHostServerInstance200Response) Get() *DescribeHostServerInstance200Response {
	return v.value
}

func (v *NullableDescribeHostServerInstance200Response) Set(val *DescribeHostServerInstance200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeHostServerInstance200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeHostServerInstance200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeHostServerInstance200Response(val *DescribeHostServerInstance200Response) *NullableDescribeHostServerInstance200Response {
	return &NullableDescribeHostServerInstance200Response{value: val, isSet: true}
}

func (v NullableDescribeHostServerInstance200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeHostServerInstance200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


