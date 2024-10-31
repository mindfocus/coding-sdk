/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ReportLittleData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReportLittleData{}

// ReportLittleData 测试报告详情
type ReportLittleData struct {
	Report *ReportLittle `json:"Report,omitempty"`
}

// NewReportLittleData instantiates a new ReportLittleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportLittleData() *ReportLittleData {
	this := ReportLittleData{}
	return &this
}

// NewReportLittleDataWithDefaults instantiates a new ReportLittleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportLittleDataWithDefaults() *ReportLittleData {
	this := ReportLittleData{}
	return &this
}

// GetReport returns the Report field value if set, zero value otherwise.
func (o *ReportLittleData) GetReport() ReportLittle {
	if o == nil || utils.IsNil(o.Report) {
		var ret ReportLittle
		return ret
	}
	return *o.Report
}

// GetReportOk returns a tuple with the Report field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportLittleData) GetReportOk() (*ReportLittle, bool) {
	if o == nil || utils.IsNil(o.Report) {
		return nil, false
	}
	return o.Report, true
}

// HasReport returns a boolean if a field has been set.
func (o *ReportLittleData) HasReport() bool {
	if o != nil && !utils.IsNil(o.Report) {
		return true
	}

	return false
}

// SetReport gets a reference to the given ReportLittle and assigns it to the Report field.
func (o *ReportLittleData) SetReport(v ReportLittle) {
	o.Report = &v
}

func (o ReportLittleData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportLittleData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Report) {
		toSerialize["Report"] = o.Report
	}
	return toSerialize, nil
}

type NullableReportLittleData struct {
	value *ReportLittleData
	isSet bool
}

func (v NullableReportLittleData) Get() *ReportLittleData {
	return v.value
}

func (v *NullableReportLittleData) Set(val *ReportLittleData) {
	v.value = val
	v.isSet = true
}

func (v NullableReportLittleData) IsSet() bool {
	return v.isSet
}

func (v *NullableReportLittleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportLittleData(val *ReportLittleData) *NullableReportLittleData {
	return &NullableReportLittleData{value: val, isSet: true}
}

func (v NullableReportLittleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportLittleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

