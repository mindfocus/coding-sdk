# ArtifactRepositoryBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevel** | Pointer to **int64** | 仓库访问权限：1-项目内;2-团队内;3-公开 | [optional] 
**CreatedAt** | Pointer to **int64** | 创建时间 | [optional] 
**Description** | Pointer to **NullableString** | 仓库描述 | [optional] [default to ""]
**Id** | Pointer to **int32** | 仓库 ID | [optional] 
**Name** | Pointer to **string** | 仓库名称 | [optional] [default to ""]
**ProjectId** | Pointer to **int32** | 项目 ID | [optional] 
**ReleaseStrategy** | Pointer to **int64** | 版本发布策略：1-允许覆盖发布;2-不允许覆盖发布;3-快照策略 | [optional] 
**TeamId** | Pointer to **int32** | 团队 ID | [optional] 
**Type** | Pointer to **int64** | 仓库类型：1-generic;2-docker;3-maven;4-npm;5-pypi;6-helm;7-composer;8-nuget;9-conan;10-cocoapods;11-rpm | [optional] 

## Methods

### NewArtifactRepositoryBean

`func NewArtifactRepositoryBean() *ArtifactRepositoryBean`

NewArtifactRepositoryBean instantiates a new ArtifactRepositoryBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactRepositoryBeanWithDefaults

`func NewArtifactRepositoryBeanWithDefaults() *ArtifactRepositoryBean`

NewArtifactRepositoryBeanWithDefaults instantiates a new ArtifactRepositoryBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevel

`func (o *ArtifactRepositoryBean) GetAccessLevel() int64`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *ArtifactRepositoryBean) GetAccessLevelOk() (*int64, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *ArtifactRepositoryBean) SetAccessLevel(v int64)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *ArtifactRepositoryBean) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ArtifactRepositoryBean) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArtifactRepositoryBean) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArtifactRepositoryBean) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArtifactRepositoryBean) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ArtifactRepositoryBean) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArtifactRepositoryBean) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArtifactRepositoryBean) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArtifactRepositoryBean) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ArtifactRepositoryBean) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ArtifactRepositoryBean) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *ArtifactRepositoryBean) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactRepositoryBean) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactRepositoryBean) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ArtifactRepositoryBean) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ArtifactRepositoryBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactRepositoryBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactRepositoryBean) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArtifactRepositoryBean) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *ArtifactRepositoryBean) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ArtifactRepositoryBean) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ArtifactRepositoryBean) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ArtifactRepositoryBean) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReleaseStrategy

`func (o *ArtifactRepositoryBean) GetReleaseStrategy() int64`

GetReleaseStrategy returns the ReleaseStrategy field if non-nil, zero value otherwise.

### GetReleaseStrategyOk

`func (o *ArtifactRepositoryBean) GetReleaseStrategyOk() (*int64, bool)`

GetReleaseStrategyOk returns a tuple with the ReleaseStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStrategy

`func (o *ArtifactRepositoryBean) SetReleaseStrategy(v int64)`

SetReleaseStrategy sets ReleaseStrategy field to given value.

### HasReleaseStrategy

`func (o *ArtifactRepositoryBean) HasReleaseStrategy() bool`

HasReleaseStrategy returns a boolean if a field has been set.

### GetTeamId

`func (o *ArtifactRepositoryBean) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *ArtifactRepositoryBean) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *ArtifactRepositoryBean) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *ArtifactRepositoryBean) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetType

`func (o *ArtifactRepositoryBean) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArtifactRepositoryBean) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArtifactRepositoryBean) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *ArtifactRepositoryBean) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


