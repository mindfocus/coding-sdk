# Run

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedToId** | Pointer to **NullableInt32** | 处理人ID | [optional] 
**BlockedCount** | Pointer to **NullableInt32** | 计划内阻塞测试数量 | [optional] 
**CompletedAt** | Pointer to **NullableString** | 归档时间 | [optional] [default to ""]
**ConfigEnvironmentId** | Pointer to **NullableInt32** | 环境标识 | [optional] 
**CreatedAt** | Pointer to **NullableString** | 创建时间 | [optional] [default to ""]
**CreatedBy** | Pointer to **NullableInt32** | 创建人ID | [optional] 
**Days** | Pointer to **NullableInt32** | 持续天数 | [optional] 
**Description** | Pointer to **NullableString** | 描述 | [optional] [default to ""]
**ExecuteType** | Pointer to **NullableInt32** | 执行方式: 1-手动执行 2-自动化流水线执行 | [optional] 
**FailedCount** | Pointer to **NullableInt32** | 计划内失败测试数量 | [optional] 
**GitDepotId** | Pointer to **NullableInt32** | 代码仓库 ID | [optional] 
**GitDepotName** | Pointer to **NullableString** | 代码仓库名名称 | [optional] [default to ""]
**GitReleaseId** | Pointer to **NullableInt32** | 发布版本 ID（资源 ID） | [optional] 
**GitReleaseName** | Pointer to **NullableString** | 发布版本名称 | [optional] [default to ""]
**GitReleaseState** | Pointer to **NullableInt32** | 发布版本名称状态：0-未发布 1-已发布 | [optional] 
**Id** | Pointer to **NullableInt32** | ID 主键 | [optional] 
**IncludeAll** | Pointer to **NullableBool** | 是否包含全部用例 | [optional] [default to false]
**IsCompleted** | Pointer to **NullableBool** | 是否归档 | [optional] [default to false]
**IterationId** | Pointer to **NullableInt64** | 所属迭代 ID | [optional] 
**IterationName** | Pointer to **NullableString** | 迭代名称 | [optional] [default to ""]
**Name** | Pointer to **NullableString** | 名称 | [optional] [default to ""]
**PassedCount** | Pointer to **NullableInt32** | 计划内通过测试数量 | [optional] 
**RetestCount** | Pointer to **NullableInt32** | 计划内重新测试数量 | [optional] 
**SectionId** | Pointer to **NullableInt64** | 分组 ID | [optional] 
**SectionName** | Pointer to **NullableString** | 分组名 | [optional] [default to ""]
**State** | Pointer to **NullableInt32** | 状态: 0-未开始 1-进行中 2-已测完 | [optional] 
**UntestedCount** | Pointer to **NullableInt32** | 计划内未测试数量 | [optional] 

## Methods

### NewRun

`func NewRun() *Run`

NewRun instantiates a new Run object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunWithDefaults

`func NewRunWithDefaults() *Run`

NewRunWithDefaults instantiates a new Run object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedToId

`func (o *Run) GetAssignedToId() int32`

GetAssignedToId returns the AssignedToId field if non-nil, zero value otherwise.

### GetAssignedToIdOk

`func (o *Run) GetAssignedToIdOk() (*int32, bool)`

GetAssignedToIdOk returns a tuple with the AssignedToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToId

`func (o *Run) SetAssignedToId(v int32)`

SetAssignedToId sets AssignedToId field to given value.

### HasAssignedToId

`func (o *Run) HasAssignedToId() bool`

HasAssignedToId returns a boolean if a field has been set.

### SetAssignedToIdNil

`func (o *Run) SetAssignedToIdNil(b bool)`

 SetAssignedToIdNil sets the value for AssignedToId to be an explicit nil

### UnsetAssignedToId
`func (o *Run) UnsetAssignedToId()`

UnsetAssignedToId ensures that no value is present for AssignedToId, not even an explicit nil
### GetBlockedCount

`func (o *Run) GetBlockedCount() int32`

GetBlockedCount returns the BlockedCount field if non-nil, zero value otherwise.

### GetBlockedCountOk

`func (o *Run) GetBlockedCountOk() (*int32, bool)`

GetBlockedCountOk returns a tuple with the BlockedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedCount

`func (o *Run) SetBlockedCount(v int32)`

SetBlockedCount sets BlockedCount field to given value.

### HasBlockedCount

`func (o *Run) HasBlockedCount() bool`

HasBlockedCount returns a boolean if a field has been set.

### SetBlockedCountNil

`func (o *Run) SetBlockedCountNil(b bool)`

 SetBlockedCountNil sets the value for BlockedCount to be an explicit nil

### UnsetBlockedCount
`func (o *Run) UnsetBlockedCount()`

UnsetBlockedCount ensures that no value is present for BlockedCount, not even an explicit nil
### GetCompletedAt

`func (o *Run) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *Run) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *Run) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *Run) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *Run) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *Run) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetConfigEnvironmentId

`func (o *Run) GetConfigEnvironmentId() int32`

GetConfigEnvironmentId returns the ConfigEnvironmentId field if non-nil, zero value otherwise.

### GetConfigEnvironmentIdOk

`func (o *Run) GetConfigEnvironmentIdOk() (*int32, bool)`

GetConfigEnvironmentIdOk returns a tuple with the ConfigEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigEnvironmentId

`func (o *Run) SetConfigEnvironmentId(v int32)`

SetConfigEnvironmentId sets ConfigEnvironmentId field to given value.

### HasConfigEnvironmentId

`func (o *Run) HasConfigEnvironmentId() bool`

HasConfigEnvironmentId returns a boolean if a field has been set.

### SetConfigEnvironmentIdNil

`func (o *Run) SetConfigEnvironmentIdNil(b bool)`

 SetConfigEnvironmentIdNil sets the value for ConfigEnvironmentId to be an explicit nil

### UnsetConfigEnvironmentId
`func (o *Run) UnsetConfigEnvironmentId()`

UnsetConfigEnvironmentId ensures that no value is present for ConfigEnvironmentId, not even an explicit nil
### GetCreatedAt

`func (o *Run) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Run) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Run) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Run) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Run) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Run) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *Run) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Run) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Run) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Run) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *Run) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *Run) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDays

`func (o *Run) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *Run) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *Run) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *Run) HasDays() bool`

HasDays returns a boolean if a field has been set.

### SetDaysNil

`func (o *Run) SetDaysNil(b bool)`

 SetDaysNil sets the value for Days to be an explicit nil

### UnsetDays
`func (o *Run) UnsetDays()`

UnsetDays ensures that no value is present for Days, not even an explicit nil
### GetDescription

`func (o *Run) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Run) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Run) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Run) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Run) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Run) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExecuteType

`func (o *Run) GetExecuteType() int32`

GetExecuteType returns the ExecuteType field if non-nil, zero value otherwise.

### GetExecuteTypeOk

`func (o *Run) GetExecuteTypeOk() (*int32, bool)`

GetExecuteTypeOk returns a tuple with the ExecuteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteType

`func (o *Run) SetExecuteType(v int32)`

SetExecuteType sets ExecuteType field to given value.

### HasExecuteType

`func (o *Run) HasExecuteType() bool`

HasExecuteType returns a boolean if a field has been set.

### SetExecuteTypeNil

`func (o *Run) SetExecuteTypeNil(b bool)`

 SetExecuteTypeNil sets the value for ExecuteType to be an explicit nil

### UnsetExecuteType
`func (o *Run) UnsetExecuteType()`

UnsetExecuteType ensures that no value is present for ExecuteType, not even an explicit nil
### GetFailedCount

`func (o *Run) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *Run) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *Run) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.

### HasFailedCount

`func (o *Run) HasFailedCount() bool`

HasFailedCount returns a boolean if a field has been set.

### SetFailedCountNil

`func (o *Run) SetFailedCountNil(b bool)`

 SetFailedCountNil sets the value for FailedCount to be an explicit nil

### UnsetFailedCount
`func (o *Run) UnsetFailedCount()`

UnsetFailedCount ensures that no value is present for FailedCount, not even an explicit nil
### GetGitDepotId

`func (o *Run) GetGitDepotId() int32`

GetGitDepotId returns the GitDepotId field if non-nil, zero value otherwise.

### GetGitDepotIdOk

`func (o *Run) GetGitDepotIdOk() (*int32, bool)`

GetGitDepotIdOk returns a tuple with the GitDepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitDepotId

`func (o *Run) SetGitDepotId(v int32)`

SetGitDepotId sets GitDepotId field to given value.

### HasGitDepotId

`func (o *Run) HasGitDepotId() bool`

HasGitDepotId returns a boolean if a field has been set.

### SetGitDepotIdNil

`func (o *Run) SetGitDepotIdNil(b bool)`

 SetGitDepotIdNil sets the value for GitDepotId to be an explicit nil

### UnsetGitDepotId
`func (o *Run) UnsetGitDepotId()`

UnsetGitDepotId ensures that no value is present for GitDepotId, not even an explicit nil
### GetGitDepotName

`func (o *Run) GetGitDepotName() string`

GetGitDepotName returns the GitDepotName field if non-nil, zero value otherwise.

### GetGitDepotNameOk

`func (o *Run) GetGitDepotNameOk() (*string, bool)`

GetGitDepotNameOk returns a tuple with the GitDepotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitDepotName

`func (o *Run) SetGitDepotName(v string)`

SetGitDepotName sets GitDepotName field to given value.

### HasGitDepotName

`func (o *Run) HasGitDepotName() bool`

HasGitDepotName returns a boolean if a field has been set.

### SetGitDepotNameNil

`func (o *Run) SetGitDepotNameNil(b bool)`

 SetGitDepotNameNil sets the value for GitDepotName to be an explicit nil

### UnsetGitDepotName
`func (o *Run) UnsetGitDepotName()`

UnsetGitDepotName ensures that no value is present for GitDepotName, not even an explicit nil
### GetGitReleaseId

`func (o *Run) GetGitReleaseId() int32`

GetGitReleaseId returns the GitReleaseId field if non-nil, zero value otherwise.

### GetGitReleaseIdOk

`func (o *Run) GetGitReleaseIdOk() (*int32, bool)`

GetGitReleaseIdOk returns a tuple with the GitReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitReleaseId

`func (o *Run) SetGitReleaseId(v int32)`

SetGitReleaseId sets GitReleaseId field to given value.

### HasGitReleaseId

`func (o *Run) HasGitReleaseId() bool`

HasGitReleaseId returns a boolean if a field has been set.

### SetGitReleaseIdNil

`func (o *Run) SetGitReleaseIdNil(b bool)`

 SetGitReleaseIdNil sets the value for GitReleaseId to be an explicit nil

### UnsetGitReleaseId
`func (o *Run) UnsetGitReleaseId()`

UnsetGitReleaseId ensures that no value is present for GitReleaseId, not even an explicit nil
### GetGitReleaseName

`func (o *Run) GetGitReleaseName() string`

GetGitReleaseName returns the GitReleaseName field if non-nil, zero value otherwise.

### GetGitReleaseNameOk

`func (o *Run) GetGitReleaseNameOk() (*string, bool)`

GetGitReleaseNameOk returns a tuple with the GitReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitReleaseName

`func (o *Run) SetGitReleaseName(v string)`

SetGitReleaseName sets GitReleaseName field to given value.

### HasGitReleaseName

`func (o *Run) HasGitReleaseName() bool`

HasGitReleaseName returns a boolean if a field has been set.

### SetGitReleaseNameNil

`func (o *Run) SetGitReleaseNameNil(b bool)`

 SetGitReleaseNameNil sets the value for GitReleaseName to be an explicit nil

### UnsetGitReleaseName
`func (o *Run) UnsetGitReleaseName()`

UnsetGitReleaseName ensures that no value is present for GitReleaseName, not even an explicit nil
### GetGitReleaseState

`func (o *Run) GetGitReleaseState() int32`

GetGitReleaseState returns the GitReleaseState field if non-nil, zero value otherwise.

### GetGitReleaseStateOk

`func (o *Run) GetGitReleaseStateOk() (*int32, bool)`

GetGitReleaseStateOk returns a tuple with the GitReleaseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitReleaseState

`func (o *Run) SetGitReleaseState(v int32)`

SetGitReleaseState sets GitReleaseState field to given value.

### HasGitReleaseState

`func (o *Run) HasGitReleaseState() bool`

HasGitReleaseState returns a boolean if a field has been set.

### SetGitReleaseStateNil

`func (o *Run) SetGitReleaseStateNil(b bool)`

 SetGitReleaseStateNil sets the value for GitReleaseState to be an explicit nil

### UnsetGitReleaseState
`func (o *Run) UnsetGitReleaseState()`

UnsetGitReleaseState ensures that no value is present for GitReleaseState, not even an explicit nil
### GetId

`func (o *Run) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Run) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Run) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Run) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Run) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Run) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIncludeAll

`func (o *Run) GetIncludeAll() bool`

GetIncludeAll returns the IncludeAll field if non-nil, zero value otherwise.

### GetIncludeAllOk

`func (o *Run) GetIncludeAllOk() (*bool, bool)`

GetIncludeAllOk returns a tuple with the IncludeAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAll

`func (o *Run) SetIncludeAll(v bool)`

SetIncludeAll sets IncludeAll field to given value.

### HasIncludeAll

`func (o *Run) HasIncludeAll() bool`

HasIncludeAll returns a boolean if a field has been set.

### SetIncludeAllNil

`func (o *Run) SetIncludeAllNil(b bool)`

 SetIncludeAllNil sets the value for IncludeAll to be an explicit nil

### UnsetIncludeAll
`func (o *Run) UnsetIncludeAll()`

UnsetIncludeAll ensures that no value is present for IncludeAll, not even an explicit nil
### GetIsCompleted

`func (o *Run) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *Run) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *Run) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.

### HasIsCompleted

`func (o *Run) HasIsCompleted() bool`

HasIsCompleted returns a boolean if a field has been set.

### SetIsCompletedNil

`func (o *Run) SetIsCompletedNil(b bool)`

 SetIsCompletedNil sets the value for IsCompleted to be an explicit nil

### UnsetIsCompleted
`func (o *Run) UnsetIsCompleted()`

UnsetIsCompleted ensures that no value is present for IsCompleted, not even an explicit nil
### GetIterationId

`func (o *Run) GetIterationId() int64`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *Run) GetIterationIdOk() (*int64, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *Run) SetIterationId(v int64)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *Run) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.

### SetIterationIdNil

`func (o *Run) SetIterationIdNil(b bool)`

 SetIterationIdNil sets the value for IterationId to be an explicit nil

### UnsetIterationId
`func (o *Run) UnsetIterationId()`

UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
### GetIterationName

`func (o *Run) GetIterationName() string`

GetIterationName returns the IterationName field if non-nil, zero value otherwise.

### GetIterationNameOk

`func (o *Run) GetIterationNameOk() (*string, bool)`

GetIterationNameOk returns a tuple with the IterationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationName

`func (o *Run) SetIterationName(v string)`

SetIterationName sets IterationName field to given value.

### HasIterationName

`func (o *Run) HasIterationName() bool`

HasIterationName returns a boolean if a field has been set.

### SetIterationNameNil

`func (o *Run) SetIterationNameNil(b bool)`

 SetIterationNameNil sets the value for IterationName to be an explicit nil

### UnsetIterationName
`func (o *Run) UnsetIterationName()`

UnsetIterationName ensures that no value is present for IterationName, not even an explicit nil
### GetName

`func (o *Run) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Run) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Run) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Run) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Run) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Run) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPassedCount

`func (o *Run) GetPassedCount() int32`

GetPassedCount returns the PassedCount field if non-nil, zero value otherwise.

### GetPassedCountOk

`func (o *Run) GetPassedCountOk() (*int32, bool)`

GetPassedCountOk returns a tuple with the PassedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassedCount

`func (o *Run) SetPassedCount(v int32)`

SetPassedCount sets PassedCount field to given value.

### HasPassedCount

`func (o *Run) HasPassedCount() bool`

HasPassedCount returns a boolean if a field has been set.

### SetPassedCountNil

`func (o *Run) SetPassedCountNil(b bool)`

 SetPassedCountNil sets the value for PassedCount to be an explicit nil

### UnsetPassedCount
`func (o *Run) UnsetPassedCount()`

UnsetPassedCount ensures that no value is present for PassedCount, not even an explicit nil
### GetRetestCount

`func (o *Run) GetRetestCount() int32`

GetRetestCount returns the RetestCount field if non-nil, zero value otherwise.

### GetRetestCountOk

`func (o *Run) GetRetestCountOk() (*int32, bool)`

GetRetestCountOk returns a tuple with the RetestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetestCount

`func (o *Run) SetRetestCount(v int32)`

SetRetestCount sets RetestCount field to given value.

### HasRetestCount

`func (o *Run) HasRetestCount() bool`

HasRetestCount returns a boolean if a field has been set.

### SetRetestCountNil

`func (o *Run) SetRetestCountNil(b bool)`

 SetRetestCountNil sets the value for RetestCount to be an explicit nil

### UnsetRetestCount
`func (o *Run) UnsetRetestCount()`

UnsetRetestCount ensures that no value is present for RetestCount, not even an explicit nil
### GetSectionId

`func (o *Run) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *Run) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *Run) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *Run) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### SetSectionIdNil

`func (o *Run) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *Run) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetSectionName

`func (o *Run) GetSectionName() string`

GetSectionName returns the SectionName field if non-nil, zero value otherwise.

### GetSectionNameOk

`func (o *Run) GetSectionNameOk() (*string, bool)`

GetSectionNameOk returns a tuple with the SectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionName

`func (o *Run) SetSectionName(v string)`

SetSectionName sets SectionName field to given value.

### HasSectionName

`func (o *Run) HasSectionName() bool`

HasSectionName returns a boolean if a field has been set.

### SetSectionNameNil

`func (o *Run) SetSectionNameNil(b bool)`

 SetSectionNameNil sets the value for SectionName to be an explicit nil

### UnsetSectionName
`func (o *Run) UnsetSectionName()`

UnsetSectionName ensures that no value is present for SectionName, not even an explicit nil
### GetState

`func (o *Run) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Run) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Run) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *Run) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *Run) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *Run) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetUntestedCount

`func (o *Run) GetUntestedCount() int32`

GetUntestedCount returns the UntestedCount field if non-nil, zero value otherwise.

### GetUntestedCountOk

`func (o *Run) GetUntestedCountOk() (*int32, bool)`

GetUntestedCountOk returns a tuple with the UntestedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntestedCount

`func (o *Run) SetUntestedCount(v int32)`

SetUntestedCount sets UntestedCount field to given value.

### HasUntestedCount

`func (o *Run) HasUntestedCount() bool`

HasUntestedCount returns a boolean if a field has been set.

### SetUntestedCountNil

`func (o *Run) SetUntestedCountNil(b bool)`

 SetUntestedCountNil sets the value for UntestedCount to be an explicit nil

### UnsetUntestedCount
`func (o *Run) UnsetUntestedCount()`

UnsetUntestedCount ensures that no value is present for UntestedCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


