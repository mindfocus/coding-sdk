# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | Pointer to **NullableBool** | 是否压缩 | [optional] [default to false]
**CreatedAt** | Pointer to **NullableInt32** | 创建时间 | [optional] 
**Description** | Pointer to **NullableString** | 描述 | [optional] [default to ""]
**DisplayName** | Pointer to **NullableString** | 显示名称 | [optional] [default to ""]
**EndDate** | Pointer to **NullableInt32** | 项目结束时间 | [optional] 
**Icon** | Pointer to **NullableString** | 图标 | [optional] [default to ""]
**Id** | Pointer to **int32** | 项目 ID | [optional] 
**IsDemo** | Pointer to **NullableBool** | 是否为模板项目 | [optional] [default to false]
**MaxMember** | Pointer to **NullableInt32** | 最大团员数 | [optional] 
**Name** | Pointer to **NullableString** | 名称 | [optional] [default to ""]
**StartDate** | Pointer to **NullableInt32** | 项目开始时间 | [optional] 
**Status** | Pointer to **NullableInt32** | 状态 默认都为1 | [optional] 
**TeamId** | Pointer to **NullableInt32** | 团队 ID | [optional] 
**TeamOwnerId** | Pointer to **NullableInt32** | 团队所有者 ID | [optional] 
**Type** | Pointer to **NullableInt32** | 项目类型, 项目集为0，公开项目为1，私密项目为2 | [optional] 
**UpdatedAt** | Pointer to **NullableInt32** | 更新时间 | [optional] 
**UserOwnerId** | Pointer to **NullableInt32** | 个人所有者 ID | [optional] 
**ProgramIds** | Pointer to **[]uint64** | 项目集id | [optional] 

## Methods

### NewProject

`func NewProject() *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *Project) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *Project) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *Project) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *Project) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### SetArchivedNil

`func (o *Project) SetArchivedNil(b bool)`

 SetArchivedNil sets the value for Archived to be an explicit nil

### UnsetArchived
`func (o *Project) UnsetArchived()`

UnsetArchived ensures that no value is present for Archived, not even an explicit nil
### GetCreatedAt

`func (o *Project) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Project) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Project) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Project) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Project) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Project) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDescription

`func (o *Project) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Project) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Project) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Project) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Project) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Project) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *Project) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Project) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Project) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Project) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Project) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Project) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEndDate

`func (o *Project) GetEndDate() int32`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Project) GetEndDateOk() (*int32, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Project) SetEndDate(v int32)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Project) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *Project) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *Project) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetIcon

`func (o *Project) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Project) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Project) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Project) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *Project) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *Project) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetId

`func (o *Project) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Project) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDemo

`func (o *Project) GetIsDemo() bool`

GetIsDemo returns the IsDemo field if non-nil, zero value otherwise.

### GetIsDemoOk

`func (o *Project) GetIsDemoOk() (*bool, bool)`

GetIsDemoOk returns a tuple with the IsDemo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDemo

`func (o *Project) SetIsDemo(v bool)`

SetIsDemo sets IsDemo field to given value.

### HasIsDemo

`func (o *Project) HasIsDemo() bool`

HasIsDemo returns a boolean if a field has been set.

### SetIsDemoNil

`func (o *Project) SetIsDemoNil(b bool)`

 SetIsDemoNil sets the value for IsDemo to be an explicit nil

### UnsetIsDemo
`func (o *Project) UnsetIsDemo()`

UnsetIsDemo ensures that no value is present for IsDemo, not even an explicit nil
### GetMaxMember

`func (o *Project) GetMaxMember() int32`

GetMaxMember returns the MaxMember field if non-nil, zero value otherwise.

### GetMaxMemberOk

`func (o *Project) GetMaxMemberOk() (*int32, bool)`

GetMaxMemberOk returns a tuple with the MaxMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMember

`func (o *Project) SetMaxMember(v int32)`

SetMaxMember sets MaxMember field to given value.

### HasMaxMember

`func (o *Project) HasMaxMember() bool`

HasMaxMember returns a boolean if a field has been set.

### SetMaxMemberNil

`func (o *Project) SetMaxMemberNil(b bool)`

 SetMaxMemberNil sets the value for MaxMember to be an explicit nil

### UnsetMaxMember
`func (o *Project) UnsetMaxMember()`

UnsetMaxMember ensures that no value is present for MaxMember, not even an explicit nil
### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Project) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Project) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Project) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStartDate

`func (o *Project) GetStartDate() int32`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Project) GetStartDateOk() (*int32, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Project) SetStartDate(v int32)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Project) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *Project) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *Project) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetStatus

`func (o *Project) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Project) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Project) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Project) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Project) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Project) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTeamId

`func (o *Project) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *Project) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *Project) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *Project) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *Project) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *Project) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil
### GetTeamOwnerId

`func (o *Project) GetTeamOwnerId() int32`

GetTeamOwnerId returns the TeamOwnerId field if non-nil, zero value otherwise.

### GetTeamOwnerIdOk

`func (o *Project) GetTeamOwnerIdOk() (*int32, bool)`

GetTeamOwnerIdOk returns a tuple with the TeamOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamOwnerId

`func (o *Project) SetTeamOwnerId(v int32)`

SetTeamOwnerId sets TeamOwnerId field to given value.

### HasTeamOwnerId

`func (o *Project) HasTeamOwnerId() bool`

HasTeamOwnerId returns a boolean if a field has been set.

### SetTeamOwnerIdNil

`func (o *Project) SetTeamOwnerIdNil(b bool)`

 SetTeamOwnerIdNil sets the value for TeamOwnerId to be an explicit nil

### UnsetTeamOwnerId
`func (o *Project) UnsetTeamOwnerId()`

UnsetTeamOwnerId ensures that no value is present for TeamOwnerId, not even an explicit nil
### GetType

`func (o *Project) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Project) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Project) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *Project) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Project) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Project) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUpdatedAt

`func (o *Project) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Project) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Project) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Project) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Project) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Project) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserOwnerId

`func (o *Project) GetUserOwnerId() int32`

GetUserOwnerId returns the UserOwnerId field if non-nil, zero value otherwise.

### GetUserOwnerIdOk

`func (o *Project) GetUserOwnerIdOk() (*int32, bool)`

GetUserOwnerIdOk returns a tuple with the UserOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOwnerId

`func (o *Project) SetUserOwnerId(v int32)`

SetUserOwnerId sets UserOwnerId field to given value.

### HasUserOwnerId

`func (o *Project) HasUserOwnerId() bool`

HasUserOwnerId returns a boolean if a field has been set.

### SetUserOwnerIdNil

`func (o *Project) SetUserOwnerIdNil(b bool)`

 SetUserOwnerIdNil sets the value for UserOwnerId to be an explicit nil

### UnsetUserOwnerId
`func (o *Project) UnsetUserOwnerId()`

UnsetUserOwnerId ensures that no value is present for UserOwnerId, not even an explicit nil
### GetProgramIds

`func (o *Project) GetProgramIds() []uint64`

GetProgramIds returns the ProgramIds field if non-nil, zero value otherwise.

### GetProgramIdsOk

`func (o *Project) GetProgramIdsOk() (*[]uint64, bool)`

GetProgramIdsOk returns a tuple with the ProgramIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramIds

`func (o *Project) SetProgramIds(v []uint64)`

SetProgramIds sets ProgramIds field to given value.

### HasProgramIds

`func (o *Project) HasProgramIds() bool`

HasProgramIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


