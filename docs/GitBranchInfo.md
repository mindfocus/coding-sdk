# GitBranchInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | Pointer to **string** | 分支名称 | [optional] [default to ""]
**DenyForcePush** | Pointer to **NullableBool** | 是否禁止强制推送 | [optional] [default to false]
**ForceSquash** | Pointer to **NullableBool** | 是否可以ForceSquash | [optional] [default to false]
**IsDefaultBranch** | Pointer to **bool** | 是否默认分支 | [optional] [default to false]
**IsProtected** | Pointer to **bool** | 是否保护分支 | [optional] [default to false]
**IsReadOnly** | Pointer to **NullableBool** | 是否只读 | [optional] [default to false]
**LastCommitDate** | Pointer to **NullableInt64** | 最后提交时间 | [optional] 
**Sha** | Pointer to **string** | 分支sha值 | [optional] [default to ""]
**SpecBranchType** | Pointer to **NullableString** | 仓库规范类型 | [optional] [default to ""]

## Methods

### NewGitBranchInfo

`func NewGitBranchInfo() *GitBranchInfo`

NewGitBranchInfo instantiates a new GitBranchInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitBranchInfoWithDefaults

`func NewGitBranchInfoWithDefaults() *GitBranchInfo`

NewGitBranchInfoWithDefaults instantiates a new GitBranchInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchName

`func (o *GitBranchInfo) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *GitBranchInfo) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *GitBranchInfo) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *GitBranchInfo) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetDenyForcePush

`func (o *GitBranchInfo) GetDenyForcePush() bool`

GetDenyForcePush returns the DenyForcePush field if non-nil, zero value otherwise.

### GetDenyForcePushOk

`func (o *GitBranchInfo) GetDenyForcePushOk() (*bool, bool)`

GetDenyForcePushOk returns a tuple with the DenyForcePush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyForcePush

`func (o *GitBranchInfo) SetDenyForcePush(v bool)`

SetDenyForcePush sets DenyForcePush field to given value.

### HasDenyForcePush

`func (o *GitBranchInfo) HasDenyForcePush() bool`

HasDenyForcePush returns a boolean if a field has been set.

### SetDenyForcePushNil

`func (o *GitBranchInfo) SetDenyForcePushNil(b bool)`

 SetDenyForcePushNil sets the value for DenyForcePush to be an explicit nil

### UnsetDenyForcePush
`func (o *GitBranchInfo) UnsetDenyForcePush()`

UnsetDenyForcePush ensures that no value is present for DenyForcePush, not even an explicit nil
### GetForceSquash

`func (o *GitBranchInfo) GetForceSquash() bool`

GetForceSquash returns the ForceSquash field if non-nil, zero value otherwise.

### GetForceSquashOk

`func (o *GitBranchInfo) GetForceSquashOk() (*bool, bool)`

GetForceSquashOk returns a tuple with the ForceSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSquash

`func (o *GitBranchInfo) SetForceSquash(v bool)`

SetForceSquash sets ForceSquash field to given value.

### HasForceSquash

`func (o *GitBranchInfo) HasForceSquash() bool`

HasForceSquash returns a boolean if a field has been set.

### SetForceSquashNil

`func (o *GitBranchInfo) SetForceSquashNil(b bool)`

 SetForceSquashNil sets the value for ForceSquash to be an explicit nil

### UnsetForceSquash
`func (o *GitBranchInfo) UnsetForceSquash()`

UnsetForceSquash ensures that no value is present for ForceSquash, not even an explicit nil
### GetIsDefaultBranch

`func (o *GitBranchInfo) GetIsDefaultBranch() bool`

GetIsDefaultBranch returns the IsDefaultBranch field if non-nil, zero value otherwise.

### GetIsDefaultBranchOk

`func (o *GitBranchInfo) GetIsDefaultBranchOk() (*bool, bool)`

GetIsDefaultBranchOk returns a tuple with the IsDefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultBranch

`func (o *GitBranchInfo) SetIsDefaultBranch(v bool)`

SetIsDefaultBranch sets IsDefaultBranch field to given value.

### HasIsDefaultBranch

`func (o *GitBranchInfo) HasIsDefaultBranch() bool`

HasIsDefaultBranch returns a boolean if a field has been set.

### GetIsProtected

`func (o *GitBranchInfo) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *GitBranchInfo) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *GitBranchInfo) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *GitBranchInfo) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *GitBranchInfo) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *GitBranchInfo) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *GitBranchInfo) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *GitBranchInfo) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### SetIsReadOnlyNil

`func (o *GitBranchInfo) SetIsReadOnlyNil(b bool)`

 SetIsReadOnlyNil sets the value for IsReadOnly to be an explicit nil

### UnsetIsReadOnly
`func (o *GitBranchInfo) UnsetIsReadOnly()`

UnsetIsReadOnly ensures that no value is present for IsReadOnly, not even an explicit nil
### GetLastCommitDate

`func (o *GitBranchInfo) GetLastCommitDate() int64`

GetLastCommitDate returns the LastCommitDate field if non-nil, zero value otherwise.

### GetLastCommitDateOk

`func (o *GitBranchInfo) GetLastCommitDateOk() (*int64, bool)`

GetLastCommitDateOk returns a tuple with the LastCommitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommitDate

`func (o *GitBranchInfo) SetLastCommitDate(v int64)`

SetLastCommitDate sets LastCommitDate field to given value.

### HasLastCommitDate

`func (o *GitBranchInfo) HasLastCommitDate() bool`

HasLastCommitDate returns a boolean if a field has been set.

### SetLastCommitDateNil

`func (o *GitBranchInfo) SetLastCommitDateNil(b bool)`

 SetLastCommitDateNil sets the value for LastCommitDate to be an explicit nil

### UnsetLastCommitDate
`func (o *GitBranchInfo) UnsetLastCommitDate()`

UnsetLastCommitDate ensures that no value is present for LastCommitDate, not even an explicit nil
### GetSha

`func (o *GitBranchInfo) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitBranchInfo) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitBranchInfo) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *GitBranchInfo) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetSpecBranchType

`func (o *GitBranchInfo) GetSpecBranchType() string`

GetSpecBranchType returns the SpecBranchType field if non-nil, zero value otherwise.

### GetSpecBranchTypeOk

`func (o *GitBranchInfo) GetSpecBranchTypeOk() (*string, bool)`

GetSpecBranchTypeOk returns a tuple with the SpecBranchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecBranchType

`func (o *GitBranchInfo) SetSpecBranchType(v string)`

SetSpecBranchType sets SpecBranchType field to given value.

### HasSpecBranchType

`func (o *GitBranchInfo) HasSpecBranchType() bool`

HasSpecBranchType returns a boolean if a field has been set.

### SetSpecBranchTypeNil

`func (o *GitBranchInfo) SetSpecBranchTypeNil(b bool)`

 SetSpecBranchTypeNil sets the value for SpecBranchType to be an explicit nil

### UnsetSpecBranchType
`func (o *GitBranchInfo) UnsetSpecBranchType()`

UnsetSpecBranchType ensures that no value is present for SpecBranchType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


