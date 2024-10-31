# MergeRequestDiff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deletions** | Pointer to **NullableInt64** | 一共减少行数 | [optional] 
**Insertions** | Pointer to **NullableInt64** | 一共新增行数 | [optional] 
**IsLarge** | Pointer to **NullableBool** | diff 信息是否过多，不宜于展示 | [optional] [default to false]
**NewSha** | Pointer to **NullableString** | 源分支对应的Sha值 | [optional] [default to ""]
**OldSha** | Pointer to **NullableString** | 目标分支对应的Sha值 | [optional] [default to ""]
**Paths** | Pointer to [**[]MergeRequestDiffFile**](MergeRequestDiffFile.md) | 文件信息 | [optional] 

## Methods

### NewMergeRequestDiff

`func NewMergeRequestDiff() *MergeRequestDiff`

NewMergeRequestDiff instantiates a new MergeRequestDiff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeRequestDiffWithDefaults

`func NewMergeRequestDiffWithDefaults() *MergeRequestDiff`

NewMergeRequestDiffWithDefaults instantiates a new MergeRequestDiff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletions

`func (o *MergeRequestDiff) GetDeletions() int64`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *MergeRequestDiff) GetDeletionsOk() (*int64, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *MergeRequestDiff) SetDeletions(v int64)`

SetDeletions sets Deletions field to given value.

### HasDeletions

`func (o *MergeRequestDiff) HasDeletions() bool`

HasDeletions returns a boolean if a field has been set.

### SetDeletionsNil

`func (o *MergeRequestDiff) SetDeletionsNil(b bool)`

 SetDeletionsNil sets the value for Deletions to be an explicit nil

### UnsetDeletions
`func (o *MergeRequestDiff) UnsetDeletions()`

UnsetDeletions ensures that no value is present for Deletions, not even an explicit nil
### GetInsertions

`func (o *MergeRequestDiff) GetInsertions() int64`

GetInsertions returns the Insertions field if non-nil, zero value otherwise.

### GetInsertionsOk

`func (o *MergeRequestDiff) GetInsertionsOk() (*int64, bool)`

GetInsertionsOk returns a tuple with the Insertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertions

`func (o *MergeRequestDiff) SetInsertions(v int64)`

SetInsertions sets Insertions field to given value.

### HasInsertions

`func (o *MergeRequestDiff) HasInsertions() bool`

HasInsertions returns a boolean if a field has been set.

### SetInsertionsNil

`func (o *MergeRequestDiff) SetInsertionsNil(b bool)`

 SetInsertionsNil sets the value for Insertions to be an explicit nil

### UnsetInsertions
`func (o *MergeRequestDiff) UnsetInsertions()`

UnsetInsertions ensures that no value is present for Insertions, not even an explicit nil
### GetIsLarge

`func (o *MergeRequestDiff) GetIsLarge() bool`

GetIsLarge returns the IsLarge field if non-nil, zero value otherwise.

### GetIsLargeOk

`func (o *MergeRequestDiff) GetIsLargeOk() (*bool, bool)`

GetIsLargeOk returns a tuple with the IsLarge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLarge

`func (o *MergeRequestDiff) SetIsLarge(v bool)`

SetIsLarge sets IsLarge field to given value.

### HasIsLarge

`func (o *MergeRequestDiff) HasIsLarge() bool`

HasIsLarge returns a boolean if a field has been set.

### SetIsLargeNil

`func (o *MergeRequestDiff) SetIsLargeNil(b bool)`

 SetIsLargeNil sets the value for IsLarge to be an explicit nil

### UnsetIsLarge
`func (o *MergeRequestDiff) UnsetIsLarge()`

UnsetIsLarge ensures that no value is present for IsLarge, not even an explicit nil
### GetNewSha

`func (o *MergeRequestDiff) GetNewSha() string`

GetNewSha returns the NewSha field if non-nil, zero value otherwise.

### GetNewShaOk

`func (o *MergeRequestDiff) GetNewShaOk() (*string, bool)`

GetNewShaOk returns a tuple with the NewSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSha

`func (o *MergeRequestDiff) SetNewSha(v string)`

SetNewSha sets NewSha field to given value.

### HasNewSha

`func (o *MergeRequestDiff) HasNewSha() bool`

HasNewSha returns a boolean if a field has been set.

### SetNewShaNil

`func (o *MergeRequestDiff) SetNewShaNil(b bool)`

 SetNewShaNil sets the value for NewSha to be an explicit nil

### UnsetNewSha
`func (o *MergeRequestDiff) UnsetNewSha()`

UnsetNewSha ensures that no value is present for NewSha, not even an explicit nil
### GetOldSha

`func (o *MergeRequestDiff) GetOldSha() string`

GetOldSha returns the OldSha field if non-nil, zero value otherwise.

### GetOldShaOk

`func (o *MergeRequestDiff) GetOldShaOk() (*string, bool)`

GetOldShaOk returns a tuple with the OldSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSha

`func (o *MergeRequestDiff) SetOldSha(v string)`

SetOldSha sets OldSha field to given value.

### HasOldSha

`func (o *MergeRequestDiff) HasOldSha() bool`

HasOldSha returns a boolean if a field has been set.

### SetOldShaNil

`func (o *MergeRequestDiff) SetOldShaNil(b bool)`

 SetOldShaNil sets the value for OldSha to be an explicit nil

### UnsetOldSha
`func (o *MergeRequestDiff) UnsetOldSha()`

UnsetOldSha ensures that no value is present for OldSha, not even an explicit nil
### GetPaths

`func (o *MergeRequestDiff) GetPaths() []MergeRequestDiffFile`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *MergeRequestDiff) GetPathsOk() (*[]MergeRequestDiffFile, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *MergeRequestDiff) SetPaths(v []MergeRequestDiffFile)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *MergeRequestDiff) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### SetPathsNil

`func (o *MergeRequestDiff) SetPathsNil(b bool)`

 SetPathsNil sets the value for Paths to be an explicit nil

### UnsetPaths
`func (o *MergeRequestDiff) UnsetPaths()`

UnsetPaths ensures that no value is present for Paths, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


