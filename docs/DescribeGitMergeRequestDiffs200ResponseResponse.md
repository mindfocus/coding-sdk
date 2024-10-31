# DescribeGitMergeRequestDiffs200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Diff** | Pointer to [**MergeRequestDiff**](MergeRequestDiff.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitMergeRequestDiffs200ResponseResponse

`func NewDescribeGitMergeRequestDiffs200ResponseResponse() *DescribeGitMergeRequestDiffs200ResponseResponse`

NewDescribeGitMergeRequestDiffs200ResponseResponse instantiates a new DescribeGitMergeRequestDiffs200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitMergeRequestDiffs200ResponseResponseWithDefaults

`func NewDescribeGitMergeRequestDiffs200ResponseResponseWithDefaults() *DescribeGitMergeRequestDiffs200ResponseResponse`

NewDescribeGitMergeRequestDiffs200ResponseResponseWithDefaults instantiates a new DescribeGitMergeRequestDiffs200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiff

`func (o *DescribeGitMergeRequestDiffs200ResponseResponse) GetDiff() MergeRequestDiff`

GetDiff returns the Diff field if non-nil, zero value otherwise.

### GetDiffOk

`func (o *DescribeGitMergeRequestDiffs200ResponseResponse) GetDiffOk() (*MergeRequestDiff, bool)`

GetDiffOk returns a tuple with the Diff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiff

`func (o *DescribeGitMergeRequestDiffs200ResponseResponse) SetDiff(v MergeRequestDiff)`

SetDiff sets Diff field to given value.

### HasDiff

`func (o *DescribeGitMergeRequestDiffs200ResponseResponse) HasDiff() bool`

HasDiff returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeGitMergeRequestDiffs200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitMergeRequestDiffs200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitMergeRequestDiffs200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitMergeRequestDiffs200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


