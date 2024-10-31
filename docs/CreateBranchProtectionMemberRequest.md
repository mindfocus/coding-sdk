# CreateBranchProtectionMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPush** | **bool** | 是否允许直接推送 | 
**BranchProtectionId** | **int64** | 保护分支规则id | 
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径 | [optional] 
**UserGlobalKey** | **string** | 用户globalkey | 

## Methods

### NewCreateBranchProtectionMemberRequest

`func NewCreateBranchProtectionMemberRequest(allowPush bool, branchProtectionId int64, depotId int64, userGlobalKey string, ) *CreateBranchProtectionMemberRequest`

NewCreateBranchProtectionMemberRequest instantiates a new CreateBranchProtectionMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBranchProtectionMemberRequestWithDefaults

`func NewCreateBranchProtectionMemberRequestWithDefaults() *CreateBranchProtectionMemberRequest`

NewCreateBranchProtectionMemberRequestWithDefaults instantiates a new CreateBranchProtectionMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPush

`func (o *CreateBranchProtectionMemberRequest) GetAllowPush() bool`

GetAllowPush returns the AllowPush field if non-nil, zero value otherwise.

### GetAllowPushOk

`func (o *CreateBranchProtectionMemberRequest) GetAllowPushOk() (*bool, bool)`

GetAllowPushOk returns a tuple with the AllowPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPush

`func (o *CreateBranchProtectionMemberRequest) SetAllowPush(v bool)`

SetAllowPush sets AllowPush field to given value.


### GetBranchProtectionId

`func (o *CreateBranchProtectionMemberRequest) GetBranchProtectionId() int64`

GetBranchProtectionId returns the BranchProtectionId field if non-nil, zero value otherwise.

### GetBranchProtectionIdOk

`func (o *CreateBranchProtectionMemberRequest) GetBranchProtectionIdOk() (*int64, bool)`

GetBranchProtectionIdOk returns a tuple with the BranchProtectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchProtectionId

`func (o *CreateBranchProtectionMemberRequest) SetBranchProtectionId(v int64)`

SetBranchProtectionId sets BranchProtectionId field to given value.


### GetDepotId

`func (o *CreateBranchProtectionMemberRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CreateBranchProtectionMemberRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CreateBranchProtectionMemberRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *CreateBranchProtectionMemberRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateBranchProtectionMemberRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateBranchProtectionMemberRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *CreateBranchProtectionMemberRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetUserGlobalKey

`func (o *CreateBranchProtectionMemberRequest) GetUserGlobalKey() string`

GetUserGlobalKey returns the UserGlobalKey field if non-nil, zero value otherwise.

### GetUserGlobalKeyOk

`func (o *CreateBranchProtectionMemberRequest) GetUserGlobalKeyOk() (*string, bool)`

GetUserGlobalKeyOk returns a tuple with the UserGlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGlobalKey

`func (o *CreateBranchProtectionMemberRequest) SetUserGlobalKey(v string)`

SetUserGlobalKey sets UserGlobalKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


