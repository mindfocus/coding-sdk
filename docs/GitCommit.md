# GitCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorEmail** | Pointer to **string** | 作者邮箱 | [optional] [default to ""]
**AuthorName** | Pointer to **string** | 作者姓名 | [optional] [default to ""]
**CommitDate** | Pointer to **int64** | 提交日期 | [optional] 
**Committer** | Pointer to [**Committer**](Committer.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** | 提交日期 | [optional] 
**Parents** | Pointer to **[]string** | 父提交 | [optional] 
**Sha** | Pointer to **string** | 提交 ID | [optional] [default to ""]
**ShortMessage** | Pointer to **string** | 提交信息 | [optional] [default to ""]

## Methods

### NewGitCommit

`func NewGitCommit() *GitCommit`

NewGitCommit instantiates a new GitCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCommitWithDefaults

`func NewGitCommitWithDefaults() *GitCommit`

NewGitCommitWithDefaults instantiates a new GitCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorEmail

`func (o *GitCommit) GetAuthorEmail() string`

GetAuthorEmail returns the AuthorEmail field if non-nil, zero value otherwise.

### GetAuthorEmailOk

`func (o *GitCommit) GetAuthorEmailOk() (*string, bool)`

GetAuthorEmailOk returns a tuple with the AuthorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorEmail

`func (o *GitCommit) SetAuthorEmail(v string)`

SetAuthorEmail sets AuthorEmail field to given value.

### HasAuthorEmail

`func (o *GitCommit) HasAuthorEmail() bool`

HasAuthorEmail returns a boolean if a field has been set.

### GetAuthorName

`func (o *GitCommit) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *GitCommit) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *GitCommit) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *GitCommit) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### GetCommitDate

`func (o *GitCommit) GetCommitDate() int64`

GetCommitDate returns the CommitDate field if non-nil, zero value otherwise.

### GetCommitDateOk

`func (o *GitCommit) GetCommitDateOk() (*int64, bool)`

GetCommitDateOk returns a tuple with the CommitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitDate

`func (o *GitCommit) SetCommitDate(v int64)`

SetCommitDate sets CommitDate field to given value.

### HasCommitDate

`func (o *GitCommit) HasCommitDate() bool`

HasCommitDate returns a boolean if a field has been set.

### GetCommitter

`func (o *GitCommit) GetCommitter() Committer`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *GitCommit) GetCommitterOk() (*Committer, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *GitCommit) SetCommitter(v Committer)`

SetCommitter sets Committer field to given value.

### HasCommitter

`func (o *GitCommit) HasCommitter() bool`

HasCommitter returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GitCommit) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GitCommit) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GitCommit) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GitCommit) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetParents

`func (o *GitCommit) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *GitCommit) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *GitCommit) SetParents(v []string)`

SetParents sets Parents field to given value.

### HasParents

`func (o *GitCommit) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetSha

`func (o *GitCommit) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitCommit) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitCommit) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *GitCommit) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetShortMessage

`func (o *GitCommit) GetShortMessage() string`

GetShortMessage returns the ShortMessage field if non-nil, zero value otherwise.

### GetShortMessageOk

`func (o *GitCommit) GetShortMessageOk() (*string, bool)`

GetShortMessageOk returns a tuple with the ShortMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortMessage

`func (o *GitCommit) SetShortMessage(v string)`

SetShortMessage sets ShortMessage field to given value.

### HasShortMessage

`func (o *GitCommit) HasShortMessage() bool`

HasShortMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


