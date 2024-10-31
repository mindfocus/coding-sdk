/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ReorderCdPipelines200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReorderCdPipelines200Response{}

// ReorderCdPipelines200Response struct for ReorderCdPipelines200Response
type ReorderCdPipelines200Response struct {
	Response *ReorderCdPipelines200ResponseResponse `json:"Response,omitempty"`
}

// NewReorderCdPipelines200Response instantiates a new ReorderCdPipelines200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReorderCdPipelines200Response() *ReorderCdPipelines200Response {
	this := ReorderCdPipelines200Response{}
	return &this
}

// NewReorderCdPipelines200ResponseWithDefaults instantiates a new ReorderCdPipelines200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReorderCdPipelines200ResponseWithDefaults() *ReorderCdPipelines200Response {
	this := ReorderCdPipelines200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ReorderCdPipelines200Response) GetResponse() ReorderCdPipelines200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret ReorderCdPipelines200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReorderCdPipelines200Response) GetResponseOk() (*ReorderCdPipelines200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ReorderCdPipelines200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ReorderCdPipelines200ResponseResponse and assigns it to the Response field.
func (o *ReorderCdPipelines200Response) SetResponse(v ReorderCdPipelines200ResponseResponse) {
	o.Response = &v
}

func (o ReorderCdPipelines200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReorderCdPipelines200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableReorderCdPipelines200Response struct {
	value *ReorderCdPipelines200Response
	isSet bool
}

func (v NullableReorderCdPipelines200Response) Get() *ReorderCdPipelines200Response {
	return v.value
}

func (v *NullableReorderCdPipelines200Response) Set(val *ReorderCdPipelines200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReorderCdPipelines200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReorderCdPipelines200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReorderCdPipelines200Response(val *ReorderCdPipelines200Response) *NullableReorderCdPipelines200Response {
	return &NullableReorderCdPipelines200Response{value: val, isSet: true}
}

func (v NullableReorderCdPipelines200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReorderCdPipelines200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

