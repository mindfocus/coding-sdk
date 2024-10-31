/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the Release type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Release{}

// Release git 版本
type Release struct {
	// 内容
	Body utils.NullableString `json:"Body,omitempty"`
	// commit Sha 值
	CommitSha utils.NullableString `json:"CommitSha,omitempty"`
	// 创建时间
	CreatedAt utils.NullableInt64 `json:"CreatedAt,omitempty"`
	// 创建者
	CreatorId utils.NullableInt64 `json:"CreatorId,omitempty"`
	// 仓库 Id
	DepotId utils.NullableInt64 `json:"DepotId,omitempty"`
	// html内容
	Html utils.NullableString `json:"Html,omitempty"`
	// 版本 Id
	Id utils.NullableInt64 `json:"Id,omitempty"`
	// 是否预发布
	Pre utils.NullableBool `json:"Pre,omitempty"`
	// 项目 Id
	ProjectId utils.NullableInt64 `json:"ProjectId,omitempty"`
	// 版本序号Id
	ReleaseId utils.NullableInt64 `json:"ReleaseId,omitempty"`
	// 标签名字
	TagName utils.NullableString `json:"TagName,omitempty"`
	// 目标 commit Sha 值
	TargetCommitBranch utils.NullableString `json:"TargetCommitBranch,omitempty"`
	// 标题
	Title utils.NullableString `json:"Title,omitempty"`
	// 更新时间
	UpdatedAt utils.NullableInt64 `json:"UpdatedAt,omitempty"`
}

// NewRelease instantiates a new Release object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelease() *Release {
	this := Release{}
	var body string = ""
	this.Body = *utils.NewNullableString(&body)
	var commitSha string = ""
	this.CommitSha = *utils.NewNullableString(&commitSha)
	var html string = ""
	this.Html = *utils.NewNullableString(&html)
	var pre bool = false
	this.Pre = *utils.NewNullableBool(&pre)
	var tagName string = ""
	this.TagName = *utils.NewNullableString(&tagName)
	var targetCommitBranch string = ""
	this.TargetCommitBranch = *utils.NewNullableString(&targetCommitBranch)
	var title string = ""
	this.Title = *utils.NewNullableString(&title)
	return &this
}

// NewReleaseWithDefaults instantiates a new Release object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseWithDefaults() *Release {
	this := Release{}
	var body string = ""
	this.Body = *utils.NewNullableString(&body)
	var commitSha string = ""
	this.CommitSha = *utils.NewNullableString(&commitSha)
	var html string = ""
	this.Html = *utils.NewNullableString(&html)
	var pre bool = false
	this.Pre = *utils.NewNullableBool(&pre)
	var tagName string = ""
	this.TagName = *utils.NewNullableString(&tagName)
	var targetCommitBranch string = ""
	this.TargetCommitBranch = *utils.NewNullableString(&targetCommitBranch)
	var title string = ""
	this.Title = *utils.NewNullableString(&title)
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetBody() string {
	if o == nil || utils.IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *Release) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given utils.NullableString and assigns it to the Body field.
func (o *Release) SetBody(v string) {
	o.Body.Set(&v)
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *Release) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *Release) UnsetBody() {
	o.Body.Unset()
}

// GetCommitSha returns the CommitSha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetCommitSha() string {
	if o == nil || utils.IsNil(o.CommitSha.Get()) {
		var ret string
		return ret
	}
	return *o.CommitSha.Get()
}

// GetCommitShaOk returns a tuple with the CommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetCommitShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitSha.Get(), o.CommitSha.IsSet()
}

// HasCommitSha returns a boolean if a field has been set.
func (o *Release) HasCommitSha() bool {
	if o != nil && o.CommitSha.IsSet() {
		return true
	}

	return false
}

// SetCommitSha gets a reference to the given utils.NullableString and assigns it to the CommitSha field.
func (o *Release) SetCommitSha(v string) {
	o.CommitSha.Set(&v)
}
// SetCommitShaNil sets the value for CommitSha to be an explicit nil
func (o *Release) SetCommitShaNil() {
	o.CommitSha.Set(nil)
}

// UnsetCommitSha ensures that no value is present for CommitSha, not even an explicit nil
func (o *Release) UnsetCommitSha() {
	o.CommitSha.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetCreatedAt() int64 {
	if o == nil || utils.IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Release) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given utils.NullableInt64 and assigns it to the CreatedAt field.
func (o *Release) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Release) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Release) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetCreatorId() int64 {
	if o == nil || utils.IsNil(o.CreatorId.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatorId.Get()
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetCreatorIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorId.Get(), o.CreatorId.IsSet()
}

// HasCreatorId returns a boolean if a field has been set.
func (o *Release) HasCreatorId() bool {
	if o != nil && o.CreatorId.IsSet() {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given utils.NullableInt64 and assigns it to the CreatorId field.
func (o *Release) SetCreatorId(v int64) {
	o.CreatorId.Set(&v)
}
// SetCreatorIdNil sets the value for CreatorId to be an explicit nil
func (o *Release) SetCreatorIdNil() {
	o.CreatorId.Set(nil)
}

// UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
func (o *Release) UnsetCreatorId() {
	o.CreatorId.Unset()
}

// GetDepotId returns the DepotId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetDepotId() int64 {
	if o == nil || utils.IsNil(o.DepotId.Get()) {
		var ret int64
		return ret
	}
	return *o.DepotId.Get()
}

// GetDepotIdOk returns a tuple with the DepotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DepotId.Get(), o.DepotId.IsSet()
}

// HasDepotId returns a boolean if a field has been set.
func (o *Release) HasDepotId() bool {
	if o != nil && o.DepotId.IsSet() {
		return true
	}

	return false
}

// SetDepotId gets a reference to the given utils.NullableInt64 and assigns it to the DepotId field.
func (o *Release) SetDepotId(v int64) {
	o.DepotId.Set(&v)
}
// SetDepotIdNil sets the value for DepotId to be an explicit nil
func (o *Release) SetDepotIdNil() {
	o.DepotId.Set(nil)
}

// UnsetDepotId ensures that no value is present for DepotId, not even an explicit nil
func (o *Release) UnsetDepotId() {
	o.DepotId.Unset()
}

// GetHtml returns the Html field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetHtml() string {
	if o == nil || utils.IsNil(o.Html.Get()) {
		var ret string
		return ret
	}
	return *o.Html.Get()
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Html.Get(), o.Html.IsSet()
}

// HasHtml returns a boolean if a field has been set.
func (o *Release) HasHtml() bool {
	if o != nil && o.Html.IsSet() {
		return true
	}

	return false
}

// SetHtml gets a reference to the given utils.NullableString and assigns it to the Html field.
func (o *Release) SetHtml(v string) {
	o.Html.Set(&v)
}
// SetHtmlNil sets the value for Html to be an explicit nil
func (o *Release) SetHtmlNil() {
	o.Html.Set(nil)
}

// UnsetHtml ensures that no value is present for Html, not even an explicit nil
func (o *Release) UnsetHtml() {
	o.Html.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetId() int64 {
	if o == nil || utils.IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Release) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given utils.NullableInt64 and assigns it to the Id field.
func (o *Release) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Release) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Release) UnsetId() {
	o.Id.Unset()
}

// GetPre returns the Pre field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetPre() bool {
	if o == nil || utils.IsNil(o.Pre.Get()) {
		var ret bool
		return ret
	}
	return *o.Pre.Get()
}

// GetPreOk returns a tuple with the Pre field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetPreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pre.Get(), o.Pre.IsSet()
}

// HasPre returns a boolean if a field has been set.
func (o *Release) HasPre() bool {
	if o != nil && o.Pre.IsSet() {
		return true
	}

	return false
}

// SetPre gets a reference to the given utils.NullableBool and assigns it to the Pre field.
func (o *Release) SetPre(v bool) {
	o.Pre.Set(&v)
}
// SetPreNil sets the value for Pre to be an explicit nil
func (o *Release) SetPreNil() {
	o.Pre.Set(nil)
}

// UnsetPre ensures that no value is present for Pre, not even an explicit nil
func (o *Release) UnsetPre() {
	o.Pre.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetProjectId() int64 {
	if o == nil || utils.IsNil(o.ProjectId.Get()) {
		var ret int64
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetProjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *Release) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given utils.NullableInt64 and assigns it to the ProjectId field.
func (o *Release) SetProjectId(v int64) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *Release) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *Release) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetReleaseId returns the ReleaseId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetReleaseId() int64 {
	if o == nil || utils.IsNil(o.ReleaseId.Get()) {
		var ret int64
		return ret
	}
	return *o.ReleaseId.Get()
}

// GetReleaseIdOk returns a tuple with the ReleaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetReleaseIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseId.Get(), o.ReleaseId.IsSet()
}

// HasReleaseId returns a boolean if a field has been set.
func (o *Release) HasReleaseId() bool {
	if o != nil && o.ReleaseId.IsSet() {
		return true
	}

	return false
}

// SetReleaseId gets a reference to the given utils.NullableInt64 and assigns it to the ReleaseId field.
func (o *Release) SetReleaseId(v int64) {
	o.ReleaseId.Set(&v)
}
// SetReleaseIdNil sets the value for ReleaseId to be an explicit nil
func (o *Release) SetReleaseIdNil() {
	o.ReleaseId.Set(nil)
}

// UnsetReleaseId ensures that no value is present for ReleaseId, not even an explicit nil
func (o *Release) UnsetReleaseId() {
	o.ReleaseId.Unset()
}

// GetTagName returns the TagName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetTagName() string {
	if o == nil || utils.IsNil(o.TagName.Get()) {
		var ret string
		return ret
	}
	return *o.TagName.Get()
}

// GetTagNameOk returns a tuple with the TagName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetTagNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TagName.Get(), o.TagName.IsSet()
}

// HasTagName returns a boolean if a field has been set.
func (o *Release) HasTagName() bool {
	if o != nil && o.TagName.IsSet() {
		return true
	}

	return false
}

// SetTagName gets a reference to the given utils.NullableString and assigns it to the TagName field.
func (o *Release) SetTagName(v string) {
	o.TagName.Set(&v)
}
// SetTagNameNil sets the value for TagName to be an explicit nil
func (o *Release) SetTagNameNil() {
	o.TagName.Set(nil)
}

// UnsetTagName ensures that no value is present for TagName, not even an explicit nil
func (o *Release) UnsetTagName() {
	o.TagName.Unset()
}

// GetTargetCommitBranch returns the TargetCommitBranch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetTargetCommitBranch() string {
	if o == nil || utils.IsNil(o.TargetCommitBranch.Get()) {
		var ret string
		return ret
	}
	return *o.TargetCommitBranch.Get()
}

// GetTargetCommitBranchOk returns a tuple with the TargetCommitBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetTargetCommitBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetCommitBranch.Get(), o.TargetCommitBranch.IsSet()
}

// HasTargetCommitBranch returns a boolean if a field has been set.
func (o *Release) HasTargetCommitBranch() bool {
	if o != nil && o.TargetCommitBranch.IsSet() {
		return true
	}

	return false
}

// SetTargetCommitBranch gets a reference to the given utils.NullableString and assigns it to the TargetCommitBranch field.
func (o *Release) SetTargetCommitBranch(v string) {
	o.TargetCommitBranch.Set(&v)
}
// SetTargetCommitBranchNil sets the value for TargetCommitBranch to be an explicit nil
func (o *Release) SetTargetCommitBranchNil() {
	o.TargetCommitBranch.Set(nil)
}

// UnsetTargetCommitBranch ensures that no value is present for TargetCommitBranch, not even an explicit nil
func (o *Release) UnsetTargetCommitBranch() {
	o.TargetCommitBranch.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetTitle() string {
	if o == nil || utils.IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Release) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given utils.NullableString and assigns it to the Title field.
func (o *Release) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Release) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Release) UnsetTitle() {
	o.Title.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetUpdatedAt() int64 {
	if o == nil || utils.IsNil(o.UpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Release) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given utils.NullableInt64 and assigns it to the UpdatedAt field.
func (o *Release) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *Release) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *Release) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

func (o Release) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Release) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Body.IsSet() {
		toSerialize["Body"] = o.Body.Get()
	}
	if o.CommitSha.IsSet() {
		toSerialize["CommitSha"] = o.CommitSha.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["CreatedAt"] = o.CreatedAt.Get()
	}
	if o.CreatorId.IsSet() {
		toSerialize["CreatorId"] = o.CreatorId.Get()
	}
	if o.DepotId.IsSet() {
		toSerialize["DepotId"] = o.DepotId.Get()
	}
	if o.Html.IsSet() {
		toSerialize["Html"] = o.Html.Get()
	}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.Pre.IsSet() {
		toSerialize["Pre"] = o.Pre.Get()
	}
	if o.ProjectId.IsSet() {
		toSerialize["ProjectId"] = o.ProjectId.Get()
	}
	if o.ReleaseId.IsSet() {
		toSerialize["ReleaseId"] = o.ReleaseId.Get()
	}
	if o.TagName.IsSet() {
		toSerialize["TagName"] = o.TagName.Get()
	}
	if o.TargetCommitBranch.IsSet() {
		toSerialize["TargetCommitBranch"] = o.TargetCommitBranch.Get()
	}
	if o.Title.IsSet() {
		toSerialize["Title"] = o.Title.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["UpdatedAt"] = o.UpdatedAt.Get()
	}
	return toSerialize, nil
}

type NullableRelease struct {
	value *Release
	isSet bool
}

func (v NullableRelease) Get() *Release {
	return v.value
}

func (v *NullableRelease) Set(val *Release) {
	v.value = val
	v.isSet = true
}

func (v NullableRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelease(val *Release) *NullableRelease {
	return &NullableRelease{value: val, isSet: true}
}

func (v NullableRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


