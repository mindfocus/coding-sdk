# ModifyGitCommitStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitSha** | **string** | 提交id | 
**Context** | **string** | 流水线文本内容 | 
**DepotId** | Pointer to **int64** | 仓库id | [optional] 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**Description** | **string** | 流水线状态描述 | 
**State** | **string** | 提交状态 | 
**TargetURL** | **string** | 流水线链接地址 | 

## Methods

### NewModifyGitCommitStatusRequest

`func NewModifyGitCommitStatusRequest(commitSha string, context string, description string, state string, targetURL string, ) *ModifyGitCommitStatusRequest`

NewModifyGitCommitStatusRequest instantiates a new ModifyGitCommitStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyGitCommitStatusRequestWithDefaults

`func NewModifyGitCommitStatusRequestWithDefaults() *ModifyGitCommitStatusRequest`

NewModifyGitCommitStatusRequestWithDefaults instantiates a new ModifyGitCommitStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitSha

`func (o *ModifyGitCommitStatusRequest) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *ModifyGitCommitStatusRequest) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *ModifyGitCommitStatusRequest) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### GetContext

`func (o *ModifyGitCommitStatusRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ModifyGitCommitStatusRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ModifyGitCommitStatusRequest) SetContext(v string)`

SetContext sets Context field to given value.


### GetDepotId

`func (o *ModifyGitCommitStatusRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ModifyGitCommitStatusRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ModifyGitCommitStatusRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.

### HasDepotId

`func (o *ModifyGitCommitStatusRequest) HasDepotId() bool`

HasDepotId returns a boolean if a field has been set.

### GetDepotPath

`func (o *ModifyGitCommitStatusRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyGitCommitStatusRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyGitCommitStatusRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *ModifyGitCommitStatusRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetDescription

`func (o *ModifyGitCommitStatusRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModifyGitCommitStatusRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModifyGitCommitStatusRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetState

`func (o *ModifyGitCommitStatusRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModifyGitCommitStatusRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModifyGitCommitStatusRequest) SetState(v string)`

SetState sets State field to given value.


### GetTargetURL

`func (o *ModifyGitCommitStatusRequest) GetTargetURL() string`

GetTargetURL returns the TargetURL field if non-nil, zero value otherwise.

### GetTargetURLOk

`func (o *ModifyGitCommitStatusRequest) GetTargetURLOk() (*string, bool)`

GetTargetURLOk returns a tuple with the TargetURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetURL

`func (o *ModifyGitCommitStatusRequest) SetTargetURL(v string)`

SetTargetURL sets TargetURL field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


