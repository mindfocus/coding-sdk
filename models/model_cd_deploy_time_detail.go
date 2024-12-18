/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CdDeployTimeDetail type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CdDeployTimeDetail{}

// CdDeployTimeDetail CdDeployTimeDetail 结构
type CdDeployTimeDetail struct {
	// 应用名称
	Application *string `json:"Application,omitempty"`
	CdDeployTime *CdDeployTime `json:"CdDeployTime,omitempty"`
}

// NewCdDeployTimeDetail instantiates a new CdDeployTimeDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdDeployTimeDetail() *CdDeployTimeDetail {
	this := CdDeployTimeDetail{}
	var application string = ""
	this.Application = &application
	return &this
}

// NewCdDeployTimeDetailWithDefaults instantiates a new CdDeployTimeDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdDeployTimeDetailWithDefaults() *CdDeployTimeDetail {
	this := CdDeployTimeDetail{}
	var application string = ""
	this.Application = &application
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *CdDeployTimeDetail) GetApplication() string {
	if o == nil || utils.IsNil(o.Application) {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdDeployTimeDetail) GetApplicationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *CdDeployTimeDetail) HasApplication() bool {
	if o != nil && !utils.IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *CdDeployTimeDetail) SetApplication(v string) {
	o.Application = &v
}

// GetCdDeployTime returns the CdDeployTime field value if set, zero value otherwise.
func (o *CdDeployTimeDetail) GetCdDeployTime() CdDeployTime {
	if o == nil || utils.IsNil(o.CdDeployTime) {
		var ret CdDeployTime
		return ret
	}
	return *o.CdDeployTime
}

// GetCdDeployTimeOk returns a tuple with the CdDeployTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdDeployTimeDetail) GetCdDeployTimeOk() (*CdDeployTime, bool) {
	if o == nil || utils.IsNil(o.CdDeployTime) {
		return nil, false
	}
	return o.CdDeployTime, true
}

// HasCdDeployTime returns a boolean if a field has been set.
func (o *CdDeployTimeDetail) HasCdDeployTime() bool {
	if o != nil && !utils.IsNil(o.CdDeployTime) {
		return true
	}

	return false
}

// SetCdDeployTime gets a reference to the given CdDeployTime and assigns it to the CdDeployTime field.
func (o *CdDeployTimeDetail) SetCdDeployTime(v CdDeployTime) {
	o.CdDeployTime = &v
}

func (o CdDeployTimeDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CdDeployTimeDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Application) {
		toSerialize["Application"] = o.Application
	}
	if !utils.IsNil(o.CdDeployTime) {
		toSerialize["CdDeployTime"] = o.CdDeployTime
	}
	return toSerialize, nil
}

type NullableCdDeployTimeDetail struct {
	value *CdDeployTimeDetail
	isSet bool
}

func (v NullableCdDeployTimeDetail) Get() *CdDeployTimeDetail {
	return v.value
}

func (v *NullableCdDeployTimeDetail) Set(val *CdDeployTimeDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableCdDeployTimeDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableCdDeployTimeDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdDeployTimeDetail(val *CdDeployTimeDetail) *NullableCdDeployTimeDetail {
	return &NullableCdDeployTimeDetail{value: val, isSet: true}
}

func (v NullableCdDeployTimeDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdDeployTimeDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


