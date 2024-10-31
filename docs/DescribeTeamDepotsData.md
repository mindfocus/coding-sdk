# DescribeTeamDepotsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotList** | Pointer to [**[]CodingCITeamDepot**](CodingCITeamDepot.md) | 仓库列表 | [optional] 
**IsBound** | Pointer to **bool** | 仓库类型是否被授权，如 Github 是否被授权 | [optional] [default to false]

## Methods

### NewDescribeTeamDepotsData

`func NewDescribeTeamDepotsData() *DescribeTeamDepotsData`

NewDescribeTeamDepotsData instantiates a new DescribeTeamDepotsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeTeamDepotsDataWithDefaults

`func NewDescribeTeamDepotsDataWithDefaults() *DescribeTeamDepotsData`

NewDescribeTeamDepotsDataWithDefaults instantiates a new DescribeTeamDepotsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotList

`func (o *DescribeTeamDepotsData) GetDepotList() []CodingCITeamDepot`

GetDepotList returns the DepotList field if non-nil, zero value otherwise.

### GetDepotListOk

`func (o *DescribeTeamDepotsData) GetDepotListOk() (*[]CodingCITeamDepot, bool)`

GetDepotListOk returns a tuple with the DepotList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotList

`func (o *DescribeTeamDepotsData) SetDepotList(v []CodingCITeamDepot)`

SetDepotList sets DepotList field to given value.

### HasDepotList

`func (o *DescribeTeamDepotsData) HasDepotList() bool`

HasDepotList returns a boolean if a field has been set.

### GetIsBound

`func (o *DescribeTeamDepotsData) GetIsBound() bool`

GetIsBound returns the IsBound field if non-nil, zero value otherwise.

### GetIsBoundOk

`func (o *DescribeTeamDepotsData) GetIsBoundOk() (*bool, bool)`

GetIsBoundOk returns a tuple with the IsBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBound

`func (o *DescribeTeamDepotsData) SetIsBound(v bool)`

SetIsBound sets IsBound field to given value.

### HasIsBound

`func (o *DescribeTeamDepotsData) HasIsBound() bool`

HasIsBound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


