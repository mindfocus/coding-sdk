# ModifyGitRebaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseBranchName** | **string** | 基础分支名字 | 
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**SrcBranchName** | **string** | 源分支名字 | 

## Methods

### NewModifyGitRebaseRequest

`func NewModifyGitRebaseRequest(baseBranchName string, depotId int64, srcBranchName string, ) *ModifyGitRebaseRequest`

NewModifyGitRebaseRequest instantiates a new ModifyGitRebaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyGitRebaseRequestWithDefaults

`func NewModifyGitRebaseRequestWithDefaults() *ModifyGitRebaseRequest`

NewModifyGitRebaseRequestWithDefaults instantiates a new ModifyGitRebaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseBranchName

`func (o *ModifyGitRebaseRequest) GetBaseBranchName() string`

GetBaseBranchName returns the BaseBranchName field if non-nil, zero value otherwise.

### GetBaseBranchNameOk

`func (o *ModifyGitRebaseRequest) GetBaseBranchNameOk() (*string, bool)`

GetBaseBranchNameOk returns a tuple with the BaseBranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseBranchName

`func (o *ModifyGitRebaseRequest) SetBaseBranchName(v string)`

SetBaseBranchName sets BaseBranchName field to given value.


### GetDepotId

`func (o *ModifyGitRebaseRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ModifyGitRebaseRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ModifyGitRebaseRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *ModifyGitRebaseRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyGitRebaseRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyGitRebaseRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *ModifyGitRebaseRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetSrcBranchName

`func (o *ModifyGitRebaseRequest) GetSrcBranchName() string`

GetSrcBranchName returns the SrcBranchName field if non-nil, zero value otherwise.

### GetSrcBranchNameOk

`func (o *ModifyGitRebaseRequest) GetSrcBranchNameOk() (*string, bool)`

GetSrcBranchNameOk returns a tuple with the SrcBranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcBranchName

`func (o *ModifyGitRebaseRequest) SetSrcBranchName(v string)`

SetSrcBranchName sets SrcBranchName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


