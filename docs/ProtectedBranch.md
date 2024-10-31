# ProtectedBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitDate** | Pointer to **int64** | 提交时间戳,单位毫秒 | [optional] 
**DenyForcePush** | Pointer to **bool** | 是否禁止强制推送 | [optional] [default to false]
**ForceSquash** | Pointer to **bool** | 是否使用 squash 推送 | [optional] [default to false]
**Name** | Pointer to **string** | 名称 | [optional] [default to ""]
**StatusCheck** | Pointer to **bool** | 是否开启状态检查 | [optional] [default to false]

## Methods

### NewProtectedBranch

`func NewProtectedBranch() *ProtectedBranch`

NewProtectedBranch instantiates a new ProtectedBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedBranchWithDefaults

`func NewProtectedBranchWithDefaults() *ProtectedBranch`

NewProtectedBranchWithDefaults instantiates a new ProtectedBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitDate

`func (o *ProtectedBranch) GetCommitDate() int64`

GetCommitDate returns the CommitDate field if non-nil, zero value otherwise.

### GetCommitDateOk

`func (o *ProtectedBranch) GetCommitDateOk() (*int64, bool)`

GetCommitDateOk returns a tuple with the CommitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitDate

`func (o *ProtectedBranch) SetCommitDate(v int64)`

SetCommitDate sets CommitDate field to given value.

### HasCommitDate

`func (o *ProtectedBranch) HasCommitDate() bool`

HasCommitDate returns a boolean if a field has been set.

### GetDenyForcePush

`func (o *ProtectedBranch) GetDenyForcePush() bool`

GetDenyForcePush returns the DenyForcePush field if non-nil, zero value otherwise.

### GetDenyForcePushOk

`func (o *ProtectedBranch) GetDenyForcePushOk() (*bool, bool)`

GetDenyForcePushOk returns a tuple with the DenyForcePush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyForcePush

`func (o *ProtectedBranch) SetDenyForcePush(v bool)`

SetDenyForcePush sets DenyForcePush field to given value.

### HasDenyForcePush

`func (o *ProtectedBranch) HasDenyForcePush() bool`

HasDenyForcePush returns a boolean if a field has been set.

### GetForceSquash

`func (o *ProtectedBranch) GetForceSquash() bool`

GetForceSquash returns the ForceSquash field if non-nil, zero value otherwise.

### GetForceSquashOk

`func (o *ProtectedBranch) GetForceSquashOk() (*bool, bool)`

GetForceSquashOk returns a tuple with the ForceSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSquash

`func (o *ProtectedBranch) SetForceSquash(v bool)`

SetForceSquash sets ForceSquash field to given value.

### HasForceSquash

`func (o *ProtectedBranch) HasForceSquash() bool`

HasForceSquash returns a boolean if a field has been set.

### GetName

`func (o *ProtectedBranch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtectedBranch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtectedBranch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProtectedBranch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatusCheck

`func (o *ProtectedBranch) GetStatusCheck() bool`

GetStatusCheck returns the StatusCheck field if non-nil, zero value otherwise.

### GetStatusCheckOk

`func (o *ProtectedBranch) GetStatusCheckOk() (*bool, bool)`

GetStatusCheckOk returns a tuple with the StatusCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCheck

`func (o *ProtectedBranch) SetStatusCheck(v bool)`

SetStatusCheck sets StatusCheck field to given value.

### HasStatusCheck

`func (o *ProtectedBranch) HasStatusCheck() bool`

HasStatusCheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


