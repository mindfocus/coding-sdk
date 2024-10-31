# GitDiff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeType** | Pointer to **string** | 文件改变类型 | [optional] [default to ""]
**Content** | Pointer to **string** | diff信息内容 | [optional] [default to ""]
**Deletions** | Pointer to **int64** | 一共删除几行 | [optional] 
**Insertions** | Pointer to **int64** | 一共新增几行 | [optional] 
**Lines** | Pointer to [**[]Line**](Line.md) | diff每行信息拆解后的集合信息 | [optional] 
**NewMode** | Pointer to **string** | 修改后文件的权限 | [optional] [default to ""]
**OldMode** | Pointer to **string** | 修改前文件的权限 | [optional] [default to ""]
**Path** | Pointer to **string** | 文件路径 | [optional] [default to ""]

## Methods

### NewGitDiff

`func NewGitDiff() *GitDiff`

NewGitDiff instantiates a new GitDiff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitDiffWithDefaults

`func NewGitDiffWithDefaults() *GitDiff`

NewGitDiffWithDefaults instantiates a new GitDiff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeType

`func (o *GitDiff) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *GitDiff) GetChangeTypeOk() (*string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *GitDiff) SetChangeType(v string)`

SetChangeType sets ChangeType field to given value.

### HasChangeType

`func (o *GitDiff) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### GetContent

`func (o *GitDiff) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GitDiff) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GitDiff) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GitDiff) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDeletions

`func (o *GitDiff) GetDeletions() int64`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *GitDiff) GetDeletionsOk() (*int64, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *GitDiff) SetDeletions(v int64)`

SetDeletions sets Deletions field to given value.

### HasDeletions

`func (o *GitDiff) HasDeletions() bool`

HasDeletions returns a boolean if a field has been set.

### GetInsertions

`func (o *GitDiff) GetInsertions() int64`

GetInsertions returns the Insertions field if non-nil, zero value otherwise.

### GetInsertionsOk

`func (o *GitDiff) GetInsertionsOk() (*int64, bool)`

GetInsertionsOk returns a tuple with the Insertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertions

`func (o *GitDiff) SetInsertions(v int64)`

SetInsertions sets Insertions field to given value.

### HasInsertions

`func (o *GitDiff) HasInsertions() bool`

HasInsertions returns a boolean if a field has been set.

### GetLines

`func (o *GitDiff) GetLines() []Line`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *GitDiff) GetLinesOk() (*[]Line, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *GitDiff) SetLines(v []Line)`

SetLines sets Lines field to given value.

### HasLines

`func (o *GitDiff) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetNewMode

`func (o *GitDiff) GetNewMode() string`

GetNewMode returns the NewMode field if non-nil, zero value otherwise.

### GetNewModeOk

`func (o *GitDiff) GetNewModeOk() (*string, bool)`

GetNewModeOk returns a tuple with the NewMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMode

`func (o *GitDiff) SetNewMode(v string)`

SetNewMode sets NewMode field to given value.

### HasNewMode

`func (o *GitDiff) HasNewMode() bool`

HasNewMode returns a boolean if a field has been set.

### GetOldMode

`func (o *GitDiff) GetOldMode() string`

GetOldMode returns the OldMode field if non-nil, zero value otherwise.

### GetOldModeOk

`func (o *GitDiff) GetOldModeOk() (*string, bool)`

GetOldModeOk returns a tuple with the OldMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldMode

`func (o *GitDiff) SetOldMode(v string)`

SetOldMode sets OldMode field to given value.

### HasOldMode

`func (o *GitDiff) HasOldMode() bool`

HasOldMode returns a boolean if a field has been set.

### GetPath

`func (o *GitDiff) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GitDiff) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GitDiff) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GitDiff) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


