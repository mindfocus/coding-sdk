# DescribeIssueCommentList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommentList** | Pointer to [**[]IssueComment**](IssueComment.md) | 事项评论列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeIssueCommentList200ResponseResponse

`func NewDescribeIssueCommentList200ResponseResponse() *DescribeIssueCommentList200ResponseResponse`

NewDescribeIssueCommentList200ResponseResponse instantiates a new DescribeIssueCommentList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIssueCommentList200ResponseResponseWithDefaults

`func NewDescribeIssueCommentList200ResponseResponseWithDefaults() *DescribeIssueCommentList200ResponseResponse`

NewDescribeIssueCommentList200ResponseResponseWithDefaults instantiates a new DescribeIssueCommentList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommentList

`func (o *DescribeIssueCommentList200ResponseResponse) GetCommentList() []IssueComment`

GetCommentList returns the CommentList field if non-nil, zero value otherwise.

### GetCommentListOk

`func (o *DescribeIssueCommentList200ResponseResponse) GetCommentListOk() (*[]IssueComment, bool)`

GetCommentListOk returns a tuple with the CommentList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentList

`func (o *DescribeIssueCommentList200ResponseResponse) SetCommentList(v []IssueComment)`

SetCommentList sets CommentList field to given value.

### HasCommentList

`func (o *DescribeIssueCommentList200ResponseResponse) HasCommentList() bool`

HasCommentList returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeIssueCommentList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeIssueCommentList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeIssueCommentList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeIssueCommentList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


