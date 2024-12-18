/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"bytes"
	"fmt"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ModifyIssueRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyIssueRequest{}

// ModifyIssueRequest struct for ModifyIssueRequest
type ModifyIssueRequest struct {
	// 处理人 Id
	AssigneeId *int64 `json:"AssigneeId,omitempty"`
	// 评论
	Comment *string `json:"Comment,omitempty"`
	// 自定义属性值列表
	CustomFieldValues []IssueCustomFieldForm `json:"CustomFieldValues,omitempty"`
	// 缺陷类型 Id
	DefectTypeId *int64 `json:"DefectTypeId,omitempty"`
	// 删除的文件 Id 列表
	DelFileIds []int64 `json:"DelFileIds,omitempty"`
	// 删除的标签 Id 列表
	DelLabelIds []int64 `json:"DelLabelIds,omitempty"`
	// 需要删除的版本jcode列表
	DelReleaseCodes []int64 `json:"DelReleaseCodes,omitempty"`
	// 删除的事项关注人 Id 列表
	DelWatcherIds []int64 `json:"DelWatcherIds,omitempty"`
	// 截止日期，格式：2021-01-01
	DueDate map[string]interface{} `json:"DueDate,omitempty"`
	// 添加的文件 Id 列表
	FileIds []int64 `json:"FileIds,omitempty"`
	// 事项 Code
	IssueCode int64 `json:"IssueCode"`
	// 迭代code
	IterationCode *int64 `json:"IterationCode,omitempty"`
	// 标签 Id 列表
	LabelIds []int64 `json:"LabelIds,omitempty"`
	// 事项名称
	Name *string `json:"Name,omitempty"`
	// 父事项 Code  敏捷模式：Type 为 SUB_TASK 时必须指定  经典模式：Type 为 MISSION、REQUIREMENT 时可指定
	ParentCode *int64 `json:"ParentCode,omitempty"`
	// 紧急程度  \"0\" - 低  \"1\" - 中  \"2\" - 高  \"3\" - 紧急
	Priority *string `json:"Priority,omitempty"`
	// 进度
	Progress *float32 `json:"Progress,omitempty"`
	// 项目模块 Id
	ProjectModuleId *int64 `json:"ProjectModuleId,omitempty"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	RecordHour *IssueRecordHourForm `json:"RecordHour,omitempty"`
	// 版本code列表
	ReleaseCodes []int64 `json:"ReleaseCodes,omitempty"`
	// 需求类型 Id
	RequirementTypeId *int64 `json:"RequirementTypeId,omitempty"`
	// 开始日期，格式：2021-01-01
	StartDate map[string]interface{} `json:"StartDate,omitempty"`
	// 事项状态 Id
	StatusId *int64 `json:"StatusId,omitempty"`
	// 故事点，例如：0.5、1
	StoryPoint *string `json:"StoryPoint,omitempty"`
	// 标签 Id 列表
	UpdateLabelIds []int64 `json:"UpdateLabelIds,omitempty"`
	// 关注人 Id 列表
	UpdateWatchers []int64 `json:"UpdateWatchers,omitempty"`
	// 添加的事项关注人 Id 列表
	WatcherIds []int64 `json:"WatcherIds,omitempty"`
	// 工时（小时数）
	WorkingHours *float32 `json:"WorkingHours,omitempty"`
}

type _ModifyIssueRequest ModifyIssueRequest

// NewModifyIssueRequest instantiates a new ModifyIssueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyIssueRequest(issueCode int64, projectName string) *ModifyIssueRequest {
	this := ModifyIssueRequest{}
	this.IssueCode = issueCode
	this.ProjectName = projectName
	return &this
}

// NewModifyIssueRequestWithDefaults instantiates a new ModifyIssueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyIssueRequestWithDefaults() *ModifyIssueRequest {
	this := ModifyIssueRequest{}
	return &this
}

// GetAssigneeId returns the AssigneeId field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetAssigneeId() int64 {
	if o == nil || utils.IsNil(o.AssigneeId) {
		var ret int64
		return ret
	}
	return *o.AssigneeId
}

// GetAssigneeIdOk returns a tuple with the AssigneeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetAssigneeIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.AssigneeId) {
		return nil, false
	}
	return o.AssigneeId, true
}

// HasAssigneeId returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasAssigneeId() bool {
	if o != nil && !utils.IsNil(o.AssigneeId) {
		return true
	}

	return false
}

// SetAssigneeId gets a reference to the given int64 and assigns it to the AssigneeId field.
func (o *ModifyIssueRequest) SetAssigneeId(v int64) {
	o.AssigneeId = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetComment() string {
	if o == nil || utils.IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetCommentOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasComment() bool {
	if o != nil && !utils.IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ModifyIssueRequest) SetComment(v string) {
	o.Comment = &v
}

// GetCustomFieldValues returns the CustomFieldValues field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetCustomFieldValues() []IssueCustomFieldForm {
	if o == nil || utils.IsNil(o.CustomFieldValues) {
		var ret []IssueCustomFieldForm
		return ret
	}
	return o.CustomFieldValues
}

// GetCustomFieldValuesOk returns a tuple with the CustomFieldValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetCustomFieldValuesOk() ([]IssueCustomFieldForm, bool) {
	if o == nil || utils.IsNil(o.CustomFieldValues) {
		return nil, false
	}
	return o.CustomFieldValues, true
}

// HasCustomFieldValues returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasCustomFieldValues() bool {
	if o != nil && !utils.IsNil(o.CustomFieldValues) {
		return true
	}

	return false
}

// SetCustomFieldValues gets a reference to the given []IssueCustomFieldForm and assigns it to the CustomFieldValues field.
func (o *ModifyIssueRequest) SetCustomFieldValues(v []IssueCustomFieldForm) {
	o.CustomFieldValues = v
}

// GetDefectTypeId returns the DefectTypeId field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetDefectTypeId() int64 {
	if o == nil || utils.IsNil(o.DefectTypeId) {
		var ret int64
		return ret
	}
	return *o.DefectTypeId
}

// GetDefectTypeIdOk returns a tuple with the DefectTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetDefectTypeIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.DefectTypeId) {
		return nil, false
	}
	return o.DefectTypeId, true
}

// HasDefectTypeId returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasDefectTypeId() bool {
	if o != nil && !utils.IsNil(o.DefectTypeId) {
		return true
	}

	return false
}

// SetDefectTypeId gets a reference to the given int64 and assigns it to the DefectTypeId field.
func (o *ModifyIssueRequest) SetDefectTypeId(v int64) {
	o.DefectTypeId = &v
}

// GetDelFileIds returns the DelFileIds field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetDelFileIds() []int64 {
	if o == nil || utils.IsNil(o.DelFileIds) {
		var ret []int64
		return ret
	}
	return o.DelFileIds
}

// GetDelFileIdsOk returns a tuple with the DelFileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetDelFileIdsOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.DelFileIds) {
		return nil, false
	}
	return o.DelFileIds, true
}

// HasDelFileIds returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasDelFileIds() bool {
	if o != nil && !utils.IsNil(o.DelFileIds) {
		return true
	}

	return false
}

// SetDelFileIds gets a reference to the given []int64 and assigns it to the DelFileIds field.
func (o *ModifyIssueRequest) SetDelFileIds(v []int64) {
	o.DelFileIds = v
}

// GetDelLabelIds returns the DelLabelIds field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetDelLabelIds() []int64 {
	if o == nil || utils.IsNil(o.DelLabelIds) {
		var ret []int64
		return ret
	}
	return o.DelLabelIds
}

// GetDelLabelIdsOk returns a tuple with the DelLabelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetDelLabelIdsOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.DelLabelIds) {
		return nil, false
	}
	return o.DelLabelIds, true
}

// HasDelLabelIds returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasDelLabelIds() bool {
	if o != nil && !utils.IsNil(o.DelLabelIds) {
		return true
	}

	return false
}

// SetDelLabelIds gets a reference to the given []int64 and assigns it to the DelLabelIds field.
func (o *ModifyIssueRequest) SetDelLabelIds(v []int64) {
	o.DelLabelIds = v
}

// GetDelReleaseCodes returns the DelReleaseCodes field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetDelReleaseCodes() []int64 {
	if o == nil || utils.IsNil(o.DelReleaseCodes) {
		var ret []int64
		return ret
	}
	return o.DelReleaseCodes
}

// GetDelReleaseCodesOk returns a tuple with the DelReleaseCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetDelReleaseCodesOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.DelReleaseCodes) {
		return nil, false
	}
	return o.DelReleaseCodes, true
}

// HasDelReleaseCodes returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasDelReleaseCodes() bool {
	if o != nil && !utils.IsNil(o.DelReleaseCodes) {
		return true
	}

	return false
}

// SetDelReleaseCodes gets a reference to the given []int64 and assigns it to the DelReleaseCodes field.
func (o *ModifyIssueRequest) SetDelReleaseCodes(v []int64) {
	o.DelReleaseCodes = v
}

// GetDelWatcherIds returns the DelWatcherIds field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetDelWatcherIds() []int64 {
	if o == nil || utils.IsNil(o.DelWatcherIds) {
		var ret []int64
		return ret
	}
	return o.DelWatcherIds
}

// GetDelWatcherIdsOk returns a tuple with the DelWatcherIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetDelWatcherIdsOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.DelWatcherIds) {
		return nil, false
	}
	return o.DelWatcherIds, true
}

// HasDelWatcherIds returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasDelWatcherIds() bool {
	if o != nil && !utils.IsNil(o.DelWatcherIds) {
		return true
	}

	return false
}

// SetDelWatcherIds gets a reference to the given []int64 and assigns it to the DelWatcherIds field.
func (o *ModifyIssueRequest) SetDelWatcherIds(v []int64) {
	o.DelWatcherIds = v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetDueDate() map[string]interface{} {
	if o == nil || utils.IsNil(o.DueDate) {
		var ret map[string]interface{}
		return ret
	}
	return o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetDueDateOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.DueDate) {
		return map[string]interface{}{}, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasDueDate() bool {
	if o != nil && !utils.IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given map[string]interface{} and assigns it to the DueDate field.
func (o *ModifyIssueRequest) SetDueDate(v map[string]interface{}) {
	o.DueDate = v
}

// GetFileIds returns the FileIds field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetFileIds() []int64 {
	if o == nil || utils.IsNil(o.FileIds) {
		var ret []int64
		return ret
	}
	return o.FileIds
}

// GetFileIdsOk returns a tuple with the FileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetFileIdsOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.FileIds) {
		return nil, false
	}
	return o.FileIds, true
}

// HasFileIds returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasFileIds() bool {
	if o != nil && !utils.IsNil(o.FileIds) {
		return true
	}

	return false
}

// SetFileIds gets a reference to the given []int64 and assigns it to the FileIds field.
func (o *ModifyIssueRequest) SetFileIds(v []int64) {
	o.FileIds = v
}

// GetIssueCode returns the IssueCode field value
func (o *ModifyIssueRequest) GetIssueCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IssueCode
}

// GetIssueCodeOk returns a tuple with the IssueCode field value
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetIssueCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueCode, true
}

// SetIssueCode sets field value
func (o *ModifyIssueRequest) SetIssueCode(v int64) {
	o.IssueCode = v
}

// GetIterationCode returns the IterationCode field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetIterationCode() int64 {
	if o == nil || utils.IsNil(o.IterationCode) {
		var ret int64
		return ret
	}
	return *o.IterationCode
}

// GetIterationCodeOk returns a tuple with the IterationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetIterationCodeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.IterationCode) {
		return nil, false
	}
	return o.IterationCode, true
}

// HasIterationCode returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasIterationCode() bool {
	if o != nil && !utils.IsNil(o.IterationCode) {
		return true
	}

	return false
}

// SetIterationCode gets a reference to the given int64 and assigns it to the IterationCode field.
func (o *ModifyIssueRequest) SetIterationCode(v int64) {
	o.IterationCode = &v
}

// GetLabelIds returns the LabelIds field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetLabelIds() []int64 {
	if o == nil || utils.IsNil(o.LabelIds) {
		var ret []int64
		return ret
	}
	return o.LabelIds
}

// GetLabelIdsOk returns a tuple with the LabelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetLabelIdsOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.LabelIds) {
		return nil, false
	}
	return o.LabelIds, true
}

// HasLabelIds returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasLabelIds() bool {
	if o != nil && !utils.IsNil(o.LabelIds) {
		return true
	}

	return false
}

// SetLabelIds gets a reference to the given []int64 and assigns it to the LabelIds field.
func (o *ModifyIssueRequest) SetLabelIds(v []int64) {
	o.LabelIds = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModifyIssueRequest) SetName(v string) {
	o.Name = &v
}

// GetParentCode returns the ParentCode field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetParentCode() int64 {
	if o == nil || utils.IsNil(o.ParentCode) {
		var ret int64
		return ret
	}
	return *o.ParentCode
}

// GetParentCodeOk returns a tuple with the ParentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetParentCodeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.ParentCode) {
		return nil, false
	}
	return o.ParentCode, true
}

// HasParentCode returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasParentCode() bool {
	if o != nil && !utils.IsNil(o.ParentCode) {
		return true
	}

	return false
}

// SetParentCode gets a reference to the given int64 and assigns it to the ParentCode field.
func (o *ModifyIssueRequest) SetParentCode(v int64) {
	o.ParentCode = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetPriority() string {
	if o == nil || utils.IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetPriorityOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasPriority() bool {
	if o != nil && !utils.IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *ModifyIssueRequest) SetPriority(v string) {
	o.Priority = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetProgress() float32 {
	if o == nil || utils.IsNil(o.Progress) {
		var ret float32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetProgressOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasProgress() bool {
	if o != nil && !utils.IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given float32 and assigns it to the Progress field.
func (o *ModifyIssueRequest) SetProgress(v float32) {
	o.Progress = &v
}

// GetProjectModuleId returns the ProjectModuleId field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetProjectModuleId() int64 {
	if o == nil || utils.IsNil(o.ProjectModuleId) {
		var ret int64
		return ret
	}
	return *o.ProjectModuleId
}

// GetProjectModuleIdOk returns a tuple with the ProjectModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetProjectModuleIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.ProjectModuleId) {
		return nil, false
	}
	return o.ProjectModuleId, true
}

// HasProjectModuleId returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasProjectModuleId() bool {
	if o != nil && !utils.IsNil(o.ProjectModuleId) {
		return true
	}

	return false
}

// SetProjectModuleId gets a reference to the given int64 and assigns it to the ProjectModuleId field.
func (o *ModifyIssueRequest) SetProjectModuleId(v int64) {
	o.ProjectModuleId = &v
}

// GetProjectName returns the ProjectName field value
func (o *ModifyIssueRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *ModifyIssueRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetRecordHour returns the RecordHour field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetRecordHour() IssueRecordHourForm {
	if o == nil || utils.IsNil(o.RecordHour) {
		var ret IssueRecordHourForm
		return ret
	}
	return *o.RecordHour
}

// GetRecordHourOk returns a tuple with the RecordHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetRecordHourOk() (*IssueRecordHourForm, bool) {
	if o == nil || utils.IsNil(o.RecordHour) {
		return nil, false
	}
	return o.RecordHour, true
}

// HasRecordHour returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasRecordHour() bool {
	if o != nil && !utils.IsNil(o.RecordHour) {
		return true
	}

	return false
}

// SetRecordHour gets a reference to the given IssueRecordHourForm and assigns it to the RecordHour field.
func (o *ModifyIssueRequest) SetRecordHour(v IssueRecordHourForm) {
	o.RecordHour = &v
}

// GetReleaseCodes returns the ReleaseCodes field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetReleaseCodes() []int64 {
	if o == nil || utils.IsNil(o.ReleaseCodes) {
		var ret []int64
		return ret
	}
	return o.ReleaseCodes
}

// GetReleaseCodesOk returns a tuple with the ReleaseCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetReleaseCodesOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.ReleaseCodes) {
		return nil, false
	}
	return o.ReleaseCodes, true
}

// HasReleaseCodes returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasReleaseCodes() bool {
	if o != nil && !utils.IsNil(o.ReleaseCodes) {
		return true
	}

	return false
}

// SetReleaseCodes gets a reference to the given []int64 and assigns it to the ReleaseCodes field.
func (o *ModifyIssueRequest) SetReleaseCodes(v []int64) {
	o.ReleaseCodes = v
}

// GetRequirementTypeId returns the RequirementTypeId field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetRequirementTypeId() int64 {
	if o == nil || utils.IsNil(o.RequirementTypeId) {
		var ret int64
		return ret
	}
	return *o.RequirementTypeId
}

// GetRequirementTypeIdOk returns a tuple with the RequirementTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetRequirementTypeIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.RequirementTypeId) {
		return nil, false
	}
	return o.RequirementTypeId, true
}

// HasRequirementTypeId returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasRequirementTypeId() bool {
	if o != nil && !utils.IsNil(o.RequirementTypeId) {
		return true
	}

	return false
}

// SetRequirementTypeId gets a reference to the given int64 and assigns it to the RequirementTypeId field.
func (o *ModifyIssueRequest) SetRequirementTypeId(v int64) {
	o.RequirementTypeId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetStartDate() map[string]interface{} {
	if o == nil || utils.IsNil(o.StartDate) {
		var ret map[string]interface{}
		return ret
	}
	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetStartDateOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.StartDate) {
		return map[string]interface{}{}, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasStartDate() bool {
	if o != nil && !utils.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given map[string]interface{} and assigns it to the StartDate field.
func (o *ModifyIssueRequest) SetStartDate(v map[string]interface{}) {
	o.StartDate = v
}

// GetStatusId returns the StatusId field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetStatusId() int64 {
	if o == nil || utils.IsNil(o.StatusId) {
		var ret int64
		return ret
	}
	return *o.StatusId
}

// GetStatusIdOk returns a tuple with the StatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetStatusIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.StatusId) {
		return nil, false
	}
	return o.StatusId, true
}

// HasStatusId returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasStatusId() bool {
	if o != nil && !utils.IsNil(o.StatusId) {
		return true
	}

	return false
}

// SetStatusId gets a reference to the given int64 and assigns it to the StatusId field.
func (o *ModifyIssueRequest) SetStatusId(v int64) {
	o.StatusId = &v
}

// GetStoryPoint returns the StoryPoint field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetStoryPoint() string {
	if o == nil || utils.IsNil(o.StoryPoint) {
		var ret string
		return ret
	}
	return *o.StoryPoint
}

// GetStoryPointOk returns a tuple with the StoryPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetStoryPointOk() (*string, bool) {
	if o == nil || utils.IsNil(o.StoryPoint) {
		return nil, false
	}
	return o.StoryPoint, true
}

// HasStoryPoint returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasStoryPoint() bool {
	if o != nil && !utils.IsNil(o.StoryPoint) {
		return true
	}

	return false
}

// SetStoryPoint gets a reference to the given string and assigns it to the StoryPoint field.
func (o *ModifyIssueRequest) SetStoryPoint(v string) {
	o.StoryPoint = &v
}

// GetUpdateLabelIds returns the UpdateLabelIds field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetUpdateLabelIds() []int64 {
	if o == nil || utils.IsNil(o.UpdateLabelIds) {
		var ret []int64
		return ret
	}
	return o.UpdateLabelIds
}

// GetUpdateLabelIdsOk returns a tuple with the UpdateLabelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetUpdateLabelIdsOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.UpdateLabelIds) {
		return nil, false
	}
	return o.UpdateLabelIds, true
}

// HasUpdateLabelIds returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasUpdateLabelIds() bool {
	if o != nil && !utils.IsNil(o.UpdateLabelIds) {
		return true
	}

	return false
}

// SetUpdateLabelIds gets a reference to the given []int64 and assigns it to the UpdateLabelIds field.
func (o *ModifyIssueRequest) SetUpdateLabelIds(v []int64) {
	o.UpdateLabelIds = v
}

// GetUpdateWatchers returns the UpdateWatchers field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetUpdateWatchers() []int64 {
	if o == nil || utils.IsNil(o.UpdateWatchers) {
		var ret []int64
		return ret
	}
	return o.UpdateWatchers
}

// GetUpdateWatchersOk returns a tuple with the UpdateWatchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetUpdateWatchersOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.UpdateWatchers) {
		return nil, false
	}
	return o.UpdateWatchers, true
}

// HasUpdateWatchers returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasUpdateWatchers() bool {
	if o != nil && !utils.IsNil(o.UpdateWatchers) {
		return true
	}

	return false
}

// SetUpdateWatchers gets a reference to the given []int64 and assigns it to the UpdateWatchers field.
func (o *ModifyIssueRequest) SetUpdateWatchers(v []int64) {
	o.UpdateWatchers = v
}

// GetWatcherIds returns the WatcherIds field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetWatcherIds() []int64 {
	if o == nil || utils.IsNil(o.WatcherIds) {
		var ret []int64
		return ret
	}
	return o.WatcherIds
}

// GetWatcherIdsOk returns a tuple with the WatcherIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetWatcherIdsOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.WatcherIds) {
		return nil, false
	}
	return o.WatcherIds, true
}

// HasWatcherIds returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasWatcherIds() bool {
	if o != nil && !utils.IsNil(o.WatcherIds) {
		return true
	}

	return false
}

// SetWatcherIds gets a reference to the given []int64 and assigns it to the WatcherIds field.
func (o *ModifyIssueRequest) SetWatcherIds(v []int64) {
	o.WatcherIds = v
}

// GetWorkingHours returns the WorkingHours field value if set, zero value otherwise.
func (o *ModifyIssueRequest) GetWorkingHours() float32 {
	if o == nil || utils.IsNil(o.WorkingHours) {
		var ret float32
		return ret
	}
	return *o.WorkingHours
}

// GetWorkingHoursOk returns a tuple with the WorkingHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyIssueRequest) GetWorkingHoursOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.WorkingHours) {
		return nil, false
	}
	return o.WorkingHours, true
}

// HasWorkingHours returns a boolean if a field has been set.
func (o *ModifyIssueRequest) HasWorkingHours() bool {
	if o != nil && !utils.IsNil(o.WorkingHours) {
		return true
	}

	return false
}

// SetWorkingHours gets a reference to the given float32 and assigns it to the WorkingHours field.
func (o *ModifyIssueRequest) SetWorkingHours(v float32) {
	o.WorkingHours = &v
}

func (o ModifyIssueRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyIssueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AssigneeId) {
		toSerialize["AssigneeId"] = o.AssigneeId
	}
	if !utils.IsNil(o.Comment) {
		toSerialize["Comment"] = o.Comment
	}
	if !utils.IsNil(o.CustomFieldValues) {
		toSerialize["CustomFieldValues"] = o.CustomFieldValues
	}
	if !utils.IsNil(o.DefectTypeId) {
		toSerialize["DefectTypeId"] = o.DefectTypeId
	}
	if !utils.IsNil(o.DelFileIds) {
		toSerialize["DelFileIds"] = o.DelFileIds
	}
	if !utils.IsNil(o.DelLabelIds) {
		toSerialize["DelLabelIds"] = o.DelLabelIds
	}
	if !utils.IsNil(o.DelReleaseCodes) {
		toSerialize["DelReleaseCodes"] = o.DelReleaseCodes
	}
	if !utils.IsNil(o.DelWatcherIds) {
		toSerialize["DelWatcherIds"] = o.DelWatcherIds
	}
	if !utils.IsNil(o.DueDate) {
		toSerialize["DueDate"] = o.DueDate
	}
	if !utils.IsNil(o.FileIds) {
		toSerialize["FileIds"] = o.FileIds
	}
	toSerialize["IssueCode"] = o.IssueCode
	if !utils.IsNil(o.IterationCode) {
		toSerialize["IterationCode"] = o.IterationCode
	}
	if !utils.IsNil(o.LabelIds) {
		toSerialize["LabelIds"] = o.LabelIds
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !utils.IsNil(o.ParentCode) {
		toSerialize["ParentCode"] = o.ParentCode
	}
	if !utils.IsNil(o.Priority) {
		toSerialize["Priority"] = o.Priority
	}
	if !utils.IsNil(o.Progress) {
		toSerialize["Progress"] = o.Progress
	}
	if !utils.IsNil(o.ProjectModuleId) {
		toSerialize["ProjectModuleId"] = o.ProjectModuleId
	}
	toSerialize["ProjectName"] = o.ProjectName
	if !utils.IsNil(o.RecordHour) {
		toSerialize["RecordHour"] = o.RecordHour
	}
	if !utils.IsNil(o.ReleaseCodes) {
		toSerialize["ReleaseCodes"] = o.ReleaseCodes
	}
	if !utils.IsNil(o.RequirementTypeId) {
		toSerialize["RequirementTypeId"] = o.RequirementTypeId
	}
	if !utils.IsNil(o.StartDate) {
		toSerialize["StartDate"] = o.StartDate
	}
	if !utils.IsNil(o.StatusId) {
		toSerialize["StatusId"] = o.StatusId
	}
	if !utils.IsNil(o.StoryPoint) {
		toSerialize["StoryPoint"] = o.StoryPoint
	}
	if !utils.IsNil(o.UpdateLabelIds) {
		toSerialize["UpdateLabelIds"] = o.UpdateLabelIds
	}
	if !utils.IsNil(o.UpdateWatchers) {
		toSerialize["UpdateWatchers"] = o.UpdateWatchers
	}
	if !utils.IsNil(o.WatcherIds) {
		toSerialize["WatcherIds"] = o.WatcherIds
	}
	if !utils.IsNil(o.WorkingHours) {
		toSerialize["WorkingHours"] = o.WorkingHours
	}
	return toSerialize, nil
}

func (o *ModifyIssueRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"IssueCode",
		"ProjectName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varModifyIssueRequest := _ModifyIssueRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyIssueRequest)

	if err != nil {
		return err
	}

	*o = ModifyIssueRequest(varModifyIssueRequest)

	return err
}

type NullableModifyIssueRequest struct {
	value *ModifyIssueRequest
	isSet bool
}

func (v NullableModifyIssueRequest) Get() *ModifyIssueRequest {
	return v.value
}

func (v *NullableModifyIssueRequest) Set(val *ModifyIssueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyIssueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyIssueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyIssueRequest(val *ModifyIssueRequest) *NullableModifyIssueRequest {
	return &NullableModifyIssueRequest{value: val, isSet: true}
}

func (v NullableModifyIssueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyIssueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


