# ModifyDepotPushSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径 | 
**Param** | [**DepotPushSetting**](DepotPushSetting.md) |  | 

## Methods

### NewModifyDepotPushSettingRequest

`func NewModifyDepotPushSettingRequest(depotPath string, param DepotPushSetting, ) *ModifyDepotPushSettingRequest`

NewModifyDepotPushSettingRequest instantiates a new ModifyDepotPushSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyDepotPushSettingRequestWithDefaults

`func NewModifyDepotPushSettingRequestWithDefaults() *ModifyDepotPushSettingRequest`

NewModifyDepotPushSettingRequestWithDefaults instantiates a new ModifyDepotPushSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *ModifyDepotPushSettingRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyDepotPushSettingRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyDepotPushSettingRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetParam

`func (o *ModifyDepotPushSettingRequest) GetParam() DepotPushSetting`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *ModifyDepotPushSettingRequest) GetParamOk() (*DepotPushSetting, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *ModifyDepotPushSettingRequest) SetParam(v DepotPushSetting)`

SetParam sets Param field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


