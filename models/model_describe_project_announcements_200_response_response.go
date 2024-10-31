/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeProjectAnnouncements200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeProjectAnnouncements200ResponseResponse{}

// DescribeProjectAnnouncements200ResponseResponse 公共返回体
type DescribeProjectAnnouncements200ResponseResponse struct {
	// 公告总数
	TotalCount *int64 `json:"TotalCount,omitempty"`
	// 每页数量
	PageSize *int64 `json:"PageSize,omitempty"`
	// 页数
	PageNumber *int64 `json:"PageNumber,omitempty"`
	// 公告列表
	List []ProjectAnnouncementProjectAnnouncement `json:"List,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeProjectAnnouncements200ResponseResponse instantiates a new DescribeProjectAnnouncements200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeProjectAnnouncements200ResponseResponse() *DescribeProjectAnnouncements200ResponseResponse {
	this := DescribeProjectAnnouncements200ResponseResponse{}
	var totalCount int64 = 0
	this.TotalCount = &totalCount
	var pageSize int64 = 0
	this.PageSize = &pageSize
	var pageNumber int64 = 0
	this.PageNumber = &pageNumber
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeProjectAnnouncements200ResponseResponseWithDefaults instantiates a new DescribeProjectAnnouncements200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeProjectAnnouncements200ResponseResponseWithDefaults() *DescribeProjectAnnouncements200ResponseResponse {
	this := DescribeProjectAnnouncements200ResponseResponse{}
	var totalCount int64 = 0
	this.TotalCount = &totalCount
	var pageSize int64 = 0
	this.PageSize = &pageSize
	var pageNumber int64 = 0
	this.PageNumber = &pageNumber
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *DescribeProjectAnnouncements200ResponseResponse) GetTotalCount() int64 {
	if o == nil || utils.IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectAnnouncements200ResponseResponse) GetTotalCountOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *DescribeProjectAnnouncements200ResponseResponse) HasTotalCount() bool {
	if o != nil && !utils.IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *DescribeProjectAnnouncements200ResponseResponse) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *DescribeProjectAnnouncements200ResponseResponse) GetPageSize() int64 {
	if o == nil || utils.IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectAnnouncements200ResponseResponse) GetPageSizeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *DescribeProjectAnnouncements200ResponseResponse) HasPageSize() bool {
	if o != nil && !utils.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *DescribeProjectAnnouncements200ResponseResponse) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *DescribeProjectAnnouncements200ResponseResponse) GetPageNumber() int64 {
	if o == nil || utils.IsNil(o.PageNumber) {
		var ret int64
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectAnnouncements200ResponseResponse) GetPageNumberOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *DescribeProjectAnnouncements200ResponseResponse) HasPageNumber() bool {
	if o != nil && !utils.IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int64 and assigns it to the PageNumber field.
func (o *DescribeProjectAnnouncements200ResponseResponse) SetPageNumber(v int64) {
	o.PageNumber = &v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *DescribeProjectAnnouncements200ResponseResponse) GetList() []ProjectAnnouncementProjectAnnouncement {
	if o == nil || utils.IsNil(o.List) {
		var ret []ProjectAnnouncementProjectAnnouncement
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectAnnouncements200ResponseResponse) GetListOk() ([]ProjectAnnouncementProjectAnnouncement, bool) {
	if o == nil || utils.IsNil(o.List) {
		return nil, false
	}
	return o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *DescribeProjectAnnouncements200ResponseResponse) HasList() bool {
	if o != nil && !utils.IsNil(o.List) {
		return true
	}

	return false
}

// SetList gets a reference to the given []ProjectAnnouncementProjectAnnouncement and assigns it to the List field.
func (o *DescribeProjectAnnouncements200ResponseResponse) SetList(v []ProjectAnnouncementProjectAnnouncement) {
	o.List = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeProjectAnnouncements200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectAnnouncements200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeProjectAnnouncements200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeProjectAnnouncements200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeProjectAnnouncements200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeProjectAnnouncements200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.TotalCount) {
		toSerialize["TotalCount"] = o.TotalCount
	}
	if !utils.IsNil(o.PageSize) {
		toSerialize["PageSize"] = o.PageSize
	}
	if !utils.IsNil(o.PageNumber) {
		toSerialize["PageNumber"] = o.PageNumber
	}
	if !utils.IsNil(o.List) {
		toSerialize["List"] = o.List
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeProjectAnnouncements200ResponseResponse struct {
	value *DescribeProjectAnnouncements200ResponseResponse
	isSet bool
}

func (v NullableDescribeProjectAnnouncements200ResponseResponse) Get() *DescribeProjectAnnouncements200ResponseResponse {
	return v.value
}

func (v *NullableDescribeProjectAnnouncements200ResponseResponse) Set(val *DescribeProjectAnnouncements200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeProjectAnnouncements200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeProjectAnnouncements200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeProjectAnnouncements200ResponseResponse(val *DescribeProjectAnnouncements200ResponseResponse) *NullableDescribeProjectAnnouncements200ResponseResponse {
	return &NullableDescribeProjectAnnouncements200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeProjectAnnouncements200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeProjectAnnouncements200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

