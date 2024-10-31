# DescribeGitMergeRequestsBySha200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]MergeRequestDetail**](MergeRequestDetail.md) | 合并请求详情 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitMergeRequestsBySha200ResponseResponse

`func NewDescribeGitMergeRequestsBySha200ResponseResponse() *DescribeGitMergeRequestsBySha200ResponseResponse`

NewDescribeGitMergeRequestsBySha200ResponseResponse instantiates a new DescribeGitMergeRequestsBySha200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitMergeRequestsBySha200ResponseResponseWithDefaults

`func NewDescribeGitMergeRequestsBySha200ResponseResponseWithDefaults() *DescribeGitMergeRequestsBySha200ResponseResponse`

NewDescribeGitMergeRequestsBySha200ResponseResponseWithDefaults instantiates a new DescribeGitMergeRequestsBySha200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *DescribeGitMergeRequestsBySha200ResponseResponse) GetDetails() []MergeRequestDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DescribeGitMergeRequestsBySha200ResponseResponse) GetDetailsOk() (*[]MergeRequestDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DescribeGitMergeRequestsBySha200ResponseResponse) SetDetails(v []MergeRequestDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *DescribeGitMergeRequestsBySha200ResponseResponse) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeGitMergeRequestsBySha200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitMergeRequestsBySha200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitMergeRequestsBySha200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitMergeRequestsBySha200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


