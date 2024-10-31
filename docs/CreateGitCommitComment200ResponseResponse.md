# CreateGitCommitComment200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to [**CommitComment**](CommitComment.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewCreateGitCommitComment200ResponseResponse

`func NewCreateGitCommitComment200ResponseResponse() *CreateGitCommitComment200ResponseResponse`

NewCreateGitCommitComment200ResponseResponse instantiates a new CreateGitCommitComment200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGitCommitComment200ResponseResponseWithDefaults

`func NewCreateGitCommitComment200ResponseResponseWithDefaults() *CreateGitCommitComment200ResponseResponse`

NewCreateGitCommitComment200ResponseResponseWithDefaults instantiates a new CreateGitCommitComment200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *CreateGitCommitComment200ResponseResponse) GetComment() CommitComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateGitCommitComment200ResponseResponse) GetCommentOk() (*CommitComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateGitCommitComment200ResponseResponse) SetComment(v CommitComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateGitCommitComment200ResponseResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetRequestId

`func (o *CreateGitCommitComment200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateGitCommitComment200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateGitCommitComment200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateGitCommitComment200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


