# CreateIssueCommentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | 评论内容 | 
**IssueCode** | **int64** | 事项 Code | 
**ParentId** | **int64** | 父评论ID | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewCreateIssueCommentRequest

`func NewCreateIssueCommentRequest(content string, issueCode int64, parentId int64, projectName string, ) *CreateIssueCommentRequest`

NewCreateIssueCommentRequest instantiates a new CreateIssueCommentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIssueCommentRequestWithDefaults

`func NewCreateIssueCommentRequestWithDefaults() *CreateIssueCommentRequest`

NewCreateIssueCommentRequestWithDefaults instantiates a new CreateIssueCommentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CreateIssueCommentRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateIssueCommentRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateIssueCommentRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetIssueCode

`func (o *CreateIssueCommentRequest) GetIssueCode() int64`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *CreateIssueCommentRequest) GetIssueCodeOk() (*int64, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *CreateIssueCommentRequest) SetIssueCode(v int64)`

SetIssueCode sets IssueCode field to given value.


### GetParentId

`func (o *CreateIssueCommentRequest) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateIssueCommentRequest) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateIssueCommentRequest) SetParentId(v int64)`

SetParentId sets ParentId field to given value.


### GetProjectName

`func (o *CreateIssueCommentRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateIssueCommentRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateIssueCommentRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


