# DescribeIssueAttachmentPreSignedUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | **string** | 文件名 | [default to ""]
**FileSize** | **int64** | 文件大小 | [default to 0]
**ProjectName** | **string** | 项目名 | [default to ""]

## Methods

### NewDescribeIssueAttachmentPreSignedUrlRequest

`func NewDescribeIssueAttachmentPreSignedUrlRequest(fileName string, fileSize int64, projectName string, ) *DescribeIssueAttachmentPreSignedUrlRequest`

NewDescribeIssueAttachmentPreSignedUrlRequest instantiates a new DescribeIssueAttachmentPreSignedUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIssueAttachmentPreSignedUrlRequestWithDefaults

`func NewDescribeIssueAttachmentPreSignedUrlRequestWithDefaults() *DescribeIssueAttachmentPreSignedUrlRequest`

NewDescribeIssueAttachmentPreSignedUrlRequestWithDefaults instantiates a new DescribeIssueAttachmentPreSignedUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *DescribeIssueAttachmentPreSignedUrlRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DescribeIssueAttachmentPreSignedUrlRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DescribeIssueAttachmentPreSignedUrlRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFileSize

`func (o *DescribeIssueAttachmentPreSignedUrlRequest) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *DescribeIssueAttachmentPreSignedUrlRequest) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *DescribeIssueAttachmentPreSignedUrlRequest) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.


### GetProjectName

`func (o *DescribeIssueAttachmentPreSignedUrlRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeIssueAttachmentPreSignedUrlRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeIssueAttachmentPreSignedUrlRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


