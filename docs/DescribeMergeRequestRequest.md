# DescribeMergeRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，与DepotId选择其一即可 | [optional] 
**MergeId** | **int64** | 合并请求iid | 

## Methods

### NewDescribeMergeRequestRequest

`func NewDescribeMergeRequestRequest(depotId int64, mergeId int64, ) *DescribeMergeRequestRequest`

NewDescribeMergeRequestRequest instantiates a new DescribeMergeRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeMergeRequestRequestWithDefaults

`func NewDescribeMergeRequestRequestWithDefaults() *DescribeMergeRequestRequest`

NewDescribeMergeRequestRequestWithDefaults instantiates a new DescribeMergeRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeMergeRequestRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeMergeRequestRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeMergeRequestRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeMergeRequestRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeMergeRequestRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeMergeRequestRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeMergeRequestRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetMergeId

`func (o *DescribeMergeRequestRequest) GetMergeId() int64`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *DescribeMergeRequestRequest) GetMergeIdOk() (*int64, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *DescribeMergeRequestRequest) SetMergeId(v int64)`

SetMergeId sets MergeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


