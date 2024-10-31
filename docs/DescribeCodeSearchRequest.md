# DescribeCodeSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | **string** | 分支名称 | 
**DepotPath** | **string** | 仓库路径 | 
**FilePath** | Pointer to **string** | 文件路径（如需查询某个路径下的代码片段，需要输入此参数） | [optional] 
**KeyWord** | **string** | 查询代码片段关键字 | 
**PageNumber** | **int64** | 页码数量 | 
**PageSize** | **int64** | 页码大小 | 

## Methods

### NewDescribeCodeSearchRequest

`func NewDescribeCodeSearchRequest(branchName string, depotPath string, keyWord string, pageNumber int64, pageSize int64, ) *DescribeCodeSearchRequest`

NewDescribeCodeSearchRequest instantiates a new DescribeCodeSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCodeSearchRequestWithDefaults

`func NewDescribeCodeSearchRequestWithDefaults() *DescribeCodeSearchRequest`

NewDescribeCodeSearchRequestWithDefaults instantiates a new DescribeCodeSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchName

`func (o *DescribeCodeSearchRequest) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *DescribeCodeSearchRequest) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *DescribeCodeSearchRequest) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.


### GetDepotPath

`func (o *DescribeCodeSearchRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeCodeSearchRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeCodeSearchRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetFilePath

`func (o *DescribeCodeSearchRequest) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *DescribeCodeSearchRequest) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *DescribeCodeSearchRequest) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *DescribeCodeSearchRequest) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetKeyWord

`func (o *DescribeCodeSearchRequest) GetKeyWord() string`

GetKeyWord returns the KeyWord field if non-nil, zero value otherwise.

### GetKeyWordOk

`func (o *DescribeCodeSearchRequest) GetKeyWordOk() (*string, bool)`

GetKeyWordOk returns a tuple with the KeyWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWord

`func (o *DescribeCodeSearchRequest) SetKeyWord(v string)`

SetKeyWord sets KeyWord field to given value.


### GetPageNumber

`func (o *DescribeCodeSearchRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeCodeSearchRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeCodeSearchRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeCodeSearchRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeCodeSearchRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeCodeSearchRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


