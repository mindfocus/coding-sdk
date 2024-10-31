# DescribeProjectMergeRequestsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAtEndDate** | Pointer to **string** | mr创建结束时间 | [optional] 
**CreatedAtStartDate** | Pointer to **string** | mr创建开始时间 | [optional] 
**CreatorEmail** | Pointer to **string** | 合并请求创建者邮箱 | [optional] 
**CreatorGlobalKey** | Pointer to **string** | 合并请求创建者全局 key | [optional] 
**IsSortDirectionAsc** | Pointer to **bool** | 是否正序排序 默认 false | [optional] 
**KeyWord** | Pointer to **string** | 关键词 | [optional] 
**Label** | Pointer to **string** | 标签 | [optional] 
**MergerEmail** | Pointer to **string** | 合并请求合并者邮箱 | [optional] 
**MergerGlobalKey** | Pointer to **string** | 合并请求合并者全局 key | [optional] 
**PageNumber** | Pointer to **int64** | 页数 默认 1 | [optional] 
**PageSize** | Pointer to **int64** | 每页条数 默认 10 | [optional] 
**ProjectId** | **int64** | 项目 Id | 
**ReviewerEmail** | Pointer to **string** | 合并请求评审者邮箱 | [optional] 
**ReviewerGlobalKey** | Pointer to **string** | 合并请求评审者全局 key | [optional] 
**Sort** | Pointer to **string** | 排序 action_at：以更新时间排序  created_at：以创建时间排序 | [optional] 
**SourceBranch** | Pointer to **string** | 源分支 | [optional] 
**Status** | Pointer to **string** | 合并请求状态 open/close/all | [optional] 
**TargetBranch** | Pointer to **string** | 欲合入分支 | [optional] 
**UpdatedAtEndDate** | Pointer to **string** | mr更新结束时间 | [optional] 
**UpdatedAtStartDate** | Pointer to **string** | mr更新开始时间 | [optional] 

## Methods

### NewDescribeProjectMergeRequestsRequest

`func NewDescribeProjectMergeRequestsRequest(projectId int64, ) *DescribeProjectMergeRequestsRequest`

NewDescribeProjectMergeRequestsRequest instantiates a new DescribeProjectMergeRequestsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProjectMergeRequestsRequestWithDefaults

`func NewDescribeProjectMergeRequestsRequestWithDefaults() *DescribeProjectMergeRequestsRequest`

NewDescribeProjectMergeRequestsRequestWithDefaults instantiates a new DescribeProjectMergeRequestsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAtEndDate

`func (o *DescribeProjectMergeRequestsRequest) GetCreatedAtEndDate() string`

GetCreatedAtEndDate returns the CreatedAtEndDate field if non-nil, zero value otherwise.

### GetCreatedAtEndDateOk

`func (o *DescribeProjectMergeRequestsRequest) GetCreatedAtEndDateOk() (*string, bool)`

GetCreatedAtEndDateOk returns a tuple with the CreatedAtEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtEndDate

`func (o *DescribeProjectMergeRequestsRequest) SetCreatedAtEndDate(v string)`

SetCreatedAtEndDate sets CreatedAtEndDate field to given value.

### HasCreatedAtEndDate

`func (o *DescribeProjectMergeRequestsRequest) HasCreatedAtEndDate() bool`

HasCreatedAtEndDate returns a boolean if a field has been set.

### GetCreatedAtStartDate

`func (o *DescribeProjectMergeRequestsRequest) GetCreatedAtStartDate() string`

GetCreatedAtStartDate returns the CreatedAtStartDate field if non-nil, zero value otherwise.

### GetCreatedAtStartDateOk

`func (o *DescribeProjectMergeRequestsRequest) GetCreatedAtStartDateOk() (*string, bool)`

GetCreatedAtStartDateOk returns a tuple with the CreatedAtStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtStartDate

`func (o *DescribeProjectMergeRequestsRequest) SetCreatedAtStartDate(v string)`

SetCreatedAtStartDate sets CreatedAtStartDate field to given value.

### HasCreatedAtStartDate

`func (o *DescribeProjectMergeRequestsRequest) HasCreatedAtStartDate() bool`

HasCreatedAtStartDate returns a boolean if a field has been set.

### GetCreatorEmail

`func (o *DescribeProjectMergeRequestsRequest) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *DescribeProjectMergeRequestsRequest) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *DescribeProjectMergeRequestsRequest) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *DescribeProjectMergeRequestsRequest) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetCreatorGlobalKey

`func (o *DescribeProjectMergeRequestsRequest) GetCreatorGlobalKey() string`

GetCreatorGlobalKey returns the CreatorGlobalKey field if non-nil, zero value otherwise.

### GetCreatorGlobalKeyOk

`func (o *DescribeProjectMergeRequestsRequest) GetCreatorGlobalKeyOk() (*string, bool)`

GetCreatorGlobalKeyOk returns a tuple with the CreatorGlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorGlobalKey

`func (o *DescribeProjectMergeRequestsRequest) SetCreatorGlobalKey(v string)`

SetCreatorGlobalKey sets CreatorGlobalKey field to given value.

### HasCreatorGlobalKey

`func (o *DescribeProjectMergeRequestsRequest) HasCreatorGlobalKey() bool`

HasCreatorGlobalKey returns a boolean if a field has been set.

### GetIsSortDirectionAsc

`func (o *DescribeProjectMergeRequestsRequest) GetIsSortDirectionAsc() bool`

GetIsSortDirectionAsc returns the IsSortDirectionAsc field if non-nil, zero value otherwise.

### GetIsSortDirectionAscOk

`func (o *DescribeProjectMergeRequestsRequest) GetIsSortDirectionAscOk() (*bool, bool)`

GetIsSortDirectionAscOk returns a tuple with the IsSortDirectionAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSortDirectionAsc

`func (o *DescribeProjectMergeRequestsRequest) SetIsSortDirectionAsc(v bool)`

SetIsSortDirectionAsc sets IsSortDirectionAsc field to given value.

### HasIsSortDirectionAsc

`func (o *DescribeProjectMergeRequestsRequest) HasIsSortDirectionAsc() bool`

HasIsSortDirectionAsc returns a boolean if a field has been set.

### GetKeyWord

`func (o *DescribeProjectMergeRequestsRequest) GetKeyWord() string`

GetKeyWord returns the KeyWord field if non-nil, zero value otherwise.

### GetKeyWordOk

`func (o *DescribeProjectMergeRequestsRequest) GetKeyWordOk() (*string, bool)`

GetKeyWordOk returns a tuple with the KeyWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWord

`func (o *DescribeProjectMergeRequestsRequest) SetKeyWord(v string)`

SetKeyWord sets KeyWord field to given value.

### HasKeyWord

`func (o *DescribeProjectMergeRequestsRequest) HasKeyWord() bool`

HasKeyWord returns a boolean if a field has been set.

### GetLabel

`func (o *DescribeProjectMergeRequestsRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DescribeProjectMergeRequestsRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DescribeProjectMergeRequestsRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DescribeProjectMergeRequestsRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMergerEmail

`func (o *DescribeProjectMergeRequestsRequest) GetMergerEmail() string`

GetMergerEmail returns the MergerEmail field if non-nil, zero value otherwise.

### GetMergerEmailOk

`func (o *DescribeProjectMergeRequestsRequest) GetMergerEmailOk() (*string, bool)`

GetMergerEmailOk returns a tuple with the MergerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergerEmail

`func (o *DescribeProjectMergeRequestsRequest) SetMergerEmail(v string)`

SetMergerEmail sets MergerEmail field to given value.

### HasMergerEmail

`func (o *DescribeProjectMergeRequestsRequest) HasMergerEmail() bool`

HasMergerEmail returns a boolean if a field has been set.

### GetMergerGlobalKey

`func (o *DescribeProjectMergeRequestsRequest) GetMergerGlobalKey() string`

GetMergerGlobalKey returns the MergerGlobalKey field if non-nil, zero value otherwise.

### GetMergerGlobalKeyOk

`func (o *DescribeProjectMergeRequestsRequest) GetMergerGlobalKeyOk() (*string, bool)`

GetMergerGlobalKeyOk returns a tuple with the MergerGlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergerGlobalKey

`func (o *DescribeProjectMergeRequestsRequest) SetMergerGlobalKey(v string)`

SetMergerGlobalKey sets MergerGlobalKey field to given value.

### HasMergerGlobalKey

`func (o *DescribeProjectMergeRequestsRequest) HasMergerGlobalKey() bool`

HasMergerGlobalKey returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeProjectMergeRequestsRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeProjectMergeRequestsRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeProjectMergeRequestsRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeProjectMergeRequestsRequest) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeProjectMergeRequestsRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeProjectMergeRequestsRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeProjectMergeRequestsRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeProjectMergeRequestsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetProjectId

`func (o *DescribeProjectMergeRequestsRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeProjectMergeRequestsRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeProjectMergeRequestsRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetReviewerEmail

`func (o *DescribeProjectMergeRequestsRequest) GetReviewerEmail() string`

GetReviewerEmail returns the ReviewerEmail field if non-nil, zero value otherwise.

### GetReviewerEmailOk

`func (o *DescribeProjectMergeRequestsRequest) GetReviewerEmailOk() (*string, bool)`

GetReviewerEmailOk returns a tuple with the ReviewerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerEmail

`func (o *DescribeProjectMergeRequestsRequest) SetReviewerEmail(v string)`

SetReviewerEmail sets ReviewerEmail field to given value.

### HasReviewerEmail

`func (o *DescribeProjectMergeRequestsRequest) HasReviewerEmail() bool`

HasReviewerEmail returns a boolean if a field has been set.

### GetReviewerGlobalKey

`func (o *DescribeProjectMergeRequestsRequest) GetReviewerGlobalKey() string`

GetReviewerGlobalKey returns the ReviewerGlobalKey field if non-nil, zero value otherwise.

### GetReviewerGlobalKeyOk

`func (o *DescribeProjectMergeRequestsRequest) GetReviewerGlobalKeyOk() (*string, bool)`

GetReviewerGlobalKeyOk returns a tuple with the ReviewerGlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGlobalKey

`func (o *DescribeProjectMergeRequestsRequest) SetReviewerGlobalKey(v string)`

SetReviewerGlobalKey sets ReviewerGlobalKey field to given value.

### HasReviewerGlobalKey

`func (o *DescribeProjectMergeRequestsRequest) HasReviewerGlobalKey() bool`

HasReviewerGlobalKey returns a boolean if a field has been set.

### GetSort

`func (o *DescribeProjectMergeRequestsRequest) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *DescribeProjectMergeRequestsRequest) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *DescribeProjectMergeRequestsRequest) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *DescribeProjectMergeRequestsRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetSourceBranch

`func (o *DescribeProjectMergeRequestsRequest) GetSourceBranch() string`

GetSourceBranch returns the SourceBranch field if non-nil, zero value otherwise.

### GetSourceBranchOk

`func (o *DescribeProjectMergeRequestsRequest) GetSourceBranchOk() (*string, bool)`

GetSourceBranchOk returns a tuple with the SourceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBranch

`func (o *DescribeProjectMergeRequestsRequest) SetSourceBranch(v string)`

SetSourceBranch sets SourceBranch field to given value.

### HasSourceBranch

`func (o *DescribeProjectMergeRequestsRequest) HasSourceBranch() bool`

HasSourceBranch returns a boolean if a field has been set.

### GetStatus

`func (o *DescribeProjectMergeRequestsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeProjectMergeRequestsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeProjectMergeRequestsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DescribeProjectMergeRequestsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetBranch

`func (o *DescribeProjectMergeRequestsRequest) GetTargetBranch() string`

GetTargetBranch returns the TargetBranch field if non-nil, zero value otherwise.

### GetTargetBranchOk

`func (o *DescribeProjectMergeRequestsRequest) GetTargetBranchOk() (*string, bool)`

GetTargetBranchOk returns a tuple with the TargetBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBranch

`func (o *DescribeProjectMergeRequestsRequest) SetTargetBranch(v string)`

SetTargetBranch sets TargetBranch field to given value.

### HasTargetBranch

`func (o *DescribeProjectMergeRequestsRequest) HasTargetBranch() bool`

HasTargetBranch returns a boolean if a field has been set.

### GetUpdatedAtEndDate

`func (o *DescribeProjectMergeRequestsRequest) GetUpdatedAtEndDate() string`

GetUpdatedAtEndDate returns the UpdatedAtEndDate field if non-nil, zero value otherwise.

### GetUpdatedAtEndDateOk

`func (o *DescribeProjectMergeRequestsRequest) GetUpdatedAtEndDateOk() (*string, bool)`

GetUpdatedAtEndDateOk returns a tuple with the UpdatedAtEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtEndDate

`func (o *DescribeProjectMergeRequestsRequest) SetUpdatedAtEndDate(v string)`

SetUpdatedAtEndDate sets UpdatedAtEndDate field to given value.

### HasUpdatedAtEndDate

`func (o *DescribeProjectMergeRequestsRequest) HasUpdatedAtEndDate() bool`

HasUpdatedAtEndDate returns a boolean if a field has been set.

### GetUpdatedAtStartDate

`func (o *DescribeProjectMergeRequestsRequest) GetUpdatedAtStartDate() string`

GetUpdatedAtStartDate returns the UpdatedAtStartDate field if non-nil, zero value otherwise.

### GetUpdatedAtStartDateOk

`func (o *DescribeProjectMergeRequestsRequest) GetUpdatedAtStartDateOk() (*string, bool)`

GetUpdatedAtStartDateOk returns a tuple with the UpdatedAtStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtStartDate

`func (o *DescribeProjectMergeRequestsRequest) SetUpdatedAtStartDate(v string)`

SetUpdatedAtStartDate sets UpdatedAtStartDate field to given value.

### HasUpdatedAtStartDate

`func (o *DescribeProjectMergeRequestsRequest) HasUpdatedAtStartDate() bool`

HasUpdatedAtStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


