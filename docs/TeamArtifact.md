# TeamArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactType** | Pointer to **int64** | 制品类型（1-generic;2-docker;3-maven;4-npm;5-pypi;6-helm;7-composer;8-nuget;9-conan;10-cocoapods;11-rpm） | [optional] 
**CreatedAt** | Pointer to **int64** | 推送时间 | [optional] 
**Description** | Pointer to **string** | 制品描述 | [optional] [default to ""]
**DownloadCount** | Pointer to **string** | 下载次数 | [optional] [default to ""]
**Hash** | Pointer to **string** | 制品hash | [optional] [default to ""]
**Package** | Pointer to **string** | 制品名称 | [optional] [default to ""]
**PackageVersion** | Pointer to **string** | 制品版本 | [optional] [default to ""]
**PkgId** | Pointer to **int64** | 制品包ID | [optional] 
**ProjectId** | Pointer to **int64** | 项目ID | [optional] 
**ProjectName** | Pointer to **bool** | 项目名称 | [optional] [default to false]
**ReleaseStatus** | Pointer to **int64** | 发布状态（1&#x3D;未发布，2&#x3D;已发布） | [optional] 
**RepoId** | Pointer to **int64** | 制品仓库ID | [optional] 
**Repository** | Pointer to **string** | 制品仓库名称 | [optional] [default to ""]
**Size** | Pointer to **float32** | 制品代销 | [optional] 
**VersionId** | Pointer to **int64** | 制品版本ID | [optional] 

## Methods

### NewTeamArtifact

`func NewTeamArtifact() *TeamArtifact`

NewTeamArtifact instantiates a new TeamArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamArtifactWithDefaults

`func NewTeamArtifactWithDefaults() *TeamArtifact`

NewTeamArtifactWithDefaults instantiates a new TeamArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactType

`func (o *TeamArtifact) GetArtifactType() int64`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *TeamArtifact) GetArtifactTypeOk() (*int64, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *TeamArtifact) SetArtifactType(v int64)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *TeamArtifact) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TeamArtifact) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TeamArtifact) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TeamArtifact) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TeamArtifact) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *TeamArtifact) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamArtifact) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamArtifact) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamArtifact) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDownloadCount

`func (o *TeamArtifact) GetDownloadCount() string`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *TeamArtifact) GetDownloadCountOk() (*string, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *TeamArtifact) SetDownloadCount(v string)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *TeamArtifact) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### GetHash

`func (o *TeamArtifact) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *TeamArtifact) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *TeamArtifact) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *TeamArtifact) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetPackage

`func (o *TeamArtifact) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *TeamArtifact) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *TeamArtifact) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *TeamArtifact) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetPackageVersion

`func (o *TeamArtifact) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *TeamArtifact) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *TeamArtifact) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *TeamArtifact) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetPkgId

`func (o *TeamArtifact) GetPkgId() int64`

GetPkgId returns the PkgId field if non-nil, zero value otherwise.

### GetPkgIdOk

`func (o *TeamArtifact) GetPkgIdOk() (*int64, bool)`

GetPkgIdOk returns a tuple with the PkgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgId

`func (o *TeamArtifact) SetPkgId(v int64)`

SetPkgId sets PkgId field to given value.

### HasPkgId

`func (o *TeamArtifact) HasPkgId() bool`

HasPkgId returns a boolean if a field has been set.

### GetProjectId

`func (o *TeamArtifact) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TeamArtifact) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TeamArtifact) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *TeamArtifact) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *TeamArtifact) GetProjectName() bool`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *TeamArtifact) GetProjectNameOk() (*bool, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *TeamArtifact) SetProjectName(v bool)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *TeamArtifact) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetReleaseStatus

`func (o *TeamArtifact) GetReleaseStatus() int64`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *TeamArtifact) GetReleaseStatusOk() (*int64, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *TeamArtifact) SetReleaseStatus(v int64)`

SetReleaseStatus sets ReleaseStatus field to given value.

### HasReleaseStatus

`func (o *TeamArtifact) HasReleaseStatus() bool`

HasReleaseStatus returns a boolean if a field has been set.

### GetRepoId

`func (o *TeamArtifact) GetRepoId() int64`

GetRepoId returns the RepoId field if non-nil, zero value otherwise.

### GetRepoIdOk

`func (o *TeamArtifact) GetRepoIdOk() (*int64, bool)`

GetRepoIdOk returns a tuple with the RepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoId

`func (o *TeamArtifact) SetRepoId(v int64)`

SetRepoId sets RepoId field to given value.

### HasRepoId

`func (o *TeamArtifact) HasRepoId() bool`

HasRepoId returns a boolean if a field has been set.

### GetRepository

`func (o *TeamArtifact) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *TeamArtifact) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *TeamArtifact) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *TeamArtifact) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetSize

`func (o *TeamArtifact) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TeamArtifact) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TeamArtifact) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *TeamArtifact) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetVersionId

`func (o *TeamArtifact) GetVersionId() int64`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *TeamArtifact) GetVersionIdOk() (*int64, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *TeamArtifact) SetVersionId(v int64)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *TeamArtifact) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


