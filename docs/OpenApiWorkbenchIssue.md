# OpenApiWorkbenchIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to [**User**](User.md) |  | [optional] 
**ChildCompletedCount** | Pointer to **NullableInt64** | 子 | [optional] 
**ChildCount** | Pointer to **NullableInt64** | 子 | [optional] 
**Code** | Pointer to **NullableInt64** | 事项code | [optional] 
**CreatedAt** | Pointer to **NullableInt64** | 创建时间 | [optional] 
**Creator** | Pointer to [**User**](User.md) |  | [optional] 
**Description** | Pointer to **NullableBool** | 描述 | [optional] [default to false]
**DueDate** | Pointer to **NullableInt64** | 截止时间 | [optional] 
**Epic** | Pointer to [**Epic**](Epic.md) |  | [optional] 
**IssueStatus** | Pointer to [**IssueStatus**](IssueStatus.md) |  | [optional] 
**IssueType** | Pointer to [**IssueTypeDetail**](IssueTypeDetail.md) |  | [optional] 
**Iteration** | Pointer to [**IterationSimple**](IterationSimple.md) |  | [optional] 
**Name** | Pointer to **NullableString** | 事项名称 | [optional] [default to ""]
**ParentCode** | Pointer to **NullableInt64** | 父事项code | [optional] 
**ParentIssue** | Pointer to [**IssueSimpleData**](IssueSimpleData.md) |  | [optional] 
**Priority** | Pointer to **NullableString** | 优先级 | [optional] [default to ""]
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**ProjectId** | Pointer to **NullableInt64** | 项目ID | [optional] 
**SubTasks** | Pointer to [**[]SubTask**](SubTask.md) | 子工作项 | [optional] 
**Type** | Pointer to **NullableString** | 事项类型 | [optional] [default to ""]
**UpdatedAt** | Pointer to **NullableInt64** | 更新时间 | [optional] 

## Methods

### NewOpenApiWorkbenchIssue

`func NewOpenApiWorkbenchIssue() *OpenApiWorkbenchIssue`

NewOpenApiWorkbenchIssue instantiates a new OpenApiWorkbenchIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiWorkbenchIssueWithDefaults

`func NewOpenApiWorkbenchIssueWithDefaults() *OpenApiWorkbenchIssue`

NewOpenApiWorkbenchIssueWithDefaults instantiates a new OpenApiWorkbenchIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *OpenApiWorkbenchIssue) GetAssignee() User`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *OpenApiWorkbenchIssue) GetAssigneeOk() (*User, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *OpenApiWorkbenchIssue) SetAssignee(v User)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *OpenApiWorkbenchIssue) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetChildCompletedCount

`func (o *OpenApiWorkbenchIssue) GetChildCompletedCount() int64`

GetChildCompletedCount returns the ChildCompletedCount field if non-nil, zero value otherwise.

### GetChildCompletedCountOk

`func (o *OpenApiWorkbenchIssue) GetChildCompletedCountOk() (*int64, bool)`

GetChildCompletedCountOk returns a tuple with the ChildCompletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCompletedCount

`func (o *OpenApiWorkbenchIssue) SetChildCompletedCount(v int64)`

SetChildCompletedCount sets ChildCompletedCount field to given value.

### HasChildCompletedCount

`func (o *OpenApiWorkbenchIssue) HasChildCompletedCount() bool`

HasChildCompletedCount returns a boolean if a field has been set.

### SetChildCompletedCountNil

`func (o *OpenApiWorkbenchIssue) SetChildCompletedCountNil(b bool)`

 SetChildCompletedCountNil sets the value for ChildCompletedCount to be an explicit nil

### UnsetChildCompletedCount
`func (o *OpenApiWorkbenchIssue) UnsetChildCompletedCount()`

UnsetChildCompletedCount ensures that no value is present for ChildCompletedCount, not even an explicit nil
### GetChildCount

`func (o *OpenApiWorkbenchIssue) GetChildCount() int64`

GetChildCount returns the ChildCount field if non-nil, zero value otherwise.

### GetChildCountOk

`func (o *OpenApiWorkbenchIssue) GetChildCountOk() (*int64, bool)`

GetChildCountOk returns a tuple with the ChildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCount

`func (o *OpenApiWorkbenchIssue) SetChildCount(v int64)`

SetChildCount sets ChildCount field to given value.

### HasChildCount

`func (o *OpenApiWorkbenchIssue) HasChildCount() bool`

HasChildCount returns a boolean if a field has been set.

### SetChildCountNil

`func (o *OpenApiWorkbenchIssue) SetChildCountNil(b bool)`

 SetChildCountNil sets the value for ChildCount to be an explicit nil

### UnsetChildCount
`func (o *OpenApiWorkbenchIssue) UnsetChildCount()`

UnsetChildCount ensures that no value is present for ChildCount, not even an explicit nil
### GetCode

`func (o *OpenApiWorkbenchIssue) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OpenApiWorkbenchIssue) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OpenApiWorkbenchIssue) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *OpenApiWorkbenchIssue) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *OpenApiWorkbenchIssue) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *OpenApiWorkbenchIssue) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCreatedAt

`func (o *OpenApiWorkbenchIssue) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OpenApiWorkbenchIssue) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OpenApiWorkbenchIssue) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OpenApiWorkbenchIssue) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *OpenApiWorkbenchIssue) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *OpenApiWorkbenchIssue) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreator

`func (o *OpenApiWorkbenchIssue) GetCreator() User`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *OpenApiWorkbenchIssue) GetCreatorOk() (*User, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *OpenApiWorkbenchIssue) SetCreator(v User)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *OpenApiWorkbenchIssue) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDescription

`func (o *OpenApiWorkbenchIssue) GetDescription() bool`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenApiWorkbenchIssue) GetDescriptionOk() (*bool, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenApiWorkbenchIssue) SetDescription(v bool)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpenApiWorkbenchIssue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OpenApiWorkbenchIssue) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OpenApiWorkbenchIssue) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDueDate

`func (o *OpenApiWorkbenchIssue) GetDueDate() int64`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *OpenApiWorkbenchIssue) GetDueDateOk() (*int64, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *OpenApiWorkbenchIssue) SetDueDate(v int64)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *OpenApiWorkbenchIssue) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *OpenApiWorkbenchIssue) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *OpenApiWorkbenchIssue) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetEpic

`func (o *OpenApiWorkbenchIssue) GetEpic() Epic`

GetEpic returns the Epic field if non-nil, zero value otherwise.

### GetEpicOk

`func (o *OpenApiWorkbenchIssue) GetEpicOk() (*Epic, bool)`

GetEpicOk returns a tuple with the Epic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpic

`func (o *OpenApiWorkbenchIssue) SetEpic(v Epic)`

SetEpic sets Epic field to given value.

### HasEpic

`func (o *OpenApiWorkbenchIssue) HasEpic() bool`

HasEpic returns a boolean if a field has been set.

### GetIssueStatus

`func (o *OpenApiWorkbenchIssue) GetIssueStatus() IssueStatus`

GetIssueStatus returns the IssueStatus field if non-nil, zero value otherwise.

### GetIssueStatusOk

`func (o *OpenApiWorkbenchIssue) GetIssueStatusOk() (*IssueStatus, bool)`

GetIssueStatusOk returns a tuple with the IssueStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueStatus

`func (o *OpenApiWorkbenchIssue) SetIssueStatus(v IssueStatus)`

SetIssueStatus sets IssueStatus field to given value.

### HasIssueStatus

`func (o *OpenApiWorkbenchIssue) HasIssueStatus() bool`

HasIssueStatus returns a boolean if a field has been set.

### GetIssueType

`func (o *OpenApiWorkbenchIssue) GetIssueType() IssueTypeDetail`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *OpenApiWorkbenchIssue) GetIssueTypeOk() (*IssueTypeDetail, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *OpenApiWorkbenchIssue) SetIssueType(v IssueTypeDetail)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *OpenApiWorkbenchIssue) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetIteration

`func (o *OpenApiWorkbenchIssue) GetIteration() IterationSimple`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *OpenApiWorkbenchIssue) GetIterationOk() (*IterationSimple, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *OpenApiWorkbenchIssue) SetIteration(v IterationSimple)`

SetIteration sets Iteration field to given value.

### HasIteration

`func (o *OpenApiWorkbenchIssue) HasIteration() bool`

HasIteration returns a boolean if a field has been set.

### GetName

`func (o *OpenApiWorkbenchIssue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenApiWorkbenchIssue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenApiWorkbenchIssue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OpenApiWorkbenchIssue) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OpenApiWorkbenchIssue) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OpenApiWorkbenchIssue) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetParentCode

`func (o *OpenApiWorkbenchIssue) GetParentCode() int64`

GetParentCode returns the ParentCode field if non-nil, zero value otherwise.

### GetParentCodeOk

`func (o *OpenApiWorkbenchIssue) GetParentCodeOk() (*int64, bool)`

GetParentCodeOk returns a tuple with the ParentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCode

`func (o *OpenApiWorkbenchIssue) SetParentCode(v int64)`

SetParentCode sets ParentCode field to given value.

### HasParentCode

`func (o *OpenApiWorkbenchIssue) HasParentCode() bool`

HasParentCode returns a boolean if a field has been set.

### SetParentCodeNil

`func (o *OpenApiWorkbenchIssue) SetParentCodeNil(b bool)`

 SetParentCodeNil sets the value for ParentCode to be an explicit nil

### UnsetParentCode
`func (o *OpenApiWorkbenchIssue) UnsetParentCode()`

UnsetParentCode ensures that no value is present for ParentCode, not even an explicit nil
### GetParentIssue

`func (o *OpenApiWorkbenchIssue) GetParentIssue() IssueSimpleData`

GetParentIssue returns the ParentIssue field if non-nil, zero value otherwise.

### GetParentIssueOk

`func (o *OpenApiWorkbenchIssue) GetParentIssueOk() (*IssueSimpleData, bool)`

GetParentIssueOk returns a tuple with the ParentIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIssue

`func (o *OpenApiWorkbenchIssue) SetParentIssue(v IssueSimpleData)`

SetParentIssue sets ParentIssue field to given value.

### HasParentIssue

`func (o *OpenApiWorkbenchIssue) HasParentIssue() bool`

HasParentIssue returns a boolean if a field has been set.

### GetPriority

`func (o *OpenApiWorkbenchIssue) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OpenApiWorkbenchIssue) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OpenApiWorkbenchIssue) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *OpenApiWorkbenchIssue) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *OpenApiWorkbenchIssue) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *OpenApiWorkbenchIssue) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetProject

`func (o *OpenApiWorkbenchIssue) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *OpenApiWorkbenchIssue) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *OpenApiWorkbenchIssue) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *OpenApiWorkbenchIssue) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectId

`func (o *OpenApiWorkbenchIssue) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *OpenApiWorkbenchIssue) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *OpenApiWorkbenchIssue) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *OpenApiWorkbenchIssue) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *OpenApiWorkbenchIssue) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *OpenApiWorkbenchIssue) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetSubTasks

`func (o *OpenApiWorkbenchIssue) GetSubTasks() []SubTask`

GetSubTasks returns the SubTasks field if non-nil, zero value otherwise.

### GetSubTasksOk

`func (o *OpenApiWorkbenchIssue) GetSubTasksOk() (*[]SubTask, bool)`

GetSubTasksOk returns a tuple with the SubTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTasks

`func (o *OpenApiWorkbenchIssue) SetSubTasks(v []SubTask)`

SetSubTasks sets SubTasks field to given value.

### HasSubTasks

`func (o *OpenApiWorkbenchIssue) HasSubTasks() bool`

HasSubTasks returns a boolean if a field has been set.

### SetSubTasksNil

`func (o *OpenApiWorkbenchIssue) SetSubTasksNil(b bool)`

 SetSubTasksNil sets the value for SubTasks to be an explicit nil

### UnsetSubTasks
`func (o *OpenApiWorkbenchIssue) UnsetSubTasks()`

UnsetSubTasks ensures that no value is present for SubTasks, not even an explicit nil
### GetType

`func (o *OpenApiWorkbenchIssue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpenApiWorkbenchIssue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpenApiWorkbenchIssue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OpenApiWorkbenchIssue) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *OpenApiWorkbenchIssue) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OpenApiWorkbenchIssue) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUpdatedAt

`func (o *OpenApiWorkbenchIssue) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OpenApiWorkbenchIssue) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OpenApiWorkbenchIssue) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OpenApiWorkbenchIssue) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *OpenApiWorkbenchIssue) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *OpenApiWorkbenchIssue) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


