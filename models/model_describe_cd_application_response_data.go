/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCdApplicationResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCdApplicationResponseData{}

// DescribeCdApplicationResponseData DescribeCdApplicationResponseData 结构
type DescribeCdApplicationResponseData struct {
	// 应用 JSON 配置
	ApplicationJsonContent *string `json:"ApplicationJsonContent,omitempty"`
}

// NewDescribeCdApplicationResponseData instantiates a new DescribeCdApplicationResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCdApplicationResponseData() *DescribeCdApplicationResponseData {
	this := DescribeCdApplicationResponseData{}
	var applicationJsonContent string = ""
	this.ApplicationJsonContent = &applicationJsonContent
	return &this
}

// NewDescribeCdApplicationResponseDataWithDefaults instantiates a new DescribeCdApplicationResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCdApplicationResponseDataWithDefaults() *DescribeCdApplicationResponseData {
	this := DescribeCdApplicationResponseData{}
	var applicationJsonContent string = ""
	this.ApplicationJsonContent = &applicationJsonContent
	return &this
}

// GetApplicationJsonContent returns the ApplicationJsonContent field value if set, zero value otherwise.
func (o *DescribeCdApplicationResponseData) GetApplicationJsonContent() string {
	if o == nil || utils.IsNil(o.ApplicationJsonContent) {
		var ret string
		return ret
	}
	return *o.ApplicationJsonContent
}

// GetApplicationJsonContentOk returns a tuple with the ApplicationJsonContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdApplicationResponseData) GetApplicationJsonContentOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ApplicationJsonContent) {
		return nil, false
	}
	return o.ApplicationJsonContent, true
}

// HasApplicationJsonContent returns a boolean if a field has been set.
func (o *DescribeCdApplicationResponseData) HasApplicationJsonContent() bool {
	if o != nil && !utils.IsNil(o.ApplicationJsonContent) {
		return true
	}

	return false
}

// SetApplicationJsonContent gets a reference to the given string and assigns it to the ApplicationJsonContent field.
func (o *DescribeCdApplicationResponseData) SetApplicationJsonContent(v string) {
	o.ApplicationJsonContent = &v
}

func (o DescribeCdApplicationResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCdApplicationResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ApplicationJsonContent) {
		toSerialize["ApplicationJsonContent"] = o.ApplicationJsonContent
	}
	return toSerialize, nil
}

type NullableDescribeCdApplicationResponseData struct {
	value *DescribeCdApplicationResponseData
	isSet bool
}

func (v NullableDescribeCdApplicationResponseData) Get() *DescribeCdApplicationResponseData {
	return v.value
}

func (v *NullableDescribeCdApplicationResponseData) Set(val *DescribeCdApplicationResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCdApplicationResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCdApplicationResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCdApplicationResponseData(val *DescribeCdApplicationResponseData) *NullableDescribeCdApplicationResponseData {
	return &NullableDescribeCdApplicationResponseData{value: val, isSet: true}
}

func (v NullableDescribeCdApplicationResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCdApplicationResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

