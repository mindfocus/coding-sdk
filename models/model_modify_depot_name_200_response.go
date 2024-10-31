/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ModifyDepotName200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyDepotName200Response{}

// ModifyDepotName200Response struct for ModifyDepotName200Response
type ModifyDepotName200Response struct {
	Response *ModifyDepotName200ResponseResponse `json:"Response,omitempty"`
}

// NewModifyDepotName200Response instantiates a new ModifyDepotName200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyDepotName200Response() *ModifyDepotName200Response {
	this := ModifyDepotName200Response{}
	return &this
}

// NewModifyDepotName200ResponseWithDefaults instantiates a new ModifyDepotName200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyDepotName200ResponseWithDefaults() *ModifyDepotName200Response {
	this := ModifyDepotName200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ModifyDepotName200Response) GetResponse() ModifyDepotName200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret ModifyDepotName200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyDepotName200Response) GetResponseOk() (*ModifyDepotName200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ModifyDepotName200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ModifyDepotName200ResponseResponse and assigns it to the Response field.
func (o *ModifyDepotName200Response) SetResponse(v ModifyDepotName200ResponseResponse) {
	o.Response = &v
}

func (o ModifyDepotName200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyDepotName200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableModifyDepotName200Response struct {
	value *ModifyDepotName200Response
	isSet bool
}

func (v NullableModifyDepotName200Response) Get() *ModifyDepotName200Response {
	return v.value
}

func (v *NullableModifyDepotName200Response) Set(val *ModifyDepotName200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyDepotName200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyDepotName200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyDepotName200Response(val *ModifyDepotName200Response) *NullableModifyDepotName200Response {
	return &NullableModifyDepotName200Response{value: val, isSet: true}
}

func (v NullableModifyDepotName200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyDepotName200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

