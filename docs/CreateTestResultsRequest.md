# CreateTestResultsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseIds** | **[]int32** | 测试用例 ID 数组 | 
**ProjectName** | **string** | 项目名称 | 
**RunId** | **int32** | 测试计划 ID | 
**Status** | **string** | 测试状态：UNTESTED:未测试,PASSED:通过,BLOCKED:阻塞,RETEST:重测,FAILED:失败 | 

## Methods

### NewCreateTestResultsRequest

`func NewCreateTestResultsRequest(caseIds []int32, projectName string, runId int32, status string, ) *CreateTestResultsRequest`

NewCreateTestResultsRequest instantiates a new CreateTestResultsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestResultsRequestWithDefaults

`func NewCreateTestResultsRequestWithDefaults() *CreateTestResultsRequest`

NewCreateTestResultsRequestWithDefaults instantiates a new CreateTestResultsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseIds

`func (o *CreateTestResultsRequest) GetCaseIds() []int32`

GetCaseIds returns the CaseIds field if non-nil, zero value otherwise.

### GetCaseIdsOk

`func (o *CreateTestResultsRequest) GetCaseIdsOk() (*[]int32, bool)`

GetCaseIdsOk returns a tuple with the CaseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseIds

`func (o *CreateTestResultsRequest) SetCaseIds(v []int32)`

SetCaseIds sets CaseIds field to given value.


### GetProjectName

`func (o *CreateTestResultsRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateTestResultsRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateTestResultsRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetRunId

`func (o *CreateTestResultsRequest) GetRunId() int32`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *CreateTestResultsRequest) GetRunIdOk() (*int32, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *CreateTestResultsRequest) SetRunId(v int32)`

SetRunId sets RunId field to given value.


### GetStatus

`func (o *CreateTestResultsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTestResultsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTestResultsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


