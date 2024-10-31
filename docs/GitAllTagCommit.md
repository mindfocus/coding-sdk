# GitAllTagCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorEmail** | Pointer to **NullableString** | 作者邮箱 | [optional] [default to ""]
**AuthorName** | Pointer to **NullableString** | 作者姓名 | [optional] [default to ""]
**CommitDate** | Pointer to **NullableInt64** | 提交时间 | [optional] 
**CommitterEmail** | Pointer to **NullableString** | 提交者邮箱 | [optional] [default to ""]
**CommitterName** | Pointer to **NullableString** | 提交者姓名 | [optional] [default to ""]
**CreatedAt** | Pointer to **NullableInt64** | 创建时间 | [optional] 
**Parents** | Pointer to **[]string** | 父提交 | [optional] 
**Sha** | Pointer to **NullableString** | 提交sha | [optional] [default to ""]
**ShortMessage** | Pointer to **NullableString** | 短描述 | [optional] [default to ""]

## Methods

### NewGitAllTagCommit

`func NewGitAllTagCommit() *GitAllTagCommit`

NewGitAllTagCommit instantiates a new GitAllTagCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitAllTagCommitWithDefaults

`func NewGitAllTagCommitWithDefaults() *GitAllTagCommit`

NewGitAllTagCommitWithDefaults instantiates a new GitAllTagCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorEmail

`func (o *GitAllTagCommit) GetAuthorEmail() string`

GetAuthorEmail returns the AuthorEmail field if non-nil, zero value otherwise.

### GetAuthorEmailOk

`func (o *GitAllTagCommit) GetAuthorEmailOk() (*string, bool)`

GetAuthorEmailOk returns a tuple with the AuthorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorEmail

`func (o *GitAllTagCommit) SetAuthorEmail(v string)`

SetAuthorEmail sets AuthorEmail field to given value.

### HasAuthorEmail

`func (o *GitAllTagCommit) HasAuthorEmail() bool`

HasAuthorEmail returns a boolean if a field has been set.

### SetAuthorEmailNil

`func (o *GitAllTagCommit) SetAuthorEmailNil(b bool)`

 SetAuthorEmailNil sets the value for AuthorEmail to be an explicit nil

### UnsetAuthorEmail
`func (o *GitAllTagCommit) UnsetAuthorEmail()`

UnsetAuthorEmail ensures that no value is present for AuthorEmail, not even an explicit nil
### GetAuthorName

`func (o *GitAllTagCommit) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *GitAllTagCommit) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *GitAllTagCommit) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *GitAllTagCommit) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### SetAuthorNameNil

`func (o *GitAllTagCommit) SetAuthorNameNil(b bool)`

 SetAuthorNameNil sets the value for AuthorName to be an explicit nil

### UnsetAuthorName
`func (o *GitAllTagCommit) UnsetAuthorName()`

UnsetAuthorName ensures that no value is present for AuthorName, not even an explicit nil
### GetCommitDate

`func (o *GitAllTagCommit) GetCommitDate() int64`

GetCommitDate returns the CommitDate field if non-nil, zero value otherwise.

### GetCommitDateOk

`func (o *GitAllTagCommit) GetCommitDateOk() (*int64, bool)`

GetCommitDateOk returns a tuple with the CommitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitDate

`func (o *GitAllTagCommit) SetCommitDate(v int64)`

SetCommitDate sets CommitDate field to given value.

### HasCommitDate

`func (o *GitAllTagCommit) HasCommitDate() bool`

HasCommitDate returns a boolean if a field has been set.

### SetCommitDateNil

`func (o *GitAllTagCommit) SetCommitDateNil(b bool)`

 SetCommitDateNil sets the value for CommitDate to be an explicit nil

### UnsetCommitDate
`func (o *GitAllTagCommit) UnsetCommitDate()`

UnsetCommitDate ensures that no value is present for CommitDate, not even an explicit nil
### GetCommitterEmail

`func (o *GitAllTagCommit) GetCommitterEmail() string`

GetCommitterEmail returns the CommitterEmail field if non-nil, zero value otherwise.

### GetCommitterEmailOk

`func (o *GitAllTagCommit) GetCommitterEmailOk() (*string, bool)`

GetCommitterEmailOk returns a tuple with the CommitterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitterEmail

`func (o *GitAllTagCommit) SetCommitterEmail(v string)`

SetCommitterEmail sets CommitterEmail field to given value.

### HasCommitterEmail

`func (o *GitAllTagCommit) HasCommitterEmail() bool`

HasCommitterEmail returns a boolean if a field has been set.

### SetCommitterEmailNil

`func (o *GitAllTagCommit) SetCommitterEmailNil(b bool)`

 SetCommitterEmailNil sets the value for CommitterEmail to be an explicit nil

### UnsetCommitterEmail
`func (o *GitAllTagCommit) UnsetCommitterEmail()`

UnsetCommitterEmail ensures that no value is present for CommitterEmail, not even an explicit nil
### GetCommitterName

`func (o *GitAllTagCommit) GetCommitterName() string`

GetCommitterName returns the CommitterName field if non-nil, zero value otherwise.

### GetCommitterNameOk

`func (o *GitAllTagCommit) GetCommitterNameOk() (*string, bool)`

GetCommitterNameOk returns a tuple with the CommitterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitterName

`func (o *GitAllTagCommit) SetCommitterName(v string)`

SetCommitterName sets CommitterName field to given value.

### HasCommitterName

`func (o *GitAllTagCommit) HasCommitterName() bool`

HasCommitterName returns a boolean if a field has been set.

### SetCommitterNameNil

`func (o *GitAllTagCommit) SetCommitterNameNil(b bool)`

 SetCommitterNameNil sets the value for CommitterName to be an explicit nil

### UnsetCommitterName
`func (o *GitAllTagCommit) UnsetCommitterName()`

UnsetCommitterName ensures that no value is present for CommitterName, not even an explicit nil
### GetCreatedAt

`func (o *GitAllTagCommit) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GitAllTagCommit) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GitAllTagCommit) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GitAllTagCommit) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GitAllTagCommit) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GitAllTagCommit) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetParents

`func (o *GitAllTagCommit) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *GitAllTagCommit) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *GitAllTagCommit) SetParents(v []string)`

SetParents sets Parents field to given value.

### HasParents

`func (o *GitAllTagCommit) HasParents() bool`

HasParents returns a boolean if a field has been set.

### SetParentsNil

`func (o *GitAllTagCommit) SetParentsNil(b bool)`

 SetParentsNil sets the value for Parents to be an explicit nil

### UnsetParents
`func (o *GitAllTagCommit) UnsetParents()`

UnsetParents ensures that no value is present for Parents, not even an explicit nil
### GetSha

`func (o *GitAllTagCommit) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitAllTagCommit) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitAllTagCommit) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *GitAllTagCommit) HasSha() bool`

HasSha returns a boolean if a field has been set.

### SetShaNil

`func (o *GitAllTagCommit) SetShaNil(b bool)`

 SetShaNil sets the value for Sha to be an explicit nil

### UnsetSha
`func (o *GitAllTagCommit) UnsetSha()`

UnsetSha ensures that no value is present for Sha, not even an explicit nil
### GetShortMessage

`func (o *GitAllTagCommit) GetShortMessage() string`

GetShortMessage returns the ShortMessage field if non-nil, zero value otherwise.

### GetShortMessageOk

`func (o *GitAllTagCommit) GetShortMessageOk() (*string, bool)`

GetShortMessageOk returns a tuple with the ShortMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortMessage

`func (o *GitAllTagCommit) SetShortMessage(v string)`

SetShortMessage sets ShortMessage field to given value.

### HasShortMessage

`func (o *GitAllTagCommit) HasShortMessage() bool`

HasShortMessage returns a boolean if a field has been set.

### SetShortMessageNil

`func (o *GitAllTagCommit) SetShortMessageNil(b bool)`

 SetShortMessageNil sets the value for ShortMessage to be an explicit nil

### UnsetShortMessage
`func (o *GitAllTagCommit) UnsetShortMessage()`

UnsetShortMessage ensures that no value is present for ShortMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


