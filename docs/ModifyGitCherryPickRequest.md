# ModifyGitCherryPickRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | **string** | 分支名称 | 
**DepotId** | **int64** | 仓库 ID | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**Sha** | **string** | 提交 ID | 

## Methods

### NewModifyGitCherryPickRequest

`func NewModifyGitCherryPickRequest(branchName string, depotId int64, sha string, ) *ModifyGitCherryPickRequest`

NewModifyGitCherryPickRequest instantiates a new ModifyGitCherryPickRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyGitCherryPickRequestWithDefaults

`func NewModifyGitCherryPickRequestWithDefaults() *ModifyGitCherryPickRequest`

NewModifyGitCherryPickRequestWithDefaults instantiates a new ModifyGitCherryPickRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchName

`func (o *ModifyGitCherryPickRequest) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *ModifyGitCherryPickRequest) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *ModifyGitCherryPickRequest) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.


### GetDepotId

`func (o *ModifyGitCherryPickRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ModifyGitCherryPickRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ModifyGitCherryPickRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *ModifyGitCherryPickRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyGitCherryPickRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyGitCherryPickRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *ModifyGitCherryPickRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetSha

`func (o *ModifyGitCherryPickRequest) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *ModifyGitCherryPickRequest) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *ModifyGitCherryPickRequest) SetSha(v string)`

SetSha sets Sha field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


