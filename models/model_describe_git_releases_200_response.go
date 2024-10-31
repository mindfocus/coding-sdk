/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeGitReleases200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitReleases200Response{}

// DescribeGitReleases200Response struct for DescribeGitReleases200Response
type DescribeGitReleases200Response struct {
	Response *DescribeGitReleases200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeGitReleases200Response instantiates a new DescribeGitReleases200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitReleases200Response() *DescribeGitReleases200Response {
	this := DescribeGitReleases200Response{}
	return &this
}

// NewDescribeGitReleases200ResponseWithDefaults instantiates a new DescribeGitReleases200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitReleases200ResponseWithDefaults() *DescribeGitReleases200Response {
	this := DescribeGitReleases200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeGitReleases200Response) GetResponse() DescribeGitReleases200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeGitReleases200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitReleases200Response) GetResponseOk() (*DescribeGitReleases200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeGitReleases200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeGitReleases200ResponseResponse and assigns it to the Response field.
func (o *DescribeGitReleases200Response) SetResponse(v DescribeGitReleases200ResponseResponse) {
	o.Response = &v
}

func (o DescribeGitReleases200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitReleases200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeGitReleases200Response struct {
	value *DescribeGitReleases200Response
	isSet bool
}

func (v NullableDescribeGitReleases200Response) Get() *DescribeGitReleases200Response {
	return v.value
}

func (v *NullableDescribeGitReleases200Response) Set(val *DescribeGitReleases200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitReleases200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitReleases200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitReleases200Response(val *DescribeGitReleases200Response) *NullableDescribeGitReleases200Response {
	return &NullableDescribeGitReleases200Response{value: val, isSet: true}
}

func (v NullableDescribeGitReleases200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitReleases200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


