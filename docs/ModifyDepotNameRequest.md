# ModifyDepotNameRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库 ID | 
**DepotName** | **string** | 仓库名称 | 
**DepotPath** | Pointer to **string** | 仓库路径，与仓库ID二选一 | [optional] 
**UserId** | **int64** | 用户 ID | 

## Methods

### NewModifyDepotNameRequest

`func NewModifyDepotNameRequest(depotId int64, depotName string, userId int64, ) *ModifyDepotNameRequest`

NewModifyDepotNameRequest instantiates a new ModifyDepotNameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyDepotNameRequestWithDefaults

`func NewModifyDepotNameRequestWithDefaults() *ModifyDepotNameRequest`

NewModifyDepotNameRequestWithDefaults instantiates a new ModifyDepotNameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *ModifyDepotNameRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ModifyDepotNameRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ModifyDepotNameRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotName

`func (o *ModifyDepotNameRequest) GetDepotName() string`

GetDepotName returns the DepotName field if non-nil, zero value otherwise.

### GetDepotNameOk

`func (o *ModifyDepotNameRequest) GetDepotNameOk() (*string, bool)`

GetDepotNameOk returns a tuple with the DepotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotName

`func (o *ModifyDepotNameRequest) SetDepotName(v string)`

SetDepotName sets DepotName field to given value.


### GetDepotPath

`func (o *ModifyDepotNameRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyDepotNameRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyDepotNameRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *ModifyDepotNameRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetUserId

`func (o *ModifyDepotNameRequest) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModifyDepotNameRequest) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModifyDepotNameRequest) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


