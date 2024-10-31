/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeProjectDepotBranches200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeProjectDepotBranches200Response{}

// DescribeProjectDepotBranches200Response struct for DescribeProjectDepotBranches200Response
type DescribeProjectDepotBranches200Response struct {
	Response *DescribeProjectDepotBranches200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeProjectDepotBranches200Response instantiates a new DescribeProjectDepotBranches200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeProjectDepotBranches200Response() *DescribeProjectDepotBranches200Response {
	this := DescribeProjectDepotBranches200Response{}
	return &this
}

// NewDescribeProjectDepotBranches200ResponseWithDefaults instantiates a new DescribeProjectDepotBranches200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeProjectDepotBranches200ResponseWithDefaults() *DescribeProjectDepotBranches200Response {
	this := DescribeProjectDepotBranches200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeProjectDepotBranches200Response) GetResponse() DescribeProjectDepotBranches200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeProjectDepotBranches200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectDepotBranches200Response) GetResponseOk() (*DescribeProjectDepotBranches200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeProjectDepotBranches200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeProjectDepotBranches200ResponseResponse and assigns it to the Response field.
func (o *DescribeProjectDepotBranches200Response) SetResponse(v DescribeProjectDepotBranches200ResponseResponse) {
	o.Response = &v
}

func (o DescribeProjectDepotBranches200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeProjectDepotBranches200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeProjectDepotBranches200Response struct {
	value *DescribeProjectDepotBranches200Response
	isSet bool
}

func (v NullableDescribeProjectDepotBranches200Response) Get() *DescribeProjectDepotBranches200Response {
	return v.value
}

func (v *NullableDescribeProjectDepotBranches200Response) Set(val *DescribeProjectDepotBranches200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeProjectDepotBranches200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeProjectDepotBranches200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeProjectDepotBranches200Response(val *DescribeProjectDepotBranches200Response) *NullableDescribeProjectDepotBranches200Response {
	return &NullableDescribeProjectDepotBranches200Response{value: val, isSet: true}
}

func (v NullableDescribeProjectDepotBranches200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeProjectDepotBranches200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

