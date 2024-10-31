# CodingCIPersonalExternalDepotData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotList** | Pointer to [**[]CodingCIPersonalExternalDepot**](CodingCIPersonalExternalDepot.md) | 仓库列表 | [optional] 
**IsBound** | Pointer to **bool** | 仓库类型是否被授权，如 Github 是否被授权 | [optional] [default to false]
**OwnerName** | Pointer to **NullableString** | 所有者用户名 | [optional] [default to ""]

## Methods

### NewCodingCIPersonalExternalDepotData

`func NewCodingCIPersonalExternalDepotData() *CodingCIPersonalExternalDepotData`

NewCodingCIPersonalExternalDepotData instantiates a new CodingCIPersonalExternalDepotData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodingCIPersonalExternalDepotDataWithDefaults

`func NewCodingCIPersonalExternalDepotDataWithDefaults() *CodingCIPersonalExternalDepotData`

NewCodingCIPersonalExternalDepotDataWithDefaults instantiates a new CodingCIPersonalExternalDepotData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotList

`func (o *CodingCIPersonalExternalDepotData) GetDepotList() []CodingCIPersonalExternalDepot`

GetDepotList returns the DepotList field if non-nil, zero value otherwise.

### GetDepotListOk

`func (o *CodingCIPersonalExternalDepotData) GetDepotListOk() (*[]CodingCIPersonalExternalDepot, bool)`

GetDepotListOk returns a tuple with the DepotList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotList

`func (o *CodingCIPersonalExternalDepotData) SetDepotList(v []CodingCIPersonalExternalDepot)`

SetDepotList sets DepotList field to given value.

### HasDepotList

`func (o *CodingCIPersonalExternalDepotData) HasDepotList() bool`

HasDepotList returns a boolean if a field has been set.

### GetIsBound

`func (o *CodingCIPersonalExternalDepotData) GetIsBound() bool`

GetIsBound returns the IsBound field if non-nil, zero value otherwise.

### GetIsBoundOk

`func (o *CodingCIPersonalExternalDepotData) GetIsBoundOk() (*bool, bool)`

GetIsBoundOk returns a tuple with the IsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBound

`func (o *CodingCIPersonalExternalDepotData) SetIsBound(v bool)`

SetIsBound sets IsBound field to given value.

### HasIsBound

`func (o *CodingCIPersonalExternalDepotData) HasIsBound() bool`

HasIsBound returns a boolean if a field has been set.

### GetOwnerName

`func (o *CodingCIPersonalExternalDepotData) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *CodingCIPersonalExternalDepotData) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *CodingCIPersonalExternalDepotData) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *CodingCIPersonalExternalDepotData) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### SetOwnerNameNil

`func (o *CodingCIPersonalExternalDepotData) SetOwnerNameNil(b bool)`

 SetOwnerNameNil sets the value for OwnerName to be an explicit nil

### UnsetOwnerName
`func (o *CodingCIPersonalExternalDepotData) UnsetOwnerName()`

UnsetOwnerName ensures that no value is present for OwnerName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


