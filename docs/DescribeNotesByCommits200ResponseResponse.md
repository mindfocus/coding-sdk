# DescribeNotesByCommits200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notes** | Pointer to [**[]CommitNote**](CommitNote.md) | commit notes 集合 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeNotesByCommits200ResponseResponse

`func NewDescribeNotesByCommits200ResponseResponse() *DescribeNotesByCommits200ResponseResponse`

NewDescribeNotesByCommits200ResponseResponse instantiates a new DescribeNotesByCommits200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeNotesByCommits200ResponseResponseWithDefaults

`func NewDescribeNotesByCommits200ResponseResponseWithDefaults() *DescribeNotesByCommits200ResponseResponse`

NewDescribeNotesByCommits200ResponseResponseWithDefaults instantiates a new DescribeNotesByCommits200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotes

`func (o *DescribeNotesByCommits200ResponseResponse) GetNotes() []CommitNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DescribeNotesByCommits200ResponseResponse) GetNotesOk() (*[]CommitNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DescribeNotesByCommits200ResponseResponse) SetNotes(v []CommitNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DescribeNotesByCommits200ResponseResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeNotesByCommits200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeNotesByCommits200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeNotesByCommits200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeNotesByCommits200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

