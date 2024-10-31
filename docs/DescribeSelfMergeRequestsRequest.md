# DescribeSelfMergeRequestsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAtEndDate** | Pointer to **string** | mr创建结束时间 | [optional] 
**CreatedAtStartDate** | Pointer to **string** | mr创建开始时间 | [optional] 
**IsSortDirectionAsc** | Pointer to **bool** | 是否正序排序 默认 false | [optional] 
**KeyWord** | Pointer to **string** | 关键词 | [optional] 
**Label** | Pointer to **string** | 标签 | [optional] 
**MergerEmail** | Pointer to **string** | 合并请求合并者邮箱 | [optional] 
**MergerGlobalKey** | Pointer to **string** | 合并请求合并者全局 key | [optional] 
**PageNumber** | Pointer to **int64** | 页数 默认 1 | [optional] 
**PageSize** | Pointer to **int64** | 每页条数 默认 10 | [optional] 
**ReviewerEmail** | Pointer to **string** | 合并请求评审者邮箱 | [optional] 
**ReviewerGlobalKey** | Pointer to **string** | 合并请求评审者全局 key | [optional] 
**Sort** | Pointer to **string** | 排序 action_at：以更新时间排序  created_at：以创建时间排序 | [optional] 
**SourceBranch** | Pointer to **string** | 源分支 | [optional] 
**Status** | Pointer to **string** | 合并请求状态 open/close/all | [optional] 
**TargetBranch** | Pointer to **string** | 欲合入分支 | [optional] 
**UpdatedAtEndDate** | Pointer to **string** | mr更新结束时间 | [optional] 
**UpdatedAtStartDate** | Pointer to **string** | mr更新开始时间 | [optional] 

## Methods

### NewDescribeSelfMergeRequestsRequest

`func NewDescribeSelfMergeRequestsRequest() *DescribeSelfMergeRequestsRequest`

NewDescribeSelfMergeRequestsRequest instantiates a new DescribeSelfMergeRequestsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeSelfMergeRequestsRequestWithDefaults

`func NewDescribeSelfMergeRequestsRequestWithDefaults() *DescribeSelfMergeRequestsRequest`

NewDescribeSelfMergeRequestsRequestWithDefaults instantiates a new DescribeSelfMergeRequestsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAtEndDate

`func (o *DescribeSelfMergeRequestsRequest) GetCreatedAtEndDate() string`

GetCreatedAtEndDate returns the CreatedAtEndDate field if non-nil, zero value otherwise.

### GetCreatedAtEndDateOk

`func (o *DescribeSelfMergeRequestsRequest) GetCreatedAtEndDateOk() (*string, bool)`

GetCreatedAtEndDateOk returns a tuple with the CreatedAtEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtEndDate

`func (o *DescribeSelfMergeRequestsRequest) SetCreatedAtEndDate(v string)`

SetCreatedAtEndDate sets CreatedAtEndDate field to given value.

### HasCreatedAtEndDate

`func (o *DescribeSelfMergeRequestsRequest) HasCreatedAtEndDate() bool`

HasCreatedAtEndDate returns a boolean if a field has been set.

### GetCreatedAtStartDate

`func (o *DescribeSelfMergeRequestsRequest) GetCreatedAtStartDate() string`

GetCreatedAtStartDate returns the CreatedAtStartDate field if non-nil, zero value otherwise.

### GetCreatedAtStartDateOk

`func (o *DescribeSelfMergeRequestsRequest) GetCreatedAtStartDateOk() (*string, bool)`

GetCreatedAtStartDateOk returns a tuple with the CreatedAtStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtStartDate

`func (o *DescribeSelfMergeRequestsRequest) SetCreatedAtStartDate(v string)`

SetCreatedAtStartDate sets CreatedAtStartDate field to given value.

### HasCreatedAtStartDate

`func (o *DescribeSelfMergeRequestsRequest) HasCreatedAtStartDate() bool`

HasCreatedAtStartDate returns a boolean if a field has been set.

### GetIsSortDirectionAsc

`func (o *DescribeSelfMergeRequestsRequest) GetIsSortDirectionAsc() bool`

GetIsSortDirectionAsc returns the IsSortDirectionAsc field if non-nil, zero value otherwise.

### GetIsSortDirectionAscOk

`func (o *DescribeSelfMergeRequestsRequest) GetIsSortDirectionAscOk() (*bool, bool)`

GetIsSortDirectionAscOk returns a tuple with the IsSortDirectionAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSortDirectionAsc

`func (o *DescribeSelfMergeRequestsRequest) SetIsSortDirectionAsc(v bool)`

SetIsSortDirectionAsc sets IsSortDirectionAsc field to given value.

### HasIsSortDirectionAsc

`func (o *DescribeSelfMergeRequestsRequest) HasIsSortDirectionAsc() bool`

HasIsSortDirectionAsc returns a boolean if a field has been set.

### GetKeyWord

`func (o *DescribeSelfMergeRequestsRequest) GetKeyWord() string`

GetKeyWord returns the KeyWord field if non-nil, zero value otherwise.

### GetKeyWordOk

`func (o *DescribeSelfMergeRequestsRequest) GetKeyWordOk() (*string, bool)`

GetKeyWordOk returns a tuple with the KeyWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWord

`func (o *DescribeSelfMergeRequestsRequest) SetKeyWord(v string)`

SetKeyWord sets KeyWord field to given value.

### HasKeyWord

`func (o *DescribeSelfMergeRequestsRequest) HasKeyWord() bool`

HasKeyWord returns a boolean if a field has been set.

### GetLabel

`func (o *DescribeSelfMergeRequestsRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DescribeSelfMergeRequestsRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DescribeSelfMergeRequestsRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DescribeSelfMergeRequestsRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMergerEmail

`func (o *DescribeSelfMergeRequestsRequest) GetMergerEmail() string`

GetMergerEmail returns the MergerEmail field if non-nil, zero value otherwise.

### GetMergerEmailOk

`func (o *DescribeSelfMergeRequestsRequest) GetMergerEmailOk() (*string, bool)`

GetMergerEmailOk returns a tuple with the MergerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergerEmail

`func (o *DescribeSelfMergeRequestsRequest) SetMergerEmail(v string)`

SetMergerEmail sets MergerEmail field to given value.

### HasMergerEmail

`func (o *DescribeSelfMergeRequestsRequest) HasMergerEmail() bool`

HasMergerEmail returns a boolean if a field has been set.

### GetMergerGlobalKey

`func (o *DescribeSelfMergeRequestsRequest) GetMergerGlobalKey() string`

GetMergerGlobalKey returns the MergerGlobalKey field if non-nil, zero value otherwise.

### GetMergerGlobalKeyOk

`func (o *DescribeSelfMergeRequestsRequest) GetMergerGlobalKeyOk() (*string, bool)`

GetMergerGlobalKeyOk returns a tuple with the MergerGlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergerGlobalKey

`func (o *DescribeSelfMergeRequestsRequest) SetMergerGlobalKey(v string)`

SetMergerGlobalKey sets MergerGlobalKey field to given value.

### HasMergerGlobalKey

`func (o *DescribeSelfMergeRequestsRequest) HasMergerGlobalKey() bool`

HasMergerGlobalKey returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeSelfMergeRequestsRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeSelfMergeRequestsRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeSelfMergeRequestsRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeSelfMergeRequestsRequest) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeSelfMergeRequestsRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeSelfMergeRequestsRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeSelfMergeRequestsRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeSelfMergeRequestsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetReviewerEmail

`func (o *DescribeSelfMergeRequestsRequest) GetReviewerEmail() string`

GetReviewerEmail returns the ReviewerEmail field if non-nil, zero value otherwise.

### GetReviewerEmailOk

`func (o *DescribeSelfMergeRequestsRequest) GetReviewerEmailOk() (*string, bool)`

GetReviewerEmailOk returns a tuple with the ReviewerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerEmail

`func (o *DescribeSelfMergeRequestsRequest) SetReviewerEmail(v string)`

SetReviewerEmail sets ReviewerEmail field to given value.

### HasReviewerEmail

`func (o *DescribeSelfMergeRequestsRequest) HasReviewerEmail() bool`

HasReviewerEmail returns a boolean if a field has been set.

### GetReviewerGlobalKey

`func (o *DescribeSelfMergeRequestsRequest) GetReviewerGlobalKey() string`

GetReviewerGlobalKey returns the ReviewerGlobalKey field if non-nil, zero value otherwise.

### GetReviewerGlobalKeyOk

`func (o *DescribeSelfMergeRequestsRequest) GetReviewerGlobalKeyOk() (*string, bool)`

GetReviewerGlobalKeyOk returns a tuple with the ReviewerGlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGlobalKey

`func (o *DescribeSelfMergeRequestsRequest) SetReviewerGlobalKey(v string)`

SetReviewerGlobalKey sets ReviewerGlobalKey field to given value.

### HasReviewerGlobalKey

`func (o *DescribeSelfMergeRequestsRequest) HasReviewerGlobalKey() bool`

HasReviewerGlobalKey returns a boolean if a field has been set.

### GetSort

`func (o *DescribeSelfMergeRequestsRequest) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *DescribeSelfMergeRequestsRequest) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *DescribeSelfMergeRequestsRequest) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *DescribeSelfMergeRequestsRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetSourceBranch

`func (o *DescribeSelfMergeRequestsRequest) GetSourceBranch() string`

GetSourceBranch returns the SourceBranch field if non-nil, zero value otherwise.

### GetSourceBranchOk

`func (o *DescribeSelfMergeRequestsRequest) GetSourceBranchOk() (*string, bool)`

GetSourceBranchOk returns a tuple with the SourceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBranch

`func (o *DescribeSelfMergeRequestsRequest) SetSourceBranch(v string)`

SetSourceBranch sets SourceBranch field to given value.

### HasSourceBranch

`func (o *DescribeSelfMergeRequestsRequest) HasSourceBranch() bool`

HasSourceBranch returns a boolean if a field has been set.

### GetStatus

`func (o *DescribeSelfMergeRequestsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeSelfMergeRequestsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeSelfMergeRequestsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DescribeSelfMergeRequestsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetBranch

`func (o *DescribeSelfMergeRequestsRequest) GetTargetBranch() string`

GetTargetBranch returns the TargetBranch field if non-nil, zero value otherwise.

### GetTargetBranchOk

`func (o *DescribeSelfMergeRequestsRequest) GetTargetBranchOk() (*string, bool)`

GetTargetBranchOk returns a tuple with the TargetBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBranch

`func (o *DescribeSelfMergeRequestsRequest) SetTargetBranch(v string)`

SetTargetBranch sets TargetBranch field to given value.

### HasTargetBranch

`func (o *DescribeSelfMergeRequestsRequest) HasTargetBranch() bool`

HasTargetBranch returns a boolean if a field has been set.

### GetUpdatedAtEndDate

`func (o *DescribeSelfMergeRequestsRequest) GetUpdatedAtEndDate() string`

GetUpdatedAtEndDate returns the UpdatedAtEndDate field if non-nil, zero value otherwise.

### GetUpdatedAtEndDateOk

`func (o *DescribeSelfMergeRequestsRequest) GetUpdatedAtEndDateOk() (*string, bool)`

GetUpdatedAtEndDateOk returns a tuple with the UpdatedAtEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtEndDate

`func (o *DescribeSelfMergeRequestsRequest) SetUpdatedAtEndDate(v string)`

SetUpdatedAtEndDate sets UpdatedAtEndDate field to given value.

### HasUpdatedAtEndDate

`func (o *DescribeSelfMergeRequestsRequest) HasUpdatedAtEndDate() bool`

HasUpdatedAtEndDate returns a boolean if a field has been set.

### GetUpdatedAtStartDate

`func (o *DescribeSelfMergeRequestsRequest) GetUpdatedAtStartDate() string`

GetUpdatedAtStartDate returns the UpdatedAtStartDate field if non-nil, zero value otherwise.

### GetUpdatedAtStartDateOk

`func (o *DescribeSelfMergeRequestsRequest) GetUpdatedAtStartDateOk() (*string, bool)`

GetUpdatedAtStartDateOk returns a tuple with the UpdatedAtStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtStartDate

`func (o *DescribeSelfMergeRequestsRequest) SetUpdatedAtStartDate(v string)`

SetUpdatedAtStartDate sets UpdatedAtStartDate field to given value.

### HasUpdatedAtStartDate

`func (o *DescribeSelfMergeRequestsRequest) HasUpdatedAtStartDate() bool`

HasUpdatedAtStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


