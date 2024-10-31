# ArtifactsOpenApiRemoteTeamData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncMessage** | **string** | 下发信息 | [default to ""]
**RemoteTeamUrl** | **string** | 团队地址 | [default to ""]
**UserName** | **string** | 用户名 | [default to ""]
**SyncStatus** | **int64** | 下发状态 | [default to 0]
**Id** | **int64** | ID | [default to 0]
**RemoteRepos** | [**[]ArtifactsOpenApiRemoteRepoData**](ArtifactsOpenApiRemoteRepoData.md) | 远程仓库列表 | 
**RemoteTeamName** | **string** | 远程团队名 | [default to ""]
**Successful** | **bool** | 是否成功 | [default to false]
**Deleting** | **bool** | 是否删除 | [default to false]

## Methods

### NewArtifactsOpenApiRemoteTeamData

`func NewArtifactsOpenApiRemoteTeamData(syncMessage string, remoteTeamUrl string, userName string, syncStatus int64, id int64, remoteRepos []ArtifactsOpenApiRemoteRepoData, remoteTeamName string, successful bool, deleting bool, ) *ArtifactsOpenApiRemoteTeamData`

NewArtifactsOpenApiRemoteTeamData instantiates a new ArtifactsOpenApiRemoteTeamData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactsOpenApiRemoteTeamDataWithDefaults

`func NewArtifactsOpenApiRemoteTeamDataWithDefaults() *ArtifactsOpenApiRemoteTeamData`

NewArtifactsOpenApiRemoteTeamDataWithDefaults instantiates a new ArtifactsOpenApiRemoteTeamData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyncMessage

`func (o *ArtifactsOpenApiRemoteTeamData) GetSyncMessage() string`

GetSyncMessage returns the SyncMessage field if non-nil, zero value otherwise.

### GetSyncMessageOk

`func (o *ArtifactsOpenApiRemoteTeamData) GetSyncMessageOk() (*string, bool)`

GetSyncMessageOk returns a tuple with the SyncMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncMessage

`func (o *ArtifactsOpenApiRemoteTeamData) SetSyncMessage(v string)`

SetSyncMessage sets SyncMessage field to given value.


### GetRemoteTeamUrl

`func (o *ArtifactsOpenApiRemoteTeamData) GetRemoteTeamUrl() string`

GetRemoteTeamUrl returns the RemoteTeamUrl field if non-nil, zero value otherwise.

### GetRemoteTeamUrlOk

`func (o *ArtifactsOpenApiRemoteTeamData) GetRemoteTeamUrlOk() (*string, bool)`

GetRemoteTeamUrlOk returns a tuple with the RemoteTeamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTeamUrl

`func (o *ArtifactsOpenApiRemoteTeamData) SetRemoteTeamUrl(v string)`

SetRemoteTeamUrl sets RemoteTeamUrl field to given value.


### GetUserName

`func (o *ArtifactsOpenApiRemoteTeamData) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ArtifactsOpenApiRemoteTeamData) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ArtifactsOpenApiRemoteTeamData) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetSyncStatus

`func (o *ArtifactsOpenApiRemoteTeamData) GetSyncStatus() int64`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *ArtifactsOpenApiRemoteTeamData) GetSyncStatusOk() (*int64, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *ArtifactsOpenApiRemoteTeamData) SetSyncStatus(v int64)`

SetSyncStatus sets SyncStatus field to given value.


### GetId

`func (o *ArtifactsOpenApiRemoteTeamData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactsOpenApiRemoteTeamData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactsOpenApiRemoteTeamData) SetId(v int64)`

SetId sets Id field to given value.


### GetRemoteRepos

`func (o *ArtifactsOpenApiRemoteTeamData) GetRemoteRepos() []ArtifactsOpenApiRemoteRepoData`

GetRemoteRepos returns the RemoteRepos field if non-nil, zero value otherwise.

### GetRemoteReposOk

`func (o *ArtifactsOpenApiRemoteTeamData) GetRemoteReposOk() (*[]ArtifactsOpenApiRemoteRepoData, bool)`

GetRemoteReposOk returns a tuple with the RemoteRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteRepos

`func (o *ArtifactsOpenApiRemoteTeamData) SetRemoteRepos(v []ArtifactsOpenApiRemoteRepoData)`

SetRemoteRepos sets RemoteRepos field to given value.


### GetRemoteTeamName

`func (o *ArtifactsOpenApiRemoteTeamData) GetRemoteTeamName() string`

GetRemoteTeamName returns the RemoteTeamName field if non-nil, zero value otherwise.

### GetRemoteTeamNameOk

`func (o *ArtifactsOpenApiRemoteTeamData) GetRemoteTeamNameOk() (*string, bool)`

GetRemoteTeamNameOk returns a tuple with the RemoteTeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTeamName

`func (o *ArtifactsOpenApiRemoteTeamData) SetRemoteTeamName(v string)`

SetRemoteTeamName sets RemoteTeamName field to given value.


### GetSuccessful

`func (o *ArtifactsOpenApiRemoteTeamData) GetSuccessful() bool`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *ArtifactsOpenApiRemoteTeamData) GetSuccessfulOk() (*bool, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *ArtifactsOpenApiRemoteTeamData) SetSuccessful(v bool)`

SetSuccessful sets Successful field to given value.


### GetDeleting

`func (o *ArtifactsOpenApiRemoteTeamData) GetDeleting() bool`

GetDeleting returns the Deleting field if non-nil, zero value otherwise.

### GetDeletingOk

`func (o *ArtifactsOpenApiRemoteTeamData) GetDeletingOk() (*bool, bool)`

GetDeletingOk returns a tuple with the Deleting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleting

`func (o *ArtifactsOpenApiRemoteTeamData) SetDeleting(v bool)`

SetDeleting sets Deleting field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


