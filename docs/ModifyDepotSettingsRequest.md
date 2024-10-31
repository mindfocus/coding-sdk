# ModifyDepotSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径， | 
**DepotStatusCheckRequired** | **bool** | 是否开启状态检查 | 

## Methods

### NewModifyDepotSettingsRequest

`func NewModifyDepotSettingsRequest(depotPath string, depotStatusCheckRequired bool, ) *ModifyDepotSettingsRequest`

NewModifyDepotSettingsRequest instantiates a new ModifyDepotSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyDepotSettingsRequestWithDefaults

`func NewModifyDepotSettingsRequestWithDefaults() *ModifyDepotSettingsRequest`

NewModifyDepotSettingsRequestWithDefaults instantiates a new ModifyDepotSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *ModifyDepotSettingsRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyDepotSettingsRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyDepotSettingsRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetDepotStatusCheckRequired

`func (o *ModifyDepotSettingsRequest) GetDepotStatusCheckRequired() bool`

GetDepotStatusCheckRequired returns the DepotStatusCheckRequired field if non-nil, zero value otherwise.

### GetDepotStatusCheckRequiredOk

`func (o *ModifyDepotSettingsRequest) GetDepotStatusCheckRequiredOk() (*bool, bool)`

GetDepotStatusCheckRequiredOk returns a tuple with the DepotStatusCheckRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotStatusCheckRequired

`func (o *ModifyDepotSettingsRequest) SetDepotStatusCheckRequired(v bool)`

SetDepotStatusCheckRequired sets DepotStatusCheckRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


