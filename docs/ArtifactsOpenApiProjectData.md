# ArtifactsOpenApiProjectData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int64** | 是否生效，0-否，1-是 | [default to 0]
**HtmlUrl** | **string** | 项目URL | [default to ""]
**UpdatedAt** | **int64** | 更新时间 | [default to 0]
**TeamId** | **int64** | 团队ID | [default to 0]
**Name** | **string** | 名称 | [default to ""]
**Type** | **int64** | 项目类型 | [default to 0]
**ApiUrl** | **string** | 接口URL | [default to ""]
**DisplayName** | **string** | 显示名称 | [default to ""]
**OwnerUserName** | **string** | 项目所属用户名 | [default to ""]
**Id** | **int64** | ID | [default to 0]
**Icon** | **string** | 图标 | [default to ""]
**TeamOwnerId** | **int64** | 团队所有者 ID | [default to 0]
**ProjectPath** | **string** | 项目位置 | [default to ""]
**PmTypeName** | **string** | 项目管理类型 | [default to ""]
**CreatedAt** | **int64** | 创建时间 | [default to 0]

## Methods

### NewArtifactsOpenApiProjectData

`func NewArtifactsOpenApiProjectData(status int64, htmlUrl string, updatedAt int64, teamId int64, name string, type_ int64, apiUrl string, displayName string, ownerUserName string, id int64, icon string, teamOwnerId int64, projectPath string, pmTypeName string, createdAt int64, ) *ArtifactsOpenApiProjectData`

NewArtifactsOpenApiProjectData instantiates a new ArtifactsOpenApiProjectData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactsOpenApiProjectDataWithDefaults

`func NewArtifactsOpenApiProjectDataWithDefaults() *ArtifactsOpenApiProjectData`

NewArtifactsOpenApiProjectDataWithDefaults instantiates a new ArtifactsOpenApiProjectData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ArtifactsOpenApiProjectData) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArtifactsOpenApiProjectData) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArtifactsOpenApiProjectData) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetHtmlUrl

`func (o *ArtifactsOpenApiProjectData) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *ArtifactsOpenApiProjectData) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *ArtifactsOpenApiProjectData) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetUpdatedAt

`func (o *ArtifactsOpenApiProjectData) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ArtifactsOpenApiProjectData) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ArtifactsOpenApiProjectData) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetTeamId

`func (o *ArtifactsOpenApiProjectData) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *ArtifactsOpenApiProjectData) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *ArtifactsOpenApiProjectData) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.


### GetName

`func (o *ArtifactsOpenApiProjectData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactsOpenApiProjectData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactsOpenApiProjectData) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ArtifactsOpenApiProjectData) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArtifactsOpenApiProjectData) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArtifactsOpenApiProjectData) SetType(v int64)`

SetType sets Type field to given value.


### GetApiUrl

`func (o *ArtifactsOpenApiProjectData) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *ArtifactsOpenApiProjectData) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *ArtifactsOpenApiProjectData) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.


### GetDisplayName

`func (o *ArtifactsOpenApiProjectData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ArtifactsOpenApiProjectData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ArtifactsOpenApiProjectData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetOwnerUserName

`func (o *ArtifactsOpenApiProjectData) GetOwnerUserName() string`

GetOwnerUserName returns the OwnerUserName field if non-nil, zero value otherwise.

### GetOwnerUserNameOk

`func (o *ArtifactsOpenApiProjectData) GetOwnerUserNameOk() (*string, bool)`

GetOwnerUserNameOk returns a tuple with the OwnerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserName

`func (o *ArtifactsOpenApiProjectData) SetOwnerUserName(v string)`

SetOwnerUserName sets OwnerUserName field to given value.


### GetId

`func (o *ArtifactsOpenApiProjectData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactsOpenApiProjectData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactsOpenApiProjectData) SetId(v int64)`

SetId sets Id field to given value.


### GetIcon

`func (o *ArtifactsOpenApiProjectData) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *ArtifactsOpenApiProjectData) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *ArtifactsOpenApiProjectData) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetTeamOwnerId

`func (o *ArtifactsOpenApiProjectData) GetTeamOwnerId() int64`

GetTeamOwnerId returns the TeamOwnerId field if non-nil, zero value otherwise.

### GetTeamOwnerIdOk

`func (o *ArtifactsOpenApiProjectData) GetTeamOwnerIdOk() (*int64, bool)`

GetTeamOwnerIdOk returns a tuple with the TeamOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamOwnerId

`func (o *ArtifactsOpenApiProjectData) SetTeamOwnerId(v int64)`

SetTeamOwnerId sets TeamOwnerId field to given value.


### GetProjectPath

`func (o *ArtifactsOpenApiProjectData) GetProjectPath() string`

GetProjectPath returns the ProjectPath field if non-nil, zero value otherwise.

### GetProjectPathOk

`func (o *ArtifactsOpenApiProjectData) GetProjectPathOk() (*string, bool)`

GetProjectPathOk returns a tuple with the ProjectPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPath

`func (o *ArtifactsOpenApiProjectData) SetProjectPath(v string)`

SetProjectPath sets ProjectPath field to given value.


### GetPmTypeName

`func (o *ArtifactsOpenApiProjectData) GetPmTypeName() string`

GetPmTypeName returns the PmTypeName field if non-nil, zero value otherwise.

### GetPmTypeNameOk

`func (o *ArtifactsOpenApiProjectData) GetPmTypeNameOk() (*string, bool)`

GetPmTypeNameOk returns a tuple with the PmTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmTypeName

`func (o *ArtifactsOpenApiProjectData) SetPmTypeName(v string)`

SetPmTypeName sets PmTypeName field to given value.


### GetCreatedAt

`func (o *ArtifactsOpenApiProjectData) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArtifactsOpenApiProjectData) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArtifactsOpenApiProjectData) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


