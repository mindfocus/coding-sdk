# DescribeGitCommitFilePathList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePaths** | Pointer to [**[]GitCommitFilePath**](GitCommitFilePath.md) | 文件路径列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitCommitFilePathList200ResponseResponse

`func NewDescribeGitCommitFilePathList200ResponseResponse() *DescribeGitCommitFilePathList200ResponseResponse`

NewDescribeGitCommitFilePathList200ResponseResponse instantiates a new DescribeGitCommitFilePathList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitCommitFilePathList200ResponseResponseWithDefaults

`func NewDescribeGitCommitFilePathList200ResponseResponseWithDefaults() *DescribeGitCommitFilePathList200ResponseResponse`

NewDescribeGitCommitFilePathList200ResponseResponseWithDefaults instantiates a new DescribeGitCommitFilePathList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePaths

`func (o *DescribeGitCommitFilePathList200ResponseResponse) GetFilePaths() []GitCommitFilePath`

GetFilePaths returns the FilePaths field if non-nil, zero value otherwise.

### GetFilePathsOk

`func (o *DescribeGitCommitFilePathList200ResponseResponse) GetFilePathsOk() (*[]GitCommitFilePath, bool)`

GetFilePathsOk returns a tuple with the FilePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePaths

`func (o *DescribeGitCommitFilePathList200ResponseResponse) SetFilePaths(v []GitCommitFilePath)`

SetFilePaths sets FilePaths field to given value.

### HasFilePaths

`func (o *DescribeGitCommitFilePathList200ResponseResponse) HasFilePaths() bool`

HasFilePaths returns a boolean if a field has been set.

### SetFilePathsNil

`func (o *DescribeGitCommitFilePathList200ResponseResponse) SetFilePathsNil(b bool)`

 SetFilePathsNil sets the value for FilePaths to be an explicit nil

### UnsetFilePaths
`func (o *DescribeGitCommitFilePathList200ResponseResponse) UnsetFilePaths()`

UnsetFilePaths ensures that no value is present for FilePaths, not even an explicit nil
### GetRequestId

`func (o *DescribeGitCommitFilePathList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitCommitFilePathList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitCommitFilePathList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitCommitFilePathList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


