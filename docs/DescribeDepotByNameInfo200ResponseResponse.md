# DescribeDepotByNameInfo200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Depot** | Pointer to [**DepotInfo**](DepotInfo.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeDepotByNameInfo200ResponseResponse

`func NewDescribeDepotByNameInfo200ResponseResponse() *DescribeDepotByNameInfo200ResponseResponse`

NewDescribeDepotByNameInfo200ResponseResponse instantiates a new DescribeDepotByNameInfo200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDepotByNameInfo200ResponseResponseWithDefaults

`func NewDescribeDepotByNameInfo200ResponseResponseWithDefaults() *DescribeDepotByNameInfo200ResponseResponse`

NewDescribeDepotByNameInfo200ResponseResponseWithDefaults instantiates a new DescribeDepotByNameInfo200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepot

`func (o *DescribeDepotByNameInfo200ResponseResponse) GetDepot() DepotInfo`

GetDepot returns the Depot field if non-nil, zero value otherwise.

### GetDepotOk

`func (o *DescribeDepotByNameInfo200ResponseResponse) GetDepotOk() (*DepotInfo, bool)`

GetDepotOk returns a tuple with the Depot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepot

`func (o *DescribeDepotByNameInfo200ResponseResponse) SetDepot(v DepotInfo)`

SetDepot sets Depot field to given value.

### HasDepot

`func (o *DescribeDepotByNameInfo200ResponseResponse) HasDepot() bool`

HasDepot returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeDepotByNameInfo200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeDepotByNameInfo200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeDepotByNameInfo200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeDepotByNameInfo200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


