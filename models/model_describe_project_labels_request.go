/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeProjectLabelsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeProjectLabelsRequest{}

// DescribeProjectLabelsRequest struct for DescribeProjectLabelsRequest
type DescribeProjectLabelsRequest struct {
	// 标签
	Label *string `json:"Label,omitempty"`
}

// NewDescribeProjectLabelsRequest instantiates a new DescribeProjectLabelsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeProjectLabelsRequest() *DescribeProjectLabelsRequest {
	this := DescribeProjectLabelsRequest{}
	return &this
}

// NewDescribeProjectLabelsRequestWithDefaults instantiates a new DescribeProjectLabelsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeProjectLabelsRequestWithDefaults() *DescribeProjectLabelsRequest {
	this := DescribeProjectLabelsRequest{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DescribeProjectLabelsRequest) GetLabel() string {
	if o == nil || utils.IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectLabelsRequest) GetLabelOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DescribeProjectLabelsRequest) HasLabel() bool {
	if o != nil && !utils.IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *DescribeProjectLabelsRequest) SetLabel(v string) {
	o.Label = &v
}

func (o DescribeProjectLabelsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeProjectLabelsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Label) {
		toSerialize["Label"] = o.Label
	}
	return toSerialize, nil
}

type NullableDescribeProjectLabelsRequest struct {
	value *DescribeProjectLabelsRequest
	isSet bool
}

func (v NullableDescribeProjectLabelsRequest) Get() *DescribeProjectLabelsRequest {
	return v.value
}

func (v *NullableDescribeProjectLabelsRequest) Set(val *DescribeProjectLabelsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeProjectLabelsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeProjectLabelsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeProjectLabelsRequest(val *DescribeProjectLabelsRequest) *NullableDescribeProjectLabelsRequest {
	return &NullableDescribeProjectLabelsRequest{value: val, isSet: true}
}

func (v NullableDescribeProjectLabelsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeProjectLabelsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

