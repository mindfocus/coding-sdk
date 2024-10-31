# CodingCIBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | 分支 | [default to ""]
**Cause** | **string** | 触发原因 | [default to ""]
**CodingCIId** | **string** | 构建唯一标识 | [default to ""]
**CommitId** | **string** | Git Commit ID | [default to ""]
**CreatedAt** | **int32** | 构建创建时间 | 
**Duration** | **int32** | 构建执行时间 QUEUED  等待构建 INITIALIZING  初始化 NOT_BUILT  准备构建 RUNNING  运行中 SUCCEED  成功 FAILED  失败 ABORTED  被取消 TIMEOUT  超时 | 
**EnvList** | Pointer to [**[]CIJobEnv**](CIJobEnv.md) | 构建使用的环境变量 | [optional] 
**FailedMessage** | **string** | 失败原因 | [default to ""]
**Id** | **int32** | 构建 ID | 
**JenkinsFileContent** | **string** | 本次构建的 Jenkinsfile | [default to ""]
**JobId** | **int32** | 构建计划 ID | 
**NodeObtainedAt** | Pointer to **NullableInt64** | 获取到执行机的时间，如果为负数表示还未获取到构建节点 | [optional] 
**Number** | **int32** | 构建序号 | 
**StartedAt** | Pointer to **NullableInt64** | 开始构建时间，如果为负数就是默认值表示还未开始 | [optional] 
**Status** | **string** | 构建当前状态 | [default to ""]
**StatusNode** | **string** | 构建进行到了哪个 stage/node | [default to ""]
**TestResult** | [**CIBuildTestResult**](CIBuildTestResult.md) |  | 

## Methods

### NewCodingCIBuild

`func NewCodingCIBuild(branch string, cause string, codingCIId string, commitId string, createdAt int32, duration int32, failedMessage string, id int32, jenkinsFileContent string, jobId int32, number int32, status string, statusNode string, testResult CIBuildTestResult, ) *CodingCIBuild`

NewCodingCIBuild instantiates a new CodingCIBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodingCIBuildWithDefaults

`func NewCodingCIBuildWithDefaults() *CodingCIBuild`

NewCodingCIBuildWithDefaults instantiates a new CodingCIBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *CodingCIBuild) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CodingCIBuild) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CodingCIBuild) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetCause

`func (o *CodingCIBuild) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *CodingCIBuild) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *CodingCIBuild) SetCause(v string)`

SetCause sets Cause field to given value.


### GetCodingCIId

`func (o *CodingCIBuild) GetCodingCIId() string`

GetCodingCIId returns the CodingCIId field if non-nil, zero value otherwise.

### GetCodingCIIdOk

`func (o *CodingCIBuild) GetCodingCIIdOk() (*string, bool)`

GetCodingCIIdOk returns a tuple with the CodingCIId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodingCIId

`func (o *CodingCIBuild) SetCodingCIId(v string)`

SetCodingCIId sets CodingCIId field to given value.


### GetCommitId

`func (o *CodingCIBuild) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *CodingCIBuild) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *CodingCIBuild) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### GetCreatedAt

`func (o *CodingCIBuild) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CodingCIBuild) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CodingCIBuild) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetDuration

`func (o *CodingCIBuild) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CodingCIBuild) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CodingCIBuild) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetEnvList

`func (o *CodingCIBuild) GetEnvList() []CIJobEnv`

GetEnvList returns the EnvList field if non-nil, zero value otherwise.

### GetEnvListOk

`func (o *CodingCIBuild) GetEnvListOk() (*[]CIJobEnv, bool)`

GetEnvListOk returns a tuple with the EnvList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvList

`func (o *CodingCIBuild) SetEnvList(v []CIJobEnv)`

SetEnvList sets EnvList field to given value.

### HasEnvList

`func (o *CodingCIBuild) HasEnvList() bool`

HasEnvList returns a boolean if a field has been set.

### SetEnvListNil

`func (o *CodingCIBuild) SetEnvListNil(b bool)`

 SetEnvListNil sets the value for EnvList to be an explicit nil

### UnsetEnvList
`func (o *CodingCIBuild) UnsetEnvList()`

UnsetEnvList ensures that no value is present for EnvList, not even an explicit nil
### GetFailedMessage

`func (o *CodingCIBuild) GetFailedMessage() string`

GetFailedMessage returns the FailedMessage field if non-nil, zero value otherwise.

### GetFailedMessageOk

`func (o *CodingCIBuild) GetFailedMessageOk() (*string, bool)`

GetFailedMessageOk returns a tuple with the FailedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMessage

`func (o *CodingCIBuild) SetFailedMessage(v string)`

SetFailedMessage sets FailedMessage field to given value.


### GetId

`func (o *CodingCIBuild) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CodingCIBuild) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CodingCIBuild) SetId(v int32)`

SetId sets Id field to given value.


### GetJenkinsFileContent

`func (o *CodingCIBuild) GetJenkinsFileContent() string`

GetJenkinsFileContent returns the JenkinsFileContent field if non-nil, zero value otherwise.

### GetJenkinsFileContentOk

`func (o *CodingCIBuild) GetJenkinsFileContentOk() (*string, bool)`

GetJenkinsFileContentOk returns a tuple with the JenkinsFileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJenkinsFileContent

`func (o *CodingCIBuild) SetJenkinsFileContent(v string)`

SetJenkinsFileContent sets JenkinsFileContent field to given value.


### GetJobId

`func (o *CodingCIBuild) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *CodingCIBuild) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *CodingCIBuild) SetJobId(v int32)`

SetJobId sets JobId field to given value.


### GetNodeObtainedAt

`func (o *CodingCIBuild) GetNodeObtainedAt() int64`

GetNodeObtainedAt returns the NodeObtainedAt field if non-nil, zero value otherwise.

### GetNodeObtainedAtOk

`func (o *CodingCIBuild) GetNodeObtainedAtOk() (*int64, bool)`

GetNodeObtainedAtOk returns a tuple with the NodeObtainedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeObtainedAt

`func (o *CodingCIBuild) SetNodeObtainedAt(v int64)`

SetNodeObtainedAt sets NodeObtainedAt field to given value.

### HasNodeObtainedAt

`func (o *CodingCIBuild) HasNodeObtainedAt() bool`

HasNodeObtainedAt returns a boolean if a field has been set.

### SetNodeObtainedAtNil

`func (o *CodingCIBuild) SetNodeObtainedAtNil(b bool)`

 SetNodeObtainedAtNil sets the value for NodeObtainedAt to be an explicit nil

### UnsetNodeObtainedAt
`func (o *CodingCIBuild) UnsetNodeObtainedAt()`

UnsetNodeObtainedAt ensures that no value is present for NodeObtainedAt, not even an explicit nil
### GetNumber

`func (o *CodingCIBuild) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CodingCIBuild) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CodingCIBuild) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetStartedAt

`func (o *CodingCIBuild) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CodingCIBuild) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CodingCIBuild) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CodingCIBuild) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *CodingCIBuild) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *CodingCIBuild) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetStatus

`func (o *CodingCIBuild) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CodingCIBuild) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CodingCIBuild) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusNode

`func (o *CodingCIBuild) GetStatusNode() string`

GetStatusNode returns the StatusNode field if non-nil, zero value otherwise.

### GetStatusNodeOk

`func (o *CodingCIBuild) GetStatusNodeOk() (*string, bool)`

GetStatusNodeOk returns a tuple with the StatusNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusNode

`func (o *CodingCIBuild) SetStatusNode(v string)`

SetStatusNode sets StatusNode field to given value.


### GetTestResult

`func (o *CodingCIBuild) GetTestResult() CIBuildTestResult`

GetTestResult returns the TestResult field if non-nil, zero value otherwise.

### GetTestResultOk

`func (o *CodingCIBuild) GetTestResultOk() (*CIBuildTestResult, bool)`

GetTestResultOk returns a tuple with the TestResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResult

`func (o *CodingCIBuild) SetTestResult(v CIBuildTestResult)`

SetTestResult sets TestResult field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


