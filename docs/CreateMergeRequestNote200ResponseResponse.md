# CreateMergeRequestNote200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Note** | Pointer to [**MergeRequestNote**](MergeRequestNote.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewCreateMergeRequestNote200ResponseResponse

`func NewCreateMergeRequestNote200ResponseResponse() *CreateMergeRequestNote200ResponseResponse`

NewCreateMergeRequestNote200ResponseResponse instantiates a new CreateMergeRequestNote200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMergeRequestNote200ResponseResponseWithDefaults

`func NewCreateMergeRequestNote200ResponseResponseWithDefaults() *CreateMergeRequestNote200ResponseResponse`

NewCreateMergeRequestNote200ResponseResponseWithDefaults instantiates a new CreateMergeRequestNote200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNote

`func (o *CreateMergeRequestNote200ResponseResponse) GetNote() MergeRequestNote`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateMergeRequestNote200ResponseResponse) GetNoteOk() (*MergeRequestNote, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateMergeRequestNote200ResponseResponse) SetNote(v MergeRequestNote)`

SetNote sets Note field to given value.

### HasNote

`func (o *CreateMergeRequestNote200ResponseResponse) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetRequestId

`func (o *CreateMergeRequestNote200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateMergeRequestNote200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateMergeRequestNote200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateMergeRequestNote200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


