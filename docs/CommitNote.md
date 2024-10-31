# CommitNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitSha** | Pointer to **string** | 请求的sha值 | [optional] [default to ""]
**NoteContent** | Pointer to **string** | note的具体值 | [optional] [default to ""]

## Methods

### NewCommitNote

`func NewCommitNote() *CommitNote`

NewCommitNote instantiates a new CommitNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitNoteWithDefaults

`func NewCommitNoteWithDefaults() *CommitNote`

NewCommitNoteWithDefaults instantiates a new CommitNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitSha

`func (o *CommitNote) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *CommitNote) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *CommitNote) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *CommitNote) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### GetNoteContent

`func (o *CommitNote) GetNoteContent() string`

GetNoteContent returns the NoteContent field if non-nil, zero value otherwise.

### GetNoteContentOk

`func (o *CommitNote) GetNoteContentOk() (*string, bool)`

GetNoteContentOk returns a tuple with the NoteContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteContent

`func (o *CommitNote) SetNoteContent(v string)`

SetNoteContent sets NoteContent field to given value.

### HasNoteContent

`func (o *CommitNote) HasNoteContent() bool`

HasNoteContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


