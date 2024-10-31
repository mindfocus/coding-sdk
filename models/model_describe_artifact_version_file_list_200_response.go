/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeArtifactVersionFileList200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeArtifactVersionFileList200Response{}

// DescribeArtifactVersionFileList200Response struct for DescribeArtifactVersionFileList200Response
type DescribeArtifactVersionFileList200Response struct {
	Response *DescribeArtifactVersionFileList200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeArtifactVersionFileList200Response instantiates a new DescribeArtifactVersionFileList200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeArtifactVersionFileList200Response() *DescribeArtifactVersionFileList200Response {
	this := DescribeArtifactVersionFileList200Response{}
	return &this
}

// NewDescribeArtifactVersionFileList200ResponseWithDefaults instantiates a new DescribeArtifactVersionFileList200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeArtifactVersionFileList200ResponseWithDefaults() *DescribeArtifactVersionFileList200Response {
	this := DescribeArtifactVersionFileList200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeArtifactVersionFileList200Response) GetResponse() DescribeArtifactVersionFileList200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeArtifactVersionFileList200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactVersionFileList200Response) GetResponseOk() (*DescribeArtifactVersionFileList200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeArtifactVersionFileList200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeArtifactVersionFileList200ResponseResponse and assigns it to the Response field.
func (o *DescribeArtifactVersionFileList200Response) SetResponse(v DescribeArtifactVersionFileList200ResponseResponse) {
	o.Response = &v
}

func (o DescribeArtifactVersionFileList200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeArtifactVersionFileList200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeArtifactVersionFileList200Response struct {
	value *DescribeArtifactVersionFileList200Response
	isSet bool
}

func (v NullableDescribeArtifactVersionFileList200Response) Get() *DescribeArtifactVersionFileList200Response {
	return v.value
}

func (v *NullableDescribeArtifactVersionFileList200Response) Set(val *DescribeArtifactVersionFileList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeArtifactVersionFileList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeArtifactVersionFileList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeArtifactVersionFileList200Response(val *DescribeArtifactVersionFileList200Response) *NullableDescribeArtifactVersionFileList200Response {
	return &NullableDescribeArtifactVersionFileList200Response{value: val, isSet: true}
}

func (v NullableDescribeArtifactVersionFileList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeArtifactVersionFileList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

