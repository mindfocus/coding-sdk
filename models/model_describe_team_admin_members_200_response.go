/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeTeamAdminMembers200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeTeamAdminMembers200Response{}

// DescribeTeamAdminMembers200Response struct for DescribeTeamAdminMembers200Response
type DescribeTeamAdminMembers200Response struct {
	Response *DescribeTeamAdminMembers200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeTeamAdminMembers200Response instantiates a new DescribeTeamAdminMembers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeTeamAdminMembers200Response() *DescribeTeamAdminMembers200Response {
	this := DescribeTeamAdminMembers200Response{}
	return &this
}

// NewDescribeTeamAdminMembers200ResponseWithDefaults instantiates a new DescribeTeamAdminMembers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeTeamAdminMembers200ResponseWithDefaults() *DescribeTeamAdminMembers200Response {
	this := DescribeTeamAdminMembers200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeTeamAdminMembers200Response) GetResponse() DescribeTeamAdminMembers200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeTeamAdminMembers200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTeamAdminMembers200Response) GetResponseOk() (*DescribeTeamAdminMembers200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeTeamAdminMembers200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeTeamAdminMembers200ResponseResponse and assigns it to the Response field.
func (o *DescribeTeamAdminMembers200Response) SetResponse(v DescribeTeamAdminMembers200ResponseResponse) {
	o.Response = &v
}

func (o DescribeTeamAdminMembers200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeTeamAdminMembers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeTeamAdminMembers200Response struct {
	value *DescribeTeamAdminMembers200Response
	isSet bool
}

func (v NullableDescribeTeamAdminMembers200Response) Get() *DescribeTeamAdminMembers200Response {
	return v.value
}

func (v *NullableDescribeTeamAdminMembers200Response) Set(val *DescribeTeamAdminMembers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeTeamAdminMembers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeTeamAdminMembers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeTeamAdminMembers200Response(val *DescribeTeamAdminMembers200Response) *NullableDescribeTeamAdminMembers200Response {
	return &NullableDescribeTeamAdminMembers200Response{value: val, isSet: true}
}

func (v NullableDescribeTeamAdminMembers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeTeamAdminMembers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


