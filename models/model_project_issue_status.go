/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ProjectIssueStatus type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ProjectIssueStatus{}

// ProjectIssueStatus 项目的事项状态
type ProjectIssueStatus struct {
	// 创建时间戳
	CreatedAt *int64 `json:"CreatedAt,omitempty"`
	// 是否是默认值
	IsDefault *bool `json:"IsDefault,omitempty"`
	IssueStatus *IssueStatus `json:"IssueStatus,omitempty"`
	// 事项状态 ID
	IssueStatusId *int64 `json:"IssueStatusId,omitempty"`
	// 事项类型
	IssueType *string `json:"IssueType,omitempty"`
	// 排序
	Sort *int64 `json:"Sort,omitempty"`
	// 修改时间戳
	UpdatedAt *int64 `json:"UpdatedAt,omitempty"`
}

// NewProjectIssueStatus instantiates a new ProjectIssueStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectIssueStatus() *ProjectIssueStatus {
	this := ProjectIssueStatus{}
	var isDefault bool = false
	this.IsDefault = &isDefault
	var issueType string = ""
	this.IssueType = &issueType
	return &this
}

// NewProjectIssueStatusWithDefaults instantiates a new ProjectIssueStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectIssueStatusWithDefaults() *ProjectIssueStatus {
	this := ProjectIssueStatus{}
	var isDefault bool = false
	this.IsDefault = &isDefault
	var issueType string = ""
	this.IssueType = &issueType
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ProjectIssueStatus) GetCreatedAt() int64 {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectIssueStatus) GetCreatedAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ProjectIssueStatus) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *ProjectIssueStatus) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *ProjectIssueStatus) GetIsDefault() bool {
	if o == nil || utils.IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectIssueStatus) GetIsDefaultOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *ProjectIssueStatus) HasIsDefault() bool {
	if o != nil && !utils.IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *ProjectIssueStatus) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetIssueStatus returns the IssueStatus field value if set, zero value otherwise.
func (o *ProjectIssueStatus) GetIssueStatus() IssueStatus {
	if o == nil || utils.IsNil(o.IssueStatus) {
		var ret IssueStatus
		return ret
	}
	return *o.IssueStatus
}

// GetIssueStatusOk returns a tuple with the IssueStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectIssueStatus) GetIssueStatusOk() (*IssueStatus, bool) {
	if o == nil || utils.IsNil(o.IssueStatus) {
		return nil, false
	}
	return o.IssueStatus, true
}

// HasIssueStatus returns a boolean if a field has been set.
func (o *ProjectIssueStatus) HasIssueStatus() bool {
	if o != nil && !utils.IsNil(o.IssueStatus) {
		return true
	}

	return false
}

// SetIssueStatus gets a reference to the given IssueStatus and assigns it to the IssueStatus field.
func (o *ProjectIssueStatus) SetIssueStatus(v IssueStatus) {
	o.IssueStatus = &v
}

// GetIssueStatusId returns the IssueStatusId field value if set, zero value otherwise.
func (o *ProjectIssueStatus) GetIssueStatusId() int64 {
	if o == nil || utils.IsNil(o.IssueStatusId) {
		var ret int64
		return ret
	}
	return *o.IssueStatusId
}

// GetIssueStatusIdOk returns a tuple with the IssueStatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectIssueStatus) GetIssueStatusIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.IssueStatusId) {
		return nil, false
	}
	return o.IssueStatusId, true
}

// HasIssueStatusId returns a boolean if a field has been set.
func (o *ProjectIssueStatus) HasIssueStatusId() bool {
	if o != nil && !utils.IsNil(o.IssueStatusId) {
		return true
	}

	return false
}

// SetIssueStatusId gets a reference to the given int64 and assigns it to the IssueStatusId field.
func (o *ProjectIssueStatus) SetIssueStatusId(v int64) {
	o.IssueStatusId = &v
}

// GetIssueType returns the IssueType field value if set, zero value otherwise.
func (o *ProjectIssueStatus) GetIssueType() string {
	if o == nil || utils.IsNil(o.IssueType) {
		var ret string
		return ret
	}
	return *o.IssueType
}

// GetIssueTypeOk returns a tuple with the IssueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectIssueStatus) GetIssueTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IssueType) {
		return nil, false
	}
	return o.IssueType, true
}

// HasIssueType returns a boolean if a field has been set.
func (o *ProjectIssueStatus) HasIssueType() bool {
	if o != nil && !utils.IsNil(o.IssueType) {
		return true
	}

	return false
}

// SetIssueType gets a reference to the given string and assigns it to the IssueType field.
func (o *ProjectIssueStatus) SetIssueType(v string) {
	o.IssueType = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ProjectIssueStatus) GetSort() int64 {
	if o == nil || utils.IsNil(o.Sort) {
		var ret int64
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectIssueStatus) GetSortOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ProjectIssueStatus) HasSort() bool {
	if o != nil && !utils.IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given int64 and assigns it to the Sort field.
func (o *ProjectIssueStatus) SetSort(v int64) {
	o.Sort = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ProjectIssueStatus) GetUpdatedAt() int64 {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectIssueStatus) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ProjectIssueStatus) HasUpdatedAt() bool {
	if o != nil && !utils.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *ProjectIssueStatus) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

func (o ProjectIssueStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectIssueStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if !utils.IsNil(o.IsDefault) {
		toSerialize["IsDefault"] = o.IsDefault
	}
	if !utils.IsNil(o.IssueStatus) {
		toSerialize["IssueStatus"] = o.IssueStatus
	}
	if !utils.IsNil(o.IssueStatusId) {
		toSerialize["IssueStatusId"] = o.IssueStatusId
	}
	if !utils.IsNil(o.IssueType) {
		toSerialize["IssueType"] = o.IssueType
	}
	if !utils.IsNil(o.Sort) {
		toSerialize["Sort"] = o.Sort
	}
	if !utils.IsNil(o.UpdatedAt) {
		toSerialize["UpdatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableProjectIssueStatus struct {
	value *ProjectIssueStatus
	isSet bool
}

func (v NullableProjectIssueStatus) Get() *ProjectIssueStatus {
	return v.value
}

func (v *NullableProjectIssueStatus) Set(val *ProjectIssueStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectIssueStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectIssueStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectIssueStatus(val *ProjectIssueStatus) *NullableProjectIssueStatus {
	return &NullableProjectIssueStatus{value: val, isSet: true}
}

func (v NullableProjectIssueStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectIssueStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

