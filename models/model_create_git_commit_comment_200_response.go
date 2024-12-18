/*
CODING OPEN API

  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CreateGitCommitComment200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateGitCommitComment200Response{}

// CreateGitCommitComment200Response struct for CreateGitCommitComment200Response
type CreateGitCommitComment200Response struct {
	Response *CreateGitCommitComment200ResponseResponse `json:"Response,omitempty"`
}

// NewCreateGitCommitComment200Response instantiates a new CreateGitCommitComment200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGitCommitComment200Response() *CreateGitCommitComment200Response {
	this := CreateGitCommitComment200Response{}
	return &this
}

// NewCreateGitCommitComment200ResponseWithDefaults instantiates a new CreateGitCommitComment200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGitCommitComment200ResponseWithDefaults() *CreateGitCommitComment200Response {
	this := CreateGitCommitComment200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *CreateGitCommitComment200Response) GetResponse() CreateGitCommitComment200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret CreateGitCommitComment200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGitCommitComment200Response) GetResponseOk() (*CreateGitCommitComment200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *CreateGitCommitComment200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given CreateGitCommitComment200ResponseResponse and assigns it to the Response field.
func (o *CreateGitCommitComment200Response) SetResponse(v CreateGitCommitComment200ResponseResponse) {
	o.Response = &v
}

func (o CreateGitCommitComment200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGitCommitComment200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableCreateGitCommitComment200Response struct {
	value *CreateGitCommitComment200Response
	isSet bool
}

func (v NullableCreateGitCommitComment200Response) Get() *CreateGitCommitComment200Response {
	return v.value
}

func (v *NullableCreateGitCommitComment200Response) Set(val *CreateGitCommitComment200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGitCommitComment200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGitCommitComment200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGitCommitComment200Response(val *CreateGitCommitComment200Response) *NullableCreateGitCommitComment200Response {
	return &NullableCreateGitCommitComment200Response{value: val, isSet: true}
}

func (v NullableCreateGitCommitComment200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGitCommitComment200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


