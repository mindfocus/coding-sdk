/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ModifyGitMergeRequest200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyGitMergeRequest200Response{}

// ModifyGitMergeRequest200Response struct for ModifyGitMergeRequest200Response
type ModifyGitMergeRequest200Response struct {
	Response *ModifyMergeMR200ResponseResponse `json:"Response,omitempty"`
}

// NewModifyGitMergeRequest200Response instantiates a new ModifyGitMergeRequest200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyGitMergeRequest200Response() *ModifyGitMergeRequest200Response {
	this := ModifyGitMergeRequest200Response{}
	return &this
}

// NewModifyGitMergeRequest200ResponseWithDefaults instantiates a new ModifyGitMergeRequest200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyGitMergeRequest200ResponseWithDefaults() *ModifyGitMergeRequest200Response {
	this := ModifyGitMergeRequest200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ModifyGitMergeRequest200Response) GetResponse() ModifyMergeMR200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret ModifyMergeMR200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyGitMergeRequest200Response) GetResponseOk() (*ModifyMergeMR200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ModifyGitMergeRequest200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ModifyMergeMR200ResponseResponse and assigns it to the Response field.
func (o *ModifyGitMergeRequest200Response) SetResponse(v ModifyMergeMR200ResponseResponse) {
	o.Response = &v
}

func (o ModifyGitMergeRequest200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyGitMergeRequest200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableModifyGitMergeRequest200Response struct {
	value *ModifyGitMergeRequest200Response
	isSet bool
}

func (v NullableModifyGitMergeRequest200Response) Get() *ModifyGitMergeRequest200Response {
	return v.value
}

func (v *NullableModifyGitMergeRequest200Response) Set(val *ModifyGitMergeRequest200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyGitMergeRequest200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyGitMergeRequest200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyGitMergeRequest200Response(val *ModifyGitMergeRequest200Response) *NullableModifyGitMergeRequest200Response {
	return &NullableModifyGitMergeRequest200Response{value: val, isSet: true}
}

func (v NullableModifyGitMergeRequest200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyGitMergeRequest200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


