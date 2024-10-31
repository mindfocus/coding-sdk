/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeSelfMergeRequestsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeSelfMergeRequestsRequest{}

// DescribeSelfMergeRequestsRequest struct for DescribeSelfMergeRequestsRequest
type DescribeSelfMergeRequestsRequest struct {
	// mr创建结束时间
	CreatedAtEndDate *string `json:"CreatedAtEndDate,omitempty"`
	// mr创建开始时间
	CreatedAtStartDate *string `json:"CreatedAtStartDate,omitempty"`
	// 是否正序排序 默认 false
	IsSortDirectionAsc *bool `json:"IsSortDirectionAsc,omitempty"`
	// 关键词
	KeyWord *string `json:"KeyWord,omitempty"`
	// 标签
	Label *string `json:"Label,omitempty"`
	// 合并请求合并者邮箱
	MergerEmail *string `json:"MergerEmail,omitempty"`
	// 合并请求合并者全局 key
	MergerGlobalKey *string `json:"MergerGlobalKey,omitempty"`
	// 页数 默认 1
	PageNumber *int64 `json:"PageNumber,omitempty"`
	// 每页条数 默认 10
	PageSize *int64 `json:"PageSize,omitempty"`
	// 合并请求评审者邮箱
	ReviewerEmail *string `json:"ReviewerEmail,omitempty"`
	// 合并请求评审者全局 key
	ReviewerGlobalKey *string `json:"ReviewerGlobalKey,omitempty"`
	// 排序 action_at：以更新时间排序  created_at：以创建时间排序
	Sort *string `json:"Sort,omitempty"`
	// 源分支
	SourceBranch *string `json:"SourceBranch,omitempty"`
	// 合并请求状态 open/close/all
	Status *string `json:"Status,omitempty"`
	// 欲合入分支
	TargetBranch *string `json:"TargetBranch,omitempty"`
	// mr更新结束时间
	UpdatedAtEndDate *string `json:"UpdatedAtEndDate,omitempty"`
	// mr更新开始时间
	UpdatedAtStartDate *string `json:"UpdatedAtStartDate,omitempty"`
}

// NewDescribeSelfMergeRequestsRequest instantiates a new DescribeSelfMergeRequestsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeSelfMergeRequestsRequest() *DescribeSelfMergeRequestsRequest {
	this := DescribeSelfMergeRequestsRequest{}
	return &this
}

// NewDescribeSelfMergeRequestsRequestWithDefaults instantiates a new DescribeSelfMergeRequestsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeSelfMergeRequestsRequestWithDefaults() *DescribeSelfMergeRequestsRequest {
	this := DescribeSelfMergeRequestsRequest{}
	return &this
}

// GetCreatedAtEndDate returns the CreatedAtEndDate field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetCreatedAtEndDate() string {
	if o == nil || utils.IsNil(o.CreatedAtEndDate) {
		var ret string
		return ret
	}
	return *o.CreatedAtEndDate
}

// GetCreatedAtEndDateOk returns a tuple with the CreatedAtEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetCreatedAtEndDateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedAtEndDate) {
		return nil, false
	}
	return o.CreatedAtEndDate, true
}

// HasCreatedAtEndDate returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasCreatedAtEndDate() bool {
	if o != nil && !utils.IsNil(o.CreatedAtEndDate) {
		return true
	}

	return false
}

// SetCreatedAtEndDate gets a reference to the given string and assigns it to the CreatedAtEndDate field.
func (o *DescribeSelfMergeRequestsRequest) SetCreatedAtEndDate(v string) {
	o.CreatedAtEndDate = &v
}

// GetCreatedAtStartDate returns the CreatedAtStartDate field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetCreatedAtStartDate() string {
	if o == nil || utils.IsNil(o.CreatedAtStartDate) {
		var ret string
		return ret
	}
	return *o.CreatedAtStartDate
}

// GetCreatedAtStartDateOk returns a tuple with the CreatedAtStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetCreatedAtStartDateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedAtStartDate) {
		return nil, false
	}
	return o.CreatedAtStartDate, true
}

// HasCreatedAtStartDate returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasCreatedAtStartDate() bool {
	if o != nil && !utils.IsNil(o.CreatedAtStartDate) {
		return true
	}

	return false
}

// SetCreatedAtStartDate gets a reference to the given string and assigns it to the CreatedAtStartDate field.
func (o *DescribeSelfMergeRequestsRequest) SetCreatedAtStartDate(v string) {
	o.CreatedAtStartDate = &v
}

// GetIsSortDirectionAsc returns the IsSortDirectionAsc field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetIsSortDirectionAsc() bool {
	if o == nil || utils.IsNil(o.IsSortDirectionAsc) {
		var ret bool
		return ret
	}
	return *o.IsSortDirectionAsc
}

// GetIsSortDirectionAscOk returns a tuple with the IsSortDirectionAsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetIsSortDirectionAscOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsSortDirectionAsc) {
		return nil, false
	}
	return o.IsSortDirectionAsc, true
}

// HasIsSortDirectionAsc returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasIsSortDirectionAsc() bool {
	if o != nil && !utils.IsNil(o.IsSortDirectionAsc) {
		return true
	}

	return false
}

// SetIsSortDirectionAsc gets a reference to the given bool and assigns it to the IsSortDirectionAsc field.
func (o *DescribeSelfMergeRequestsRequest) SetIsSortDirectionAsc(v bool) {
	o.IsSortDirectionAsc = &v
}

// GetKeyWord returns the KeyWord field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetKeyWord() string {
	if o == nil || utils.IsNil(o.KeyWord) {
		var ret string
		return ret
	}
	return *o.KeyWord
}

// GetKeyWordOk returns a tuple with the KeyWord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetKeyWordOk() (*string, bool) {
	if o == nil || utils.IsNil(o.KeyWord) {
		return nil, false
	}
	return o.KeyWord, true
}

// HasKeyWord returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasKeyWord() bool {
	if o != nil && !utils.IsNil(o.KeyWord) {
		return true
	}

	return false
}

// SetKeyWord gets a reference to the given string and assigns it to the KeyWord field.
func (o *DescribeSelfMergeRequestsRequest) SetKeyWord(v string) {
	o.KeyWord = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetLabel() string {
	if o == nil || utils.IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetLabelOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasLabel() bool {
	if o != nil && !utils.IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *DescribeSelfMergeRequestsRequest) SetLabel(v string) {
	o.Label = &v
}

// GetMergerEmail returns the MergerEmail field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetMergerEmail() string {
	if o == nil || utils.IsNil(o.MergerEmail) {
		var ret string
		return ret
	}
	return *o.MergerEmail
}

// GetMergerEmailOk returns a tuple with the MergerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetMergerEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MergerEmail) {
		return nil, false
	}
	return o.MergerEmail, true
}

// HasMergerEmail returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasMergerEmail() bool {
	if o != nil && !utils.IsNil(o.MergerEmail) {
		return true
	}

	return false
}

// SetMergerEmail gets a reference to the given string and assigns it to the MergerEmail field.
func (o *DescribeSelfMergeRequestsRequest) SetMergerEmail(v string) {
	o.MergerEmail = &v
}

// GetMergerGlobalKey returns the MergerGlobalKey field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetMergerGlobalKey() string {
	if o == nil || utils.IsNil(o.MergerGlobalKey) {
		var ret string
		return ret
	}
	return *o.MergerGlobalKey
}

// GetMergerGlobalKeyOk returns a tuple with the MergerGlobalKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetMergerGlobalKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.MergerGlobalKey) {
		return nil, false
	}
	return o.MergerGlobalKey, true
}

// HasMergerGlobalKey returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasMergerGlobalKey() bool {
	if o != nil && !utils.IsNil(o.MergerGlobalKey) {
		return true
	}

	return false
}

// SetMergerGlobalKey gets a reference to the given string and assigns it to the MergerGlobalKey field.
func (o *DescribeSelfMergeRequestsRequest) SetMergerGlobalKey(v string) {
	o.MergerGlobalKey = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetPageNumber() int64 {
	if o == nil || utils.IsNil(o.PageNumber) {
		var ret int64
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetPageNumberOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasPageNumber() bool {
	if o != nil && !utils.IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int64 and assigns it to the PageNumber field.
func (o *DescribeSelfMergeRequestsRequest) SetPageNumber(v int64) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetPageSize() int64 {
	if o == nil || utils.IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasPageSize() bool {
	if o != nil && !utils.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *DescribeSelfMergeRequestsRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetReviewerEmail returns the ReviewerEmail field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetReviewerEmail() string {
	if o == nil || utils.IsNil(o.ReviewerEmail) {
		var ret string
		return ret
	}
	return *o.ReviewerEmail
}

// GetReviewerEmailOk returns a tuple with the ReviewerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetReviewerEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReviewerEmail) {
		return nil, false
	}
	return o.ReviewerEmail, true
}

// HasReviewerEmail returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasReviewerEmail() bool {
	if o != nil && !utils.IsNil(o.ReviewerEmail) {
		return true
	}

	return false
}

// SetReviewerEmail gets a reference to the given string and assigns it to the ReviewerEmail field.
func (o *DescribeSelfMergeRequestsRequest) SetReviewerEmail(v string) {
	o.ReviewerEmail = &v
}

// GetReviewerGlobalKey returns the ReviewerGlobalKey field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetReviewerGlobalKey() string {
	if o == nil || utils.IsNil(o.ReviewerGlobalKey) {
		var ret string
		return ret
	}
	return *o.ReviewerGlobalKey
}

// GetReviewerGlobalKeyOk returns a tuple with the ReviewerGlobalKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetReviewerGlobalKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ReviewerGlobalKey) {
		return nil, false
	}
	return o.ReviewerGlobalKey, true
}

// HasReviewerGlobalKey returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasReviewerGlobalKey() bool {
	if o != nil && !utils.IsNil(o.ReviewerGlobalKey) {
		return true
	}

	return false
}

// SetReviewerGlobalKey gets a reference to the given string and assigns it to the ReviewerGlobalKey field.
func (o *DescribeSelfMergeRequestsRequest) SetReviewerGlobalKey(v string) {
	o.ReviewerGlobalKey = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetSort() string {
	if o == nil || utils.IsNil(o.Sort) {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetSortOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasSort() bool {
	if o != nil && !utils.IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *DescribeSelfMergeRequestsRequest) SetSort(v string) {
	o.Sort = &v
}

// GetSourceBranch returns the SourceBranch field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetSourceBranch() string {
	if o == nil || utils.IsNil(o.SourceBranch) {
		var ret string
		return ret
	}
	return *o.SourceBranch
}

// GetSourceBranchOk returns a tuple with the SourceBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetSourceBranchOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SourceBranch) {
		return nil, false
	}
	return o.SourceBranch, true
}

// HasSourceBranch returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasSourceBranch() bool {
	if o != nil && !utils.IsNil(o.SourceBranch) {
		return true
	}

	return false
}

// SetSourceBranch gets a reference to the given string and assigns it to the SourceBranch field.
func (o *DescribeSelfMergeRequestsRequest) SetSourceBranch(v string) {
	o.SourceBranch = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DescribeSelfMergeRequestsRequest) SetStatus(v string) {
	o.Status = &v
}

// GetTargetBranch returns the TargetBranch field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetTargetBranch() string {
	if o == nil || utils.IsNil(o.TargetBranch) {
		var ret string
		return ret
	}
	return *o.TargetBranch
}

// GetTargetBranchOk returns a tuple with the TargetBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetTargetBranchOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TargetBranch) {
		return nil, false
	}
	return o.TargetBranch, true
}

// HasTargetBranch returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasTargetBranch() bool {
	if o != nil && !utils.IsNil(o.TargetBranch) {
		return true
	}

	return false
}

// SetTargetBranch gets a reference to the given string and assigns it to the TargetBranch field.
func (o *DescribeSelfMergeRequestsRequest) SetTargetBranch(v string) {
	o.TargetBranch = &v
}

// GetUpdatedAtEndDate returns the UpdatedAtEndDate field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetUpdatedAtEndDate() string {
	if o == nil || utils.IsNil(o.UpdatedAtEndDate) {
		var ret string
		return ret
	}
	return *o.UpdatedAtEndDate
}

// GetUpdatedAtEndDateOk returns a tuple with the UpdatedAtEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetUpdatedAtEndDateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.UpdatedAtEndDate) {
		return nil, false
	}
	return o.UpdatedAtEndDate, true
}

// HasUpdatedAtEndDate returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasUpdatedAtEndDate() bool {
	if o != nil && !utils.IsNil(o.UpdatedAtEndDate) {
		return true
	}

	return false
}

// SetUpdatedAtEndDate gets a reference to the given string and assigns it to the UpdatedAtEndDate field.
func (o *DescribeSelfMergeRequestsRequest) SetUpdatedAtEndDate(v string) {
	o.UpdatedAtEndDate = &v
}

// GetUpdatedAtStartDate returns the UpdatedAtStartDate field value if set, zero value otherwise.
func (o *DescribeSelfMergeRequestsRequest) GetUpdatedAtStartDate() string {
	if o == nil || utils.IsNil(o.UpdatedAtStartDate) {
		var ret string
		return ret
	}
	return *o.UpdatedAtStartDate
}

// GetUpdatedAtStartDateOk returns a tuple with the UpdatedAtStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSelfMergeRequestsRequest) GetUpdatedAtStartDateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.UpdatedAtStartDate) {
		return nil, false
	}
	return o.UpdatedAtStartDate, true
}

// HasUpdatedAtStartDate returns a boolean if a field has been set.
func (o *DescribeSelfMergeRequestsRequest) HasUpdatedAtStartDate() bool {
	if o != nil && !utils.IsNil(o.UpdatedAtStartDate) {
		return true
	}

	return false
}

// SetUpdatedAtStartDate gets a reference to the given string and assigns it to the UpdatedAtStartDate field.
func (o *DescribeSelfMergeRequestsRequest) SetUpdatedAtStartDate(v string) {
	o.UpdatedAtStartDate = &v
}

func (o DescribeSelfMergeRequestsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeSelfMergeRequestsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CreatedAtEndDate) {
		toSerialize["CreatedAtEndDate"] = o.CreatedAtEndDate
	}
	if !utils.IsNil(o.CreatedAtStartDate) {
		toSerialize["CreatedAtStartDate"] = o.CreatedAtStartDate
	}
	if !utils.IsNil(o.IsSortDirectionAsc) {
		toSerialize["IsSortDirectionAsc"] = o.IsSortDirectionAsc
	}
	if !utils.IsNil(o.KeyWord) {
		toSerialize["KeyWord"] = o.KeyWord
	}
	if !utils.IsNil(o.Label) {
		toSerialize["Label"] = o.Label
	}
	if !utils.IsNil(o.MergerEmail) {
		toSerialize["MergerEmail"] = o.MergerEmail
	}
	if !utils.IsNil(o.MergerGlobalKey) {
		toSerialize["MergerGlobalKey"] = o.MergerGlobalKey
	}
	if !utils.IsNil(o.PageNumber) {
		toSerialize["PageNumber"] = o.PageNumber
	}
	if !utils.IsNil(o.PageSize) {
		toSerialize["PageSize"] = o.PageSize
	}
	if !utils.IsNil(o.ReviewerEmail) {
		toSerialize["ReviewerEmail"] = o.ReviewerEmail
	}
	if !utils.IsNil(o.ReviewerGlobalKey) {
		toSerialize["ReviewerGlobalKey"] = o.ReviewerGlobalKey
	}
	if !utils.IsNil(o.Sort) {
		toSerialize["Sort"] = o.Sort
	}
	if !utils.IsNil(o.SourceBranch) {
		toSerialize["SourceBranch"] = o.SourceBranch
	}
	if !utils.IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !utils.IsNil(o.TargetBranch) {
		toSerialize["TargetBranch"] = o.TargetBranch
	}
	if !utils.IsNil(o.UpdatedAtEndDate) {
		toSerialize["UpdatedAtEndDate"] = o.UpdatedAtEndDate
	}
	if !utils.IsNil(o.UpdatedAtStartDate) {
		toSerialize["UpdatedAtStartDate"] = o.UpdatedAtStartDate
	}
	return toSerialize, nil
}

type NullableDescribeSelfMergeRequestsRequest struct {
	value *DescribeSelfMergeRequestsRequest
	isSet bool
}

func (v NullableDescribeSelfMergeRequestsRequest) Get() *DescribeSelfMergeRequestsRequest {
	return v.value
}

func (v *NullableDescribeSelfMergeRequestsRequest) Set(val *DescribeSelfMergeRequestsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeSelfMergeRequestsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeSelfMergeRequestsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeSelfMergeRequestsRequest(val *DescribeSelfMergeRequestsRequest) *NullableDescribeSelfMergeRequestsRequest {
	return &NullableDescribeSelfMergeRequestsRequest{value: val, isSet: true}
}

func (v NullableDescribeSelfMergeRequestsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeSelfMergeRequestsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

