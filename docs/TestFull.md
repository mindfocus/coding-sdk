# TestFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedToId** | Pointer to **NullableInt32** | 处理人 ID | [optional] 
**Case** | Pointer to [**Case**](Case.md) |  | [optional] 
**CaseId** | Pointer to **NullableInt32** | 用例 ID | [optional] 
**Id** | Pointer to **NullableInt32** | 测试任务 ID | [optional] 
**IsCompleted** | Pointer to **NullableBool** | 是否归档 | [optional] [default to false]
**Priority** | Pointer to **NullableInt32** | 优先级 | [optional] 
**SectionId** | Pointer to **NullableInt32** | 分组 ID | [optional] 
**SectionPath** | Pointer to **[]string** | 用例所属分组（按层级由上到下排序） | [optional] 
**Sort** | Pointer to **NullableInt32** | 排序 | [optional] 
**Status** | Pointer to **NullableString** | 状态 | [optional] [default to ""]
**TestedAt** | Pointer to **NullableString** | 测试时间 | [optional] [default to ""]
**TestedBy** | Pointer to **NullableInt32** | 测试人 ID | [optional] 
**Title** | Pointer to **NullableString** | 标题 | [optional] [default to ""]

## Methods

### NewTestFull

`func NewTestFull() *TestFull`

NewTestFull instantiates a new TestFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestFullWithDefaults

`func NewTestFullWithDefaults() *TestFull`

NewTestFullWithDefaults instantiates a new TestFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedToId

`func (o *TestFull) GetAssignedToId() int32`

GetAssignedToId returns the AssignedToId field if non-nil, zero value otherwise.

### GetAssignedToIdOk

`func (o *TestFull) GetAssignedToIdOk() (*int32, bool)`

GetAssignedToIdOk returns a tuple with the AssignedToId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToId

`func (o *TestFull) SetAssignedToId(v int32)`

SetAssignedToId sets AssignedToId field to given value.

### HasAssignedToId

`func (o *TestFull) HasAssignedToId() bool`

HasAssignedToId returns a boolean if a field has been set.

### SetAssignedToIdNil

`func (o *TestFull) SetAssignedToIdNil(b bool)`

 SetAssignedToIdNil sets the value for AssignedToId to be an explicit nil

### UnsetAssignedToId
`func (o *TestFull) UnsetAssignedToId()`

UnsetAssignedToId ensures that no value is present for AssignedToId, not even an explicit nil
### GetCase

`func (o *TestFull) GetCase() Case`

GetCase returns the Case field if non-nil, zero value otherwise.

### GetCaseOk

`func (o *TestFull) GetCaseOk() (*Case, bool)`

GetCaseOk returns a tuple with the Case field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCase

`func (o *TestFull) SetCase(v Case)`

SetCase sets Case field to given value.

### HasCase

`func (o *TestFull) HasCase() bool`

HasCase returns a boolean if a field has been set.

### GetCaseId

`func (o *TestFull) GetCaseId() int32`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *TestFull) GetCaseIdOk() (*int32, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *TestFull) SetCaseId(v int32)`

SetCaseId sets CaseId field to given value.

### HasCaseId

`func (o *TestFull) HasCaseId() bool`

HasCaseId returns a boolean if a field has been set.

### SetCaseIdNil

`func (o *TestFull) SetCaseIdNil(b bool)`

 SetCaseIdNil sets the value for CaseId to be an explicit nil

### UnsetCaseId
`func (o *TestFull) UnsetCaseId()`

UnsetCaseId ensures that no value is present for CaseId, not even an explicit nil
### GetId

`func (o *TestFull) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestFull) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestFull) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TestFull) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TestFull) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TestFull) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsCompleted

`func (o *TestFull) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *TestFull) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *TestFull) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.

### HasIsCompleted

`func (o *TestFull) HasIsCompleted() bool`

HasIsCompleted returns a boolean if a field has been set.

### SetIsCompletedNil

`func (o *TestFull) SetIsCompletedNil(b bool)`

 SetIsCompletedNil sets the value for IsCompleted to be an explicit nil

### UnsetIsCompleted
`func (o *TestFull) UnsetIsCompleted()`

UnsetIsCompleted ensures that no value is present for IsCompleted, not even an explicit nil
### GetPriority

`func (o *TestFull) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TestFull) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TestFull) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TestFull) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *TestFull) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *TestFull) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetSectionId

`func (o *TestFull) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *TestFull) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *TestFull) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *TestFull) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### SetSectionIdNil

`func (o *TestFull) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *TestFull) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetSectionPath

`func (o *TestFull) GetSectionPath() []string`

GetSectionPath returns the SectionPath field if non-nil, zero value otherwise.

### GetSectionPathOk

`func (o *TestFull) GetSectionPathOk() (*[]string, bool)`

GetSectionPathOk returns a tuple with the SectionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionPath

`func (o *TestFull) SetSectionPath(v []string)`

SetSectionPath sets SectionPath field to given value.

### HasSectionPath

`func (o *TestFull) HasSectionPath() bool`

HasSectionPath returns a boolean if a field has been set.

### SetSectionPathNil

`func (o *TestFull) SetSectionPathNil(b bool)`

 SetSectionPathNil sets the value for SectionPath to be an explicit nil

### UnsetSectionPath
`func (o *TestFull) UnsetSectionPath()`

UnsetSectionPath ensures that no value is present for SectionPath, not even an explicit nil
### GetSort

`func (o *TestFull) GetSort() int32`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TestFull) GetSortOk() (*int32, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TestFull) SetSort(v int32)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TestFull) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *TestFull) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *TestFull) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetStatus

`func (o *TestFull) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestFull) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestFull) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestFull) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TestFull) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TestFull) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTestedAt

`func (o *TestFull) GetTestedAt() string`

GetTestedAt returns the TestedAt field if non-nil, zero value otherwise.

### GetTestedAtOk

`func (o *TestFull) GetTestedAtOk() (*string, bool)`

GetTestedAtOk returns a tuple with the TestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestedAt

`func (o *TestFull) SetTestedAt(v string)`

SetTestedAt sets TestedAt field to given value.

### HasTestedAt

`func (o *TestFull) HasTestedAt() bool`

HasTestedAt returns a boolean if a field has been set.

### SetTestedAtNil

`func (o *TestFull) SetTestedAtNil(b bool)`

 SetTestedAtNil sets the value for TestedAt to be an explicit nil

### UnsetTestedAt
`func (o *TestFull) UnsetTestedAt()`

UnsetTestedAt ensures that no value is present for TestedAt, not even an explicit nil
### GetTestedBy

`func (o *TestFull) GetTestedBy() int32`

GetTestedBy returns the TestedBy field if non-nil, zero value otherwise.

### GetTestedByOk

`func (o *TestFull) GetTestedByOk() (*int32, bool)`

GetTestedByOk returns a tuple with the TestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestedBy

`func (o *TestFull) SetTestedBy(v int32)`

SetTestedBy sets TestedBy field to given value.

### HasTestedBy

`func (o *TestFull) HasTestedBy() bool`

HasTestedBy returns a boolean if a field has been set.

### SetTestedByNil

`func (o *TestFull) SetTestedByNil(b bool)`

 SetTestedByNil sets the value for TestedBy to be an explicit nil

### UnsetTestedBy
`func (o *TestFull) UnsetTestedBy()`

UnsetTestedBy ensures that no value is present for TestedBy, not even an explicit nil
### GetTitle

`func (o *TestFull) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TestFull) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TestFull) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TestFull) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *TestFull) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *TestFull) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


