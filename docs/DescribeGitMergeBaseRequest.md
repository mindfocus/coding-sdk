# DescribeGitMergeBaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库 ID | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可  | [optional] 
**DestRef** | **string** | 分支2 | 
**SrcRef** | **string** | 分支1 | 

## Methods

### NewDescribeGitMergeBaseRequest

`func NewDescribeGitMergeBaseRequest(depotId int64, destRef string, srcRef string, ) *DescribeGitMergeBaseRequest`

NewDescribeGitMergeBaseRequest instantiates a new DescribeGitMergeBaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitMergeBaseRequestWithDefaults

`func NewDescribeGitMergeBaseRequestWithDefaults() *DescribeGitMergeBaseRequest`

NewDescribeGitMergeBaseRequestWithDefaults instantiates a new DescribeGitMergeBaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeGitMergeBaseRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitMergeBaseRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitMergeBaseRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitMergeBaseRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitMergeBaseRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitMergeBaseRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitMergeBaseRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetDestRef

`func (o *DescribeGitMergeBaseRequest) GetDestRef() string`

GetDestRef returns the DestRef field if non-nil, zero value otherwise.

### GetDestRefOk

`func (o *DescribeGitMergeBaseRequest) GetDestRefOk() (*string, bool)`

GetDestRefOk returns a tuple with the DestRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestRef

`func (o *DescribeGitMergeBaseRequest) SetDestRef(v string)`

SetDestRef sets DestRef field to given value.


### GetSrcRef

`func (o *DescribeGitMergeBaseRequest) GetSrcRef() string`

GetSrcRef returns the SrcRef field if non-nil, zero value otherwise.

### GetSrcRefOk

`func (o *DescribeGitMergeBaseRequest) GetSrcRefOk() (*string, bool)`

GetSrcRefOk returns a tuple with the SrcRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcRef

`func (o *DescribeGitMergeBaseRequest) SetSrcRef(v string)`

SetSrcRef sets SrcRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


