# DescribeGitCommitCommentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径 | 
**PageNumber** | **int64** | 页码数量 | 
**PageSize** | **int64** | 页码大小 | 
**Sha** | **string** | 提交Sha | 

## Methods

### NewDescribeGitCommitCommentsRequest

`func NewDescribeGitCommitCommentsRequest(depotPath string, pageNumber int64, pageSize int64, sha string, ) *DescribeGitCommitCommentsRequest`

NewDescribeGitCommitCommentsRequest instantiates a new DescribeGitCommitCommentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitCommitCommentsRequestWithDefaults

`func NewDescribeGitCommitCommentsRequestWithDefaults() *DescribeGitCommitCommentsRequest`

NewDescribeGitCommitCommentsRequestWithDefaults instantiates a new DescribeGitCommitCommentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *DescribeGitCommitCommentsRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitCommitCommentsRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitCommitCommentsRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetPageNumber

`func (o *DescribeGitCommitCommentsRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeGitCommitCommentsRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeGitCommitCommentsRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeGitCommitCommentsRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeGitCommitCommentsRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeGitCommitCommentsRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetSha

`func (o *DescribeGitCommitCommentsRequest) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *DescribeGitCommitCommentsRequest) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *DescribeGitCommitCommentsRequest) SetSha(v string)`

SetSha sets Sha field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


