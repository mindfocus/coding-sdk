# ModifyGitMergeBranchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitMessage** | **string** | 提交信息 | 
**DepotPath** | **string** | 仓库路径或仓库ID | 
**FromBranch** | **string** | 源分支 | 
**ToBranch** | **string** | 目标分支 | 

## Methods

### NewModifyGitMergeBranchRequest

`func NewModifyGitMergeBranchRequest(commitMessage string, depotPath string, fromBranch string, toBranch string, ) *ModifyGitMergeBranchRequest`

NewModifyGitMergeBranchRequest instantiates a new ModifyGitMergeBranchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyGitMergeBranchRequestWithDefaults

`func NewModifyGitMergeBranchRequestWithDefaults() *ModifyGitMergeBranchRequest`

NewModifyGitMergeBranchRequestWithDefaults instantiates a new ModifyGitMergeBranchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitMessage

`func (o *ModifyGitMergeBranchRequest) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *ModifyGitMergeBranchRequest) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *ModifyGitMergeBranchRequest) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.


### GetDepotPath

`func (o *ModifyGitMergeBranchRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyGitMergeBranchRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyGitMergeBranchRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetFromBranch

`func (o *ModifyGitMergeBranchRequest) GetFromBranch() string`

GetFromBranch returns the FromBranch field if non-nil, zero value otherwise.

### GetFromBranchOk

`func (o *ModifyGitMergeBranchRequest) GetFromBranchOk() (*string, bool)`

GetFromBranchOk returns a tuple with the FromBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromBranch

`func (o *ModifyGitMergeBranchRequest) SetFromBranch(v string)`

SetFromBranch sets FromBranch field to given value.


### GetToBranch

`func (o *ModifyGitMergeBranchRequest) GetToBranch() string`

GetToBranch returns the ToBranch field if non-nil, zero value otherwise.

### GetToBranchOk

`func (o *ModifyGitMergeBranchRequest) GetToBranchOk() (*string, bool)`

GetToBranchOk returns a tuple with the ToBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToBranch

`func (o *ModifyGitMergeBranchRequest) SetToBranch(v string)`

SetToBranch sets ToBranch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


