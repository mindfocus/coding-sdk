# MergeRequestNoteList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Note** | Pointer to [**[]MergeRequestNote**](MergeRequestNote.md) | 合并请求中父子评论列表、diff信息同行评论列表 | [optional] 

## Methods

### NewMergeRequestNoteList

`func NewMergeRequestNoteList() *MergeRequestNoteList`

NewMergeRequestNoteList instantiates a new MergeRequestNoteList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeRequestNoteListWithDefaults

`func NewMergeRequestNoteListWithDefaults() *MergeRequestNoteList`

NewMergeRequestNoteListWithDefaults instantiates a new MergeRequestNoteList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNote

`func (o *MergeRequestNoteList) GetNote() []MergeRequestNote`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *MergeRequestNoteList) GetNoteOk() (*[]MergeRequestNote, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *MergeRequestNoteList) SetNote(v []MergeRequestNote)`

SetNote sets Note field to given value.

### HasNote

`func (o *MergeRequestNoteList) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


