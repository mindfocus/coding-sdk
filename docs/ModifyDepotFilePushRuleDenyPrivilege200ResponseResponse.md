# ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitFilePushRule** | Pointer to [**[]GitFilePushRule**](GitFilePushRule.md) | git 仓库文件推送规则列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewModifyDepotFilePushRuleDenyPrivilege200ResponseResponse

`func NewModifyDepotFilePushRuleDenyPrivilege200ResponseResponse() *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse`

NewModifyDepotFilePushRuleDenyPrivilege200ResponseResponse instantiates a new ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyDepotFilePushRuleDenyPrivilege200ResponseResponseWithDefaults

`func NewModifyDepotFilePushRuleDenyPrivilege200ResponseResponseWithDefaults() *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse`

NewModifyDepotFilePushRuleDenyPrivilege200ResponseResponseWithDefaults instantiates a new ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitFilePushRule

`func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) GetGitFilePushRule() []GitFilePushRule`

GetGitFilePushRule returns the GitFilePushRule field if non-nil, zero value otherwise.

### GetGitFilePushRuleOk

`func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) GetGitFilePushRuleOk() (*[]GitFilePushRule, bool)`

GetGitFilePushRuleOk returns a tuple with the GitFilePushRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitFilePushRule

`func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) SetGitFilePushRule(v []GitFilePushRule)`

SetGitFilePushRule sets GitFilePushRule field to given value.

### HasGitFilePushRule

`func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) HasGitFilePushRule() bool`

HasGitFilePushRule returns a boolean if a field has been set.

### SetGitFilePushRuleNil

`func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) SetGitFilePushRuleNil(b bool)`

 SetGitFilePushRuleNil sets the value for GitFilePushRule to be an explicit nil

### UnsetGitFilePushRule
`func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) UnsetGitFilePushRule()`

UnsetGitFilePushRule ensures that no value is present for GitFilePushRule, not even an explicit nil
### GetRequestId

`func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


