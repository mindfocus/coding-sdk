# ModifyGitTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库 ID | 
**DepotPath** | Pointer to **string** | 仓库路径，与仓库ID二选一 | [optional] 
**TargetProjectId** | **int64** | 目标项目 ID | 

## Methods

### NewModifyGitTransferRequest

`func NewModifyGitTransferRequest(depotId int64, targetProjectId int64, ) *ModifyGitTransferRequest`

NewModifyGitTransferRequest instantiates a new ModifyGitTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyGitTransferRequestWithDefaults

`func NewModifyGitTransferRequestWithDefaults() *ModifyGitTransferRequest`

NewModifyGitTransferRequestWithDefaults instantiates a new ModifyGitTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *ModifyGitTransferRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ModifyGitTransferRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ModifyGitTransferRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *ModifyGitTransferRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyGitTransferRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyGitTransferRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *ModifyGitTransferRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetTargetProjectId

`func (o *ModifyGitTransferRequest) GetTargetProjectId() int64`

GetTargetProjectId returns the TargetProjectId field if non-nil, zero value otherwise.

### GetTargetProjectIdOk

`func (o *ModifyGitTransferRequest) GetTargetProjectIdOk() (*int64, bool)`

GetTargetProjectIdOk returns a tuple with the TargetProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProjectId

`func (o *ModifyGitTransferRequest) SetTargetProjectId(v int64)`

SetTargetProjectId sets TargetProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

