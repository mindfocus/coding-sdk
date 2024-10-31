# DescribeSingeMergeRequestNotesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径 | 
**MergeId** | **int32** | 合并请求的Iid | 

## Methods

### NewDescribeSingeMergeRequestNotesRequest

`func NewDescribeSingeMergeRequestNotesRequest(depotPath string, mergeId int32, ) *DescribeSingeMergeRequestNotesRequest`

NewDescribeSingeMergeRequestNotesRequest instantiates a new DescribeSingeMergeRequestNotesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeSingeMergeRequestNotesRequestWithDefaults

`func NewDescribeSingeMergeRequestNotesRequestWithDefaults() *DescribeSingeMergeRequestNotesRequest`

NewDescribeSingeMergeRequestNotesRequestWithDefaults instantiates a new DescribeSingeMergeRequestNotesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *DescribeSingeMergeRequestNotesRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeSingeMergeRequestNotesRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeSingeMergeRequestNotesRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetMergeId

`func (o *DescribeSingeMergeRequestNotesRequest) GetMergeId() int32`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *DescribeSingeMergeRequestNotesRequest) GetMergeIdOk() (*int32, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *DescribeSingeMergeRequestNotesRequest) SetMergeId(v int32)`

SetMergeId sets MergeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


