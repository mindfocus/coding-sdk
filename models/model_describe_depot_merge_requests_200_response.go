/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeDepotMergeRequests200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeDepotMergeRequests200Response{}

// DescribeDepotMergeRequests200Response struct for DescribeDepotMergeRequests200Response
type DescribeDepotMergeRequests200Response struct {
	Response *DescribeDepotMergeRequests200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeDepotMergeRequests200Response instantiates a new DescribeDepotMergeRequests200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeDepotMergeRequests200Response() *DescribeDepotMergeRequests200Response {
	this := DescribeDepotMergeRequests200Response{}
	return &this
}

// NewDescribeDepotMergeRequests200ResponseWithDefaults instantiates a new DescribeDepotMergeRequests200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeDepotMergeRequests200ResponseWithDefaults() *DescribeDepotMergeRequests200Response {
	this := DescribeDepotMergeRequests200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequests200Response) GetResponse() DescribeDepotMergeRequests200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeDepotMergeRequests200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequests200Response) GetResponseOk() (*DescribeDepotMergeRequests200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequests200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeDepotMergeRequests200ResponseResponse and assigns it to the Response field.
func (o *DescribeDepotMergeRequests200Response) SetResponse(v DescribeDepotMergeRequests200ResponseResponse) {
	o.Response = &v
}

func (o DescribeDepotMergeRequests200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeDepotMergeRequests200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeDepotMergeRequests200Response struct {
	value *DescribeDepotMergeRequests200Response
	isSet bool
}

func (v NullableDescribeDepotMergeRequests200Response) Get() *DescribeDepotMergeRequests200Response {
	return v.value
}

func (v *NullableDescribeDepotMergeRequests200Response) Set(val *DescribeDepotMergeRequests200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeDepotMergeRequests200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeDepotMergeRequests200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeDepotMergeRequests200Response(val *DescribeDepotMergeRequests200Response) *NullableDescribeDepotMergeRequests200Response {
	return &NullableDescribeDepotMergeRequests200Response{value: val, isSet: true}
}

func (v NullableDescribeDepotMergeRequests200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeDepotMergeRequests200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

