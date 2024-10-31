# DescribeDepotSpecs200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotSpecs** | Pointer to [**[]DepotSpec**](DepotSpec.md) | 仓库规范列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeDepotSpecs200ResponseResponse

`func NewDescribeDepotSpecs200ResponseResponse() *DescribeDepotSpecs200ResponseResponse`

NewDescribeDepotSpecs200ResponseResponse instantiates a new DescribeDepotSpecs200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDepotSpecs200ResponseResponseWithDefaults

`func NewDescribeDepotSpecs200ResponseResponseWithDefaults() *DescribeDepotSpecs200ResponseResponse`

NewDescribeDepotSpecs200ResponseResponseWithDefaults instantiates a new DescribeDepotSpecs200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotSpecs

`func (o *DescribeDepotSpecs200ResponseResponse) GetDepotSpecs() []DepotSpec`

GetDepotSpecs returns the DepotSpecs field if non-nil, zero value otherwise.

### GetDepotSpecsOk

`func (o *DescribeDepotSpecs200ResponseResponse) GetDepotSpecsOk() (*[]DepotSpec, bool)`

GetDepotSpecsOk returns a tuple with the DepotSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotSpecs

`func (o *DescribeDepotSpecs200ResponseResponse) SetDepotSpecs(v []DepotSpec)`

SetDepotSpecs sets DepotSpecs field to given value.

### HasDepotSpecs

`func (o *DescribeDepotSpecs200ResponseResponse) HasDepotSpecs() bool`

HasDepotSpecs returns a boolean if a field has been set.

### SetDepotSpecsNil

`func (o *DescribeDepotSpecs200ResponseResponse) SetDepotSpecsNil(b bool)`

 SetDepotSpecsNil sets the value for DepotSpecs to be an explicit nil

### UnsetDepotSpecs
`func (o *DescribeDepotSpecs200ResponseResponse) UnsetDepotSpecs()`

UnsetDepotSpecs ensures that no value is present for DepotSpecs, not even an explicit nil
### GetRequestId

`func (o *DescribeDepotSpecs200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeDepotSpecs200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeDepotSpecs200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeDepotSpecs200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


