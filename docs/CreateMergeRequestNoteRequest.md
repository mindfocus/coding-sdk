# CreateMergeRequestNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | 行评论内容 | 
**DepotPath** | **string** | 仓库路径 | 
**Form** | Pointer to [**MergeRequestDiffNoteForm**](MergeRequestDiffNoteForm.md) |  | [optional] 
**MergeId** | **int32** | 合并请求的Iid | 
**ParentId** | **int32** | 行评论的父亲Id | 

## Methods

### NewCreateMergeRequestNoteRequest

`func NewCreateMergeRequestNoteRequest(content string, depotPath string, mergeId int32, parentId int32, ) *CreateMergeRequestNoteRequest`

NewCreateMergeRequestNoteRequest instantiates a new CreateMergeRequestNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMergeRequestNoteRequestWithDefaults

`func NewCreateMergeRequestNoteRequestWithDefaults() *CreateMergeRequestNoteRequest`

NewCreateMergeRequestNoteRequestWithDefaults instantiates a new CreateMergeRequestNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CreateMergeRequestNoteRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateMergeRequestNoteRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateMergeRequestNoteRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetDepotPath

`func (o *CreateMergeRequestNoteRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateMergeRequestNoteRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateMergeRequestNoteRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetForm

`func (o *CreateMergeRequestNoteRequest) GetForm() MergeRequestDiffNoteForm`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *CreateMergeRequestNoteRequest) GetFormOk() (*MergeRequestDiffNoteForm, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *CreateMergeRequestNoteRequest) SetForm(v MergeRequestDiffNoteForm)`

SetForm sets Form field to given value.

### HasForm

`func (o *CreateMergeRequestNoteRequest) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetMergeId

`func (o *CreateMergeRequestNoteRequest) GetMergeId() int32`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *CreateMergeRequestNoteRequest) GetMergeIdOk() (*int32, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *CreateMergeRequestNoteRequest) SetMergeId(v int32)`

SetMergeId sets MergeId field to given value.


### GetParentId

`func (o *CreateMergeRequestNoteRequest) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateMergeRequestNoteRequest) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateMergeRequestNoteRequest) SetParentId(v int32)`

SetParentId sets ParentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


