/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CreateServiceHook200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateServiceHook200Response{}

// CreateServiceHook200Response struct for CreateServiceHook200Response
type CreateServiceHook200Response struct {
	Response *ModifyServiceHook200ResponseResponse `json:"Response,omitempty"`
}

// NewCreateServiceHook200Response instantiates a new CreateServiceHook200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateServiceHook200Response() *CreateServiceHook200Response {
	this := CreateServiceHook200Response{}
	return &this
}

// NewCreateServiceHook200ResponseWithDefaults instantiates a new CreateServiceHook200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateServiceHook200ResponseWithDefaults() *CreateServiceHook200Response {
	this := CreateServiceHook200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *CreateServiceHook200Response) GetResponse() ModifyServiceHook200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret ModifyServiceHook200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateServiceHook200Response) GetResponseOk() (*ModifyServiceHook200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *CreateServiceHook200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ModifyServiceHook200ResponseResponse and assigns it to the Response field.
func (o *CreateServiceHook200Response) SetResponse(v ModifyServiceHook200ResponseResponse) {
	o.Response = &v
}

func (o CreateServiceHook200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateServiceHook200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableCreateServiceHook200Response struct {
	value *CreateServiceHook200Response
	isSet bool
}

func (v NullableCreateServiceHook200Response) Get() *CreateServiceHook200Response {
	return v.value
}

func (v *NullableCreateServiceHook200Response) Set(val *CreateServiceHook200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateServiceHook200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateServiceHook200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateServiceHook200Response(val *CreateServiceHook200Response) *NullableCreateServiceHook200Response {
	return &NullableCreateServiceHook200Response{value: val, isSet: true}
}

func (v NullableCreateServiceHook200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateServiceHook200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

