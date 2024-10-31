/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DeleteProjectLabel200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DeleteProjectLabel200Response{}

// DeleteProjectLabel200Response struct for DeleteProjectLabel200Response
type DeleteProjectLabel200Response struct {
	Response *EnabledServiceHook200ResponseResponse `json:"Response,omitempty"`
}

// NewDeleteProjectLabel200Response instantiates a new DeleteProjectLabel200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteProjectLabel200Response() *DeleteProjectLabel200Response {
	this := DeleteProjectLabel200Response{}
	return &this
}

// NewDeleteProjectLabel200ResponseWithDefaults instantiates a new DeleteProjectLabel200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteProjectLabel200ResponseWithDefaults() *DeleteProjectLabel200Response {
	this := DeleteProjectLabel200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DeleteProjectLabel200Response) GetResponse() EnabledServiceHook200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret EnabledServiceHook200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteProjectLabel200Response) GetResponseOk() (*EnabledServiceHook200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DeleteProjectLabel200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given EnabledServiceHook200ResponseResponse and assigns it to the Response field.
func (o *DeleteProjectLabel200Response) SetResponse(v EnabledServiceHook200ResponseResponse) {
	o.Response = &v
}

func (o DeleteProjectLabel200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteProjectLabel200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDeleteProjectLabel200Response struct {
	value *DeleteProjectLabel200Response
	isSet bool
}

func (v NullableDeleteProjectLabel200Response) Get() *DeleteProjectLabel200Response {
	return v.value
}

func (v *NullableDeleteProjectLabel200Response) Set(val *DeleteProjectLabel200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteProjectLabel200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteProjectLabel200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteProjectLabel200Response(val *DeleteProjectLabel200Response) *NullableDeleteProjectLabel200Response {
	return &NullableDeleteProjectLabel200Response{value: val, isSet: true}
}

func (v NullableDeleteProjectLabel200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteProjectLabel200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

