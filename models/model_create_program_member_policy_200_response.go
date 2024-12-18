/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CreateProgramMemberPolicy200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateProgramMemberPolicy200Response{}

// CreateProgramMemberPolicy200Response struct for CreateProgramMemberPolicy200Response
type CreateProgramMemberPolicy200Response struct {
	Response *CreateProgramMemberPolicy200ResponseResponse `json:"Response,omitempty"`
}

// NewCreateProgramMemberPolicy200Response instantiates a new CreateProgramMemberPolicy200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProgramMemberPolicy200Response() *CreateProgramMemberPolicy200Response {
	this := CreateProgramMemberPolicy200Response{}
	return &this
}

// NewCreateProgramMemberPolicy200ResponseWithDefaults instantiates a new CreateProgramMemberPolicy200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProgramMemberPolicy200ResponseWithDefaults() *CreateProgramMemberPolicy200Response {
	this := CreateProgramMemberPolicy200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *CreateProgramMemberPolicy200Response) GetResponse() CreateProgramMemberPolicy200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret CreateProgramMemberPolicy200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProgramMemberPolicy200Response) GetResponseOk() (*CreateProgramMemberPolicy200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *CreateProgramMemberPolicy200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given CreateProgramMemberPolicy200ResponseResponse and assigns it to the Response field.
func (o *CreateProgramMemberPolicy200Response) SetResponse(v CreateProgramMemberPolicy200ResponseResponse) {
	o.Response = &v
}

func (o CreateProgramMemberPolicy200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProgramMemberPolicy200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableCreateProgramMemberPolicy200Response struct {
	value *CreateProgramMemberPolicy200Response
	isSet bool
}

func (v NullableCreateProgramMemberPolicy200Response) Get() *CreateProgramMemberPolicy200Response {
	return v.value
}

func (v *NullableCreateProgramMemberPolicy200Response) Set(val *CreateProgramMemberPolicy200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProgramMemberPolicy200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProgramMemberPolicy200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProgramMemberPolicy200Response(val *CreateProgramMemberPolicy200Response) *NullableCreateProgramMemberPolicy200Response {
	return &NullableCreateProgramMemberPolicy200Response{value: val, isSet: true}
}

func (v NullableCreateProgramMemberPolicy200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProgramMemberPolicy200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


