# ModifyIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssigneeId** | Pointer to **int64** | 处理人 Id | [optional] 
**Comment** | Pointer to **string** | 评论 | [optional] 
**CustomFieldValues** | Pointer to [**[]IssueCustomFieldForm**](IssueCustomFieldForm.md) | 自定义属性值列表 | [optional] 
**DefectTypeId** | Pointer to **int64** | 缺陷类型 Id | [optional] 
**DelFileIds** | Pointer to **[]int64** | 删除的文件 Id 列表 | [optional] 
**DelLabelIds** | Pointer to **[]int64** | 删除的标签 Id 列表 | [optional] 
**DelReleaseCodes** | Pointer to **[]int64** | 需要删除的版本jcode列表 | [optional] 
**DelWatcherIds** | Pointer to **[]int64** | 删除的事项关注人 Id 列表 | [optional] 
**DueDate** | Pointer to **map[string]interface{}** | 截止日期，格式：2021-01-01 | [optional] 
**FileIds** | Pointer to **[]int64** | 添加的文件 Id 列表 | [optional] 
**IssueCode** | **int64** | 事项 Code | 
**IterationCode** | Pointer to **int64** | 迭代code | [optional] 
**LabelIds** | Pointer to **[]int64** | 标签 Id 列表 | [optional] 
**Name** | Pointer to **string** | 事项名称 | [optional] 
**ParentCode** | Pointer to **int64** | 父事项 Code  敏捷模式：Type 为 SUB_TASK 时必须指定  经典模式：Type 为 MISSION、REQUIREMENT 时可指定 | [optional] 
**Priority** | Pointer to **string** | 紧急程度  \&quot;0\&quot; - 低  \&quot;1\&quot; - 中  \&quot;2\&quot; - 高  \&quot;3\&quot; - 紧急 | [optional] 
**Progress** | Pointer to **float32** | 进度 | [optional] 
**ProjectModuleId** | Pointer to **int64** | 项目模块 Id | [optional] 
**ProjectName** | **string** | 项目名称 | 
**RecordHour** | Pointer to [**IssueRecordHourForm**](IssueRecordHourForm.md) |  | [optional] 
**ReleaseCodes** | Pointer to **[]int64** | 版本code列表 | [optional] 
**RequirementTypeId** | Pointer to **int64** | 需求类型 Id | [optional] 
**StartDate** | Pointer to **map[string]interface{}** | 开始日期，格式：2021-01-01 | [optional] 
**StatusId** | Pointer to **int64** | 事项状态 Id | [optional] 
**StoryPoint** | Pointer to **string** | 故事点，例如：0.5、1 | [optional] 
**UpdateLabelIds** | Pointer to **[]int64** | 标签 Id 列表 | [optional] 
**UpdateWatchers** | Pointer to **[]int64** | 关注人 Id 列表 | [optional] 
**WatcherIds** | Pointer to **[]int64** | 添加的事项关注人 Id 列表 | [optional] 
**WorkingHours** | Pointer to **float32** | 工时（小时数） | [optional] 

## Methods

### NewModifyIssueRequest

`func NewModifyIssueRequest(issueCode int64, projectName string, ) *ModifyIssueRequest`

NewModifyIssueRequest instantiates a new ModifyIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyIssueRequestWithDefaults

`func NewModifyIssueRequestWithDefaults() *ModifyIssueRequest`

NewModifyIssueRequestWithDefaults instantiates a new ModifyIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigneeId

`func (o *ModifyIssueRequest) GetAssigneeId() int64`

GetAssigneeId returns the AssigneeId field if non-nil, zero value otherwise.

### GetAssigneeIdOk

`func (o *ModifyIssueRequest) GetAssigneeIdOk() (*int64, bool)`

GetAssigneeIdOk returns a tuple with the AssigneeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigneeId

`func (o *ModifyIssueRequest) SetAssigneeId(v int64)`

SetAssigneeId sets AssigneeId field to given value.

### HasAssigneeId

`func (o *ModifyIssueRequest) HasAssigneeId() bool`

HasAssigneeId returns a boolean if a field has been set.

### GetComment

`func (o *ModifyIssueRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ModifyIssueRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ModifyIssueRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ModifyIssueRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCustomFieldValues

`func (o *ModifyIssueRequest) GetCustomFieldValues() []IssueCustomFieldForm`

GetCustomFieldValues returns the CustomFieldValues field if non-nil, zero value otherwise.

### GetCustomFieldValuesOk

`func (o *ModifyIssueRequest) GetCustomFieldValuesOk() (*[]IssueCustomFieldForm, bool)`

GetCustomFieldValuesOk returns a tuple with the CustomFieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldValues

`func (o *ModifyIssueRequest) SetCustomFieldValues(v []IssueCustomFieldForm)`

SetCustomFieldValues sets CustomFieldValues field to given value.

### HasCustomFieldValues

`func (o *ModifyIssueRequest) HasCustomFieldValues() bool`

HasCustomFieldValues returns a boolean if a field has been set.

### GetDefectTypeId

`func (o *ModifyIssueRequest) GetDefectTypeId() int64`

GetDefectTypeId returns the DefectTypeId field if non-nil, zero value otherwise.

### GetDefectTypeIdOk

`func (o *ModifyIssueRequest) GetDefectTypeIdOk() (*int64, bool)`

GetDefectTypeIdOk returns a tuple with the DefectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectTypeId

`func (o *ModifyIssueRequest) SetDefectTypeId(v int64)`

SetDefectTypeId sets DefectTypeId field to given value.

### HasDefectTypeId

`func (o *ModifyIssueRequest) HasDefectTypeId() bool`

HasDefectTypeId returns a boolean if a field has been set.

### GetDelFileIds

`func (o *ModifyIssueRequest) GetDelFileIds() []int64`

GetDelFileIds returns the DelFileIds field if non-nil, zero value otherwise.

### GetDelFileIdsOk

`func (o *ModifyIssueRequest) GetDelFileIdsOk() (*[]int64, bool)`

GetDelFileIdsOk returns a tuple with the DelFileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelFileIds

`func (o *ModifyIssueRequest) SetDelFileIds(v []int64)`

SetDelFileIds sets DelFileIds field to given value.

### HasDelFileIds

`func (o *ModifyIssueRequest) HasDelFileIds() bool`

HasDelFileIds returns a boolean if a field has been set.

### GetDelLabelIds

`func (o *ModifyIssueRequest) GetDelLabelIds() []int64`

GetDelLabelIds returns the DelLabelIds field if non-nil, zero value otherwise.

### GetDelLabelIdsOk

`func (o *ModifyIssueRequest) GetDelLabelIdsOk() (*[]int64, bool)`

GetDelLabelIdsOk returns a tuple with the DelLabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelLabelIds

`func (o *ModifyIssueRequest) SetDelLabelIds(v []int64)`

SetDelLabelIds sets DelLabelIds field to given value.

### HasDelLabelIds

`func (o *ModifyIssueRequest) HasDelLabelIds() bool`

HasDelLabelIds returns a boolean if a field has been set.

### GetDelReleaseCodes

`func (o *ModifyIssueRequest) GetDelReleaseCodes() []int64`

GetDelReleaseCodes returns the DelReleaseCodes field if non-nil, zero value otherwise.

### GetDelReleaseCodesOk

`func (o *ModifyIssueRequest) GetDelReleaseCodesOk() (*[]int64, bool)`

GetDelReleaseCodesOk returns a tuple with the DelReleaseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelReleaseCodes

`func (o *ModifyIssueRequest) SetDelReleaseCodes(v []int64)`

SetDelReleaseCodes sets DelReleaseCodes field to given value.

### HasDelReleaseCodes

`func (o *ModifyIssueRequest) HasDelReleaseCodes() bool`

HasDelReleaseCodes returns a boolean if a field has been set.

### GetDelWatcherIds

`func (o *ModifyIssueRequest) GetDelWatcherIds() []int64`

GetDelWatcherIds returns the DelWatcherIds field if non-nil, zero value otherwise.

### GetDelWatcherIdsOk

`func (o *ModifyIssueRequest) GetDelWatcherIdsOk() (*[]int64, bool)`

GetDelWatcherIdsOk returns a tuple with the DelWatcherIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelWatcherIds

`func (o *ModifyIssueRequest) SetDelWatcherIds(v []int64)`

SetDelWatcherIds sets DelWatcherIds field to given value.

### HasDelWatcherIds

`func (o *ModifyIssueRequest) HasDelWatcherIds() bool`

HasDelWatcherIds returns a boolean if a field has been set.

### GetDueDate

`func (o *ModifyIssueRequest) GetDueDate() map[string]interface{}`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *ModifyIssueRequest) GetDueDateOk() (*map[string]interface{}, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *ModifyIssueRequest) SetDueDate(v map[string]interface{})`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *ModifyIssueRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetFileIds

`func (o *ModifyIssueRequest) GetFileIds() []int64`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *ModifyIssueRequest) GetFileIdsOk() (*[]int64, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *ModifyIssueRequest) SetFileIds(v []int64)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *ModifyIssueRequest) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.

### GetIssueCode

`func (o *ModifyIssueRequest) GetIssueCode() int64`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *ModifyIssueRequest) GetIssueCodeOk() (*int64, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *ModifyIssueRequest) SetIssueCode(v int64)`

SetIssueCode sets IssueCode field to given value.


### GetIterationCode

`func (o *ModifyIssueRequest) GetIterationCode() int64`

GetIterationCode returns the IterationCode field if non-nil, zero value otherwise.

### GetIterationCodeOk

`func (o *ModifyIssueRequest) GetIterationCodeOk() (*int64, bool)`

GetIterationCodeOk returns a tuple with the IterationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCode

`func (o *ModifyIssueRequest) SetIterationCode(v int64)`

SetIterationCode sets IterationCode field to given value.

### HasIterationCode

`func (o *ModifyIssueRequest) HasIterationCode() bool`

HasIterationCode returns a boolean if a field has been set.

### GetLabelIds

`func (o *ModifyIssueRequest) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *ModifyIssueRequest) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *ModifyIssueRequest) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *ModifyIssueRequest) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetName

`func (o *ModifyIssueRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyIssueRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyIssueRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModifyIssueRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentCode

`func (o *ModifyIssueRequest) GetParentCode() int64`

GetParentCode returns the ParentCode field if non-nil, zero value otherwise.

### GetParentCodeOk

`func (o *ModifyIssueRequest) GetParentCodeOk() (*int64, bool)`

GetParentCodeOk returns a tuple with the ParentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCode

`func (o *ModifyIssueRequest) SetParentCode(v int64)`

SetParentCode sets ParentCode field to given value.

### HasParentCode

`func (o *ModifyIssueRequest) HasParentCode() bool`

HasParentCode returns a boolean if a field has been set.

### GetPriority

`func (o *ModifyIssueRequest) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ModifyIssueRequest) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ModifyIssueRequest) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ModifyIssueRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProgress

`func (o *ModifyIssueRequest) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ModifyIssueRequest) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ModifyIssueRequest) SetProgress(v float32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ModifyIssueRequest) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetProjectModuleId

`func (o *ModifyIssueRequest) GetProjectModuleId() int64`

GetProjectModuleId returns the ProjectModuleId field if non-nil, zero value otherwise.

### GetProjectModuleIdOk

`func (o *ModifyIssueRequest) GetProjectModuleIdOk() (*int64, bool)`

GetProjectModuleIdOk returns a tuple with the ProjectModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectModuleId

`func (o *ModifyIssueRequest) SetProjectModuleId(v int64)`

SetProjectModuleId sets ProjectModuleId field to given value.

### HasProjectModuleId

`func (o *ModifyIssueRequest) HasProjectModuleId() bool`

HasProjectModuleId returns a boolean if a field has been set.

### GetProjectName

`func (o *ModifyIssueRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ModifyIssueRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ModifyIssueRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetRecordHour

`func (o *ModifyIssueRequest) GetRecordHour() IssueRecordHourForm`

GetRecordHour returns the RecordHour field if non-nil, zero value otherwise.

### GetRecordHourOk

`func (o *ModifyIssueRequest) GetRecordHourOk() (*IssueRecordHourForm, bool)`

GetRecordHourOk returns a tuple with the RecordHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordHour

`func (o *ModifyIssueRequest) SetRecordHour(v IssueRecordHourForm)`

SetRecordHour sets RecordHour field to given value.

### HasRecordHour

`func (o *ModifyIssueRequest) HasRecordHour() bool`

HasRecordHour returns a boolean if a field has been set.

### GetReleaseCodes

`func (o *ModifyIssueRequest) GetReleaseCodes() []int64`

GetReleaseCodes returns the ReleaseCodes field if non-nil, zero value otherwise.

### GetReleaseCodesOk

`func (o *ModifyIssueRequest) GetReleaseCodesOk() (*[]int64, bool)`

GetReleaseCodesOk returns a tuple with the ReleaseCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseCodes

`func (o *ModifyIssueRequest) SetReleaseCodes(v []int64)`

SetReleaseCodes sets ReleaseCodes field to given value.

### HasReleaseCodes

`func (o *ModifyIssueRequest) HasReleaseCodes() bool`

HasReleaseCodes returns a boolean if a field has been set.

### GetRequirementTypeId

`func (o *ModifyIssueRequest) GetRequirementTypeId() int64`

GetRequirementTypeId returns the RequirementTypeId field if non-nil, zero value otherwise.

### GetRequirementTypeIdOk

`func (o *ModifyIssueRequest) GetRequirementTypeIdOk() (*int64, bool)`

GetRequirementTypeIdOk returns a tuple with the RequirementTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementTypeId

`func (o *ModifyIssueRequest) SetRequirementTypeId(v int64)`

SetRequirementTypeId sets RequirementTypeId field to given value.

### HasRequirementTypeId

`func (o *ModifyIssueRequest) HasRequirementTypeId() bool`

HasRequirementTypeId returns a boolean if a field has been set.

### GetStartDate

`func (o *ModifyIssueRequest) GetStartDate() map[string]interface{}`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ModifyIssueRequest) GetStartDateOk() (*map[string]interface{}, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ModifyIssueRequest) SetStartDate(v map[string]interface{})`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ModifyIssueRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatusId

`func (o *ModifyIssueRequest) GetStatusId() int64`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *ModifyIssueRequest) GetStatusIdOk() (*int64, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *ModifyIssueRequest) SetStatusId(v int64)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *ModifyIssueRequest) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetStoryPoint

`func (o *ModifyIssueRequest) GetStoryPoint() string`

GetStoryPoint returns the StoryPoint field if non-nil, zero value otherwise.

### GetStoryPointOk

`func (o *ModifyIssueRequest) GetStoryPointOk() (*string, bool)`

GetStoryPointOk returns a tuple with the StoryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryPoint

`func (o *ModifyIssueRequest) SetStoryPoint(v string)`

SetStoryPoint sets StoryPoint field to given value.

### HasStoryPoint

`func (o *ModifyIssueRequest) HasStoryPoint() bool`

HasStoryPoint returns a boolean if a field has been set.

### GetUpdateLabelIds

`func (o *ModifyIssueRequest) GetUpdateLabelIds() []int64`

GetUpdateLabelIds returns the UpdateLabelIds field if non-nil, zero value otherwise.

### GetUpdateLabelIdsOk

`func (o *ModifyIssueRequest) GetUpdateLabelIdsOk() (*[]int64, bool)`

GetUpdateLabelIdsOk returns a tuple with the UpdateLabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateLabelIds

`func (o *ModifyIssueRequest) SetUpdateLabelIds(v []int64)`

SetUpdateLabelIds sets UpdateLabelIds field to given value.

### HasUpdateLabelIds

`func (o *ModifyIssueRequest) HasUpdateLabelIds() bool`

HasUpdateLabelIds returns a boolean if a field has been set.

### GetUpdateWatchers

`func (o *ModifyIssueRequest) GetUpdateWatchers() []int64`

GetUpdateWatchers returns the UpdateWatchers field if non-nil, zero value otherwise.

### GetUpdateWatchersOk

`func (o *ModifyIssueRequest) GetUpdateWatchersOk() (*[]int64, bool)`

GetUpdateWatchersOk returns a tuple with the UpdateWatchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateWatchers

`func (o *ModifyIssueRequest) SetUpdateWatchers(v []int64)`

SetUpdateWatchers sets UpdateWatchers field to given value.

### HasUpdateWatchers

`func (o *ModifyIssueRequest) HasUpdateWatchers() bool`

HasUpdateWatchers returns a boolean if a field has been set.

### GetWatcherIds

`func (o *ModifyIssueRequest) GetWatcherIds() []int64`

GetWatcherIds returns the WatcherIds field if non-nil, zero value otherwise.

### GetWatcherIdsOk

`func (o *ModifyIssueRequest) GetWatcherIdsOk() (*[]int64, bool)`

GetWatcherIdsOk returns a tuple with the WatcherIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatcherIds

`func (o *ModifyIssueRequest) SetWatcherIds(v []int64)`

SetWatcherIds sets WatcherIds field to given value.

### HasWatcherIds

`func (o *ModifyIssueRequest) HasWatcherIds() bool`

HasWatcherIds returns a boolean if a field has been set.

### GetWorkingHours

`func (o *ModifyIssueRequest) GetWorkingHours() float32`

GetWorkingHours returns the WorkingHours field if non-nil, zero value otherwise.

### GetWorkingHoursOk

`func (o *ModifyIssueRequest) GetWorkingHoursOk() (*float32, bool)`

GetWorkingHoursOk returns a tuple with the WorkingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingHours

`func (o *ModifyIssueRequest) SetWorkingHours(v float32)`

SetWorkingHours sets WorkingHours field to given value.

### HasWorkingHours

`func (o *ModifyIssueRequest) HasWorkingHours() bool`

HasWorkingHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


