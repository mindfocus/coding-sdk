# DescribeGitMergeRequestDiffDetailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**MergeId** | **int64** | 合并请求id | 
**Path** | **string** | 文件路径 | 

## Methods

### NewDescribeGitMergeRequestDiffDetailRequest

`func NewDescribeGitMergeRequestDiffDetailRequest(depotId int64, mergeId int64, path string, ) *DescribeGitMergeRequestDiffDetailRequest`

NewDescribeGitMergeRequestDiffDetailRequest instantiates a new DescribeGitMergeRequestDiffDetailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitMergeRequestDiffDetailRequestWithDefaults

`func NewDescribeGitMergeRequestDiffDetailRequestWithDefaults() *DescribeGitMergeRequestDiffDetailRequest`

NewDescribeGitMergeRequestDiffDetailRequestWithDefaults instantiates a new DescribeGitMergeRequestDiffDetailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeGitMergeRequestDiffDetailRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitMergeRequestDiffDetailRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitMergeRequestDiffDetailRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitMergeRequestDiffDetailRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitMergeRequestDiffDetailRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitMergeRequestDiffDetailRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitMergeRequestDiffDetailRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetMergeId

`func (o *DescribeGitMergeRequestDiffDetailRequest) GetMergeId() int64`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *DescribeGitMergeRequestDiffDetailRequest) GetMergeIdOk() (*int64, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *DescribeGitMergeRequestDiffDetailRequest) SetMergeId(v int64)`

SetMergeId sets MergeId field to given value.


### GetPath

`func (o *DescribeGitMergeRequestDiffDetailRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DescribeGitMergeRequestDiffDetailRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DescribeGitMergeRequestDiffDetailRequest) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


