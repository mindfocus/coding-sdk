# DescribeGitMergeRequestParticipantsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库 Id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**MergeId** | **int64** | 合并请求 IId | 

## Methods

### NewDescribeGitMergeRequestParticipantsRequest

`func NewDescribeGitMergeRequestParticipantsRequest(depotId int64, mergeId int64, ) *DescribeGitMergeRequestParticipantsRequest`

NewDescribeGitMergeRequestParticipantsRequest instantiates a new DescribeGitMergeRequestParticipantsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitMergeRequestParticipantsRequestWithDefaults

`func NewDescribeGitMergeRequestParticipantsRequestWithDefaults() *DescribeGitMergeRequestParticipantsRequest`

NewDescribeGitMergeRequestParticipantsRequestWithDefaults instantiates a new DescribeGitMergeRequestParticipantsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeGitMergeRequestParticipantsRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitMergeRequestParticipantsRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitMergeRequestParticipantsRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitMergeRequestParticipantsRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitMergeRequestParticipantsRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitMergeRequestParticipantsRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitMergeRequestParticipantsRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetMergeId

`func (o *DescribeGitMergeRequestParticipantsRequest) GetMergeId() int64`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *DescribeGitMergeRequestParticipantsRequest) GetMergeIdOk() (*int64, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *DescribeGitMergeRequestParticipantsRequest) SetMergeId(v int64)`

SetMergeId sets MergeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


