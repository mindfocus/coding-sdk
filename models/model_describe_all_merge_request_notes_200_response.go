/*
CODING OPEN API

  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeAllMergeRequestNotes200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeAllMergeRequestNotes200Response{}

// DescribeAllMergeRequestNotes200Response struct for DescribeAllMergeRequestNotes200Response
type DescribeAllMergeRequestNotes200Response struct {
	Response *DescribeAllMergeRequestNotes200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeAllMergeRequestNotes200Response instantiates a new DescribeAllMergeRequestNotes200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeAllMergeRequestNotes200Response() *DescribeAllMergeRequestNotes200Response {
	this := DescribeAllMergeRequestNotes200Response{}
	return &this
}

// NewDescribeAllMergeRequestNotes200ResponseWithDefaults instantiates a new DescribeAllMergeRequestNotes200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeAllMergeRequestNotes200ResponseWithDefaults() *DescribeAllMergeRequestNotes200Response {
	this := DescribeAllMergeRequestNotes200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeAllMergeRequestNotes200Response) GetResponse() DescribeAllMergeRequestNotes200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeAllMergeRequestNotes200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeAllMergeRequestNotes200Response) GetResponseOk() (*DescribeAllMergeRequestNotes200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeAllMergeRequestNotes200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeAllMergeRequestNotes200ResponseResponse and assigns it to the Response field.
func (o *DescribeAllMergeRequestNotes200Response) SetResponse(v DescribeAllMergeRequestNotes200ResponseResponse) {
	o.Response = &v
}

func (o DescribeAllMergeRequestNotes200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeAllMergeRequestNotes200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeAllMergeRequestNotes200Response struct {
	value *DescribeAllMergeRequestNotes200Response
	isSet bool
}

func (v NullableDescribeAllMergeRequestNotes200Response) Get() *DescribeAllMergeRequestNotes200Response {
	return v.value
}

func (v *NullableDescribeAllMergeRequestNotes200Response) Set(val *DescribeAllMergeRequestNotes200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeAllMergeRequestNotes200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeAllMergeRequestNotes200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeAllMergeRequestNotes200Response(val *DescribeAllMergeRequestNotes200Response) *NullableDescribeAllMergeRequestNotes200Response {
	return &NullableDescribeAllMergeRequestNotes200Response{value: val, isSet: true}
}

func (v NullableDescribeAllMergeRequestNotes200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeAllMergeRequestNotes200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


