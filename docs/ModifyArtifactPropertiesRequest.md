# ModifyArtifactPropertiesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Package** | **string** | 包名 | 
**PackageVersion** | **string** | 版本号 | 
**ProjectId** | **int32** | 项目 ID | 
**PropertySet** | [**[]ArtifactPropertyBean**](ArtifactPropertyBean.md) | 属性列表 | 
**Repository** | **string** | 仓库名 | 

## Methods

### NewModifyArtifactPropertiesRequest

`func NewModifyArtifactPropertiesRequest(package_ string, packageVersion string, projectId int32, propertySet []ArtifactPropertyBean, repository string, ) *ModifyArtifactPropertiesRequest`

NewModifyArtifactPropertiesRequest instantiates a new ModifyArtifactPropertiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyArtifactPropertiesRequestWithDefaults

`func NewModifyArtifactPropertiesRequestWithDefaults() *ModifyArtifactPropertiesRequest`

NewModifyArtifactPropertiesRequestWithDefaults instantiates a new ModifyArtifactPropertiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *ModifyArtifactPropertiesRequest) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *ModifyArtifactPropertiesRequest) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *ModifyArtifactPropertiesRequest) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetPackageVersion

`func (o *ModifyArtifactPropertiesRequest) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *ModifyArtifactPropertiesRequest) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *ModifyArtifactPropertiesRequest) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.


### GetProjectId

`func (o *ModifyArtifactPropertiesRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModifyArtifactPropertiesRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModifyArtifactPropertiesRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetPropertySet

`func (o *ModifyArtifactPropertiesRequest) GetPropertySet() []ArtifactPropertyBean`

GetPropertySet returns the PropertySet field if non-nil, zero value otherwise.

### GetPropertySetOk

`func (o *ModifyArtifactPropertiesRequest) GetPropertySetOk() (*[]ArtifactPropertyBean, bool)`

GetPropertySetOk returns a tuple with the PropertySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySet

`func (o *ModifyArtifactPropertiesRequest) SetPropertySet(v []ArtifactPropertyBean)`

SetPropertySet sets PropertySet field to given value.


### GetRepository

`func (o *ModifyArtifactPropertiesRequest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ModifyArtifactPropertiesRequest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ModifyArtifactPropertiesRequest) SetRepository(v string)`

SetRepository sets Repository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


