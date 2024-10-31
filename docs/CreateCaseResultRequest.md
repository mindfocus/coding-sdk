# CreateCaseResultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseId** | **int32** | 测试任务 ID | 
**CustomStepStatus** | Pointer to **[]string** | 每一步的测试结果（步骤用例时需要本参数） | [optional] 
**ProjectName** | **string** | 项目名称 | 
**RunId** | **int32** | 测试计划 ID | 
**Status** | **string** | 该任务的测试结果，可选值：UNTESTED:未测试,PASSED:通过,BLOCKED:阻塞,RETEST:重测,FAILED:失败 | 

## Methods

### NewCreateCaseResultRequest

`func NewCreateCaseResultRequest(caseId int32, projectName string, runId int32, status string, ) *CreateCaseResultRequest`

NewCreateCaseResultRequest instantiates a new CreateCaseResultRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCaseResultRequestWithDefaults

`func NewCreateCaseResultRequestWithDefaults() *CreateCaseResultRequest`

NewCreateCaseResultRequestWithDefaults instantiates a new CreateCaseResultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseId

`func (o *CreateCaseResultRequest) GetCaseId() int32`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *CreateCaseResultRequest) GetCaseIdOk() (*int32, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *CreateCaseResultRequest) SetCaseId(v int32)`

SetCaseId sets CaseId field to given value.


### GetCustomStepStatus

`func (o *CreateCaseResultRequest) GetCustomStepStatus() []string`

GetCustomStepStatus returns the CustomStepStatus field if non-nil, zero value otherwise.

### GetCustomStepStatusOk

`func (o *CreateCaseResultRequest) GetCustomStepStatusOk() (*[]string, bool)`

GetCustomStepStatusOk returns a tuple with the CustomStepStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStepStatus

`func (o *CreateCaseResultRequest) SetCustomStepStatus(v []string)`

SetCustomStepStatus sets CustomStepStatus field to given value.

### HasCustomStepStatus

`func (o *CreateCaseResultRequest) HasCustomStepStatus() bool`

HasCustomStepStatus returns a boolean if a field has been set.

### GetProjectName

`func (o *CreateCaseResultRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateCaseResultRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateCaseResultRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetRunId

`func (o *CreateCaseResultRequest) GetRunId() int32`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *CreateCaseResultRequest) GetRunIdOk() (*int32, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *CreateCaseResultRequest) SetRunId(v int32)`

SetRunId sets RunId field to given value.


### GetStatus

`func (o *CreateCaseResultRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateCaseResultRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateCaseResultRequest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


