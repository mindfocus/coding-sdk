/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeArtifactProperties200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeArtifactProperties200Response{}

// DescribeArtifactProperties200Response struct for DescribeArtifactProperties200Response
type DescribeArtifactProperties200Response struct {
	Response *DescribeArtifactProperties200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeArtifactProperties200Response instantiates a new DescribeArtifactProperties200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeArtifactProperties200Response() *DescribeArtifactProperties200Response {
	this := DescribeArtifactProperties200Response{}
	return &this
}

// NewDescribeArtifactProperties200ResponseWithDefaults instantiates a new DescribeArtifactProperties200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeArtifactProperties200ResponseWithDefaults() *DescribeArtifactProperties200Response {
	this := DescribeArtifactProperties200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeArtifactProperties200Response) GetResponse() DescribeArtifactProperties200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeArtifactProperties200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactProperties200Response) GetResponseOk() (*DescribeArtifactProperties200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeArtifactProperties200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeArtifactProperties200ResponseResponse and assigns it to the Response field.
func (o *DescribeArtifactProperties200Response) SetResponse(v DescribeArtifactProperties200ResponseResponse) {
	o.Response = &v
}

func (o DescribeArtifactProperties200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeArtifactProperties200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeArtifactProperties200Response struct {
	value *DescribeArtifactProperties200Response
	isSet bool
}

func (v NullableDescribeArtifactProperties200Response) Get() *DescribeArtifactProperties200Response {
	return v.value
}

func (v *NullableDescribeArtifactProperties200Response) Set(val *DescribeArtifactProperties200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeArtifactProperties200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeArtifactProperties200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeArtifactProperties200Response(val *DescribeArtifactProperties200Response) *NullableDescribeArtifactProperties200Response {
	return &NullableDescribeArtifactProperties200Response{value: val, isSet: true}
}

func (v NullableDescribeArtifactProperties200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeArtifactProperties200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


