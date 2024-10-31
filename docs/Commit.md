# Commit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**Committer**](Committer.md) |  | [optional] 
**AuthorEmail** | Pointer to **string** | 作者邮箱 | [optional] [default to ""]
**AuthorName** | Pointer to **string** | 作者姓名 | [optional] [default to ""]
**CommitDate** | Pointer to **int64** | 提交时间的时间戳 | [optional] 
**Committer** | Pointer to [**Committer**](Committer.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** | 提交时间的时间戳 | [optional] 
**FullMessage** | Pointer to **string** | 提交附带的全部信息 | [optional] [default to ""]
**Parents** | Pointer to **[]string** | 父提交 | [optional] 
**Sha** | Pointer to **string** | 提交的sha值(commitId) | [optional] [default to ""]
**ShortMessage** | Pointer to **string** | 提交附带的message | [optional] [default to ""]

## Methods

### NewCommit

`func NewCommit() *Commit`

NewCommit instantiates a new Commit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitWithDefaults

`func NewCommitWithDefaults() *Commit`

NewCommitWithDefaults instantiates a new Commit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *Commit) GetAuthor() Committer`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Commit) GetAuthorOk() (*Committer, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Commit) SetAuthor(v Committer)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Commit) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetAuthorEmail

`func (o *Commit) GetAuthorEmail() string`

GetAuthorEmail returns the AuthorEmail field if non-nil, zero value otherwise.

### GetAuthorEmailOk

`func (o *Commit) GetAuthorEmailOk() (*string, bool)`

GetAuthorEmailOk returns a tuple with the AuthorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorEmail

`func (o *Commit) SetAuthorEmail(v string)`

SetAuthorEmail sets AuthorEmail field to given value.

### HasAuthorEmail

`func (o *Commit) HasAuthorEmail() bool`

HasAuthorEmail returns a boolean if a field has been set.

### GetAuthorName

`func (o *Commit) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *Commit) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *Commit) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *Commit) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### GetCommitDate

`func (o *Commit) GetCommitDate() int64`

GetCommitDate returns the CommitDate field if non-nil, zero value otherwise.

### GetCommitDateOk

`func (o *Commit) GetCommitDateOk() (*int64, bool)`

GetCommitDateOk returns a tuple with the CommitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitDate

`func (o *Commit) SetCommitDate(v int64)`

SetCommitDate sets CommitDate field to given value.

### HasCommitDate

`func (o *Commit) HasCommitDate() bool`

HasCommitDate returns a boolean if a field has been set.

### GetCommitter

`func (o *Commit) GetCommitter() Committer`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *Commit) GetCommitterOk() (*Committer, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *Commit) SetCommitter(v Committer)`

SetCommitter sets Committer field to given value.

### HasCommitter

`func (o *Commit) HasCommitter() bool`

HasCommitter returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Commit) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Commit) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Commit) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Commit) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFullMessage

`func (o *Commit) GetFullMessage() string`

GetFullMessage returns the FullMessage field if non-nil, zero value otherwise.

### GetFullMessageOk

`func (o *Commit) GetFullMessageOk() (*string, bool)`

GetFullMessageOk returns a tuple with the FullMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullMessage

`func (o *Commit) SetFullMessage(v string)`

SetFullMessage sets FullMessage field to given value.

### HasFullMessage

`func (o *Commit) HasFullMessage() bool`

HasFullMessage returns a boolean if a field has been set.

### GetParents

`func (o *Commit) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *Commit) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *Commit) SetParents(v []string)`

SetParents sets Parents field to given value.

### HasParents

`func (o *Commit) HasParents() bool`

HasParents returns a boolean if a field has been set.

### SetParentsNil

`func (o *Commit) SetParentsNil(b bool)`

 SetParentsNil sets the value for Parents to be an explicit nil

### UnsetParents
`func (o *Commit) UnsetParents()`

UnsetParents ensures that no value is present for Parents, not even an explicit nil
### GetSha

`func (o *Commit) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *Commit) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *Commit) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *Commit) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetShortMessage

`func (o *Commit) GetShortMessage() string`

GetShortMessage returns the ShortMessage field if non-nil, zero value otherwise.

### GetShortMessageOk

`func (o *Commit) GetShortMessageOk() (*string, bool)`

GetShortMessageOk returns a tuple with the ShortMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortMessage

`func (o *Commit) SetShortMessage(v string)`

SetShortMessage sets ShortMessage field to given value.

### HasShortMessage

`func (o *Commit) HasShortMessage() bool`

HasShortMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


