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

// checks if the CreateIssueRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateIssueRequest{}

// CreateIssueRequest struct for CreateIssueRequest
type CreateIssueRequest struct {
	// 处理人 Id
	AssigneeId *int64 `json:"AssigneeId,omitempty"`
	// 自定义属性值列表
	CustomFieldValues []IssueCustomFieldForm `json:"CustomFieldValues,omitempty"`
	// 缺陷类型 Id
	DefectTypeId *int64 `json:"DefectTypeId,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty"`
	// 截止日期，格式：2021-01-01
	DueDate map[string]interface{} `json:"DueDate,omitempty"`
	// 史诗 Code，Type 为 EPIC 或 SUB_TASK时，不传该值
	EpicCode *int64 `json:"EpicCode,omitempty"`
	// 附件的文件 Id 列表
	FileIds []int64 `json:"FileIds,omitempty"`
	// 事项类型 Id
	IssueTypeId *int64 `json:"IssueTypeId,omitempty"`
	// 迭代 Code，Type 为 EPIC 或 SUB_TASK时，忽略该值
	IterationCode *int64 `json:"IterationCode,omitempty"`
	// 标签 Id 列表
	LabelIds []int64 `json:"LabelIds,omitempty"`
	// 事项名称
	Name string `json:"Name"`
	// 父事项 Code  敏捷模式：Type 为 SUB_TASK 时必须指定  经典模式：Type 为 MISSION、REQUIREMENT 时可指定
	ParentCode *int64 `json:"ParentCode,omitempty"`
	// 紧急程度  \"0\" - 低  \"1\" - 中  \"2\" - 高  \"3\" - 紧急
	Priority string `json:"Priority"`
	// 项目模块 Id
	ProjectModuleId *int64 `json:"ProjectModuleId,omitempty"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
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
	// 排序目标位置的事项 code  可不填，排在底位
	TargetSortCode *int64 `json:"TargetSortCode,omitempty"`
	// 第三方链接列表
	ThirdLinks []IssueThirdLinkForm `json:"ThirdLinks,omitempty"`
	// 事项类型  DEFECT - 缺陷  REQUIREMENT - 需求  MISSION - 任务  EPIC - 史诗  SUB_TASK - 子任务
	Type string `json:"Type"`
	// 关注人 Id 列表
	WatcherIds []int64 `json:"WatcherIds,omitempty"`
	// 工时（小时数）
	WorkingHours *float32 `json:"WorkingHours,omitempty"`
}

type _CreateIssueRequest CreateIssueRequest

// NewCreateIssueRequest instantiates a new CreateIssueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIssueRequest(name string, priority string, projectName string, type_ string) *CreateIssueRequest {
	this := CreateIssueRequest{}
	this.Name = name
	this.Priority = priority
	this.ProjectName = projectName
	this.Type = type_
	return &this
}

// NewCreateIssueRequestWithDefaults instantiates a new CreateIssueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIssueRequestWithDefaults() *CreateIssueRequest {
	this := CreateIssueRequest{}
	return &this
}

// GetAssigneeId returns the AssigneeId field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetAssigneeId() int64 {
	if o == nil || utils.IsNil(o.AssigneeId) {
		var ret int64
		return ret
	}
	return *o.AssigneeId
}

// GetAssigneeIdOk returns a tuple with the AssigneeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetAssigneeIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.AssigneeId) {
		return nil, false
	}
	return o.AssigneeId, true
}

// HasAssigneeId returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasAssigneeId() bool {
	if o != nil && !utils.IsNil(o.AssigneeId) {
		return true
	}

	return false
}

// SetAssigneeId gets a reference to the given int64 and assigns it to the AssigneeId field.
func (o *CreateIssueRequest) SetAssigneeId(v int64) {
	o.AssigneeId = &v
}

// GetCustomFieldValues returns the CustomFieldValues field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetCustomFieldValues() []IssueCustomFieldForm {
	if o == nil || utils.IsNil(o.CustomFieldValues) {
		var ret []IssueCustomFieldForm
		return ret
	}
	return o.CustomFieldValues
}

// GetCustomFieldValuesOk returns a tuple with the CustomFieldValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetCustomFieldValuesOk() ([]IssueCustomFieldForm, bool) {
	if o == nil || utils.IsNil(o.CustomFieldValues) {
		return nil, false
	}
	return o.CustomFieldValues, true
}

// HasCustomFieldValues returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasCustomFieldValues() bool {
	if o != nil && !utils.IsNil(o.CustomFieldValues) {
		return true
	}

	return false
}

// SetCustomFieldValues gets a reference to the given []IssueCustomFieldForm and assigns it to the CustomFieldValues field.
func (o *CreateIssueRequest) SetCustomFieldValues(v []IssueCustomFieldForm) {
	o.CustomFieldValues = v
}

// GetDefectTypeId returns the DefectTypeId field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetDefectTypeId() int64 {
	if o == nil || utils.IsNil(o.DefectTypeId) {
		var ret int64
		return ret
	}
	return *o.DefectTypeId
}

// GetDefectTypeIdOk returns a tuple with the DefectTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetDefectTypeIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.DefectTypeId) {
		return nil, false
	}
	return o.DefectTypeId, true
}

// HasDefectTypeId returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasDefectTypeId() bool {
	if o != nil && !utils.IsNil(o.DefectTypeId) {
		return true
	}

	return false
}

// SetDefectTypeId gets a reference to the given int64 and assigns it to the DefectTypeId field.
func (o *CreateIssueRequest) SetDefectTypeId(v int64) {
	o.DefectTypeId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateIssueRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetDueDate() map[string]interface{} {
	if o == nil || utils.IsNil(o.DueDate) {
		var ret map[string]interface{}
		return ret
	}
	return o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetDueDateOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.DueDate) {
		return map[string]interface{}{}, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasDueDate() bool {
	if o != nil && !utils.IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given map[string]interface{} and assigns it to the DueDate field.
func (o *CreateIssueRequest) SetDueDate(v map[string]interface{}) {
	o.DueDate = v
}

// GetEpicCode returns the EpicCode field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetEpicCode() int64 {
	if o == nil || utils.IsNil(o.EpicCode) {
		var ret int64
		return ret
	}
	return *o.EpicCode
}

// GetEpicCodeOk returns a tuple with the EpicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetEpicCodeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.EpicCode) {
		return nil, false
	}
	return o.EpicCode, true
}

// HasEpicCode returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasEpicCode() bool {
	if o != nil && !utils.IsNil(o.EpicCode) {
		return true
	}

	return false
}

// SetEpicCode gets a reference to the given int64 and assigns it to the EpicCode field.
func (o *CreateIssueRequest) SetEpicCode(v int64) {
	o.EpicCode = &v
}

// GetFileIds returns the FileIds field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetFileIds() []int64 {
	if o == nil || utils.IsNil(o.FileIds) {
		var ret []int64
		return ret
	}
	return o.FileIds
}

// GetFileIdsOk returns a tuple with the FileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetFileIdsOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.FileIds) {
		return nil, false
	}
	return o.FileIds, true
}

// HasFileIds returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasFileIds() bool {
	if o != nil && !utils.IsNil(o.FileIds) {
		return true
	}

	return false
}

// SetFileIds gets a reference to the given []int64 and assigns it to the FileIds field.
func (o *CreateIssueRequest) SetFileIds(v []int64) {
	o.FileIds = v
}

// GetIssueTypeId returns the IssueTypeId field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetIssueTypeId() int64 {
	if o == nil || utils.IsNil(o.IssueTypeId) {
		var ret int64
		return ret
	}
	return *o.IssueTypeId
}

// GetIssueTypeIdOk returns a tuple with the IssueTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetIssueTypeIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.IssueTypeId) {
		return nil, false
	}
	return o.IssueTypeId, true
}

// HasIssueTypeId returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasIssueTypeId() bool {
	if o != nil && !utils.IsNil(o.IssueTypeId) {
		return true
	}

	return false
}

// SetIssueTypeId gets a reference to the given int64 and assigns it to the IssueTypeId field.
func (o *CreateIssueRequest) SetIssueTypeId(v int64) {
	o.IssueTypeId = &v
}

// GetIterationCode returns the IterationCode field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetIterationCode() int64 {
	if o == nil || utils.IsNil(o.IterationCode) {
		var ret int64
		return ret
	}
	return *o.IterationCode
}

// GetIterationCodeOk returns a tuple with the IterationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetIterationCodeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.IterationCode) {
		return nil, false
	}
	return o.IterationCode, true
}

// HasIterationCode returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasIterationCode() bool {
	if o != nil && !utils.IsNil(o.IterationCode) {
		return true
	}

	return false
}

// SetIterationCode gets a reference to the given int64 and assigns it to the IterationCode field.
func (o *CreateIssueRequest) SetIterationCode(v int64) {
	o.IterationCode = &v
}

// GetLabelIds returns the LabelIds field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetLabelIds() []int64 {
	if o == nil || utils.IsNil(o.LabelIds) {
		var ret []int64
		return ret
	}
	return o.LabelIds
}

// GetLabelIdsOk returns a tuple with the LabelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetLabelIdsOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.LabelIds) {
		return nil, false
	}
	return o.LabelIds, true
}

// HasLabelIds returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasLabelIds() bool {
	if o != nil && !utils.IsNil(o.LabelIds) {
		return true
	}

	return false
}

// SetLabelIds gets a reference to the given []int64 and assigns it to the LabelIds field.
func (o *CreateIssueRequest) SetLabelIds(v []int64) {
	o.LabelIds = v
}

// GetName returns the Name field value
func (o *CreateIssueRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateIssueRequest) SetName(v string) {
	o.Name = v
}

// GetParentCode returns the ParentCode field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetParentCode() int64 {
	if o == nil || utils.IsNil(o.ParentCode) {
		var ret int64
		return ret
	}
	return *o.ParentCode
}

// GetParentCodeOk returns a tuple with the ParentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetParentCodeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.ParentCode) {
		return nil, false
	}
	return o.ParentCode, true
}

// HasParentCode returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasParentCode() bool {
	if o != nil && !utils.IsNil(o.ParentCode) {
		return true
	}

	return false
}

// SetParentCode gets a reference to the given int64 and assigns it to the ParentCode field.
func (o *CreateIssueRequest) SetParentCode(v int64) {
	o.ParentCode = &v
}

// GetPriority returns the Priority field value
func (o *CreateIssueRequest) GetPriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *CreateIssueRequest) SetPriority(v string) {
	o.Priority = v
}

// GetProjectModuleId returns the ProjectModuleId field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetProjectModuleId() int64 {
	if o == nil || utils.IsNil(o.ProjectModuleId) {
		var ret int64
		return ret
	}
	return *o.ProjectModuleId
}

// GetProjectModuleIdOk returns a tuple with the ProjectModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetProjectModuleIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.ProjectModuleId) {
		return nil, false
	}
	return o.ProjectModuleId, true
}

// HasProjectModuleId returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasProjectModuleId() bool {
	if o != nil && !utils.IsNil(o.ProjectModuleId) {
		return true
	}

	return false
}

// SetProjectModuleId gets a reference to the given int64 and assigns it to the ProjectModuleId field.
func (o *CreateIssueRequest) SetProjectModuleId(v int64) {
	o.ProjectModuleId = &v
}

// GetProjectName returns the ProjectName field value
func (o *CreateIssueRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *CreateIssueRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetReleaseCodes returns the ReleaseCodes field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetReleaseCodes() []int64 {
	if o == nil || utils.IsNil(o.ReleaseCodes) {
		var ret []int64
		return ret
	}
	return o.ReleaseCodes
}

// GetReleaseCodesOk returns a tuple with the ReleaseCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetReleaseCodesOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.ReleaseCodes) {
		return nil, false
	}
	return o.ReleaseCodes, true
}

// HasReleaseCodes returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasReleaseCodes() bool {
	if o != nil && !utils.IsNil(o.ReleaseCodes) {
		return true
	}

	return false
}

// SetReleaseCodes gets a reference to the given []int64 and assigns it to the ReleaseCodes field.
func (o *CreateIssueRequest) SetReleaseCodes(v []int64) {
	o.ReleaseCodes = v
}

// GetRequirementTypeId returns the RequirementTypeId field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetRequirementTypeId() int64 {
	if o == nil || utils.IsNil(o.RequirementTypeId) {
		var ret int64
		return ret
	}
	return *o.RequirementTypeId
}

// GetRequirementTypeIdOk returns a tuple with the RequirementTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetRequirementTypeIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.RequirementTypeId) {
		return nil, false
	}
	return o.RequirementTypeId, true
}

// HasRequirementTypeId returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasRequirementTypeId() bool {
	if o != nil && !utils.IsNil(o.RequirementTypeId) {
		return true
	}

	return false
}

// SetRequirementTypeId gets a reference to the given int64 and assigns it to the RequirementTypeId field.
func (o *CreateIssueRequest) SetRequirementTypeId(v int64) {
	o.RequirementTypeId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetStartDate() map[string]interface{} {
	if o == nil || utils.IsNil(o.StartDate) {
		var ret map[string]interface{}
		return ret
	}
	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetStartDateOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.StartDate) {
		return map[string]interface{}{}, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasStartDate() bool {
	if o != nil && !utils.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given map[string]interface{} and assigns it to the StartDate field.
func (o *CreateIssueRequest) SetStartDate(v map[string]interface{}) {
	o.StartDate = v
}

// GetStatusId returns the StatusId field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetStatusId() int64 {
	if o == nil || utils.IsNil(o.StatusId) {
		var ret int64
		return ret
	}
	return *o.StatusId
}

// GetStatusIdOk returns a tuple with the StatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetStatusIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.StatusId) {
		return nil, false
	}
	return o.StatusId, true
}

// HasStatusId returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasStatusId() bool {
	if o != nil && !utils.IsNil(o.StatusId) {
		return true
	}

	return false
}

// SetStatusId gets a reference to the given int64 and assigns it to the StatusId field.
func (o *CreateIssueRequest) SetStatusId(v int64) {
	o.StatusId = &v
}

// GetStoryPoint returns the StoryPoint field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetStoryPoint() string {
	if o == nil || utils.IsNil(o.StoryPoint) {
		var ret string
		return ret
	}
	return *o.StoryPoint
}

// GetStoryPointOk returns a tuple with the StoryPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetStoryPointOk() (*string, bool) {
	if o == nil || utils.IsNil(o.StoryPoint) {
		return nil, false
	}
	return o.StoryPoint, true
}

// HasStoryPoint returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasStoryPoint() bool {
	if o != nil && !utils.IsNil(o.StoryPoint) {
		return true
	}

	return false
}

// SetStoryPoint gets a reference to the given string and assigns it to the StoryPoint field.
func (o *CreateIssueRequest) SetStoryPoint(v string) {
	o.StoryPoint = &v
}

// GetTargetSortCode returns the TargetSortCode field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetTargetSortCode() int64 {
	if o == nil || utils.IsNil(o.TargetSortCode) {
		var ret int64
		return ret
	}
	return *o.TargetSortCode
}

// GetTargetSortCodeOk returns a tuple with the TargetSortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetTargetSortCodeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.TargetSortCode) {
		return nil, false
	}
	return o.TargetSortCode, true
}

// HasTargetSortCode returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasTargetSortCode() bool {
	if o != nil && !utils.IsNil(o.TargetSortCode) {
		return true
	}

	return false
}

// SetTargetSortCode gets a reference to the given int64 and assigns it to the TargetSortCode field.
func (o *CreateIssueRequest) SetTargetSortCode(v int64) {
	o.TargetSortCode = &v
}

// GetThirdLinks returns the ThirdLinks field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetThirdLinks() []IssueThirdLinkForm {
	if o == nil || utils.IsNil(o.ThirdLinks) {
		var ret []IssueThirdLinkForm
		return ret
	}
	return o.ThirdLinks
}

// GetThirdLinksOk returns a tuple with the ThirdLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetThirdLinksOk() ([]IssueThirdLinkForm, bool) {
	if o == nil || utils.IsNil(o.ThirdLinks) {
		return nil, false
	}
	return o.ThirdLinks, true
}

// HasThirdLinks returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasThirdLinks() bool {
	if o != nil && !utils.IsNil(o.ThirdLinks) {
		return true
	}

	return false
}

// SetThirdLinks gets a reference to the given []IssueThirdLinkForm and assigns it to the ThirdLinks field.
func (o *CreateIssueRequest) SetThirdLinks(v []IssueThirdLinkForm) {
	o.ThirdLinks = v
}

// GetType returns the Type field value
func (o *CreateIssueRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateIssueRequest) SetType(v string) {
	o.Type = v
}

// GetWatcherIds returns the WatcherIds field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetWatcherIds() []int64 {
	if o == nil || utils.IsNil(o.WatcherIds) {
		var ret []int64
		return ret
	}
	return o.WatcherIds
}

// GetWatcherIdsOk returns a tuple with the WatcherIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetWatcherIdsOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.WatcherIds) {
		return nil, false
	}
	return o.WatcherIds, true
}

// HasWatcherIds returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasWatcherIds() bool {
	if o != nil && !utils.IsNil(o.WatcherIds) {
		return true
	}

	return false
}

// SetWatcherIds gets a reference to the given []int64 and assigns it to the WatcherIds field.
func (o *CreateIssueRequest) SetWatcherIds(v []int64) {
	o.WatcherIds = v
}

// GetWorkingHours returns the WorkingHours field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetWorkingHours() float32 {
	if o == nil || utils.IsNil(o.WorkingHours) {
		var ret float32
		return ret
	}
	return *o.WorkingHours
}

// GetWorkingHoursOk returns a tuple with the WorkingHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetWorkingHoursOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.WorkingHours) {
		return nil, false
	}
	return o.WorkingHours, true
}

// HasWorkingHours returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasWorkingHours() bool {
	if o != nil && !utils.IsNil(o.WorkingHours) {
		return true
	}

	return false
}

// SetWorkingHours gets a reference to the given float32 and assigns it to the WorkingHours field.
func (o *CreateIssueRequest) SetWorkingHours(v float32) {
	o.WorkingHours = &v
}

func (o CreateIssueRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIssueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AssigneeId) {
		toSerialize["AssigneeId"] = o.AssigneeId
	}
	if !utils.IsNil(o.CustomFieldValues) {
		toSerialize["CustomFieldValues"] = o.CustomFieldValues
	}
	if !utils.IsNil(o.DefectTypeId) {
		toSerialize["DefectTypeId"] = o.DefectTypeId
	}
	if !utils.IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !utils.IsNil(o.DueDate) {
		toSerialize["DueDate"] = o.DueDate
	}
	if !utils.IsNil(o.EpicCode) {
		toSerialize["EpicCode"] = o.EpicCode
	}
	if !utils.IsNil(o.FileIds) {
		toSerialize["FileIds"] = o.FileIds
	}
	if !utils.IsNil(o.IssueTypeId) {
		toSerialize["IssueTypeId"] = o.IssueTypeId
	}
	if !utils.IsNil(o.IterationCode) {
		toSerialize["IterationCode"] = o.IterationCode
	}
	if !utils.IsNil(o.LabelIds) {
		toSerialize["LabelIds"] = o.LabelIds
	}
	toSerialize["Name"] = o.Name
	if !utils.IsNil(o.ParentCode) {
		toSerialize["ParentCode"] = o.ParentCode
	}
	toSerialize["Priority"] = o.Priority
	if !utils.IsNil(o.ProjectModuleId) {
		toSerialize["ProjectModuleId"] = o.ProjectModuleId
	}
	toSerialize["ProjectName"] = o.ProjectName
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
	if !utils.IsNil(o.TargetSortCode) {
		toSerialize["TargetSortCode"] = o.TargetSortCode
	}
	if !utils.IsNil(o.ThirdLinks) {
		toSerialize["ThirdLinks"] = o.ThirdLinks
	}
	toSerialize["Type"] = o.Type
	if !utils.IsNil(o.WatcherIds) {
		toSerialize["WatcherIds"] = o.WatcherIds
	}
	if !utils.IsNil(o.WorkingHours) {
		toSerialize["WorkingHours"] = o.WorkingHours
	}
	return toSerialize, nil
}

func (o *CreateIssueRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Name",
		"Priority",
		"ProjectName",
		"Type",
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

	varCreateIssueRequest := _CreateIssueRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateIssueRequest)

	if err != nil {
		return err
	}

	*o = CreateIssueRequest(varCreateIssueRequest)

	return err
}

type NullableCreateIssueRequest struct {
	value *CreateIssueRequest
	isSet bool
}

func (v NullableCreateIssueRequest) Get() *CreateIssueRequest {
	return v.value
}

func (v *NullableCreateIssueRequest) Set(val *CreateIssueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIssueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIssueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIssueRequest(val *CreateIssueRequest) *NullableCreateIssueRequest {
	return &NullableCreateIssueRequest{value: val, isSet: true}
}

func (v NullableCreateIssueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIssueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


