# DescribeGitTreeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库Id | 
**DepotPath** | Pointer to **string** | 仓库路径，与仓库ID二选一 | [optional] 
**IsRecursive** | **bool** | 是否递归 | 
**Path** | Pointer to **string** | 文件路径 | [optional] 
**Ref** | **string** | 分支或标签名，默认 HEAD | 

## Methods

### NewDescribeGitTreeRequest

`func NewDescribeGitTreeRequest(depotId int64, isRecursive bool, ref string, ) *DescribeGitTreeRequest`

NewDescribeGitTreeRequest instantiates a new DescribeGitTreeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitTreeRequestWithDefaults

`func NewDescribeGitTreeRequestWithDefaults() *DescribeGitTreeRequest`

NewDescribeGitTreeRequestWithDefaults instantiates a new DescribeGitTreeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeGitTreeRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitTreeRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitTreeRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitTreeRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitTreeRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitTreeRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitTreeRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetIsRecursive

`func (o *DescribeGitTreeRequest) GetIsRecursive() bool`

GetIsRecursive returns the IsRecursive field if non-nil, zero value otherwise.

### GetIsRecursiveOk

`func (o *DescribeGitTreeRequest) GetIsRecursiveOk() (*bool, bool)`

GetIsRecursiveOk returns a tuple with the IsRecursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecursive

`func (o *DescribeGitTreeRequest) SetIsRecursive(v bool)`

SetIsRecursive sets IsRecursive field to given value.


### GetPath

`func (o *DescribeGitTreeRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DescribeGitTreeRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DescribeGitTreeRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DescribeGitTreeRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRef

`func (o *DescribeGitTreeRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DescribeGitTreeRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DescribeGitTreeRequest) SetRef(v string)`

SetRef sets Ref field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


