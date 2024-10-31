# DescribeGitCommitDiff200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Diffs** | Pointer to [**[]GitDiff**](GitDiff.md) | diff信息详情 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitCommitDiff200ResponseResponse

`func NewDescribeGitCommitDiff200ResponseResponse() *DescribeGitCommitDiff200ResponseResponse`

NewDescribeGitCommitDiff200ResponseResponse instantiates a new DescribeGitCommitDiff200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitCommitDiff200ResponseResponseWithDefaults

`func NewDescribeGitCommitDiff200ResponseResponseWithDefaults() *DescribeGitCommitDiff200ResponseResponse`

NewDescribeGitCommitDiff200ResponseResponseWithDefaults instantiates a new DescribeGitCommitDiff200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiffs

`func (o *DescribeGitCommitDiff200ResponseResponse) GetDiffs() []GitDiff`

GetDiffs returns the Diffs field if non-nil, zero value otherwise.

### GetDiffsOk

`func (o *DescribeGitCommitDiff200ResponseResponse) GetDiffsOk() (*[]GitDiff, bool)`

GetDiffsOk returns a tuple with the Diffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffs

`func (o *DescribeGitCommitDiff200ResponseResponse) SetDiffs(v []GitDiff)`

SetDiffs sets Diffs field to given value.

### HasDiffs

`func (o *DescribeGitCommitDiff200ResponseResponse) HasDiffs() bool`

HasDiffs returns a boolean if a field has been set.

### SetDiffsNil

`func (o *DescribeGitCommitDiff200ResponseResponse) SetDiffsNil(b bool)`

 SetDiffsNil sets the value for Diffs to be an explicit nil

### UnsetDiffs
`func (o *DescribeGitCommitDiff200ResponseResponse) UnsetDiffs()`

UnsetDiffs ensures that no value is present for Diffs, not even an explicit nil
### GetRequestId

`func (o *DescribeGitCommitDiff200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitCommitDiff200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitCommitDiff200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitCommitDiff200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


