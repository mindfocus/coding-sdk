/*
CODING OPEN API

  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the MergeRequestData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MergeRequestData{}

// MergeRequestData 合并请求数据
type MergeRequestData struct {
	// 合并时间
	ActionAt utils.NullableInt64 `json:"ActionAt,omitempty"`
	ActionAuthor *CodingUser `json:"ActionAuthor,omitempty"`
	Author *CodingUser `json:"Author,omitempty"`
	// 基础sha
	BaseSha utils.NullableString `json:"BaseSha,omitempty"`
	// 评论数
	CommentCount utils.NullableInt64 `json:"CommentCount,omitempty"`
	// 创建时间
	CreatedAt utils.NullableInt64 `json:"CreatedAt,omitempty"`
	// 仓库 Id
	DepotId *int64 `json:"DepotId,omitempty"`
	// MR描述
	Describe utils.NullableString `json:"Describe,omitempty"`
	// 授权数
	Granted utils.NullableInt64 `json:"Granted,omitempty"`
	// 合并请求 Id
	Id *int64 `json:"Id,omitempty"`
	// 标签列表
	Labels []string `json:"Labels,omitempty"`
	// 合并Commit Sha
	MergeCommitSha utils.NullableString `json:"MergeCommitSha,omitempty"`
	// 合并请求 IId
	MergeId *int64 `json:"MergeId,omitempty"`
	// 合并请求 web 页面路径
	Path utils.NullableString `json:"Path,omitempty"`
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty"`
	// 是否提醒
	Reminded utils.NullableBool `json:"Reminded,omitempty"`
	// 评审者列表
	Reviewers []CodingUser `json:"Reviewers,omitempty"`
	// 源分支
	SourceBranch utils.NullableString `json:"SourceBranch,omitempty"`
	// 源分支Commit Sha
	SourceBranchSha utils.NullableString `json:"SourceBranchSha,omitempty"`
	// 合并请求状态
	Status utils.NullableString `json:"Status,omitempty"`
	// MR阻塞点
	StickingPoint utils.NullableString `json:"StickingPoint,omitempty"`
	// 目标分支
	TargetBranch utils.NullableString `json:"TargetBranch,omitempty"`
	// 目标分支是否保护分支
	TargetBranchProtected utils.NullableBool `json:"TargetBranchProtected,omitempty"`
	// 目标分支Commit Sha
	TargetBranchSha utils.NullableString `json:"TargetBranchSha,omitempty"`
	// 合并请求标题
	Title *string `json:"Title,omitempty"`
	// 更新时间
	UpdateAt utils.NullableInt64 `json:"UpdateAt,omitempty"`
}

// NewMergeRequestData instantiates a new MergeRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeRequestData() *MergeRequestData {
	this := MergeRequestData{}
	var baseSha string = ""
	this.BaseSha = *utils.NewNullableString(&baseSha)
	var describe string = ""
	this.Describe = *utils.NewNullableString(&describe)
	var mergeCommitSha string = ""
	this.MergeCommitSha = *utils.NewNullableString(&mergeCommitSha)
	var path string = ""
	this.Path = *utils.NewNullableString(&path)
	var reminded bool = false
	this.Reminded = *utils.NewNullableBool(&reminded)
	var sourceBranch string = ""
	this.SourceBranch = *utils.NewNullableString(&sourceBranch)
	var sourceBranchSha string = ""
	this.SourceBranchSha = *utils.NewNullableString(&sourceBranchSha)
	var status string = ""
	this.Status = *utils.NewNullableString(&status)
	var stickingPoint string = ""
	this.StickingPoint = *utils.NewNullableString(&stickingPoint)
	var targetBranch string = ""
	this.TargetBranch = *utils.NewNullableString(&targetBranch)
	var targetBranchProtected bool = false
	this.TargetBranchProtected = *utils.NewNullableBool(&targetBranchProtected)
	var targetBranchSha string = ""
	this.TargetBranchSha = *utils.NewNullableString(&targetBranchSha)
	var title string = ""
	this.Title = &title
	return &this
}

// NewMergeRequestDataWithDefaults instantiates a new MergeRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeRequestDataWithDefaults() *MergeRequestData {
	this := MergeRequestData{}
	var baseSha string = ""
	this.BaseSha = *utils.NewNullableString(&baseSha)
	var describe string = ""
	this.Describe = *utils.NewNullableString(&describe)
	var mergeCommitSha string = ""
	this.MergeCommitSha = *utils.NewNullableString(&mergeCommitSha)
	var path string = ""
	this.Path = *utils.NewNullableString(&path)
	var reminded bool = false
	this.Reminded = *utils.NewNullableBool(&reminded)
	var sourceBranch string = ""
	this.SourceBranch = *utils.NewNullableString(&sourceBranch)
	var sourceBranchSha string = ""
	this.SourceBranchSha = *utils.NewNullableString(&sourceBranchSha)
	var status string = ""
	this.Status = *utils.NewNullableString(&status)
	var stickingPoint string = ""
	this.StickingPoint = *utils.NewNullableString(&stickingPoint)
	var targetBranch string = ""
	this.TargetBranch = *utils.NewNullableString(&targetBranch)
	var targetBranchProtected bool = false
	this.TargetBranchProtected = *utils.NewNullableBool(&targetBranchProtected)
	var targetBranchSha string = ""
	this.TargetBranchSha = *utils.NewNullableString(&targetBranchSha)
	var title string = ""
	this.Title = &title
	return &this
}

// GetActionAt returns the ActionAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetActionAt() int64 {
	if o == nil || utils.IsNil(o.ActionAt.Get()) {
		var ret int64
		return ret
	}
	return *o.ActionAt.Get()
}

// GetActionAtOk returns a tuple with the ActionAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetActionAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionAt.Get(), o.ActionAt.IsSet()
}

// HasActionAt returns a boolean if a field has been set.
func (o *MergeRequestData) HasActionAt() bool {
	if o != nil && o.ActionAt.IsSet() {
		return true
	}

	return false
}

// SetActionAt gets a reference to the given utils.NullableInt64 and assigns it to the ActionAt field.
func (o *MergeRequestData) SetActionAt(v int64) {
	o.ActionAt.Set(&v)
}
// SetActionAtNil sets the value for ActionAt to be an explicit nil
func (o *MergeRequestData) SetActionAtNil() {
	o.ActionAt.Set(nil)
}

// UnsetActionAt ensures that no value is present for ActionAt, not even an explicit nil
func (o *MergeRequestData) UnsetActionAt() {
	o.ActionAt.Unset()
}

// GetActionAuthor returns the ActionAuthor field value if set, zero value otherwise.
func (o *MergeRequestData) GetActionAuthor() CodingUser {
	if o == nil || utils.IsNil(o.ActionAuthor) {
		var ret CodingUser
		return ret
	}
	return *o.ActionAuthor
}

// GetActionAuthorOk returns a tuple with the ActionAuthor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestData) GetActionAuthorOk() (*CodingUser, bool) {
	if o == nil || utils.IsNil(o.ActionAuthor) {
		return nil, false
	}
	return o.ActionAuthor, true
}

// HasActionAuthor returns a boolean if a field has been set.
func (o *MergeRequestData) HasActionAuthor() bool {
	if o != nil && !utils.IsNil(o.ActionAuthor) {
		return true
	}

	return false
}

// SetActionAuthor gets a reference to the given CodingUser and assigns it to the ActionAuthor field.
func (o *MergeRequestData) SetActionAuthor(v CodingUser) {
	o.ActionAuthor = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *MergeRequestData) GetAuthor() CodingUser {
	if o == nil || utils.IsNil(o.Author) {
		var ret CodingUser
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestData) GetAuthorOk() (*CodingUser, bool) {
	if o == nil || utils.IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *MergeRequestData) HasAuthor() bool {
	if o != nil && !utils.IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given CodingUser and assigns it to the Author field.
func (o *MergeRequestData) SetAuthor(v CodingUser) {
	o.Author = &v
}

// GetBaseSha returns the BaseSha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetBaseSha() string {
	if o == nil || utils.IsNil(o.BaseSha.Get()) {
		var ret string
		return ret
	}
	return *o.BaseSha.Get()
}

// GetBaseShaOk returns a tuple with the BaseSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetBaseShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseSha.Get(), o.BaseSha.IsSet()
}

// HasBaseSha returns a boolean if a field has been set.
func (o *MergeRequestData) HasBaseSha() bool {
	if o != nil && o.BaseSha.IsSet() {
		return true
	}

	return false
}

// SetBaseSha gets a reference to the given utils.NullableString and assigns it to the BaseSha field.
func (o *MergeRequestData) SetBaseSha(v string) {
	o.BaseSha.Set(&v)
}
// SetBaseShaNil sets the value for BaseSha to be an explicit nil
func (o *MergeRequestData) SetBaseShaNil() {
	o.BaseSha.Set(nil)
}

// UnsetBaseSha ensures that no value is present for BaseSha, not even an explicit nil
func (o *MergeRequestData) UnsetBaseSha() {
	o.BaseSha.Unset()
}

// GetCommentCount returns the CommentCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetCommentCount() int64 {
	if o == nil || utils.IsNil(o.CommentCount.Get()) {
		var ret int64
		return ret
	}
	return *o.CommentCount.Get()
}

// GetCommentCountOk returns a tuple with the CommentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetCommentCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommentCount.Get(), o.CommentCount.IsSet()
}

// HasCommentCount returns a boolean if a field has been set.
func (o *MergeRequestData) HasCommentCount() bool {
	if o != nil && o.CommentCount.IsSet() {
		return true
	}

	return false
}

// SetCommentCount gets a reference to the given utils.NullableInt64 and assigns it to the CommentCount field.
func (o *MergeRequestData) SetCommentCount(v int64) {
	o.CommentCount.Set(&v)
}
// SetCommentCountNil sets the value for CommentCount to be an explicit nil
func (o *MergeRequestData) SetCommentCountNil() {
	o.CommentCount.Set(nil)
}

// UnsetCommentCount ensures that no value is present for CommentCount, not even an explicit nil
func (o *MergeRequestData) UnsetCommentCount() {
	o.CommentCount.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetCreatedAt() int64 {
	if o == nil || utils.IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MergeRequestData) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given utils.NullableInt64 and assigns it to the CreatedAt field.
func (o *MergeRequestData) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *MergeRequestData) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *MergeRequestData) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDepotId returns the DepotId field value if set, zero value otherwise.
func (o *MergeRequestData) GetDepotId() int64 {
	if o == nil || utils.IsNil(o.DepotId) {
		var ret int64
		return ret
	}
	return *o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestData) GetDepotIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.DepotId) {
		return nil, false
	}
	return o.DepotId, true
}

// HasDepotId returns a boolean if a field has been set.
func (o *MergeRequestData) HasDepotId() bool {
	if o != nil && !utils.IsNil(o.DepotId) {
		return true
	}

	return false
}

// SetDepotId gets a reference to the given int64 and assigns it to the DepotId field.
func (o *MergeRequestData) SetDepotId(v int64) {
	o.DepotId = &v
}

// GetDescribe returns the Describe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetDescribe() string {
	if o == nil || utils.IsNil(o.Describe.Get()) {
		var ret string
		return ret
	}
	return *o.Describe.Get()
}

// GetDescribeOk returns a tuple with the Describe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetDescribeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Describe.Get(), o.Describe.IsSet()
}

// HasDescribe returns a boolean if a field has been set.
func (o *MergeRequestData) HasDescribe() bool {
	if o != nil && o.Describe.IsSet() {
		return true
	}

	return false
}

// SetDescribe gets a reference to the given utils.NullableString and assigns it to the Describe field.
func (o *MergeRequestData) SetDescribe(v string) {
	o.Describe.Set(&v)
}
// SetDescribeNil sets the value for Describe to be an explicit nil
func (o *MergeRequestData) SetDescribeNil() {
	o.Describe.Set(nil)
}

// UnsetDescribe ensures that no value is present for Describe, not even an explicit nil
func (o *MergeRequestData) UnsetDescribe() {
	o.Describe.Unset()
}

// GetGranted returns the Granted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetGranted() int64 {
	if o == nil || utils.IsNil(o.Granted.Get()) {
		var ret int64
		return ret
	}
	return *o.Granted.Get()
}

// GetGrantedOk returns a tuple with the Granted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetGrantedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Granted.Get(), o.Granted.IsSet()
}

// HasGranted returns a boolean if a field has been set.
func (o *MergeRequestData) HasGranted() bool {
	if o != nil && o.Granted.IsSet() {
		return true
	}

	return false
}

// SetGranted gets a reference to the given utils.NullableInt64 and assigns it to the Granted field.
func (o *MergeRequestData) SetGranted(v int64) {
	o.Granted.Set(&v)
}
// SetGrantedNil sets the value for Granted to be an explicit nil
func (o *MergeRequestData) SetGrantedNil() {
	o.Granted.Set(nil)
}

// UnsetGranted ensures that no value is present for Granted, not even an explicit nil
func (o *MergeRequestData) UnsetGranted() {
	o.Granted.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MergeRequestData) GetId() int64 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestData) GetIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MergeRequestData) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *MergeRequestData) SetId(v int64) {
	o.Id = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetLabelsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *MergeRequestData) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *MergeRequestData) SetLabels(v []string) {
	o.Labels = v
}

// GetMergeCommitSha returns the MergeCommitSha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetMergeCommitSha() string {
	if o == nil || utils.IsNil(o.MergeCommitSha.Get()) {
		var ret string
		return ret
	}
	return *o.MergeCommitSha.Get()
}

// GetMergeCommitShaOk returns a tuple with the MergeCommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetMergeCommitShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MergeCommitSha.Get(), o.MergeCommitSha.IsSet()
}

// HasMergeCommitSha returns a boolean if a field has been set.
func (o *MergeRequestData) HasMergeCommitSha() bool {
	if o != nil && o.MergeCommitSha.IsSet() {
		return true
	}

	return false
}

// SetMergeCommitSha gets a reference to the given utils.NullableString and assigns it to the MergeCommitSha field.
func (o *MergeRequestData) SetMergeCommitSha(v string) {
	o.MergeCommitSha.Set(&v)
}
// SetMergeCommitShaNil sets the value for MergeCommitSha to be an explicit nil
func (o *MergeRequestData) SetMergeCommitShaNil() {
	o.MergeCommitSha.Set(nil)
}

// UnsetMergeCommitSha ensures that no value is present for MergeCommitSha, not even an explicit nil
func (o *MergeRequestData) UnsetMergeCommitSha() {
	o.MergeCommitSha.Unset()
}

// GetMergeId returns the MergeId field value if set, zero value otherwise.
func (o *MergeRequestData) GetMergeId() int64 {
	if o == nil || utils.IsNil(o.MergeId) {
		var ret int64
		return ret
	}
	return *o.MergeId
}

// GetMergeIdOk returns a tuple with the MergeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestData) GetMergeIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.MergeId) {
		return nil, false
	}
	return o.MergeId, true
}

// HasMergeId returns a boolean if a field has been set.
func (o *MergeRequestData) HasMergeId() bool {
	if o != nil && !utils.IsNil(o.MergeId) {
		return true
	}

	return false
}

// SetMergeId gets a reference to the given int64 and assigns it to the MergeId field.
func (o *MergeRequestData) SetMergeId(v int64) {
	o.MergeId = &v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetPath() string {
	if o == nil || utils.IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *MergeRequestData) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given utils.NullableString and assigns it to the Path field.
func (o *MergeRequestData) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *MergeRequestData) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *MergeRequestData) UnsetPath() {
	o.Path.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *MergeRequestData) GetProjectId() int64 {
	if o == nil || utils.IsNil(o.ProjectId) {
		var ret int64
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestData) GetProjectIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *MergeRequestData) HasProjectId() bool {
	if o != nil && !utils.IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int64 and assigns it to the ProjectId field.
func (o *MergeRequestData) SetProjectId(v int64) {
	o.ProjectId = &v
}

// GetReminded returns the Reminded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetReminded() bool {
	if o == nil || utils.IsNil(o.Reminded.Get()) {
		var ret bool
		return ret
	}
	return *o.Reminded.Get()
}

// GetRemindedOk returns a tuple with the Reminded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetRemindedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reminded.Get(), o.Reminded.IsSet()
}

// HasReminded returns a boolean if a field has been set.
func (o *MergeRequestData) HasReminded() bool {
	if o != nil && o.Reminded.IsSet() {
		return true
	}

	return false
}

// SetReminded gets a reference to the given utils.NullableBool and assigns it to the Reminded field.
func (o *MergeRequestData) SetReminded(v bool) {
	o.Reminded.Set(&v)
}
// SetRemindedNil sets the value for Reminded to be an explicit nil
func (o *MergeRequestData) SetRemindedNil() {
	o.Reminded.Set(nil)
}

// UnsetReminded ensures that no value is present for Reminded, not even an explicit nil
func (o *MergeRequestData) UnsetReminded() {
	o.Reminded.Unset()
}

// GetReviewers returns the Reviewers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetReviewers() []CodingUser {
	if o == nil {
		var ret []CodingUser
		return ret
	}
	return o.Reviewers
}

// GetReviewersOk returns a tuple with the Reviewers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetReviewersOk() ([]CodingUser, bool) {
	if o == nil || utils.IsNil(o.Reviewers) {
		return nil, false
	}
	return o.Reviewers, true
}

// HasReviewers returns a boolean if a field has been set.
func (o *MergeRequestData) HasReviewers() bool {
	if o != nil && !utils.IsNil(o.Reviewers) {
		return true
	}

	return false
}

// SetReviewers gets a reference to the given []CodingUser and assigns it to the Reviewers field.
func (o *MergeRequestData) SetReviewers(v []CodingUser) {
	o.Reviewers = v
}

// GetSourceBranch returns the SourceBranch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetSourceBranch() string {
	if o == nil || utils.IsNil(o.SourceBranch.Get()) {
		var ret string
		return ret
	}
	return *o.SourceBranch.Get()
}

// GetSourceBranchOk returns a tuple with the SourceBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetSourceBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceBranch.Get(), o.SourceBranch.IsSet()
}

// HasSourceBranch returns a boolean if a field has been set.
func (o *MergeRequestData) HasSourceBranch() bool {
	if o != nil && o.SourceBranch.IsSet() {
		return true
	}

	return false
}

// SetSourceBranch gets a reference to the given utils.NullableString and assigns it to the SourceBranch field.
func (o *MergeRequestData) SetSourceBranch(v string) {
	o.SourceBranch.Set(&v)
}
// SetSourceBranchNil sets the value for SourceBranch to be an explicit nil
func (o *MergeRequestData) SetSourceBranchNil() {
	o.SourceBranch.Set(nil)
}

// UnsetSourceBranch ensures that no value is present for SourceBranch, not even an explicit nil
func (o *MergeRequestData) UnsetSourceBranch() {
	o.SourceBranch.Unset()
}

// GetSourceBranchSha returns the SourceBranchSha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetSourceBranchSha() string {
	if o == nil || utils.IsNil(o.SourceBranchSha.Get()) {
		var ret string
		return ret
	}
	return *o.SourceBranchSha.Get()
}

// GetSourceBranchShaOk returns a tuple with the SourceBranchSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetSourceBranchShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceBranchSha.Get(), o.SourceBranchSha.IsSet()
}

// HasSourceBranchSha returns a boolean if a field has been set.
func (o *MergeRequestData) HasSourceBranchSha() bool {
	if o != nil && o.SourceBranchSha.IsSet() {
		return true
	}

	return false
}

// SetSourceBranchSha gets a reference to the given utils.NullableString and assigns it to the SourceBranchSha field.
func (o *MergeRequestData) SetSourceBranchSha(v string) {
	o.SourceBranchSha.Set(&v)
}
// SetSourceBranchShaNil sets the value for SourceBranchSha to be an explicit nil
func (o *MergeRequestData) SetSourceBranchShaNil() {
	o.SourceBranchSha.Set(nil)
}

// UnsetSourceBranchSha ensures that no value is present for SourceBranchSha, not even an explicit nil
func (o *MergeRequestData) UnsetSourceBranchSha() {
	o.SourceBranchSha.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetStatus() string {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *MergeRequestData) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given utils.NullableString and assigns it to the Status field.
func (o *MergeRequestData) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *MergeRequestData) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *MergeRequestData) UnsetStatus() {
	o.Status.Unset()
}

// GetStickingPoint returns the StickingPoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetStickingPoint() string {
	if o == nil || utils.IsNil(o.StickingPoint.Get()) {
		var ret string
		return ret
	}
	return *o.StickingPoint.Get()
}

// GetStickingPointOk returns a tuple with the StickingPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetStickingPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StickingPoint.Get(), o.StickingPoint.IsSet()
}

// HasStickingPoint returns a boolean if a field has been set.
func (o *MergeRequestData) HasStickingPoint() bool {
	if o != nil && o.StickingPoint.IsSet() {
		return true
	}

	return false
}

// SetStickingPoint gets a reference to the given utils.NullableString and assigns it to the StickingPoint field.
func (o *MergeRequestData) SetStickingPoint(v string) {
	o.StickingPoint.Set(&v)
}
// SetStickingPointNil sets the value for StickingPoint to be an explicit nil
func (o *MergeRequestData) SetStickingPointNil() {
	o.StickingPoint.Set(nil)
}

// UnsetStickingPoint ensures that no value is present for StickingPoint, not even an explicit nil
func (o *MergeRequestData) UnsetStickingPoint() {
	o.StickingPoint.Unset()
}

// GetTargetBranch returns the TargetBranch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetTargetBranch() string {
	if o == nil || utils.IsNil(o.TargetBranch.Get()) {
		var ret string
		return ret
	}
	return *o.TargetBranch.Get()
}

// GetTargetBranchOk returns a tuple with the TargetBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetTargetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetBranch.Get(), o.TargetBranch.IsSet()
}

// HasTargetBranch returns a boolean if a field has been set.
func (o *MergeRequestData) HasTargetBranch() bool {
	if o != nil && o.TargetBranch.IsSet() {
		return true
	}

	return false
}

// SetTargetBranch gets a reference to the given utils.NullableString and assigns it to the TargetBranch field.
func (o *MergeRequestData) SetTargetBranch(v string) {
	o.TargetBranch.Set(&v)
}
// SetTargetBranchNil sets the value for TargetBranch to be an explicit nil
func (o *MergeRequestData) SetTargetBranchNil() {
	o.TargetBranch.Set(nil)
}

// UnsetTargetBranch ensures that no value is present for TargetBranch, not even an explicit nil
func (o *MergeRequestData) UnsetTargetBranch() {
	o.TargetBranch.Unset()
}

// GetTargetBranchProtected returns the TargetBranchProtected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetTargetBranchProtected() bool {
	if o == nil || utils.IsNil(o.TargetBranchProtected.Get()) {
		var ret bool
		return ret
	}
	return *o.TargetBranchProtected.Get()
}

// GetTargetBranchProtectedOk returns a tuple with the TargetBranchProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetTargetBranchProtectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetBranchProtected.Get(), o.TargetBranchProtected.IsSet()
}

// HasTargetBranchProtected returns a boolean if a field has been set.
func (o *MergeRequestData) HasTargetBranchProtected() bool {
	if o != nil && o.TargetBranchProtected.IsSet() {
		return true
	}

	return false
}

// SetTargetBranchProtected gets a reference to the given utils.NullableBool and assigns it to the TargetBranchProtected field.
func (o *MergeRequestData) SetTargetBranchProtected(v bool) {
	o.TargetBranchProtected.Set(&v)
}
// SetTargetBranchProtectedNil sets the value for TargetBranchProtected to be an explicit nil
func (o *MergeRequestData) SetTargetBranchProtectedNil() {
	o.TargetBranchProtected.Set(nil)
}

// UnsetTargetBranchProtected ensures that no value is present for TargetBranchProtected, not even an explicit nil
func (o *MergeRequestData) UnsetTargetBranchProtected() {
	o.TargetBranchProtected.Unset()
}

// GetTargetBranchSha returns the TargetBranchSha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetTargetBranchSha() string {
	if o == nil || utils.IsNil(o.TargetBranchSha.Get()) {
		var ret string
		return ret
	}
	return *o.TargetBranchSha.Get()
}

// GetTargetBranchShaOk returns a tuple with the TargetBranchSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetTargetBranchShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetBranchSha.Get(), o.TargetBranchSha.IsSet()
}

// HasTargetBranchSha returns a boolean if a field has been set.
func (o *MergeRequestData) HasTargetBranchSha() bool {
	if o != nil && o.TargetBranchSha.IsSet() {
		return true
	}

	return false
}

// SetTargetBranchSha gets a reference to the given utils.NullableString and assigns it to the TargetBranchSha field.
func (o *MergeRequestData) SetTargetBranchSha(v string) {
	o.TargetBranchSha.Set(&v)
}
// SetTargetBranchShaNil sets the value for TargetBranchSha to be an explicit nil
func (o *MergeRequestData) SetTargetBranchShaNil() {
	o.TargetBranchSha.Set(nil)
}

// UnsetTargetBranchSha ensures that no value is present for TargetBranchSha, not even an explicit nil
func (o *MergeRequestData) UnsetTargetBranchSha() {
	o.TargetBranchSha.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *MergeRequestData) GetTitle() string {
	if o == nil || utils.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestData) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *MergeRequestData) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *MergeRequestData) SetTitle(v string) {
	o.Title = &v
}

// GetUpdateAt returns the UpdateAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestData) GetUpdateAt() int64 {
	if o == nil || utils.IsNil(o.UpdateAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdateAt.Get()
}

// GetUpdateAtOk returns a tuple with the UpdateAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestData) GetUpdateAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdateAt.Get(), o.UpdateAt.IsSet()
}

// HasUpdateAt returns a boolean if a field has been set.
func (o *MergeRequestData) HasUpdateAt() bool {
	if o != nil && o.UpdateAt.IsSet() {
		return true
	}

	return false
}

// SetUpdateAt gets a reference to the given utils.NullableInt64 and assigns it to the UpdateAt field.
func (o *MergeRequestData) SetUpdateAt(v int64) {
	o.UpdateAt.Set(&v)
}
// SetUpdateAtNil sets the value for UpdateAt to be an explicit nil
func (o *MergeRequestData) SetUpdateAtNil() {
	o.UpdateAt.Set(nil)
}

// UnsetUpdateAt ensures that no value is present for UpdateAt, not even an explicit nil
func (o *MergeRequestData) UnsetUpdateAt() {
	o.UpdateAt.Unset()
}

func (o MergeRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MergeRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ActionAt.IsSet() {
		toSerialize["ActionAt"] = o.ActionAt.Get()
	}
	if !utils.IsNil(o.ActionAuthor) {
		toSerialize["ActionAuthor"] = o.ActionAuthor
	}
	if !utils.IsNil(o.Author) {
		toSerialize["Author"] = o.Author
	}
	if o.BaseSha.IsSet() {
		toSerialize["BaseSha"] = o.BaseSha.Get()
	}
	if o.CommentCount.IsSet() {
		toSerialize["CommentCount"] = o.CommentCount.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["CreatedAt"] = o.CreatedAt.Get()
	}
	if !utils.IsNil(o.DepotId) {
		toSerialize["DepotId"] = o.DepotId
	}
	if o.Describe.IsSet() {
		toSerialize["Describe"] = o.Describe.Get()
	}
	if o.Granted.IsSet() {
		toSerialize["Granted"] = o.Granted.Get()
	}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if o.Labels != nil {
		toSerialize["Labels"] = o.Labels
	}
	if o.MergeCommitSha.IsSet() {
		toSerialize["MergeCommitSha"] = o.MergeCommitSha.Get()
	}
	if !utils.IsNil(o.MergeId) {
		toSerialize["MergeId"] = o.MergeId
	}
	if o.Path.IsSet() {
		toSerialize["Path"] = o.Path.Get()
	}
	if !utils.IsNil(o.ProjectId) {
		toSerialize["ProjectId"] = o.ProjectId
	}
	if o.Reminded.IsSet() {
		toSerialize["Reminded"] = o.Reminded.Get()
	}
	if o.Reviewers != nil {
		toSerialize["Reviewers"] = o.Reviewers
	}
	if o.SourceBranch.IsSet() {
		toSerialize["SourceBranch"] = o.SourceBranch.Get()
	}
	if o.SourceBranchSha.IsSet() {
		toSerialize["SourceBranchSha"] = o.SourceBranchSha.Get()
	}
	if o.Status.IsSet() {
		toSerialize["Status"] = o.Status.Get()
	}
	if o.StickingPoint.IsSet() {
		toSerialize["StickingPoint"] = o.StickingPoint.Get()
	}
	if o.TargetBranch.IsSet() {
		toSerialize["TargetBranch"] = o.TargetBranch.Get()
	}
	if o.TargetBranchProtected.IsSet() {
		toSerialize["TargetBranchProtected"] = o.TargetBranchProtected.Get()
	}
	if o.TargetBranchSha.IsSet() {
		toSerialize["TargetBranchSha"] = o.TargetBranchSha.Get()
	}
	if !utils.IsNil(o.Title) {
		toSerialize["Title"] = o.Title
	}
	if o.UpdateAt.IsSet() {
		toSerialize["UpdateAt"] = o.UpdateAt.Get()
	}
	return toSerialize, nil
}

type NullableMergeRequestData struct {
	value *MergeRequestData
	isSet bool
}

func (v NullableMergeRequestData) Get() *MergeRequestData {
	return v.value
}

func (v *NullableMergeRequestData) Set(val *MergeRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeRequestData(val *MergeRequestData) *NullableMergeRequestData {
	return &NullableMergeRequestData{value: val, isSet: true}
}

func (v NullableMergeRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

