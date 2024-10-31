# DescribeGitCommitNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitSha** | **string** | 提交的 Sha | 
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**NotesRef** | **string** | 注释 分支 Ref | 

## Methods

### NewDescribeGitCommitNoteRequest

`func NewDescribeGitCommitNoteRequest(commitSha string, depotId int64, notesRef string, ) *DescribeGitCommitNoteRequest`

NewDescribeGitCommitNoteRequest instantiates a new DescribeGitCommitNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitCommitNoteRequestWithDefaults

`func NewDescribeGitCommitNoteRequestWithDefaults() *DescribeGitCommitNoteRequest`

NewDescribeGitCommitNoteRequestWithDefaults instantiates a new DescribeGitCommitNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitSha

`func (o *DescribeGitCommitNoteRequest) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *DescribeGitCommitNoteRequest) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *DescribeGitCommitNoteRequest) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### GetDepotId

`func (o *DescribeGitCommitNoteRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitCommitNoteRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitCommitNoteRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitCommitNoteRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitCommitNoteRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitCommitNoteRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitCommitNoteRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetNotesRef

`func (o *DescribeGitCommitNoteRequest) GetNotesRef() string`

GetNotesRef returns the NotesRef field if non-nil, zero value otherwise.

### GetNotesRefOk

`func (o *DescribeGitCommitNoteRequest) GetNotesRefOk() (*string, bool)`

GetNotesRefOk returns a tuple with the NotesRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesRef

`func (o *DescribeGitCommitNoteRequest) SetNotesRef(v string)`

SetNotesRef sets NotesRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


