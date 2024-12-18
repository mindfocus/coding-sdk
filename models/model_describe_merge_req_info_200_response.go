/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeMergeReqInfo200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeMergeReqInfo200Response{}

// DescribeMergeReqInfo200Response struct for DescribeMergeReqInfo200Response
type DescribeMergeReqInfo200Response struct {
	Response *DescribeMergeReqInfo200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeMergeReqInfo200Response instantiates a new DescribeMergeReqInfo200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeMergeReqInfo200Response() *DescribeMergeReqInfo200Response {
	this := DescribeMergeReqInfo200Response{}
	return &this
}

// NewDescribeMergeReqInfo200ResponseWithDefaults instantiates a new DescribeMergeReqInfo200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeMergeReqInfo200ResponseWithDefaults() *DescribeMergeReqInfo200Response {
	this := DescribeMergeReqInfo200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeMergeReqInfo200Response) GetResponse() DescribeMergeReqInfo200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeMergeReqInfo200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeMergeReqInfo200Response) GetResponseOk() (*DescribeMergeReqInfo200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeMergeReqInfo200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeMergeReqInfo200ResponseResponse and assigns it to the Response field.
func (o *DescribeMergeReqInfo200Response) SetResponse(v DescribeMergeReqInfo200ResponseResponse) {
	o.Response = &v
}

func (o DescribeMergeReqInfo200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeMergeReqInfo200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeMergeReqInfo200Response struct {
	value *DescribeMergeReqInfo200Response
	isSet bool
}

func (v NullableDescribeMergeReqInfo200Response) Get() *DescribeMergeReqInfo200Response {
	return v.value
}

func (v *NullableDescribeMergeReqInfo200Response) Set(val *DescribeMergeReqInfo200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeMergeReqInfo200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeMergeReqInfo200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeMergeReqInfo200Response(val *DescribeMergeReqInfo200Response) *NullableDescribeMergeReqInfo200Response {
	return &NullableDescribeMergeReqInfo200Response{value: val, isSet: true}
}

func (v NullableDescribeMergeReqInfo200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeMergeReqInfo200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


