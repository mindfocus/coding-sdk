# DepotData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Depots** | Pointer to [**[]DepotInfo**](DepotInfo.md) | 仓库信息列表 | [optional] 
**Page** | Pointer to [**PageInfo**](PageInfo.md) |  | [optional] 

## Methods

### NewDepotData

`func NewDepotData() *DepotData`

NewDepotData instantiates a new DepotData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepotDataWithDefaults

`func NewDepotDataWithDefaults() *DepotData`

NewDepotDataWithDefaults instantiates a new DepotData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepots

`func (o *DepotData) GetDepots() []DepotInfo`

GetDepots returns the Depots field if non-nil, zero value otherwise.

### GetDepotsOk

`func (o *DepotData) GetDepotsOk() (*[]DepotInfo, bool)`

GetDepotsOk returns a tuple with the Depots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepots

`func (o *DepotData) SetDepots(v []DepotInfo)`

SetDepots sets Depots field to given value.

### HasDepots

`func (o *DepotData) HasDepots() bool`

HasDepots returns a boolean if a field has been set.

### GetPage

`func (o *DepotData) GetPage() PageInfo`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DepotData) GetPageOk() (*PageInfo, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DepotData) SetPage(v PageInfo)`

SetPage sets Page field to given value.

### HasPage

`func (o *DepotData) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


