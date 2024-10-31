# CreateTestResultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomStepStatus** | Pointer to **[]string** | 每一步的测试结果（步骤用例时需要本参数） | [optional] 
**ProjectName** | **string** | 项目名称 | 
**Status** | **string** | 该任务的测试结果，可选值：UNTESTED:未测试,PASSED:通过,BLOCKED:阻塞,RETEST:重测,FAILED:失败 | 
**TestId** | **int32** | 测试任务 ID | 

## Methods

### NewCreateTestResultRequest

`func NewCreateTestResultRequest(projectName string, status string, testId int32, ) *CreateTestResultRequest`

NewCreateTestResultRequest instantiates a new CreateTestResultRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestResultRequestWithDefaults

`func NewCreateTestResultRequestWithDefaults() *CreateTestResultRequest`

NewCreateTestResultRequestWithDefaults instantiates a new CreateTestResultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomStepStatus

`func (o *CreateTestResultRequest) GetCustomStepStatus() []string`

GetCustomStepStatus returns the CustomStepStatus field if non-nil, zero value otherwise.

### GetCustomStepStatusOk

`func (o *CreateTestResultRequest) GetCustomStepStatusOk() (*[]string, bool)`

GetCustomStepStatusOk returns a tuple with the CustomStepStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStepStatus

`func (o *CreateTestResultRequest) SetCustomStepStatus(v []string)`

SetCustomStepStatus sets CustomStepStatus field to given value.

### HasCustomStepStatus

`func (o *CreateTestResultRequest) HasCustomStepStatus() bool`

HasCustomStepStatus returns a boolean if a field has been set.

### GetProjectName

`func (o *CreateTestResultRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateTestResultRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateTestResultRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetStatus

`func (o *CreateTestResultRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTestResultRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTestResultRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTestId

`func (o *CreateTestResultRequest) GetTestId() int32`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *CreateTestResultRequest) GetTestIdOk() (*int32, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *CreateTestResultRequest) SetTestId(v int32)`

SetTestId sets TestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


