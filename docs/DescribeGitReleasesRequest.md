# DescribeGitReleasesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**FromDate** | Pointer to **string** | 搜索条件:起始日期 | [optional] 
**PageNumber** | Pointer to **int64** | 页数 | [optional] 
**PageSize** | Pointer to **int64** | 每页条数 | [optional] 
**Status** | Pointer to **int64** | 搜索条件:版本状态(0:全部 1:已发布 2:预发布) | [optional] 
**TagName** | Pointer to **string** | 搜索条件:标签名字 | [optional] 
**ToDate** | Pointer to **string** | 搜索条件:终止日期 | [optional] 

## Methods

### NewDescribeGitReleasesRequest

`func NewDescribeGitReleasesRequest(depotId int64, ) *DescribeGitReleasesRequest`

NewDescribeGitReleasesRequest instantiates a new DescribeGitReleasesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitReleasesRequestWithDefaults

`func NewDescribeGitReleasesRequestWithDefaults() *DescribeGitReleasesRequest`

NewDescribeGitReleasesRequestWithDefaults instantiates a new DescribeGitReleasesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeGitReleasesRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitReleasesRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitReleasesRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitReleasesRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitReleasesRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitReleasesRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitReleasesRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetFromDate

`func (o *DescribeGitReleasesRequest) GetFromDate() string`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *DescribeGitReleasesRequest) GetFromDateOk() (*string, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *DescribeGitReleasesRequest) SetFromDate(v string)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *DescribeGitReleasesRequest) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeGitReleasesRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeGitReleasesRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeGitReleasesRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeGitReleasesRequest) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeGitReleasesRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeGitReleasesRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeGitReleasesRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeGitReleasesRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetStatus

`func (o *DescribeGitReleasesRequest) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeGitReleasesRequest) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeGitReleasesRequest) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DescribeGitReleasesRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTagName

`func (o *DescribeGitReleasesRequest) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *DescribeGitReleasesRequest) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *DescribeGitReleasesRequest) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *DescribeGitReleasesRequest) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### GetToDate

`func (o *DescribeGitReleasesRequest) GetToDate() string`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *DescribeGitReleasesRequest) GetToDateOk() (*string, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *DescribeGitReleasesRequest) SetToDate(v string)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *DescribeGitReleasesRequest) HasToDate() bool`

HasToDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


