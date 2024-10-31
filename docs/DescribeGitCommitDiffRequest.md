# DescribeGitCommitDiffRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**Path** | Pointer to **string** | 查询指定文件时填写 | [optional] 
**Sha** | **string** | 提交id | 

## Methods

### NewDescribeGitCommitDiffRequest

`func NewDescribeGitCommitDiffRequest(depotId int64, sha string, ) *DescribeGitCommitDiffRequest`

NewDescribeGitCommitDiffRequest instantiates a new DescribeGitCommitDiffRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitCommitDiffRequestWithDefaults

`func NewDescribeGitCommitDiffRequestWithDefaults() *DescribeGitCommitDiffRequest`

NewDescribeGitCommitDiffRequestWithDefaults instantiates a new DescribeGitCommitDiffRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeGitCommitDiffRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitCommitDiffRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitCommitDiffRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitCommitDiffRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitCommitDiffRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitCommitDiffRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitCommitDiffRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetPath

`func (o *DescribeGitCommitDiffRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DescribeGitCommitDiffRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DescribeGitCommitDiffRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DescribeGitCommitDiffRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSha

`func (o *DescribeGitCommitDiffRequest) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *DescribeGitCommitDiffRequest) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *DescribeGitCommitDiffRequest) SetSha(v string)`

SetSha sets Sha field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


