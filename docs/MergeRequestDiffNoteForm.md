# MergeRequestDiffNoteForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitSha** | **string** | 提交Id | [default to ""]
**Index** | **int32** | Diff信息行数 | 
**Path** | **string** | 改动文件路径 | [default to ""]

## Methods

### NewMergeRequestDiffNoteForm

`func NewMergeRequestDiffNoteForm(commitSha string, index int32, path string, ) *MergeRequestDiffNoteForm`

NewMergeRequestDiffNoteForm instantiates a new MergeRequestDiffNoteForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeRequestDiffNoteFormWithDefaults

`func NewMergeRequestDiffNoteFormWithDefaults() *MergeRequestDiffNoteForm`

NewMergeRequestDiffNoteFormWithDefaults instantiates a new MergeRequestDiffNoteForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitSha

`func (o *MergeRequestDiffNoteForm) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *MergeRequestDiffNoteForm) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *MergeRequestDiffNoteForm) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### GetIndex

`func (o *MergeRequestDiffNoteForm) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MergeRequestDiffNoteForm) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *MergeRequestDiffNoteForm) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetPath

`func (o *MergeRequestDiffNoteForm) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MergeRequestDiffNoteForm) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MergeRequestDiffNoteForm) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


