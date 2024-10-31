# CreateTestStepResultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actual** | Pointer to **string** | 该步骤的实际测试结果 | [optional] 
**ProjectName** | **string** | 项目名称 | 
**Status** | Pointer to **string** | 该任务的测试结果，可选值：UNTESTED:未测试,PASSED:通过,BLOCKED:阻塞,RETEST:重测,FAILED:失败 | [optional] 
**StepIndex** | **int32** | 步骤的索引顺序，起始值为 1 | 
**StepStatus** | **string** | 该步骤的测试结果，可选值：PASSED:通过,FAILED:失败 | 
**TestId** | **int32** | 测试任务 ID | 

## Methods

### NewCreateTestStepResultRequest

`func NewCreateTestStepResultRequest(projectName string, stepIndex int32, stepStatus string, testId int32, ) *CreateTestStepResultRequest`

NewCreateTestStepResultRequest instantiates a new CreateTestStepResultRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestStepResultRequestWithDefaults

`func NewCreateTestStepResultRequestWithDefaults() *CreateTestStepResultRequest`

NewCreateTestStepResultRequestWithDefaults instantiates a new CreateTestStepResultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActual

`func (o *CreateTestStepResultRequest) GetActual() string`

GetActual returns the Actual field if non-nil, zero value otherwise.

### GetActualOk

`func (o *CreateTestStepResultRequest) GetActualOk() (*string, bool)`

GetActualOk returns a tuple with the Actual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActual

`func (o *CreateTestStepResultRequest) SetActual(v string)`

SetActual sets Actual field to given value.

### HasActual

`func (o *CreateTestStepResultRequest) HasActual() bool`

HasActual returns a boolean if a field has been set.

### GetProjectName

`func (o *CreateTestStepResultRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateTestStepResultRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateTestStepResultRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetStatus

`func (o *CreateTestStepResultRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTestStepResultRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTestStepResultRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateTestStepResultRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStepIndex

`func (o *CreateTestStepResultRequest) GetStepIndex() int32`

GetStepIndex returns the StepIndex field if non-nil, zero value otherwise.

### GetStepIndexOk

`func (o *CreateTestStepResultRequest) GetStepIndexOk() (*int32, bool)`

GetStepIndexOk returns a tuple with the StepIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepIndex

`func (o *CreateTestStepResultRequest) SetStepIndex(v int32)`

SetStepIndex sets StepIndex field to given value.


### GetStepStatus

`func (o *CreateTestStepResultRequest) GetStepStatus() string`

GetStepStatus returns the StepStatus field if non-nil, zero value otherwise.

### GetStepStatusOk

`func (o *CreateTestStepResultRequest) GetStepStatusOk() (*string, bool)`

GetStepStatusOk returns a tuple with the StepStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepStatus

`func (o *CreateTestStepResultRequest) SetStepStatus(v string)`

SetStepStatus sets StepStatus field to given value.


### GetTestId

`func (o *CreateTestStepResultRequest) GetTestId() int32`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *CreateTestStepResultRequest) GetTestIdOk() (*int32, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *CreateTestStepResultRequest) SetTestId(v int32)`

SetTestId sets TestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


