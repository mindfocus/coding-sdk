# DescribeGitBlameInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitSha** | **string** | 提交sha | 
**DepotPath** | **string** | 仓库路径 | 
**FilePath** | **string** | 文件路径 | 
**LineEnd** | Pointer to **int64** | 结束行 | [optional] 
**LineSnat** | Pointer to **int64** | 开始行 | [optional] 

## Methods

### NewDescribeGitBlameInfoRequest

`func NewDescribeGitBlameInfoRequest(commitSha string, depotPath string, filePath string, ) *DescribeGitBlameInfoRequest`

NewDescribeGitBlameInfoRequest instantiates a new DescribeGitBlameInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitBlameInfoRequestWithDefaults

`func NewDescribeGitBlameInfoRequestWithDefaults() *DescribeGitBlameInfoRequest`

NewDescribeGitBlameInfoRequestWithDefaults instantiates a new DescribeGitBlameInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitSha

`func (o *DescribeGitBlameInfoRequest) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *DescribeGitBlameInfoRequest) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *DescribeGitBlameInfoRequest) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### GetDepotPath

`func (o *DescribeGitBlameInfoRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitBlameInfoRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitBlameInfoRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetFilePath

`func (o *DescribeGitBlameInfoRequest) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *DescribeGitBlameInfoRequest) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *DescribeGitBlameInfoRequest) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.


### GetLineEnd

`func (o *DescribeGitBlameInfoRequest) GetLineEnd() int64`

GetLineEnd returns the LineEnd field if non-nil, zero value otherwise.

### GetLineEndOk

`func (o *DescribeGitBlameInfoRequest) GetLineEndOk() (*int64, bool)`

GetLineEndOk returns a tuple with the LineEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineEnd

`func (o *DescribeGitBlameInfoRequest) SetLineEnd(v int64)`

SetLineEnd sets LineEnd field to given value.

### HasLineEnd

`func (o *DescribeGitBlameInfoRequest) HasLineEnd() bool`

HasLineEnd returns a boolean if a field has been set.

### GetLineSnat

`func (o *DescribeGitBlameInfoRequest) GetLineSnat() int64`

GetLineSnat returns the LineSnat field if non-nil, zero value otherwise.

### GetLineSnatOk

`func (o *DescribeGitBlameInfoRequest) GetLineSnatOk() (*int64, bool)`

GetLineSnatOk returns a tuple with the LineSnat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineSnat

`func (o *DescribeGitBlameInfoRequest) SetLineSnat(v int64)`

SetLineSnat sets LineSnat field to given value.

### HasLineSnat

`func (o *DescribeGitBlameInfoRequest) HasLineSnat() bool`

HasLineSnat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


