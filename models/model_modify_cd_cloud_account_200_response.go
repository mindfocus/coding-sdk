/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ModifyCdCloudAccount200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyCdCloudAccount200Response{}

// ModifyCdCloudAccount200Response struct for ModifyCdCloudAccount200Response
type ModifyCdCloudAccount200Response struct {
	Response *ModifyCdCloudAccount200ResponseResponse `json:"Response,omitempty"`
}

// NewModifyCdCloudAccount200Response instantiates a new ModifyCdCloudAccount200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyCdCloudAccount200Response() *ModifyCdCloudAccount200Response {
	this := ModifyCdCloudAccount200Response{}
	return &this
}

// NewModifyCdCloudAccount200ResponseWithDefaults instantiates a new ModifyCdCloudAccount200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyCdCloudAccount200ResponseWithDefaults() *ModifyCdCloudAccount200Response {
	this := ModifyCdCloudAccount200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ModifyCdCloudAccount200Response) GetResponse() ModifyCdCloudAccount200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret ModifyCdCloudAccount200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyCdCloudAccount200Response) GetResponseOk() (*ModifyCdCloudAccount200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ModifyCdCloudAccount200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ModifyCdCloudAccount200ResponseResponse and assigns it to the Response field.
func (o *ModifyCdCloudAccount200Response) SetResponse(v ModifyCdCloudAccount200ResponseResponse) {
	o.Response = &v
}

func (o ModifyCdCloudAccount200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyCdCloudAccount200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableModifyCdCloudAccount200Response struct {
	value *ModifyCdCloudAccount200Response
	isSet bool
}

func (v NullableModifyCdCloudAccount200Response) Get() *ModifyCdCloudAccount200Response {
	return v.value
}

func (v *NullableModifyCdCloudAccount200Response) Set(val *ModifyCdCloudAccount200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyCdCloudAccount200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyCdCloudAccount200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyCdCloudAccount200Response(val *ModifyCdCloudAccount200Response) *NullableModifyCdCloudAccount200Response {
	return &NullableModifyCdCloudAccount200Response{value: val, isSet: true}
}

func (v NullableModifyCdCloudAccount200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyCdCloudAccount200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


