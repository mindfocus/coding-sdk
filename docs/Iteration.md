# Iteration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **int64** | 处理人 ID ，为 0 代表没有设置 | [optional] 
**Code** | Pointer to **int64** | 迭代编号，项目内唯一 | [optional] 
**CompletedCount** | Pointer to **int64** | 迭代中完成事项总数 | [optional] 
**CompletedPercent** | Pointer to **float32** | 迭代中事项完成比率 | [optional] 
**Completer** | Pointer to **int64** | 完成人 ID | [optional] 
**CreatedAt** | Pointer to **NullableInt64** | 创建时间 | [optional] 
**Creator** | Pointer to **int64** | 创建人 ID | [optional] 
**EndAt** | Pointer to **NullableInt64** | 结束时间，时间戳，-28800000 代表没有设置 | [optional] 
**Goal** | Pointer to **NullableString** | 迭代目标 | [optional] [default to ""]
**Name** | Pointer to **string** | 标题 | [optional] [default to ""]
**ProcessingCount** | Pointer to **int64** | 迭代中进行中事项总数 | [optional] 
**StartAt** | Pointer to **NullableInt64** | 开始时间，时间戳，-28800000 代表没有设置 | [optional] 
**Starter** | Pointer to **int64** | 开始人 ID | [optional] 
**Status** | Pointer to **string** | 迭代状态：WAIT_PROCESS,PROCESSING,COMPLETED | [optional] [default to ""]
**UpdatedAt** | Pointer to **NullableInt64** | 修改时间 | [optional] 
**WaitProcessCount** | Pointer to **int64** | 迭代中待处理事项总数 | [optional] 

## Methods

### NewIteration

`func NewIteration() *Iteration`

NewIteration instantiates a new Iteration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIterationWithDefaults

`func NewIterationWithDefaults() *Iteration`

NewIterationWithDefaults instantiates a new Iteration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *Iteration) GetAssignee() int64`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *Iteration) GetAssigneeOk() (*int64, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *Iteration) SetAssignee(v int64)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *Iteration) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCode

`func (o *Iteration) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Iteration) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Iteration) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *Iteration) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCompletedCount

`func (o *Iteration) GetCompletedCount() int64`

GetCompletedCount returns the CompletedCount field if non-nil, zero value otherwise.

### GetCompletedCountOk

`func (o *Iteration) GetCompletedCountOk() (*int64, bool)`

GetCompletedCountOk returns a tuple with the CompletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedCount

`func (o *Iteration) SetCompletedCount(v int64)`

SetCompletedCount sets CompletedCount field to given value.

### HasCompletedCount

`func (o *Iteration) HasCompletedCount() bool`

HasCompletedCount returns a boolean if a field has been set.

### GetCompletedPercent

`func (o *Iteration) GetCompletedPercent() float32`

GetCompletedPercent returns the CompletedPercent field if non-nil, zero value otherwise.

### GetCompletedPercentOk

`func (o *Iteration) GetCompletedPercentOk() (*float32, bool)`

GetCompletedPercentOk returns a tuple with the CompletedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedPercent

`func (o *Iteration) SetCompletedPercent(v float32)`

SetCompletedPercent sets CompletedPercent field to given value.

### HasCompletedPercent

`func (o *Iteration) HasCompletedPercent() bool`

HasCompletedPercent returns a boolean if a field has been set.

### GetCompleter

`func (o *Iteration) GetCompleter() int64`

GetCompleter returns the Completer field if non-nil, zero value otherwise.

### GetCompleterOk

`func (o *Iteration) GetCompleterOk() (*int64, bool)`

GetCompleterOk returns a tuple with the Completer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleter

`func (o *Iteration) SetCompleter(v int64)`

SetCompleter sets Completer field to given value.

### HasCompleter

`func (o *Iteration) HasCompleter() bool`

HasCompleter returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Iteration) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Iteration) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Iteration) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Iteration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Iteration) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Iteration) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreator

`func (o *Iteration) GetCreator() int64`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *Iteration) GetCreatorOk() (*int64, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *Iteration) SetCreator(v int64)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *Iteration) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetEndAt

`func (o *Iteration) GetEndAt() int64`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *Iteration) GetEndAtOk() (*int64, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *Iteration) SetEndAt(v int64)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *Iteration) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### SetEndAtNil

`func (o *Iteration) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *Iteration) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
### GetGoal

`func (o *Iteration) GetGoal() string`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *Iteration) GetGoalOk() (*string, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *Iteration) SetGoal(v string)`

SetGoal sets Goal field to given value.

### HasGoal

`func (o *Iteration) HasGoal() bool`

HasGoal returns a boolean if a field has been set.

### SetGoalNil

`func (o *Iteration) SetGoalNil(b bool)`

 SetGoalNil sets the value for Goal to be an explicit nil

### UnsetGoal
`func (o *Iteration) UnsetGoal()`

UnsetGoal ensures that no value is present for Goal, not even an explicit nil
### GetName

`func (o *Iteration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Iteration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Iteration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Iteration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessingCount

`func (o *Iteration) GetProcessingCount() int64`

GetProcessingCount returns the ProcessingCount field if non-nil, zero value otherwise.

### GetProcessingCountOk

`func (o *Iteration) GetProcessingCountOk() (*int64, bool)`

GetProcessingCountOk returns a tuple with the ProcessingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingCount

`func (o *Iteration) SetProcessingCount(v int64)`

SetProcessingCount sets ProcessingCount field to given value.

### HasProcessingCount

`func (o *Iteration) HasProcessingCount() bool`

HasProcessingCount returns a boolean if a field has been set.

### GetStartAt

`func (o *Iteration) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *Iteration) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *Iteration) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *Iteration) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### SetStartAtNil

`func (o *Iteration) SetStartAtNil(b bool)`

 SetStartAtNil sets the value for StartAt to be an explicit nil

### UnsetStartAt
`func (o *Iteration) UnsetStartAt()`

UnsetStartAt ensures that no value is present for StartAt, not even an explicit nil
### GetStarter

`func (o *Iteration) GetStarter() int64`

GetStarter returns the Starter field if non-nil, zero value otherwise.

### GetStarterOk

`func (o *Iteration) GetStarterOk() (*int64, bool)`

GetStarterOk returns a tuple with the Starter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarter

`func (o *Iteration) SetStarter(v int64)`

SetStarter sets Starter field to given value.

### HasStarter

`func (o *Iteration) HasStarter() bool`

HasStarter returns a boolean if a field has been set.

### GetStatus

`func (o *Iteration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Iteration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Iteration) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Iteration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Iteration) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Iteration) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Iteration) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Iteration) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Iteration) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Iteration) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetWaitProcessCount

`func (o *Iteration) GetWaitProcessCount() int64`

GetWaitProcessCount returns the WaitProcessCount field if non-nil, zero value otherwise.

### GetWaitProcessCountOk

`func (o *Iteration) GetWaitProcessCountOk() (*int64, bool)`

GetWaitProcessCountOk returns a tuple with the WaitProcessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitProcessCount

`func (o *Iteration) SetWaitProcessCount(v int64)`

SetWaitProcessCount sets WaitProcessCount field to given value.

### HasWaitProcessCount

`func (o *Iteration) HasWaitProcessCount() bool`

HasWaitProcessCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


