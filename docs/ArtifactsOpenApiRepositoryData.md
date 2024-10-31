# ArtifactsOpenApiRepositoryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int64** | 制品类型(1-generic;2-docker;3-maven;4-npm;5-pypi;6-helm;7-composer;8-nuget;9-conan,10-cocoapods,11-rpm,12-Go) | [default to 0]
**StorageRule** | **int64** | 存储规则 | [default to 0]
**Description** | **string** | 项目描述 | [default to ""]
**CreatorId** | **int64** | 创建人ID | [default to 0]
**ProjectId** | **int64** | 项目ID | [default to 0]
**Id** | **int64** | 制品仓库ID | [default to 0]
**TeamId** | **int64** | 团队ID | [default to 0]
**ReleaseStrategy** | **int64** | 版本发布策略：1-允许覆盖发布;2-不允许覆盖发布;3-快照策略 | [default to 0]
**Deleting** | **bool** | 是否删除 | [default to false]
**Name** | **string** | 制品库名称 | [default to ""]

## Methods

### NewArtifactsOpenApiRepositoryData

`func NewArtifactsOpenApiRepositoryData(type_ int64, storageRule int64, description string, creatorId int64, projectId int64, id int64, teamId int64, releaseStrategy int64, deleting bool, name string, ) *ArtifactsOpenApiRepositoryData`

NewArtifactsOpenApiRepositoryData instantiates a new ArtifactsOpenApiRepositoryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactsOpenApiRepositoryDataWithDefaults

`func NewArtifactsOpenApiRepositoryDataWithDefaults() *ArtifactsOpenApiRepositoryData`

NewArtifactsOpenApiRepositoryDataWithDefaults instantiates a new ArtifactsOpenApiRepositoryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ArtifactsOpenApiRepositoryData) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArtifactsOpenApiRepositoryData) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArtifactsOpenApiRepositoryData) SetType(v int64)`

SetType sets Type field to given value.


### GetStorageRule

`func (o *ArtifactsOpenApiRepositoryData) GetStorageRule() int64`

GetStorageRule returns the StorageRule field if non-nil, zero value otherwise.

### GetStorageRuleOk

`func (o *ArtifactsOpenApiRepositoryData) GetStorageRuleOk() (*int64, bool)`

GetStorageRuleOk returns a tuple with the StorageRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageRule

`func (o *ArtifactsOpenApiRepositoryData) SetStorageRule(v int64)`

SetStorageRule sets StorageRule field to given value.


### GetDescription

`func (o *ArtifactsOpenApiRepositoryData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArtifactsOpenApiRepositoryData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArtifactsOpenApiRepositoryData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatorId

`func (o *ArtifactsOpenApiRepositoryData) GetCreatorId() int64`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ArtifactsOpenApiRepositoryData) GetCreatorIdOk() (*int64, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ArtifactsOpenApiRepositoryData) SetCreatorId(v int64)`

SetCreatorId sets CreatorId field to given value.


### GetProjectId

`func (o *ArtifactsOpenApiRepositoryData) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ArtifactsOpenApiRepositoryData) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ArtifactsOpenApiRepositoryData) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetId

`func (o *ArtifactsOpenApiRepositoryData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactsOpenApiRepositoryData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactsOpenApiRepositoryData) SetId(v int64)`

SetId sets Id field to given value.


### GetTeamId

`func (o *ArtifactsOpenApiRepositoryData) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *ArtifactsOpenApiRepositoryData) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *ArtifactsOpenApiRepositoryData) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.


### GetReleaseStrategy

`func (o *ArtifactsOpenApiRepositoryData) GetReleaseStrategy() int64`

GetReleaseStrategy returns the ReleaseStrategy field if non-nil, zero value otherwise.

### GetReleaseStrategyOk

`func (o *ArtifactsOpenApiRepositoryData) GetReleaseStrategyOk() (*int64, bool)`

GetReleaseStrategyOk returns a tuple with the ReleaseStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStrategy

`func (o *ArtifactsOpenApiRepositoryData) SetReleaseStrategy(v int64)`

SetReleaseStrategy sets ReleaseStrategy field to given value.


### GetDeleting

`func (o *ArtifactsOpenApiRepositoryData) GetDeleting() bool`

GetDeleting returns the Deleting field if non-nil, zero value otherwise.

### GetDeletingOk

`func (o *ArtifactsOpenApiRepositoryData) GetDeletingOk() (*bool, bool)`

GetDeletingOk returns a tuple with the Deleting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleting

`func (o *ArtifactsOpenApiRepositoryData) SetDeleting(v bool)`

SetDeleting sets Deleting field to given value.


### GetName

`func (o *ArtifactsOpenApiRepositoryData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactsOpenApiRepositoryData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactsOpenApiRepositoryData) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


