/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DepotInfo type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DepotInfo{}

// DepotInfo 仓库信息
type DepotInfo struct {
	// 创建时间
	CreatedAt *int64 `json:"CreatedAt,omitempty"`
	// 默认分支
	DefaultBranch utils.NullableString `json:"DefaultBranch,omitempty"`
	// 仓库描述
	Description *string `json:"Description,omitempty"`
	// 项目组Id
	GroupId utils.NullableInt64 `json:"GroupId,omitempty"`
	// 项目名称
	GroupName utils.NullableString `json:"GroupName,omitempty"`
	// 项目类型
	GroupType utils.NullableString `json:"GroupType,omitempty"`
	// 仓库的https地址
	HttpsUrl *string `json:"HttpsUrl,omitempty"`
	// 仓库id
	Id *int64 `json:"Id,omitempty"`
	// 最终push时间
	LastPushAt *int64 `json:"LastPushAt,omitempty"`
	// 仓库名称
	Name *string `json:"Name,omitempty"`
	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty"`
	// 仓库类型,具体值为git或者svn
	RepoType *string `json:"RepoType,omitempty"`
	// 仓库的ssh地址
	SshUrl *string `json:"SshUrl,omitempty"`
	// 仓库类型,具体值为git或者svn
	VcsType *string `json:"VcsType,omitempty"`
	// 仓库webURL
	WebUrl *string `json:"WebUrl,omitempty"`
}

// NewDepotInfo instantiates a new DepotInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepotInfo() *DepotInfo {
	this := DepotInfo{}
	var defaultBranch string = ""
	this.DefaultBranch = *utils.NewNullableString(&defaultBranch)
	var description string = ""
	this.Description = &description
	var groupName string = ""
	this.GroupName = *utils.NewNullableString(&groupName)
	var groupType string = ""
	this.GroupType = *utils.NewNullableString(&groupType)
	var httpsUrl string = ""
	this.HttpsUrl = &httpsUrl
	var name string = ""
	this.Name = &name
	var projectName string = ""
	this.ProjectName = &projectName
	var repoType string = ""
	this.RepoType = &repoType
	var sshUrl string = ""
	this.SshUrl = &sshUrl
	var vcsType string = ""
	this.VcsType = &vcsType
	var webUrl string = ""
	this.WebUrl = &webUrl
	return &this
}

// NewDepotInfoWithDefaults instantiates a new DepotInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepotInfoWithDefaults() *DepotInfo {
	this := DepotInfo{}
	var defaultBranch string = ""
	this.DefaultBranch = *utils.NewNullableString(&defaultBranch)
	var description string = ""
	this.Description = &description
	var groupName string = ""
	this.GroupName = *utils.NewNullableString(&groupName)
	var groupType string = ""
	this.GroupType = *utils.NewNullableString(&groupType)
	var httpsUrl string = ""
	this.HttpsUrl = &httpsUrl
	var name string = ""
	this.Name = &name
	var projectName string = ""
	this.ProjectName = &projectName
	var repoType string = ""
	this.RepoType = &repoType
	var sshUrl string = ""
	this.SshUrl = &sshUrl
	var vcsType string = ""
	this.VcsType = &vcsType
	var webUrl string = ""
	this.WebUrl = &webUrl
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DepotInfo) GetCreatedAt() int64 {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetCreatedAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DepotInfo) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *DepotInfo) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetDefaultBranch returns the DefaultBranch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepotInfo) GetDefaultBranch() string {
	if o == nil || utils.IsNil(o.DefaultBranch.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultBranch.Get()
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepotInfo) GetDefaultBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultBranch.Get(), o.DefaultBranch.IsSet()
}

// HasDefaultBranch returns a boolean if a field has been set.
func (o *DepotInfo) HasDefaultBranch() bool {
	if o != nil && o.DefaultBranch.IsSet() {
		return true
	}

	return false
}

// SetDefaultBranch gets a reference to the given utils.NullableString and assigns it to the DefaultBranch field.
func (o *DepotInfo) SetDefaultBranch(v string) {
	o.DefaultBranch.Set(&v)
}
// SetDefaultBranchNil sets the value for DefaultBranch to be an explicit nil
func (o *DepotInfo) SetDefaultBranchNil() {
	o.DefaultBranch.Set(nil)
}

// UnsetDefaultBranch ensures that no value is present for DefaultBranch, not even an explicit nil
func (o *DepotInfo) UnsetDefaultBranch() {
	o.DefaultBranch.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DepotInfo) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DepotInfo) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DepotInfo) SetDescription(v string) {
	o.Description = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepotInfo) GetGroupId() int64 {
	if o == nil || utils.IsNil(o.GroupId.Get()) {
		var ret int64
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepotInfo) GetGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *DepotInfo) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given utils.NullableInt64 and assigns it to the GroupId field.
func (o *DepotInfo) SetGroupId(v int64) {
	o.GroupId.Set(&v)
}
// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *DepotInfo) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *DepotInfo) UnsetGroupId() {
	o.GroupId.Unset()
}

// GetGroupName returns the GroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepotInfo) GetGroupName() string {
	if o == nil || utils.IsNil(o.GroupName.Get()) {
		var ret string
		return ret
	}
	return *o.GroupName.Get()
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepotInfo) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupName.Get(), o.GroupName.IsSet()
}

// HasGroupName returns a boolean if a field has been set.
func (o *DepotInfo) HasGroupName() bool {
	if o != nil && o.GroupName.IsSet() {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given utils.NullableString and assigns it to the GroupName field.
func (o *DepotInfo) SetGroupName(v string) {
	o.GroupName.Set(&v)
}
// SetGroupNameNil sets the value for GroupName to be an explicit nil
func (o *DepotInfo) SetGroupNameNil() {
	o.GroupName.Set(nil)
}

// UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
func (o *DepotInfo) UnsetGroupName() {
	o.GroupName.Unset()
}

// GetGroupType returns the GroupType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepotInfo) GetGroupType() string {
	if o == nil || utils.IsNil(o.GroupType.Get()) {
		var ret string
		return ret
	}
	return *o.GroupType.Get()
}

// GetGroupTypeOk returns a tuple with the GroupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepotInfo) GetGroupTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupType.Get(), o.GroupType.IsSet()
}

// HasGroupType returns a boolean if a field has been set.
func (o *DepotInfo) HasGroupType() bool {
	if o != nil && o.GroupType.IsSet() {
		return true
	}

	return false
}

// SetGroupType gets a reference to the given utils.NullableString and assigns it to the GroupType field.
func (o *DepotInfo) SetGroupType(v string) {
	o.GroupType.Set(&v)
}
// SetGroupTypeNil sets the value for GroupType to be an explicit nil
func (o *DepotInfo) SetGroupTypeNil() {
	o.GroupType.Set(nil)
}

// UnsetGroupType ensures that no value is present for GroupType, not even an explicit nil
func (o *DepotInfo) UnsetGroupType() {
	o.GroupType.Unset()
}

// GetHttpsUrl returns the HttpsUrl field value if set, zero value otherwise.
func (o *DepotInfo) GetHttpsUrl() string {
	if o == nil || utils.IsNil(o.HttpsUrl) {
		var ret string
		return ret
	}
	return *o.HttpsUrl
}

// GetHttpsUrlOk returns a tuple with the HttpsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetHttpsUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.HttpsUrl) {
		return nil, false
	}
	return o.HttpsUrl, true
}

// HasHttpsUrl returns a boolean if a field has been set.
func (o *DepotInfo) HasHttpsUrl() bool {
	if o != nil && !utils.IsNil(o.HttpsUrl) {
		return true
	}

	return false
}

// SetHttpsUrl gets a reference to the given string and assigns it to the HttpsUrl field.
func (o *DepotInfo) SetHttpsUrl(v string) {
	o.HttpsUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DepotInfo) GetId() int64 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DepotInfo) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DepotInfo) SetId(v int64) {
	o.Id = &v
}

// GetLastPushAt returns the LastPushAt field value if set, zero value otherwise.
func (o *DepotInfo) GetLastPushAt() int64 {
	if o == nil || utils.IsNil(o.LastPushAt) {
		var ret int64
		return ret
	}
	return *o.LastPushAt
}

// GetLastPushAtOk returns a tuple with the LastPushAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetLastPushAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.LastPushAt) {
		return nil, false
	}
	return o.LastPushAt, true
}

// HasLastPushAt returns a boolean if a field has been set.
func (o *DepotInfo) HasLastPushAt() bool {
	if o != nil && !utils.IsNil(o.LastPushAt) {
		return true
	}

	return false
}

// SetLastPushAt gets a reference to the given int64 and assigns it to the LastPushAt field.
func (o *DepotInfo) SetLastPushAt(v int64) {
	o.LastPushAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DepotInfo) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DepotInfo) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DepotInfo) SetName(v string) {
	o.Name = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DepotInfo) GetProjectId() int64 {
	if o == nil || utils.IsNil(o.ProjectId) {
		var ret int64
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetProjectIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DepotInfo) HasProjectId() bool {
	if o != nil && !utils.IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int64 and assigns it to the ProjectId field.
func (o *DepotInfo) SetProjectId(v int64) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *DepotInfo) GetProjectName() string {
	if o == nil || utils.IsNil(o.ProjectName) {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetProjectNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ProjectName) {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *DepotInfo) HasProjectName() bool {
	if o != nil && !utils.IsNil(o.ProjectName) {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *DepotInfo) SetProjectName(v string) {
	o.ProjectName = &v
}

// GetRepoType returns the RepoType field value if set, zero value otherwise.
func (o *DepotInfo) GetRepoType() string {
	if o == nil || utils.IsNil(o.RepoType) {
		var ret string
		return ret
	}
	return *o.RepoType
}

// GetRepoTypeOk returns a tuple with the RepoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetRepoTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RepoType) {
		return nil, false
	}
	return o.RepoType, true
}

// HasRepoType returns a boolean if a field has been set.
func (o *DepotInfo) HasRepoType() bool {
	if o != nil && !utils.IsNil(o.RepoType) {
		return true
	}

	return false
}

// SetRepoType gets a reference to the given string and assigns it to the RepoType field.
func (o *DepotInfo) SetRepoType(v string) {
	o.RepoType = &v
}

// GetSshUrl returns the SshUrl field value if set, zero value otherwise.
func (o *DepotInfo) GetSshUrl() string {
	if o == nil || utils.IsNil(o.SshUrl) {
		var ret string
		return ret
	}
	return *o.SshUrl
}

// GetSshUrlOk returns a tuple with the SshUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetSshUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SshUrl) {
		return nil, false
	}
	return o.SshUrl, true
}

// HasSshUrl returns a boolean if a field has been set.
func (o *DepotInfo) HasSshUrl() bool {
	if o != nil && !utils.IsNil(o.SshUrl) {
		return true
	}

	return false
}

// SetSshUrl gets a reference to the given string and assigns it to the SshUrl field.
func (o *DepotInfo) SetSshUrl(v string) {
	o.SshUrl = &v
}

// GetVcsType returns the VcsType field value if set, zero value otherwise.
func (o *DepotInfo) GetVcsType() string {
	if o == nil || utils.IsNil(o.VcsType) {
		var ret string
		return ret
	}
	return *o.VcsType
}

// GetVcsTypeOk returns a tuple with the VcsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetVcsTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.VcsType) {
		return nil, false
	}
	return o.VcsType, true
}

// HasVcsType returns a boolean if a field has been set.
func (o *DepotInfo) HasVcsType() bool {
	if o != nil && !utils.IsNil(o.VcsType) {
		return true
	}

	return false
}

// SetVcsType gets a reference to the given string and assigns it to the VcsType field.
func (o *DepotInfo) SetVcsType(v string) {
	o.VcsType = &v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise.
func (o *DepotInfo) GetWebUrl() string {
	if o == nil || utils.IsNil(o.WebUrl) {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetWebUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.WebUrl) {
		return nil, false
	}
	return o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *DepotInfo) HasWebUrl() bool {
	if o != nil && !utils.IsNil(o.WebUrl) {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *DepotInfo) SetWebUrl(v string) {
	o.WebUrl = &v
}

func (o DepotInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepotInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if o.DefaultBranch.IsSet() {
		toSerialize["DefaultBranch"] = o.DefaultBranch.Get()
	}
	if !utils.IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if o.GroupId.IsSet() {
		toSerialize["GroupId"] = o.GroupId.Get()
	}
	if o.GroupName.IsSet() {
		toSerialize["GroupName"] = o.GroupName.Get()
	}
	if o.GroupType.IsSet() {
		toSerialize["GroupType"] = o.GroupType.Get()
	}
	if !utils.IsNil(o.HttpsUrl) {
		toSerialize["HttpsUrl"] = o.HttpsUrl
	}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !utils.IsNil(o.LastPushAt) {
		toSerialize["LastPushAt"] = o.LastPushAt
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !utils.IsNil(o.ProjectId) {
		toSerialize["ProjectId"] = o.ProjectId
	}
	if !utils.IsNil(o.ProjectName) {
		toSerialize["ProjectName"] = o.ProjectName
	}
	if !utils.IsNil(o.RepoType) {
		toSerialize["RepoType"] = o.RepoType
	}
	if !utils.IsNil(o.SshUrl) {
		toSerialize["SshUrl"] = o.SshUrl
	}
	if !utils.IsNil(o.VcsType) {
		toSerialize["VcsType"] = o.VcsType
	}
	if !utils.IsNil(o.WebUrl) {
		toSerialize["WebUrl"] = o.WebUrl
	}
	return toSerialize, nil
}

type NullableDepotInfo struct {
	value *DepotInfo
	isSet bool
}

func (v NullableDepotInfo) Get() *DepotInfo {
	return v.value
}

func (v *NullableDepotInfo) Set(val *DepotInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDepotInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDepotInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepotInfo(val *DepotInfo) *NullableDepotInfo {
	return &NullableDepotInfo{value: val, isSet: true}
}

func (v NullableDepotInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepotInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


