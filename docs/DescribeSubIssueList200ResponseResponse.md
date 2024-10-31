# DescribeSubIssueList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubIssueList** | Pointer to [**[]IssueListData**](IssueListData.md) | 子事项列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeSubIssueList200ResponseResponse

`func NewDescribeSubIssueList200ResponseResponse() *DescribeSubIssueList200ResponseResponse`

NewDescribeSubIssueList200ResponseResponse instantiates a new DescribeSubIssueList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeSubIssueList200ResponseResponseWithDefaults

`func NewDescribeSubIssueList200ResponseResponseWithDefaults() *DescribeSubIssueList200ResponseResponse`

NewDescribeSubIssueList200ResponseResponseWithDefaults instantiates a new DescribeSubIssueList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubIssueList

`func (o *DescribeSubIssueList200ResponseResponse) GetSubIssueList() []IssueListData`

GetSubIssueList returns the SubIssueList field if non-nil, zero value otherwise.

### GetSubIssueListOk

`func (o *DescribeSubIssueList200ResponseResponse) GetSubIssueListOk() (*[]IssueListData, bool)`

GetSubIssueListOk returns a tuple with the SubIssueList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubIssueList

`func (o *DescribeSubIssueList200ResponseResponse) SetSubIssueList(v []IssueListData)`

SetSubIssueList sets SubIssueList field to given value.

### HasSubIssueList

`func (o *DescribeSubIssueList200ResponseResponse) HasSubIssueList() bool`

HasSubIssueList returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeSubIssueList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeSubIssueList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeSubIssueList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeSubIssueList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


