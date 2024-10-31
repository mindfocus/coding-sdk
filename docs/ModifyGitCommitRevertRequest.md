# ModifyGitCommitRevertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | **string** | 分支名 | 
**DepotId** | **int64** | 仓库 ID | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**Message** | **string** | 提交描述 | 
**Sha** | **string** | 欲还原的提交 ID | 

## Methods

### NewModifyGitCommitRevertRequest

`func NewModifyGitCommitRevertRequest(branchName string, depotId int64, message string, sha string, ) *ModifyGitCommitRevertRequest`

NewModifyGitCommitRevertRequest instantiates a new ModifyGitCommitRevertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyGitCommitRevertRequestWithDefaults

`func NewModifyGitCommitRevertRequestWithDefaults() *ModifyGitCommitRevertRequest`

NewModifyGitCommitRevertRequestWithDefaults instantiates a new ModifyGitCommitRevertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchName

`func (o *ModifyGitCommitRevertRequest) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *ModifyGitCommitRevertRequest) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *ModifyGitCommitRevertRequest) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.


### GetDepotId

`func (o *ModifyGitCommitRevertRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ModifyGitCommitRevertRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ModifyGitCommitRevertRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *ModifyGitCommitRevertRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyGitCommitRevertRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyGitCommitRevertRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *ModifyGitCommitRevertRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetMessage

`func (o *ModifyGitCommitRevertRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModifyGitCommitRevertRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModifyGitCommitRevertRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSha

`func (o *ModifyGitCommitRevertRequest) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *ModifyGitCommitRevertRequest) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *ModifyGitCommitRevertRequest) SetSha(v string)`

SetSha sets Sha field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


