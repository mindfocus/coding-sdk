# DescribeProtectedBranchesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | **string** | 仓库路径 | 

## Methods

### NewDescribeProtectedBranchesRequest

`func NewDescribeProtectedBranchesRequest(depotId int64, depotPath string, ) *DescribeProtectedBranchesRequest`

NewDescribeProtectedBranchesRequest instantiates a new DescribeProtectedBranchesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProtectedBranchesRequestWithDefaults

`func NewDescribeProtectedBranchesRequestWithDefaults() *DescribeProtectedBranchesRequest`

NewDescribeProtectedBranchesRequestWithDefaults instantiates a new DescribeProtectedBranchesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeProtectedBranchesRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeProtectedBranchesRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeProtectedBranchesRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeProtectedBranchesRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeProtectedBranchesRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeProtectedBranchesRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


