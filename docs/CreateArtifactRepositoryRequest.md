# CreateArtifactRepositoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevel** | Pointer to **int32** | 仓库权限范围:1-项目内;2-团队内;3-公开，默认：1-项目内 | [optional] 
**AllowProxy** | Pointer to **bool** | 是否开启代理，仅支持当 Type 为 3-maven;4-npm; 5-PyPI;7-composer;10-cocoapods 时可为 true，默认：false | [optional] 
**Description** | Pointer to **string** | 仓库描述信息 | [optional] 
**ProjectId** | **int32** | 项目 ID | 
**RepositoryName** | **string** | 仓库名称 | 
**Type** | **int32** | 仓库类型:1-generic;2-docker;3-maven;4-npm;5-pypi;6-helm;7-composer;8-nuget;9-conan;10-cocoapods;11-rpm | 

## Methods

### NewCreateArtifactRepositoryRequest

`func NewCreateArtifactRepositoryRequest(projectId int32, repositoryName string, type_ int32, ) *CreateArtifactRepositoryRequest`

NewCreateArtifactRepositoryRequest instantiates a new CreateArtifactRepositoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateArtifactRepositoryRequestWithDefaults

`func NewCreateArtifactRepositoryRequestWithDefaults() *CreateArtifactRepositoryRequest`

NewCreateArtifactRepositoryRequestWithDefaults instantiates a new CreateArtifactRepositoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevel

`func (o *CreateArtifactRepositoryRequest) GetAccessLevel() int32`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *CreateArtifactRepositoryRequest) GetAccessLevelOk() (*int32, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *CreateArtifactRepositoryRequest) SetAccessLevel(v int32)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *CreateArtifactRepositoryRequest) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetAllowProxy

`func (o *CreateArtifactRepositoryRequest) GetAllowProxy() bool`

GetAllowProxy returns the AllowProxy field if non-nil, zero value otherwise.

### GetAllowProxyOk

`func (o *CreateArtifactRepositoryRequest) GetAllowProxyOk() (*bool, bool)`

GetAllowProxyOk returns a tuple with the AllowProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProxy

`func (o *CreateArtifactRepositoryRequest) SetAllowProxy(v bool)`

SetAllowProxy sets AllowProxy field to given value.

### HasAllowProxy

`func (o *CreateArtifactRepositoryRequest) HasAllowProxy() bool`

HasAllowProxy returns a boolean if a field has been set.

### GetDescription

`func (o *CreateArtifactRepositoryRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateArtifactRepositoryRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateArtifactRepositoryRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateArtifactRepositoryRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateArtifactRepositoryRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateArtifactRepositoryRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateArtifactRepositoryRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetRepositoryName

`func (o *CreateArtifactRepositoryRequest) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *CreateArtifactRepositoryRequest) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *CreateArtifactRepositoryRequest) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.


### GetType

`func (o *CreateArtifactRepositoryRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateArtifactRepositoryRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateArtifactRepositoryRequest) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


