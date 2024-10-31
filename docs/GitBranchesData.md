# GitBranchesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branches** | Pointer to [**[]GitBranchInfo**](GitBranchInfo.md) | 分支列表 | [optional] 
**TotalCount** | Pointer to **int64** | Branches的总条数 | [optional] 

## Methods

### NewGitBranchesData

`func NewGitBranchesData() *GitBranchesData`

NewGitBranchesData instantiates a new GitBranchesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitBranchesDataWithDefaults

`func NewGitBranchesDataWithDefaults() *GitBranchesData`

NewGitBranchesDataWithDefaults instantiates a new GitBranchesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranches

`func (o *GitBranchesData) GetBranches() []GitBranchInfo`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *GitBranchesData) GetBranchesOk() (*[]GitBranchInfo, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *GitBranchesData) SetBranches(v []GitBranchInfo)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *GitBranchesData) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### GetTotalCount

`func (o *GitBranchesData) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GitBranchesData) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GitBranchesData) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GitBranchesData) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


