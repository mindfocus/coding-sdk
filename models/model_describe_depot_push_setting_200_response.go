/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeDepotPushSetting200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeDepotPushSetting200Response{}

// DescribeDepotPushSetting200Response struct for DescribeDepotPushSetting200Response
type DescribeDepotPushSetting200Response struct {
	Response *ModifyDepotPushSetting200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeDepotPushSetting200Response instantiates a new DescribeDepotPushSetting200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeDepotPushSetting200Response() *DescribeDepotPushSetting200Response {
	this := DescribeDepotPushSetting200Response{}
	return &this
}

// NewDescribeDepotPushSetting200ResponseWithDefaults instantiates a new DescribeDepotPushSetting200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeDepotPushSetting200ResponseWithDefaults() *DescribeDepotPushSetting200Response {
	this := DescribeDepotPushSetting200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeDepotPushSetting200Response) GetResponse() ModifyDepotPushSetting200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret ModifyDepotPushSetting200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotPushSetting200Response) GetResponseOk() (*ModifyDepotPushSetting200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeDepotPushSetting200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ModifyDepotPushSetting200ResponseResponse and assigns it to the Response field.
func (o *DescribeDepotPushSetting200Response) SetResponse(v ModifyDepotPushSetting200ResponseResponse) {
	o.Response = &v
}

func (o DescribeDepotPushSetting200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeDepotPushSetting200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeDepotPushSetting200Response struct {
	value *DescribeDepotPushSetting200Response
	isSet bool
}

func (v NullableDescribeDepotPushSetting200Response) Get() *DescribeDepotPushSetting200Response {
	return v.value
}

func (v *NullableDescribeDepotPushSetting200Response) Set(val *DescribeDepotPushSetting200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeDepotPushSetting200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeDepotPushSetting200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeDepotPushSetting200Response(val *DescribeDepotPushSetting200Response) *NullableDescribeDepotPushSetting200Response {
	return &NullableDescribeDepotPushSetting200Response{value: val, isSet: true}
}

func (v NullableDescribeDepotPushSetting200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeDepotPushSetting200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


