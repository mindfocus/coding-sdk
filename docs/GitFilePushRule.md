# GitFilePushRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilePushRuleId** | Pointer to **int64** | 文件推送规则 ID | [optional] 
**IsDenyForAllUser** | Pointer to **bool** | 拒绝所有人推送 | [optional] [default to false]
**Pattern** | Pointer to **string** | 文件路径 | [optional] [default to ""]
**Privileges** | Pointer to [**[]GitFilePushRulePrivilege**](GitFilePushRulePrivilege.md) | 特权者列表 | [optional] 

## Methods

### NewGitFilePushRule

`func NewGitFilePushRule() *GitFilePushRule`

NewGitFilePushRule instantiates a new GitFilePushRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitFilePushRuleWithDefaults

`func NewGitFilePushRuleWithDefaults() *GitFilePushRule`

NewGitFilePushRuleWithDefaults instantiates a new GitFilePushRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilePushRuleId

`func (o *GitFilePushRule) GetFilePushRuleId() int64`

GetFilePushRuleId returns the FilePushRuleId field if non-nil, zero value otherwise.

### GetFilePushRuleIdOk

`func (o *GitFilePushRule) GetFilePushRuleIdOk() (*int64, bool)`

GetFilePushRuleIdOk returns a tuple with the FilePushRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePushRuleId

`func (o *GitFilePushRule) SetFilePushRuleId(v int64)`

SetFilePushRuleId sets FilePushRuleId field to given value.

### HasFilePushRuleId

`func (o *GitFilePushRule) HasFilePushRuleId() bool`

HasFilePushRuleId returns a boolean if a field has been set.

### GetIsDenyForAllUser

`func (o *GitFilePushRule) GetIsDenyForAllUser() bool`

GetIsDenyForAllUser returns the IsDenyForAllUser field if non-nil, zero value otherwise.

### GetIsDenyForAllUserOk

`func (o *GitFilePushRule) GetIsDenyForAllUserOk() (*bool, bool)`

GetIsDenyForAllUserOk returns a tuple with the IsDenyForAllUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDenyForAllUser

`func (o *GitFilePushRule) SetIsDenyForAllUser(v bool)`

SetIsDenyForAllUser sets IsDenyForAllUser field to given value.

### HasIsDenyForAllUser

`func (o *GitFilePushRule) HasIsDenyForAllUser() bool`

HasIsDenyForAllUser returns a boolean if a field has been set.

### GetPattern

`func (o *GitFilePushRule) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *GitFilePushRule) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *GitFilePushRule) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *GitFilePushRule) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetPrivileges

`func (o *GitFilePushRule) GetPrivileges() []GitFilePushRulePrivilege`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *GitFilePushRule) GetPrivilegesOk() (*[]GitFilePushRulePrivilege, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *GitFilePushRule) SetPrivileges(v []GitFilePushRulePrivilege)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *GitFilePushRule) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### SetPrivilegesNil

`func (o *GitFilePushRule) SetPrivilegesNil(b bool)`

 SetPrivilegesNil sets the value for Privileges to be an explicit nil

### UnsetPrivileges
`func (o *GitFilePushRule) UnsetPrivileges()`

UnsetPrivileges ensures that no value is present for Privileges, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


