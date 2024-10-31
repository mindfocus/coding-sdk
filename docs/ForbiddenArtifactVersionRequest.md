# ForbiddenArtifactVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForbiddenAction** | **string** | FORBIDDEN 禁止下载，UNFORBIDDEN 解除禁止下载 | 
**ForbiddenNote** | Pointer to **string** | 禁止下载说明 | [optional] 
**Package** | **string** | 包名 | 
**PackageVersion** | **string** | 版本号 | 
**ProjectId** | **int32** | 项目 ID | 
**Repository** | **string** | 仓库名 | 

## Methods

### NewForbiddenArtifactVersionRequest

`func NewForbiddenArtifactVersionRequest(forbiddenAction string, package_ string, packageVersion string, projectId int32, repository string, ) *ForbiddenArtifactVersionRequest`

NewForbiddenArtifactVersionRequest instantiates a new ForbiddenArtifactVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForbiddenArtifactVersionRequestWithDefaults

`func NewForbiddenArtifactVersionRequestWithDefaults() *ForbiddenArtifactVersionRequest`

NewForbiddenArtifactVersionRequestWithDefaults instantiates a new ForbiddenArtifactVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForbiddenAction

`func (o *ForbiddenArtifactVersionRequest) GetForbiddenAction() string`

GetForbiddenAction returns the ForbiddenAction field if non-nil, zero value otherwise.

### GetForbiddenActionOk

`func (o *ForbiddenArtifactVersionRequest) GetForbiddenActionOk() (*string, bool)`

GetForbiddenActionOk returns a tuple with the ForbiddenAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenAction

`func (o *ForbiddenArtifactVersionRequest) SetForbiddenAction(v string)`

SetForbiddenAction sets ForbiddenAction field to given value.


### GetForbiddenNote

`func (o *ForbiddenArtifactVersionRequest) GetForbiddenNote() string`

GetForbiddenNote returns the ForbiddenNote field if non-nil, zero value otherwise.

### GetForbiddenNoteOk

`func (o *ForbiddenArtifactVersionRequest) GetForbiddenNoteOk() (*string, bool)`

GetForbiddenNoteOk returns a tuple with the ForbiddenNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenNote

`func (o *ForbiddenArtifactVersionRequest) SetForbiddenNote(v string)`

SetForbiddenNote sets ForbiddenNote field to given value.

### HasForbiddenNote

`func (o *ForbiddenArtifactVersionRequest) HasForbiddenNote() bool`

HasForbiddenNote returns a boolean if a field has been set.

### GetPackage

`func (o *ForbiddenArtifactVersionRequest) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *ForbiddenArtifactVersionRequest) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *ForbiddenArtifactVersionRequest) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetPackageVersion

`func (o *ForbiddenArtifactVersionRequest) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *ForbiddenArtifactVersionRequest) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *ForbiddenArtifactVersionRequest) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.


### GetProjectId

`func (o *ForbiddenArtifactVersionRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ForbiddenArtifactVersionRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ForbiddenArtifactVersionRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetRepository

`func (o *ForbiddenArtifactVersionRequest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ForbiddenArtifactVersionRequest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ForbiddenArtifactVersionRequest) SetRepository(v string)`

SetRepository sets Repository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


