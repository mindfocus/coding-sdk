# PlanIterationIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueCode** | **[]int64** | 需要规划的事项编号列表 | 
**IterationCode** | **int64** | 迭代编号，将事项移出迭代请传 0 | 
**ProjectName** | **string** | 项目名称 | 
**SingleMode** | Pointer to **bool** | 单个事项移动模式，即未规划的子事项的迭代跟随传入的事项一同规划，否则仅规划传入的事项 | [optional] 

## Methods

### NewPlanIterationIssueRequest

`func NewPlanIterationIssueRequest(issueCode []int64, iterationCode int64, projectName string, ) *PlanIterationIssueRequest`

NewPlanIterationIssueRequest instantiates a new PlanIterationIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanIterationIssueRequestWithDefaults

`func NewPlanIterationIssueRequestWithDefaults() *PlanIterationIssueRequest`

NewPlanIterationIssueRequestWithDefaults instantiates a new PlanIterationIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueCode

`func (o *PlanIterationIssueRequest) GetIssueCode() []int64`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *PlanIterationIssueRequest) GetIssueCodeOk() (*[]int64, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *PlanIterationIssueRequest) SetIssueCode(v []int64)`

SetIssueCode sets IssueCode field to given value.


### GetIterationCode

`func (o *PlanIterationIssueRequest) GetIterationCode() int64`

GetIterationCode returns the IterationCode field if non-nil, zero value otherwise.

### GetIterationCodeOk

`func (o *PlanIterationIssueRequest) GetIterationCodeOk() (*int64, bool)`

GetIterationCodeOk returns a tuple with the IterationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCode

`func (o *PlanIterationIssueRequest) SetIterationCode(v int64)`

SetIterationCode sets IterationCode field to given value.


### GetProjectName

`func (o *PlanIterationIssueRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *PlanIterationIssueRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *PlanIterationIssueRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetSingleMode

`func (o *PlanIterationIssueRequest) GetSingleMode() bool`

GetSingleMode returns the SingleMode field if non-nil, zero value otherwise.

### GetSingleModeOk

`func (o *PlanIterationIssueRequest) GetSingleModeOk() (*bool, bool)`

GetSingleModeOk returns a tuple with the SingleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleMode

`func (o *PlanIterationIssueRequest) SetSingleMode(v bool)`

SetSingleMode sets SingleMode field to given value.

### HasSingleMode

`func (o *PlanIterationIssueRequest) HasSingleMode() bool`

HasSingleMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


