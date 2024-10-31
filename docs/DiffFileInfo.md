# DiffFileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deletions** | Pointer to **int64** | 总删除行数 | [optional] 
**DifferentLines** | Pointer to [**[]DifferentLine**](DifferentLine.md) | 差异信息 | [optional] 
**Insertions** | Pointer to **int64** | 总新增行数 | [optional] 

## Methods

### NewDiffFileInfo

`func NewDiffFileInfo() *DiffFileInfo`

NewDiffFileInfo instantiates a new DiffFileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiffFileInfoWithDefaults

`func NewDiffFileInfoWithDefaults() *DiffFileInfo`

NewDiffFileInfoWithDefaults instantiates a new DiffFileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletions

`func (o *DiffFileInfo) GetDeletions() int64`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *DiffFileInfo) GetDeletionsOk() (*int64, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *DiffFileInfo) SetDeletions(v int64)`

SetDeletions sets Deletions field to given value.

### HasDeletions

`func (o *DiffFileInfo) HasDeletions() bool`

HasDeletions returns a boolean if a field has been set.

### GetDifferentLines

`func (o *DiffFileInfo) GetDifferentLines() []DifferentLine`

GetDifferentLines returns the DifferentLines field if non-nil, zero value otherwise.

### GetDifferentLinesOk

`func (o *DiffFileInfo) GetDifferentLinesOk() (*[]DifferentLine, bool)`

GetDifferentLinesOk returns a tuple with the DifferentLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifferentLines

`func (o *DiffFileInfo) SetDifferentLines(v []DifferentLine)`

SetDifferentLines sets DifferentLines field to given value.

### HasDifferentLines

`func (o *DiffFileInfo) HasDifferentLines() bool`

HasDifferentLines returns a boolean if a field has been set.

### GetInsertions

`func (o *DiffFileInfo) GetInsertions() int64`

GetInsertions returns the Insertions field if non-nil, zero value otherwise.

### GetInsertionsOk

`func (o *DiffFileInfo) GetInsertionsOk() (*int64, bool)`

GetInsertionsOk returns a tuple with the Insertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertions

`func (o *DiffFileInfo) SetInsertions(v int64)`

SetInsertions sets Insertions field to given value.

### HasInsertions

`func (o *DiffFileInfo) HasInsertions() bool`

HasInsertions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


