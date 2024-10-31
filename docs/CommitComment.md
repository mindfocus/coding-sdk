# CommitComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**CodingUser**](CodingUser.md) |  | [optional] 
**CommitSha** | Pointer to **string** | 提交 ID | [optional] [default to ""]
**Content** | Pointer to **string** | 评论内容 | [optional] [default to ""]
**CreatedAt** | Pointer to **int64** | 创建时间 | [optional] 
**DepotId** | Pointer to **int64** | 仓库 ID | [optional] 
**Id** | Pointer to **int64** | 提交评论 ID | [optional] 
**Index** | Pointer to **int64** | git diff 信息的第几行 | [optional] 
**Path** | Pointer to **string** | 评论的文件路径 | [optional] [default to ""]

## Methods

### NewCommitComment

`func NewCommitComment() *CommitComment`

NewCommitComment instantiates a new CommitComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitCommentWithDefaults

`func NewCommitCommentWithDefaults() *CommitComment`

NewCommitCommentWithDefaults instantiates a new CommitComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *CommitComment) GetAuthor() CodingUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CommitComment) GetAuthorOk() (*CodingUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CommitComment) SetAuthor(v CodingUser)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *CommitComment) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCommitSha

`func (o *CommitComment) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *CommitComment) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *CommitComment) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *CommitComment) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### GetContent

`func (o *CommitComment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CommitComment) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CommitComment) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CommitComment) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CommitComment) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CommitComment) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CommitComment) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CommitComment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDepotId

`func (o *CommitComment) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CommitComment) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CommitComment) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.

### HasDepotId

`func (o *CommitComment) HasDepotId() bool`

HasDepotId returns a boolean if a field has been set.

### GetId

`func (o *CommitComment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommitComment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommitComment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CommitComment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *CommitComment) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *CommitComment) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *CommitComment) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *CommitComment) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetPath

`func (o *CommitComment) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CommitComment) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CommitComment) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CommitComment) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


