# CreateCodingCIJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlwaysUserIdList** | Pointer to **[]int32** | 不管构建成功还是失败总是通知的用户 | [optional] 
**AutoCancelSameMergeRequest** | **bool** | 自动取消相同 MR | 
**AutoCancelSameRevision** | **bool** | 自动取消相同版本 | 
**BranchRegex** | Pointer to **string** | hookType 为 CUSTOME 时须指定 | [optional] 
**BranchSelector** | Pointer to **string** | hookType 为 DEFAULT 时须指定 | [optional] 
**BuildFailUserIdList** | Pointer to **[]int32** | 仅构建失败时要通知的用户 | [optional] 
**CachePathList** | Pointer to [**[]CodingCIJobCachePath**](CodingCIJobCachePath.md) | 任务缓存目录配置 | [optional] 
**DepotId** | **int32** | 仓库 ID | 
**DepotType** | **string** | 仓库类型 CODING_OTHER_PROJ,CODING,TGIT,GITHUB,GITLAB,GITLAB_PRIVATE,GITEE,COMMON_GIT_REPO,NONE | 
**EnvList** | Pointer to [**[]CIJobEnv**](CIJobEnv.md) | 环境变量配置 | [optional] 
**ExecuteIn** | **string** | 执行方式 CVM | STATIC | AGENT | 
**ExecutedAgentPoolId** | Pointer to **int32** | 自定义构建节点池 ID，ExecuteIn 为 AGENT 必填 | [optional] 
**HookType** | **string** | 代码更新触发匹配规则 DEFAULT,TAG,BRANCH,CUSTOM | 
**JenkinsFileFromType** | **string** | STATIC，SCM 从代码库读取 | 
**JenkinsFilePath** | Pointer to **string** | JenkinsFileFromType 为 SCM 必填 | [optional] 
**JenkinsFileStaticContent** | Pointer to **string** | JenkinsFileFromType 为 STATIC 必填 | [optional] 
**JobFromType** | **string** | 构建计划来源 TKE TCB | 
**Name** | **string** | 构建计划名称 | 
**ProjectId** | **int32** | 项目 ID | 
**ScheduleList** | Pointer to [**[]CodingCIJobSchedule**](CodingCIJobSchedule.md) | 针对 CRON triggerMethod 的 schedule 规则配置, 暂只用于添加 | [optional] 
**TriggerMethodList** | Pointer to **[]string** | REF_CHANGE 代码更新触发      CRON &#x3D; 1 定时触发      MR_CHANGE  MR变动触发 | [optional] 
**TriggerRemind** | Pointer to **string** | 构建结果通知触发者机制  ALWAYS -总是通知;  BUILD_FAIL -仅构建失败时通知; | [optional] 

## Methods

### NewCreateCodingCIJobRequest

`func NewCreateCodingCIJobRequest(autoCancelSameMergeRequest bool, autoCancelSameRevision bool, depotId int32, depotType string, executeIn string, hookType string, jenkinsFileFromType string, jobFromType string, name string, projectId int32, ) *CreateCodingCIJobRequest`

NewCreateCodingCIJobRequest instantiates a new CreateCodingCIJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCodingCIJobRequestWithDefaults

`func NewCreateCodingCIJobRequestWithDefaults() *CreateCodingCIJobRequest`

NewCreateCodingCIJobRequestWithDefaults instantiates a new CreateCodingCIJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlwaysUserIdList

`func (o *CreateCodingCIJobRequest) GetAlwaysUserIdList() []int32`

GetAlwaysUserIdList returns the AlwaysUserIdList field if non-nil, zero value otherwise.

### GetAlwaysUserIdListOk

`func (o *CreateCodingCIJobRequest) GetAlwaysUserIdListOk() (*[]int32, bool)`

GetAlwaysUserIdListOk returns a tuple with the AlwaysUserIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysUserIdList

`func (o *CreateCodingCIJobRequest) SetAlwaysUserIdList(v []int32)`

SetAlwaysUserIdList sets AlwaysUserIdList field to given value.

### HasAlwaysUserIdList

`func (o *CreateCodingCIJobRequest) HasAlwaysUserIdList() bool`

HasAlwaysUserIdList returns a boolean if a field has been set.

### GetAutoCancelSameMergeRequest

`func (o *CreateCodingCIJobRequest) GetAutoCancelSameMergeRequest() bool`

GetAutoCancelSameMergeRequest returns the AutoCancelSameMergeRequest field if non-nil, zero value otherwise.

### GetAutoCancelSameMergeRequestOk

`func (o *CreateCodingCIJobRequest) GetAutoCancelSameMergeRequestOk() (*bool, bool)`

GetAutoCancelSameMergeRequestOk returns a tuple with the AutoCancelSameMergeRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCancelSameMergeRequest

`func (o *CreateCodingCIJobRequest) SetAutoCancelSameMergeRequest(v bool)`

SetAutoCancelSameMergeRequest sets AutoCancelSameMergeRequest field to given value.


### GetAutoCancelSameRevision

`func (o *CreateCodingCIJobRequest) GetAutoCancelSameRevision() bool`

GetAutoCancelSameRevision returns the AutoCancelSameRevision field if non-nil, zero value otherwise.

### GetAutoCancelSameRevisionOk

`func (o *CreateCodingCIJobRequest) GetAutoCancelSameRevisionOk() (*bool, bool)`

GetAutoCancelSameRevisionOk returns a tuple with the AutoCancelSameRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCancelSameRevision

`func (o *CreateCodingCIJobRequest) SetAutoCancelSameRevision(v bool)`

SetAutoCancelSameRevision sets AutoCancelSameRevision field to given value.


### GetBranchRegex

`func (o *CreateCodingCIJobRequest) GetBranchRegex() string`

GetBranchRegex returns the BranchRegex field if non-nil, zero value otherwise.

### GetBranchRegexOk

`func (o *CreateCodingCIJobRequest) GetBranchRegexOk() (*string, bool)`

GetBranchRegexOk returns a tuple with the BranchRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchRegex

`func (o *CreateCodingCIJobRequest) SetBranchRegex(v string)`

SetBranchRegex sets BranchRegex field to given value.

### HasBranchRegex

`func (o *CreateCodingCIJobRequest) HasBranchRegex() bool`

HasBranchRegex returns a boolean if a field has been set.

### GetBranchSelector

`func (o *CreateCodingCIJobRequest) GetBranchSelector() string`

GetBranchSelector returns the BranchSelector field if non-nil, zero value otherwise.

### GetBranchSelectorOk

`func (o *CreateCodingCIJobRequest) GetBranchSelectorOk() (*string, bool)`

GetBranchSelectorOk returns a tuple with the BranchSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchSelector

`func (o *CreateCodingCIJobRequest) SetBranchSelector(v string)`

SetBranchSelector sets BranchSelector field to given value.

### HasBranchSelector

`func (o *CreateCodingCIJobRequest) HasBranchSelector() bool`

HasBranchSelector returns a boolean if a field has been set.

### GetBuildFailUserIdList

`func (o *CreateCodingCIJobRequest) GetBuildFailUserIdList() []int32`

GetBuildFailUserIdList returns the BuildFailUserIdList field if non-nil, zero value otherwise.

### GetBuildFailUserIdListOk

`func (o *CreateCodingCIJobRequest) GetBuildFailUserIdListOk() (*[]int32, bool)`

GetBuildFailUserIdListOk returns a tuple with the BuildFailUserIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildFailUserIdList

`func (o *CreateCodingCIJobRequest) SetBuildFailUserIdList(v []int32)`

SetBuildFailUserIdList sets BuildFailUserIdList field to given value.

### HasBuildFailUserIdList

`func (o *CreateCodingCIJobRequest) HasBuildFailUserIdList() bool`

HasBuildFailUserIdList returns a boolean if a field has been set.

### GetCachePathList

`func (o *CreateCodingCIJobRequest) GetCachePathList() []CodingCIJobCachePath`

GetCachePathList returns the CachePathList field if non-nil, zero value otherwise.

### GetCachePathListOk

`func (o *CreateCodingCIJobRequest) GetCachePathListOk() (*[]CodingCIJobCachePath, bool)`

GetCachePathListOk returns a tuple with the CachePathList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePathList

`func (o *CreateCodingCIJobRequest) SetCachePathList(v []CodingCIJobCachePath)`

SetCachePathList sets CachePathList field to given value.

### HasCachePathList

`func (o *CreateCodingCIJobRequest) HasCachePathList() bool`

HasCachePathList returns a boolean if a field has been set.

### GetDepotId

`func (o *CreateCodingCIJobRequest) GetDepotId() int32`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CreateCodingCIJobRequest) GetDepotIdOk() (*int32, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CreateCodingCIJobRequest) SetDepotId(v int32)`

SetDepotId sets DepotId field to given value.


### GetDepotType

`func (o *CreateCodingCIJobRequest) GetDepotType() string`

GetDepotType returns the DepotType field if non-nil, zero value otherwise.

### GetDepotTypeOk

`func (o *CreateCodingCIJobRequest) GetDepotTypeOk() (*string, bool)`

GetDepotTypeOk returns a tuple with the DepotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotType

`func (o *CreateCodingCIJobRequest) SetDepotType(v string)`

SetDepotType sets DepotType field to given value.


### GetEnvList

`func (o *CreateCodingCIJobRequest) GetEnvList() []CIJobEnv`

GetEnvList returns the EnvList field if non-nil, zero value otherwise.

### GetEnvListOk

`func (o *CreateCodingCIJobRequest) GetEnvListOk() (*[]CIJobEnv, bool)`

GetEnvListOk returns a tuple with the EnvList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvList

`func (o *CreateCodingCIJobRequest) SetEnvList(v []CIJobEnv)`

SetEnvList sets EnvList field to given value.

### HasEnvList

`func (o *CreateCodingCIJobRequest) HasEnvList() bool`

HasEnvList returns a boolean if a field has been set.

### GetExecuteIn

`func (o *CreateCodingCIJobRequest) GetExecuteIn() string`

GetExecuteIn returns the ExecuteIn field if non-nil, zero value otherwise.

### GetExecuteInOk

`func (o *CreateCodingCIJobRequest) GetExecuteInOk() (*string, bool)`

GetExecuteInOk returns a tuple with the ExecuteIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteIn

`func (o *CreateCodingCIJobRequest) SetExecuteIn(v string)`

SetExecuteIn sets ExecuteIn field to given value.


### GetExecutedAgentPoolId

`func (o *CreateCodingCIJobRequest) GetExecutedAgentPoolId() int32`

GetExecutedAgentPoolId returns the ExecutedAgentPoolId field if non-nil, zero value otherwise.

### GetExecutedAgentPoolIdOk

`func (o *CreateCodingCIJobRequest) GetExecutedAgentPoolIdOk() (*int32, bool)`

GetExecutedAgentPoolIdOk returns a tuple with the ExecutedAgentPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAgentPoolId

`func (o *CreateCodingCIJobRequest) SetExecutedAgentPoolId(v int32)`

SetExecutedAgentPoolId sets ExecutedAgentPoolId field to given value.

### HasExecutedAgentPoolId

`func (o *CreateCodingCIJobRequest) HasExecutedAgentPoolId() bool`

HasExecutedAgentPoolId returns a boolean if a field has been set.

### GetHookType

`func (o *CreateCodingCIJobRequest) GetHookType() string`

GetHookType returns the HookType field if non-nil, zero value otherwise.

### GetHookTypeOk

`func (o *CreateCodingCIJobRequest) GetHookTypeOk() (*string, bool)`

GetHookTypeOk returns a tuple with the HookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookType

`func (o *CreateCodingCIJobRequest) SetHookType(v string)`

SetHookType sets HookType field to given value.


### GetJenkinsFileFromType

`func (o *CreateCodingCIJobRequest) GetJenkinsFileFromType() string`

GetJenkinsFileFromType returns the JenkinsFileFromType field if non-nil, zero value otherwise.

### GetJenkinsFileFromTypeOk

`func (o *CreateCodingCIJobRequest) GetJenkinsFileFromTypeOk() (*string, bool)`

GetJenkinsFileFromTypeOk returns a tuple with the JenkinsFileFromType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJenkinsFileFromType

`func (o *CreateCodingCIJobRequest) SetJenkinsFileFromType(v string)`

SetJenkinsFileFromType sets JenkinsFileFromType field to given value.


### GetJenkinsFilePath

`func (o *CreateCodingCIJobRequest) GetJenkinsFilePath() string`

GetJenkinsFilePath returns the JenkinsFilePath field if non-nil, zero value otherwise.

### GetJenkinsFilePathOk

`func (o *CreateCodingCIJobRequest) GetJenkinsFilePathOk() (*string, bool)`

GetJenkinsFilePathOk returns a tuple with the JenkinsFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJenkinsFilePath

`func (o *CreateCodingCIJobRequest) SetJenkinsFilePath(v string)`

SetJenkinsFilePath sets JenkinsFilePath field to given value.

### HasJenkinsFilePath

`func (o *CreateCodingCIJobRequest) HasJenkinsFilePath() bool`

HasJenkinsFilePath returns a boolean if a field has been set.

### GetJenkinsFileStaticContent

`func (o *CreateCodingCIJobRequest) GetJenkinsFileStaticContent() string`

GetJenkinsFileStaticContent returns the JenkinsFileStaticContent field if non-nil, zero value otherwise.

### GetJenkinsFileStaticContentOk

`func (o *CreateCodingCIJobRequest) GetJenkinsFileStaticContentOk() (*string, bool)`

GetJenkinsFileStaticContentOk returns a tuple with the JenkinsFileStaticContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJenkinsFileStaticContent

`func (o *CreateCodingCIJobRequest) SetJenkinsFileStaticContent(v string)`

SetJenkinsFileStaticContent sets JenkinsFileStaticContent field to given value.

### HasJenkinsFileStaticContent

`func (o *CreateCodingCIJobRequest) HasJenkinsFileStaticContent() bool`

HasJenkinsFileStaticContent returns a boolean if a field has been set.

### GetJobFromType

`func (o *CreateCodingCIJobRequest) GetJobFromType() string`

GetJobFromType returns the JobFromType field if non-nil, zero value otherwise.

### GetJobFromTypeOk

`func (o *CreateCodingCIJobRequest) GetJobFromTypeOk() (*string, bool)`

GetJobFromTypeOk returns a tuple with the JobFromType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFromType

`func (o *CreateCodingCIJobRequest) SetJobFromType(v string)`

SetJobFromType sets JobFromType field to given value.


### GetName

`func (o *CreateCodingCIJobRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCodingCIJobRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCodingCIJobRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *CreateCodingCIJobRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateCodingCIJobRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateCodingCIJobRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetScheduleList

`func (o *CreateCodingCIJobRequest) GetScheduleList() []CodingCIJobSchedule`

GetScheduleList returns the ScheduleList field if non-nil, zero value otherwise.

### GetScheduleListOk

`func (o *CreateCodingCIJobRequest) GetScheduleListOk() (*[]CodingCIJobSchedule, bool)`

GetScheduleListOk returns a tuple with the ScheduleList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleList

`func (o *CreateCodingCIJobRequest) SetScheduleList(v []CodingCIJobSchedule)`

SetScheduleList sets ScheduleList field to given value.

### HasScheduleList

`func (o *CreateCodingCIJobRequest) HasScheduleList() bool`

HasScheduleList returns a boolean if a field has been set.

### GetTriggerMethodList

`func (o *CreateCodingCIJobRequest) GetTriggerMethodList() []string`

GetTriggerMethodList returns the TriggerMethodList field if non-nil, zero value otherwise.

### GetTriggerMethodListOk

`func (o *CreateCodingCIJobRequest) GetTriggerMethodListOk() (*[]string, bool)`

GetTriggerMethodListOk returns a tuple with the TriggerMethodList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMethodList

`func (o *CreateCodingCIJobRequest) SetTriggerMethodList(v []string)`

SetTriggerMethodList sets TriggerMethodList field to given value.

### HasTriggerMethodList

`func (o *CreateCodingCIJobRequest) HasTriggerMethodList() bool`

HasTriggerMethodList returns a boolean if a field has been set.

### GetTriggerRemind

`func (o *CreateCodingCIJobRequest) GetTriggerRemind() string`

GetTriggerRemind returns the TriggerRemind field if non-nil, zero value otherwise.

### GetTriggerRemindOk

`func (o *CreateCodingCIJobRequest) GetTriggerRemindOk() (*string, bool)`

GetTriggerRemindOk returns a tuple with the TriggerRemind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerRemind

`func (o *CreateCodingCIJobRequest) SetTriggerRemind(v string)`

SetTriggerRemind sets TriggerRemind field to given value.

### HasTriggerRemind

`func (o *CreateCodingCIJobRequest) HasTriggerRemind() bool`

HasTriggerRemind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


