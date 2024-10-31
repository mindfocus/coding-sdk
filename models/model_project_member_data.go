/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ProjectMemberData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ProjectMemberData{}

// ProjectMemberData 项目成员分页信息
type ProjectMemberData struct {
	// 第几页
	PageNumber *int32 `json:"PageNumber,omitempty"`
	// 每页条数
	PageSize *int32 `json:"PageSize,omitempty"`
	// 项目成员列表信息
	ProjectMembers []ProjectMemberUserData `json:"ProjectMembers,omitempty"`
	// 总条数
	TotalCount *int64 `json:"TotalCount,omitempty"`
}

// NewProjectMemberData instantiates a new ProjectMemberData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMemberData() *ProjectMemberData {
	this := ProjectMemberData{}
	return &this
}

// NewProjectMemberDataWithDefaults instantiates a new ProjectMemberData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMemberDataWithDefaults() *ProjectMemberData {
	this := ProjectMemberData{}
	return &this
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *ProjectMemberData) GetPageNumber() int32 {
	if o == nil || utils.IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberData) GetPageNumberOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *ProjectMemberData) HasPageNumber() bool {
	if o != nil && !utils.IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *ProjectMemberData) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ProjectMemberData) GetPageSize() int32 {
	if o == nil || utils.IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberData) GetPageSizeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ProjectMemberData) HasPageSize() bool {
	if o != nil && !utils.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ProjectMemberData) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetProjectMembers returns the ProjectMembers field value if set, zero value otherwise.
func (o *ProjectMemberData) GetProjectMembers() []ProjectMemberUserData {
	if o == nil || utils.IsNil(o.ProjectMembers) {
		var ret []ProjectMemberUserData
		return ret
	}
	return o.ProjectMembers
}

// GetProjectMembersOk returns a tuple with the ProjectMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberData) GetProjectMembersOk() ([]ProjectMemberUserData, bool) {
	if o == nil || utils.IsNil(o.ProjectMembers) {
		return nil, false
	}
	return o.ProjectMembers, true
}

// HasProjectMembers returns a boolean if a field has been set.
func (o *ProjectMemberData) HasProjectMembers() bool {
	if o != nil && !utils.IsNil(o.ProjectMembers) {
		return true
	}

	return false
}

// SetProjectMembers gets a reference to the given []ProjectMemberUserData and assigns it to the ProjectMembers field.
func (o *ProjectMemberData) SetProjectMembers(v []ProjectMemberUserData) {
	o.ProjectMembers = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ProjectMemberData) GetTotalCount() int64 {
	if o == nil || utils.IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberData) GetTotalCountOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ProjectMemberData) HasTotalCount() bool {
	if o != nil && !utils.IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *ProjectMemberData) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o ProjectMemberData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMemberData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.PageNumber) {
		toSerialize["PageNumber"] = o.PageNumber
	}
	if !utils.IsNil(o.PageSize) {
		toSerialize["PageSize"] = o.PageSize
	}
	if !utils.IsNil(o.ProjectMembers) {
		toSerialize["ProjectMembers"] = o.ProjectMembers
	}
	if !utils.IsNil(o.TotalCount) {
		toSerialize["TotalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableProjectMemberData struct {
	value *ProjectMemberData
	isSet bool
}

func (v NullableProjectMemberData) Get() *ProjectMemberData {
	return v.value
}

func (v *NullableProjectMemberData) Set(val *ProjectMemberData) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMemberData) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMemberData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMemberData(val *ProjectMemberData) *NullableProjectMemberData {
	return &NullableProjectMemberData{value: val, isSet: true}
}

func (v NullableProjectMemberData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMemberData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

