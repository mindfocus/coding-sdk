# ModifyMergeMRRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**IsDelSourceBranch** | **bool** | 是否删除源分支 | 
**IsFastForward** | **bool** | 是否 Fast Forward 模式 | 
**MergeId** | **int64** | 合并请求id | 
**Message** | **string** | 合并信息 | 
**Squash** | **bool** | 是否需要对mr进行 Squash | 

## Methods

### NewModifyMergeMRRequest

`func NewModifyMergeMRRequest(depotId int64, isDelSourceBranch bool, isFastForward bool, mergeId int64, message string, squash bool, ) *ModifyMergeMRRequest`

NewModifyMergeMRRequest instantiates a new ModifyMergeMRRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyMergeMRRequestWithDefaults

`func NewModifyMergeMRRequestWithDefaults() *ModifyMergeMRRequest`

NewModifyMergeMRRequestWithDefaults instantiates a new ModifyMergeMRRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *ModifyMergeMRRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ModifyMergeMRRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ModifyMergeMRRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *ModifyMergeMRRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyMergeMRRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyMergeMRRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *ModifyMergeMRRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetIsDelSourceBranch

`func (o *ModifyMergeMRRequest) GetIsDelSourceBranch() bool`

GetIsDelSourceBranch returns the IsDelSourceBranch field if non-nil, zero value otherwise.

### GetIsDelSourceBranchOk

`func (o *ModifyMergeMRRequest) GetIsDelSourceBranchOk() (*bool, bool)`

GetIsDelSourceBranchOk returns a tuple with the IsDelSourceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDelSourceBranch

`func (o *ModifyMergeMRRequest) SetIsDelSourceBranch(v bool)`

SetIsDelSourceBranch sets IsDelSourceBranch field to given value.


### GetIsFastForward

`func (o *ModifyMergeMRRequest) GetIsFastForward() bool`

GetIsFastForward returns the IsFastForward field if non-nil, zero value otherwise.

### GetIsFastForwardOk

`func (o *ModifyMergeMRRequest) GetIsFastForwardOk() (*bool, bool)`

GetIsFastForwardOk returns a tuple with the IsFastForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFastForward

`func (o *ModifyMergeMRRequest) SetIsFastForward(v bool)`

SetIsFastForward sets IsFastForward field to given value.


### GetMergeId

`func (o *ModifyMergeMRRequest) GetMergeId() int64`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *ModifyMergeMRRequest) GetMergeIdOk() (*int64, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *ModifyMergeMRRequest) SetMergeId(v int64)`

SetMergeId sets MergeId field to given value.


### GetMessage

`func (o *ModifyMergeMRRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModifyMergeMRRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModifyMergeMRRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSquash

`func (o *ModifyMergeMRRequest) GetSquash() bool`

GetSquash returns the Squash field if non-nil, zero value otherwise.

### GetSquashOk

`func (o *ModifyMergeMRRequest) GetSquashOk() (*bool, bool)`

GetSquashOk returns a tuple with the Squash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquash

`func (o *ModifyMergeMRRequest) SetSquash(v bool)`

SetSquash sets Squash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


