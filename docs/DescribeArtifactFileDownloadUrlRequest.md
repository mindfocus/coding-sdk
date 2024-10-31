# DescribeArtifactFileDownloadUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | **string** | 文件名称 | 
**Package** | **string** | 包名 | 
**PackageVersion** | **string** | 版本号 | 
**ProjectId** | **int32** | 项目 ID | 
**Repository** | **string** | 仓库名 | 
**Timeout** | Pointer to **int32** | 下载链接超时时间（单位：秒），默认：300 | [optional] 

## Methods

### NewDescribeArtifactFileDownloadUrlRequest

`func NewDescribeArtifactFileDownloadUrlRequest(fileName string, package_ string, packageVersion string, projectId int32, repository string, ) *DescribeArtifactFileDownloadUrlRequest`

NewDescribeArtifactFileDownloadUrlRequest instantiates a new DescribeArtifactFileDownloadUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeArtifactFileDownloadUrlRequestWithDefaults

`func NewDescribeArtifactFileDownloadUrlRequestWithDefaults() *DescribeArtifactFileDownloadUrlRequest`

NewDescribeArtifactFileDownloadUrlRequestWithDefaults instantiates a new DescribeArtifactFileDownloadUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *DescribeArtifactFileDownloadUrlRequest) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DescribeArtifactFileDownloadUrlRequest) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DescribeArtifactFileDownloadUrlRequest) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetPackage

`func (o *DescribeArtifactFileDownloadUrlRequest) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *DescribeArtifactFileDownloadUrlRequest) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *DescribeArtifactFileDownloadUrlRequest) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetPackageVersion

`func (o *DescribeArtifactFileDownloadUrlRequest) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *DescribeArtifactFileDownloadUrlRequest) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *DescribeArtifactFileDownloadUrlRequest) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.


### GetProjectId

`func (o *DescribeArtifactFileDownloadUrlRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeArtifactFileDownloadUrlRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeArtifactFileDownloadUrlRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetRepository

`func (o *DescribeArtifactFileDownloadUrlRequest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DescribeArtifactFileDownloadUrlRequest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DescribeArtifactFileDownloadUrlRequest) SetRepository(v string)`

SetRepository sets Repository field to given value.


### GetTimeout

`func (o *DescribeArtifactFileDownloadUrlRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DescribeArtifactFileDownloadUrlRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DescribeArtifactFileDownloadUrlRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DescribeArtifactFileDownloadUrlRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


