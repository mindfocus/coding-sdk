# DescribeDepotSpecDetailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | Pointer to **string** | 仓库路径（和仓库规范名字二选一选填） | [optional] 
**DepotSpecName** | Pointer to **string** | 仓库规范名字（和仓库路径二选一选填） | [optional] 

## Methods

### NewDescribeDepotSpecDetailRequest

`func NewDescribeDepotSpecDetailRequest() *DescribeDepotSpecDetailRequest`

NewDescribeDepotSpecDetailRequest instantiates a new DescribeDepotSpecDetailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDepotSpecDetailRequestWithDefaults

`func NewDescribeDepotSpecDetailRequestWithDefaults() *DescribeDepotSpecDetailRequest`

NewDescribeDepotSpecDetailRequestWithDefaults instantiates a new DescribeDepotSpecDetailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *DescribeDepotSpecDetailRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeDepotSpecDetailRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeDepotSpecDetailRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeDepotSpecDetailRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetDepotSpecName

`func (o *DescribeDepotSpecDetailRequest) GetDepotSpecName() string`

GetDepotSpecName returns the DepotSpecName field if non-nil, zero value otherwise.

### GetDepotSpecNameOk

`func (o *DescribeDepotSpecDetailRequest) GetDepotSpecNameOk() (*string, bool)`

GetDepotSpecNameOk returns a tuple with the DepotSpecName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotSpecName

`func (o *DescribeDepotSpecDetailRequest) SetDepotSpecName(v string)`

SetDepotSpecName sets DepotSpecName field to given value.

### HasDepotSpecName

`func (o *DescribeDepotSpecDetailRequest) HasDepotSpecName() bool`

HasDepotSpecName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


