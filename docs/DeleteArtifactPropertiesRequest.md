# DeleteArtifactPropertiesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Package** | **string** | 包名 | 
**PackageVersion** | **string** | 版本号 | 
**ProjectId** | **int32** | 项目 ID | 
**PropertyNameSet** | **[]string** | 属性名称列表（ 以 ‘coding.’ 作为属性名称开头的属性，将不可删除） | 
**Repository** | **string** | 仓库名 | 

## Methods

### NewDeleteArtifactPropertiesRequest

`func NewDeleteArtifactPropertiesRequest(package_ string, packageVersion string, projectId int32, propertyNameSet []string, repository string, ) *DeleteArtifactPropertiesRequest`

NewDeleteArtifactPropertiesRequest instantiates a new DeleteArtifactPropertiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteArtifactPropertiesRequestWithDefaults

`func NewDeleteArtifactPropertiesRequestWithDefaults() *DeleteArtifactPropertiesRequest`

NewDeleteArtifactPropertiesRequestWithDefaults instantiates a new DeleteArtifactPropertiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *DeleteArtifactPropertiesRequest) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *DeleteArtifactPropertiesRequest) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *DeleteArtifactPropertiesRequest) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetPackageVersion

`func (o *DeleteArtifactPropertiesRequest) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *DeleteArtifactPropertiesRequest) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *DeleteArtifactPropertiesRequest) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.


### GetProjectId

`func (o *DeleteArtifactPropertiesRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DeleteArtifactPropertiesRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DeleteArtifactPropertiesRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetPropertyNameSet

`func (o *DeleteArtifactPropertiesRequest) GetPropertyNameSet() []string`

GetPropertyNameSet returns the PropertyNameSet field if non-nil, zero value otherwise.

### GetPropertyNameSetOk

`func (o *DeleteArtifactPropertiesRequest) GetPropertyNameSetOk() (*[]string, bool)`

GetPropertyNameSetOk returns a tuple with the PropertyNameSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyNameSet

`func (o *DeleteArtifactPropertiesRequest) SetPropertyNameSet(v []string)`

SetPropertyNameSet sets PropertyNameSet field to given value.


### GetRepository

`func (o *DeleteArtifactPropertiesRequest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DeleteArtifactPropertiesRequest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DeleteArtifactPropertiesRequest) SetRepository(v string)`

SetRepository sets Repository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


