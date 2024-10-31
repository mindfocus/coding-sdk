# FileDiff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeType** | Pointer to **string** | 操作类型,具体值为: ADD(添加), MODIFY(修改), DELETE(删除), REPLACE(替换) | [optional] [default to ""]
**Deletions** | Pointer to **int64** | 删除的行数 | [optional] 
**DiffLines** | Pointer to [**[]DifferentLine**](DifferentLine.md) | Diff Line | [optional] 
**Insertions** | Pointer to **int64** | 新增的行数 | [optional] 
**ObjectId** | Pointer to **string** | 请求的objectid | [optional] [default to ""]
**Path** | Pointer to **string** | 文件路径 | [optional] [default to ""]

## Methods

### NewFileDiff

`func NewFileDiff() *FileDiff`

NewFileDiff instantiates a new FileDiff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileDiffWithDefaults

`func NewFileDiffWithDefaults() *FileDiff`

NewFileDiffWithDefaults instantiates a new FileDiff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeType

`func (o *FileDiff) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *FileDiff) GetChangeTypeOk() (*string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *FileDiff) SetChangeType(v string)`

SetChangeType sets ChangeType field to given value.

### HasChangeType

`func (o *FileDiff) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### GetDeletions

`func (o *FileDiff) GetDeletions() int64`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *FileDiff) GetDeletionsOk() (*int64, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *FileDiff) SetDeletions(v int64)`

SetDeletions sets Deletions field to given value.

### HasDeletions

`func (o *FileDiff) HasDeletions() bool`

HasDeletions returns a boolean if a field has been set.

### GetDiffLines

`func (o *FileDiff) GetDiffLines() []DifferentLine`

GetDiffLines returns the DiffLines field if non-nil, zero value otherwise.

### GetDiffLinesOk

`func (o *FileDiff) GetDiffLinesOk() (*[]DifferentLine, bool)`

GetDiffLinesOk returns a tuple with the DiffLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffLines

`func (o *FileDiff) SetDiffLines(v []DifferentLine)`

SetDiffLines sets DiffLines field to given value.

### HasDiffLines

`func (o *FileDiff) HasDiffLines() bool`

HasDiffLines returns a boolean if a field has been set.

### GetInsertions

`func (o *FileDiff) GetInsertions() int64`

GetInsertions returns the Insertions field if non-nil, zero value otherwise.

### GetInsertionsOk

`func (o *FileDiff) GetInsertionsOk() (*int64, bool)`

GetInsertionsOk returns a tuple with the Insertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertions

`func (o *FileDiff) SetInsertions(v int64)`

SetInsertions sets Insertions field to given value.

### HasInsertions

`func (o *FileDiff) HasInsertions() bool`

HasInsertions returns a boolean if a field has been set.

### GetObjectId

`func (o *FileDiff) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *FileDiff) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *FileDiff) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *FileDiff) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetPath

`func (o *FileDiff) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileDiff) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileDiff) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FileDiff) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


