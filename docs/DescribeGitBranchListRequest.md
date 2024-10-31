# DescribeGitBranchListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径 | 
**KeyWord** | Pointer to **string** | 查询的关键词 | [optional] 
**PageNumber** | **int64** | 分页页码 | 
**PageSize** | **int64** | 分页页距,默认为10 | 

## Methods

### NewDescribeGitBranchListRequest

`func NewDescribeGitBranchListRequest(depotPath string, pageNumber int64, pageSize int64, ) *DescribeGitBranchListRequest`

NewDescribeGitBranchListRequest instantiates a new DescribeGitBranchListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitBranchListRequestWithDefaults

`func NewDescribeGitBranchListRequestWithDefaults() *DescribeGitBranchListRequest`

NewDescribeGitBranchListRequestWithDefaults instantiates a new DescribeGitBranchListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *DescribeGitBranchListRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitBranchListRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitBranchListRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetKeyWord

`func (o *DescribeGitBranchListRequest) GetKeyWord() string`

GetKeyWord returns the KeyWord field if non-nil, zero value otherwise.

### GetKeyWordOk

`func (o *DescribeGitBranchListRequest) GetKeyWordOk() (*string, bool)`

GetKeyWordOk returns a tuple with the KeyWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWord

`func (o *DescribeGitBranchListRequest) SetKeyWord(v string)`

SetKeyWord sets KeyWord field to given value.

### HasKeyWord

`func (o *DescribeGitBranchListRequest) HasKeyWord() bool`

HasKeyWord returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeGitBranchListRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeGitBranchListRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeGitBranchListRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeGitBranchListRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeGitBranchListRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeGitBranchListRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


