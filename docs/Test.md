# Test

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedToId** | Pointer to **NullableInt32** | 处理人 ID | [optional] 
**CaseId** | Pointer to **NullableInt32** | 用例 ID | [optional] 
**Id** | Pointer to **NullableInt32** | 测试任务 ID | [optional] 
**IsCompleted** | Pointer to **NullableBool** | 是否归档 | [optional] [default to false]
**Priority** | Pointer to **NullableInt32** | 优先级 | [optional] 
**SectionId** | Pointer to **NullableInt32** | 分组 ID | [optional] 
**Sort** | Pointer to **NullableInt32** | 排序 | [optional] 
**Status** | Pointer to **NullableString** | 状态 | [optional] [default to ""]
**TestedAt** | Pointer to **NullableString** | 测试时间 | [optional] [default to ""]
**Title** | Pointer to **NullableString** | 标题 | [optional] [default to ""]

## Methods

### NewTest

`func NewTest() *Test`

NewTest instantiates a new Test object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestWithDefaults

`func NewTestWithDefaults() *Test`

NewTestWithDefaults instantiates a new Test object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedToId

`func (o *Test) GetAssignedToId() int32`

GetAssignedToId returns the AssignedToId field if non-nil, zero value otherwise.

### GetAssignedToIdOk

`func (o *Test) GetAssignedToIdOk() (*int32, bool)`

GetAssignedToIdOk returns a tuple with the AssignedToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToId

`func (o *Test) SetAssignedToId(v int32)`

SetAssignedToId sets AssignedToId field to given value.

### HasAssignedToId

`func (o *Test) HasAssignedToId() bool`

HasAssignedToId returns a boolean if a field has been set.

### SetAssignedToIdNil

`func (o *Test) SetAssignedToIdNil(b bool)`

 SetAssignedToIdNil sets the value for AssignedToId to be an explicit nil

### UnsetAssignedToId
`func (o *Test) UnsetAssignedToId()`

UnsetAssignedToId ensures that no value is present for AssignedToId, not even an explicit nil
### GetCaseId

`func (o *Test) GetCaseId() int32`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *Test) GetCaseIdOk() (*int32, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *Test) SetCaseId(v int32)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *Test) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### SetCaseIdNil

`func (o *Test) SetCaseIdNil(b bool)`

 SetCaseIdNil sets the value for CaseId to be an explicit nil

### UnsetCaseId
`func (o *Test) UnsetCaseId()`

UnsetCaseId ensures that no value is present for CaseId, not even an explicit nil
### GetId

`func (o *Test) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Test) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Test) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Test) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Test) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Test) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsCompleted

`func (o *Test) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *Test) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *Test) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.

### HasIsCompleted

`func (o *Test) HasIsCompleted() bool`

HasIsCompleted returns a boolean if a field has been set.

### SetIsCompletedNil

`func (o *Test) SetIsCompletedNil(b bool)`

 SetIsCompletedNil sets the value for IsCompleted to be an explicit nil

### UnsetIsCompleted
`func (o *Test) UnsetIsCompleted()`

UnsetIsCompleted ensures that no value is present for IsCompleted, not even an explicit nil
### GetPriority

`func (o *Test) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Test) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Test) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Test) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *Test) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *Test) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetSectionId

`func (o *Test) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *Test) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *Test) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *Test) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### SetSectionIdNil

`func (o *Test) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *Test) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetSort

`func (o *Test) GetSort() int32`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *Test) GetSortOk() (*int32, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *Test) SetSort(v int32)`

SetSort sets Sort field to given value.

### HasSort

`func (o *Test) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *Test) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *Test) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetStatus

`func (o *Test) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Test) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Test) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Test) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Test) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Test) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTestedAt

`func (o *Test) GetTestedAt() string`

GetTestedAt returns the TestedAt field if non-nil, zero value otherwise.

### GetTestedAtOk

`func (o *Test) GetTestedAtOk() (*string, bool)`

GetTestedAtOk returns a tuple with the TestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestedAt

`func (o *Test) SetTestedAt(v string)`

SetTestedAt sets TestedAt field to given value.

### HasTestedAt

`func (o *Test) HasTestedAt() bool`

HasTestedAt returns a boolean if a field has been set.

### SetTestedAtNil

`func (o *Test) SetTestedAtNil(b bool)`

 SetTestedAtNil sets the value for TestedAt to be an explicit nil

### UnsetTestedAt
`func (o *Test) UnsetTestedAt()`

UnsetTestedAt ensures that no value is present for TestedAt, not even an explicit nil
### GetTitle

`func (o *Test) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Test) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Test) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Test) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Test) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Test) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


