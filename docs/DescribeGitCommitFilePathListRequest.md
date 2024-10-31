# DescribeGitCommitFilePathListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitSha** | **string** | 提交的ID | 
**DepotPath** | **string** | 仓库路径，也可填写仓库ID的字符串形式 | 

## Methods

### NewDescribeGitCommitFilePathListRequest

`func NewDescribeGitCommitFilePathListRequest(commitSha string, depotPath string, ) *DescribeGitCommitFilePathListRequest`

NewDescribeGitCommitFilePathListRequest instantiates a new DescribeGitCommitFilePathListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitCommitFilePathListRequestWithDefaults

`func NewDescribeGitCommitFilePathListRequestWithDefaults() *DescribeGitCommitFilePathListRequest`

NewDescribeGitCommitFilePathListRequestWithDefaults instantiates a new DescribeGitCommitFilePathListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitSha

`func (o *DescribeGitCommitFilePathListRequest) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *DescribeGitCommitFilePathListRequest) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *DescribeGitCommitFilePathListRequest) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### GetDepotPath

`func (o *DescribeGitCommitFilePathListRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitCommitFilePathListRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitCommitFilePathListRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


