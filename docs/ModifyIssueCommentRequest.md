# ModifyIssueCommentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommentId** | **int64** | 评论 ID | 
**Content** | **string** | 评论内容 | 
**IssueCode** | **int64** | 事项 Code | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewModifyIssueCommentRequest

`func NewModifyIssueCommentRequest(commentId int64, content string, issueCode int64, projectName string, ) *ModifyIssueCommentRequest`

NewModifyIssueCommentRequest instantiates a new ModifyIssueCommentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyIssueCommentRequestWithDefaults

`func NewModifyIssueCommentRequestWithDefaults() *ModifyIssueCommentRequest`

NewModifyIssueCommentRequestWithDefaults instantiates a new ModifyIssueCommentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommentId

`func (o *ModifyIssueCommentRequest) GetCommentId() int64`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *ModifyIssueCommentRequest) GetCommentIdOk() (*int64, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *ModifyIssueCommentRequest) SetCommentId(v int64)`

SetCommentId sets CommentId field to given value.


### GetContent

`func (o *ModifyIssueCommentRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ModifyIssueCommentRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ModifyIssueCommentRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetIssueCode

`func (o *ModifyIssueCommentRequest) GetIssueCode() int64`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *ModifyIssueCommentRequest) GetIssueCodeOk() (*int64, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *ModifyIssueCommentRequest) SetIssueCode(v int64)`

SetIssueCode sets IssueCode field to given value.


### GetProjectName

`func (o *ModifyIssueCommentRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ModifyIssueCommentRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ModifyIssueCommentRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


