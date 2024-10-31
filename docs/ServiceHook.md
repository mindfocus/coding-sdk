# ServiceHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | 发送类型 | [optional] [default to ""]
**ActionLabel** | Pointer to **string** | 发送类型描述 | [optional] [default to ""]
**ActionProperties** | Pointer to **string** | 发送行为属性 | [optional] [default to ""]
**CreatedAt** | Pointer to **int64** | 创建时间 | [optional] 
**CreatorBy** | Pointer to **int64** | 创建者编号 | [optional] 
**CreatorByUser** | Pointer to [**ServiceHookUser**](ServiceHookUser.md) |  | [optional] 
**Enabled** | Pointer to **bool** | 事件开关 | [optional] [default to false]
**Event** | Pointer to **[]string** | 事件列表 | [optional] 
**EventLabel** | Pointer to **[]string** | 事件描述列表 | [optional] 
**FilterProperties** | Pointer to **string** | 过滤器属性 | [optional] [default to ""]
**Id** | Pointer to **string** | Service Hook 编号 | [optional] [default to ""]
**LastSucceedAt** | Pointer to **int64** | 最近发送成功时间 | [optional] 
**Name** | Pointer to **string** | Service Hook 名称 | [optional] [default to ""]
**Service** | Pointer to **string** | 服务类型 | [optional] [default to ""]
**ServiceName** | Pointer to **string** | 服务名 | [optional] [default to ""]
**Status** | Pointer to **int64** | 发送状态 | [optional] 
**TargetId** | Pointer to **int64** | 目标数据编号 | [optional] 
**TargetType** | Pointer to **string** | 目标数据类型：Project、Team,Space,Program | [optional] [default to ""]
**UpdatedAt** | Pointer to **int64** | 更新时间 | [optional] 
**UpdatedBy** | Pointer to **int64** | 更新者编号 | [optional] 
**UpdatedByUser** | Pointer to [**ServiceHookUser**](ServiceHookUser.md) |  | [optional] 
**Version** | Pointer to **int64** | 版本 | [optional] 

## Methods

### NewServiceHook

`func NewServiceHook() *ServiceHook`

NewServiceHook instantiates a new ServiceHook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceHookWithDefaults

`func NewServiceHookWithDefaults() *ServiceHook`

NewServiceHookWithDefaults instantiates a new ServiceHook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ServiceHook) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ServiceHook) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ServiceHook) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ServiceHook) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActionLabel

`func (o *ServiceHook) GetActionLabel() string`

GetActionLabel returns the ActionLabel field if non-nil, zero value otherwise.

### GetActionLabelOk

`func (o *ServiceHook) GetActionLabelOk() (*string, bool)`

GetActionLabelOk returns a tuple with the ActionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionLabel

`func (o *ServiceHook) SetActionLabel(v string)`

SetActionLabel sets ActionLabel field to given value.

### HasActionLabel

`func (o *ServiceHook) HasActionLabel() bool`

HasActionLabel returns a boolean if a field has been set.

### GetActionProperties

`func (o *ServiceHook) GetActionProperties() string`

GetActionProperties returns the ActionProperties field if non-nil, zero value otherwise.

### GetActionPropertiesOk

`func (o *ServiceHook) GetActionPropertiesOk() (*string, bool)`

GetActionPropertiesOk returns a tuple with the ActionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionProperties

`func (o *ServiceHook) SetActionProperties(v string)`

SetActionProperties sets ActionProperties field to given value.

### HasActionProperties

`func (o *ServiceHook) HasActionProperties() bool`

HasActionProperties returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ServiceHook) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceHook) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceHook) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceHook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatorBy

`func (o *ServiceHook) GetCreatorBy() int64`

GetCreatorBy returns the CreatorBy field if non-nil, zero value otherwise.

### GetCreatorByOk

`func (o *ServiceHook) GetCreatorByOk() (*int64, bool)`

GetCreatorByOk returns a tuple with the CreatorBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorBy

`func (o *ServiceHook) SetCreatorBy(v int64)`

SetCreatorBy sets CreatorBy field to given value.

### HasCreatorBy

`func (o *ServiceHook) HasCreatorBy() bool`

HasCreatorBy returns a boolean if a field has been set.

### GetCreatorByUser

`func (o *ServiceHook) GetCreatorByUser() ServiceHookUser`

GetCreatorByUser returns the CreatorByUser field if non-nil, zero value otherwise.

### GetCreatorByUserOk

`func (o *ServiceHook) GetCreatorByUserOk() (*ServiceHookUser, bool)`

GetCreatorByUserOk returns a tuple with the CreatorByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorByUser

`func (o *ServiceHook) SetCreatorByUser(v ServiceHookUser)`

SetCreatorByUser sets CreatorByUser field to given value.

### HasCreatorByUser

`func (o *ServiceHook) HasCreatorByUser() bool`

HasCreatorByUser returns a boolean if a field has been set.

### GetEnabled

`func (o *ServiceHook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceHook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceHook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServiceHook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEvent

`func (o *ServiceHook) GetEvent() []string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ServiceHook) GetEventOk() (*[]string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ServiceHook) SetEvent(v []string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *ServiceHook) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetEventLabel

`func (o *ServiceHook) GetEventLabel() []string`

GetEventLabel returns the EventLabel field if non-nil, zero value otherwise.

### GetEventLabelOk

`func (o *ServiceHook) GetEventLabelOk() (*[]string, bool)`

GetEventLabelOk returns a tuple with the EventLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLabel

`func (o *ServiceHook) SetEventLabel(v []string)`

SetEventLabel sets EventLabel field to given value.

### HasEventLabel

`func (o *ServiceHook) HasEventLabel() bool`

HasEventLabel returns a boolean if a field has been set.

### GetFilterProperties

`func (o *ServiceHook) GetFilterProperties() string`

GetFilterProperties returns the FilterProperties field if non-nil, zero value otherwise.

### GetFilterPropertiesOk

`func (o *ServiceHook) GetFilterPropertiesOk() (*string, bool)`

GetFilterPropertiesOk returns a tuple with the FilterProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterProperties

`func (o *ServiceHook) SetFilterProperties(v string)`

SetFilterProperties sets FilterProperties field to given value.

### HasFilterProperties

`func (o *ServiceHook) HasFilterProperties() bool`

HasFilterProperties returns a boolean if a field has been set.

### GetId

`func (o *ServiceHook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceHook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceHook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceHook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastSucceedAt

`func (o *ServiceHook) GetLastSucceedAt() int64`

GetLastSucceedAt returns the LastSucceedAt field if non-nil, zero value otherwise.

### GetLastSucceedAtOk

`func (o *ServiceHook) GetLastSucceedAtOk() (*int64, bool)`

GetLastSucceedAtOk returns a tuple with the LastSucceedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSucceedAt

`func (o *ServiceHook) SetLastSucceedAt(v int64)`

SetLastSucceedAt sets LastSucceedAt field to given value.

### HasLastSucceedAt

`func (o *ServiceHook) HasLastSucceedAt() bool`

HasLastSucceedAt returns a boolean if a field has been set.

### GetName

`func (o *ServiceHook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceHook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceHook) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceHook) HasName() bool`

HasName returns a boolean if a field has been set.

### GetService

`func (o *ServiceHook) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceHook) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceHook) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ServiceHook) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceName

`func (o *ServiceHook) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceHook) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceHook) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServiceHook) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceHook) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceHook) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceHook) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceHook) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetId

`func (o *ServiceHook) GetTargetId() int64`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ServiceHook) GetTargetIdOk() (*int64, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ServiceHook) SetTargetId(v int64)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ServiceHook) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetTargetType

`func (o *ServiceHook) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ServiceHook) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ServiceHook) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ServiceHook) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ServiceHook) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ServiceHook) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ServiceHook) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ServiceHook) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ServiceHook) GetUpdatedBy() int64`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ServiceHook) GetUpdatedByOk() (*int64, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ServiceHook) SetUpdatedBy(v int64)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ServiceHook) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedByUser

`func (o *ServiceHook) GetUpdatedByUser() ServiceHookUser`

GetUpdatedByUser returns the UpdatedByUser field if non-nil, zero value otherwise.

### GetUpdatedByUserOk

`func (o *ServiceHook) GetUpdatedByUserOk() (*ServiceHookUser, bool)`

GetUpdatedByUserOk returns a tuple with the UpdatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUser

`func (o *ServiceHook) SetUpdatedByUser(v ServiceHookUser)`

SetUpdatedByUser sets UpdatedByUser field to given value.

### HasUpdatedByUser

`func (o *ServiceHook) HasUpdatedByUser() bool`

HasUpdatedByUser returns a boolean if a field has been set.

### GetVersion

`func (o *ServiceHook) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceHook) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceHook) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServiceHook) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


