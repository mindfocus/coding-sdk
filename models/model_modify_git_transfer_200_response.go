/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ModifyGitTransfer200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyGitTransfer200Response{}

// ModifyGitTransfer200Response struct for ModifyGitTransfer200Response
type ModifyGitTransfer200Response struct {
	Response *ModifyGitTransfer200ResponseResponse `json:"Response,omitempty"`
}

// NewModifyGitTransfer200Response instantiates a new ModifyGitTransfer200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyGitTransfer200Response() *ModifyGitTransfer200Response {
	this := ModifyGitTransfer200Response{}
	return &this
}

// NewModifyGitTransfer200ResponseWithDefaults instantiates a new ModifyGitTransfer200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyGitTransfer200ResponseWithDefaults() *ModifyGitTransfer200Response {
	this := ModifyGitTransfer200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ModifyGitTransfer200Response) GetResponse() ModifyGitTransfer200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret ModifyGitTransfer200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyGitTransfer200Response) GetResponseOk() (*ModifyGitTransfer200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ModifyGitTransfer200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ModifyGitTransfer200ResponseResponse and assigns it to the Response field.
func (o *ModifyGitTransfer200Response) SetResponse(v ModifyGitTransfer200ResponseResponse) {
	o.Response = &v
}

func (o ModifyGitTransfer200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyGitTransfer200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableModifyGitTransfer200Response struct {
	value *ModifyGitTransfer200Response
	isSet bool
}

func (v NullableModifyGitTransfer200Response) Get() *ModifyGitTransfer200Response {
	return v.value
}

func (v *NullableModifyGitTransfer200Response) Set(val *ModifyGitTransfer200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyGitTransfer200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyGitTransfer200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyGitTransfer200Response(val *ModifyGitTransfer200Response) *NullableModifyGitTransfer200Response {
	return &NullableModifyGitTransfer200Response{value: val, isSet: true}
}

func (v NullableModifyGitTransfer200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyGitTransfer200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


