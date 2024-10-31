# DescribeIssueList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueList** | Pointer to [**[]IssueListData**](IssueListData.md) | 事项列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeIssueList200ResponseResponse

`func NewDescribeIssueList200ResponseResponse() *DescribeIssueList200ResponseResponse`

NewDescribeIssueList200ResponseResponse instantiates a new DescribeIssueList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIssueList200ResponseResponseWithDefaults

`func NewDescribeIssueList200ResponseResponseWithDefaults() *DescribeIssueList200ResponseResponse`

NewDescribeIssueList200ResponseResponseWithDefaults instantiates a new DescribeIssueList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueList

`func (o *DescribeIssueList200ResponseResponse) GetIssueList() []IssueListData`

GetIssueList returns the IssueList field if non-nil, zero value otherwise.

### GetIssueListOk

`func (o *DescribeIssueList200ResponseResponse) GetIssueListOk() (*[]IssueListData, bool)`

GetIssueListOk returns a tuple with the IssueList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueList

`func (o *DescribeIssueList200ResponseResponse) SetIssueList(v []IssueListData)`

SetIssueList sets IssueList field to given value.

### HasIssueList

`func (o *DescribeIssueList200ResponseResponse) HasIssueList() bool`

HasIssueList returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeIssueList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeIssueList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeIssueList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeIssueList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


