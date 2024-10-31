# GitFilePushRulePrivilege

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDeny** | Pointer to **bool** | 拒绝推送文件 | [optional] [default to false]
**IsRole** | Pointer to **bool** | 特权者是角色 | [optional] [default to false]
**IsUser** | Pointer to **bool** | 特权者是用户 | [optional] [default to false]
**Role** | Pointer to [**GitFilePushRuleRole**](GitFilePushRuleRole.md) |  | [optional] 
**User** | Pointer to [**GitFilePushRuleUser**](GitFilePushRuleUser.md) |  | [optional] 

## Methods

### NewGitFilePushRulePrivilege

`func NewGitFilePushRulePrivilege() *GitFilePushRulePrivilege`

NewGitFilePushRulePrivilege instantiates a new GitFilePushRulePrivilege object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitFilePushRulePrivilegeWithDefaults

`func NewGitFilePushRulePrivilegeWithDefaults() *GitFilePushRulePrivilege`

NewGitFilePushRulePrivilegeWithDefaults instantiates a new GitFilePushRulePrivilege object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDeny

`func (o *GitFilePushRulePrivilege) GetIsDeny() bool`

GetIsDeny returns the IsDeny field if non-nil, zero value otherwise.

### GetIsDenyOk

`func (o *GitFilePushRulePrivilege) GetIsDenyOk() (*bool, bool)`

GetIsDenyOk returns a tuple with the IsDeny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeny

`func (o *GitFilePushRulePrivilege) SetIsDeny(v bool)`

SetIsDeny sets IsDeny field to given value.

### HasIsDeny

`func (o *GitFilePushRulePrivilege) HasIsDeny() bool`

HasIsDeny returns a boolean if a field has been set.

### GetIsRole

`func (o *GitFilePushRulePrivilege) GetIsRole() bool`

GetIsRole returns the IsRole field if non-nil, zero value otherwise.

### GetIsRoleOk

`func (o *GitFilePushRulePrivilege) GetIsRoleOk() (*bool, bool)`

GetIsRoleOk returns a tuple with the IsRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRole

`func (o *GitFilePushRulePrivilege) SetIsRole(v bool)`

SetIsRole sets IsRole field to given value.

### HasIsRole

`func (o *GitFilePushRulePrivilege) HasIsRole() bool`

HasIsRole returns a boolean if a field has been set.

### GetIsUser

`func (o *GitFilePushRulePrivilege) GetIsUser() bool`

GetIsUser returns the IsUser field if non-nil, zero value otherwise.

### GetIsUserOk

`func (o *GitFilePushRulePrivilege) GetIsUserOk() (*bool, bool)`

GetIsUserOk returns a tuple with the IsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUser

`func (o *GitFilePushRulePrivilege) SetIsUser(v bool)`

SetIsUser sets IsUser field to given value.

### HasIsUser

`func (o *GitFilePushRulePrivilege) HasIsUser() bool`

HasIsUser returns a boolean if a field has been set.

### GetRole

`func (o *GitFilePushRulePrivilege) GetRole() GitFilePushRuleRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GitFilePushRulePrivilege) GetRoleOk() (*GitFilePushRuleRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GitFilePushRulePrivilege) SetRole(v GitFilePushRuleRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *GitFilePushRulePrivilege) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUser

`func (o *GitFilePushRulePrivilege) GetUser() GitFilePushRuleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GitFilePushRulePrivilege) GetUserOk() (*GitFilePushRuleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GitFilePushRulePrivilege) SetUser(v GitFilePushRuleUser)`

SetUser sets User field to given value.

### HasUser

`func (o *GitFilePushRulePrivilege) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


