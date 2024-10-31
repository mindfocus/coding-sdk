# DescribeGitCommitsInPageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**EndDate** | Pointer to **string** | 查询截止日志 | [optional] 
**KeyWord** | Pointer to **string** | 提交信息中的关键词 | [optional] 
**PageNumber** | Pointer to **int64** | 分页页码 | [optional] 
**PageSize** | Pointer to **int64** | 分页页距 | [optional] 
**Path** | Pointer to **string** | 文件路径 | [optional] 
**Ref** | **string** | 分支名称 | 
**StartDate** | Pointer to **string** | 查询起始日期 | [optional] 

## Methods

### NewDescribeGitCommitsInPageRequest

`func NewDescribeGitCommitsInPageRequest(depotId int64, ref string, ) *DescribeGitCommitsInPageRequest`

NewDescribeGitCommitsInPageRequest instantiates a new DescribeGitCommitsInPageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitCommitsInPageRequestWithDefaults

`func NewDescribeGitCommitsInPageRequestWithDefaults() *DescribeGitCommitsInPageRequest`

NewDescribeGitCommitsInPageRequestWithDefaults instantiates a new DescribeGitCommitsInPageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeGitCommitsInPageRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitCommitsInPageRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitCommitsInPageRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitCommitsInPageRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitCommitsInPageRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitCommitsInPageRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitCommitsInPageRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetEndDate

`func (o *DescribeGitCommitsInPageRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DescribeGitCommitsInPageRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DescribeGitCommitsInPageRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DescribeGitCommitsInPageRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetKeyWord

`func (o *DescribeGitCommitsInPageRequest) GetKeyWord() string`

GetKeyWord returns the KeyWord field if non-nil, zero value otherwise.

### GetKeyWordOk

`func (o *DescribeGitCommitsInPageRequest) GetKeyWordOk() (*string, bool)`

GetKeyWordOk returns a tuple with the KeyWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWord

`func (o *DescribeGitCommitsInPageRequest) SetKeyWord(v string)`

SetKeyWord sets KeyWord field to given value.

### HasKeyWord

`func (o *DescribeGitCommitsInPageRequest) HasKeyWord() bool`

HasKeyWord returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeGitCommitsInPageRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeGitCommitsInPageRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeGitCommitsInPageRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeGitCommitsInPageRequest) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeGitCommitsInPageRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeGitCommitsInPageRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeGitCommitsInPageRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeGitCommitsInPageRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPath

`func (o *DescribeGitCommitsInPageRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DescribeGitCommitsInPageRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DescribeGitCommitsInPageRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DescribeGitCommitsInPageRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRef

`func (o *DescribeGitCommitsInPageRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DescribeGitCommitsInPageRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DescribeGitCommitsInPageRequest) SetRef(v string)`

SetRef sets Ref field to given value.


### GetStartDate

`func (o *DescribeGitCommitsInPageRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DescribeGitCommitsInPageRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DescribeGitCommitsInPageRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DescribeGitCommitsInPageRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


