# DescribeArtifactChecksumsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Package** | **string** | 包名 | 
**PackageVersion** | **string** | 版本号 | 
**ProjectId** | **int32** | 项目 ID | 
**Repository** | **string** | 仓库名 | 

## Methods

### NewDescribeArtifactChecksumsRequest

`func NewDescribeArtifactChecksumsRequest(package_ string, packageVersion string, projectId int32, repository string, ) *DescribeArtifactChecksumsRequest`

NewDescribeArtifactChecksumsRequest instantiates a new DescribeArtifactChecksumsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeArtifactChecksumsRequestWithDefaults

`func NewDescribeArtifactChecksumsRequestWithDefaults() *DescribeArtifactChecksumsRequest`

NewDescribeArtifactChecksumsRequestWithDefaults instantiates a new DescribeArtifactChecksumsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *DescribeArtifactChecksumsRequest) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *DescribeArtifactChecksumsRequest) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *DescribeArtifactChecksumsRequest) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetPackageVersion

`func (o *DescribeArtifactChecksumsRequest) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *DescribeArtifactChecksumsRequest) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *DescribeArtifactChecksumsRequest) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.


### GetProjectId

`func (o *DescribeArtifactChecksumsRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeArtifactChecksumsRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeArtifactChecksumsRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetRepository

`func (o *DescribeArtifactChecksumsRequest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DescribeArtifactChecksumsRequest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DescribeArtifactChecksumsRequest) SetRepository(v string)`

SetRepository sets Repository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


