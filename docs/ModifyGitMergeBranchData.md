# ModifyGitMergeBranchData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MergeCommit** | Pointer to [**GitCommit**](GitCommit.md) |  | [optional] 
**MergeStatus** | Pointer to **NullableString** | 合并状态 MERGED 合并成功 FAILED 合并失败，异常原因 NOT_MERGEABLE 不可合并，有代码冲突 ALREADY_MERGED 两个分支已经合并，两个分支是一样的或者目标分支的已经合并进当前分支了 | [optional] [default to ""]

## Methods

### NewModifyGitMergeBranchData

`func NewModifyGitMergeBranchData() *ModifyGitMergeBranchData`

NewModifyGitMergeBranchData instantiates a new ModifyGitMergeBranchData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyGitMergeBranchDataWithDefaults

`func NewModifyGitMergeBranchDataWithDefaults() *ModifyGitMergeBranchData`

NewModifyGitMergeBranchDataWithDefaults instantiates a new ModifyGitMergeBranchData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMergeCommit

`func (o *ModifyGitMergeBranchData) GetMergeCommit() GitCommit`

GetMergeCommit returns the MergeCommit field if non-nil, zero value otherwise.

### GetMergeCommitOk

`func (o *ModifyGitMergeBranchData) GetMergeCommitOk() (*GitCommit, bool)`

GetMergeCommitOk returns a tuple with the MergeCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommit

`func (o *ModifyGitMergeBranchData) SetMergeCommit(v GitCommit)`

SetMergeCommit sets MergeCommit field to given value.

### HasMergeCommit

`func (o *ModifyGitMergeBranchData) HasMergeCommit() bool`

HasMergeCommit returns a boolean if a field has been set.

### GetMergeStatus

`func (o *ModifyGitMergeBranchData) GetMergeStatus() string`

GetMergeStatus returns the MergeStatus field if non-nil, zero value otherwise.

### GetMergeStatusOk

`func (o *ModifyGitMergeBranchData) GetMergeStatusOk() (*string, bool)`

GetMergeStatusOk returns a tuple with the MergeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeStatus

`func (o *ModifyGitMergeBranchData) SetMergeStatus(v string)`

SetMergeStatus sets MergeStatus field to given value.

### HasMergeStatus

`func (o *ModifyGitMergeBranchData) HasMergeStatus() bool`

HasMergeStatus returns a boolean if a field has been set.

### SetMergeStatusNil

`func (o *ModifyGitMergeBranchData) SetMergeStatusNil(b bool)`

 SetMergeStatusNil sets the value for MergeStatus to be an explicit nil

### UnsetMergeStatus
`func (o *ModifyGitMergeBranchData) UnsetMergeStatus()`

UnsetMergeStatus ensures that no value is present for MergeStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


