# DescribeMergeReqCommitsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 与仓库路径二选一即可 | [optional] 
**MergeId** | **int64** | 合并请求id | 

## Methods

### NewDescribeMergeReqCommitsRequest

`func NewDescribeMergeReqCommitsRequest(depotId int64, mergeId int64, ) *DescribeMergeReqCommitsRequest`

NewDescribeMergeReqCommitsRequest instantiates a new DescribeMergeReqCommitsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeMergeReqCommitsRequestWithDefaults

`func NewDescribeMergeReqCommitsRequestWithDefaults() *DescribeMergeReqCommitsRequest`

NewDescribeMergeReqCommitsRequestWithDefaults instantiates a new DescribeMergeReqCommitsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeMergeReqCommitsRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeMergeReqCommitsRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeMergeReqCommitsRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeMergeReqCommitsRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeMergeReqCommitsRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeMergeReqCommitsRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeMergeReqCommitsRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetMergeId

`func (o *DescribeMergeReqCommitsRequest) GetMergeId() int64`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *DescribeMergeReqCommitsRequest) GetMergeIdOk() (*int64, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *DescribeMergeReqCommitsRequest) SetMergeId(v int64)`

SetMergeId sets MergeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


