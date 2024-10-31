/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeMergeReqCommits200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeMergeReqCommits200Response{}

// DescribeMergeReqCommits200Response struct for DescribeMergeReqCommits200Response
type DescribeMergeReqCommits200Response struct {
	Response *DescribeMergeReqCommits200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeMergeReqCommits200Response instantiates a new DescribeMergeReqCommits200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeMergeReqCommits200Response() *DescribeMergeReqCommits200Response {
	this := DescribeMergeReqCommits200Response{}
	return &this
}

// NewDescribeMergeReqCommits200ResponseWithDefaults instantiates a new DescribeMergeReqCommits200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeMergeReqCommits200ResponseWithDefaults() *DescribeMergeReqCommits200Response {
	this := DescribeMergeReqCommits200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeMergeReqCommits200Response) GetResponse() DescribeMergeReqCommits200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeMergeReqCommits200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeMergeReqCommits200Response) GetResponseOk() (*DescribeMergeReqCommits200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeMergeReqCommits200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeMergeReqCommits200ResponseResponse and assigns it to the Response field.
func (o *DescribeMergeReqCommits200Response) SetResponse(v DescribeMergeReqCommits200ResponseResponse) {
	o.Response = &v
}

func (o DescribeMergeReqCommits200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeMergeReqCommits200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeMergeReqCommits200Response struct {
	value *DescribeMergeReqCommits200Response
	isSet bool
}

func (v NullableDescribeMergeReqCommits200Response) Get() *DescribeMergeReqCommits200Response {
	return v.value
}

func (v *NullableDescribeMergeReqCommits200Response) Set(val *DescribeMergeReqCommits200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeMergeReqCommits200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeMergeReqCommits200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeMergeReqCommits200Response(val *DescribeMergeReqCommits200Response) *NullableDescribeMergeReqCommits200Response {
	return &NullableDescribeMergeReqCommits200Response{value: val, isSet: true}
}

func (v NullableDescribeMergeReqCommits200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeMergeReqCommits200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


