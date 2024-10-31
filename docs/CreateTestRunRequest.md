# CreateTestRunRequest

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
**IncludeAll** | **bool** | 是否包含全部用例 | 
**Name** | **string** | 标题 | 
**ProjectName** | **string** | 项目名称 | 
**SectionId** | Pointer to **int64** | 分组 ID | [optional] 

## Methods

### NewCreateTestRunRequest

`func NewCreateTestRunRequest(includeAll bool, name string, projectName string, ) *CreateTestRunRequest`

NewCreateTestRunRequest instantiates a new CreateTestRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestRunRequestWithDefaults

`func NewCreateTestRunRequestWithDefaults() *CreateTestRunRequest`

NewCreateTestRunRequestWithDefaults instantiates a new CreateTestRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedToId

`func (o *CreateTestRunRequest) GetAssignedToId() int32`

GetAssignedToId returns the AssignedToId field if non-nil, zero value otherwise.

### GetAssignedToIdOk

`func (o *CreateTestRunRequest) GetAssignedToIdOk() (*int32, bool)`

GetAssignedToIdOk returns a tuple with the AssignedToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToId

`func (o *CreateTestRunRequest) SetAssignedToId(v int32)`

SetAssignedToId sets AssignedToId field to given value.

### HasAssignedToId

`func (o *CreateTestRunRequest) HasAssignedToId() bool`

HasAssignedToId returns a boolean if a field has been set.

### GetCases

`func (o *CreateTestRunRequest) GetCases() []int32`

GetCases returns the Cases field if non-nil, zero value otherwise.

### GetCasesOk

`func (o *CreateTestRunRequest) GetCasesOk() (*[]int32, bool)`

GetCasesOk returns a tuple with the Cases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCases

`func (o *CreateTestRunRequest) SetCases(v []int32)`

SetCases sets Cases field to given value.

### HasCases

`func (o *CreateTestRunRequest) HasCases() bool`

HasCases returns a boolean if a field has been set.

### GetConfigEnvironmentId

`func (o *CreateTestRunRequest) GetConfigEnvironmentId() int32`

GetConfigEnvironmentId returns the ConfigEnvironmentId field if non-nil, zero value otherwise.

### GetConfigEnvironmentIdOk

`func (o *CreateTestRunRequest) GetConfigEnvironmentIdOk() (*int32, bool)`

GetConfigEnvironmentIdOk returns a tuple with the ConfigEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigEnvironmentId

`func (o *CreateTestRunRequest) SetConfigEnvironmentId(v int32)`

SetConfigEnvironmentId sets ConfigEnvironmentId field to given value.

### HasConfigEnvironmentId

`func (o *CreateTestRunRequest) HasConfigEnvironmentId() bool`

HasConfigEnvironmentId returns a boolean if a field has been set.

### GetDescription

`func (o *CreateTestRunRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTestRunRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTestRunRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTestRunRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExecuteType

`func (o *CreateTestRunRequest) GetExecuteType() int32`

GetExecuteType returns the ExecuteType field if non-nil, zero value otherwise.

### GetExecuteTypeOk

`func (o *CreateTestRunRequest) GetExecuteTypeOk() (*int32, bool)`

GetExecuteTypeOk returns a tuple with the ExecuteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteType

`func (o *CreateTestRunRequest) SetExecuteType(v int32)`

SetExecuteType sets ExecuteType field to given value.

### HasExecuteType

`func (o *CreateTestRunRequest) HasExecuteType() bool`

HasExecuteType returns a boolean if a field has been set.

### GetGitDepotId

`func (o *CreateTestRunRequest) GetGitDepotId() int32`

GetGitDepotId returns the GitDepotId field if non-nil, zero value otherwise.

### GetGitDepotIdOk

`func (o *CreateTestRunRequest) GetGitDepotIdOk() (*int32, bool)`

GetGitDepotIdOk returns a tuple with the GitDepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitDepotId

`func (o *CreateTestRunRequest) SetGitDepotId(v int32)`

SetGitDepotId sets GitDepotId field to given value.

### HasGitDepotId

`func (o *CreateTestRunRequest) HasGitDepotId() bool`

HasGitDepotId returns a boolean if a field has been set.

### GetGitReleaseId

`func (o *CreateTestRunRequest) GetGitReleaseId() int32`

GetGitReleaseId returns the GitReleaseId field if non-nil, zero value otherwise.

### GetGitReleaseIdOk

`func (o *CreateTestRunRequest) GetGitReleaseIdOk() (*int32, bool)`

GetGitReleaseIdOk returns a tuple with the GitReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitReleaseId

`func (o *CreateTestRunRequest) SetGitReleaseId(v int32)`

SetGitReleaseId sets GitReleaseId field to given value.

### HasGitReleaseId

`func (o *CreateTestRunRequest) HasGitReleaseId() bool`

HasGitReleaseId returns a boolean if a field has been set.

### GetIncludeAll

`func (o *CreateTestRunRequest) GetIncludeAll() bool`

GetIncludeAll returns the IncludeAll field if non-nil, zero value otherwise.

### GetIncludeAllOk

`func (o *CreateTestRunRequest) GetIncludeAllOk() (*bool, bool)`

GetIncludeAllOk returns a tuple with the IncludeAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAll

`func (o *CreateTestRunRequest) SetIncludeAll(v bool)`

SetIncludeAll sets IncludeAll field to given value.


### GetName

`func (o *CreateTestRunRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTestRunRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTestRunRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProjectName

`func (o *CreateTestRunRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateTestRunRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateTestRunRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetSectionId

`func (o *CreateTestRunRequest) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *CreateTestRunRequest) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *CreateTestRunRequest) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *CreateTestRunRequest) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


