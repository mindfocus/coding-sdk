# DescribeGitMergeBase200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | Pointer to [**GitCommit**](GitCommit.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitMergeBase200ResponseResponse

`func NewDescribeGitMergeBase200ResponseResponse() *DescribeGitMergeBase200ResponseResponse`

NewDescribeGitMergeBase200ResponseResponse instantiates a new DescribeGitMergeBase200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitMergeBase200ResponseResponseWithDefaults

`func NewDescribeGitMergeBase200ResponseResponseWithDefaults() *DescribeGitMergeBase200ResponseResponse`

NewDescribeGitMergeBase200ResponseResponseWithDefaults instantiates a new DescribeGitMergeBase200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *DescribeGitMergeBase200ResponseResponse) GetCommit() GitCommit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *DescribeGitMergeBase200ResponseResponse) GetCommitOk() (*GitCommit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *DescribeGitMergeBase200ResponseResponse) SetCommit(v GitCommit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *DescribeGitMergeBase200ResponseResponse) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeGitMergeBase200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitMergeBase200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitMergeBase200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitMergeBase200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


