/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeProjectsByFeature200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeProjectsByFeature200Response{}

// DescribeProjectsByFeature200Response struct for DescribeProjectsByFeature200Response
type DescribeProjectsByFeature200Response struct {
	Response *DescribeProjectsByFeature200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeProjectsByFeature200Response instantiates a new DescribeProjectsByFeature200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeProjectsByFeature200Response() *DescribeProjectsByFeature200Response {
	this := DescribeProjectsByFeature200Response{}
	return &this
}

// NewDescribeProjectsByFeature200ResponseWithDefaults instantiates a new DescribeProjectsByFeature200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeProjectsByFeature200ResponseWithDefaults() *DescribeProjectsByFeature200Response {
	this := DescribeProjectsByFeature200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeProjectsByFeature200Response) GetResponse() DescribeProjectsByFeature200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeProjectsByFeature200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectsByFeature200Response) GetResponseOk() (*DescribeProjectsByFeature200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeProjectsByFeature200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeProjectsByFeature200ResponseResponse and assigns it to the Response field.
func (o *DescribeProjectsByFeature200Response) SetResponse(v DescribeProjectsByFeature200ResponseResponse) {
	o.Response = &v
}

func (o DescribeProjectsByFeature200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeProjectsByFeature200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeProjectsByFeature200Response struct {
	value *DescribeProjectsByFeature200Response
	isSet bool
}

func (v NullableDescribeProjectsByFeature200Response) Get() *DescribeProjectsByFeature200Response {
	return v.value
}

func (v *NullableDescribeProjectsByFeature200Response) Set(val *DescribeProjectsByFeature200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeProjectsByFeature200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeProjectsByFeature200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeProjectsByFeature200Response(val *DescribeProjectsByFeature200Response) *NullableDescribeProjectsByFeature200Response {
	return &NullableDescribeProjectsByFeature200Response{value: val, isSet: true}
}

func (v NullableDescribeProjectsByFeature200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeProjectsByFeature200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

