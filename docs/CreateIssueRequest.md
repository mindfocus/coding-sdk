# CreateIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssigneeId** | Pointer to **int64** | 处理人 Id | [optional] 
**CustomFieldValues** | Pointer to [**[]IssueCustomFieldForm**](IssueCustomFieldForm.md) | 自定义属性值列表 | [optional] 
**DefectTypeId** | Pointer to **int64** | 缺陷类型 Id | [optional] 
**Description** | Pointer to **string** | 描述 | [optional] 
**DueDate** | Pointer to **map[string]interface{}** | 截止日期，格式：2021-01-01 | [optional] 
**EpicCode** | Pointer to **int64** | 史诗 Code，Type 为 EPIC 或 SUB_TASK时，不传该值 | [optional] 
**FileIds** | Pointer to **[]int64** | 附件的文件 Id 列表 | [optional] 
**IssueTypeId** | Pointer to **int64** | 事项类型 Id | [optional] 
**IterationCode** | Pointer to **int64** | 迭代 Code，Type 为 EPIC 或 SUB_TASK时，忽略该值 | [optional] 
**LabelIds** | Pointer to **[]int64** | 标签 Id 列表 | [optional] 
**Name** | **string** | 事项名称 | 
**ParentCode** | Pointer to **int64** | 父事项 Code  敏捷模式：Type 为 SUB_TASK 时必须指定  经典模式：Type 为 MISSION、REQUIREMENT 时可指定 | [optional] 
**Priority** | **string** | 紧急程度  \&quot;0\&quot; - 低  \&quot;1\&quot; - 中  \&quot;2\&quot; - 高  \&quot;3\&quot; - 紧急 | 
**ProjectModuleId** | Pointer to **int64** | 项目模块 Id | [optional] 
**ProjectName** | **string** | 项目名称 | 
**ReleaseCodes** | Pointer to **[]int64** | 版本code列表 | [optional] 
**RequirementTypeId** | Pointer to **int64** | 需求类型 Id | [optional] 
**StartDate** | Pointer to **map[string]interface{}** | 开始日期，格式：2021-01-01 | [optional] 
**StatusId** | Pointer to **int64** | 事项状态 Id | [optional] 
**StoryPoint** | Pointer to **string** | 故事点，例如：0.5、1 | [optional] 
**TargetSortCode** | Pointer to **int64** | 排序目标位置的事项 code  可不填，排在底位 | [optional] 
**ThirdLinks** | Pointer to [**[]IssueThirdLinkForm**](IssueThirdLinkForm.md) | 第三方链接列表 | [optional] 
**Type** | **string** | 事项类型  DEFECT - 缺陷  REQUIREMENT - 需求  MISSION - 任务  EPIC - 史诗  SUB_TASK - 子任务 | 
**WatcherIds** | Pointer to **[]int64** | 关注人 Id 列表 | [optional] 
**WorkingHours** | Pointer to **float32** | 工时（小时数） | [optional] 

## Methods

### NewCreateIssueRequest

`func NewCreateIssueRequest(name string, priority string, projectName string, type_ string, ) *CreateIssueRequest`

NewCreateIssueRequest instantiates a new CreateIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIssueRequestWithDefaults

`func NewCreateIssueRequestWithDefaults() *CreateIssueRequest`

NewCreateIssueRequestWithDefaults instantiates a new CreateIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigneeId

`func (o *CreateIssueRequest) GetAssigneeId() int64`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *CreateIssueRequest) GetAssigneeIdOk() (*int64, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *CreateIssueRequest) SetAssigneeId(v int64)`

SetAssigneeId sets AssigneeId field to given value.

### HasAssigneeId

`func (o *CreateIssueRequest) HasAssigneeId() bool`

HasAssigneeId returns a boolean if a field has been set.

### GetCustomFieldValues

`func (o *CreateIssueRequest) GetCustomFieldValues() []IssueCustomFieldForm`

GetCustomFieldValues returns the CustomFieldValues field if non-nil, zero value otherwise.

### GetCustomFieldValuesOk

`func (o *CreateIssueRequest) GetCustomFieldValuesOk() (*[]IssueCustomFieldForm, bool)`

GetCustomFieldValuesOk returns a tuple with the CustomFieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldValues

`func (o *CreateIssueRequest) SetCustomFieldValues(v []IssueCustomFieldForm)`

SetCustomFieldValues sets CustomFieldValues field to given value.

### HasCustomFieldValues

`func (o *CreateIssueRequest) HasCustomFieldValues() bool`

HasCustomFieldValues returns a boolean if a field has been set.

### GetDefectTypeId

`func (o *CreateIssueRequest) GetDefectTypeId() int64`

GetDefectTypeId returns the DefectTypeId field if non-nil, zero value otherwise.

### GetDefectTypeIdOk

`func (o *CreateIssueRequest) GetDefectTypeIdOk() (*int64, bool)`

GetDefectTypeIdOk returns a tuple with the DefectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectTypeId

`func (o *CreateIssueRequest) SetDefectTypeId(v int64)`

SetDefectTypeId sets DefectTypeId field to given value.

### HasDefectTypeId

`func (o *CreateIssueRequest) HasDefectTypeId() bool`

HasDefectTypeId returns a boolean if a field has been set.

### GetDescription

`func (o *CreateIssueRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateIssueRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateIssueRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateIssueRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDueDate

`func (o *CreateIssueRequest) GetDueDate() map[string]interface{}`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *CreateIssueRequest) GetDueDateOk() (*map[string]interface{}, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *CreateIssueRequest) SetDueDate(v map[string]interface{})`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *CreateIssueRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetEpicCode

`func (o *CreateIssueRequest) GetEpicCode() int64`

GetEpicCode returns the EpicCode field if non-nil, zero value otherwise.

### GetEpicCodeOk

`func (o *CreateIssueRequest) GetEpicCodeOk() (*int64, bool)`

GetEpicCodeOk returns a tuple with the EpicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpicCode

`func (o *CreateIssueRequest) SetEpicCode(v int64)`

SetEpicCode sets EpicCode field to given value.

### HasEpicCode

`func (o *CreateIssueRequest) HasEpicCode() bool`

HasEpicCode returns a boolean if a field has been set.

### GetFileIds

`func (o *CreateIssueRequest) GetFileIds() []int64`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *CreateIssueRequest) GetFileIdsOk() (*[]int64, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *CreateIssueRequest) SetFileIds(v []int64)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *CreateIssueRequest) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.

### GetIssueTypeId

`func (o *CreateIssueRequest) GetIssueTypeId() int64`

GetIssueTypeId returns the IssueTypeId field if non-nil, zero value otherwise.

### GetIssueTypeIdOk

`func (o *CreateIssueRequest) GetIssueTypeIdOk() (*int64, bool)`

GetIssueTypeIdOk returns a tuple with the IssueTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeId

`func (o *CreateIssueRequest) SetIssueTypeId(v int64)`

SetIssueTypeId sets IssueTypeId field to given value.

### HasIssueTypeId

`func (o *CreateIssueRequest) HasIssueTypeId() bool`

HasIssueTypeId returns a boolean if a field has been set.

### GetIterationCode

`func (o *CreateIssueRequest) GetIterationCode() int64`

GetIterationCode returns the IterationCode field if non-nil, zero value otherwise.

### GetIterationCodeOk

`func (o *CreateIssueRequest) GetIterationCodeOk() (*int64, bool)`

GetIterationCodeOk returns a tuple with the IterationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCode

`func (o *CreateIssueRequest) SetIterationCode(v int64)`

SetIterationCode sets IterationCode field to given value.

### HasIterationCode

`func (o *CreateIssueRequest) HasIterationCode() bool`

HasIterationCode returns a boolean if a field has been set.

### GetLabelIds

`func (o *CreateIssueRequest) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *CreateIssueRequest) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *CreateIssueRequest) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *CreateIssueRequest) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetName

`func (o *CreateIssueRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIssueRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIssueRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParentCode

`func (o *CreateIssueRequest) GetParentCode() int64`

GetParentCode returns the ParentCode field if non-nil, zero value otherwise.

### GetParentCodeOk

`func (o *CreateIssueRequest) GetParentCodeOk() (*int64, bool)`

GetParentCodeOk returns a tuple with the ParentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCode

`func (o *CreateIssueRequest) SetParentCode(v int64)`

SetParentCode sets ParentCode field to given value.

### HasParentCode

`func (o *CreateIssueRequest) HasParentCode() bool`

HasParentCode returns a boolean if a field has been set.

### GetPriority

`func (o *CreateIssueRequest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateIssueRequest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateIssueRequest) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetProjectModuleId

`func (o *CreateIssueRequest) GetProjectModuleId() int64`

GetProjectModuleId returns the ProjectModuleId field if non-nil, zero value otherwise.

### GetProjectModuleIdOk

`func (o *CreateIssueRequest) GetProjectModuleIdOk() (*int64, bool)`

GetProjectModuleIdOk returns a tuple with the ProjectModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectModuleId

`func (o *CreateIssueRequest) SetProjectModuleId(v int64)`

SetProjectModuleId sets ProjectModuleId field to given value.

### HasProjectModuleId

`func (o *CreateIssueRequest) HasProjectModuleId() bool`

HasProjectModuleId returns a boolean if a field has been set.

### GetProjectName

`func (o *CreateIssueRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateIssueRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateIssueRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetReleaseCodes

`func (o *CreateIssueRequest) GetReleaseCodes() []int64`

GetReleaseCodes returns the ReleaseCodes field if non-nil, zero value otherwise.

### GetReleaseCodesOk

`func (o *CreateIssueRequest) GetReleaseCodesOk() (*[]int64, bool)`

GetReleaseCodesOk returns a tuple with the ReleaseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseCodes

`func (o *CreateIssueRequest) SetReleaseCodes(v []int64)`

SetReleaseCodes sets ReleaseCodes field to given value.

### HasReleaseCodes

`func (o *CreateIssueRequest) HasReleaseCodes() bool`

HasReleaseCodes returns a boolean if a field has been set.

### GetRequirementTypeId

`func (o *CreateIssueRequest) GetRequirementTypeId() int64`

GetRequirementTypeId returns the RequirementTypeId field if non-nil, zero value otherwise.

### GetRequirementTypeIdOk

`func (o *CreateIssueRequest) GetRequirementTypeIdOk() (*int64, bool)`

GetRequirementTypeIdOk returns a tuple with the RequirementTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementTypeId

`func (o *CreateIssueRequest) SetRequirementTypeId(v int64)`

SetRequirementTypeId sets RequirementTypeId field to given value.

### HasRequirementTypeId

`func (o *CreateIssueRequest) HasRequirementTypeId() bool`

HasRequirementTypeId returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateIssueRequest) GetStartDate() map[string]interface{}`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateIssueRequest) GetStartDateOk() (*map[string]interface{}, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateIssueRequest) SetStartDate(v map[string]interface{})`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateIssueRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatusId

`func (o *CreateIssueRequest) GetStatusId() int64`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *CreateIssueRequest) GetStatusIdOk() (*int64, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *CreateIssueRequest) SetStatusId(v int64)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *CreateIssueRequest) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetStoryPoint

`func (o *CreateIssueRequest) GetStoryPoint() string`

GetStoryPoint returns the StoryPoint field if non-nil, zero value otherwise.

### GetStoryPointOk

`func (o *CreateIssueRequest) GetStoryPointOk() (*string, bool)`

GetStoryPointOk returns a tuple with the StoryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryPoint

`func (o *CreateIssueRequest) SetStoryPoint(v string)`

SetStoryPoint sets StoryPoint field to given value.

### HasStoryPoint

`func (o *CreateIssueRequest) HasStoryPoint() bool`

HasStoryPoint returns a boolean if a field has been set.

### GetTargetSortCode

`func (o *CreateIssueRequest) GetTargetSortCode() int64`

GetTargetSortCode returns the TargetSortCode field if non-nil, zero value otherwise.

### GetTargetSortCodeOk

`func (o *CreateIssueRequest) GetTargetSortCodeOk() (*int64, bool)`

GetTargetSortCodeOk returns a tuple with the TargetSortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSortCode

`func (o *CreateIssueRequest) SetTargetSortCode(v int64)`

SetTargetSortCode sets TargetSortCode field to given value.

### HasTargetSortCode

`func (o *CreateIssueRequest) HasTargetSortCode() bool`

HasTargetSortCode returns a boolean if a field has been set.

### GetThirdLinks

`func (o *CreateIssueRequest) GetThirdLinks() []IssueThirdLinkForm`

GetThirdLinks returns the ThirdLinks field if non-nil, zero value otherwise.

### GetThirdLinksOk

`func (o *CreateIssueRequest) GetThirdLinksOk() (*[]IssueThirdLinkForm, bool)`

GetThirdLinksOk returns a tuple with the ThirdLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdLinks

`func (o *CreateIssueRequest) SetThirdLinks(v []IssueThirdLinkForm)`

SetThirdLinks sets ThirdLinks field to given value.

### HasThirdLinks

`func (o *CreateIssueRequest) HasThirdLinks() bool`

HasThirdLinks returns a boolean if a field has been set.

### GetType

`func (o *CreateIssueRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateIssueRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateIssueRequest) SetType(v string)`

SetType sets Type field to given value.


### GetWatcherIds

`func (o *CreateIssueRequest) GetWatcherIds() []int64`

GetWatcherIds returns the WatcherIds field if non-nil, zero value otherwise.

### GetWatcherIdsOk

`func (o *CreateIssueRequest) GetWatcherIdsOk() (*[]int64, bool)`

GetWatcherIdsOk returns a tuple with the WatcherIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatcherIds

`func (o *CreateIssueRequest) SetWatcherIds(v []int64)`

SetWatcherIds sets WatcherIds field to given value.

### HasWatcherIds

`func (o *CreateIssueRequest) HasWatcherIds() bool`

HasWatcherIds returns a boolean if a field has been set.

### GetWorkingHours

`func (o *CreateIssueRequest) GetWorkingHours() float32`

GetWorkingHours returns the WorkingHours field if non-nil, zero value otherwise.

### GetWorkingHoursOk

`func (o *CreateIssueRequest) GetWorkingHoursOk() (*float32, bool)`

GetWorkingHoursOk returns a tuple with the WorkingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHours

`func (o *CreateIssueRequest) SetWorkingHours(v float32)`

SetWorkingHours sets WorkingHours field to given value.

### HasWorkingHours

`func (o *CreateIssueRequest) HasWorkingHours() bool`

HasWorkingHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


