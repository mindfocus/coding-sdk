# CommitInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitDate** | Pointer to **NullableInt64** | 提交日期 | [optional] 
**CommitSha** | Pointer to **NullableString** | 提交信息 | [optional] [default to ""]
**Committer** | Pointer to [**CodingUser**](CodingUser.md) |  | [optional] 
**LineNumber** | Pointer to **NullableInt64** | 行数 | [optional] 

## Methods

### NewCommitInfo

`func NewCommitInfo() *CommitInfo`

NewCommitInfo instantiates a new CommitInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitInfoWithDefaults

`func NewCommitInfoWithDefaults() *CommitInfo`

NewCommitInfoWithDefaults instantiates a new CommitInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitDate

`func (o *CommitInfo) GetCommitDate() int64`

GetCommitDate returns the CommitDate field if non-nil, zero value otherwise.

### GetCommitDateOk

`func (o *CommitInfo) GetCommitDateOk() (*int64, bool)`

GetCommitDateOk returns a tuple with the CommitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitDate

`func (o *CommitInfo) SetCommitDate(v int64)`

SetCommitDate sets CommitDate field to given value.

### HasCommitDate

`func (o *CommitInfo) HasCommitDate() bool`

HasCommitDate returns a boolean if a field has been set.

### SetCommitDateNil

`func (o *CommitInfo) SetCommitDateNil(b bool)`

 SetCommitDateNil sets the value for CommitDate to be an explicit nil

### UnsetCommitDate
`func (o *CommitInfo) UnsetCommitDate()`

UnsetCommitDate ensures that no value is present for CommitDate, not even an explicit nil
### GetCommitSha

`func (o *CommitInfo) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *CommitInfo) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *CommitInfo) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *CommitInfo) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### SetCommitShaNil

`func (o *CommitInfo) SetCommitShaNil(b bool)`

 SetCommitShaNil sets the value for CommitSha to be an explicit nil

### UnsetCommitSha
`func (o *CommitInfo) UnsetCommitSha()`

UnsetCommitSha ensures that no value is present for CommitSha, not even an explicit nil
### GetCommitter

`func (o *CommitInfo) GetCommitter() CodingUser`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *CommitInfo) GetCommitterOk() (*CodingUser, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *CommitInfo) SetCommitter(v CodingUser)`

SetCommitter sets Committer field to given value.

### HasCommitter

`func (o *CommitInfo) HasCommitter() bool`

HasCommitter returns a boolean if a field has been set.

### GetLineNumber

`func (o *CommitInfo) GetLineNumber() int64`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *CommitInfo) GetLineNumberOk() (*int64, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *CommitInfo) SetLineNumber(v int64)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *CommitInfo) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### SetLineNumberNil

`func (o *CommitInfo) SetLineNumberNil(b bool)`

 SetLineNumberNil sets the value for LineNumber to be an explicit nil

### UnsetLineNumber
`func (o *CommitInfo) UnsetLineNumber()`

UnsetLineNumber ensures that no value is present for LineNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


