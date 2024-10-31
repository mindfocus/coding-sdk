# CodingCIJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlwaysUserIdList** | **[]int32** | 不管构建成功还是失败总是通知的用户 | 
**AutoCancelSameMergeRequest** | **bool** | 自动取消相同 MR | [default to false]
**AutoCancelSameRevision** | **bool** | 自动取消相同版本 | [default to false]
**BranchRegex** | **string** | 分支匹配正则 | [default to ""]
**BranchSelector** | **string** | 触发构建的分支 | [default to ""]
**BuildFailUserIdList** | **[]int32** | 仅构建失败时要通知的用户 | 
**CachePathList** | [**[]CIJobCachePath**](CIJobCachePath.md) | 任务缓存目录配置 | 
**CacheSize** | Pointer to **int32** | 构建缓存大小 | [optional] 
**CreatedAt** | **int32** | 创建时间 | 
**CreatorId** | **int32** | 创建者 | 
**DepotHttpsUrl** | Pointer to **string** | 仓库库的 Https 地址 | [optional] [default to ""]
**DepotId** | **int32** | 仓库ID | 
**DepotName** | Pointer to **string** | 仓库名称 | [optional] [default to ""]
**DepotSshUrl** | Pointer to **string** | 仓库库的 SSH 地址 | [optional] [default to ""]
**DepotType** | **string** | 仓库类型 | [default to ""]
**DepotWebUrl** | Pointer to **string** | 仓库库的 Web 页面 | [optional] [default to ""]
**EnvList** | [**[]CIJobEnv**](CIJobEnv.md) | 环境变量配置 | 
**ExecuteIn** | **string** | 执行方式 | [default to ""]
**ExecutedAgentPoolId** | Pointer to **int32** | 自定义构建节点池 ID | [optional] 
**HookType** | **string** | 代码更新触发匹配规则 | [default to ""]
**Id** | Pointer to **int32** | 构建计划ID | [optional] 
**JenkinsFileFromType** | **string** | Jenkinsfile 来源 | [default to ""]
**JenkinsFilePath** | **string** | Jenkinsfile 在仓库中的文件路径 | [default to ""]
**JenkinsFileStaticContent** | **string** | Jenkinsfile 文件内容 | [default to ""]
**JobFromType** | **string** | 构建计划创建来源 | [default to ""]
**Name** | **string** | 构建计划名称 | [default to ""]
**ProjectId** | **int32** | 项目ID | 
**ProjectName** | Pointer to **string** | 项目名称 | [optional] [default to ""]
**ScheduleList** | [**[]CIJobSchedule**](CIJobSchedule.md) | 针对 CRON triggerMethod 的 schedule 规则配置 | 
**TriggerMethodList** | **[]string** | 构建计划触发方式 | 
**TriggerRemind** | **string** | 构建结果通知触发者机制 | [default to ""]
**UpdatedAt** | **int32** | 最后更新时间 | 

## Methods

### NewCodingCIJob

`func NewCodingCIJob(alwaysUserIdList []int32, autoCancelSameMergeRequest bool, autoCancelSameRevision bool, branchRegex string, branchSelector string, buildFailUserIdList []int32, cachePathList []CIJobCachePath, createdAt int32, creatorId int32, depotId int32, depotType string, envList []CIJobEnv, executeIn string, hookType string, jenkinsFileFromType string, jenkinsFilePath string, jenkinsFileStaticContent string, jobFromType string, name string, projectId int32, scheduleList []CIJobSchedule, triggerMethodList []string, triggerRemind string, updatedAt int32, ) *CodingCIJob`

NewCodingCIJob instantiates a new CodingCIJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodingCIJobWithDefaults

`func NewCodingCIJobWithDefaults() *CodingCIJob`

NewCodingCIJobWithDefaults instantiates a new CodingCIJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlwaysUserIdList

`func (o *CodingCIJob) GetAlwaysUserIdList() []int32`

GetAlwaysUserIdList returns the AlwaysUserIdList field if non-nil, zero value otherwise.

### GetAlwaysUserIdListOk

`func (o *CodingCIJob) GetAlwaysUserIdListOk() (*[]int32, bool)`

GetAlwaysUserIdListOk returns a tuple with the AlwaysUserIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysUserIdList

`func (o *CodingCIJob) SetAlwaysUserIdList(v []int32)`

SetAlwaysUserIdList sets AlwaysUserIdList field to given value.


### GetAutoCancelSameMergeRequest

`func (o *CodingCIJob) GetAutoCancelSameMergeRequest() bool`

GetAutoCancelSameMergeRequest returns the AutoCancelSameMergeRequest field if non-nil, zero value otherwise.

### GetAutoCancelSameMergeRequestOk

`func (o *CodingCIJob) GetAutoCancelSameMergeRequestOk() (*bool, bool)`

GetAutoCancelSameMergeRequestOk returns a tuple with the AutoCancelSameMergeRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCancelSameMergeRequest

`func (o *CodingCIJob) SetAutoCancelSameMergeRequest(v bool)`

SetAutoCancelSameMergeRequest sets AutoCancelSameMergeRequest field to given value.


### GetAutoCancelSameRevision

`func (o *CodingCIJob) GetAutoCancelSameRevision() bool`

GetAutoCancelSameRevision returns the AutoCancelSameRevision field if non-nil, zero value otherwise.

### GetAutoCancelSameRevisionOk

`func (o *CodingCIJob) GetAutoCancelSameRevisionOk() (*bool, bool)`

GetAutoCancelSameRevisionOk returns a tuple with the AutoCancelSameRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCancelSameRevision

`func (o *CodingCIJob) SetAutoCancelSameRevision(v bool)`

SetAutoCancelSameRevision sets AutoCancelSameRevision field to given value.


### GetBranchRegex

`func (o *CodingCIJob) GetBranchRegex() string`

GetBranchRegex returns the BranchRegex field if non-nil, zero value otherwise.

### GetBranchRegexOk

`func (o *CodingCIJob) GetBranchRegexOk() (*string, bool)`

GetBranchRegexOk returns a tuple with the BranchRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchRegex

`func (o *CodingCIJob) SetBranchRegex(v string)`

SetBranchRegex sets BranchRegex field to given value.


### GetBranchSelector

`func (o *CodingCIJob) GetBranchSelector() string`

GetBranchSelector returns the BranchSelector field if non-nil, zero value otherwise.

### GetBranchSelectorOk

`func (o *CodingCIJob) GetBranchSelectorOk() (*string, bool)`

GetBranchSelectorOk returns a tuple with the BranchSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchSelector

`func (o *CodingCIJob) SetBranchSelector(v string)`

SetBranchSelector sets BranchSelector field to given value.


### GetBuildFailUserIdList

`func (o *CodingCIJob) GetBuildFailUserIdList() []int32`

GetBuildFailUserIdList returns the BuildFailUserIdList field if non-nil, zero value otherwise.

### GetBuildFailUserIdListOk

`func (o *CodingCIJob) GetBuildFailUserIdListOk() (*[]int32, bool)`

GetBuildFailUserIdListOk returns a tuple with the BuildFailUserIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildFailUserIdList

`func (o *CodingCIJob) SetBuildFailUserIdList(v []int32)`

SetBuildFailUserIdList sets BuildFailUserIdList field to given value.


### GetCachePathList

`func (o *CodingCIJob) GetCachePathList() []CIJobCachePath`

GetCachePathList returns the CachePathList field if non-nil, zero value otherwise.

### GetCachePathListOk

`func (o *CodingCIJob) GetCachePathListOk() (*[]CIJobCachePath, bool)`

GetCachePathListOk returns a tuple with the CachePathList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePathList

`func (o *CodingCIJob) SetCachePathList(v []CIJobCachePath)`

SetCachePathList sets CachePathList field to given value.


### GetCacheSize

`func (o *CodingCIJob) GetCacheSize() int32`

GetCacheSize returns the CacheSize field if non-nil, zero value otherwise.

### GetCacheSizeOk

`func (o *CodingCIJob) GetCacheSizeOk() (*int32, bool)`

GetCacheSizeOk returns a tuple with the CacheSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSize

`func (o *CodingCIJob) SetCacheSize(v int32)`

SetCacheSize sets CacheSize field to given value.

### HasCacheSize

`func (o *CodingCIJob) HasCacheSize() bool`

HasCacheSize returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CodingCIJob) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CodingCIJob) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CodingCIJob) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatorId

`func (o *CodingCIJob) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *CodingCIJob) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *CodingCIJob) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetDepotHttpsUrl

`func (o *CodingCIJob) GetDepotHttpsUrl() string`

GetDepotHttpsUrl returns the DepotHttpsUrl field if non-nil, zero value otherwise.

### GetDepotHttpsUrlOk

`func (o *CodingCIJob) GetDepotHttpsUrlOk() (*string, bool)`

GetDepotHttpsUrlOk returns a tuple with the DepotHttpsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotHttpsUrl

`func (o *CodingCIJob) SetDepotHttpsUrl(v string)`

SetDepotHttpsUrl sets DepotHttpsUrl field to given value.

### HasDepotHttpsUrl

`func (o *CodingCIJob) HasDepotHttpsUrl() bool`

HasDepotHttpsUrl returns a boolean if a field has been set.

### GetDepotId

`func (o *CodingCIJob) GetDepotId() int32`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CodingCIJob) GetDepotIdOk() (*int32, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CodingCIJob) SetDepotId(v int32)`

SetDepotId sets DepotId field to given value.


### GetDepotName

`func (o *CodingCIJob) GetDepotName() string`

GetDepotName returns the DepotName field if non-nil, zero value otherwise.

### GetDepotNameOk

`func (o *CodingCIJob) GetDepotNameOk() (*string, bool)`

GetDepotNameOk returns a tuple with the DepotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotName

`func (o *CodingCIJob) SetDepotName(v string)`

SetDepotName sets DepotName field to given value.

### HasDepotName

`func (o *CodingCIJob) HasDepotName() bool`

HasDepotName returns a boolean if a field has been set.

### GetDepotSshUrl

`func (o *CodingCIJob) GetDepotSshUrl() string`

GetDepotSshUrl returns the DepotSshUrl field if non-nil, zero value otherwise.

### GetDepotSshUrlOk

`func (o *CodingCIJob) GetDepotSshUrlOk() (*string, bool)`

GetDepotSshUrlOk returns a tuple with the DepotSshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotSshUrl

`func (o *CodingCIJob) SetDepotSshUrl(v string)`

SetDepotSshUrl sets DepotSshUrl field to given value.

### HasDepotSshUrl

`func (o *CodingCIJob) HasDepotSshUrl() bool`

HasDepotSshUrl returns a boolean if a field has been set.

### GetDepotType

`func (o *CodingCIJob) GetDepotType() string`

GetDepotType returns the DepotType field if non-nil, zero value otherwise.

### GetDepotTypeOk

`func (o *CodingCIJob) GetDepotTypeOk() (*string, bool)`

GetDepotTypeOk returns a tuple with the DepotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotType

`func (o *CodingCIJob) SetDepotType(v string)`

SetDepotType sets DepotType field to given value.


### GetDepotWebUrl

`func (o *CodingCIJob) GetDepotWebUrl() string`

GetDepotWebUrl returns the DepotWebUrl field if non-nil, zero value otherwise.

### GetDepotWebUrlOk

`func (o *CodingCIJob) GetDepotWebUrlOk() (*string, bool)`

GetDepotWebUrlOk returns a tuple with the DepotWebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotWebUrl

`func (o *CodingCIJob) SetDepotWebUrl(v string)`

SetDepotWebUrl sets DepotWebUrl field to given value.

### HasDepotWebUrl

`func (o *CodingCIJob) HasDepotWebUrl() bool`

HasDepotWebUrl returns a boolean if a field has been set.

### GetEnvList

`func (o *CodingCIJob) GetEnvList() []CIJobEnv`

GetEnvList returns the EnvList field if non-nil, zero value otherwise.

### GetEnvListOk

`func (o *CodingCIJob) GetEnvListOk() (*[]CIJobEnv, bool)`

GetEnvListOk returns a tuple with the EnvList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvList

`func (o *CodingCIJob) SetEnvList(v []CIJobEnv)`

SetEnvList sets EnvList field to given value.


### GetExecuteIn

`func (o *CodingCIJob) GetExecuteIn() string`

GetExecuteIn returns the ExecuteIn field if non-nil, zero value otherwise.

### GetExecuteInOk

`func (o *CodingCIJob) GetExecuteInOk() (*string, bool)`

GetExecuteInOk returns a tuple with the ExecuteIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteIn

`func (o *CodingCIJob) SetExecuteIn(v string)`

SetExecuteIn sets ExecuteIn field to given value.


### GetExecutedAgentPoolId

`func (o *CodingCIJob) GetExecutedAgentPoolId() int32`

GetExecutedAgentPoolId returns the ExecutedAgentPoolId field if non-nil, zero value otherwise.

### GetExecutedAgentPoolIdOk

`func (o *CodingCIJob) GetExecutedAgentPoolIdOk() (*int32, bool)`

GetExecutedAgentPoolIdOk returns a tuple with the ExecutedAgentPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAgentPoolId

`func (o *CodingCIJob) SetExecutedAgentPoolId(v int32)`

SetExecutedAgentPoolId sets ExecutedAgentPoolId field to given value.

### HasExecutedAgentPoolId

`func (o *CodingCIJob) HasExecutedAgentPoolId() bool`

HasExecutedAgentPoolId returns a boolean if a field has been set.

### GetHookType

`func (o *CodingCIJob) GetHookType() string`

GetHookType returns the HookType field if non-nil, zero value otherwise.

### GetHookTypeOk

`func (o *CodingCIJob) GetHookTypeOk() (*string, bool)`

GetHookTypeOk returns a tuple with the HookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookType

`func (o *CodingCIJob) SetHookType(v string)`

SetHookType sets HookType field to given value.


### GetId

`func (o *CodingCIJob) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CodingCIJob) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CodingCIJob) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CodingCIJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJenkinsFileFromType

`func (o *CodingCIJob) GetJenkinsFileFromType() string`

GetJenkinsFileFromType returns the JenkinsFileFromType field if non-nil, zero value otherwise.

### GetJenkinsFileFromTypeOk

`func (o *CodingCIJob) GetJenkinsFileFromTypeOk() (*string, bool)`

GetJenkinsFileFromTypeOk returns a tuple with the JenkinsFileFromType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJenkinsFileFromType

`func (o *CodingCIJob) SetJenkinsFileFromType(v string)`

SetJenkinsFileFromType sets JenkinsFileFromType field to given value.


### GetJenkinsFilePath

`func (o *CodingCIJob) GetJenkinsFilePath() string`

GetJenkinsFilePath returns the JenkinsFilePath field if non-nil, zero value otherwise.

### GetJenkinsFilePathOk

`func (o *CodingCIJob) GetJenkinsFilePathOk() (*string, bool)`

GetJenkinsFilePathOk returns a tuple with the JenkinsFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJenkinsFilePath

`func (o *CodingCIJob) SetJenkinsFilePath(v string)`

SetJenkinsFilePath sets JenkinsFilePath field to given value.


### GetJenkinsFileStaticContent

`func (o *CodingCIJob) GetJenkinsFileStaticContent() string`

GetJenkinsFileStaticContent returns the JenkinsFileStaticContent field if non-nil, zero value otherwise.

### GetJenkinsFileStaticContentOk

`func (o *CodingCIJob) GetJenkinsFileStaticContentOk() (*string, bool)`

GetJenkinsFileStaticContentOk returns a tuple with the JenkinsFileStaticContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJenkinsFileStaticContent

`func (o *CodingCIJob) SetJenkinsFileStaticContent(v string)`

SetJenkinsFileStaticContent sets JenkinsFileStaticContent field to given value.


### GetJobFromType

`func (o *CodingCIJob) GetJobFromType() string`

GetJobFromType returns the JobFromType field if non-nil, zero value otherwise.

### GetJobFromTypeOk

`func (o *CodingCIJob) GetJobFromTypeOk() (*string, bool)`

GetJobFromTypeOk returns a tuple with the JobFromType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFromType

`func (o *CodingCIJob) SetJobFromType(v string)`

SetJobFromType sets JobFromType field to given value.


### GetName

`func (o *CodingCIJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodingCIJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodingCIJob) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *CodingCIJob) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CodingCIJob) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CodingCIJob) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetProjectName

`func (o *CodingCIJob) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CodingCIJob) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CodingCIJob) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *CodingCIJob) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetScheduleList

`func (o *CodingCIJob) GetScheduleList() []CIJobSchedule`

GetScheduleList returns the ScheduleList field if non-nil, zero value otherwise.

### GetScheduleListOk

`func (o *CodingCIJob) GetScheduleListOk() (*[]CIJobSchedule, bool)`

GetScheduleListOk returns a tuple with the ScheduleList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleList

`func (o *CodingCIJob) SetScheduleList(v []CIJobSchedule)`

SetScheduleList sets ScheduleList field to given value.


### GetTriggerMethodList

`func (o *CodingCIJob) GetTriggerMethodList() []string`

GetTriggerMethodList returns the TriggerMethodList field if non-nil, zero value otherwise.

### GetTriggerMethodListOk

`func (o *CodingCIJob) GetTriggerMethodListOk() (*[]string, bool)`

GetTriggerMethodListOk returns a tuple with the TriggerMethodList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMethodList

`func (o *CodingCIJob) SetTriggerMethodList(v []string)`

SetTriggerMethodList sets TriggerMethodList field to given value.


### GetTriggerRemind

`func (o *CodingCIJob) GetTriggerRemind() string`

GetTriggerRemind returns the TriggerRemind field if non-nil, zero value otherwise.

### GetTriggerRemindOk

`func (o *CodingCIJob) GetTriggerRemindOk() (*string, bool)`

GetTriggerRemindOk returns a tuple with the TriggerRemind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerRemind

`func (o *CodingCIJob) SetTriggerRemind(v string)`

SetTriggerRemind sets TriggerRemind field to given value.


### GetUpdatedAt

`func (o *CodingCIJob) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CodingCIJob) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CodingCIJob) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


