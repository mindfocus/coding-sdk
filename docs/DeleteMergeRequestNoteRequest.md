# DeleteMergeRequestNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径 | 
**LineNoteId** | **int32** | 行评论Id | 
**MergeId** | **int32** | 合并请求的Iid | 

## Methods

### NewDeleteMergeRequestNoteRequest

`func NewDeleteMergeRequestNoteRequest(depotPath string, lineNoteId int32, mergeId int32, ) *DeleteMergeRequestNoteRequest`

NewDeleteMergeRequestNoteRequest instantiates a new DeleteMergeRequestNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteMergeRequestNoteRequestWithDefaults

`func NewDeleteMergeRequestNoteRequestWithDefaults() *DeleteMergeRequestNoteRequest`

NewDeleteMergeRequestNoteRequestWithDefaults instantiates a new DeleteMergeRequestNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *DeleteMergeRequestNoteRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DeleteMergeRequestNoteRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DeleteMergeRequestNoteRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetLineNoteId

`func (o *DeleteMergeRequestNoteRequest) GetLineNoteId() int32`

GetLineNoteId returns the LineNoteId field if non-nil, zero value otherwise.

### GetLineNoteIdOk

`func (o *DeleteMergeRequestNoteRequest) GetLineNoteIdOk() (*int32, bool)`

GetLineNoteIdOk returns a tuple with the LineNoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNoteId

`func (o *DeleteMergeRequestNoteRequest) SetLineNoteId(v int32)`

SetLineNoteId sets LineNoteId field to given value.


### GetMergeId

`func (o *DeleteMergeRequestNoteRequest) GetMergeId() int32`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *DeleteMergeRequestNoteRequest) GetMergeIdOk() (*int32, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *DeleteMergeRequestNoteRequest) SetMergeId(v int32)`

SetMergeId sets MergeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


