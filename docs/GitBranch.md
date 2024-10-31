# GitBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | Pointer to **string** | 分支名称 | [optional] [default to ""]
**IsDefaultBranch** | Pointer to **bool** | 是否为默认分支 | [optional] [default to false]
**IsProtected** | Pointer to **bool** | 是否为保护分支 | [optional] [default to false]
**LastCommit** | Pointer to [**GitCommit**](GitCommit.md) |  | [optional] 
**Sha** | Pointer to **string** | 分支的sha值 | [optional] [default to ""]
**Content** | Pointer to **string** | 分支的备注信息 | [optional] [default to ""]

## Methods

### NewGitBranch

`func NewGitBranch() *GitBranch`

NewGitBranch instantiates a new GitBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitBranchWithDefaults

`func NewGitBranchWithDefaults() *GitBranch`

NewGitBranchWithDefaults instantiates a new GitBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchName

`func (o *GitBranch) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *GitBranch) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *GitBranch) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *GitBranch) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetIsDefaultBranch

`func (o *GitBranch) GetIsDefaultBranch() bool`

GetIsDefaultBranch returns the IsDefaultBranch field if non-nil, zero value otherwise.

### GetIsDefaultBranchOk

`func (o *GitBranch) GetIsDefaultBranchOk() (*bool, bool)`

GetIsDefaultBranchOk returns a tuple with the IsDefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultBranch

`func (o *GitBranch) SetIsDefaultBranch(v bool)`

SetIsDefaultBranch sets IsDefaultBranch field to given value.

### HasIsDefaultBranch

`func (o *GitBranch) HasIsDefaultBranch() bool`

HasIsDefaultBranch returns a boolean if a field has been set.

### GetIsProtected

`func (o *GitBranch) GetIsProtected() bool`

GetIsProtected returns the IsProtected field if non-nil, zero value otherwise.

### GetIsProtectedOk

`func (o *GitBranch) GetIsProtectedOk() (*bool, bool)`

GetIsProtectedOk returns a tuple with the IsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProtected

`func (o *GitBranch) SetIsProtected(v bool)`

SetIsProtected sets IsProtected field to given value.

### HasIsProtected

`func (o *GitBranch) HasIsProtected() bool`

HasIsProtected returns a boolean if a field has been set.

### GetLastCommit

`func (o *GitBranch) GetLastCommit() GitCommit`

GetLastCommit returns the LastCommit field if non-nil, zero value otherwise.

### GetLastCommitOk

`func (o *GitBranch) GetLastCommitOk() (*GitCommit, bool)`

GetLastCommitOk returns a tuple with the LastCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommit

`func (o *GitBranch) SetLastCommit(v GitCommit)`

SetLastCommit sets LastCommit field to given value.

### HasLastCommit

`func (o *GitBranch) HasLastCommit() bool`

HasLastCommit returns a boolean if a field has been set.

### GetSha

`func (o *GitBranch) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitBranch) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitBranch) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *GitBranch) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetContent

`func (o *GitBranch) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GitBranch) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GitBranch) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GitBranch) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


