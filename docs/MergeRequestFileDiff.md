# MergeRequestFileDiff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deletions** | Pointer to **int64** | 总删除条数 | [optional] 
**FileDiffs** | Pointer to [**[]FileDiff**](FileDiff.md) | 文件差异列表 | [optional] 
**Insertions** | Pointer to **int64** | 总新增条数 | [optional] 
**NewSha** | Pointer to **string** | 新请求的 sha 值 | [optional] [default to ""]
**OldSha** | Pointer to **string** | 旧请求的 sha 值 | [optional] [default to ""]

## Methods

### NewMergeRequestFileDiff

`func NewMergeRequestFileDiff() *MergeRequestFileDiff`

NewMergeRequestFileDiff instantiates a new MergeRequestFileDiff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeRequestFileDiffWithDefaults

`func NewMergeRequestFileDiffWithDefaults() *MergeRequestFileDiff`

NewMergeRequestFileDiffWithDefaults instantiates a new MergeRequestFileDiff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletions

`func (o *MergeRequestFileDiff) GetDeletions() int64`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *MergeRequestFileDiff) GetDeletionsOk() (*int64, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *MergeRequestFileDiff) SetDeletions(v int64)`

SetDeletions sets Deletions field to given value.

### HasDeletions

`func (o *MergeRequestFileDiff) HasDeletions() bool`

HasDeletions returns a boolean if a field has been set.

### GetFileDiffs

`func (o *MergeRequestFileDiff) GetFileDiffs() []FileDiff`

GetFileDiffs returns the FileDiffs field if non-nil, zero value otherwise.

### GetFileDiffsOk

`func (o *MergeRequestFileDiff) GetFileDiffsOk() (*[]FileDiff, bool)`

GetFileDiffsOk returns a tuple with the FileDiffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDiffs

`func (o *MergeRequestFileDiff) SetFileDiffs(v []FileDiff)`

SetFileDiffs sets FileDiffs field to given value.

### HasFileDiffs

`func (o *MergeRequestFileDiff) HasFileDiffs() bool`

HasFileDiffs returns a boolean if a field has been set.

### GetInsertions

`func (o *MergeRequestFileDiff) GetInsertions() int64`

GetInsertions returns the Insertions field if non-nil, zero value otherwise.

### GetInsertionsOk

`func (o *MergeRequestFileDiff) GetInsertionsOk() (*int64, bool)`

GetInsertionsOk returns a tuple with the Insertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertions

`func (o *MergeRequestFileDiff) SetInsertions(v int64)`

SetInsertions sets Insertions field to given value.

### HasInsertions

`func (o *MergeRequestFileDiff) HasInsertions() bool`

HasInsertions returns a boolean if a field has been set.

### GetNewSha

`func (o *MergeRequestFileDiff) GetNewSha() string`

GetNewSha returns the NewSha field if non-nil, zero value otherwise.

### GetNewShaOk

`func (o *MergeRequestFileDiff) GetNewShaOk() (*string, bool)`

GetNewShaOk returns a tuple with the NewSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSha

`func (o *MergeRequestFileDiff) SetNewSha(v string)`

SetNewSha sets NewSha field to given value.

### HasNewSha

`func (o *MergeRequestFileDiff) HasNewSha() bool`

HasNewSha returns a boolean if a field has been set.

### GetOldSha

`func (o *MergeRequestFileDiff) GetOldSha() string`

GetOldSha returns the OldSha field if non-nil, zero value otherwise.

### GetOldShaOk

`func (o *MergeRequestFileDiff) GetOldShaOk() (*string, bool)`

GetOldShaOk returns a tuple with the OldSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSha

`func (o *MergeRequestFileDiff) SetOldSha(v string)`

SetOldSha sets OldSha field to given value.

### HasOldSha

`func (o *MergeRequestFileDiff) HasOldSha() bool`

HasOldSha returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


