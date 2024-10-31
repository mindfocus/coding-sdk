# IssueListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssigneeId** | Pointer to **int64** | 处理人 Id | [optional] 
**Code** | Pointer to **int64** | 事项 Code | [optional] 
**CompletedAt** | Pointer to **NullableInt64** | 完成时间时间戳 | [optional] 
**CreatedAt** | Pointer to **NullableInt64** | 创建时间时间戳 | [optional] 
**CreatorId** | Pointer to **int64** | 创建人 Id | [optional] 
**CustomFields** | Pointer to [**[]IssueCustomField**](IssueCustomField.md) | 自定义属性 | [optional] 
**Description** | Pointer to **string** | 描述 | [optional] [default to ""]
**DueDate** | Pointer to **NullableInt64** | 截止日期时间戳 | [optional] 
**IssueStatusId** | Pointer to **int64** | 状态 Id | [optional] 
**IssueStatusName** | Pointer to **string** | 状态名称 | [optional] [default to ""]
**IssueStatusType** | Pointer to **string** | 状态类型：  TODO｜PROCESSING｜COMPLETED | [optional] [default to ""]
**IssueTypeDetail** | Pointer to [**IssueTypeDetail**](IssueTypeDetail.md) |  | [optional] 
**IssueTypeId** | Pointer to **NullableInt64** | 事项类型 ID | [optional] 
**Iteration** | Pointer to [**IterationSimple**](IterationSimple.md) |  | [optional] 
**IterationId** | Pointer to **int64** | 迭代 Id | [optional] 
**Labels** | Pointer to [**[]IssueProjectLabel**](IssueProjectLabel.md) | 事项标签 | [optional] 
**Name** | Pointer to **string** | 名称 | [optional] [default to ""]
**ParentCode** | Pointer to **NullableInt64** | 父事项code | [optional] 
**ParentId** | Pointer to **NullableInt64** | 父事项ID | [optional] 
**ParentType** | Pointer to **string** | 父事项类型：  DEFECT - 缺陷，  REQUIREMENT - 需求，  MISSION - 任务，  EPIC - 史诗 | [optional] [default to ""]
**Priority** | Pointer to **string** | 优先级：  \&quot;0\&quot; - 低，  \&quot;1\&quot; - 中，  \&quot;2\&quot; - 高，  \&quot;3\&quot; - 紧急，  \&quot;\&quot; - 未指定 | [optional] [default to ""]
**StartDate** | Pointer to **NullableInt64** | 开始日期时间戳 | [optional] 
**StoryPoint** | Pointer to **NullableString** | 故事点，例如：\&quot;0.5\&quot;、\&quot;0\&quot;，  空字符串 \&quot;\&quot; 表示未指定 | [optional] [default to ""]
**Type** | Pointer to **string** | 事项类型：  DEFECT - 缺陷，  REQUIREMENT - 需求，  MISSION - 任务，  EPIC - 史诗，  SUB_TASK - 子工作项 | [optional] [default to ""]
**UpdatedAt** | Pointer to **NullableInt64** | 修改时间时间戳 | [optional] 
**WorkingHours** | Pointer to **float32** | 工时（小时） | [optional] 
**Assignees** | Pointer to [**[]User**](User.md) | 经办人列表 | [optional] 

## Methods

### NewIssueListData

`func NewIssueListData() *IssueListData`

NewIssueListData instantiates a new IssueListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueListDataWithDefaults

`func NewIssueListDataWithDefaults() *IssueListData`

NewIssueListDataWithDefaults instantiates a new IssueListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigneeId

`func (o *IssueListData) GetAssigneeId() int64`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *IssueListData) GetAssigneeIdOk() (*int64, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *IssueListData) SetAssigneeId(v int64)`

SetAssigneeId sets AssigneeId field to given value.

### HasAssigneeId

`func (o *IssueListData) HasAssigneeId() bool`

HasAssigneeId returns a boolean if a field has been set.

### GetCode

`func (o *IssueListData) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IssueListData) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IssueListData) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *IssueListData) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCompletedAt

`func (o *IssueListData) GetCompletedAt() int64`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *IssueListData) GetCompletedAtOk() (*int64, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *IssueListData) SetCompletedAt(v int64)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *IssueListData) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *IssueListData) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *IssueListData) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetCreatedAt

`func (o *IssueListData) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueListData) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueListData) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssueListData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *IssueListData) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *IssueListData) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatorId

`func (o *IssueListData) GetCreatorId() int64`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *IssueListData) GetCreatorIdOk() (*int64, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *IssueListData) SetCreatorId(v int64)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *IssueListData) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetCustomFields

`func (o *IssueListData) GetCustomFields() []IssueCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IssueListData) GetCustomFieldsOk() (*[]IssueCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IssueListData) SetCustomFields(v []IssueCustomField)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IssueListData) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### SetCustomFieldsNil

`func (o *IssueListData) SetCustomFieldsNil(b bool)`

 SetCustomFieldsNil sets the value for CustomFields to be an explicit nil

### UnsetCustomFields
`func (o *IssueListData) UnsetCustomFields()`

UnsetCustomFields ensures that no value is present for CustomFields, not even an explicit nil
### GetDescription

`func (o *IssueListData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueListData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueListData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueListData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueDate

`func (o *IssueListData) GetDueDate() int64`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *IssueListData) GetDueDateOk() (*int64, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *IssueListData) SetDueDate(v int64)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *IssueListData) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### SetDueDateNil

`func (o *IssueListData) SetDueDateNil(b bool)`

 SetDueDateNil sets the value for DueDate to be an explicit nil

### UnsetDueDate
`func (o *IssueListData) UnsetDueDate()`

UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
### GetIssueStatusId

`func (o *IssueListData) GetIssueStatusId() int64`

GetIssueStatusId returns the IssueStatusId field if non-nil, zero value otherwise.

### GetIssueStatusIdOk

`func (o *IssueListData) GetIssueStatusIdOk() (*int64, bool)`

GetIssueStatusIdOk returns a tuple with the IssueStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueStatusId

`func (o *IssueListData) SetIssueStatusId(v int64)`

SetIssueStatusId sets IssueStatusId field to given value.

### HasIssueStatusId

`func (o *IssueListData) HasIssueStatusId() bool`

HasIssueStatusId returns a boolean if a field has been set.

### GetIssueStatusName

`func (o *IssueListData) GetIssueStatusName() string`

GetIssueStatusName returns the IssueStatusName field if non-nil, zero value otherwise.

### GetIssueStatusNameOk

`func (o *IssueListData) GetIssueStatusNameOk() (*string, bool)`

GetIssueStatusNameOk returns a tuple with the IssueStatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueStatusName

`func (o *IssueListData) SetIssueStatusName(v string)`

SetIssueStatusName sets IssueStatusName field to given value.

### HasIssueStatusName

`func (o *IssueListData) HasIssueStatusName() bool`

HasIssueStatusName returns a boolean if a field has been set.

### GetIssueStatusType

`func (o *IssueListData) GetIssueStatusType() string`

GetIssueStatusType returns the IssueStatusType field if non-nil, zero value otherwise.

### GetIssueStatusTypeOk

`func (o *IssueListData) GetIssueStatusTypeOk() (*string, bool)`

GetIssueStatusTypeOk returns a tuple with the IssueStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueStatusType

`func (o *IssueListData) SetIssueStatusType(v string)`

SetIssueStatusType sets IssueStatusType field to given value.

### HasIssueStatusType

`func (o *IssueListData) HasIssueStatusType() bool`

HasIssueStatusType returns a boolean if a field has been set.

### GetIssueTypeDetail

`func (o *IssueListData) GetIssueTypeDetail() IssueTypeDetail`

GetIssueTypeDetail returns the IssueTypeDetail field if non-nil, zero value otherwise.

### GetIssueTypeDetailOk

`func (o *IssueListData) GetIssueTypeDetailOk() (*IssueTypeDetail, bool)`

GetIssueTypeDetailOk returns a tuple with the IssueTypeDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeDetail

`func (o *IssueListData) SetIssueTypeDetail(v IssueTypeDetail)`

SetIssueTypeDetail sets IssueTypeDetail field to given value.

### HasIssueTypeDetail

`func (o *IssueListData) HasIssueTypeDetail() bool`

HasIssueTypeDetail returns a boolean if a field has been set.

### GetIssueTypeId

`func (o *IssueListData) GetIssueTypeId() int64`

GetIssueTypeId returns the IssueTypeId field if non-nil, zero value otherwise.

### GetIssueTypeIdOk

`func (o *IssueListData) GetIssueTypeIdOk() (*int64, bool)`

GetIssueTypeIdOk returns a tuple with the IssueTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeId

`func (o *IssueListData) SetIssueTypeId(v int64)`

SetIssueTypeId sets IssueTypeId field to given value.

### HasIssueTypeId

`func (o *IssueListData) HasIssueTypeId() bool`

HasIssueTypeId returns a boolean if a field has been set.

### SetIssueTypeIdNil

`func (o *IssueListData) SetIssueTypeIdNil(b bool)`

 SetIssueTypeIdNil sets the value for IssueTypeId to be an explicit nil

### UnsetIssueTypeId
`func (o *IssueListData) UnsetIssueTypeId()`

UnsetIssueTypeId ensures that no value is present for IssueTypeId, not even an explicit nil
### GetIteration

`func (o *IssueListData) GetIteration() IterationSimple`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *IssueListData) GetIterationOk() (*IterationSimple, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *IssueListData) SetIteration(v IterationSimple)`

SetIteration sets Iteration field to given value.

### HasIteration

`func (o *IssueListData) HasIteration() bool`

HasIteration returns a boolean if a field has been set.

### GetIterationId

`func (o *IssueListData) GetIterationId() int64`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *IssueListData) GetIterationIdOk() (*int64, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *IssueListData) SetIterationId(v int64)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *IssueListData) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.

### GetLabels

`func (o *IssueListData) GetLabels() []IssueProjectLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IssueListData) GetLabelsOk() (*[]IssueProjectLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IssueListData) SetLabels(v []IssueProjectLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *IssueListData) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *IssueListData) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *IssueListData) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetName

`func (o *IssueListData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueListData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueListData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueListData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentCode

`func (o *IssueListData) GetParentCode() int64`

GetParentCode returns the ParentCode field if non-nil, zero value otherwise.

### GetParentCodeOk

`func (o *IssueListData) GetParentCodeOk() (*int64, bool)`

GetParentCodeOk returns a tuple with the ParentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCode

`func (o *IssueListData) SetParentCode(v int64)`

SetParentCode sets ParentCode field to given value.

### HasParentCode

`func (o *IssueListData) HasParentCode() bool`

HasParentCode returns a boolean if a field has been set.

### SetParentCodeNil

`func (o *IssueListData) SetParentCodeNil(b bool)`

 SetParentCodeNil sets the value for ParentCode to be an explicit nil

### UnsetParentCode
`func (o *IssueListData) UnsetParentCode()`

UnsetParentCode ensures that no value is present for ParentCode, not even an explicit nil
### GetParentId

`func (o *IssueListData) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *IssueListData) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *IssueListData) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *IssueListData) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *IssueListData) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *IssueListData) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetParentType

`func (o *IssueListData) GetParentType() string`

GetParentType returns the ParentType field if non-nil, zero value otherwise.

### GetParentTypeOk

`func (o *IssueListData) GetParentTypeOk() (*string, bool)`

GetParentTypeOk returns a tuple with the ParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentType

`func (o *IssueListData) SetParentType(v string)`

SetParentType sets ParentType field to given value.

### HasParentType

`func (o *IssueListData) HasParentType() bool`

HasParentType returns a boolean if a field has been set.

### GetPriority

`func (o *IssueListData) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IssueListData) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IssueListData) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IssueListData) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStartDate

`func (o *IssueListData) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *IssueListData) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *IssueListData) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *IssueListData) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *IssueListData) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *IssueListData) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetStoryPoint

`func (o *IssueListData) GetStoryPoint() string`

GetStoryPoint returns the StoryPoint field if non-nil, zero value otherwise.

### GetStoryPointOk

`func (o *IssueListData) GetStoryPointOk() (*string, bool)`

GetStoryPointOk returns a tuple with the StoryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryPoint

`func (o *IssueListData) SetStoryPoint(v string)`

SetStoryPoint sets StoryPoint field to given value.

### HasStoryPoint

`func (o *IssueListData) HasStoryPoint() bool`

HasStoryPoint returns a boolean if a field has been set.

### SetStoryPointNil

`func (o *IssueListData) SetStoryPointNil(b bool)`

 SetStoryPointNil sets the value for StoryPoint to be an explicit nil

### UnsetStoryPoint
`func (o *IssueListData) UnsetStoryPoint()`

UnsetStoryPoint ensures that no value is present for StoryPoint, not even an explicit nil
### GetType

`func (o *IssueListData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssueListData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssueListData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IssueListData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IssueListData) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssueListData) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssueListData) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IssueListData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *IssueListData) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *IssueListData) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetWorkingHours

`func (o *IssueListData) GetWorkingHours() float32`

GetWorkingHours returns the WorkingHours field if non-nil, zero value otherwise.

### GetWorkingHoursOk

`func (o *IssueListData) GetWorkingHoursOk() (*float32, bool)`

GetWorkingHoursOk returns a tuple with the WorkingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHours

`func (o *IssueListData) SetWorkingHours(v float32)`

SetWorkingHours sets WorkingHours field to given value.

### HasWorkingHours

`func (o *IssueListData) HasWorkingHours() bool`

HasWorkingHours returns a boolean if a field has been set.

### GetAssignees

`func (o *IssueListData) GetAssignees() []User`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *IssueListData) GetAssigneesOk() (*[]User, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *IssueListData) SetAssignees(v []User)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *IssueListData) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


