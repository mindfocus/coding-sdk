/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCdDeployCountByApplications200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCdDeployCountByApplications200Response{}

// DescribeCdDeployCountByApplications200Response struct for DescribeCdDeployCountByApplications200Response
type DescribeCdDeployCountByApplications200Response struct {
	Response *DescribeCdDeployCountByApplications200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeCdDeployCountByApplications200Response instantiates a new DescribeCdDeployCountByApplications200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCdDeployCountByApplications200Response() *DescribeCdDeployCountByApplications200Response {
	this := DescribeCdDeployCountByApplications200Response{}
	return &this
}

// NewDescribeCdDeployCountByApplications200ResponseWithDefaults instantiates a new DescribeCdDeployCountByApplications200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCdDeployCountByApplications200ResponseWithDefaults() *DescribeCdDeployCountByApplications200Response {
	this := DescribeCdDeployCountByApplications200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeCdDeployCountByApplications200Response) GetResponse() DescribeCdDeployCountByApplications200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeCdDeployCountByApplications200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdDeployCountByApplications200Response) GetResponseOk() (*DescribeCdDeployCountByApplications200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeCdDeployCountByApplications200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeCdDeployCountByApplications200ResponseResponse and assigns it to the Response field.
func (o *DescribeCdDeployCountByApplications200Response) SetResponse(v DescribeCdDeployCountByApplications200ResponseResponse) {
	o.Response = &v
}

func (o DescribeCdDeployCountByApplications200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCdDeployCountByApplications200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeCdDeployCountByApplications200Response struct {
	value *DescribeCdDeployCountByApplications200Response
	isSet bool
}

func (v NullableDescribeCdDeployCountByApplications200Response) Get() *DescribeCdDeployCountByApplications200Response {
	return v.value
}

func (v *NullableDescribeCdDeployCountByApplications200Response) Set(val *DescribeCdDeployCountByApplications200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCdDeployCountByApplications200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCdDeployCountByApplications200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCdDeployCountByApplications200Response(val *DescribeCdDeployCountByApplications200Response) *NullableDescribeCdDeployCountByApplications200Response {
	return &NullableDescribeCdDeployCountByApplications200Response{value: val, isSet: true}
}

func (v NullableDescribeCdDeployCountByApplications200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCdDeployCountByApplications200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


