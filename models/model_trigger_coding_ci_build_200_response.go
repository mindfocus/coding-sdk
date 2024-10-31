/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the TriggerCodingCIBuild200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TriggerCodingCIBuild200Response{}

// TriggerCodingCIBuild200Response struct for TriggerCodingCIBuild200Response
type TriggerCodingCIBuild200Response struct {
	Response *TriggerCodingCIBuild200ResponseResponse `json:"Response,omitempty"`
}

// NewTriggerCodingCIBuild200Response instantiates a new TriggerCodingCIBuild200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerCodingCIBuild200Response() *TriggerCodingCIBuild200Response {
	this := TriggerCodingCIBuild200Response{}
	return &this
}

// NewTriggerCodingCIBuild200ResponseWithDefaults instantiates a new TriggerCodingCIBuild200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerCodingCIBuild200ResponseWithDefaults() *TriggerCodingCIBuild200Response {
	this := TriggerCodingCIBuild200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *TriggerCodingCIBuild200Response) GetResponse() TriggerCodingCIBuild200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret TriggerCodingCIBuild200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerCodingCIBuild200Response) GetResponseOk() (*TriggerCodingCIBuild200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *TriggerCodingCIBuild200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given TriggerCodingCIBuild200ResponseResponse and assigns it to the Response field.
func (o *TriggerCodingCIBuild200Response) SetResponse(v TriggerCodingCIBuild200ResponseResponse) {
	o.Response = &v
}

func (o TriggerCodingCIBuild200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerCodingCIBuild200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableTriggerCodingCIBuild200Response struct {
	value *TriggerCodingCIBuild200Response
	isSet bool
}

func (v NullableTriggerCodingCIBuild200Response) Get() *TriggerCodingCIBuild200Response {
	return v.value
}

func (v *NullableTriggerCodingCIBuild200Response) Set(val *TriggerCodingCIBuild200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerCodingCIBuild200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerCodingCIBuild200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerCodingCIBuild200Response(val *TriggerCodingCIBuild200Response) *NullableTriggerCodingCIBuild200Response {
	return &NullableTriggerCodingCIBuild200Response{value: val, isSet: true}
}

func (v NullableTriggerCodingCIBuild200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerCodingCIBuild200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

