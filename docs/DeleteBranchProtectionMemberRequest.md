# DeleteBranchProtectionMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchProtectionId** | **int64** | 保护分支规则id | 
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径 | [optional] 
**UserGlobalKey** | **string** | 用户globalkey | 

## Methods

### NewDeleteBranchProtectionMemberRequest

`func NewDeleteBranchProtectionMemberRequest(branchProtectionId int64, depotId int64, userGlobalKey string, ) *DeleteBranchProtectionMemberRequest`

NewDeleteBranchProtectionMemberRequest instantiates a new DeleteBranchProtectionMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteBranchProtectionMemberRequestWithDefaults

`func NewDeleteBranchProtectionMemberRequestWithDefaults() *DeleteBranchProtectionMemberRequest`

NewDeleteBranchProtectionMemberRequestWithDefaults instantiates a new DeleteBranchProtectionMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchProtectionId

`func (o *DeleteBranchProtectionMemberRequest) GetBranchProtectionId() int64`

GetBranchProtectionId returns the BranchProtectionId field if non-nil, zero value otherwise.

### GetBranchProtectionIdOk

`func (o *DeleteBranchProtectionMemberRequest) GetBranchProtectionIdOk() (*int64, bool)`

GetBranchProtectionIdOk returns a tuple with the BranchProtectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchProtectionId

`func (o *DeleteBranchProtectionMemberRequest) SetBranchProtectionId(v int64)`

SetBranchProtectionId sets BranchProtectionId field to given value.


### GetDepotId

`func (o *DeleteBranchProtectionMemberRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DeleteBranchProtectionMemberRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DeleteBranchProtectionMemberRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DeleteBranchProtectionMemberRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DeleteBranchProtectionMemberRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DeleteBranchProtectionMemberRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DeleteBranchProtectionMemberRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetUserGlobalKey

`func (o *DeleteBranchProtectionMemberRequest) GetUserGlobalKey() string`

GetUserGlobalKey returns the UserGlobalKey field if non-nil, zero value otherwise.

### GetUserGlobalKeyOk

`func (o *DeleteBranchProtectionMemberRequest) GetUserGlobalKeyOk() (*string, bool)`

GetUserGlobalKeyOk returns a tuple with the UserGlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGlobalKey

`func (o *DeleteBranchProtectionMemberRequest) SetUserGlobalKey(v string)`

SetUserGlobalKey sets UserGlobalKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


