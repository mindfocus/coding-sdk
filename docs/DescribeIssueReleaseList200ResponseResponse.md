# DescribeIssueReleaseList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseList** | Pointer to [**[]OpenApiRelease**](OpenApiRelease.md) | 版本列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeIssueReleaseList200ResponseResponse

`func NewDescribeIssueReleaseList200ResponseResponse() *DescribeIssueReleaseList200ResponseResponse`

NewDescribeIssueReleaseList200ResponseResponse instantiates a new DescribeIssueReleaseList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIssueReleaseList200ResponseResponseWithDefaults

`func NewDescribeIssueReleaseList200ResponseResponseWithDefaults() *DescribeIssueReleaseList200ResponseResponse`

NewDescribeIssueReleaseList200ResponseResponseWithDefaults instantiates a new DescribeIssueReleaseList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseList

`func (o *DescribeIssueReleaseList200ResponseResponse) GetReleaseList() []OpenApiRelease`

GetReleaseList returns the ReleaseList field if non-nil, zero value otherwise.

### GetReleaseListOk

`func (o *DescribeIssueReleaseList200ResponseResponse) GetReleaseListOk() (*[]OpenApiRelease, bool)`

GetReleaseListOk returns a tuple with the ReleaseList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseList

`func (o *DescribeIssueReleaseList200ResponseResponse) SetReleaseList(v []OpenApiRelease)`

SetReleaseList sets ReleaseList field to given value.

### HasReleaseList

`func (o *DescribeIssueReleaseList200ResponseResponse) HasReleaseList() bool`

HasReleaseList returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeIssueReleaseList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeIssueReleaseList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeIssueReleaseList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeIssueReleaseList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


