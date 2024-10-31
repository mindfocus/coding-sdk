/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCdDeployCountByProject200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCdDeployCountByProject200Response{}

// DescribeCdDeployCountByProject200Response struct for DescribeCdDeployCountByProject200Response
type DescribeCdDeployCountByProject200Response struct {
	Response *DescribeCdDeployCountByProject200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeCdDeployCountByProject200Response instantiates a new DescribeCdDeployCountByProject200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCdDeployCountByProject200Response() *DescribeCdDeployCountByProject200Response {
	this := DescribeCdDeployCountByProject200Response{}
	return &this
}

// NewDescribeCdDeployCountByProject200ResponseWithDefaults instantiates a new DescribeCdDeployCountByProject200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCdDeployCountByProject200ResponseWithDefaults() *DescribeCdDeployCountByProject200Response {
	this := DescribeCdDeployCountByProject200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeCdDeployCountByProject200Response) GetResponse() DescribeCdDeployCountByProject200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeCdDeployCountByProject200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdDeployCountByProject200Response) GetResponseOk() (*DescribeCdDeployCountByProject200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeCdDeployCountByProject200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeCdDeployCountByProject200ResponseResponse and assigns it to the Response field.
func (o *DescribeCdDeployCountByProject200Response) SetResponse(v DescribeCdDeployCountByProject200ResponseResponse) {
	o.Response = &v
}

func (o DescribeCdDeployCountByProject200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCdDeployCountByProject200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeCdDeployCountByProject200Response struct {
	value *DescribeCdDeployCountByProject200Response
	isSet bool
}

func (v NullableDescribeCdDeployCountByProject200Response) Get() *DescribeCdDeployCountByProject200Response {
	return v.value
}

func (v *NullableDescribeCdDeployCountByProject200Response) Set(val *DescribeCdDeployCountByProject200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCdDeployCountByProject200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCdDeployCountByProject200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCdDeployCountByProject200Response(val *DescribeCdDeployCountByProject200Response) *NullableDescribeCdDeployCountByProject200Response {
	return &NullableDescribeCdDeployCountByProject200Response{value: val, isSet: true}
}

func (v NullableDescribeCdDeployCountByProject200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCdDeployCountByProject200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


