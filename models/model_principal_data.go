/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the PrincipalData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PrincipalData{}

// PrincipalData 查询项目成员主体 响应信息
type PrincipalData struct {
	// 页数
	PageNumber *int32 `json:"PageNumber,omitempty"`
	// 条数
	PageSize *int32 `json:"PageSize,omitempty"`
	// 成员主体信息
	Principals []PrincipalResp `json:"Principals,omitempty"`
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty"`
}

// NewPrincipalData instantiates a new PrincipalData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalData() *PrincipalData {
	this := PrincipalData{}
	return &this
}

// NewPrincipalDataWithDefaults instantiates a new PrincipalData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalDataWithDefaults() *PrincipalData {
	this := PrincipalData{}
	return &this
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *PrincipalData) GetPageNumber() int32 {
	if o == nil || utils.IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalData) GetPageNumberOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *PrincipalData) HasPageNumber() bool {
	if o != nil && !utils.IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *PrincipalData) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *PrincipalData) GetPageSize() int32 {
	if o == nil || utils.IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalData) GetPageSizeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *PrincipalData) HasPageSize() bool {
	if o != nil && !utils.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *PrincipalData) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPrincipals returns the Principals field value if set, zero value otherwise.
func (o *PrincipalData) GetPrincipals() []PrincipalResp {
	if o == nil || utils.IsNil(o.Principals) {
		var ret []PrincipalResp
		return ret
	}
	return o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalData) GetPrincipalsOk() ([]PrincipalResp, bool) {
	if o == nil || utils.IsNil(o.Principals) {
		return nil, false
	}
	return o.Principals, true
}

// HasPrincipals returns a boolean if a field has been set.
func (o *PrincipalData) HasPrincipals() bool {
	if o != nil && !utils.IsNil(o.Principals) {
		return true
	}

	return false
}

// SetPrincipals gets a reference to the given []PrincipalResp and assigns it to the Principals field.
func (o *PrincipalData) SetPrincipals(v []PrincipalResp) {
	o.Principals = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *PrincipalData) GetTotalCount() int64 {
	if o == nil || utils.IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalData) GetTotalCountOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *PrincipalData) HasTotalCount() bool {
	if o != nil && !utils.IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *PrincipalData) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o PrincipalData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.PageNumber) {
		toSerialize["PageNumber"] = o.PageNumber
	}
	if !utils.IsNil(o.PageSize) {
		toSerialize["PageSize"] = o.PageSize
	}
	if !utils.IsNil(o.Principals) {
		toSerialize["Principals"] = o.Principals
	}
	if !utils.IsNil(o.TotalCount) {
		toSerialize["TotalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullablePrincipalData struct {
	value *PrincipalData
	isSet bool
}

func (v NullablePrincipalData) Get() *PrincipalData {
	return v.value
}

func (v *NullablePrincipalData) Set(val *PrincipalData) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalData) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalData(val *PrincipalData) *NullablePrincipalData {
	return &NullablePrincipalData{value: val, isSet: true}
}

func (v NullablePrincipalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


