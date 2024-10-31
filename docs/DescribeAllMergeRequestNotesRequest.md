# DescribeAllMergeRequestNotesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAtEndDate** | Pointer to **string** | 创建结束时间 | [optional] 
**CreatedAtStartDate** | Pointer to **string** | 创建开始时间 | [optional] 
**DepotPath** | **string** | 仓库路径 | 
**MergeIds** | Pointer to **[]int32** | 合并请求的Iid列表 | [optional] 
**MrStatuses** | Pointer to **[]string** | 合并请求的状态列表 | [optional] 
**PageNumber** | **int32** | 页码 | 
**PageSize** | **int32** | 页数量 | 
**ReporterIds** | Pointer to **[]int32** | 合并请求的发起者列表 | [optional] 

## Methods

### NewDescribeAllMergeRequestNotesRequest

`func NewDescribeAllMergeRequestNotesRequest(depotPath string, pageNumber int32, pageSize int32, ) *DescribeAllMergeRequestNotesRequest`

NewDescribeAllMergeRequestNotesRequest instantiates a new DescribeAllMergeRequestNotesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAllMergeRequestNotesRequestWithDefaults

`func NewDescribeAllMergeRequestNotesRequestWithDefaults() *DescribeAllMergeRequestNotesRequest`

NewDescribeAllMergeRequestNotesRequestWithDefaults instantiates a new DescribeAllMergeRequestNotesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAtEndDate

`func (o *DescribeAllMergeRequestNotesRequest) GetCreatedAtEndDate() string`

GetCreatedAtEndDate returns the CreatedAtEndDate field if non-nil, zero value otherwise.

### GetCreatedAtEndDateOk

`func (o *DescribeAllMergeRequestNotesRequest) GetCreatedAtEndDateOk() (*string, bool)`

GetCreatedAtEndDateOk returns a tuple with the CreatedAtEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtEndDate

`func (o *DescribeAllMergeRequestNotesRequest) SetCreatedAtEndDate(v string)`

SetCreatedAtEndDate sets CreatedAtEndDate field to given value.

### HasCreatedAtEndDate

`func (o *DescribeAllMergeRequestNotesRequest) HasCreatedAtEndDate() bool`

HasCreatedAtEndDate returns a boolean if a field has been set.

### GetCreatedAtStartDate

`func (o *DescribeAllMergeRequestNotesRequest) GetCreatedAtStartDate() string`

GetCreatedAtStartDate returns the CreatedAtStartDate field if non-nil, zero value otherwise.

### GetCreatedAtStartDateOk

`func (o *DescribeAllMergeRequestNotesRequest) GetCreatedAtStartDateOk() (*string, bool)`

GetCreatedAtStartDateOk returns a tuple with the CreatedAtStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtStartDate

`func (o *DescribeAllMergeRequestNotesRequest) SetCreatedAtStartDate(v string)`

SetCreatedAtStartDate sets CreatedAtStartDate field to given value.

### HasCreatedAtStartDate

`func (o *DescribeAllMergeRequestNotesRequest) HasCreatedAtStartDate() bool`

HasCreatedAtStartDate returns a boolean if a field has been set.

### GetDepotPath

`func (o *DescribeAllMergeRequestNotesRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeAllMergeRequestNotesRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeAllMergeRequestNotesRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetMergeIds

`func (o *DescribeAllMergeRequestNotesRequest) GetMergeIds() []int32`

GetMergeIds returns the MergeIds field if non-nil, zero value otherwise.

### GetMergeIdsOk

`func (o *DescribeAllMergeRequestNotesRequest) GetMergeIdsOk() (*[]int32, bool)`

GetMergeIdsOk returns a tuple with the MergeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeIds

`func (o *DescribeAllMergeRequestNotesRequest) SetMergeIds(v []int32)`

SetMergeIds sets MergeIds field to given value.

### HasMergeIds

`func (o *DescribeAllMergeRequestNotesRequest) HasMergeIds() bool`

HasMergeIds returns a boolean if a field has been set.

### GetMrStatuses

`func (o *DescribeAllMergeRequestNotesRequest) GetMrStatuses() []string`

GetMrStatuses returns the MrStatuses field if non-nil, zero value otherwise.

### GetMrStatusesOk

`func (o *DescribeAllMergeRequestNotesRequest) GetMrStatusesOk() (*[]string, bool)`

GetMrStatusesOk returns a tuple with the MrStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrStatuses

`func (o *DescribeAllMergeRequestNotesRequest) SetMrStatuses(v []string)`

SetMrStatuses sets MrStatuses field to given value.

### HasMrStatuses

`func (o *DescribeAllMergeRequestNotesRequest) HasMrStatuses() bool`

HasMrStatuses returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeAllMergeRequestNotesRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeAllMergeRequestNotesRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeAllMergeRequestNotesRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeAllMergeRequestNotesRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeAllMergeRequestNotesRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeAllMergeRequestNotesRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetReporterIds

`func (o *DescribeAllMergeRequestNotesRequest) GetReporterIds() []int32`

GetReporterIds returns the ReporterIds field if non-nil, zero value otherwise.

### GetReporterIdsOk

`func (o *DescribeAllMergeRequestNotesRequest) GetReporterIdsOk() (*[]int32, bool)`

GetReporterIdsOk returns a tuple with the ReporterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReporterIds

`func (o *DescribeAllMergeRequestNotesRequest) SetReporterIds(v []int32)`

SetReporterIds sets ReporterIds field to given value.

### HasReporterIds

`func (o *DescribeAllMergeRequestNotesRequest) HasReporterIds() bool`

HasReporterIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


