/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ServiceHook type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ServiceHook{}

// ServiceHook ServiceHook 对象
type ServiceHook struct {
	// 发送类型
	Action *string `json:"Action,omitempty"`
	// 发送类型描述
	ActionLabel *string `json:"ActionLabel,omitempty"`
	// 发送行为属性
	ActionProperties *string `json:"ActionProperties,omitempty"`
	// 创建时间
	CreatedAt *int64 `json:"CreatedAt,omitempty"`
	// 创建者编号
	CreatorBy *int64 `json:"CreatorBy,omitempty"`
	CreatorByUser *ServiceHookUser `json:"CreatorByUser,omitempty"`
	// 事件开关
	Enabled *bool `json:"Enabled,omitempty"`
	// 事件列表
	Event []string `json:"Event,omitempty"`
	// 事件描述列表
	EventLabel []string `json:"EventLabel,omitempty"`
	// 过滤器属性
	FilterProperties *string `json:"FilterProperties,omitempty"`
	// Service Hook 编号
	Id *string `json:"Id,omitempty"`
	// 最近发送成功时间
	LastSucceedAt *int64 `json:"LastSucceedAt,omitempty"`
	// Service Hook 名称
	Name *string `json:"Name,omitempty"`
	// 服务类型
	Service *string `json:"Service,omitempty"`
	// 服务名
	ServiceName *string `json:"ServiceName,omitempty"`
	// 发送状态
	Status *int64 `json:"Status,omitempty"`
	// 目标数据编号
	TargetId *int64 `json:"TargetId,omitempty"`
	// 目标数据类型：Project、Team,Space,Program
	TargetType *string `json:"TargetType,omitempty"`
	// 更新时间
	UpdatedAt *int64 `json:"UpdatedAt,omitempty"`
	// 更新者编号
	UpdatedBy *int64 `json:"UpdatedBy,omitempty"`
	UpdatedByUser *ServiceHookUser `json:"UpdatedByUser,omitempty"`
	// 版本
	Version *int64 `json:"Version,omitempty"`
}

// NewServiceHook instantiates a new ServiceHook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceHook() *ServiceHook {
	this := ServiceHook{}
	var action string = ""
	this.Action = &action
	var actionLabel string = ""
	this.ActionLabel = &actionLabel
	var actionProperties string = ""
	this.ActionProperties = &actionProperties
	var enabled bool = false
	this.Enabled = &enabled
	var filterProperties string = ""
	this.FilterProperties = &filterProperties
	var id string = ""
	this.Id = &id
	var name string = ""
	this.Name = &name
	var service string = ""
	this.Service = &service
	var serviceName string = ""
	this.ServiceName = &serviceName
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// NewServiceHookWithDefaults instantiates a new ServiceHook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceHookWithDefaults() *ServiceHook {
	this := ServiceHook{}
	var action string = ""
	this.Action = &action
	var actionLabel string = ""
	this.ActionLabel = &actionLabel
	var actionProperties string = ""
	this.ActionProperties = &actionProperties
	var enabled bool = false
	this.Enabled = &enabled
	var filterProperties string = ""
	this.FilterProperties = &filterProperties
	var id string = ""
	this.Id = &id
	var name string = ""
	this.Name = &name
	var service string = ""
	this.Service = &service
	var serviceName string = ""
	this.ServiceName = &serviceName
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ServiceHook) GetAction() string {
	if o == nil || utils.IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetActionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ServiceHook) HasAction() bool {
	if o != nil && !utils.IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ServiceHook) SetAction(v string) {
	o.Action = &v
}

// GetActionLabel returns the ActionLabel field value if set, zero value otherwise.
func (o *ServiceHook) GetActionLabel() string {
	if o == nil || utils.IsNil(o.ActionLabel) {
		var ret string
		return ret
	}
	return *o.ActionLabel
}

// GetActionLabelOk returns a tuple with the ActionLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetActionLabelOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ActionLabel) {
		return nil, false
	}
	return o.ActionLabel, true
}

// HasActionLabel returns a boolean if a field has been set.
func (o *ServiceHook) HasActionLabel() bool {
	if o != nil && !utils.IsNil(o.ActionLabel) {
		return true
	}

	return false
}

// SetActionLabel gets a reference to the given string and assigns it to the ActionLabel field.
func (o *ServiceHook) SetActionLabel(v string) {
	o.ActionLabel = &v
}

// GetActionProperties returns the ActionProperties field value if set, zero value otherwise.
func (o *ServiceHook) GetActionProperties() string {
	if o == nil || utils.IsNil(o.ActionProperties) {
		var ret string
		return ret
	}
	return *o.ActionProperties
}

// GetActionPropertiesOk returns a tuple with the ActionProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetActionPropertiesOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ActionProperties) {
		return nil, false
	}
	return o.ActionProperties, true
}

// HasActionProperties returns a boolean if a field has been set.
func (o *ServiceHook) HasActionProperties() bool {
	if o != nil && !utils.IsNil(o.ActionProperties) {
		return true
	}

	return false
}

// SetActionProperties gets a reference to the given string and assigns it to the ActionProperties field.
func (o *ServiceHook) SetActionProperties(v string) {
	o.ActionProperties = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServiceHook) GetCreatedAt() int64 {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetCreatedAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceHook) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *ServiceHook) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreatorBy returns the CreatorBy field value if set, zero value otherwise.
func (o *ServiceHook) GetCreatorBy() int64 {
	if o == nil || utils.IsNil(o.CreatorBy) {
		var ret int64
		return ret
	}
	return *o.CreatorBy
}

// GetCreatorByOk returns a tuple with the CreatorBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetCreatorByOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.CreatorBy) {
		return nil, false
	}
	return o.CreatorBy, true
}

// HasCreatorBy returns a boolean if a field has been set.
func (o *ServiceHook) HasCreatorBy() bool {
	if o != nil && !utils.IsNil(o.CreatorBy) {
		return true
	}

	return false
}

// SetCreatorBy gets a reference to the given int64 and assigns it to the CreatorBy field.
func (o *ServiceHook) SetCreatorBy(v int64) {
	o.CreatorBy = &v
}

// GetCreatorByUser returns the CreatorByUser field value if set, zero value otherwise.
func (o *ServiceHook) GetCreatorByUser() ServiceHookUser {
	if o == nil || utils.IsNil(o.CreatorByUser) {
		var ret ServiceHookUser
		return ret
	}
	return *o.CreatorByUser
}

// GetCreatorByUserOk returns a tuple with the CreatorByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetCreatorByUserOk() (*ServiceHookUser, bool) {
	if o == nil || utils.IsNil(o.CreatorByUser) {
		return nil, false
	}
	return o.CreatorByUser, true
}

// HasCreatorByUser returns a boolean if a field has been set.
func (o *ServiceHook) HasCreatorByUser() bool {
	if o != nil && !utils.IsNil(o.CreatorByUser) {
		return true
	}

	return false
}

// SetCreatorByUser gets a reference to the given ServiceHookUser and assigns it to the CreatorByUser field.
func (o *ServiceHook) SetCreatorByUser(v ServiceHookUser) {
	o.CreatorByUser = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ServiceHook) GetEnabled() bool {
	if o == nil || utils.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ServiceHook) HasEnabled() bool {
	if o != nil && !utils.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ServiceHook) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *ServiceHook) GetEvent() []string {
	if o == nil || utils.IsNil(o.Event) {
		var ret []string
		return ret
	}
	return o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetEventOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *ServiceHook) HasEvent() bool {
	if o != nil && !utils.IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given []string and assigns it to the Event field.
func (o *ServiceHook) SetEvent(v []string) {
	o.Event = v
}

// GetEventLabel returns the EventLabel field value if set, zero value otherwise.
func (o *ServiceHook) GetEventLabel() []string {
	if o == nil || utils.IsNil(o.EventLabel) {
		var ret []string
		return ret
	}
	return o.EventLabel
}

// GetEventLabelOk returns a tuple with the EventLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetEventLabelOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.EventLabel) {
		return nil, false
	}
	return o.EventLabel, true
}

// HasEventLabel returns a boolean if a field has been set.
func (o *ServiceHook) HasEventLabel() bool {
	if o != nil && !utils.IsNil(o.EventLabel) {
		return true
	}

	return false
}

// SetEventLabel gets a reference to the given []string and assigns it to the EventLabel field.
func (o *ServiceHook) SetEventLabel(v []string) {
	o.EventLabel = v
}

// GetFilterProperties returns the FilterProperties field value if set, zero value otherwise.
func (o *ServiceHook) GetFilterProperties() string {
	if o == nil || utils.IsNil(o.FilterProperties) {
		var ret string
		return ret
	}
	return *o.FilterProperties
}

// GetFilterPropertiesOk returns a tuple with the FilterProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetFilterPropertiesOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FilterProperties) {
		return nil, false
	}
	return o.FilterProperties, true
}

// HasFilterProperties returns a boolean if a field has been set.
func (o *ServiceHook) HasFilterProperties() bool {
	if o != nil && !utils.IsNil(o.FilterProperties) {
		return true
	}

	return false
}

// SetFilterProperties gets a reference to the given string and assigns it to the FilterProperties field.
func (o *ServiceHook) SetFilterProperties(v string) {
	o.FilterProperties = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceHook) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceHook) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceHook) SetId(v string) {
	o.Id = &v
}

// GetLastSucceedAt returns the LastSucceedAt field value if set, zero value otherwise.
func (o *ServiceHook) GetLastSucceedAt() int64 {
	if o == nil || utils.IsNil(o.LastSucceedAt) {
		var ret int64
		return ret
	}
	return *o.LastSucceedAt
}

// GetLastSucceedAtOk returns a tuple with the LastSucceedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetLastSucceedAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.LastSucceedAt) {
		return nil, false
	}
	return o.LastSucceedAt, true
}

// HasLastSucceedAt returns a boolean if a field has been set.
func (o *ServiceHook) HasLastSucceedAt() bool {
	if o != nil && !utils.IsNil(o.LastSucceedAt) {
		return true
	}

	return false
}

// SetLastSucceedAt gets a reference to the given int64 and assigns it to the LastSucceedAt field.
func (o *ServiceHook) SetLastSucceedAt(v int64) {
	o.LastSucceedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceHook) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceHook) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceHook) SetName(v string) {
	o.Name = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ServiceHook) GetService() string {
	if o == nil || utils.IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetServiceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ServiceHook) HasService() bool {
	if o != nil && !utils.IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ServiceHook) SetService(v string) {
	o.Service = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *ServiceHook) GetServiceName() string {
	if o == nil || utils.IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetServiceNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *ServiceHook) HasServiceName() bool {
	if o != nil && !utils.IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *ServiceHook) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServiceHook) GetStatus() int64 {
	if o == nil || utils.IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetStatusOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServiceHook) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *ServiceHook) SetStatus(v int64) {
	o.Status = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *ServiceHook) GetTargetId() int64 {
	if o == nil || utils.IsNil(o.TargetId) {
		var ret int64
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetTargetIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *ServiceHook) HasTargetId() bool {
	if o != nil && !utils.IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given int64 and assigns it to the TargetId field.
func (o *ServiceHook) SetTargetId(v int64) {
	o.TargetId = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *ServiceHook) GetTargetType() string {
	if o == nil || utils.IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetTargetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *ServiceHook) HasTargetType() bool {
	if o != nil && !utils.IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *ServiceHook) SetTargetType(v string) {
	o.TargetType = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ServiceHook) GetUpdatedAt() int64 {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ServiceHook) HasUpdatedAt() bool {
	if o != nil && !utils.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *ServiceHook) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ServiceHook) GetUpdatedBy() int64 {
	if o == nil || utils.IsNil(o.UpdatedBy) {
		var ret int64
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetUpdatedByOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ServiceHook) HasUpdatedBy() bool {
	if o != nil && !utils.IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given int64 and assigns it to the UpdatedBy field.
func (o *ServiceHook) SetUpdatedBy(v int64) {
	o.UpdatedBy = &v
}

// GetUpdatedByUser returns the UpdatedByUser field value if set, zero value otherwise.
func (o *ServiceHook) GetUpdatedByUser() ServiceHookUser {
	if o == nil || utils.IsNil(o.UpdatedByUser) {
		var ret ServiceHookUser
		return ret
	}
	return *o.UpdatedByUser
}

// GetUpdatedByUserOk returns a tuple with the UpdatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetUpdatedByUserOk() (*ServiceHookUser, bool) {
	if o == nil || utils.IsNil(o.UpdatedByUser) {
		return nil, false
	}
	return o.UpdatedByUser, true
}

// HasUpdatedByUser returns a boolean if a field has been set.
func (o *ServiceHook) HasUpdatedByUser() bool {
	if o != nil && !utils.IsNil(o.UpdatedByUser) {
		return true
	}

	return false
}

// SetUpdatedByUser gets a reference to the given ServiceHookUser and assigns it to the UpdatedByUser field.
func (o *ServiceHook) SetUpdatedByUser(v ServiceHookUser) {
	o.UpdatedByUser = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ServiceHook) GetVersion() int64 {
	if o == nil || utils.IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetVersionOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ServiceHook) HasVersion() bool {
	if o != nil && !utils.IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *ServiceHook) SetVersion(v int64) {
	o.Version = &v
}

func (o ServiceHook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceHook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Action) {
		toSerialize["Action"] = o.Action
	}
	if !utils.IsNil(o.ActionLabel) {
		toSerialize["ActionLabel"] = o.ActionLabel
	}
	if !utils.IsNil(o.ActionProperties) {
		toSerialize["ActionProperties"] = o.ActionProperties
	}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if !utils.IsNil(o.CreatorBy) {
		toSerialize["CreatorBy"] = o.CreatorBy
	}
	if !utils.IsNil(o.CreatorByUser) {
		toSerialize["CreatorByUser"] = o.CreatorByUser
	}
	if !utils.IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !utils.IsNil(o.Event) {
		toSerialize["Event"] = o.Event
	}
	if !utils.IsNil(o.EventLabel) {
		toSerialize["EventLabel"] = o.EventLabel
	}
	if !utils.IsNil(o.FilterProperties) {
		toSerialize["FilterProperties"] = o.FilterProperties
	}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !utils.IsNil(o.LastSucceedAt) {
		toSerialize["LastSucceedAt"] = o.LastSucceedAt
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !utils.IsNil(o.Service) {
		toSerialize["Service"] = o.Service
	}
	if !utils.IsNil(o.ServiceName) {
		toSerialize["ServiceName"] = o.ServiceName
	}
	if !utils.IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !utils.IsNil(o.TargetId) {
		toSerialize["TargetId"] = o.TargetId
	}
	if !utils.IsNil(o.TargetType) {
		toSerialize["TargetType"] = o.TargetType
	}
	if !utils.IsNil(o.UpdatedAt) {
		toSerialize["UpdatedAt"] = o.UpdatedAt
	}
	if !utils.IsNil(o.UpdatedBy) {
		toSerialize["UpdatedBy"] = o.UpdatedBy
	}
	if !utils.IsNil(o.UpdatedByUser) {
		toSerialize["UpdatedByUser"] = o.UpdatedByUser
	}
	if !utils.IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	return toSerialize, nil
}

type NullableServiceHook struct {
	value *ServiceHook
	isSet bool
}

func (v NullableServiceHook) Get() *ServiceHook {
	return v.value
}

func (v *NullableServiceHook) Set(val *ServiceHook) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceHook) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceHook(val *ServiceHook) *NullableServiceHook {
	return &NullableServiceHook{value: val, isSet: true}
}

func (v NullableServiceHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

