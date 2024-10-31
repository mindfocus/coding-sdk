# DescribeGitFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库Id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**Path** | Pointer to **string** | 文件路径 | [optional] 
**Ref** | **string** | 分支名 | 

## Methods

### NewDescribeGitFileRequest

`func NewDescribeGitFileRequest(depotId int64, ref string, ) *DescribeGitFileRequest`

NewDescribeGitFileRequest instantiates a new DescribeGitFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitFileRequestWithDefaults

`func NewDescribeGitFileRequestWithDefaults() *DescribeGitFileRequest`

NewDescribeGitFileRequestWithDefaults instantiates a new DescribeGitFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeGitFileRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitFileRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitFileRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitFileRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitFileRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitFileRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitFileRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetPath

`func (o *DescribeGitFileRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DescribeGitFileRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DescribeGitFileRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DescribeGitFileRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRef

`func (o *DescribeGitFileRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DescribeGitFileRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DescribeGitFileRequest) SetRef(v string)`

SetRef sets Ref field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


