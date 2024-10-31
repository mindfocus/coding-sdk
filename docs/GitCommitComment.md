# GitCommitComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitComments** | Pointer to [**[]CommitComment**](CommitComment.md) | 提交评论详细信息 | [optional] 
**Page** | Pointer to [**PageInfo**](PageInfo.md) |  | [optional] 

## Methods

### NewGitCommitComment

`func NewGitCommitComment() *GitCommitComment`

NewGitCommitComment instantiates a new GitCommitComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCommitCommentWithDefaults

`func NewGitCommitCommentWithDefaults() *GitCommitComment`

NewGitCommitCommentWithDefaults instantiates a new GitCommitComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitComments

`func (o *GitCommitComment) GetCommitComments() []CommitComment`

GetCommitComments returns the CommitComments field if non-nil, zero value otherwise.

### GetCommitCommentsOk

`func (o *GitCommitComment) GetCommitCommentsOk() (*[]CommitComment, bool)`

GetCommitCommentsOk returns a tuple with the CommitComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitComments

`func (o *GitCommitComment) SetCommitComments(v []CommitComment)`

SetCommitComments sets CommitComments field to given value.

### HasCommitComments

`func (o *GitCommitComment) HasCommitComments() bool`

HasCommitComments returns a boolean if a field has been set.

### GetPage

`func (o *GitCommitComment) GetPage() PageInfo`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GitCommitComment) GetPageOk() (*PageInfo, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GitCommitComment) SetPage(v PageInfo)`

SetPage sets Page field to given value.

### HasPage

`func (o *GitCommitComment) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


