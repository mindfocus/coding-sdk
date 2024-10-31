# DescribeTestRunListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecuteType** | Pointer to **int32** | 执行方式: 1-手动执行 2-自动化流水线执行 | [optional] 
**GitReleaseState** | Pointer to **int32** | 发布版本状态：0-未发布 1-已发布（与参数IterationId、IterationStatus、SectionId互斥） | [optional] 
**IsCompleted** | Pointer to **bool** | 是否已经归档 | [optional] 
**IterationId** | Pointer to **[]int32** | 迭代 ID（与参数IterationStatus、GitReleaseState、SectionId互斥） | [optional] 
**IterationStatus** | Pointer to **[]string** | 迭代状态: WAIT_PROCESS、PROCESSING、COMPLETED（与参数IterationId、GitReleaseState、SectionId互斥） | [optional] 
**Keyword** | Pointer to **string** | 计划名称 | [optional] 
**ProjectName** | **string** | 项目名称 | 
**SectionId** | Pointer to **int64** | 分组 ID（与参数IterationId、IterationStatus、GitReleaseState互斥） | [optional] 
**State** | Pointer to **int32** | 状态: 0-未开始 1-进行中 2-已测完 | [optional] 

## Methods

### NewDescribeTestRunListRequest

`func NewDescribeTestRunListRequest(projectName string, ) *DescribeTestRunListRequest`

NewDescribeTestRunListRequest instantiates a new DescribeTestRunListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeTestRunListRequestWithDefaults

`func NewDescribeTestRunListRequestWithDefaults() *DescribeTestRunListRequest`

NewDescribeTestRunListRequestWithDefaults instantiates a new DescribeTestRunListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecuteType

`func (o *DescribeTestRunListRequest) GetExecuteType() int32`

GetExecuteType returns the ExecuteType field if non-nil, zero value otherwise.

### GetExecuteTypeOk

`func (o *DescribeTestRunListRequest) GetExecuteTypeOk() (*int32, bool)`

GetExecuteTypeOk returns a tuple with the ExecuteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecuteType

`func (o *DescribeTestRunListRequest) SetExecuteType(v int32)`

SetExecuteType sets ExecuteType field to given value.

### HasExecuteType

`func (o *DescribeTestRunListRequest) HasExecuteType() bool`

HasExecuteType returns a boolean if a field has been set.

### GetGitReleaseState

`func (o *DescribeTestRunListRequest) GetGitReleaseState() int32`

GetGitReleaseState returns the GitReleaseState field if non-nil, zero value otherwise.

### GetGitReleaseStateOk

`func (o *DescribeTestRunListRequest) GetGitReleaseStateOk() (*int32, bool)`

GetGitReleaseStateOk returns a tuple with the GitReleaseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitReleaseState

`func (o *DescribeTestRunListRequest) SetGitReleaseState(v int32)`

SetGitReleaseState sets GitReleaseState field to given value.

### HasGitReleaseState

`func (o *DescribeTestRunListRequest) HasGitReleaseState() bool`

HasGitReleaseState returns a boolean if a field has been set.

### GetIsCompleted

`func (o *DescribeTestRunListRequest) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *DescribeTestRunListRequest) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *DescribeTestRunListRequest) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.

### HasIsCompleted

`func (o *DescribeTestRunListRequest) HasIsCompleted() bool`

HasIsCompleted returns a boolean if a field has been set.

### GetIterationId

`func (o *DescribeTestRunListRequest) GetIterationId() []int32`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *DescribeTestRunListRequest) GetIterationIdOk() (*[]int32, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *DescribeTestRunListRequest) SetIterationId(v []int32)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *DescribeTestRunListRequest) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.

### GetIterationStatus

`func (o *DescribeTestRunListRequest) GetIterationStatus() []string`

GetIterationStatus returns the IterationStatus field if non-nil, zero value otherwise.

### GetIterationStatusOk

`func (o *DescribeTestRunListRequest) GetIterationStatusOk() (*[]string, bool)`

GetIterationStatusOk returns a tuple with the IterationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationStatus

`func (o *DescribeTestRunListRequest) SetIterationStatus(v []string)`

SetIterationStatus sets IterationStatus field to given value.

### HasIterationStatus

`func (o *DescribeTestRunListRequest) HasIterationStatus() bool`

HasIterationStatus returns a boolean if a field has been set.

### GetKeyword

`func (o *DescribeTestRunListRequest) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *DescribeTestRunListRequest) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *DescribeTestRunListRequest) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *DescribeTestRunListRequest) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetProjectName

`func (o *DescribeTestRunListRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeTestRunListRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeTestRunListRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetSectionId

`func (o *DescribeTestRunListRequest) GetSectionId() int64`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *DescribeTestRunListRequest) GetSectionIdOk() (*int64, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *DescribeTestRunListRequest) SetSectionId(v int64)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *DescribeTestRunListRequest) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### GetState

`func (o *DescribeTestRunListRequest) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DescribeTestRunListRequest) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DescribeTestRunListRequest) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *DescribeTestRunListRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


