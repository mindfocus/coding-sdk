# ModifyDepotDescriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库 ID | 
**DepotPath** | Pointer to **string** | 仓库路径，与仓库ID二选一 | [optional] 
**Description** | **string** | 仓库描述信息 | 
**UserId** | **int64** | 用户 ID | 

## Methods

### NewModifyDepotDescriptionRequest

`func NewModifyDepotDescriptionRequest(depotId int64, description string, userId int64, ) *ModifyDepotDescriptionRequest`

NewModifyDepotDescriptionRequest instantiates a new ModifyDepotDescriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyDepotDescriptionRequestWithDefaults

`func NewModifyDepotDescriptionRequestWithDefaults() *ModifyDepotDescriptionRequest`

NewModifyDepotDescriptionRequestWithDefaults instantiates a new ModifyDepotDescriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *ModifyDepotDescriptionRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ModifyDepotDescriptionRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ModifyDepotDescriptionRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *ModifyDepotDescriptionRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyDepotDescriptionRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyDepotDescriptionRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *ModifyDepotDescriptionRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetDescription

`func (o *ModifyDepotDescriptionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModifyDepotDescriptionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModifyDepotDescriptionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUserId

`func (o *ModifyDepotDescriptionRequest) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModifyDepotDescriptionRequest) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModifyDepotDescriptionRequest) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


