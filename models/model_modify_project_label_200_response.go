/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ModifyProjectLabel200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyProjectLabel200Response{}

// ModifyProjectLabel200Response struct for ModifyProjectLabel200Response
type ModifyProjectLabel200Response struct {
	Response *ModifyProjectLabel200ResponseResponse `json:"Response,omitempty"`
}

// NewModifyProjectLabel200Response instantiates a new ModifyProjectLabel200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyProjectLabel200Response() *ModifyProjectLabel200Response {
	this := ModifyProjectLabel200Response{}
	return &this
}

// NewModifyProjectLabel200ResponseWithDefaults instantiates a new ModifyProjectLabel200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyProjectLabel200ResponseWithDefaults() *ModifyProjectLabel200Response {
	this := ModifyProjectLabel200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ModifyProjectLabel200Response) GetResponse() ModifyProjectLabel200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret ModifyProjectLabel200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyProjectLabel200Response) GetResponseOk() (*ModifyProjectLabel200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ModifyProjectLabel200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ModifyProjectLabel200ResponseResponse and assigns it to the Response field.
func (o *ModifyProjectLabel200Response) SetResponse(v ModifyProjectLabel200ResponseResponse) {
	o.Response = &v
}

func (o ModifyProjectLabel200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyProjectLabel200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableModifyProjectLabel200Response struct {
	value *ModifyProjectLabel200Response
	isSet bool
}

func (v NullableModifyProjectLabel200Response) Get() *ModifyProjectLabel200Response {
	return v.value
}

func (v *NullableModifyProjectLabel200Response) Set(val *ModifyProjectLabel200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyProjectLabel200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyProjectLabel200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyProjectLabel200Response(val *ModifyProjectLabel200Response) *NullableModifyProjectLabel200Response {
	return &NullableModifyProjectLabel200Response{value: val, isSet: true}
}

func (v NullableModifyProjectLabel200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyProjectLabel200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


