# MergeRequestDiffFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobSha** | Pointer to **NullableString** | 文件对应的 blob Id | [optional] [default to ""]
**ChangeType** | Pointer to **NullableString** | 文件改变类型 | [optional] [default to ""]
**Deletions** | Pointer to **NullableInt64** | 一共删除几行 | [optional] 
**Insertions** | Pointer to **NullableInt64** | 一共新增几行 | [optional] 
**Path** | Pointer to **NullableString** | 文件路径 | [optional] [default to ""]
**Size** | Pointer to **NullableInt64** | 文件大小（字节） | [optional] 

## Methods

### NewMergeRequestDiffFile

`func NewMergeRequestDiffFile() *MergeRequestDiffFile`

NewMergeRequestDiffFile instantiates a new MergeRequestDiffFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeRequestDiffFileWithDefaults

`func NewMergeRequestDiffFileWithDefaults() *MergeRequestDiffFile`

NewMergeRequestDiffFileWithDefaults instantiates a new MergeRequestDiffFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobSha

`func (o *MergeRequestDiffFile) GetBlobSha() string`

GetBlobSha returns the BlobSha field if non-nil, zero value otherwise.

### GetBlobShaOk

`func (o *MergeRequestDiffFile) GetBlobShaOk() (*string, bool)`

GetBlobShaOk returns a tuple with the BlobSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobSha

`func (o *MergeRequestDiffFile) SetBlobSha(v string)`

SetBlobSha sets BlobSha field to given value.

### HasBlobSha

`func (o *MergeRequestDiffFile) HasBlobSha() bool`

HasBlobSha returns a boolean if a field has been set.

### SetBlobShaNil

`func (o *MergeRequestDiffFile) SetBlobShaNil(b bool)`

 SetBlobShaNil sets the value for BlobSha to be an explicit nil

### UnsetBlobSha
`func (o *MergeRequestDiffFile) UnsetBlobSha()`

UnsetBlobSha ensures that no value is present for BlobSha, not even an explicit nil
### GetChangeType

`func (o *MergeRequestDiffFile) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *MergeRequestDiffFile) GetChangeTypeOk() (*string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *MergeRequestDiffFile) SetChangeType(v string)`

SetChangeType sets ChangeType field to given value.

### HasChangeType

`func (o *MergeRequestDiffFile) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### SetChangeTypeNil

`func (o *MergeRequestDiffFile) SetChangeTypeNil(b bool)`

 SetChangeTypeNil sets the value for ChangeType to be an explicit nil

### UnsetChangeType
`func (o *MergeRequestDiffFile) UnsetChangeType()`

UnsetChangeType ensures that no value is present for ChangeType, not even an explicit nil
### GetDeletions

`func (o *MergeRequestDiffFile) GetDeletions() int64`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *MergeRequestDiffFile) GetDeletionsOk() (*int64, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *MergeRequestDiffFile) SetDeletions(v int64)`

SetDeletions sets Deletions field to given value.

### HasDeletions

`func (o *MergeRequestDiffFile) HasDeletions() bool`

HasDeletions returns a boolean if a field has been set.

### SetDeletionsNil

`func (o *MergeRequestDiffFile) SetDeletionsNil(b bool)`

 SetDeletionsNil sets the value for Deletions to be an explicit nil

### UnsetDeletions
`func (o *MergeRequestDiffFile) UnsetDeletions()`

UnsetDeletions ensures that no value is present for Deletions, not even an explicit nil
### GetInsertions

`func (o *MergeRequestDiffFile) GetInsertions() int64`

GetInsertions returns the Insertions field if non-nil, zero value otherwise.

### GetInsertionsOk

`func (o *MergeRequestDiffFile) GetInsertionsOk() (*int64, bool)`

GetInsertionsOk returns a tuple with the Insertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertions

`func (o *MergeRequestDiffFile) SetInsertions(v int64)`

SetInsertions sets Insertions field to given value.

### HasInsertions

`func (o *MergeRequestDiffFile) HasInsertions() bool`

HasInsertions returns a boolean if a field has been set.

### SetInsertionsNil

`func (o *MergeRequestDiffFile) SetInsertionsNil(b bool)`

 SetInsertionsNil sets the value for Insertions to be an explicit nil

### UnsetInsertions
`func (o *MergeRequestDiffFile) UnsetInsertions()`

UnsetInsertions ensures that no value is present for Insertions, not even an explicit nil
### GetPath

`func (o *MergeRequestDiffFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MergeRequestDiffFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MergeRequestDiffFile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MergeRequestDiffFile) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *MergeRequestDiffFile) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *MergeRequestDiffFile) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetSize

`func (o *MergeRequestDiffFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MergeRequestDiffFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MergeRequestDiffFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MergeRequestDiffFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *MergeRequestDiffFile) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *MergeRequestDiffFile) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


