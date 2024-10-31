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

// checks if the CodingCIJob type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CodingCIJob{}

// CodingCIJob 构建计划信息
type CodingCIJob struct {
	// 不管构建成功还是失败总是通知的用户
	AlwaysUserIdList []int32 `json:"AlwaysUserIdList"`
	// 自动取消相同 MR
	AutoCancelSameMergeRequest bool `json:"AutoCancelSameMergeRequest"`
	// 自动取消相同版本
	AutoCancelSameRevision bool `json:"AutoCancelSameRevision"`
	// 分支匹配正则
	BranchRegex string `json:"BranchRegex"`
	// 触发构建的分支
	BranchSelector string `json:"BranchSelector"`
	// 仅构建失败时要通知的用户
	BuildFailUserIdList []int32 `json:"BuildFailUserIdList"`
	// 任务缓存目录配置
	CachePathList []CIJobCachePath `json:"CachePathList"`
	// 构建缓存大小
	CacheSize *int32 `json:"CacheSize,omitempty"`
	// 创建时间
	CreatedAt int32 `json:"CreatedAt"`
	// 创建者
	CreatorId int32 `json:"CreatorId"`
	// 仓库库的 Https 地址
	DepotHttpsUrl *string `json:"DepotHttpsUrl,omitempty"`
	// 仓库ID
	DepotId int32 `json:"DepotId"`
	// 仓库名称
	DepotName *string `json:"DepotName,omitempty"`
	// 仓库库的 SSH 地址
	DepotSshUrl *string `json:"DepotSshUrl,omitempty"`
	// 仓库类型
	DepotType string `json:"DepotType"`
	// 仓库库的 Web 页面
	DepotWebUrl *string `json:"DepotWebUrl,omitempty"`
	// 环境变量配置
	EnvList []CIJobEnv `json:"EnvList"`
	// 执行方式
	ExecuteIn string `json:"ExecuteIn"`
	// 自定义构建节点池 ID
	ExecutedAgentPoolId *int32 `json:"ExecutedAgentPoolId,omitempty"`
	// 代码更新触发匹配规则
	HookType string `json:"HookType"`
	// 构建计划ID
	Id *int32 `json:"Id,omitempty"`
	// Jenkinsfile 来源
	JenkinsFileFromType string `json:"JenkinsFileFromType"`
	// Jenkinsfile 在仓库中的文件路径
	JenkinsFilePath string `json:"JenkinsFilePath"`
	// Jenkinsfile 文件内容
	JenkinsFileStaticContent string `json:"JenkinsFileStaticContent"`
	// 构建计划创建来源
	JobFromType string `json:"JobFromType"`
	// 构建计划名称
	Name string `json:"Name"`
	// 项目ID
	ProjectId int32 `json:"ProjectId"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty"`
	// 针对 CRON triggerMethod 的 schedule 规则配置
	ScheduleList []CIJobSchedule `json:"ScheduleList"`
	// 构建计划触发方式
	TriggerMethodList []string `json:"TriggerMethodList"`
	// 构建结果通知触发者机制
	TriggerRemind string `json:"TriggerRemind"`
	// 最后更新时间
	UpdatedAt int32 `json:"UpdatedAt"`
}

type _CodingCIJob CodingCIJob

// NewCodingCIJob instantiates a new CodingCIJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodingCIJob(alwaysUserIdList []int32, autoCancelSameMergeRequest bool, autoCancelSameRevision bool, branchRegex string, branchSelector string, buildFailUserIdList []int32, cachePathList []CIJobCachePath, createdAt int32, creatorId int32, depotId int32, depotType string, envList []CIJobEnv, executeIn string, hookType string, jenkinsFileFromType string, jenkinsFilePath string, jenkinsFileStaticContent string, jobFromType string, name string, projectId int32, scheduleList []CIJobSchedule, triggerMethodList []string, triggerRemind string, updatedAt int32) *CodingCIJob {
	this := CodingCIJob{}
	this.AlwaysUserIdList = alwaysUserIdList
	this.AutoCancelSameMergeRequest = autoCancelSameMergeRequest
	this.AutoCancelSameRevision = autoCancelSameRevision
	this.BranchRegex = branchRegex
	this.BranchSelector = branchSelector
	this.BuildFailUserIdList = buildFailUserIdList
	this.CachePathList = cachePathList
	this.CreatedAt = createdAt
	this.CreatorId = creatorId
	var depotHttpsUrl string = ""
	this.DepotHttpsUrl = &depotHttpsUrl
	this.DepotId = depotId
	var depotName string = ""
	this.DepotName = &depotName
	var depotSshUrl string = ""
	this.DepotSshUrl = &depotSshUrl
	this.DepotType = depotType
	var depotWebUrl string = ""
	this.DepotWebUrl = &depotWebUrl
	this.EnvList = envList
	this.ExecuteIn = executeIn
	this.HookType = hookType
	this.JenkinsFileFromType = jenkinsFileFromType
	this.JenkinsFilePath = jenkinsFilePath
	this.JenkinsFileStaticContent = jenkinsFileStaticContent
	this.JobFromType = jobFromType
	this.Name = name
	this.ProjectId = projectId
	var projectName string = ""
	this.ProjectName = &projectName
	this.ScheduleList = scheduleList
	this.TriggerMethodList = triggerMethodList
	this.TriggerRemind = triggerRemind
	this.UpdatedAt = updatedAt
	return &this
}

// NewCodingCIJobWithDefaults instantiates a new CodingCIJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodingCIJobWithDefaults() *CodingCIJob {
	this := CodingCIJob{}
	var autoCancelSameMergeRequest bool = false
	this.AutoCancelSameMergeRequest = autoCancelSameMergeRequest
	var autoCancelSameRevision bool = false
	this.AutoCancelSameRevision = autoCancelSameRevision
	var branchRegex string = ""
	this.BranchRegex = branchRegex
	var branchSelector string = ""
	this.BranchSelector = branchSelector
	var depotHttpsUrl string = ""
	this.DepotHttpsUrl = &depotHttpsUrl
	var depotName string = ""
	this.DepotName = &depotName
	var depotSshUrl string = ""
	this.DepotSshUrl = &depotSshUrl
	var depotType string = ""
	this.DepotType = depotType
	var depotWebUrl string = ""
	this.DepotWebUrl = &depotWebUrl
	var executeIn string = ""
	this.ExecuteIn = executeIn
	var hookType string = ""
	this.HookType = hookType
	var jenkinsFileFromType string = ""
	this.JenkinsFileFromType = jenkinsFileFromType
	var jenkinsFilePath string = ""
	this.JenkinsFilePath = jenkinsFilePath
	var jenkinsFileStaticContent string = ""
	this.JenkinsFileStaticContent = jenkinsFileStaticContent
	var jobFromType string = ""
	this.JobFromType = jobFromType
	var name string = ""
	this.Name = name
	var projectName string = ""
	this.ProjectName = &projectName
	var triggerRemind string = ""
	this.TriggerRemind = triggerRemind
	return &this
}

// GetAlwaysUserIdList returns the AlwaysUserIdList field value
func (o *CodingCIJob) GetAlwaysUserIdList() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.AlwaysUserIdList
}

// GetAlwaysUserIdListOk returns a tuple with the AlwaysUserIdList field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetAlwaysUserIdListOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlwaysUserIdList, true
}

// SetAlwaysUserIdList sets field value
func (o *CodingCIJob) SetAlwaysUserIdList(v []int32) {
	o.AlwaysUserIdList = v
}

// GetAutoCancelSameMergeRequest returns the AutoCancelSameMergeRequest field value
func (o *CodingCIJob) GetAutoCancelSameMergeRequest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoCancelSameMergeRequest
}

// GetAutoCancelSameMergeRequestOk returns a tuple with the AutoCancelSameMergeRequest field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetAutoCancelSameMergeRequestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoCancelSameMergeRequest, true
}

// SetAutoCancelSameMergeRequest sets field value
func (o *CodingCIJob) SetAutoCancelSameMergeRequest(v bool) {
	o.AutoCancelSameMergeRequest = v
}

// GetAutoCancelSameRevision returns the AutoCancelSameRevision field value
func (o *CodingCIJob) GetAutoCancelSameRevision() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoCancelSameRevision
}

// GetAutoCancelSameRevisionOk returns a tuple with the AutoCancelSameRevision field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetAutoCancelSameRevisionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoCancelSameRevision, true
}

// SetAutoCancelSameRevision sets field value
func (o *CodingCIJob) SetAutoCancelSameRevision(v bool) {
	o.AutoCancelSameRevision = v
}

// GetBranchRegex returns the BranchRegex field value
func (o *CodingCIJob) GetBranchRegex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BranchRegex
}

// GetBranchRegexOk returns a tuple with the BranchRegex field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetBranchRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchRegex, true
}

// SetBranchRegex sets field value
func (o *CodingCIJob) SetBranchRegex(v string) {
	o.BranchRegex = v
}

// GetBranchSelector returns the BranchSelector field value
func (o *CodingCIJob) GetBranchSelector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BranchSelector
}

// GetBranchSelectorOk returns a tuple with the BranchSelector field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetBranchSelectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchSelector, true
}

// SetBranchSelector sets field value
func (o *CodingCIJob) SetBranchSelector(v string) {
	o.BranchSelector = v
}

// GetBuildFailUserIdList returns the BuildFailUserIdList field value
func (o *CodingCIJob) GetBuildFailUserIdList() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.BuildFailUserIdList
}

// GetBuildFailUserIdListOk returns a tuple with the BuildFailUserIdList field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetBuildFailUserIdListOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildFailUserIdList, true
}

// SetBuildFailUserIdList sets field value
func (o *CodingCIJob) SetBuildFailUserIdList(v []int32) {
	o.BuildFailUserIdList = v
}

// GetCachePathList returns the CachePathList field value
func (o *CodingCIJob) GetCachePathList() []CIJobCachePath {
	if o == nil {
		var ret []CIJobCachePath
		return ret
	}

	return o.CachePathList
}

// GetCachePathListOk returns a tuple with the CachePathList field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetCachePathListOk() ([]CIJobCachePath, bool) {
	if o == nil {
		return nil, false
	}
	return o.CachePathList, true
}

// SetCachePathList sets field value
func (o *CodingCIJob) SetCachePathList(v []CIJobCachePath) {
	o.CachePathList = v
}

// GetCacheSize returns the CacheSize field value if set, zero value otherwise.
func (o *CodingCIJob) GetCacheSize() int32 {
	if o == nil || utils.IsNil(o.CacheSize) {
		var ret int32
		return ret
	}
	return *o.CacheSize
}

// GetCacheSizeOk returns a tuple with the CacheSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetCacheSizeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.CacheSize) {
		return nil, false
	}
	return o.CacheSize, true
}

// HasCacheSize returns a boolean if a field has been set.
func (o *CodingCIJob) HasCacheSize() bool {
	if o != nil && !utils.IsNil(o.CacheSize) {
		return true
	}

	return false
}

// SetCacheSize gets a reference to the given int32 and assigns it to the CacheSize field.
func (o *CodingCIJob) SetCacheSize(v int32) {
	o.CacheSize = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CodingCIJob) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CodingCIJob) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetCreatorId returns the CreatorId field value
func (o *CodingCIJob) GetCreatorId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetCreatorIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *CodingCIJob) SetCreatorId(v int32) {
	o.CreatorId = v
}

// GetDepotHttpsUrl returns the DepotHttpsUrl field value if set, zero value otherwise.
func (o *CodingCIJob) GetDepotHttpsUrl() string {
	if o == nil || utils.IsNil(o.DepotHttpsUrl) {
		var ret string
		return ret
	}
	return *o.DepotHttpsUrl
}

// GetDepotHttpsUrlOk returns a tuple with the DepotHttpsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetDepotHttpsUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotHttpsUrl) {
		return nil, false
	}
	return o.DepotHttpsUrl, true
}

// HasDepotHttpsUrl returns a boolean if a field has been set.
func (o *CodingCIJob) HasDepotHttpsUrl() bool {
	if o != nil && !utils.IsNil(o.DepotHttpsUrl) {
		return true
	}

	return false
}

// SetDepotHttpsUrl gets a reference to the given string and assigns it to the DepotHttpsUrl field.
func (o *CodingCIJob) SetDepotHttpsUrl(v string) {
	o.DepotHttpsUrl = &v
}

// GetDepotId returns the DepotId field value
func (o *CodingCIJob) GetDepotId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetDepotIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *CodingCIJob) SetDepotId(v int32) {
	o.DepotId = v
}

// GetDepotName returns the DepotName field value if set, zero value otherwise.
func (o *CodingCIJob) GetDepotName() string {
	if o == nil || utils.IsNil(o.DepotName) {
		var ret string
		return ret
	}
	return *o.DepotName
}

// GetDepotNameOk returns a tuple with the DepotName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetDepotNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotName) {
		return nil, false
	}
	return o.DepotName, true
}

// HasDepotName returns a boolean if a field has been set.
func (o *CodingCIJob) HasDepotName() bool {
	if o != nil && !utils.IsNil(o.DepotName) {
		return true
	}

	return false
}

// SetDepotName gets a reference to the given string and assigns it to the DepotName field.
func (o *CodingCIJob) SetDepotName(v string) {
	o.DepotName = &v
}

// GetDepotSshUrl returns the DepotSshUrl field value if set, zero value otherwise.
func (o *CodingCIJob) GetDepotSshUrl() string {
	if o == nil || utils.IsNil(o.DepotSshUrl) {
		var ret string
		return ret
	}
	return *o.DepotSshUrl
}

// GetDepotSshUrlOk returns a tuple with the DepotSshUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetDepotSshUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotSshUrl) {
		return nil, false
	}
	return o.DepotSshUrl, true
}

// HasDepotSshUrl returns a boolean if a field has been set.
func (o *CodingCIJob) HasDepotSshUrl() bool {
	if o != nil && !utils.IsNil(o.DepotSshUrl) {
		return true
	}

	return false
}

// SetDepotSshUrl gets a reference to the given string and assigns it to the DepotSshUrl field.
func (o *CodingCIJob) SetDepotSshUrl(v string) {
	o.DepotSshUrl = &v
}

// GetDepotType returns the DepotType field value
func (o *CodingCIJob) GetDepotType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepotType
}

// GetDepotTypeOk returns a tuple with the DepotType field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetDepotTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotType, true
}

// SetDepotType sets field value
func (o *CodingCIJob) SetDepotType(v string) {
	o.DepotType = v
}

// GetDepotWebUrl returns the DepotWebUrl field value if set, zero value otherwise.
func (o *CodingCIJob) GetDepotWebUrl() string {
	if o == nil || utils.IsNil(o.DepotWebUrl) {
		var ret string
		return ret
	}
	return *o.DepotWebUrl
}

// GetDepotWebUrlOk returns a tuple with the DepotWebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetDepotWebUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotWebUrl) {
		return nil, false
	}
	return o.DepotWebUrl, true
}

// HasDepotWebUrl returns a boolean if a field has been set.
func (o *CodingCIJob) HasDepotWebUrl() bool {
	if o != nil && !utils.IsNil(o.DepotWebUrl) {
		return true
	}

	return false
}

// SetDepotWebUrl gets a reference to the given string and assigns it to the DepotWebUrl field.
func (o *CodingCIJob) SetDepotWebUrl(v string) {
	o.DepotWebUrl = &v
}

// GetEnvList returns the EnvList field value
func (o *CodingCIJob) GetEnvList() []CIJobEnv {
	if o == nil {
		var ret []CIJobEnv
		return ret
	}

	return o.EnvList
}

// GetEnvListOk returns a tuple with the EnvList field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetEnvListOk() ([]CIJobEnv, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvList, true
}

// SetEnvList sets field value
func (o *CodingCIJob) SetEnvList(v []CIJobEnv) {
	o.EnvList = v
}

// GetExecuteIn returns the ExecuteIn field value
func (o *CodingCIJob) GetExecuteIn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExecuteIn
}

// GetExecuteInOk returns a tuple with the ExecuteIn field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetExecuteInOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecuteIn, true
}

// SetExecuteIn sets field value
func (o *CodingCIJob) SetExecuteIn(v string) {
	o.ExecuteIn = v
}

// GetExecutedAgentPoolId returns the ExecutedAgentPoolId field value if set, zero value otherwise.
func (o *CodingCIJob) GetExecutedAgentPoolId() int32 {
	if o == nil || utils.IsNil(o.ExecutedAgentPoolId) {
		var ret int32
		return ret
	}
	return *o.ExecutedAgentPoolId
}

// GetExecutedAgentPoolIdOk returns a tuple with the ExecutedAgentPoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetExecutedAgentPoolIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ExecutedAgentPoolId) {
		return nil, false
	}
	return o.ExecutedAgentPoolId, true
}

// HasExecutedAgentPoolId returns a boolean if a field has been set.
func (o *CodingCIJob) HasExecutedAgentPoolId() bool {
	if o != nil && !utils.IsNil(o.ExecutedAgentPoolId) {
		return true
	}

	return false
}

// SetExecutedAgentPoolId gets a reference to the given int32 and assigns it to the ExecutedAgentPoolId field.
func (o *CodingCIJob) SetExecutedAgentPoolId(v int32) {
	o.ExecutedAgentPoolId = &v
}

// GetHookType returns the HookType field value
func (o *CodingCIJob) GetHookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HookType
}

// GetHookTypeOk returns a tuple with the HookType field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetHookTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HookType, true
}

// SetHookType sets field value
func (o *CodingCIJob) SetHookType(v string) {
	o.HookType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CodingCIJob) GetId() int32 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CodingCIJob) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CodingCIJob) SetId(v int32) {
	o.Id = &v
}

// GetJenkinsFileFromType returns the JenkinsFileFromType field value
func (o *CodingCIJob) GetJenkinsFileFromType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JenkinsFileFromType
}

// GetJenkinsFileFromTypeOk returns a tuple with the JenkinsFileFromType field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetJenkinsFileFromTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JenkinsFileFromType, true
}

// SetJenkinsFileFromType sets field value
func (o *CodingCIJob) SetJenkinsFileFromType(v string) {
	o.JenkinsFileFromType = v
}

// GetJenkinsFilePath returns the JenkinsFilePath field value
func (o *CodingCIJob) GetJenkinsFilePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JenkinsFilePath
}

// GetJenkinsFilePathOk returns a tuple with the JenkinsFilePath field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetJenkinsFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JenkinsFilePath, true
}

// SetJenkinsFilePath sets field value
func (o *CodingCIJob) SetJenkinsFilePath(v string) {
	o.JenkinsFilePath = v
}

// GetJenkinsFileStaticContent returns the JenkinsFileStaticContent field value
func (o *CodingCIJob) GetJenkinsFileStaticContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JenkinsFileStaticContent
}

// GetJenkinsFileStaticContentOk returns a tuple with the JenkinsFileStaticContent field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetJenkinsFileStaticContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JenkinsFileStaticContent, true
}

// SetJenkinsFileStaticContent sets field value
func (o *CodingCIJob) SetJenkinsFileStaticContent(v string) {
	o.JenkinsFileStaticContent = v
}

// GetJobFromType returns the JobFromType field value
func (o *CodingCIJob) GetJobFromType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobFromType
}

// GetJobFromTypeOk returns a tuple with the JobFromType field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetJobFromTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobFromType, true
}

// SetJobFromType sets field value
func (o *CodingCIJob) SetJobFromType(v string) {
	o.JobFromType = v
}

// GetName returns the Name field value
func (o *CodingCIJob) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CodingCIJob) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value
func (o *CodingCIJob) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *CodingCIJob) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *CodingCIJob) GetProjectName() string {
	if o == nil || utils.IsNil(o.ProjectName) {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetProjectNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ProjectName) {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *CodingCIJob) HasProjectName() bool {
	if o != nil && !utils.IsNil(o.ProjectName) {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *CodingCIJob) SetProjectName(v string) {
	o.ProjectName = &v
}

// GetScheduleList returns the ScheduleList field value
func (o *CodingCIJob) GetScheduleList() []CIJobSchedule {
	if o == nil {
		var ret []CIJobSchedule
		return ret
	}

	return o.ScheduleList
}

// GetScheduleListOk returns a tuple with the ScheduleList field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetScheduleListOk() ([]CIJobSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduleList, true
}

// SetScheduleList sets field value
func (o *CodingCIJob) SetScheduleList(v []CIJobSchedule) {
	o.ScheduleList = v
}

// GetTriggerMethodList returns the TriggerMethodList field value
func (o *CodingCIJob) GetTriggerMethodList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TriggerMethodList
}

// GetTriggerMethodListOk returns a tuple with the TriggerMethodList field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetTriggerMethodListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TriggerMethodList, true
}

// SetTriggerMethodList sets field value
func (o *CodingCIJob) SetTriggerMethodList(v []string) {
	o.TriggerMethodList = v
}

// GetTriggerRemind returns the TriggerRemind field value
func (o *CodingCIJob) GetTriggerRemind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggerRemind
}

// GetTriggerRemindOk returns a tuple with the TriggerRemind field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetTriggerRemindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerRemind, true
}

// SetTriggerRemind sets field value
func (o *CodingCIJob) SetTriggerRemind(v string) {
	o.TriggerRemind = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CodingCIJob) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CodingCIJob) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

func (o CodingCIJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CodingCIJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AlwaysUserIdList"] = o.AlwaysUserIdList
	toSerialize["AutoCancelSameMergeRequest"] = o.AutoCancelSameMergeRequest
	toSerialize["AutoCancelSameRevision"] = o.AutoCancelSameRevision
	toSerialize["BranchRegex"] = o.BranchRegex
	toSerialize["BranchSelector"] = o.BranchSelector
	toSerialize["BuildFailUserIdList"] = o.BuildFailUserIdList
	toSerialize["CachePathList"] = o.CachePathList
	if !utils.IsNil(o.CacheSize) {
		toSerialize["CacheSize"] = o.CacheSize
	}
	toSerialize["CreatedAt"] = o.CreatedAt
	toSerialize["CreatorId"] = o.CreatorId
	if !utils.IsNil(o.DepotHttpsUrl) {
		toSerialize["DepotHttpsUrl"] = o.DepotHttpsUrl
	}
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotName) {
		toSerialize["DepotName"] = o.DepotName
	}
	if !utils.IsNil(o.DepotSshUrl) {
		toSerialize["DepotSshUrl"] = o.DepotSshUrl
	}
	toSerialize["DepotType"] = o.DepotType
	if !utils.IsNil(o.DepotWebUrl) {
		toSerialize["DepotWebUrl"] = o.DepotWebUrl
	}
	toSerialize["EnvList"] = o.EnvList
	toSerialize["ExecuteIn"] = o.ExecuteIn
	if !utils.IsNil(o.ExecutedAgentPoolId) {
		toSerialize["ExecutedAgentPoolId"] = o.ExecutedAgentPoolId
	}
	toSerialize["HookType"] = o.HookType
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	toSerialize["JenkinsFileFromType"] = o.JenkinsFileFromType
	toSerialize["JenkinsFilePath"] = o.JenkinsFilePath
	toSerialize["JenkinsFileStaticContent"] = o.JenkinsFileStaticContent
	toSerialize["JobFromType"] = o.JobFromType
	toSerialize["Name"] = o.Name
	toSerialize["ProjectId"] = o.ProjectId
	if !utils.IsNil(o.ProjectName) {
		toSerialize["ProjectName"] = o.ProjectName
	}
	toSerialize["ScheduleList"] = o.ScheduleList
	toSerialize["TriggerMethodList"] = o.TriggerMethodList
	toSerialize["TriggerRemind"] = o.TriggerRemind
	toSerialize["UpdatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *CodingCIJob) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"AlwaysUserIdList",
		"AutoCancelSameMergeRequest",
		"AutoCancelSameRevision",
		"BranchRegex",
		"BranchSelector",
		"BuildFailUserIdList",
		"CachePathList",
		"CreatedAt",
		"CreatorId",
		"DepotId",
		"DepotType",
		"EnvList",
		"ExecuteIn",
		"HookType",
		"JenkinsFileFromType",
		"JenkinsFilePath",
		"JenkinsFileStaticContent",
		"JobFromType",
		"Name",
		"ProjectId",
		"ScheduleList",
		"TriggerMethodList",
		"TriggerRemind",
		"UpdatedAt",
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

	varCodingCIJob := _CodingCIJob{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCodingCIJob)

	if err != nil {
		return err
	}

	*o = CodingCIJob(varCodingCIJob)

	return err
}

type NullableCodingCIJob struct {
	value *CodingCIJob
	isSet bool
}

func (v NullableCodingCIJob) Get() *CodingCIJob {
	return v.value
}

func (v *NullableCodingCIJob) Set(val *CodingCIJob) {
	v.value = val
	v.isSet = true
}

func (v NullableCodingCIJob) IsSet() bool {
	return v.isSet
}

func (v *NullableCodingCIJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodingCIJob(val *CodingCIJob) *NullableCodingCIJob {
	return &NullableCodingCIJob{value: val, isSet: true}
}

func (v NullableCodingCIJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodingCIJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

