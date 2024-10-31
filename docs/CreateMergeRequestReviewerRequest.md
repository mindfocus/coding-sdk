# CreateMergeRequestReviewerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**MergeId** | **int64** | 合并请求id | 
**ReviewerGlobalKey** | **string** | 评审者 | 

## Methods

### NewCreateMergeRequestReviewerRequest

`func NewCreateMergeRequestReviewerRequest(depotId int64, mergeId int64, reviewerGlobalKey string, ) *CreateMergeRequestReviewerRequest`

NewCreateMergeRequestReviewerRequest instantiates a new CreateMergeRequestReviewerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMergeRequestReviewerRequestWithDefaults

`func NewCreateMergeRequestReviewerRequestWithDefaults() *CreateMergeRequestReviewerRequest`

NewCreateMergeRequestReviewerRequestWithDefaults instantiates a new CreateMergeRequestReviewerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *CreateMergeRequestReviewerRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CreateMergeRequestReviewerRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CreateMergeRequestReviewerRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *CreateMergeRequestReviewerRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateMergeRequestReviewerRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateMergeRequestReviewerRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *CreateMergeRequestReviewerRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetMergeId

`func (o *CreateMergeRequestReviewerRequest) GetMergeId() int64`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *CreateMergeRequestReviewerRequest) GetMergeIdOk() (*int64, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *CreateMergeRequestReviewerRequest) SetMergeId(v int64)`

SetMergeId sets MergeId field to given value.


### GetReviewerGlobalKey

`func (o *CreateMergeRequestReviewerRequest) GetReviewerGlobalKey() string`

GetReviewerGlobalKey returns the ReviewerGlobalKey field if non-nil, zero value otherwise.

### GetReviewerGlobalKeyOk

`func (o *CreateMergeRequestReviewerRequest) GetReviewerGlobalKeyOk() (*string, bool)`

GetReviewerGlobalKeyOk returns a tuple with the ReviewerGlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerGlobalKey

`func (o *CreateMergeRequestReviewerRequest) SetReviewerGlobalKey(v string)`

SetReviewerGlobalKey sets ReviewerGlobalKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


