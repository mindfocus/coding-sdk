/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCdApplicationsResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCdApplicationsResponseData{}

// DescribeCdApplicationsResponseData DescribeCdApplicationsResponseData 结构
type DescribeCdApplicationsResponseData struct {
	// CD 应用列表
	Applications []CdApplication `json:"Applications,omitempty"`
}

// NewDescribeCdApplicationsResponseData instantiates a new DescribeCdApplicationsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCdApplicationsResponseData() *DescribeCdApplicationsResponseData {
	this := DescribeCdApplicationsResponseData{}
	return &this
}

// NewDescribeCdApplicationsResponseDataWithDefaults instantiates a new DescribeCdApplicationsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCdApplicationsResponseDataWithDefaults() *DescribeCdApplicationsResponseData {
	this := DescribeCdApplicationsResponseData{}
	return &this
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *DescribeCdApplicationsResponseData) GetApplications() []CdApplication {
	if o == nil || utils.IsNil(o.Applications) {
		var ret []CdApplication
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdApplicationsResponseData) GetApplicationsOk() ([]CdApplication, bool) {
	if o == nil || utils.IsNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *DescribeCdApplicationsResponseData) HasApplications() bool {
	if o != nil && !utils.IsNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []CdApplication and assigns it to the Applications field.
func (o *DescribeCdApplicationsResponseData) SetApplications(v []CdApplication) {
	o.Applications = v
}

func (o DescribeCdApplicationsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCdApplicationsResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Applications) {
		toSerialize["Applications"] = o.Applications
	}
	return toSerialize, nil
}

type NullableDescribeCdApplicationsResponseData struct {
	value *DescribeCdApplicationsResponseData
	isSet bool
}

func (v NullableDescribeCdApplicationsResponseData) Get() *DescribeCdApplicationsResponseData {
	return v.value
}

func (v *NullableDescribeCdApplicationsResponseData) Set(val *DescribeCdApplicationsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCdApplicationsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCdApplicationsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCdApplicationsResponseData(val *DescribeCdApplicationsResponseData) *NullableDescribeCdApplicationsResponseData {
	return &NullableDescribeCdApplicationsResponseData{value: val, isSet: true}
}

func (v NullableDescribeCdApplicationsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCdApplicationsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

