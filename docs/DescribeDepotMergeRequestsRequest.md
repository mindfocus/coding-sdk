# DescribeDepotMergeRequestsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAtEndDate** | Pointer to **string** | mr创建结束时间 | [optional] 
**CreatedAtStartDate** | Pointer to **string** | mr创建开始时间 | [optional] 
**CreatorEmails** | Pointer to **[]string** | 合并请求创建者邮箱列表 | [optional] 
**CreatorGlobalKeys** | Pointer to **[]string** | 合并请求创建者 Global Key 列表 | [optional] 
**DepotId** | **int64** | 仓库 ID | 
**DepotPath** | Pointer to **string** | 仓库路径 | [optional] 
**IsSortDirectionAsc** | Pointer to **bool** | 是否升序 | [optional] 
**KeyWord** | Pointer to **string** | 关键词 | [optional] 
**Labels** | Pointer to **[]string** | 关联标签 | [optional] 
**MergerEmails** | Pointer to **[]string** | 合并请求合并者邮箱列表 | [optional] 
**MergerGlobalKeys** | Pointer to **[]string** | 合并请求合并者 Global Key 列表 | [optional] 
**PageNumber** | Pointer to **int64** | 页数 默认为1 | [optional] 
**PageSize** | Pointer to **int64** | 每页数量 默认为10 | [optional] 
**ReviewerEmails** | Pointer to **[]string** | 合并请求评审者邮箱列表 | [optional] 
**ReviewerGlobalKeys** | Pointer to **[]string** | 评审者 Global Key 列表 | [optional] 
**Sort** | Pointer to **string** | 排序 created_at merged_at action_at | [optional] 
**SourceBranches** | Pointer to **[]string** | 源分支 | [optional] 
**Status** | Pointer to **string** | 合并请求状态 OPEN CLOSE ALL ACCEPTED | [optional] 
**TargetBranches** | Pointer to **[]string** | 目标分支 | [optional] 
**UpdatedAtEndDate** | Pointer to **string** | mr更新结束时间 | [optional] 
**UpdatedAtStartDate** | Pointer to **string** | mr更新开始时间 | [optional] 

## Methods

### NewDescribeDepotMergeRequestsRequest

`func NewDescribeDepotMergeRequestsRequest(depotId int64, ) *DescribeDepotMergeRequestsRequest`

NewDescribeDepotMergeRequestsRequest instantiates a new DescribeDepotMergeRequestsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDepotMergeRequestsRequestWithDefaults

`func NewDescribeDepotMergeRequestsRequestWithDefaults() *DescribeDepotMergeRequestsRequest`

NewDescribeDepotMergeRequestsRequestWithDefaults instantiates a new DescribeDepotMergeRequestsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAtEndDate

`func (o *DescribeDepotMergeRequestsRequest) GetCreatedAtEndDate() string`

GetCreatedAtEndDate returns the CreatedAtEndDate field if non-nil, zero value otherwise.

### GetCreatedAtEndDateOk

`func (o *DescribeDepotMergeRequestsRequest) GetCreatedAtEndDateOk() (*string, bool)`

GetCreatedAtEndDateOk returns a tuple with the CreatedAtEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtEndDate

`func (o *DescribeDepotMergeRequestsRequest) SetCreatedAtEndDate(v string)`

SetCreatedAtEndDate sets CreatedAtEndDate field to given value.

### HasCreatedAtEndDate

`func (o *DescribeDepotMergeRequestsRequest) HasCreatedAtEndDate() bool`

HasCreatedAtEndDate returns a boolean if a field has been set.

### GetCreatedAtStartDate

`func (o *DescribeDepotMergeRequestsRequest) GetCreatedAtStartDate() string`

GetCreatedAtStartDate returns the CreatedAtStartDate field if non-nil, zero value otherwise.

### GetCreatedAtStartDateOk

`func (o *DescribeDepotMergeRequestsRequest) GetCreatedAtStartDateOk() (*string, bool)`

GetCreatedAtStartDateOk returns a tuple with the CreatedAtStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtStartDate

`func (o *DescribeDepotMergeRequestsRequest) SetCreatedAtStartDate(v string)`

SetCreatedAtStartDate sets CreatedAtStartDate field to given value.

### HasCreatedAtStartDate

`func (o *DescribeDepotMergeRequestsRequest) HasCreatedAtStartDate() bool`

HasCreatedAtStartDate returns a boolean if a field has been set.

### GetCreatorEmails

`func (o *DescribeDepotMergeRequestsRequest) GetCreatorEmails() []string`

GetCreatorEmails returns the CreatorEmails field if non-nil, zero value otherwise.

### GetCreatorEmailsOk

`func (o *DescribeDepotMergeRequestsRequest) GetCreatorEmailsOk() (*[]string, bool)`

GetCreatorEmailsOk returns a tuple with the CreatorEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmails

`func (o *DescribeDepotMergeRequestsRequest) SetCreatorEmails(v []string)`

SetCreatorEmails sets CreatorEmails field to given value.

### HasCreatorEmails

`func (o *DescribeDepotMergeRequestsRequest) HasCreatorEmails() bool`

HasCreatorEmails returns a boolean if a field has been set.

### GetCreatorGlobalKeys

`func (o *DescribeDepotMergeRequestsRequest) GetCreatorGlobalKeys() []string`

GetCreatorGlobalKeys returns the CreatorGlobalKeys field if non-nil, zero value otherwise.

### GetCreatorGlobalKeysOk

`func (o *DescribeDepotMergeRequestsRequest) GetCreatorGlobalKeysOk() (*[]string, bool)`

GetCreatorGlobalKeysOk returns a tuple with the CreatorGlobalKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorGlobalKeys

`func (o *DescribeDepotMergeRequestsRequest) SetCreatorGlobalKeys(v []string)`

SetCreatorGlobalKeys sets CreatorGlobalKeys field to given value.

### HasCreatorGlobalKeys

`func (o *DescribeDepotMergeRequestsRequest) HasCreatorGlobalKeys() bool`

HasCreatorGlobalKeys returns a boolean if a field has been set.

### GetDepotId

`func (o *DescribeDepotMergeRequestsRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeDepotMergeRequestsRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeDepotMergeRequestsRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeDepotMergeRequestsRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeDepotMergeRequestsRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeDepotMergeRequestsRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeDepotMergeRequestsRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetIsSortDirectionAsc

`func (o *DescribeDepotMergeRequestsRequest) GetIsSortDirectionAsc() bool`

GetIsSortDirectionAsc returns the IsSortDirectionAsc field if non-nil, zero value otherwise.

### GetIsSortDirectionAscOk

`func (o *DescribeDepotMergeRequestsRequest) GetIsSortDirectionAscOk() (*bool, bool)`

GetIsSortDirectionAscOk returns a tuple with the IsSortDirectionAsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSortDirectionAsc

`func (o *DescribeDepotMergeRequestsRequest) SetIsSortDirectionAsc(v bool)`

SetIsSortDirectionAsc sets IsSortDirectionAsc field to given value.

### HasIsSortDirectionAsc

`func (o *DescribeDepotMergeRequestsRequest) HasIsSortDirectionAsc() bool`

HasIsSortDirectionAsc returns a boolean if a field has been set.

### GetKeyWord

`func (o *DescribeDepotMergeRequestsRequest) GetKeyWord() string`

GetKeyWord returns the KeyWord field if non-nil, zero value otherwise.

### GetKeyWordOk

`func (o *DescribeDepotMergeRequestsRequest) GetKeyWordOk() (*string, bool)`

GetKeyWordOk returns a tuple with the KeyWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWord

`func (o *DescribeDepotMergeRequestsRequest) SetKeyWord(v string)`

SetKeyWord sets KeyWord field to given value.

### HasKeyWord

`func (o *DescribeDepotMergeRequestsRequest) HasKeyWord() bool`

HasKeyWord returns a boolean if a field has been set.

### GetLabels

`func (o *DescribeDepotMergeRequestsRequest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *DescribeDepotMergeRequestsRequest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *DescribeDepotMergeRequestsRequest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *DescribeDepotMergeRequestsRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMergerEmails

`func (o *DescribeDepotMergeRequestsRequest) GetMergerEmails() []string`

GetMergerEmails returns the MergerEmails field if non-nil, zero value otherwise.

### GetMergerEmailsOk

`func (o *DescribeDepotMergeRequestsRequest) GetMergerEmailsOk() (*[]string, bool)`

GetMergerEmailsOk returns a tuple with the MergerEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergerEmails

`func (o *DescribeDepotMergeRequestsRequest) SetMergerEmails(v []string)`

SetMergerEmails sets MergerEmails field to given value.

### HasMergerEmails

`func (o *DescribeDepotMergeRequestsRequest) HasMergerEmails() bool`

HasMergerEmails returns a boolean if a field has been set.

### GetMergerGlobalKeys

`func (o *DescribeDepotMergeRequestsRequest) GetMergerGlobalKeys() []string`

GetMergerGlobalKeys returns the MergerGlobalKeys field if non-nil, zero value otherwise.

### GetMergerGlobalKeysOk

`func (o *DescribeDepotMergeRequestsRequest) GetMergerGlobalKeysOk() (*[]string, bool)`

GetMergerGlobalKeysOk returns a tuple with the MergerGlobalKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergerGlobalKeys

`func (o *DescribeDepotMergeRequestsRequest) SetMergerGlobalKeys(v []string)`

SetMergerGlobalKeys sets MergerGlobalKeys field to given value.

### HasMergerGlobalKeys

`func (o *DescribeDepotMergeRequestsRequest) HasMergerGlobalKeys() bool`

HasMergerGlobalKeys returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeDepotMergeRequestsRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeDepotMergeRequestsRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeDepotMergeRequestsRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeDepotMergeRequestsRequest) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeDepotMergeRequestsRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeDepotMergeRequestsRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeDepotMergeRequestsRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeDepotMergeRequestsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetReviewerEmails

`func (o *DescribeDepotMergeRequestsRequest) GetReviewerEmails() []string`

GetReviewerEmails returns the ReviewerEmails field if non-nil, zero value otherwise.

### GetReviewerEmailsOk

`func (o *DescribeDepotMergeRequestsRequest) GetReviewerEmailsOk() (*[]string, bool)`

GetReviewerEmailsOk returns a tuple with the ReviewerEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerEmails

`func (o *DescribeDepotMergeRequestsRequest) SetReviewerEmails(v []string)`

SetReviewerEmails sets ReviewerEmails field to given value.

### HasReviewerEmails

`func (o *DescribeDepotMergeRequestsRequest) HasReviewerEmails() bool`

HasReviewerEmails returns a boolean if a field has been set.

### GetReviewerGlobalKeys

`func (o *DescribeDepotMergeRequestsRequest) GetReviewerGlobalKeys() []string`

GetReviewerGlobalKeys returns the ReviewerGlobalKeys field if non-nil, zero value otherwise.

### GetReviewerGlobalKeysOk

`func (o *DescribeDepotMergeRequestsRequest) GetReviewerGlobalKeysOk() (*[]string, bool)`

GetReviewerGlobalKeysOk returns a tuple with the ReviewerGlobalKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGlobalKeys

`func (o *DescribeDepotMergeRequestsRequest) SetReviewerGlobalKeys(v []string)`

SetReviewerGlobalKeys sets ReviewerGlobalKeys field to given value.

### HasReviewerGlobalKeys

`func (o *DescribeDepotMergeRequestsRequest) HasReviewerGlobalKeys() bool`

HasReviewerGlobalKeys returns a boolean if a field has been set.

### GetSort

`func (o *DescribeDepotMergeRequestsRequest) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *DescribeDepotMergeRequestsRequest) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *DescribeDepotMergeRequestsRequest) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *DescribeDepotMergeRequestsRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetSourceBranches

`func (o *DescribeDepotMergeRequestsRequest) GetSourceBranches() []string`

GetSourceBranches returns the SourceBranches field if non-nil, zero value otherwise.

### GetSourceBranchesOk

`func (o *DescribeDepotMergeRequestsRequest) GetSourceBranchesOk() (*[]string, bool)`

GetSourceBranchesOk returns a tuple with the SourceBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBranches

`func (o *DescribeDepotMergeRequestsRequest) SetSourceBranches(v []string)`

SetSourceBranches sets SourceBranches field to given value.

### HasSourceBranches

`func (o *DescribeDepotMergeRequestsRequest) HasSourceBranches() bool`

HasSourceBranches returns a boolean if a field has been set.

### GetStatus

`func (o *DescribeDepotMergeRequestsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeDepotMergeRequestsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeDepotMergeRequestsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DescribeDepotMergeRequestsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetBranches

`func (o *DescribeDepotMergeRequestsRequest) GetTargetBranches() []string`

GetTargetBranches returns the TargetBranches field if non-nil, zero value otherwise.

### GetTargetBranchesOk

`func (o *DescribeDepotMergeRequestsRequest) GetTargetBranchesOk() (*[]string, bool)`

GetTargetBranchesOk returns a tuple with the TargetBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBranches

`func (o *DescribeDepotMergeRequestsRequest) SetTargetBranches(v []string)`

SetTargetBranches sets TargetBranches field to given value.

### HasTargetBranches

`func (o *DescribeDepotMergeRequestsRequest) HasTargetBranches() bool`

HasTargetBranches returns a boolean if a field has been set.

### GetUpdatedAtEndDate

`func (o *DescribeDepotMergeRequestsRequest) GetUpdatedAtEndDate() string`

GetUpdatedAtEndDate returns the UpdatedAtEndDate field if non-nil, zero value otherwise.

### GetUpdatedAtEndDateOk

`func (o *DescribeDepotMergeRequestsRequest) GetUpdatedAtEndDateOk() (*string, bool)`

GetUpdatedAtEndDateOk returns a tuple with the UpdatedAtEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtEndDate

`func (o *DescribeDepotMergeRequestsRequest) SetUpdatedAtEndDate(v string)`

SetUpdatedAtEndDate sets UpdatedAtEndDate field to given value.

### HasUpdatedAtEndDate

`func (o *DescribeDepotMergeRequestsRequest) HasUpdatedAtEndDate() bool`

HasUpdatedAtEndDate returns a boolean if a field has been set.

### GetUpdatedAtStartDate

`func (o *DescribeDepotMergeRequestsRequest) GetUpdatedAtStartDate() string`

GetUpdatedAtStartDate returns the UpdatedAtStartDate field if non-nil, zero value otherwise.

### GetUpdatedAtStartDateOk

`func (o *DescribeDepotMergeRequestsRequest) GetUpdatedAtStartDateOk() (*string, bool)`

GetUpdatedAtStartDateOk returns a tuple with the UpdatedAtStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtStartDate

`func (o *DescribeDepotMergeRequestsRequest) SetUpdatedAtStartDate(v string)`

SetUpdatedAtStartDate sets UpdatedAtStartDate field to given value.

### HasUpdatedAtStartDate

`func (o *DescribeDepotMergeRequestsRequest) HasUpdatedAtStartDate() bool`

HasUpdatedAtStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


