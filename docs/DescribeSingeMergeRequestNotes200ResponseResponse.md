# DescribeSingeMergeRequestNotes200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notes** | Pointer to [**[]MergeRequestNoteList**](MergeRequestNoteList.md) | 单个合并请求行评论内容 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeSingeMergeRequestNotes200ResponseResponse

`func NewDescribeSingeMergeRequestNotes200ResponseResponse() *DescribeSingeMergeRequestNotes200ResponseResponse`

NewDescribeSingeMergeRequestNotes200ResponseResponse instantiates a new DescribeSingeMergeRequestNotes200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeSingeMergeRequestNotes200ResponseResponseWithDefaults

`func NewDescribeSingeMergeRequestNotes200ResponseResponseWithDefaults() *DescribeSingeMergeRequestNotes200ResponseResponse`

NewDescribeSingeMergeRequestNotes200ResponseResponseWithDefaults instantiates a new DescribeSingeMergeRequestNotes200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotes

`func (o *DescribeSingeMergeRequestNotes200ResponseResponse) GetNotes() []MergeRequestNoteList`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DescribeSingeMergeRequestNotes200ResponseResponse) GetNotesOk() (*[]MergeRequestNoteList, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DescribeSingeMergeRequestNotes200ResponseResponse) SetNotes(v []MergeRequestNoteList)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DescribeSingeMergeRequestNotes200ResponseResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeSingeMergeRequestNotes200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeSingeMergeRequestNotes200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeSingeMergeRequestNotes200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeSingeMergeRequestNotes200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

