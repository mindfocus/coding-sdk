# DifferentOfCommitDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commits** | Pointer to [**[]GitCommit**](GitCommit.md) | 请求列表 | [optional] 
**Deletions** | Pointer to **int64** | 总删除行数 | [optional] 
**DifferentOfCommits** | Pointer to [**[]DifferentOfCommit**](DifferentOfCommit.md) | 差异文件列表 | [optional] 
**Insertions** | Pointer to **int64** | 总新增行数 | [optional] 
**UpdateFileNum** | Pointer to **int64** | 总文件修改数 | [optional] 

## Methods

### NewDifferentOfCommitDetail

`func NewDifferentOfCommitDetail() *DifferentOfCommitDetail`

NewDifferentOfCommitDetail instantiates a new DifferentOfCommitDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDifferentOfCommitDetailWithDefaults

`func NewDifferentOfCommitDetailWithDefaults() *DifferentOfCommitDetail`

NewDifferentOfCommitDetailWithDefaults instantiates a new DifferentOfCommitDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommits

`func (o *DifferentOfCommitDetail) GetCommits() []GitCommit`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *DifferentOfCommitDetail) GetCommitsOk() (*[]GitCommit, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *DifferentOfCommitDetail) SetCommits(v []GitCommit)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *DifferentOfCommitDetail) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetDeletions

`func (o *DifferentOfCommitDetail) GetDeletions() int64`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *DifferentOfCommitDetail) GetDeletionsOk() (*int64, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *DifferentOfCommitDetail) SetDeletions(v int64)`

SetDeletions sets Deletions field to given value.

### HasDeletions

`func (o *DifferentOfCommitDetail) HasDeletions() bool`

HasDeletions returns a boolean if a field has been set.

### GetDifferentOfCommits

`func (o *DifferentOfCommitDetail) GetDifferentOfCommits() []DifferentOfCommit`

GetDifferentOfCommits returns the DifferentOfCommits field if non-nil, zero value otherwise.

### GetDifferentOfCommitsOk

`func (o *DifferentOfCommitDetail) GetDifferentOfCommitsOk() (*[]DifferentOfCommit, bool)`

GetDifferentOfCommitsOk returns a tuple with the DifferentOfCommits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifferentOfCommits

`func (o *DifferentOfCommitDetail) SetDifferentOfCommits(v []DifferentOfCommit)`

SetDifferentOfCommits sets DifferentOfCommits field to given value.

### HasDifferentOfCommits

`func (o *DifferentOfCommitDetail) HasDifferentOfCommits() bool`

HasDifferentOfCommits returns a boolean if a field has been set.

### GetInsertions

`func (o *DifferentOfCommitDetail) GetInsertions() int64`

GetInsertions returns the Insertions field if non-nil, zero value otherwise.

### GetInsertionsOk

`func (o *DifferentOfCommitDetail) GetInsertionsOk() (*int64, bool)`

GetInsertionsOk returns a tuple with the Insertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertions

`func (o *DifferentOfCommitDetail) SetInsertions(v int64)`

SetInsertions sets Insertions field to given value.

### HasInsertions

`func (o *DifferentOfCommitDetail) HasInsertions() bool`

HasInsertions returns a boolean if a field has been set.

### GetUpdateFileNum

`func (o *DifferentOfCommitDetail) GetUpdateFileNum() int64`

GetUpdateFileNum returns the UpdateFileNum field if non-nil, zero value otherwise.

### GetUpdateFileNumOk

`func (o *DifferentOfCommitDetail) GetUpdateFileNumOk() (*int64, bool)`

GetUpdateFileNumOk returns a tuple with the UpdateFileNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateFileNum

`func (o *DifferentOfCommitDetail) SetUpdateFileNum(v int64)`

SetUpdateFileNum sets UpdateFileNum field to given value.

### HasUpdateFileNum

`func (o *DifferentOfCommitDetail) HasUpdateFileNum() bool`

HasUpdateFileNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


