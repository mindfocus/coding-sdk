# MergeRequestNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**CodingUser**](CodingUser.md) |  | [optional] 
**CommitSha** | Pointer to **NullableString** | 合并请求中提交Id | [optional] [default to ""]
**Content** | Pointer to **string** | 行评论的内容 | [optional] [default to ""]
**CreatedAt** | Pointer to **int32** | 行评论创建时间 | [optional] 
**Id** | Pointer to **int32** | 行评论的Id | [optional] 
**Index** | Pointer to **NullableInt32** | diff信息中的行数 | [optional] 
**MergeId** | Pointer to **int32** | 合并请求的Iid | [optional] 
**ParentId** | Pointer to **int32** | 子评论的父Id | [optional] 
**Path** | Pointer to **NullableString** | 改动文件的路径 | [optional] [default to ""]
**UpdatedAt** | Pointer to **int32** | 行评论更新时间 | [optional] 

## Methods

### NewMergeRequestNote

`func NewMergeRequestNote() *MergeRequestNote`

NewMergeRequestNote instantiates a new MergeRequestNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeRequestNoteWithDefaults

`func NewMergeRequestNoteWithDefaults() *MergeRequestNote`

NewMergeRequestNoteWithDefaults instantiates a new MergeRequestNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *MergeRequestNote) GetAuthor() CodingUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *MergeRequestNote) GetAuthorOk() (*CodingUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *MergeRequestNote) SetAuthor(v CodingUser)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *MergeRequestNote) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCommitSha

`func (o *MergeRequestNote) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *MergeRequestNote) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *MergeRequestNote) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *MergeRequestNote) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### SetCommitShaNil

`func (o *MergeRequestNote) SetCommitShaNil(b bool)`

 SetCommitShaNil sets the value for CommitSha to be an explicit nil

### UnsetCommitSha
`func (o *MergeRequestNote) UnsetCommitSha()`

UnsetCommitSha ensures that no value is present for CommitSha, not even an explicit nil
### GetContent

`func (o *MergeRequestNote) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MergeRequestNote) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MergeRequestNote) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MergeRequestNote) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MergeRequestNote) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MergeRequestNote) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MergeRequestNote) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MergeRequestNote) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *MergeRequestNote) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MergeRequestNote) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MergeRequestNote) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MergeRequestNote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *MergeRequestNote) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MergeRequestNote) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *MergeRequestNote) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *MergeRequestNote) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *MergeRequestNote) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *MergeRequestNote) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil
### GetMergeId

`func (o *MergeRequestNote) GetMergeId() int32`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *MergeRequestNote) GetMergeIdOk() (*int32, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *MergeRequestNote) SetMergeId(v int32)`

SetMergeId sets MergeId field to given value.

### HasMergeId

`func (o *MergeRequestNote) HasMergeId() bool`

HasMergeId returns a boolean if a field has been set.

### GetParentId

`func (o *MergeRequestNote) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *MergeRequestNote) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *MergeRequestNote) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *MergeRequestNote) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPath

`func (o *MergeRequestNote) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MergeRequestNote) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MergeRequestNote) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MergeRequestNote) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *MergeRequestNote) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *MergeRequestNote) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetUpdatedAt

`func (o *MergeRequestNote) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MergeRequestNote) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MergeRequestNote) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MergeRequestNote) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


