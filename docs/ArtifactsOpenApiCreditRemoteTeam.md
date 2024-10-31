# ArtifactsOpenApiCreditRemoteTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteTeamUrl** | **string** | 远程团队地址 | [default to ""]
**UserName** | **string** | 用户名 | [default to ""]
**RemoteRepos** | [**[]ArtifactsOpenApiRemoteRepoData**](ArtifactsOpenApiRemoteRepoData.md) | 远程仓库列表 | 
**Password** | **string** | 个人令牌base64编码 | [default to ""]

## Methods

### NewArtifactsOpenApiCreditRemoteTeam

`func NewArtifactsOpenApiCreditRemoteTeam(remoteTeamUrl string, userName string, remoteRepos []ArtifactsOpenApiRemoteRepoData, password string, ) *ArtifactsOpenApiCreditRemoteTeam`

NewArtifactsOpenApiCreditRemoteTeam instantiates a new ArtifactsOpenApiCreditRemoteTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactsOpenApiCreditRemoteTeamWithDefaults

`func NewArtifactsOpenApiCreditRemoteTeamWithDefaults() *ArtifactsOpenApiCreditRemoteTeam`

NewArtifactsOpenApiCreditRemoteTeamWithDefaults instantiates a new ArtifactsOpenApiCreditRemoteTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteTeamUrl

`func (o *ArtifactsOpenApiCreditRemoteTeam) GetRemoteTeamUrl() string`

GetRemoteTeamUrl returns the RemoteTeamUrl field if non-nil, zero value otherwise.

### GetRemoteTeamUrlOk

`func (o *ArtifactsOpenApiCreditRemoteTeam) GetRemoteTeamUrlOk() (*string, bool)`

GetRemoteTeamUrlOk returns a tuple with the RemoteTeamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTeamUrl

`func (o *ArtifactsOpenApiCreditRemoteTeam) SetRemoteTeamUrl(v string)`

SetRemoteTeamUrl sets RemoteTeamUrl field to given value.


### GetUserName

`func (o *ArtifactsOpenApiCreditRemoteTeam) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ArtifactsOpenApiCreditRemoteTeam) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ArtifactsOpenApiCreditRemoteTeam) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetRemoteRepos

`func (o *ArtifactsOpenApiCreditRemoteTeam) GetRemoteRepos() []ArtifactsOpenApiRemoteRepoData`

GetRemoteRepos returns the RemoteRepos field if non-nil, zero value otherwise.

### GetRemoteReposOk

`func (o *ArtifactsOpenApiCreditRemoteTeam) GetRemoteReposOk() (*[]ArtifactsOpenApiRemoteRepoData, bool)`

GetRemoteReposOk returns a tuple with the RemoteRepos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteRepos

`func (o *ArtifactsOpenApiCreditRemoteTeam) SetRemoteRepos(v []ArtifactsOpenApiRemoteRepoData)`

SetRemoteRepos sets RemoteRepos field to given value.


### GetPassword

`func (o *ArtifactsOpenApiCreditRemoteTeam) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ArtifactsOpenApiCreditRemoteTeam) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ArtifactsOpenApiCreditRemoteTeam) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


