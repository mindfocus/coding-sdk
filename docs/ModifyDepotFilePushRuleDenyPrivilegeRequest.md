# ModifyDepotFilePushRuleDenyPrivilegeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径 | 
**FilePushRuleId** | **int64** | 文件推送规则 ID | 
**IsDeny** | **bool** | 拒绝推送 | 
**IsRole** | **bool** | 特权者是角色（IsUser和IsRole有且只能有一个为true） | 
**IsUser** | **bool** | 特权者是用户（IsUser和IsRole有且只能有一个为true） | 
**RoleId** | Pointer to **int64** | 角色 ID | [optional] 
**UserGlobalKey** | **string** | 用户全局 key | 

## Methods

### NewModifyDepotFilePushRuleDenyPrivilegeRequest

`func NewModifyDepotFilePushRuleDenyPrivilegeRequest(depotPath string, filePushRuleId int64, isDeny bool, isRole bool, isUser bool, userGlobalKey string, ) *ModifyDepotFilePushRuleDenyPrivilegeRequest`

NewModifyDepotFilePushRuleDenyPrivilegeRequest instantiates a new ModifyDepotFilePushRuleDenyPrivilegeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyDepotFilePushRuleDenyPrivilegeRequestWithDefaults

`func NewModifyDepotFilePushRuleDenyPrivilegeRequestWithDefaults() *ModifyDepotFilePushRuleDenyPrivilegeRequest`

NewModifyDepotFilePushRuleDenyPrivilegeRequestWithDefaults instantiates a new ModifyDepotFilePushRuleDenyPrivilegeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetFilePushRuleId

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) GetFilePushRuleId() int64`

GetFilePushRuleId returns the FilePushRuleId field if non-nil, zero value otherwise.

### GetFilePushRuleIdOk

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) GetFilePushRuleIdOk() (*int64, bool)`

GetFilePushRuleIdOk returns a tuple with the FilePushRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePushRuleId

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) SetFilePushRuleId(v int64)`

SetFilePushRuleId sets FilePushRuleId field to given value.


### GetIsDeny

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) GetIsDeny() bool`

GetIsDeny returns the IsDeny field if non-nil, zero value otherwise.

### GetIsDenyOk

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) GetIsDenyOk() (*bool, bool)`

GetIsDenyOk returns a tuple with the IsDeny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeny

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) SetIsDeny(v bool)`

SetIsDeny sets IsDeny field to given value.


### GetIsRole

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) GetIsRole() bool`

GetIsRole returns the IsRole field if non-nil, zero value otherwise.

### GetIsRoleOk

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) GetIsRoleOk() (*bool, bool)`

GetIsRoleOk returns a tuple with the IsRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRole

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) SetIsRole(v bool)`

SetIsRole sets IsRole field to given value.


### GetIsUser

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) GetIsUser() bool`

GetIsUser returns the IsUser field if non-nil, zero value otherwise.

### GetIsUserOk

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) GetIsUserOk() (*bool, bool)`

GetIsUserOk returns a tuple with the IsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUser

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) SetIsUser(v bool)`

SetIsUser sets IsUser field to given value.


### GetRoleId

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) GetRoleId() int64`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) GetRoleIdOk() (*int64, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) SetRoleId(v int64)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetUserGlobalKey

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) GetUserGlobalKey() string`

GetUserGlobalKey returns the UserGlobalKey field if non-nil, zero value otherwise.

### GetUserGlobalKeyOk

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) GetUserGlobalKeyOk() (*string, bool)`

GetUserGlobalKeyOk returns a tuple with the UserGlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGlobalKey

`func (o *ModifyDepotFilePushRuleDenyPrivilegeRequest) SetUserGlobalKey(v string)`

SetUserGlobalKey sets UserGlobalKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


