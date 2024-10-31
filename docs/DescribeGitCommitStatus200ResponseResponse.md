# DescribeGitCommitStatus200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCheckResults** | Pointer to [**[]StatusCheckResult**](StatusCheckResult.md) |  提交流水线状态列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitCommitStatus200ResponseResponse

`func NewDescribeGitCommitStatus200ResponseResponse() *DescribeGitCommitStatus200ResponseResponse`

NewDescribeGitCommitStatus200ResponseResponse instantiates a new DescribeGitCommitStatus200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitCommitStatus200ResponseResponseWithDefaults

`func NewDescribeGitCommitStatus200ResponseResponseWithDefaults() *DescribeGitCommitStatus200ResponseResponse`

NewDescribeGitCommitStatus200ResponseResponseWithDefaults instantiates a new DescribeGitCommitStatus200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCheckResults

`func (o *DescribeGitCommitStatus200ResponseResponse) GetStatusCheckResults() []StatusCheckResult`

GetStatusCheckResults returns the StatusCheckResults field if non-nil, zero value otherwise.

### GetStatusCheckResultsOk

`func (o *DescribeGitCommitStatus200ResponseResponse) GetStatusCheckResultsOk() (*[]StatusCheckResult, bool)`

GetStatusCheckResultsOk returns a tuple with the StatusCheckResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCheckResults

`func (o *DescribeGitCommitStatus200ResponseResponse) SetStatusCheckResults(v []StatusCheckResult)`

SetStatusCheckResults sets StatusCheckResults field to given value.

### HasStatusCheckResults

`func (o *DescribeGitCommitStatus200ResponseResponse) HasStatusCheckResults() bool`

HasStatusCheckResults returns a boolean if a field has been set.

### SetStatusCheckResultsNil

`func (o *DescribeGitCommitStatus200ResponseResponse) SetStatusCheckResultsNil(b bool)`

 SetStatusCheckResultsNil sets the value for StatusCheckResults to be an explicit nil

### UnsetStatusCheckResults
`func (o *DescribeGitCommitStatus200ResponseResponse) UnsetStatusCheckResults()`

UnsetStatusCheckResults ensures that no value is present for StatusCheckResults, not even an explicit nil
### GetRequestId

`func (o *DescribeGitCommitStatus200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitCommitStatus200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitCommitStatus200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitCommitStatus200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


