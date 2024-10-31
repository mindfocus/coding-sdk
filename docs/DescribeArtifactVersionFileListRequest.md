# DescribeArtifactVersionFileListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Maven** | Pointer to [**Maven**](Maven.md) |  | [optional] 
**Package** | **string** | 包名 | 
**PackageVersion** | **string** | 版本号 | 
**ProjectId** | **int32** | 项目 ID | 
**Repository** | **string** | 仓库名 | 

## Methods

### NewDescribeArtifactVersionFileListRequest

`func NewDescribeArtifactVersionFileListRequest(package_ string, packageVersion string, projectId int32, repository string, ) *DescribeArtifactVersionFileListRequest`

NewDescribeArtifactVersionFileListRequest instantiates a new DescribeArtifactVersionFileListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeArtifactVersionFileListRequestWithDefaults

`func NewDescribeArtifactVersionFileListRequestWithDefaults() *DescribeArtifactVersionFileListRequest`

NewDescribeArtifactVersionFileListRequestWithDefaults instantiates a new DescribeArtifactVersionFileListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaven

`func (o *DescribeArtifactVersionFileListRequest) GetMaven() Maven`

GetMaven returns the Maven field if non-nil, zero value otherwise.

### GetMavenOk

`func (o *DescribeArtifactVersionFileListRequest) GetMavenOk() (*Maven, bool)`

GetMavenOk returns a tuple with the Maven field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaven

`func (o *DescribeArtifactVersionFileListRequest) SetMaven(v Maven)`

SetMaven sets Maven field to given value.

### HasMaven

`func (o *DescribeArtifactVersionFileListRequest) HasMaven() bool`

HasMaven returns a boolean if a field has been set.

### GetPackage

`func (o *DescribeArtifactVersionFileListRequest) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *DescribeArtifactVersionFileListRequest) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *DescribeArtifactVersionFileListRequest) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetPackageVersion

`func (o *DescribeArtifactVersionFileListRequest) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *DescribeArtifactVersionFileListRequest) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *DescribeArtifactVersionFileListRequest) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.


### GetProjectId

`func (o *DescribeArtifactVersionFileListRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeArtifactVersionFileListRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeArtifactVersionFileListRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetRepository

`func (o *DescribeArtifactVersionFileListRequest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DescribeArtifactVersionFileListRequest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DescribeArtifactVersionFileListRequest) SetRepository(v string)`

SetRepository sets Repository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


