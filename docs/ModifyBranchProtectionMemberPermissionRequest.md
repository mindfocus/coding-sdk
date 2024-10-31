# ModifyBranchProtectionMemberPermissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPush** | **bool** | 是否直接推送 | 
**BranchRuleName** | **string** | 分支规则名称 | 
**DepotPath** | **string** | 仓库路径 | 
**UserGlobalKey** | **string** | 用户GlobalKey | 

## Methods

### NewModifyBranchProtectionMemberPermissionRequest

`func NewModifyBranchProtectionMemberPermissionRequest(allowPush bool, branchRuleName string, depotPath string, userGlobalKey string, ) *ModifyBranchProtectionMemberPermissionRequest`

NewModifyBranchProtectionMemberPermissionRequest instantiates a new ModifyBranchProtectionMemberPermissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyBranchProtectionMemberPermissionRequestWithDefaults

`func NewModifyBranchProtectionMemberPermissionRequestWithDefaults() *ModifyBranchProtectionMemberPermissionRequest`

NewModifyBranchProtectionMemberPermissionRequestWithDefaults instantiates a new ModifyBranchProtectionMemberPermissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPush

`func (o *ModifyBranchProtectionMemberPermissionRequest) GetAllowPush() bool`

GetAllowPush returns the AllowPush field if non-nil, zero value otherwise.

### GetAllowPushOk

`func (o *ModifyBranchProtectionMemberPermissionRequest) GetAllowPushOk() (*bool, bool)`

GetAllowPushOk returns a tuple with the AllowPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPush

`func (o *ModifyBranchProtectionMemberPermissionRequest) SetAllowPush(v bool)`

SetAllowPush sets AllowPush field to given value.


### GetBranchRuleName

`func (o *ModifyBranchProtectionMemberPermissionRequest) GetBranchRuleName() string`

GetBranchRuleName returns the BranchRuleName field if non-nil, zero value otherwise.

### GetBranchRuleNameOk

`func (o *ModifyBranchProtectionMemberPermissionRequest) GetBranchRuleNameOk() (*string, bool)`

GetBranchRuleNameOk returns a tuple with the BranchRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchRuleName

`func (o *ModifyBranchProtectionMemberPermissionRequest) SetBranchRuleName(v string)`

SetBranchRuleName sets BranchRuleName field to given value.


### GetDepotPath

`func (o *ModifyBranchProtectionMemberPermissionRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyBranchProtectionMemberPermissionRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyBranchProtectionMemberPermissionRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetUserGlobalKey

`func (o *ModifyBranchProtectionMemberPermissionRequest) GetUserGlobalKey() string`

GetUserGlobalKey returns the UserGlobalKey field if non-nil, zero value otherwise.

### GetUserGlobalKeyOk

`func (o *ModifyBranchProtectionMemberPermissionRequest) GetUserGlobalKeyOk() (*string, bool)`

GetUserGlobalKeyOk returns a tuple with the UserGlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGlobalKey

`func (o *ModifyBranchProtectionMemberPermissionRequest) SetUserGlobalKey(v string)`

SetUserGlobalKey sets UserGlobalKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


