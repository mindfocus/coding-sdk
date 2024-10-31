# DescribeGitBranchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | **string** | 分支名称 | 
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，与DepotId选其一就可以 | [optional] 

## Methods

### NewDescribeGitBranchRequest

`func NewDescribeGitBranchRequest(branchName string, depotId int64, ) *DescribeGitBranchRequest`

NewDescribeGitBranchRequest instantiates a new DescribeGitBranchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitBranchRequestWithDefaults

`func NewDescribeGitBranchRequestWithDefaults() *DescribeGitBranchRequest`

NewDescribeGitBranchRequestWithDefaults instantiates a new DescribeGitBranchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchName

`func (o *DescribeGitBranchRequest) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *DescribeGitBranchRequest) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *DescribeGitBranchRequest) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.


### GetDepotId

`func (o *DescribeGitBranchRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitBranchRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitBranchRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitBranchRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitBranchRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitBranchRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitBranchRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


