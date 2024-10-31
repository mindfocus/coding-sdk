# ModifyDepotQuotaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径 | 
**DepotQuotaSize** | **string** | 仓库容量 | 

## Methods

### NewModifyDepotQuotaRequest

`func NewModifyDepotQuotaRequest(depotPath string, depotQuotaSize string, ) *ModifyDepotQuotaRequest`

NewModifyDepotQuotaRequest instantiates a new ModifyDepotQuotaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyDepotQuotaRequestWithDefaults

`func NewModifyDepotQuotaRequestWithDefaults() *ModifyDepotQuotaRequest`

NewModifyDepotQuotaRequestWithDefaults instantiates a new ModifyDepotQuotaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *ModifyDepotQuotaRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyDepotQuotaRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyDepotQuotaRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetDepotQuotaSize

`func (o *ModifyDepotQuotaRequest) GetDepotQuotaSize() string`

GetDepotQuotaSize returns the DepotQuotaSize field if non-nil, zero value otherwise.

### GetDepotQuotaSizeOk

`func (o *ModifyDepotQuotaRequest) GetDepotQuotaSizeOk() (*string, bool)`

GetDepotQuotaSizeOk returns a tuple with the DepotQuotaSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotQuotaSize

`func (o *ModifyDepotQuotaRequest) SetDepotQuotaSize(v string)`

SetDepotQuotaSize sets DepotQuotaSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


