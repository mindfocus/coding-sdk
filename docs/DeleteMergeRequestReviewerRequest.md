# DeleteMergeRequestReviewerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库 ID | 
**DepotPath** | Pointer to **string** | 仓库路径，与仓库ID二选一 | [optional] 
**MergeId** | **int64** | iid | 
**ReviewerGlobalKey** | **string** | 评审者的GK | 

## Methods

### NewDeleteMergeRequestReviewerRequest

`func NewDeleteMergeRequestReviewerRequest(depotId int64, mergeId int64, reviewerGlobalKey string, ) *DeleteMergeRequestReviewerRequest`

NewDeleteMergeRequestReviewerRequest instantiates a new DeleteMergeRequestReviewerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteMergeRequestReviewerRequestWithDefaults

`func NewDeleteMergeRequestReviewerRequestWithDefaults() *DeleteMergeRequestReviewerRequest`

NewDeleteMergeRequestReviewerRequestWithDefaults instantiates a new DeleteMergeRequestReviewerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DeleteMergeRequestReviewerRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DeleteMergeRequestReviewerRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DeleteMergeRequestReviewerRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DeleteMergeRequestReviewerRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DeleteMergeRequestReviewerRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DeleteMergeRequestReviewerRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DeleteMergeRequestReviewerRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetMergeId

`func (o *DeleteMergeRequestReviewerRequest) GetMergeId() int64`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *DeleteMergeRequestReviewerRequest) GetMergeIdOk() (*int64, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *DeleteMergeRequestReviewerRequest) SetMergeId(v int64)`

SetMergeId sets MergeId field to given value.


### GetReviewerGlobalKey

`func (o *DeleteMergeRequestReviewerRequest) GetReviewerGlobalKey() string`

GetReviewerGlobalKey returns the ReviewerGlobalKey field if non-nil, zero value otherwise.

### GetReviewerGlobalKeyOk

`func (o *DeleteMergeRequestReviewerRequest) GetReviewerGlobalKeyOk() (*string, bool)`

GetReviewerGlobalKeyOk returns a tuple with the ReviewerGlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGlobalKey

`func (o *DeleteMergeRequestReviewerRequest) SetReviewerGlobalKey(v string)`

SetReviewerGlobalKey sets ReviewerGlobalKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


