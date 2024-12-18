/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ModifyWiki200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyWiki200Response{}

// ModifyWiki200Response struct for ModifyWiki200Response
type ModifyWiki200Response struct {
	Response *ModifyWiki200ResponseResponse `json:"Response,omitempty"`
}

// NewModifyWiki200Response instantiates a new ModifyWiki200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyWiki200Response() *ModifyWiki200Response {
	this := ModifyWiki200Response{}
	return &this
}

// NewModifyWiki200ResponseWithDefaults instantiates a new ModifyWiki200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyWiki200ResponseWithDefaults() *ModifyWiki200Response {
	this := ModifyWiki200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ModifyWiki200Response) GetResponse() ModifyWiki200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret ModifyWiki200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyWiki200Response) GetResponseOk() (*ModifyWiki200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ModifyWiki200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ModifyWiki200ResponseResponse and assigns it to the Response field.
func (o *ModifyWiki200Response) SetResponse(v ModifyWiki200ResponseResponse) {
	o.Response = &v
}

func (o ModifyWiki200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyWiki200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableModifyWiki200Response struct {
	value *ModifyWiki200Response
	isSet bool
}

func (v NullableModifyWiki200Response) Get() *ModifyWiki200Response {
	return v.value
}

func (v *NullableModifyWiki200Response) Set(val *ModifyWiki200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyWiki200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyWiki200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyWiki200Response(val *ModifyWiki200Response) *NullableModifyWiki200Response {
	return &NullableModifyWiki200Response{value: val, isSet: true}
}

func (v NullableModifyWiki200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyWiki200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


