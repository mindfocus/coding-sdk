# IssueComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommentId** | Pointer to **int64** | 评论 ID | [optional] 
**Content** | Pointer to **string** | 解析后的内容 | [optional] [default to ""]
**CreatedAt** | Pointer to **int64** | 创建时间戳 | [optional] 
**CreatorId** | Pointer to **int64** | 创建人 ID | [optional] 
**ParentId** | Pointer to **int64** | 父评论 ID | [optional] 
**RawContent** | Pointer to **string** | 内容 | [optional] [default to ""]
**UpdatedAt** | Pointer to **int64** | 更新时间戳 | [optional] 

## Methods

### NewIssueComment

`func NewIssueComment() *IssueComment`

NewIssueComment instantiates a new IssueComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueCommentWithDefaults

`func NewIssueCommentWithDefaults() *IssueComment`

NewIssueCommentWithDefaults instantiates a new IssueComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommentId

`func (o *IssueComment) GetCommentId() int64`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *IssueComment) GetCommentIdOk() (*int64, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *IssueComment) SetCommentId(v int64)`

SetCommentId sets CommentId field to given value.

### HasCommentId

`func (o *IssueComment) HasCommentId() bool`

HasCommentId returns a boolean if a field has been set.

### GetContent

`func (o *IssueComment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IssueComment) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IssueComment) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *IssueComment) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IssueComment) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueComment) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueComment) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssueComment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatorId

`func (o *IssueComment) GetCreatorId() int64`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *IssueComment) GetCreatorIdOk() (*int64, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *IssueComment) SetCreatorId(v int64)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *IssueComment) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetParentId

`func (o *IssueComment) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *IssueComment) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *IssueComment) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *IssueComment) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetRawContent

`func (o *IssueComment) GetRawContent() string`

GetRawContent returns the RawContent field if non-nil, zero value otherwise.

### GetRawContentOk

`func (o *IssueComment) GetRawContentOk() (*string, bool)`

GetRawContentOk returns a tuple with the RawContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawContent

`func (o *IssueComment) SetRawContent(v string)`

SetRawContent sets RawContent field to given value.

### HasRawContent

`func (o *IssueComment) HasRawContent() bool`

HasRawContent returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IssueComment) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssueComment) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssueComment) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IssueComment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


