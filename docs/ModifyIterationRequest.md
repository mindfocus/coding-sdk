# ModifyIterationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **int32** | 处理人 ID | [optional] 
**EndAt** | Pointer to **map[string]interface{}** | 结束时间，格式：2020-01-01 | [optional] 
**Goal** | Pointer to **string** | 目标 | [optional] 
**IterationCode** | **int64** | 迭代编号 | 
**Name** | **string** | 标题 | 
**ProjectName** | **string** | 项目名称 | 
**StartAt** | Pointer to **map[string]interface{}** | 开始时间，格式：2020-01-01 | [optional] 

## Methods

### NewModifyIterationRequest

`func NewModifyIterationRequest(iterationCode int64, name string, projectName string, ) *ModifyIterationRequest`

NewModifyIterationRequest instantiates a new ModifyIterationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyIterationRequestWithDefaults

`func NewModifyIterationRequestWithDefaults() *ModifyIterationRequest`

NewModifyIterationRequestWithDefaults instantiates a new ModifyIterationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *ModifyIterationRequest) GetAssignee() int32`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *ModifyIterationRequest) GetAssigneeOk() (*int32, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *ModifyIterationRequest) SetAssignee(v int32)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *ModifyIterationRequest) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetEndAt

`func (o *ModifyIterationRequest) GetEndAt() map[string]interface{}`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *ModifyIterationRequest) GetEndAtOk() (*map[string]interface{}, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *ModifyIterationRequest) SetEndAt(v map[string]interface{})`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *ModifyIterationRequest) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### GetGoal

`func (o *ModifyIterationRequest) GetGoal() string`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *ModifyIterationRequest) GetGoalOk() (*string, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *ModifyIterationRequest) SetGoal(v string)`

SetGoal sets Goal field to given value.

### HasGoal

`func (o *ModifyIterationRequest) HasGoal() bool`

HasGoal returns a boolean if a field has been set.

### GetIterationCode

`func (o *ModifyIterationRequest) GetIterationCode() int64`

GetIterationCode returns the IterationCode field if non-nil, zero value otherwise.

### GetIterationCodeOk

`func (o *ModifyIterationRequest) GetIterationCodeOk() (*int64, bool)`

GetIterationCodeOk returns a tuple with the IterationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCode

`func (o *ModifyIterationRequest) SetIterationCode(v int64)`

SetIterationCode sets IterationCode field to given value.


### GetName

`func (o *ModifyIterationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyIterationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyIterationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProjectName

`func (o *ModifyIterationRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ModifyIterationRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ModifyIterationRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetStartAt

`func (o *ModifyIterationRequest) GetStartAt() map[string]interface{}`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *ModifyIterationRequest) GetStartAtOk() (*map[string]interface{}, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *ModifyIterationRequest) SetStartAt(v map[string]interface{})`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *ModifyIterationRequest) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


