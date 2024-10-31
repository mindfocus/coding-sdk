/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CreateAttachmentPrepareSignUrl200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateAttachmentPrepareSignUrl200Response{}

// CreateAttachmentPrepareSignUrl200Response struct for CreateAttachmentPrepareSignUrl200Response
type CreateAttachmentPrepareSignUrl200Response struct {
	Response *CreateAttachmentPrepareSignUrl200ResponseResponse `json:"Response,omitempty"`
}

// NewCreateAttachmentPrepareSignUrl200Response instantiates a new CreateAttachmentPrepareSignUrl200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAttachmentPrepareSignUrl200Response() *CreateAttachmentPrepareSignUrl200Response {
	this := CreateAttachmentPrepareSignUrl200Response{}
	return &this
}

// NewCreateAttachmentPrepareSignUrl200ResponseWithDefaults instantiates a new CreateAttachmentPrepareSignUrl200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAttachmentPrepareSignUrl200ResponseWithDefaults() *CreateAttachmentPrepareSignUrl200Response {
	this := CreateAttachmentPrepareSignUrl200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *CreateAttachmentPrepareSignUrl200Response) GetResponse() CreateAttachmentPrepareSignUrl200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret CreateAttachmentPrepareSignUrl200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAttachmentPrepareSignUrl200Response) GetResponseOk() (*CreateAttachmentPrepareSignUrl200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *CreateAttachmentPrepareSignUrl200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given CreateAttachmentPrepareSignUrl200ResponseResponse and assigns it to the Response field.
func (o *CreateAttachmentPrepareSignUrl200Response) SetResponse(v CreateAttachmentPrepareSignUrl200ResponseResponse) {
	o.Response = &v
}

func (o CreateAttachmentPrepareSignUrl200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAttachmentPrepareSignUrl200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableCreateAttachmentPrepareSignUrl200Response struct {
	value *CreateAttachmentPrepareSignUrl200Response
	isSet bool
}

func (v NullableCreateAttachmentPrepareSignUrl200Response) Get() *CreateAttachmentPrepareSignUrl200Response {
	return v.value
}

func (v *NullableCreateAttachmentPrepareSignUrl200Response) Set(val *CreateAttachmentPrepareSignUrl200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAttachmentPrepareSignUrl200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAttachmentPrepareSignUrl200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAttachmentPrepareSignUrl200Response(val *CreateAttachmentPrepareSignUrl200Response) *NullableCreateAttachmentPrepareSignUrl200Response {
	return &NullableCreateAttachmentPrepareSignUrl200Response{value: val, isSet: true}
}

func (v NullableCreateAttachmentPrepareSignUrl200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAttachmentPrepareSignUrl200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


