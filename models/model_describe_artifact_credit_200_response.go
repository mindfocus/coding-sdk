/*
CODING OPEN API

  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeArtifactCredit200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeArtifactCredit200Response{}

// DescribeArtifactCredit200Response struct for DescribeArtifactCredit200Response
type DescribeArtifactCredit200Response struct {
	Response *DescribeArtifactCredit200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeArtifactCredit200Response instantiates a new DescribeArtifactCredit200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeArtifactCredit200Response() *DescribeArtifactCredit200Response {
	this := DescribeArtifactCredit200Response{}
	return &this
}

// NewDescribeArtifactCredit200ResponseWithDefaults instantiates a new DescribeArtifactCredit200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeArtifactCredit200ResponseWithDefaults() *DescribeArtifactCredit200Response {
	this := DescribeArtifactCredit200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeArtifactCredit200Response) GetResponse() DescribeArtifactCredit200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeArtifactCredit200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactCredit200Response) GetResponseOk() (*DescribeArtifactCredit200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeArtifactCredit200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeArtifactCredit200ResponseResponse and assigns it to the Response field.
func (o *DescribeArtifactCredit200Response) SetResponse(v DescribeArtifactCredit200ResponseResponse) {
	o.Response = &v
}

func (o DescribeArtifactCredit200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeArtifactCredit200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeArtifactCredit200Response struct {
	value *DescribeArtifactCredit200Response
	isSet bool
}

func (v NullableDescribeArtifactCredit200Response) Get() *DescribeArtifactCredit200Response {
	return v.value
}

func (v *NullableDescribeArtifactCredit200Response) Set(val *DescribeArtifactCredit200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeArtifactCredit200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeArtifactCredit200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeArtifactCredit200Response(val *DescribeArtifactCredit200Response) *NullableDescribeArtifactCredit200Response {
	return &NullableDescribeArtifactCredit200Response{value: val, isSet: true}
}

func (v NullableDescribeArtifactCredit200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeArtifactCredit200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

