/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ReleaseListPage type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReleaseListPage{}

// ReleaseListPage git 版本分页列表信息
type ReleaseListPage struct {
	// 版本分页列表
	Releases []Release `json:"Releases,omitempty"`
	// 总共数量
	TotalCount *int64 `json:"TotalCount,omitempty"`
}

// NewReleaseListPage instantiates a new ReleaseListPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseListPage() *ReleaseListPage {
	this := ReleaseListPage{}
	return &this
}

// NewReleaseListPageWithDefaults instantiates a new ReleaseListPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseListPageWithDefaults() *ReleaseListPage {
	this := ReleaseListPage{}
	return &this
}

// GetReleases returns the Releases field value if set, zero value otherwise.
func (o *ReleaseListPage) GetReleases() []Release {
	if o == nil || utils.IsNil(o.Releases) {
		var ret []Release
		return ret
	}
	return o.Releases
}

// GetReleasesOk returns a tuple with the Releases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseListPage) GetReleasesOk() ([]Release, bool) {
	if o == nil || utils.IsNil(o.Releases) {
		return nil, false
	}
	return o.Releases, true
}

// HasReleases returns a boolean if a field has been set.
func (o *ReleaseListPage) HasReleases() bool {
	if o != nil && !utils.IsNil(o.Releases) {
		return true
	}

	return false
}

// SetReleases gets a reference to the given []Release and assigns it to the Releases field.
func (o *ReleaseListPage) SetReleases(v []Release) {
	o.Releases = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ReleaseListPage) GetTotalCount() int64 {
	if o == nil || utils.IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseListPage) GetTotalCountOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ReleaseListPage) HasTotalCount() bool {
	if o != nil && !utils.IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *ReleaseListPage) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o ReleaseListPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleaseListPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Releases) {
		toSerialize["Releases"] = o.Releases
	}
	if !utils.IsNil(o.TotalCount) {
		toSerialize["TotalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableReleaseListPage struct {
	value *ReleaseListPage
	isSet bool
}

func (v NullableReleaseListPage) Get() *ReleaseListPage {
	return v.value
}

func (v *NullableReleaseListPage) Set(val *ReleaseListPage) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseListPage) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseListPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseListPage(val *ReleaseListPage) *NullableReleaseListPage {
	return &NullableReleaseListPage{value: val, isSet: true}
}

func (v NullableReleaseListPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseListPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

