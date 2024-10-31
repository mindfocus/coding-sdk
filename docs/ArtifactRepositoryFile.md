# ArtifactRepositoryFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactType** | Pointer to **string** | 制品类型 | [optional] [default to ""]
**DownloadUrl** | Pointer to **string** | 符合对应制品标准协议的下载链接（有效期：300 s） | [optional] [default to ""]
**Hash** | Pointer to **string** | 制品版本 Hash | [optional] [default to ""]
**Host** | Pointer to **string** | 制品仓库 Host | [optional] [default to ""]
**PackageName** | Pointer to **string** | 制品包名称 | [optional] [default to ""]
**Path** | Pointer to **string** | 相对于仓库级别的文件路径 | [optional] [default to ""]
**Project** | Pointer to **string** | 项目名称 | [optional] [default to ""]
**Repository** | Pointer to **string** | 制品仓库名称 | [optional] [default to ""]
**VersionName** | Pointer to **string** | 制品版本 | [optional] [default to ""]

## Methods

### NewArtifactRepositoryFile

`func NewArtifactRepositoryFile() *ArtifactRepositoryFile`

NewArtifactRepositoryFile instantiates a new ArtifactRepositoryFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactRepositoryFileWithDefaults

`func NewArtifactRepositoryFileWithDefaults() *ArtifactRepositoryFile`

NewArtifactRepositoryFileWithDefaults instantiates a new ArtifactRepositoryFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactType

`func (o *ArtifactRepositoryFile) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *ArtifactRepositoryFile) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *ArtifactRepositoryFile) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *ArtifactRepositoryFile) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *ArtifactRepositoryFile) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *ArtifactRepositoryFile) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *ArtifactRepositoryFile) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *ArtifactRepositoryFile) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetHash

`func (o *ArtifactRepositoryFile) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ArtifactRepositoryFile) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ArtifactRepositoryFile) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ArtifactRepositoryFile) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetHost

`func (o *ArtifactRepositoryFile) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ArtifactRepositoryFile) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ArtifactRepositoryFile) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ArtifactRepositoryFile) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPackageName

`func (o *ArtifactRepositoryFile) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *ArtifactRepositoryFile) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *ArtifactRepositoryFile) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *ArtifactRepositoryFile) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetPath

`func (o *ArtifactRepositoryFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ArtifactRepositoryFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ArtifactRepositoryFile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ArtifactRepositoryFile) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProject

`func (o *ArtifactRepositoryFile) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ArtifactRepositoryFile) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ArtifactRepositoryFile) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ArtifactRepositoryFile) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRepository

`func (o *ArtifactRepositoryFile) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ArtifactRepositoryFile) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ArtifactRepositoryFile) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ArtifactRepositoryFile) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetVersionName

`func (o *ArtifactRepositoryFile) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *ArtifactRepositoryFile) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *ArtifactRepositoryFile) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *ArtifactRepositoryFile) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


