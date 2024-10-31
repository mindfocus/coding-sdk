/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ModifyDepartment200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyDepartment200Response{}

// ModifyDepartment200Response struct for ModifyDepartment200Response
type ModifyDepartment200Response struct {
	Response *ModifyDepartment200ResponseResponse `json:"Response,omitempty"`
}

// NewModifyDepartment200Response instantiates a new ModifyDepartment200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyDepartment200Response() *ModifyDepartment200Response {
	this := ModifyDepartment200Response{}
	return &this
}

// NewModifyDepartment200ResponseWithDefaults instantiates a new ModifyDepartment200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyDepartment200ResponseWithDefaults() *ModifyDepartment200Response {
	this := ModifyDepartment200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ModifyDepartment200Response) GetResponse() ModifyDepartment200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret ModifyDepartment200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyDepartment200Response) GetResponseOk() (*ModifyDepartment200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ModifyDepartment200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ModifyDepartment200ResponseResponse and assigns it to the Response field.
func (o *ModifyDepartment200Response) SetResponse(v ModifyDepartment200ResponseResponse) {
	o.Response = &v
}

func (o ModifyDepartment200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyDepartment200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableModifyDepartment200Response struct {
	value *ModifyDepartment200Response
	isSet bool
}

func (v NullableModifyDepartment200Response) Get() *ModifyDepartment200Response {
	return v.value
}

func (v *NullableModifyDepartment200Response) Set(val *ModifyDepartment200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyDepartment200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyDepartment200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyDepartment200Response(val *ModifyDepartment200Response) *NullableModifyDepartment200Response {
	return &NullableModifyDepartment200Response{value: val, isSet: true}
}

func (v NullableModifyDepartment200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyDepartment200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


