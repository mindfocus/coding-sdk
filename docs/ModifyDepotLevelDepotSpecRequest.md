# ModifyDepotLevelDepotSpecRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径 | 
**Param** | [**DepotSpecDepotLevelParam**](DepotSpecDepotLevelParam.md) |  | 

## Methods

### NewModifyDepotLevelDepotSpecRequest

`func NewModifyDepotLevelDepotSpecRequest(depotPath string, param DepotSpecDepotLevelParam, ) *ModifyDepotLevelDepotSpecRequest`

NewModifyDepotLevelDepotSpecRequest instantiates a new ModifyDepotLevelDepotSpecRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyDepotLevelDepotSpecRequestWithDefaults

`func NewModifyDepotLevelDepotSpecRequestWithDefaults() *ModifyDepotLevelDepotSpecRequest`

NewModifyDepotLevelDepotSpecRequestWithDefaults instantiates a new ModifyDepotLevelDepotSpecRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *ModifyDepotLevelDepotSpecRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyDepotLevelDepotSpecRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyDepotLevelDepotSpecRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetParam

`func (o *ModifyDepotLevelDepotSpecRequest) GetParam() DepotSpecDepotLevelParam`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *ModifyDepotLevelDepotSpecRequest) GetParamOk() (*DepotSpecDepotLevelParam, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *ModifyDepotLevelDepotSpecRequest) SetParam(v DepotSpecDepotLevelParam)`

SetParam sets Param field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


