# ArtifactPackageBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | 创建时间 | [optional] 
**Description** | Pointer to **NullableString** | 制品包描述 | [optional] [default to ""]
**Id** | Pointer to **int32** | 制品包 ID | [optional] 
**LatestVersionId** | Pointer to **int32** | 最新推送版本的版本号 ID | [optional] 
**LatestVersionName** | Pointer to **string** | 最新推送版本的版本号 | [optional] [default to ""]
**LatestVersionReleaseStatus** | Pointer to **int64** | 最新推送版本的版本发布状态：1-未发布;2-已发布 | [optional] 
**Name** | Pointer to **string** | 制品包名称 | [optional] [default to ""]
**ReleaseStrategy** | Pointer to **int64** | 版本发布策略：1-允许覆盖发布;2-不允许覆盖发布;3-快照策略;4-继承于仓库的策略 | [optional] 
**RepoId** | Pointer to **int32** | 制品仓库 ID | [optional] 
**VersionCount** | Pointer to **int32** | 包下的版本数量 | [optional] 

## Methods

### NewArtifactPackageBean

`func NewArtifactPackageBean() *ArtifactPackageBean`

NewArtifactPackageBean instantiates a new ArtifactPackageBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactPackageBeanWithDefaults

`func NewArtifactPackageBeanWithDefaults() *ArtifactPackageBean`

NewArtifactPackageBeanWithDefaults instantiates a new ArtifactPackageBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ArtifactPackageBean) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArtifactPackageBean) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArtifactPackageBean) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArtifactPackageBean) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ArtifactPackageBean) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArtifactPackageBean) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArtifactPackageBean) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArtifactPackageBean) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ArtifactPackageBean) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ArtifactPackageBean) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *ArtifactPackageBean) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactPackageBean) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactPackageBean) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ArtifactPackageBean) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatestVersionId

`func (o *ArtifactPackageBean) GetLatestVersionId() int32`

GetLatestVersionId returns the LatestVersionId field if non-nil, zero value otherwise.

### GetLatestVersionIdOk

`func (o *ArtifactPackageBean) GetLatestVersionIdOk() (*int32, bool)`

GetLatestVersionIdOk returns a tuple with the LatestVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersionId

`func (o *ArtifactPackageBean) SetLatestVersionId(v int32)`

SetLatestVersionId sets LatestVersionId field to given value.

### HasLatestVersionId

`func (o *ArtifactPackageBean) HasLatestVersionId() bool`

HasLatestVersionId returns a boolean if a field has been set.

### GetLatestVersionName

`func (o *ArtifactPackageBean) GetLatestVersionName() string`

GetLatestVersionName returns the LatestVersionName field if non-nil, zero value otherwise.

### GetLatestVersionNameOk

`func (o *ArtifactPackageBean) GetLatestVersionNameOk() (*string, bool)`

GetLatestVersionNameOk returns a tuple with the LatestVersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersionName

`func (o *ArtifactPackageBean) SetLatestVersionName(v string)`

SetLatestVersionName sets LatestVersionName field to given value.

### HasLatestVersionName

`func (o *ArtifactPackageBean) HasLatestVersionName() bool`

HasLatestVersionName returns a boolean if a field has been set.

### GetLatestVersionReleaseStatus

`func (o *ArtifactPackageBean) GetLatestVersionReleaseStatus() int64`

GetLatestVersionReleaseStatus returns the LatestVersionReleaseStatus field if non-nil, zero value otherwise.

### GetLatestVersionReleaseStatusOk

`func (o *ArtifactPackageBean) GetLatestVersionReleaseStatusOk() (*int64, bool)`

GetLatestVersionReleaseStatusOk returns a tuple with the LatestVersionReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersionReleaseStatus

`func (o *ArtifactPackageBean) SetLatestVersionReleaseStatus(v int64)`

SetLatestVersionReleaseStatus sets LatestVersionReleaseStatus field to given value.

### HasLatestVersionReleaseStatus

`func (o *ArtifactPackageBean) HasLatestVersionReleaseStatus() bool`

HasLatestVersionReleaseStatus returns a boolean if a field has been set.

### GetName

`func (o *ArtifactPackageBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactPackageBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactPackageBean) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArtifactPackageBean) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReleaseStrategy

`func (o *ArtifactPackageBean) GetReleaseStrategy() int64`

GetReleaseStrategy returns the ReleaseStrategy field if non-nil, zero value otherwise.

### GetReleaseStrategyOk

`func (o *ArtifactPackageBean) GetReleaseStrategyOk() (*int64, bool)`

GetReleaseStrategyOk returns a tuple with the ReleaseStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStrategy

`func (o *ArtifactPackageBean) SetReleaseStrategy(v int64)`

SetReleaseStrategy sets ReleaseStrategy field to given value.

### HasReleaseStrategy

`func (o *ArtifactPackageBean) HasReleaseStrategy() bool`

HasReleaseStrategy returns a boolean if a field has been set.

### GetRepoId

`func (o *ArtifactPackageBean) GetRepoId() int32`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *ArtifactPackageBean) GetRepoIdOk() (*int32, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *ArtifactPackageBean) SetRepoId(v int32)`

SetRepoId sets RepoId field to given value.

### HasRepoId

`func (o *ArtifactPackageBean) HasRepoId() bool`

HasRepoId returns a boolean if a field has been set.

### GetVersionCount

`func (o *ArtifactPackageBean) GetVersionCount() int32`

GetVersionCount returns the VersionCount field if non-nil, zero value otherwise.

### GetVersionCountOk

`func (o *ArtifactPackageBean) GetVersionCountOk() (*int32, bool)`

GetVersionCountOk returns a tuple with the VersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCount

`func (o *ArtifactPackageBean) SetVersionCount(v int32)`

SetVersionCount sets VersionCount field to given value.

### HasVersionCount

`func (o *ArtifactPackageBean) HasVersionCount() bool`

HasVersionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


