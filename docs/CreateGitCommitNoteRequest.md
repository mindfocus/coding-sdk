# CreateGitCommitNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitMessage** | **string** | 提交信息 | 
**CommitSha** | **string** | 提交的 Sha | 
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**Note** | **string** | 注释的详情信息 | 
**NotesRef** | Pointer to **string** | 注释分支 ref。建议默认不填 | [optional] 

## Methods

### NewCreateGitCommitNoteRequest

`func NewCreateGitCommitNoteRequest(commitMessage string, commitSha string, depotId int64, note string, ) *CreateGitCommitNoteRequest`

NewCreateGitCommitNoteRequest instantiates a new CreateGitCommitNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGitCommitNoteRequestWithDefaults

`func NewCreateGitCommitNoteRequestWithDefaults() *CreateGitCommitNoteRequest`

NewCreateGitCommitNoteRequestWithDefaults instantiates a new CreateGitCommitNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitMessage

`func (o *CreateGitCommitNoteRequest) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *CreateGitCommitNoteRequest) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *CreateGitCommitNoteRequest) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.


### GetCommitSha

`func (o *CreateGitCommitNoteRequest) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *CreateGitCommitNoteRequest) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *CreateGitCommitNoteRequest) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### GetDepotId

`func (o *CreateGitCommitNoteRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CreateGitCommitNoteRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CreateGitCommitNoteRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *CreateGitCommitNoteRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateGitCommitNoteRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateGitCommitNoteRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *CreateGitCommitNoteRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetNote

`func (o *CreateGitCommitNoteRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateGitCommitNoteRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateGitCommitNoteRequest) SetNote(v string)`

SetNote sets Note field to given value.


### GetNotesRef

`func (o *CreateGitCommitNoteRequest) GetNotesRef() string`

GetNotesRef returns the NotesRef field if non-nil, zero value otherwise.

### GetNotesRefOk

`func (o *CreateGitCommitNoteRequest) GetNotesRefOk() (*string, bool)`

GetNotesRefOk returns a tuple with the NotesRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesRef

`func (o *CreateGitCommitNoteRequest) SetNotesRef(v string)`

SetNotesRef sets NotesRef field to given value.

### HasNotesRef

`func (o *CreateGitCommitNoteRequest) HasNotesRef() bool`

HasNotesRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


