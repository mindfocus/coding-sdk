# ModifyDepotSharedSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径 | 
**IsShared** | **bool** | 仓库是否开源 | 

## Methods

### NewModifyDepotSharedSettingRequest

`func NewModifyDepotSharedSettingRequest(depotPath string, isShared bool, ) *ModifyDepotSharedSettingRequest`

NewModifyDepotSharedSettingRequest instantiates a new ModifyDepotSharedSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyDepotSharedSettingRequestWithDefaults

`func NewModifyDepotSharedSettingRequestWithDefaults() *ModifyDepotSharedSettingRequest`

NewModifyDepotSharedSettingRequestWithDefaults instantiates a new ModifyDepotSharedSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *ModifyDepotSharedSettingRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyDepotSharedSettingRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyDepotSharedSettingRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetIsShared

`func (o *ModifyDepotSharedSettingRequest) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *ModifyDepotSharedSettingRequest) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *ModifyDepotSharedSettingRequest) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


