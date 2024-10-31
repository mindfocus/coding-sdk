# DescribeGitCommitComments200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitCommitComments** | Pointer to [**GitCommitComment**](GitCommitComment.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitCommitComments200ResponseResponse

`func NewDescribeGitCommitComments200ResponseResponse() *DescribeGitCommitComments200ResponseResponse`

NewDescribeGitCommitComments200ResponseResponse instantiates a new DescribeGitCommitComments200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitCommitComments200ResponseResponseWithDefaults

`func NewDescribeGitCommitComments200ResponseResponseWithDefaults() *DescribeGitCommitComments200ResponseResponse`

NewDescribeGitCommitComments200ResponseResponseWithDefaults instantiates a new DescribeGitCommitComments200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitCommitComments

`func (o *DescribeGitCommitComments200ResponseResponse) GetGitCommitComments() GitCommitComment`

GetGitCommitComments returns the GitCommitComments field if non-nil, zero value otherwise.

### GetGitCommitCommentsOk

`func (o *DescribeGitCommitComments200ResponseResponse) GetGitCommitCommentsOk() (*GitCommitComment, bool)`

GetGitCommitCommentsOk returns a tuple with the GitCommitComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitComments

`func (o *DescribeGitCommitComments200ResponseResponse) SetGitCommitComments(v GitCommitComment)`

SetGitCommitComments sets GitCommitComments field to given value.

### HasGitCommitComments

`func (o *DescribeGitCommitComments200ResponseResponse) HasGitCommitComments() bool`

HasGitCommitComments returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeGitCommitComments200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitCommitComments200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitCommitComments200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitCommitComments200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


