# DescribeMergeRequestReviewersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库 Id | 
**MergeId** | **int64** | 合并请求 IId | 

## Methods

### NewDescribeMergeRequestReviewersRequest

`func NewDescribeMergeRequestReviewersRequest(depotId int64, mergeId int64, ) *DescribeMergeRequestReviewersRequest`

NewDescribeMergeRequestReviewersRequest instantiates a new DescribeMergeRequestReviewersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeMergeRequestReviewersRequestWithDefaults

`func NewDescribeMergeRequestReviewersRequestWithDefaults() *DescribeMergeRequestReviewersRequest`

NewDescribeMergeRequestReviewersRequestWithDefaults instantiates a new DescribeMergeRequestReviewersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeMergeRequestReviewersRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeMergeRequestReviewersRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeMergeRequestReviewersRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetMergeId

`func (o *DescribeMergeRequestReviewersRequest) GetMergeId() int64`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *DescribeMergeRequestReviewersRequest) GetMergeIdOk() (*int64, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *DescribeMergeRequestReviewersRequest) SetMergeId(v int64)`

SetMergeId sets MergeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


