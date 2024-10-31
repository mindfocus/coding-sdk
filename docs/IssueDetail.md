# IssueDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to [**User**](User.md) |  | [optional] 
**Code** | Pointer to **int64** | 事项 Code | [optional] 
**CompletedAt** | Pointer to **int64** | 完成时间戳 | [optional] 
**CreatedAt** | Pointer to **int64** | 创建时间戳 | [optional] 
**Creator** | Pointer to [**User**](User.md) |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFields**](CustomFields.md) | 自定义属性列表 | [optional] 
**DefectType** | Pointer to [**DefectType**](DefectType.md) |  | [optional] 
**Description** | Pointer to **string** | 描述 | [optional] [default to ""]
**DueDate** | Pointer to **int64** | 截止日期时间戳 | [optional] 
**Epic** | Pointer to [**Epic**](Epic.md) |  | [optional] 
**Files** | Pointer to [**[]IssueFile**](IssueFile.md) | 附件列表 | [optional] 
**IssueStatusId** | Pointer to **int64** | 事项状态 Id | [optional] 
**IssueStatusName** | Pointer to **string** | 事项状态名称 | [optional] [default to ""]
**IssueStatusType** | Pointer to **string** | 事项状态类型 | [optional] [default to ""]
**IssueTypeDetail** | Pointer to [**IssueTypeDetail**](IssueTypeDetail.md) |  | [optional] 
**IssueTypeId** | Pointer to **int64** | 事项类型 ID | [optional] 
**Iteration** | Pointer to [**IterationSimple**](IterationSimple.md) |  | [optional] 
**IterationId** | Pointer to **int64** | 迭代 Id | [optional] 
**Labels** | Pointer to [**[]IssueProjectLabel**](IssueProjectLabel.md) | 标签列表 | [optional] 
**Name** | Pointer to **string** | 名称 | [optional] [default to ""]
**Parent** | Pointer to [**IssueSimpleData**](IssueSimpleData.md) |  | [optional] 
**ParentType** | Pointer to **string** | 父事项类型 | [optional] [default to ""]
**Priority** | Pointer to **string** | 优先级：  0 - 低，  1 - 中，  2 - 高，  3 - 紧急，  \&quot;\&quot; - 未指定 | [optional] [default to ""]
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**ProjectModule** | Pointer to [**IssueProjectModule**](IssueProjectModule.md) |  | [optional] 
**RequirementType** | Pointer to [**RequirementType**](RequirementType.md) |  | [optional] 
**StartDate** | Pointer to **int64** | 开始日期时间戳 | [optional] 
**StoryPoint** | Pointer to **string** | 故事点，例如：0.5、0、1  空字符串 \&quot;\&quot; 表示未指定。 | [optional] [default to ""]
**SubTasks** | Pointer to [**[]SubTask**](SubTask.md) | 子工作项列表 | [optional] 
**ThirdLinks** | Pointer to [**[]IssueThirdLink**](IssueThirdLink.md) | 第三方链接列表 | [optional] 
**Type** | Pointer to **string** | 事项类型：  DEFECT - 缺陷，  REQUIREMENT - 需求，  MISSION - 任务，  EPIC - 史诗，  SUB_TASK - 子工作项 | [optional] [default to ""]
**UpdatedAt** | Pointer to **int64** | 修改时间戳 | [optional] 
**Watchers** | Pointer to [**[]User**](User.md) | 关注人列表 | [optional] 
**WorkingHours** | Pointer to **float32** | 工时（小时数） | [optional] 
**Assignees** | Pointer to [**[]User**](User.md) | 经办人列表 | [optional] 

## Methods

### NewIssueDetail

`func NewIssueDetail() *IssueDetail`

NewIssueDetail instantiates a new IssueDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueDetailWithDefaults

`func NewIssueDetailWithDefaults() *IssueDetail`

NewIssueDetailWithDefaults instantiates a new IssueDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *IssueDetail) GetAssignee() User`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *IssueDetail) GetAssigneeOk() (*User, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *IssueDetail) SetAssignee(v User)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *IssueDetail) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCode

`func (o *IssueDetail) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IssueDetail) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IssueDetail) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *IssueDetail) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCompletedAt

`func (o *IssueDetail) GetCompletedAt() int64`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *IssueDetail) GetCompletedAtOk() (*int64, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *IssueDetail) SetCompletedAt(v int64)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *IssueDetail) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IssueDetail) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueDetail) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueDetail) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssueDetail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *IssueDetail) GetCreator() User`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *IssueDetail) GetCreatorOk() (*User, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *IssueDetail) SetCreator(v User)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *IssueDetail) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetCustomFields

`func (o *IssueDetail) GetCustomFields() []CustomFields`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IssueDetail) GetCustomFieldsOk() (*[]CustomFields, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IssueDetail) SetCustomFields(v []CustomFields)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IssueDetail) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDefectType

`func (o *IssueDetail) GetDefectType() DefectType`

GetDefectType returns the DefectType field if non-nil, zero value otherwise.

### GetDefectTypeOk

`func (o *IssueDetail) GetDefectTypeOk() (*DefectType, bool)`

GetDefectTypeOk returns a tuple with the DefectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectType

`func (o *IssueDetail) SetDefectType(v DefectType)`

SetDefectType sets DefectType field to given value.

### HasDefectType

`func (o *IssueDetail) HasDefectType() bool`

HasDefectType returns a boolean if a field has been set.

### GetDescription

`func (o *IssueDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueDate

`func (o *IssueDetail) GetDueDate() int64`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *IssueDetail) GetDueDateOk() (*int64, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *IssueDetail) SetDueDate(v int64)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *IssueDetail) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEpic

`func (o *IssueDetail) GetEpic() Epic`

GetEpic returns the Epic field if non-nil, zero value otherwise.

### GetEpicOk

`func (o *IssueDetail) GetEpicOk() (*Epic, bool)`

GetEpicOk returns a tuple with the Epic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpic

`func (o *IssueDetail) SetEpic(v Epic)`

SetEpic sets Epic field to given value.

### HasEpic

`func (o *IssueDetail) HasEpic() bool`

HasEpic returns a boolean if a field has been set.

### GetFiles

`func (o *IssueDetail) GetFiles() []IssueFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *IssueDetail) GetFilesOk() (*[]IssueFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *IssueDetail) SetFiles(v []IssueFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *IssueDetail) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetIssueStatusId

`func (o *IssueDetail) GetIssueStatusId() int64`

GetIssueStatusId returns the IssueStatusId field if non-nil, zero value otherwise.

### GetIssueStatusIdOk

`func (o *IssueDetail) GetIssueStatusIdOk() (*int64, bool)`

GetIssueStatusIdOk returns a tuple with the IssueStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueStatusId

`func (o *IssueDetail) SetIssueStatusId(v int64)`

SetIssueStatusId sets IssueStatusId field to given value.

### HasIssueStatusId

`func (o *IssueDetail) HasIssueStatusId() bool`

HasIssueStatusId returns a boolean if a field has been set.

### GetIssueStatusName

`func (o *IssueDetail) GetIssueStatusName() string`

GetIssueStatusName returns the IssueStatusName field if non-nil, zero value otherwise.

### GetIssueStatusNameOk

`func (o *IssueDetail) GetIssueStatusNameOk() (*string, bool)`

GetIssueStatusNameOk returns a tuple with the IssueStatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueStatusName

`func (o *IssueDetail) SetIssueStatusName(v string)`

SetIssueStatusName sets IssueStatusName field to given value.

### HasIssueStatusName

`func (o *IssueDetail) HasIssueStatusName() bool`

HasIssueStatusName returns a boolean if a field has been set.

### GetIssueStatusType

`func (o *IssueDetail) GetIssueStatusType() string`

GetIssueStatusType returns the IssueStatusType field if non-nil, zero value otherwise.

### GetIssueStatusTypeOk

`func (o *IssueDetail) GetIssueStatusTypeOk() (*string, bool)`

GetIssueStatusTypeOk returns a tuple with the IssueStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueStatusType

`func (o *IssueDetail) SetIssueStatusType(v string)`

SetIssueStatusType sets IssueStatusType field to given value.

### HasIssueStatusType

`func (o *IssueDetail) HasIssueStatusType() bool`

HasIssueStatusType returns a boolean if a field has been set.

### GetIssueTypeDetail

`func (o *IssueDetail) GetIssueTypeDetail() IssueTypeDetail`

GetIssueTypeDetail returns the IssueTypeDetail field if non-nil, zero value otherwise.

### GetIssueTypeDetailOk

`func (o *IssueDetail) GetIssueTypeDetailOk() (*IssueTypeDetail, bool)`

GetIssueTypeDetailOk returns a tuple with the IssueTypeDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeDetail

`func (o *IssueDetail) SetIssueTypeDetail(v IssueTypeDetail)`

SetIssueTypeDetail sets IssueTypeDetail field to given value.

### HasIssueTypeDetail

`func (o *IssueDetail) HasIssueTypeDetail() bool`

HasIssueTypeDetail returns a boolean if a field has been set.

### GetIssueTypeId

`func (o *IssueDetail) GetIssueTypeId() int64`

GetIssueTypeId returns the IssueTypeId field if non-nil, zero value otherwise.

### GetIssueTypeIdOk

`func (o *IssueDetail) GetIssueTypeIdOk() (*int64, bool)`

GetIssueTypeIdOk returns a tuple with the IssueTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeId

`func (o *IssueDetail) SetIssueTypeId(v int64)`

SetIssueTypeId sets IssueTypeId field to given value.

### HasIssueTypeId

`func (o *IssueDetail) HasIssueTypeId() bool`

HasIssueTypeId returns a boolean if a field has been set.

### GetIteration

`func (o *IssueDetail) GetIteration() IterationSimple`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *IssueDetail) GetIterationOk() (*IterationSimple, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *IssueDetail) SetIteration(v IterationSimple)`

SetIteration sets Iteration field to given value.

### HasIteration

`func (o *IssueDetail) HasIteration() bool`

HasIteration returns a boolean if a field has been set.

### GetIterationId

`func (o *IssueDetail) GetIterationId() int64`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *IssueDetail) GetIterationIdOk() (*int64, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *IssueDetail) SetIterationId(v int64)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *IssueDetail) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.

### GetLabels

`func (o *IssueDetail) GetLabels() []IssueProjectLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IssueDetail) GetLabelsOk() (*[]IssueProjectLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IssueDetail) SetLabels(v []IssueProjectLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *IssueDetail) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *IssueDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *IssueDetail) GetParent() IssueSimpleData`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IssueDetail) GetParentOk() (*IssueSimpleData, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IssueDetail) SetParent(v IssueSimpleData)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IssueDetail) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetParentType

`func (o *IssueDetail) GetParentType() string`

GetParentType returns the ParentType field if non-nil, zero value otherwise.

### GetParentTypeOk

`func (o *IssueDetail) GetParentTypeOk() (*string, bool)`

GetParentTypeOk returns a tuple with the ParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentType

`func (o *IssueDetail) SetParentType(v string)`

SetParentType sets ParentType field to given value.

### HasParentType

`func (o *IssueDetail) HasParentType() bool`

HasParentType returns a boolean if a field has been set.

### GetPriority

`func (o *IssueDetail) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IssueDetail) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IssueDetail) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IssueDetail) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProject

`func (o *IssueDetail) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *IssueDetail) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *IssueDetail) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *IssueDetail) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectModule

`func (o *IssueDetail) GetProjectModule() IssueProjectModule`

GetProjectModule returns the ProjectModule field if non-nil, zero value otherwise.

### GetProjectModuleOk

`func (o *IssueDetail) GetProjectModuleOk() (*IssueProjectModule, bool)`

GetProjectModuleOk returns a tuple with the ProjectModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectModule

`func (o *IssueDetail) SetProjectModule(v IssueProjectModule)`

SetProjectModule sets ProjectModule field to given value.

### HasProjectModule

`func (o *IssueDetail) HasProjectModule() bool`

HasProjectModule returns a boolean if a field has been set.

### GetRequirementType

`func (o *IssueDetail) GetRequirementType() RequirementType`

GetRequirementType returns the RequirementType field if non-nil, zero value otherwise.

### GetRequirementTypeOk

`func (o *IssueDetail) GetRequirementTypeOk() (*RequirementType, bool)`

GetRequirementTypeOk returns a tuple with the RequirementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementType

`func (o *IssueDetail) SetRequirementType(v RequirementType)`

SetRequirementType sets RequirementType field to given value.

### HasRequirementType

`func (o *IssueDetail) HasRequirementType() bool`

HasRequirementType returns a boolean if a field has been set.

### GetStartDate

`func (o *IssueDetail) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *IssueDetail) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *IssueDetail) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *IssueDetail) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStoryPoint

`func (o *IssueDetail) GetStoryPoint() string`

GetStoryPoint returns the StoryPoint field if non-nil, zero value otherwise.

### GetStoryPointOk

`func (o *IssueDetail) GetStoryPointOk() (*string, bool)`

GetStoryPointOk returns a tuple with the StoryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryPoint

`func (o *IssueDetail) SetStoryPoint(v string)`

SetStoryPoint sets StoryPoint field to given value.

### HasStoryPoint

`func (o *IssueDetail) HasStoryPoint() bool`

HasStoryPoint returns a boolean if a field has been set.

### GetSubTasks

`func (o *IssueDetail) GetSubTasks() []SubTask`

GetSubTasks returns the SubTasks field if non-nil, zero value otherwise.

### GetSubTasksOk

`func (o *IssueDetail) GetSubTasksOk() (*[]SubTask, bool)`

GetSubTasksOk returns a tuple with the SubTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTasks

`func (o *IssueDetail) SetSubTasks(v []SubTask)`

SetSubTasks sets SubTasks field to given value.

### HasSubTasks

`func (o *IssueDetail) HasSubTasks() bool`

HasSubTasks returns a boolean if a field has been set.

### GetThirdLinks

`func (o *IssueDetail) GetThirdLinks() []IssueThirdLink`

GetThirdLinks returns the ThirdLinks field if non-nil, zero value otherwise.

### GetThirdLinksOk

`func (o *IssueDetail) GetThirdLinksOk() (*[]IssueThirdLink, bool)`

GetThirdLinksOk returns a tuple with the ThirdLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdLinks

`func (o *IssueDetail) SetThirdLinks(v []IssueThirdLink)`

SetThirdLinks sets ThirdLinks field to given value.

### HasThirdLinks

`func (o *IssueDetail) HasThirdLinks() bool`

HasThirdLinks returns a boolean if a field has been set.

### GetType

`func (o *IssueDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssueDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssueDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IssueDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IssueDetail) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssueDetail) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssueDetail) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IssueDetail) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWatchers

`func (o *IssueDetail) GetWatchers() []User`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *IssueDetail) GetWatchersOk() (*[]User, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *IssueDetail) SetWatchers(v []User)`

SetWatchers sets Watchers field to given value.

### HasWatchers

`func (o *IssueDetail) HasWatchers() bool`

HasWatchers returns a boolean if a field has been set.

### GetWorkingHours

`func (o *IssueDetail) GetWorkingHours() float32`

GetWorkingHours returns the WorkingHours field if non-nil, zero value otherwise.

### GetWorkingHoursOk

`func (o *IssueDetail) GetWorkingHoursOk() (*float32, bool)`

GetWorkingHoursOk returns a tuple with the WorkingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHours

`func (o *IssueDetail) SetWorkingHours(v float32)`

SetWorkingHours sets WorkingHours field to given value.

### HasWorkingHours

`func (o *IssueDetail) HasWorkingHours() bool`

HasWorkingHours returns a boolean if a field has been set.

### GetAssignees

`func (o *IssueDetail) GetAssignees() []User`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *IssueDetail) GetAssigneesOk() (*[]User, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *IssueDetail) SetAssignees(v []User)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *IssueDetail) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


