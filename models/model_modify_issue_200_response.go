/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ModifyIssue200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyIssue200Response{}

// ModifyIssue200Response struct for ModifyIssue200Response
type ModifyIssue200Response struct {
	Response *ModifyIssue200ResponseResponse `json:"Response,omitempty"`
}

// NewModifyIssue200Response instantiates a new ModifyIssue200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyIssue200Response() *ModifyIssue200Response {
	this := ModifyIssue200Response{}
	return &this
}

// NewModifyIssue200ResponseWithDefaults instantiates a new ModifyIssue200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyIssue200ResponseWithDefaults() *ModifyIssue200Response {
	this := ModifyIssue200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ModifyIssue200Response) GetResponse() ModifyIssue200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret ModifyIssue200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssue200Response) GetResponseOk() (*ModifyIssue200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ModifyIssue200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ModifyIssue200ResponseResponse and assigns it to the Response field.
func (o *ModifyIssue200Response) SetResponse(v ModifyIssue200ResponseResponse) {
	o.Response = &v
}

func (o ModifyIssue200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyIssue200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableModifyIssue200Response struct {
	value *ModifyIssue200Response
	isSet bool
}

func (v NullableModifyIssue200Response) Get() *ModifyIssue200Response {
	return v.value
}

func (v *NullableModifyIssue200Response) Set(val *ModifyIssue200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyIssue200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyIssue200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyIssue200Response(val *ModifyIssue200Response) *NullableModifyIssue200Response {
	return &NullableModifyIssue200Response{value: val, isSet: true}
}

func (v NullableModifyIssue200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyIssue200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

