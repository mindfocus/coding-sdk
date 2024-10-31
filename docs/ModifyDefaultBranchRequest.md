# ModifyDefaultBranchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | **string** | 分支名 | 
**DepotId** | **int64** | 仓库Id | 
**DepotPath** | Pointer to **string** | 仓库路径 | [optional] 
**UserId** | **int64** | 用户Id | 

## Methods

### NewModifyDefaultBranchRequest

`func NewModifyDefaultBranchRequest(branchName string, depotId int64, userId int64, ) *ModifyDefaultBranchRequest`

NewModifyDefaultBranchRequest instantiates a new ModifyDefaultBranchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyDefaultBranchRequestWithDefaults

`func NewModifyDefaultBranchRequestWithDefaults() *ModifyDefaultBranchRequest`

NewModifyDefaultBranchRequestWithDefaults instantiates a new ModifyDefaultBranchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchName

`func (o *ModifyDefaultBranchRequest) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *ModifyDefaultBranchRequest) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *ModifyDefaultBranchRequest) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.


### GetDepotId

`func (o *ModifyDefaultBranchRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ModifyDefaultBranchRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ModifyDefaultBranchRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *ModifyDefaultBranchRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyDefaultBranchRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyDefaultBranchRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *ModifyDefaultBranchRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetUserId

`func (o *ModifyDefaultBranchRequest) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModifyDefaultBranchRequest) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModifyDefaultBranchRequest) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


