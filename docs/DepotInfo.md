# DepotInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | 创建时间 | [optional] 
**DefaultBranch** | Pointer to **NullableString** | 默认分支 | [optional] [default to ""]
**Description** | Pointer to **string** | 仓库描述 | [optional] [default to ""]
**GroupId** | Pointer to **NullableInt64** | 项目组Id | [optional] 
**GroupName** | Pointer to **NullableString** | 项目名称 | [optional] [default to ""]
**GroupType** | Pointer to **NullableString** | 项目类型 | [optional] [default to ""]
**HttpsUrl** | Pointer to **string** | 仓库的https地址 | [optional] [default to ""]
**Id** | Pointer to **int64** | 仓库id | [optional] 
**LastPushAt** | Pointer to **int64** | 最终push时间 | [optional] 
**Name** | Pointer to **string** | 仓库名称 | [optional] [default to ""]
**ProjectId** | Pointer to **int64** | 项目Id | [optional] 
**ProjectName** | Pointer to **string** | 项目名称 | [optional] [default to ""]
**RepoType** | Pointer to **string** | 仓库类型,具体值为git或者svn | [optional] [default to ""]
**SshUrl** | Pointer to **string** | 仓库的ssh地址 | [optional] [default to ""]
**VcsType** | Pointer to **string** | 仓库类型,具体值为git或者svn | [optional] [default to ""]
**WebUrl** | Pointer to **string** | 仓库webURL | [optional] [default to ""]

## Methods

### NewDepotInfo

`func NewDepotInfo() *DepotInfo`

NewDepotInfo instantiates a new DepotInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepotInfoWithDefaults

`func NewDepotInfoWithDefaults() *DepotInfo`

NewDepotInfoWithDefaults instantiates a new DepotInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DepotInfo) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DepotInfo) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DepotInfo) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DepotInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *DepotInfo) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *DepotInfo) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *DepotInfo) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *DepotInfo) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### SetDefaultBranchNil

`func (o *DepotInfo) SetDefaultBranchNil(b bool)`

 SetDefaultBranchNil sets the value for DefaultBranch to be an explicit nil

### UnsetDefaultBranch
`func (o *DepotInfo) UnsetDefaultBranch()`

UnsetDefaultBranch ensures that no value is present for DefaultBranch, not even an explicit nil
### GetDescription

`func (o *DepotInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DepotInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DepotInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DepotInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupId

`func (o *DepotInfo) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DepotInfo) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DepotInfo) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *DepotInfo) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *DepotInfo) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *DepotInfo) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetGroupName

`func (o *DepotInfo) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *DepotInfo) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *DepotInfo) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *DepotInfo) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *DepotInfo) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *DepotInfo) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetGroupType

`func (o *DepotInfo) GetGroupType() string`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *DepotInfo) GetGroupTypeOk() (*string, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *DepotInfo) SetGroupType(v string)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *DepotInfo) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### SetGroupTypeNil

`func (o *DepotInfo) SetGroupTypeNil(b bool)`

 SetGroupTypeNil sets the value for GroupType to be an explicit nil

### UnsetGroupType
`func (o *DepotInfo) UnsetGroupType()`

UnsetGroupType ensures that no value is present for GroupType, not even an explicit nil
### GetHttpsUrl

`func (o *DepotInfo) GetHttpsUrl() string`

GetHttpsUrl returns the HttpsUrl field if non-nil, zero value otherwise.

### GetHttpsUrlOk

`func (o *DepotInfo) GetHttpsUrlOk() (*string, bool)`

GetHttpsUrlOk returns a tuple with the HttpsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsUrl

`func (o *DepotInfo) SetHttpsUrl(v string)`

SetHttpsUrl sets HttpsUrl field to given value.

### HasHttpsUrl

`func (o *DepotInfo) HasHttpsUrl() bool`

HasHttpsUrl returns a boolean if a field has been set.

### GetId

`func (o *DepotInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DepotInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DepotInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DepotInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastPushAt

`func (o *DepotInfo) GetLastPushAt() int64`

GetLastPushAt returns the LastPushAt field if non-nil, zero value otherwise.

### GetLastPushAtOk

`func (o *DepotInfo) GetLastPushAtOk() (*int64, bool)`

GetLastPushAtOk returns a tuple with the LastPushAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPushAt

`func (o *DepotInfo) SetLastPushAt(v int64)`

SetLastPushAt sets LastPushAt field to given value.

### HasLastPushAt

`func (o *DepotInfo) HasLastPushAt() bool`

HasLastPushAt returns a boolean if a field has been set.

### GetName

`func (o *DepotInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DepotInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DepotInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DepotInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *DepotInfo) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DepotInfo) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DepotInfo) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DepotInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *DepotInfo) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DepotInfo) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DepotInfo) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *DepotInfo) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetRepoType

`func (o *DepotInfo) GetRepoType() string`

GetRepoType returns the RepoType field if non-nil, zero value otherwise.

### GetRepoTypeOk

`func (o *DepotInfo) GetRepoTypeOk() (*string, bool)`

GetRepoTypeOk returns a tuple with the RepoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoType

`func (o *DepotInfo) SetRepoType(v string)`

SetRepoType sets RepoType field to given value.

### HasRepoType

`func (o *DepotInfo) HasRepoType() bool`

HasRepoType returns a boolean if a field has been set.

### GetSshUrl

`func (o *DepotInfo) GetSshUrl() string`

GetSshUrl returns the SshUrl field if non-nil, zero value otherwise.

### GetSshUrlOk

`func (o *DepotInfo) GetSshUrlOk() (*string, bool)`

GetSshUrlOk returns a tuple with the SshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUrl

`func (o *DepotInfo) SetSshUrl(v string)`

SetSshUrl sets SshUrl field to given value.

### HasSshUrl

`func (o *DepotInfo) HasSshUrl() bool`

HasSshUrl returns a boolean if a field has been set.

### GetVcsType

`func (o *DepotInfo) GetVcsType() string`

GetVcsType returns the VcsType field if non-nil, zero value otherwise.

### GetVcsTypeOk

`func (o *DepotInfo) GetVcsTypeOk() (*string, bool)`

GetVcsTypeOk returns a tuple with the VcsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsType

`func (o *DepotInfo) SetVcsType(v string)`

SetVcsType sets VcsType field to given value.

### HasVcsType

`func (o *DepotInfo) HasVcsType() bool`

HasVcsType returns a boolean if a field has been set.

### GetWebUrl

`func (o *DepotInfo) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *DepotInfo) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *DepotInfo) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *DepotInfo) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


