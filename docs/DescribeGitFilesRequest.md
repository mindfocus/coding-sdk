# DescribeGitFilesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库Id | 
**DepotPath** | Pointer to **string** | 仓库路径，与仓库ID二选一 | [optional] 
**Path** | Pointer to **string** | 路径，不传即为根目录 | [optional] 
**Ref** | **string** | 分支名 | 

## Methods

### NewDescribeGitFilesRequest

`func NewDescribeGitFilesRequest(depotId int64, ref string, ) *DescribeGitFilesRequest`

NewDescribeGitFilesRequest instantiates a new DescribeGitFilesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitFilesRequestWithDefaults

`func NewDescribeGitFilesRequestWithDefaults() *DescribeGitFilesRequest`

NewDescribeGitFilesRequestWithDefaults instantiates a new DescribeGitFilesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeGitFilesRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitFilesRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitFilesRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitFilesRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitFilesRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitFilesRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitFilesRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetPath

`func (o *DescribeGitFilesRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DescribeGitFilesRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DescribeGitFilesRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DescribeGitFilesRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRef

`func (o *DescribeGitFilesRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DescribeGitFilesRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DescribeGitFilesRequest) SetRef(v string)`

SetRef sets Ref field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


