/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the IssueDetailListDataWithPage type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &IssueDetailListDataWithPage{}

// IssueDetailListDataWithPage 分页查询版本发布范围
type IssueDetailListDataWithPage struct {
	// 事项列表 
	List []IssueDetail `json:"List,omitempty"`
	// 页码
	PageNumber *int64 `json:"PageNumber,omitempty"`
	// 页数
	PageSize *int64 `json:"PageSize,omitempty"`
	// 总共几条事项 
	TotalCount *int64 `json:"TotalCount,omitempty"`
	// 总共几页 
	TotalPage *int64 `json:"TotalPage,omitempty"`
}

// NewIssueDetailListDataWithPage instantiates a new IssueDetailListDataWithPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueDetailListDataWithPage() *IssueDetailListDataWithPage {
	this := IssueDetailListDataWithPage{}
	return &this
}

// NewIssueDetailListDataWithPageWithDefaults instantiates a new IssueDetailListDataWithPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueDetailListDataWithPageWithDefaults() *IssueDetailListDataWithPage {
	this := IssueDetailListDataWithPage{}
	return &this
}

// GetList returns the List field value if set, zero value otherwise.
func (o *IssueDetailListDataWithPage) GetList() []IssueDetail {
	if o == nil || utils.IsNil(o.List) {
		var ret []IssueDetail
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetailListDataWithPage) GetListOk() ([]IssueDetail, bool) {
	if o == nil || utils.IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *IssueDetailListDataWithPage) HasList() bool {
	if o != nil && !utils.IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []IssueDetail and assigns it to the List field.
func (o *IssueDetailListDataWithPage) SetList(v []IssueDetail) {
	o.List = v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *IssueDetailListDataWithPage) GetPageNumber() int64 {
	if o == nil || utils.IsNil(o.PageNumber) {
		var ret int64
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetailListDataWithPage) GetPageNumberOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *IssueDetailListDataWithPage) HasPageNumber() bool {
	if o != nil && !utils.IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int64 and assigns it to the PageNumber field.
func (o *IssueDetailListDataWithPage) SetPageNumber(v int64) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *IssueDetailListDataWithPage) GetPageSize() int64 {
	if o == nil || utils.IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetailListDataWithPage) GetPageSizeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *IssueDetailListDataWithPage) HasPageSize() bool {
	if o != nil && !utils.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *IssueDetailListDataWithPage) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *IssueDetailListDataWithPage) GetTotalCount() int64 {
	if o == nil || utils.IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetailListDataWithPage) GetTotalCountOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *IssueDetailListDataWithPage) HasTotalCount() bool {
	if o != nil && !utils.IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *IssueDetailListDataWithPage) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetTotalPage returns the TotalPage field value if set, zero value otherwise.
func (o *IssueDetailListDataWithPage) GetTotalPage() int64 {
	if o == nil || utils.IsNil(o.TotalPage) {
		var ret int64
		return ret
	}
	return *o.TotalPage
}

// GetTotalPageOk returns a tuple with the TotalPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetailListDataWithPage) GetTotalPageOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.TotalPage) {
		return nil, false
	}
	return o.TotalPage, true
}

// HasTotalPage returns a boolean if a field has been set.
func (o *IssueDetailListDataWithPage) HasTotalPage() bool {
	if o != nil && !utils.IsNil(o.TotalPage) {
		return true
	}

	return false
}

// SetTotalPage gets a reference to the given int64 and assigns it to the TotalPage field.
func (o *IssueDetailListDataWithPage) SetTotalPage(v int64) {
	o.TotalPage = &v
}

func (o IssueDetailListDataWithPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueDetailListDataWithPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.List) {
		toSerialize["List"] = o.List
	}
	if !utils.IsNil(o.PageNumber) {
		toSerialize["PageNumber"] = o.PageNumber
	}
	if !utils.IsNil(o.PageSize) {
		toSerialize["PageSize"] = o.PageSize
	}
	if !utils.IsNil(o.TotalCount) {
		toSerialize["TotalCount"] = o.TotalCount
	}
	if !utils.IsNil(o.TotalPage) {
		toSerialize["TotalPage"] = o.TotalPage
	}
	return toSerialize, nil
}

type NullableIssueDetailListDataWithPage struct {
	value *IssueDetailListDataWithPage
	isSet bool
}

func (v NullableIssueDetailListDataWithPage) Get() *IssueDetailListDataWithPage {
	return v.value
}

func (v *NullableIssueDetailListDataWithPage) Set(val *IssueDetailListDataWithPage) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueDetailListDataWithPage) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueDetailListDataWithPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueDetailListDataWithPage(val *IssueDetailListDataWithPage) *NullableIssueDetailListDataWithPage {
	return &NullableIssueDetailListDataWithPage{value: val, isSet: true}
}

func (v NullableIssueDetailListDataWithPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueDetailListDataWithPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


