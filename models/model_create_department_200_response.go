/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CreateDepartment200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateDepartment200Response{}

// CreateDepartment200Response struct for CreateDepartment200Response
type CreateDepartment200Response struct {
	Response *ModifyDepartment200ResponseResponse `json:"Response,omitempty"`
}

// NewCreateDepartment200Response instantiates a new CreateDepartment200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDepartment200Response() *CreateDepartment200Response {
	this := CreateDepartment200Response{}
	return &this
}

// NewCreateDepartment200ResponseWithDefaults instantiates a new CreateDepartment200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDepartment200ResponseWithDefaults() *CreateDepartment200Response {
	this := CreateDepartment200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *CreateDepartment200Response) GetResponse() ModifyDepartment200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret ModifyDepartment200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDepartment200Response) GetResponseOk() (*ModifyDepartment200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *CreateDepartment200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ModifyDepartment200ResponseResponse and assigns it to the Response field.
func (o *CreateDepartment200Response) SetResponse(v ModifyDepartment200ResponseResponse) {
	o.Response = &v
}

func (o CreateDepartment200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDepartment200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableCreateDepartment200Response struct {
	value *CreateDepartment200Response
	isSet bool
}

func (v NullableCreateDepartment200Response) Get() *CreateDepartment200Response {
	return v.value
}

func (v *NullableCreateDepartment200Response) Set(val *CreateDepartment200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDepartment200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDepartment200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDepartment200Response(val *CreateDepartment200Response) *NullableCreateDepartment200Response {
	return &NullableCreateDepartment200Response{value: val, isSet: true}
}

func (v NullableCreateDepartment200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDepartment200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


