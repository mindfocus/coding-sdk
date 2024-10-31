# OpenApiRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **int64** | 处理人 | [optional] 
**Code** | Pointer to **int64** | 版本code | [optional] 
**CompletedCount** | Pointer to **int64** | 已完成事项数目  | [optional] 
**CompletedPercent** | Pointer to **NullableFloat32** | 事项完成率  | [optional] 
**CreatedAt** | Pointer to **int64** | 创建时间 | [optional] 
**Creator** | Pointer to **int64** | 创建人 | [optional] 
**Description** | Pointer to **NullableString** | 描述 | [optional] [default to ""]
**Id** | Pointer to **int64** | 版本ID | [optional] 
**Iterations** | Pointer to [**[]IterationSimple**](IterationSimple.md) | 版本关联迭代列表  | [optional] 
**Name** | Pointer to **string** | 版本名称 | [optional] [default to ""]
**ProcessingCount** | Pointer to **int64** | 处理中事项数目  | [optional] 
**ProjectId** | Pointer to **int64** | 项目ID | [optional] 
**ReleaseDate** | Pointer to **int64** | 发哺乳期 | [optional] 
**StartDate** | Pointer to **int64** | 开始日期 | [optional] 
**Status** | Pointer to **string** | 状态 | [optional] [default to ""]
**UpdatedAt** | Pointer to **int64** | 更新时间 | [optional] 
**WaitProcessCount** | Pointer to **int64** | 待处理事项数目  | [optional] 

## Methods

### NewOpenApiRelease

`func NewOpenApiRelease() *OpenApiRelease`

NewOpenApiRelease instantiates a new OpenApiRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiReleaseWithDefaults

`func NewOpenApiReleaseWithDefaults() *OpenApiRelease`

NewOpenApiReleaseWithDefaults instantiates a new OpenApiRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *OpenApiRelease) GetAssignee() int64`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *OpenApiRelease) GetAssigneeOk() (*int64, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *OpenApiRelease) SetAssignee(v int64)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *OpenApiRelease) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCode

`func (o *OpenApiRelease) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OpenApiRelease) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OpenApiRelease) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *OpenApiRelease) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCompletedCount

`func (o *OpenApiRelease) GetCompletedCount() int64`

GetCompletedCount returns the CompletedCount field if non-nil, zero value otherwise.

### GetCompletedCountOk

`func (o *OpenApiRelease) GetCompletedCountOk() (*int64, bool)`

GetCompletedCountOk returns a tuple with the CompletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedCount

`func (o *OpenApiRelease) SetCompletedCount(v int64)`

SetCompletedCount sets CompletedCount field to given value.

### HasCompletedCount

`func (o *OpenApiRelease) HasCompletedCount() bool`

HasCompletedCount returns a boolean if a field has been set.

### GetCompletedPercent

`func (o *OpenApiRelease) GetCompletedPercent() float32`

GetCompletedPercent returns the CompletedPercent field if non-nil, zero value otherwise.

### GetCompletedPercentOk

`func (o *OpenApiRelease) GetCompletedPercentOk() (*float32, bool)`

GetCompletedPercentOk returns a tuple with the CompletedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedPercent

`func (o *OpenApiRelease) SetCompletedPercent(v float32)`

SetCompletedPercent sets CompletedPercent field to given value.

### HasCompletedPercent

`func (o *OpenApiRelease) HasCompletedPercent() bool`

HasCompletedPercent returns a boolean if a field has been set.

### SetCompletedPercentNil

`func (o *OpenApiRelease) SetCompletedPercentNil(b bool)`

 SetCompletedPercentNil sets the value for CompletedPercent to be an explicit nil

### UnsetCompletedPercent
`func (o *OpenApiRelease) UnsetCompletedPercent()`

UnsetCompletedPercent ensures that no value is present for CompletedPercent, not even an explicit nil
### GetCreatedAt

`func (o *OpenApiRelease) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpenApiRelease) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpenApiRelease) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OpenApiRelease) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *OpenApiRelease) GetCreator() int64`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *OpenApiRelease) GetCreatorOk() (*int64, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *OpenApiRelease) SetCreator(v int64)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *OpenApiRelease) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *OpenApiRelease) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenApiRelease) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenApiRelease) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpenApiRelease) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OpenApiRelease) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OpenApiRelease) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *OpenApiRelease) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenApiRelease) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenApiRelease) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OpenApiRelease) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIterations

`func (o *OpenApiRelease) GetIterations() []IterationSimple`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *OpenApiRelease) GetIterationsOk() (*[]IterationSimple, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *OpenApiRelease) SetIterations(v []IterationSimple)`

SetIterations sets Iterations field to given value.

### HasIterations

`func (o *OpenApiRelease) HasIterations() bool`

HasIterations returns a boolean if a field has been set.

### GetName

`func (o *OpenApiRelease) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenApiRelease) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenApiRelease) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenApiRelease) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProcessingCount

`func (o *OpenApiRelease) GetProcessingCount() int64`

GetProcessingCount returns the ProcessingCount field if non-nil, zero value otherwise.

### GetProcessingCountOk

`func (o *OpenApiRelease) GetProcessingCountOk() (*int64, bool)`

GetProcessingCountOk returns a tuple with the ProcessingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingCount

`func (o *OpenApiRelease) SetProcessingCount(v int64)`

SetProcessingCount sets ProcessingCount field to given value.

### HasProcessingCount

`func (o *OpenApiRelease) HasProcessingCount() bool`

HasProcessingCount returns a boolean if a field has been set.

### GetProjectId

`func (o *OpenApiRelease) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *OpenApiRelease) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *OpenApiRelease) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *OpenApiRelease) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReleaseDate

`func (o *OpenApiRelease) GetReleaseDate() int64`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *OpenApiRelease) GetReleaseDateOk() (*int64, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *OpenApiRelease) SetReleaseDate(v int64)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *OpenApiRelease) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetStartDate

`func (o *OpenApiRelease) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *OpenApiRelease) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *OpenApiRelease) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *OpenApiRelease) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *OpenApiRelease) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OpenApiRelease) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OpenApiRelease) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OpenApiRelease) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OpenApiRelease) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OpenApiRelease) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OpenApiRelease) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OpenApiRelease) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWaitProcessCount

`func (o *OpenApiRelease) GetWaitProcessCount() int64`

GetWaitProcessCount returns the WaitProcessCount field if non-nil, zero value otherwise.

### GetWaitProcessCountOk

`func (o *OpenApiRelease) GetWaitProcessCountOk() (*int64, bool)`

GetWaitProcessCountOk returns a tuple with the WaitProcessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitProcessCount

`func (o *OpenApiRelease) SetWaitProcessCount(v int64)`

SetWaitProcessCount sets WaitProcessCount field to given value.

### HasWaitProcessCount

`func (o *OpenApiRelease) HasWaitProcessCount() bool`

HasWaitProcessCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


