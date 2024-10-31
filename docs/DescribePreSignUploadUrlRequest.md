# DescribePreSignUploadUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** | 内容类型，和web端的文件上传时content-type是一样的 | [default to ""]
**FileName** | **string** | 文件名 | [default to ""]
**FolderId** | **int64** | 若上传项目协同附件场景时，不需要配置或默认为0即可。若用于上传到文件网盘，可以设置文件夹ID, 用于文件存放位置，ID值通过open api【创建文件夹】获取 | [default to 0]
**FolderType** | **int64** | 文件夹类型，0: 常规文件夹, 1:隐藏文件夹。 如果是用于项目协同上传附件的场景，配置为1；如果是上传到文件网盘则配置为0。 | [default to 0]
**ProjectName** | **string** | 项目名称 | [default to ""]

## Methods

### NewDescribePreSignUploadUrlRequest

`func NewDescribePreSignUploadUrlRequest(contentType string, fileName string, folderId int64, folderType int64, projectName string, ) *DescribePreSignUploadUrlRequest`

NewDescribePreSignUploadUrlRequest instantiates a new DescribePreSignUploadUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribePreSignUploadUrlRequestWithDefaults

`func NewDescribePreSignUploadUrlRequestWithDefaults() *DescribePreSignUploadUrlRequest`

NewDescribePreSignUploadUrlRequestWithDefaults instantiates a new DescribePreSignUploadUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *DescribePreSignUploadUrlRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *DescribePreSignUploadUrlRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *DescribePreSignUploadUrlRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetFileName

`func (o *DescribePreSignUploadUrlRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DescribePreSignUploadUrlRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DescribePreSignUploadUrlRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFolderId

`func (o *DescribePreSignUploadUrlRequest) GetFolderId() int64`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *DescribePreSignUploadUrlRequest) GetFolderIdOk() (*int64, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *DescribePreSignUploadUrlRequest) SetFolderId(v int64)`

SetFolderId sets FolderId field to given value.


### GetFolderType

`func (o *DescribePreSignUploadUrlRequest) GetFolderType() int64`

GetFolderType returns the FolderType field if non-nil, zero value otherwise.

### GetFolderTypeOk

`func (o *DescribePreSignUploadUrlRequest) GetFolderTypeOk() (*int64, bool)`

GetFolderTypeOk returns a tuple with the FolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderType

`func (o *DescribePreSignUploadUrlRequest) SetFolderType(v int64)`

SetFolderType sets FolderType field to given value.


### GetProjectName

`func (o *DescribePreSignUploadUrlRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribePreSignUploadUrlRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribePreSignUploadUrlRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


