# ModifyCodingCIJobRequest

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
**DepotType** | **string** | 仓库类型 CODING,TGIT,GITHUB,GITLAB,GITLAB_PRIVATE,GITEE,NONE | 
**EnvList** | Pointer to [**[]CodingCIJobEnv**](CodingCIJobEnv.md) | 环境变量配置 | [optional] 
**ExecuteIn** | **string** | 执行方式 CVM | STATIC | 
**ExecutedAgentPoolId** | Pointer to **int32** | 自定义构建节点池 ID，ExecuteIn 为 AGENT 必填 | [optional] 
**HookType** | **string** | 代码更新触发匹配规则 | 
**Id** | **int32** | 构建计划 ID | 
**JenkinsFileFromType** | **string** | STATIC，SCM 从代码库读取 | 
**JenkinsFilePath** | Pointer to **string** | JenkinsFileFromType 为 SCM 必填 | [optional] 
**JenkinsFileStaticContent** | Pointer to **string** | JenkinsFileFromType 为 STATIC 必填 | [optional] 
**JobFromType** | **string** | 构建计划来源 TKE TCB | 
**Name** | Pointer to **string** | 构建计划名称 | [optional] 
**ProjectId** | **int32** | 项目 ID | 
**ScheduleList** | Pointer to [**[]CodingCIJobSchedule**](CodingCIJobSchedule.md) | 针对 CRON triggerMethod 的 schedule 规则配置, 暂只用于添加 | [optional] 
**TriggerMethodList** | Pointer to **[]string** | REF_CHANGE 代码更新触发      CRON &#x3D; 1 定时触发      MR_CHANGE  MR变动触发  TKE 对接他们传空数组老是有问题，遂改成非必填 | [optional] 
**TriggerRemind** | **string** | 构建结果通知触发者机制  ALWAYS -总是通知;  BUILD_FAIL -仅构建失败时通知; | 

## Methods

### NewModifyCodingCIJobRequest

`func NewModifyCodingCIJobRequest(autoCancelSameMergeRequest bool, autoCancelSameRevision bool, depotId int32, depotType string, executeIn string, hookType string, id int32, jenkinsFileFromType string, jobFromType string, projectId int32, triggerRemind string, ) *ModifyCodingCIJobRequest`

NewModifyCodingCIJobRequest instantiates a new ModifyCodingCIJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyCodingCIJobRequestWithDefaults

`func NewModifyCodingCIJobRequestWithDefaults() *ModifyCodingCIJobRequest`

NewModifyCodingCIJobRequestWithDefaults instantiates a new ModifyCodingCIJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlwaysUserIdList

`func (o *ModifyCodingCIJobRequest) GetAlwaysUserIdList() []int32`

GetAlwaysUserIdList returns the AlwaysUserIdList field if non-nil, zero value otherwise.

### GetAlwaysUserIdListOk

`func (o *ModifyCodingCIJobRequest) GetAlwaysUserIdListOk() (*[]int32, bool)`

GetAlwaysUserIdListOk returns a tuple with the AlwaysUserIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysUserIdList

`func (o *ModifyCodingCIJobRequest) SetAlwaysUserIdList(v []int32)`

SetAlwaysUserIdList sets AlwaysUserIdList field to given value.

### HasAlwaysUserIdList

`func (o *ModifyCodingCIJobRequest) HasAlwaysUserIdList() bool`

HasAlwaysUserIdList returns a boolean if a field has been set.

### GetAutoCancelSameMergeRequest

`func (o *ModifyCodingCIJobRequest) GetAutoCancelSameMergeRequest() bool`

GetAutoCancelSameMergeRequest returns the AutoCancelSameMergeRequest field if non-nil, zero value otherwise.

### GetAutoCancelSameMergeRequestOk

`func (o *ModifyCodingCIJobRequest) GetAutoCancelSameMergeRequestOk() (*bool, bool)`

GetAutoCancelSameMergeRequestOk returns a tuple with the AutoCancelSameMergeRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCancelSameMergeRequest

`func (o *ModifyCodingCIJobRequest) SetAutoCancelSameMergeRequest(v bool)`

SetAutoCancelSameMergeRequest sets AutoCancelSameMergeRequest field to given value.


### GetAutoCancelSameRevision

`func (o *ModifyCodingCIJobRequest) GetAutoCancelSameRevision() bool`

GetAutoCancelSameRevision returns the AutoCancelSameRevision field if non-nil, zero value otherwise.

### GetAutoCancelSameRevisionOk

`func (o *ModifyCodingCIJobRequest) GetAutoCancelSameRevisionOk() (*bool, bool)`

GetAutoCancelSameRevisionOk returns a tuple with the AutoCancelSameRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCancelSameRevision

`func (o *ModifyCodingCIJobRequest) SetAutoCancelSameRevision(v bool)`

SetAutoCancelSameRevision sets AutoCancelSameRevision field to given value.


### GetBranchRegex

`func (o *ModifyCodingCIJobRequest) GetBranchRegex() string`

GetBranchRegex returns the BranchRegex field if non-nil, zero value otherwise.

### GetBranchRegexOk

`func (o *ModifyCodingCIJobRequest) GetBranchRegexOk() (*string, bool)`

GetBranchRegexOk returns a tuple with the BranchRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchRegex

`func (o *ModifyCodingCIJobRequest) SetBranchRegex(v string)`

SetBranchRegex sets BranchRegex field to given value.

### HasBranchRegex

`func (o *ModifyCodingCIJobRequest) HasBranchRegex() bool`

HasBranchRegex returns a boolean if a field has been set.

### GetBranchSelector

`func (o *ModifyCodingCIJobRequest) GetBranchSelector() string`

GetBranchSelector returns the BranchSelector field if non-nil, zero value otherwise.

### GetBranchSelectorOk

`func (o *ModifyCodingCIJobRequest) GetBranchSelectorOk() (*string, bool)`

GetBranchSelectorOk returns a tuple with the BranchSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchSelector

`func (o *ModifyCodingCIJobRequest) SetBranchSelector(v string)`

SetBranchSelector sets BranchSelector field to given value.

### HasBranchSelector

`func (o *ModifyCodingCIJobRequest) HasBranchSelector() bool`

HasBranchSelector returns a boolean if a field has been set.

### GetBuildFailUserIdList

`func (o *ModifyCodingCIJobRequest) GetBuildFailUserIdList() []int32`

GetBuildFailUserIdList returns the BuildFailUserIdList field if non-nil, zero value otherwise.

### GetBuildFailUserIdListOk

`func (o *ModifyCodingCIJobRequest) GetBuildFailUserIdListOk() (*[]int32, bool)`

GetBuildFailUserIdListOk returns a tuple with the BuildFailUserIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildFailUserIdList

`func (o *ModifyCodingCIJobRequest) SetBuildFailUserIdList(v []int32)`

SetBuildFailUserIdList sets BuildFailUserIdList field to given value.

### HasBuildFailUserIdList

`func (o *ModifyCodingCIJobRequest) HasBuildFailUserIdList() bool`

HasBuildFailUserIdList returns a boolean if a field has been set.

### GetCachePathList

`func (o *ModifyCodingCIJobRequest) GetCachePathList() []CodingCIJobCachePath`

GetCachePathList returns the CachePathList field if non-nil, zero value otherwise.

### GetCachePathListOk

`func (o *ModifyCodingCIJobRequest) GetCachePathListOk() (*[]CodingCIJobCachePath, bool)`

GetCachePathListOk returns a tuple with the CachePathList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePathList

`func (o *ModifyCodingCIJobRequest) SetCachePathList(v []CodingCIJobCachePath)`

SetCachePathList sets CachePathList field to given value.

### HasCachePathList

`func (o *ModifyCodingCIJobRequest) HasCachePathList() bool`

HasCachePathList returns a boolean if a field has been set.

### GetDepotId

`func (o *ModifyCodingCIJobRequest) GetDepotId() int32`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ModifyCodingCIJobRequest) GetDepotIdOk() (*int32, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ModifyCodingCIJobRequest) SetDepotId(v int32)`

SetDepotId sets DepotId field to given value.


### GetDepotType

`func (o *ModifyCodingCIJobRequest) GetDepotType() string`

GetDepotType returns the DepotType field if non-nil, zero value otherwise.

### GetDepotTypeOk

`func (o *ModifyCodingCIJobRequest) GetDepotTypeOk() (*string, bool)`

GetDepotTypeOk returns a tuple with the DepotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotType

`func (o *ModifyCodingCIJobRequest) SetDepotType(v string)`

SetDepotType sets DepotType field to given value.


### GetEnvList

`func (o *ModifyCodingCIJobRequest) GetEnvList() []CodingCIJobEnv`

GetEnvList returns the EnvList field if non-nil, zero value otherwise.

### GetEnvListOk

`func (o *ModifyCodingCIJobRequest) GetEnvListOk() (*[]CodingCIJobEnv, bool)`

GetEnvListOk returns a tuple with the EnvList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvList

`func (o *ModifyCodingCIJobRequest) SetEnvList(v []CodingCIJobEnv)`

SetEnvList sets EnvList field to given value.

### HasEnvList

`func (o *ModifyCodingCIJobRequest) HasEnvList() bool`

HasEnvList returns a boolean if a field has been set.

### GetExecuteIn

`func (o *ModifyCodingCIJobRequest) GetExecuteIn() string`

GetExecuteIn returns the ExecuteIn field if non-nil, zero value otherwise.

### GetExecuteInOk

`func (o *ModifyCodingCIJobRequest) GetExecuteInOk() (*string, bool)`

GetExecuteInOk returns a tuple with the ExecuteIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteIn

`func (o *ModifyCodingCIJobRequest) SetExecuteIn(v string)`

SetExecuteIn sets ExecuteIn field to given value.


### GetExecutedAgentPoolId

`func (o *ModifyCodingCIJobRequest) GetExecutedAgentPoolId() int32`

GetExecutedAgentPoolId returns the ExecutedAgentPoolId field if non-nil, zero value otherwise.

### GetExecutedAgentPoolIdOk

`func (o *ModifyCodingCIJobRequest) GetExecutedAgentPoolIdOk() (*int32, bool)`

GetExecutedAgentPoolIdOk returns a tuple with the ExecutedAgentPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAgentPoolId

`func (o *ModifyCodingCIJobRequest) SetExecutedAgentPoolId(v int32)`

SetExecutedAgentPoolId sets ExecutedAgentPoolId field to given value.

### HasExecutedAgentPoolId

`func (o *ModifyCodingCIJobRequest) HasExecutedAgentPoolId() bool`

HasExecutedAgentPoolId returns a boolean if a field has been set.

### GetHookType

`func (o *ModifyCodingCIJobRequest) GetHookType() string`

GetHookType returns the HookType field if non-nil, zero value otherwise.

### GetHookTypeOk

`func (o *ModifyCodingCIJobRequest) GetHookTypeOk() (*string, bool)`

GetHookTypeOk returns a tuple with the HookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookType

`func (o *ModifyCodingCIJobRequest) SetHookType(v string)`

SetHookType sets HookType field to given value.


### GetId

`func (o *ModifyCodingCIJobRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModifyCodingCIJobRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModifyCodingCIJobRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetJenkinsFileFromType

`func (o *ModifyCodingCIJobRequest) GetJenkinsFileFromType() string`

GetJenkinsFileFromType returns the JenkinsFileFromType field if non-nil, zero value otherwise.

### GetJenkinsFileFromTypeOk

`func (o *ModifyCodingCIJobRequest) GetJenkinsFileFromTypeOk() (*string, bool)`

GetJenkinsFileFromTypeOk returns a tuple with the JenkinsFileFromType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJenkinsFileFromType

`func (o *ModifyCodingCIJobRequest) SetJenkinsFileFromType(v string)`

SetJenkinsFileFromType sets JenkinsFileFromType field to given value.


### GetJenkinsFilePath

`func (o *ModifyCodingCIJobRequest) GetJenkinsFilePath() string`

GetJenkinsFilePath returns the JenkinsFilePath field if non-nil, zero value otherwise.

### GetJenkinsFilePathOk

`func (o *ModifyCodingCIJobRequest) GetJenkinsFilePathOk() (*string, bool)`

GetJenkinsFilePathOk returns a tuple with the JenkinsFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJenkinsFilePath

`func (o *ModifyCodingCIJobRequest) SetJenkinsFilePath(v string)`

SetJenkinsFilePath sets JenkinsFilePath field to given value.

### HasJenkinsFilePath

`func (o *ModifyCodingCIJobRequest) HasJenkinsFilePath() bool`

HasJenkinsFilePath returns a boolean if a field has been set.

### GetJenkinsFileStaticContent

`func (o *ModifyCodingCIJobRequest) GetJenkinsFileStaticContent() string`

GetJenkinsFileStaticContent returns the JenkinsFileStaticContent field if non-nil, zero value otherwise.

### GetJenkinsFileStaticContentOk

`func (o *ModifyCodingCIJobRequest) GetJenkinsFileStaticContentOk() (*string, bool)`

GetJenkinsFileStaticContentOk returns a tuple with the JenkinsFileStaticContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJenkinsFileStaticContent

`func (o *ModifyCodingCIJobRequest) SetJenkinsFileStaticContent(v string)`

SetJenkinsFileStaticContent sets JenkinsFileStaticContent field to given value.

### HasJenkinsFileStaticContent

`func (o *ModifyCodingCIJobRequest) HasJenkinsFileStaticContent() bool`

HasJenkinsFileStaticContent returns a boolean if a field has been set.

### GetJobFromType

`func (o *ModifyCodingCIJobRequest) GetJobFromType() string`

GetJobFromType returns the JobFromType field if non-nil, zero value otherwise.

### GetJobFromTypeOk

`func (o *ModifyCodingCIJobRequest) GetJobFromTypeOk() (*string, bool)`

GetJobFromTypeOk returns a tuple with the JobFromType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFromType

`func (o *ModifyCodingCIJobRequest) SetJobFromType(v string)`

SetJobFromType sets JobFromType field to given value.


### GetName

`func (o *ModifyCodingCIJobRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyCodingCIJobRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyCodingCIJobRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModifyCodingCIJobRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *ModifyCodingCIJobRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModifyCodingCIJobRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModifyCodingCIJobRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetScheduleList

`func (o *ModifyCodingCIJobRequest) GetScheduleList() []CodingCIJobSchedule`

GetScheduleList returns the ScheduleList field if non-nil, zero value otherwise.

### GetScheduleListOk

`func (o *ModifyCodingCIJobRequest) GetScheduleListOk() (*[]CodingCIJobSchedule, bool)`

GetScheduleListOk returns a tuple with the ScheduleList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleList

`func (o *ModifyCodingCIJobRequest) SetScheduleList(v []CodingCIJobSchedule)`

SetScheduleList sets ScheduleList field to given value.

### HasScheduleList

`func (o *ModifyCodingCIJobRequest) HasScheduleList() bool`

HasScheduleList returns a boolean if a field has been set.

### GetTriggerMethodList

`func (o *ModifyCodingCIJobRequest) GetTriggerMethodList() []string`

GetTriggerMethodList returns the TriggerMethodList field if non-nil, zero value otherwise.

### GetTriggerMethodListOk

`func (o *ModifyCodingCIJobRequest) GetTriggerMethodListOk() (*[]string, bool)`

GetTriggerMethodListOk returns a tuple with the TriggerMethodList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMethodList

`func (o *ModifyCodingCIJobRequest) SetTriggerMethodList(v []string)`

SetTriggerMethodList sets TriggerMethodList field to given value.

### HasTriggerMethodList

`func (o *ModifyCodingCIJobRequest) HasTriggerMethodList() bool`

HasTriggerMethodList returns a boolean if a field has been set.

### GetTriggerRemind

`func (o *ModifyCodingCIJobRequest) GetTriggerRemind() string`

GetTriggerRemind returns the TriggerRemind field if non-nil, zero value otherwise.

### GetTriggerRemindOk

`func (o *ModifyCodingCIJobRequest) GetTriggerRemindOk() (*string, bool)`

GetTriggerRemindOk returns a tuple with the TriggerRemind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerRemind

`func (o *ModifyCodingCIJobRequest) SetTriggerRemind(v string)`

SetTriggerRemind sets TriggerRemind field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


