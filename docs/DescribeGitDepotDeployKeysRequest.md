# DescribeGitDepotDeployKeysRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库 Id | 
**DepotPath** | Pointer to **string** | 仓库路径，与仓库ID二选一 | [optional] 

## Methods

### NewDescribeGitDepotDeployKeysRequest

`func NewDescribeGitDepotDeployKeysRequest(depotId int64, ) *DescribeGitDepotDeployKeysRequest`

NewDescribeGitDepotDeployKeysRequest instantiates a new DescribeGitDepotDeployKeysRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitDepotDeployKeysRequestWithDefaults

`func NewDescribeGitDepotDeployKeysRequestWithDefaults() *DescribeGitDepotDeployKeysRequest`

NewDescribeGitDepotDeployKeysRequestWithDefaults instantiates a new DescribeGitDepotDeployKeysRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeGitDepotDeployKeysRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitDepotDeployKeysRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitDepotDeployKeysRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitDepotDeployKeysRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitDepotDeployKeysRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitDepotDeployKeysRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitDepotDeployKeysRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


