# DescribeIssueFileUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | **int64** | 文件 ID | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewDescribeIssueFileUrlRequest

`func NewDescribeIssueFileUrlRequest(fileId int64, projectName string, ) *DescribeIssueFileUrlRequest`

NewDescribeIssueFileUrlRequest instantiates a new DescribeIssueFileUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIssueFileUrlRequestWithDefaults

`func NewDescribeIssueFileUrlRequestWithDefaults() *DescribeIssueFileUrlRequest`

NewDescribeIssueFileUrlRequestWithDefaults instantiates a new DescribeIssueFileUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *DescribeIssueFileUrlRequest) GetFileId() int64`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *DescribeIssueFileUrlRequest) GetFileIdOk() (*int64, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *DescribeIssueFileUrlRequest) SetFileId(v int64)`

SetFileId sets FileId field to given value.


### GetProjectName

`func (o *DescribeIssueFileUrlRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeIssueFileUrlRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeIssueFileUrlRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


