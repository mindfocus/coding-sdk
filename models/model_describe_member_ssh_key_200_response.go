/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeMemberSshKey200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeMemberSshKey200Response{}

// DescribeMemberSshKey200Response struct for DescribeMemberSshKey200Response
type DescribeMemberSshKey200Response struct {
	Response *DescribeMemberSshKey200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeMemberSshKey200Response instantiates a new DescribeMemberSshKey200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeMemberSshKey200Response() *DescribeMemberSshKey200Response {
	this := DescribeMemberSshKey200Response{}
	return &this
}

// NewDescribeMemberSshKey200ResponseWithDefaults instantiates a new DescribeMemberSshKey200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeMemberSshKey200ResponseWithDefaults() *DescribeMemberSshKey200Response {
	this := DescribeMemberSshKey200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeMemberSshKey200Response) GetResponse() DescribeMemberSshKey200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeMemberSshKey200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeMemberSshKey200Response) GetResponseOk() (*DescribeMemberSshKey200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeMemberSshKey200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeMemberSshKey200ResponseResponse and assigns it to the Response field.
func (o *DescribeMemberSshKey200Response) SetResponse(v DescribeMemberSshKey200ResponseResponse) {
	o.Response = &v
}

func (o DescribeMemberSshKey200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeMemberSshKey200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeMemberSshKey200Response struct {
	value *DescribeMemberSshKey200Response
	isSet bool
}

func (v NullableDescribeMemberSshKey200Response) Get() *DescribeMemberSshKey200Response {
	return v.value
}

func (v *NullableDescribeMemberSshKey200Response) Set(val *DescribeMemberSshKey200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeMemberSshKey200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeMemberSshKey200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeMemberSshKey200Response(val *DescribeMemberSshKey200Response) *NullableDescribeMemberSshKey200Response {
	return &NullableDescribeMemberSshKey200Response{value: val, isSet: true}
}

func (v NullableDescribeMemberSshKey200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeMemberSshKey200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


