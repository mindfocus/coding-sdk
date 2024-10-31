# DeleteDepotFilePushRuleDenyPrivilegeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径 | 
**FilePushRuleId** | **int64** | 文件推送规则 ID | 
**IsRole** | **bool** | 特权者是角色（IsUser和IsRole有且只能有一个为true） | 
**IsUser** | **bool** | 特权者是用户（IsUser和IsRole有且只能有一个为true） | 
**RoleId** | Pointer to **int64** | 角色 ID | [optional] 
**UserGlobalKey** | **string** | 用户全局 key | 

## Methods

### NewDeleteDepotFilePushRuleDenyPrivilegeRequest

`func NewDeleteDepotFilePushRuleDenyPrivilegeRequest(depotPath string, filePushRuleId int64, isRole bool, isUser bool, userGlobalKey string, ) *DeleteDepotFilePushRuleDenyPrivilegeRequest`

NewDeleteDepotFilePushRuleDenyPrivilegeRequest instantiates a new DeleteDepotFilePushRuleDenyPrivilegeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDepotFilePushRuleDenyPrivilegeRequestWithDefaults

`func NewDeleteDepotFilePushRuleDenyPrivilegeRequestWithDefaults() *DeleteDepotFilePushRuleDenyPrivilegeRequest`

NewDeleteDepotFilePushRuleDenyPrivilegeRequestWithDefaults instantiates a new DeleteDepotFilePushRuleDenyPrivilegeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetFilePushRuleId

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) GetFilePushRuleId() int64`

GetFilePushRuleId returns the FilePushRuleId field if non-nil, zero value otherwise.

### GetFilePushRuleIdOk

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) GetFilePushRuleIdOk() (*int64, bool)`

GetFilePushRuleIdOk returns a tuple with the FilePushRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePushRuleId

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) SetFilePushRuleId(v int64)`

SetFilePushRuleId sets FilePushRuleId field to given value.


### GetIsRole

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) GetIsRole() bool`

GetIsRole returns the IsRole field if non-nil, zero value otherwise.

### GetIsRoleOk

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) GetIsRoleOk() (*bool, bool)`

GetIsRoleOk returns a tuple with the IsRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRole

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) SetIsRole(v bool)`

SetIsRole sets IsRole field to given value.


### GetIsUser

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) GetIsUser() bool`

GetIsUser returns the IsUser field if non-nil, zero value otherwise.

### GetIsUserOk

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) GetIsUserOk() (*bool, bool)`

GetIsUserOk returns a tuple with the IsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUser

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) SetIsUser(v bool)`

SetIsUser sets IsUser field to given value.


### GetRoleId

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) GetRoleId() int64`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) GetRoleIdOk() (*int64, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) SetRoleId(v int64)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetUserGlobalKey

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) GetUserGlobalKey() string`

GetUserGlobalKey returns the UserGlobalKey field if non-nil, zero value otherwise.

### GetUserGlobalKeyOk

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) GetUserGlobalKeyOk() (*string, bool)`

GetUserGlobalKeyOk returns a tuple with the UserGlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGlobalKey

`func (o *DeleteDepotFilePushRuleDenyPrivilegeRequest) SetUserGlobalKey(v string)`

SetUserGlobalKey sets UserGlobalKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


