# DescribeGitMergeRequestDiffDetail200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | Pointer to [**GitDiff**](GitDiff.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitMergeRequestDiffDetail200ResponseResponse

`func NewDescribeGitMergeRequestDiffDetail200ResponseResponse() *DescribeGitMergeRequestDiffDetail200ResponseResponse`

NewDescribeGitMergeRequestDiffDetail200ResponseResponse instantiates a new DescribeGitMergeRequestDiffDetail200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitMergeRequestDiffDetail200ResponseResponseWithDefaults

`func NewDescribeGitMergeRequestDiffDetail200ResponseResponseWithDefaults() *DescribeGitMergeRequestDiffDetail200ResponseResponse`

NewDescribeGitMergeRequestDiffDetail200ResponseResponseWithDefaults instantiates a new DescribeGitMergeRequestDiffDetail200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *DescribeGitMergeRequestDiffDetail200ResponseResponse) GetDetail() GitDiff`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *DescribeGitMergeRequestDiffDetail200ResponseResponse) GetDetailOk() (*GitDiff, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *DescribeGitMergeRequestDiffDetail200ResponseResponse) SetDetail(v GitDiff)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *DescribeGitMergeRequestDiffDetail200ResponseResponse) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeGitMergeRequestDiffDetail200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitMergeRequestDiffDetail200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitMergeRequestDiffDetail200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitMergeRequestDiffDetail200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


