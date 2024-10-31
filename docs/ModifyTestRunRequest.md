# ModifyTestRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedToId** | Pointer to **int32** | 处理人 ID | [optional] 
**Cases** | Pointer to **[]int32** | 包含的用例 ID 列表，IncludeAll&#x3D;false 必填 | [optional] 
**ConfigEnvironmentId** | Pointer to **int32** | 环境标识 | [optional] 
**Description** | Pointer to **string** | 描述 | [optional] 
**ExecuteType** | Pointer to **int32** | 执行类型：1-手动执行 2-自动化流水线执行 | [optional] 
**GitDepotId** | Pointer to **int32** | 项目代码库 ID | [optional] 
**GitReleaseId** | Pointer to **int32** | 发布版本 ID | [optional] 
**IncludeAll** | Pointer to **bool** | 是否包含全部用例 | [optional] 
**Name** | Pointer to **string** | 标题 | [optional] 
**ProjectName** | **string** | 项目名称 | 
**RunId** | **int32** | 计划 ID | 
**SectionId** | Pointer to **int64** | 分组 ID | [optional] 

## Methods

### NewModifyTestRunRequest

`func NewModifyTestRunRequest(projectName string, runId int32, ) *ModifyTestRunRequest`

NewModifyTestRunRequest instantiates a new ModifyTestRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyTestRunRequestWithDefaults

`func NewModifyTestRunRequestWithDefaults() *ModifyTestRunRequest`

NewModifyTestRunRequestWithDefaults instantiates a new ModifyTestRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedToId

`func (o *ModifyTestRunRequest) GetAssignedToId() int32`

GetAssignedToId returns the AssignedToId field if non-nil, zero value otherwise.

### GetAssignedToIdOk

`func (o *ModifyTestRunRequest) GetAssignedToIdOk() (*int32, bool)`

GetAssignedToIdOk returns a tuple with the AssignedToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToId

`func (o *ModifyTestRunRequest) SetAssignedToId(v int32)`

SetAssignedToId sets AssignedToId field to given value.

### HasAssignedToId

`func (o *ModifyTestRunRequest) HasAssignedToId() bool`

HasAssignedToId returns a boolean if a field has been set.

### GetCases

`func (o *ModifyTestRunRequest) GetCases() []int32`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *ModifyTestRunRequest) GetCasesOk() (*[]int32, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *ModifyTestRunRequest) SetCases(v []int32)`

SetCases sets Cases field to given value.

### HasCases

`func (o *ModifyTestRunRequest) HasCases() bool`

HasCases returns a boolean if a field has been set.

### GetConfigEnvironmentId

`func (o *ModifyTestRunRequest) GetConfigEnvironmentId() int32`

GetConfigEnvironmentId returns the ConfigEnvironmentId field if non-nil, zero value otherwise.

### GetConfigEnvironmentIdOk

`func (o *ModifyTestRunRequest) GetConfigEnvironmentIdOk() (*int32, bool)`

GetConfigEnvironmentIdOk returns a tuple with the ConfigEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigEnvironmentId

`func (o *ModifyTestRunRequest) SetConfigEnvironmentId(v int32)`

SetConfigEnvironmentId sets ConfigEnvironmentId field to given value.

### HasConfigEnvironmentId

`func (o *ModifyTestRunRequest) HasConfigEnvironmentId() bool`

HasConfigEnvironmentId returns a boolean if a field has been set.

### GetDescription

`func (o *ModifyTestRunRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModifyTestRunRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModifyTestRunRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModifyTestRunRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecuteType

`func (o *ModifyTestRunRequest) GetExecuteType() int32`

GetExecuteType returns the ExecuteType field if non-nil, zero value otherwise.

### GetExecuteTypeOk

`func (o *ModifyTestRunRequest) GetExecuteTypeOk() (*int32, bool)`

GetExecuteTypeOk returns a tuple with the ExecuteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteType

`func (o *ModifyTestRunRequest) SetExecuteType(v int32)`

SetExecuteType sets ExecuteType field to given value.

### HasExecuteType

`func (o *ModifyTestRunRequest) HasExecuteType() bool`

HasExecuteType returns a boolean if a field has been set.

### GetGitDepotId

`func (o *ModifyTestRunRequest) GetGitDepotId() int32`

GetGitDepotId returns the GitDepotId field if non-nil, zero value otherwise.

### GetGitDepotIdOk

`func (o *ModifyTestRunRequest) GetGitDepotIdOk() (*int32, bool)`

GetGitDepotIdOk returns a tuple with the GitDepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitDepotId

`func (o *ModifyTestRunRequest) SetGitDepotId(v int32)`

SetGitDepotId sets GitDepotId field to given value.

### HasGitDepotId

`func (o *ModifyTestRunRequest) HasGitDepotId() bool`

HasGitDepotId returns a boolean if a field has been set.

### GetGitReleaseId

`func (o *ModifyTestRunRequest) GetGitReleaseId() int32`

GetGitReleaseId returns the GitReleaseId field if non-nil, zero value otherwise.

### GetGitReleaseIdOk

`func (o *ModifyTestRunRequest) GetGitReleaseIdOk() (*int32, bool)`

GetGitReleaseIdOk returns a tuple with the GitReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitReleaseId

`func (o *ModifyTestRunRequest) SetGitReleaseId(v int32)`

SetGitReleaseId sets GitReleaseId field to given value.

### HasGitReleaseId

`func (o *ModifyTestRunRequest) HasGitReleaseId() bool`

HasGitReleaseId returns a boolean if a field has been set.

### GetIncludeAll

`func (o *ModifyTestRunRequest) GetIncludeAll() bool`

GetIncludeAll returns the IncludeAll field if non-nil, zero value otherwise.

### GetIncludeAllOk

`func (o *ModifyTestRunRequest) GetIncludeAllOk() (*bool, bool)`

GetIncludeAllOk returns a tuple with the IncludeAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAll

`func (o *ModifyTestRunRequest) SetIncludeAll(v bool)`

SetIncludeAll sets IncludeAll field to given value.

### HasIncludeAll

`func (o *ModifyTestRunRequest) HasIncludeAll() bool`

HasIncludeAll returns a boolean if a field has been set.

### GetName

`func (o *ModifyTestRunRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyTestRunRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyTestRunRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModifyTestRunRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectName

`func (o *ModifyTestRunRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ModifyTestRunRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ModifyTestRunRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetRunId

`func (o *ModifyTestRunRequest) GetRunId() int32`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ModifyTestRunRequest) GetRunIdOk() (*int32, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ModifyTestRunRequest) SetRunId(v int32)`

SetRunId sets RunId field to given value.


### GetSectionId

`func (o *ModifyTestRunRequest) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *ModifyTestRunRequest) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *ModifyTestRunRequest) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *ModifyTestRunRequest) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


