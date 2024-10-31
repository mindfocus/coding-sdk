/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CdDeployTrendDetail type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CdDeployTrendDetail{}

// CdDeployTrendDetail CdDeployTrendDetail 结构
type CdDeployTrendDetail struct {
	// 应用名称
	Application *string `json:"Application,omitempty"`
	// 应用发布趋势统计
	CdDeployTrend []CdDeployTrend `json:"CdDeployTrend,omitempty"`
	// 成功发布次数
	SuccessCount *int64 `json:"SuccessCount,omitempty"`
	// 发布总次数
	TotalCount *int64 `json:"TotalCount,omitempty"`
}

// NewCdDeployTrendDetail instantiates a new CdDeployTrendDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdDeployTrendDetail() *CdDeployTrendDetail {
	this := CdDeployTrendDetail{}
	var application string = ""
	this.Application = &application
	return &this
}

// NewCdDeployTrendDetailWithDefaults instantiates a new CdDeployTrendDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdDeployTrendDetailWithDefaults() *CdDeployTrendDetail {
	this := CdDeployTrendDetail{}
	var application string = ""
	this.Application = &application
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *CdDeployTrendDetail) GetApplication() string {
	if o == nil || utils.IsNil(o.Application) {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdDeployTrendDetail) GetApplicationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *CdDeployTrendDetail) HasApplication() bool {
	if o != nil && !utils.IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *CdDeployTrendDetail) SetApplication(v string) {
	o.Application = &v
}

// GetCdDeployTrend returns the CdDeployTrend field value if set, zero value otherwise.
func (o *CdDeployTrendDetail) GetCdDeployTrend() []CdDeployTrend {
	if o == nil || utils.IsNil(o.CdDeployTrend) {
		var ret []CdDeployTrend
		return ret
	}
	return o.CdDeployTrend
}

// GetCdDeployTrendOk returns a tuple with the CdDeployTrend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdDeployTrendDetail) GetCdDeployTrendOk() ([]CdDeployTrend, bool) {
	if o == nil || utils.IsNil(o.CdDeployTrend) {
		return nil, false
	}
	return o.CdDeployTrend, true
}

// HasCdDeployTrend returns a boolean if a field has been set.
func (o *CdDeployTrendDetail) HasCdDeployTrend() bool {
	if o != nil && !utils.IsNil(o.CdDeployTrend) {
		return true
	}

	return false
}

// SetCdDeployTrend gets a reference to the given []CdDeployTrend and assigns it to the CdDeployTrend field.
func (o *CdDeployTrendDetail) SetCdDeployTrend(v []CdDeployTrend) {
	o.CdDeployTrend = v
}

// GetSuccessCount returns the SuccessCount field value if set, zero value otherwise.
func (o *CdDeployTrendDetail) GetSuccessCount() int64 {
	if o == nil || utils.IsNil(o.SuccessCount) {
		var ret int64
		return ret
	}
	return *o.SuccessCount
}

// GetSuccessCountOk returns a tuple with the SuccessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdDeployTrendDetail) GetSuccessCountOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.SuccessCount) {
		return nil, false
	}
	return o.SuccessCount, true
}

// HasSuccessCount returns a boolean if a field has been set.
func (o *CdDeployTrendDetail) HasSuccessCount() bool {
	if o != nil && !utils.IsNil(o.SuccessCount) {
		return true
	}

	return false
}

// SetSuccessCount gets a reference to the given int64 and assigns it to the SuccessCount field.
func (o *CdDeployTrendDetail) SetSuccessCount(v int64) {
	o.SuccessCount = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *CdDeployTrendDetail) GetTotalCount() int64 {
	if o == nil || utils.IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdDeployTrendDetail) GetTotalCountOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *CdDeployTrendDetail) HasTotalCount() bool {
	if o != nil && !utils.IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *CdDeployTrendDetail) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o CdDeployTrendDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CdDeployTrendDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Application) {
		toSerialize["Application"] = o.Application
	}
	if !utils.IsNil(o.CdDeployTrend) {
		toSerialize["CdDeployTrend"] = o.CdDeployTrend
	}
	if !utils.IsNil(o.SuccessCount) {
		toSerialize["SuccessCount"] = o.SuccessCount
	}
	if !utils.IsNil(o.TotalCount) {
		toSerialize["TotalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableCdDeployTrendDetail struct {
	value *CdDeployTrendDetail
	isSet bool
}

func (v NullableCdDeployTrendDetail) Get() *CdDeployTrendDetail {
	return v.value
}

func (v *NullableCdDeployTrendDetail) Set(val *CdDeployTrendDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableCdDeployTrendDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableCdDeployTrendDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdDeployTrendDetail(val *CdDeployTrendDetail) *NullableCdDeployTrendDetail {
	return &NullableCdDeployTrendDetail{value: val, isSet: true}
}

func (v NullableCdDeployTrendDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdDeployTrendDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

