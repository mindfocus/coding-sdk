# QcloudApiGitDepotDiffFileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deletions** | **int64** | 总删除行数 | [default to 0]
**DifferentLines** | [**[]QcloudApiGitDepotDifferentLine**](QcloudApiGitDepotDifferentLine.md) |  | 
**Insertions** | **int64** | 总新增行数 | [default to 0]

## Methods

### NewQcloudApiGitDepotDiffFileInfo

`func NewQcloudApiGitDepotDiffFileInfo(deletions int64, differentLines []QcloudApiGitDepotDifferentLine, insertions int64, ) *QcloudApiGitDepotDiffFileInfo`

NewQcloudApiGitDepotDiffFileInfo instantiates a new QcloudApiGitDepotDiffFileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQcloudApiGitDepotDiffFileInfoWithDefaults

`func NewQcloudApiGitDepotDiffFileInfoWithDefaults() *QcloudApiGitDepotDiffFileInfo`

NewQcloudApiGitDepotDiffFileInfoWithDefaults instantiates a new QcloudApiGitDepotDiffFileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletions

`func (o *QcloudApiGitDepotDiffFileInfo) GetDeletions() int64`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *QcloudApiGitDepotDiffFileInfo) GetDeletionsOk() (*int64, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *QcloudApiGitDepotDiffFileInfo) SetDeletions(v int64)`

SetDeletions sets Deletions field to given value.


### GetDifferentLines

`func (o *QcloudApiGitDepotDiffFileInfo) GetDifferentLines() []QcloudApiGitDepotDifferentLine`

GetDifferentLines returns the DifferentLines field if non-nil, zero value otherwise.

### GetDifferentLinesOk

`func (o *QcloudApiGitDepotDiffFileInfo) GetDifferentLinesOk() (*[]QcloudApiGitDepotDifferentLine, bool)`

GetDifferentLinesOk returns a tuple with the DifferentLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifferentLines

`func (o *QcloudApiGitDepotDiffFileInfo) SetDifferentLines(v []QcloudApiGitDepotDifferentLine)`

SetDifferentLines sets DifferentLines field to given value.


### GetInsertions

`func (o *QcloudApiGitDepotDiffFileInfo) GetInsertions() int64`

GetInsertions returns the Insertions field if non-nil, zero value otherwise.

### GetInsertionsOk

`func (o *QcloudApiGitDepotDiffFileInfo) GetInsertionsOk() (*int64, bool)`

GetInsertionsOk returns a tuple with the Insertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertions

`func (o *QcloudApiGitDepotDiffFileInfo) SetInsertions(v int64)`

SetInsertions sets Insertions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


