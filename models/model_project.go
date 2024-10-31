/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"github.com/mindfocus/coding-sdk/utils"
)

// checks if the Project type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Project{}

// Project 项目信息
type Project struct {
	// 是否压缩
	Archived utils.NullableBool `json:"Archived,omitempty"`
	// 创建时间
	CreatedAt utils.NullableInt64 `json:"CreatedAt,omitempty"`
	// 描述
	Description utils.NullableString `json:"Description,omitempty"`
	// 显示名称
	DisplayName utils.NullableString `json:"DisplayName,omitempty"`
	// 项目结束时间
	EndDate utils.NullableInt64 `json:"EndDate,omitempty"`
	// 图标
	Icon utils.NullableString `json:"Icon,omitempty"`
	// 项目 ID
	Id *int32 `json:"Id,omitempty"`
	// 是否为模板项目
	IsDemo utils.NullableBool `json:"IsDemo,omitempty"`
	// 最大团员数
	MaxMember utils.NullableInt32 `json:"MaxMember,omitempty"`
	// 名称
	Name utils.NullableString `json:"Name,omitempty"`
	// 项目开始时间
	StartDate utils.NullableInt32 `json:"StartDate,omitempty"`
	// 状态 默认都为1
	Status utils.NullableInt32 `json:"Status,omitempty"`
	// 团队 ID
	TeamId utils.NullableInt32 `json:"TeamId,omitempty"`
	// 团队所有者 ID
	TeamOwnerId utils.NullableInt32 `json:"TeamOwnerId,omitempty"`
	// 项目类型, 项目集为0，公开项目为1，私密项目为2
	Type utils.NullableInt32 `json:"Type,omitempty"`
	// 更新时间
	UpdatedAt utils.NullableInt64 `json:"UpdatedAt,omitempty"`
	// 个人所有者 ID
	UserOwnerId utils.NullableInt32 `json:"UserOwnerId,omitempty"`
	// 项目集id
	ProgramIds []uint64 `json:"ProgramIds,omitempty"`
}

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject() *Project {
	this := Project{}
	var archived bool = false
	this.Archived = *utils.NewNullableBool(&archived)
	var description string = ""
	this.Description = *utils.NewNullableString(&description)
	var displayName string = ""
	this.DisplayName = *utils.NewNullableString(&displayName)
	var icon string = ""
	this.Icon = *utils.NewNullableString(&icon)
	var isDemo bool = false
	this.IsDemo = *utils.NewNullableBool(&isDemo)
	var name string = ""
	this.Name = *utils.NewNullableString(&name)
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	var archived bool = false
	this.Archived = *utils.NewNullableBool(&archived)
	var description string = ""
	this.Description = *utils.NewNullableString(&description)
	var displayName string = ""
	this.DisplayName = *utils.NewNullableString(&displayName)
	var icon string = ""
	this.Icon = *utils.NewNullableString(&icon)
	var isDemo bool = false
	this.IsDemo = *utils.NewNullableBool(&isDemo)
	var name string = ""
	this.Name = *utils.NewNullableString(&name)
	return &this
}

// GetArchived returns the Archived field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetArchived() bool {
	if o == nil || utils.IsNil(o.Archived.Get()) {
		var ret bool
		return ret
	}
	return *o.Archived.Get()
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Archived.Get(), o.Archived.IsSet()
}

// HasArchived returns a boolean if a field has been set.
func (o *Project) HasArchived() bool {
	if o != nil && o.Archived.IsSet() {
		return true
	}

	return false
}

// SetArchived gets a reference to the given utils.NullableBool and assigns it to the Archived field.
func (o *Project) SetArchived(v bool) {
	o.Archived.Set(&v)
}
// SetArchivedNil sets the value for Archived to be an explicit nil
func (o *Project) SetArchivedNil() {
	o.Archived.Set(nil)
}

// UnsetArchived ensures that no value is present for Archived, not even an explicit nil
func (o *Project) UnsetArchived() {
	o.Archived.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetCreatedAt() *int64 {
	if o == nil || utils.IsNil(o.CreatedAt.Get()) {
		var ret *int64 = nil
		return ret
	}
	return o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Project) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given utils.NullableInt32 and assigns it to the CreatedAt field.
func (o *Project) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Project) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Project) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Project) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given utils.NullableString and assigns it to the Description field.
func (o *Project) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Project) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Project) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetDisplayName() string {
	if o == nil || utils.IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Project) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given utils.NullableString and assigns it to the DisplayName field.
func (o *Project) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *Project) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *Project) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetEndDate() *int64 {
	if o == nil || utils.IsNil(o.EndDate.Get()) {
		var ret *int64 = nil
		return ret
	}
	return o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetEndDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *Project) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given utils.NullableInt32 and assigns it to the EndDate field.
func (o *Project) SetEndDate(v int64) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *Project) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *Project) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetIcon() string {
	if o == nil || utils.IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *Project) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given utils.NullableString and assigns it to the Icon field.
func (o *Project) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *Project) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *Project) UnsetIcon() {
	o.Icon.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Project) GetId() int32 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Project) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Project) SetId(v int32) {
	o.Id = &v
}

// GetIsDemo returns the IsDemo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetIsDemo() bool {
	if o == nil || utils.IsNil(o.IsDemo.Get()) {
		var ret bool
		return ret
	}
	return *o.IsDemo.Get()
}

// GetIsDemoOk returns a tuple with the IsDemo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetIsDemoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDemo.Get(), o.IsDemo.IsSet()
}

// HasIsDemo returns a boolean if a field has been set.
func (o *Project) HasIsDemo() bool {
	if o != nil && o.IsDemo.IsSet() {
		return true
	}

	return false
}

// SetIsDemo gets a reference to the given utils.NullableBool and assigns it to the IsDemo field.
func (o *Project) SetIsDemo(v bool) {
	o.IsDemo.Set(&v)
}
// SetIsDemoNil sets the value for IsDemo to be an explicit nil
func (o *Project) SetIsDemoNil() {
	o.IsDemo.Set(nil)
}

// UnsetIsDemo ensures that no value is present for IsDemo, not even an explicit nil
func (o *Project) UnsetIsDemo() {
	o.IsDemo.Unset()
}

// GetMaxMember returns the MaxMember field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetMaxMember() int32 {
	if o == nil || utils.IsNil(o.MaxMember.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxMember.Get()
}

// GetMaxMemberOk returns a tuple with the MaxMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetMaxMemberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxMember.Get(), o.MaxMember.IsSet()
}

// HasMaxMember returns a boolean if a field has been set.
func (o *Project) HasMaxMember() bool {
	if o != nil && o.MaxMember.IsSet() {
		return true
	}

	return false
}

// SetMaxMember gets a reference to the given utils.NullableInt32 and assigns it to the MaxMember field.
func (o *Project) SetMaxMember(v int32) {
	o.MaxMember.Set(&v)
}
// SetMaxMemberNil sets the value for MaxMember to be an explicit nil
func (o *Project) SetMaxMemberNil() {
	o.MaxMember.Set(nil)
}

// UnsetMaxMember ensures that no value is present for MaxMember, not even an explicit nil
func (o *Project) UnsetMaxMember() {
	o.MaxMember.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetName() string {
	if o == nil || utils.IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Project) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given utils.NullableString and assigns it to the Name field.
func (o *Project) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Project) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Project) UnsetName() {
	o.Name.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetStartDate() int32 {
	if o == nil || utils.IsNil(o.StartDate.Get()) {
		var ret int32
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetStartDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *Project) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given utils.NullableInt32 and assigns it to the StartDate field.
func (o *Project) SetStartDate(v int32) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *Project) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *Project) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetStatus() int32 {
	if o == nil || utils.IsNil(o.Status.Get()) {
		var ret int32
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Project) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given utils.NullableInt32 and assigns it to the Status field.
func (o *Project) SetStatus(v int32) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Project) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Project) UnsetStatus() {
	o.Status.Unset()
}

// GetTeamId returns the TeamId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetTeamId() int32 {
	if o == nil || utils.IsNil(o.TeamId.Get()) {
		var ret int32
		return ret
	}
	return *o.TeamId.Get()
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetTeamIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TeamId.Get(), o.TeamId.IsSet()
}

// HasTeamId returns a boolean if a field has been set.
func (o *Project) HasTeamId() bool {
	if o != nil && o.TeamId.IsSet() {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given utils.NullableInt32 and assigns it to the TeamId field.
func (o *Project) SetTeamId(v int32) {
	o.TeamId.Set(&v)
}
// SetTeamIdNil sets the value for TeamId to be an explicit nil
func (o *Project) SetTeamIdNil() {
	o.TeamId.Set(nil)
}

// UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil
func (o *Project) UnsetTeamId() {
	o.TeamId.Unset()
}

// GetTeamOwnerId returns the TeamOwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetTeamOwnerId() int32 {
	if o == nil || utils.IsNil(o.TeamOwnerId.Get()) {
		var ret int32
		return ret
	}
	return *o.TeamOwnerId.Get()
}

// GetTeamOwnerIdOk returns a tuple with the TeamOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetTeamOwnerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TeamOwnerId.Get(), o.TeamOwnerId.IsSet()
}

// HasTeamOwnerId returns a boolean if a field has been set.
func (o *Project) HasTeamOwnerId() bool {
	if o != nil && o.TeamOwnerId.IsSet() {
		return true
	}

	return false
}

// SetTeamOwnerId gets a reference to the given utils.NullableInt32 and assigns it to the TeamOwnerId field.
func (o *Project) SetTeamOwnerId(v int32) {
	o.TeamOwnerId.Set(&v)
}
// SetTeamOwnerIdNil sets the value for TeamOwnerId to be an explicit nil
func (o *Project) SetTeamOwnerIdNil() {
	o.TeamOwnerId.Set(nil)
}

// UnsetTeamOwnerId ensures that no value is present for TeamOwnerId, not even an explicit nil
func (o *Project) UnsetTeamOwnerId() {
	o.TeamOwnerId.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetType() int32 {
	if o == nil || utils.IsNil(o.Type.Get()) {
		var ret int32
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Project) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given utils.NullableInt32 and assigns it to the Type field.
func (o *Project) SetType(v int32) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Project) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Project) UnsetType() {
	o.Type.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetUpdatedAt() *int64 {
	if o == nil || utils.IsNil(o.UpdatedAt.Get()) {
		var ret *int64 = nil
		return ret
	}
	return o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Project) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given utils.NullableInt32 and assigns it to the UpdatedAt field.
func (o *Project) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *Project) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *Project) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUserOwnerId returns the UserOwnerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Project) GetUserOwnerId() int32 {
	if o == nil || utils.IsNil(o.UserOwnerId.Get()) {
		var ret *int32 = nil
		return *ret
	}
	return *o.UserOwnerId.Get()
}

// GetUserOwnerIdOk returns a tuple with the UserOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetUserOwnerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserOwnerId.Get(), o.UserOwnerId.IsSet()
}

// HasUserOwnerId returns a boolean if a field has been set.
func (o *Project) HasUserOwnerId() bool {
	if o != nil && o.UserOwnerId.IsSet() {
		return true
	}

	return false
}

// SetUserOwnerId gets a reference to the given utils.NullableInt32 and assigns it to the UserOwnerId field.
func (o *Project) SetUserOwnerId(v int32) {
	o.UserOwnerId.Set(&v)
}
// SetUserOwnerIdNil sets the value for UserOwnerId to be an explicit nil
func (o *Project) SetUserOwnerIdNil() {
	o.UserOwnerId.Set(nil)
}

// UnsetUserOwnerId ensures that no value is present for UserOwnerId, not even an explicit nil
func (o *Project) UnsetUserOwnerId() {
	o.UserOwnerId.Unset()
}

// GetProgramIds returns the ProgramIds field value if set, zero value otherwise.
func (o *Project) GetProgramIds() []uint64 {
	if o == nil || utils.IsNil(o.ProgramIds) {
		var ret []uint64
		return ret
	}
	return o.ProgramIds
}

// GetProgramIdsOk returns a tuple with the ProgramIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetProgramIdsOk() ([]uint64, bool) {
	if o == nil || utils.IsNil(o.ProgramIds) {
		return nil, false
	}
	return o.ProgramIds, true
}

// HasProgramIds returns a boolean if a field has been set.
func (o *Project) HasProgramIds() bool {
	if o != nil && !utils.IsNil(o.ProgramIds) {
		return true
	}

	return false
}

// SetProgramIds gets a reference to the given []uint64 and assigns it to the ProgramIds field.
func (o *Project) SetProgramIds(v []uint64) {
	o.ProgramIds = v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Project) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Archived.IsSet() {
		toSerialize["Archived"] = o.Archived.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["CreatedAt"] = o.CreatedAt.Get()
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["DisplayName"] = o.DisplayName.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["EndDate"] = o.EndDate.Get()
	}
	if o.Icon.IsSet() {
		toSerialize["Icon"] = o.Icon.Get()
	}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if o.IsDemo.IsSet() {
		toSerialize["IsDemo"] = o.IsDemo.Get()
	}
	if o.MaxMember.IsSet() {
		toSerialize["MaxMember"] = o.MaxMember.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.StartDate.IsSet() {
		toSerialize["StartDate"] = o.StartDate.Get()
	}
	if o.Status.IsSet() {
		toSerialize["Status"] = o.Status.Get()
	}
	if o.TeamId.IsSet() {
		toSerialize["TeamId"] = o.TeamId.Get()
	}
	if o.TeamOwnerId.IsSet() {
		toSerialize["TeamOwnerId"] = o.TeamOwnerId.Get()
	}
	if o.Type.IsSet() {
		toSerialize["Type"] = o.Type.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["UpdatedAt"] = o.UpdatedAt.Get()
	}
	if o.UserOwnerId.IsSet() {
		toSerialize["UserOwnerId"] = o.UserOwnerId.Get()
	}
	if !utils.IsNil(o.ProgramIds) {
		toSerialize["ProgramIds"] = o.ProgramIds
	}
	return toSerialize, nil
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

