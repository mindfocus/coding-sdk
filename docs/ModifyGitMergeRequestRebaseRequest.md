# ModifyGitMergeRequestRebaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库 ID | 
**DepotPat** | Pointer to **string** | 仓库路径，与仓库ID二选一 | [optional] 
**MergeId** | **int64** | 合并请求 资源ID | 

## Methods

### NewModifyGitMergeRequestRebaseRequest

`func NewModifyGitMergeRequestRebaseRequest(depotId int64, mergeId int64, ) *ModifyGitMergeRequestRebaseRequest`

NewModifyGitMergeRequestRebaseRequest instantiates a new ModifyGitMergeRequestRebaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyGitMergeRequestRebaseRequestWithDefaults

`func NewModifyGitMergeRequestRebaseRequestWithDefaults() *ModifyGitMergeRequestRebaseRequest`

NewModifyGitMergeRequestRebaseRequestWithDefaults instantiates a new ModifyGitMergeRequestRebaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *ModifyGitMergeRequestRebaseRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ModifyGitMergeRequestRebaseRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ModifyGitMergeRequestRebaseRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPat

`func (o *ModifyGitMergeRequestRebaseRequest) GetDepotPat() string`

GetDepotPat returns the DepotPat field if non-nil, zero value otherwise.

### GetDepotPatOk

`func (o *ModifyGitMergeRequestRebaseRequest) GetDepotPatOk() (*string, bool)`

GetDepotPatOk returns a tuple with the DepotPat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPat

`func (o *ModifyGitMergeRequestRebaseRequest) SetDepotPat(v string)`

SetDepotPat sets DepotPat field to given value.

### HasDepotPat

`func (o *ModifyGitMergeRequestRebaseRequest) HasDepotPat() bool`

HasDepotPat returns a boolean if a field has been set.

### GetMergeId

`func (o *ModifyGitMergeRequestRebaseRequest) GetMergeId() int64`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *ModifyGitMergeRequestRebaseRequest) GetMergeIdOk() (*int64, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *ModifyGitMergeRequestRebaseRequest) SetMergeId(v int64)`

SetMergeId sets MergeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


