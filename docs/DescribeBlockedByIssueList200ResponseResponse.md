# DescribeBlockedByIssueList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issues** | Pointer to [**[]IssueSimpleData**](IssueSimpleData.md) | 前置事项列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeBlockedByIssueList200ResponseResponse

`func NewDescribeBlockedByIssueList200ResponseResponse() *DescribeBlockedByIssueList200ResponseResponse`

NewDescribeBlockedByIssueList200ResponseResponse instantiates a new DescribeBlockedByIssueList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeBlockedByIssueList200ResponseResponseWithDefaults

`func NewDescribeBlockedByIssueList200ResponseResponseWithDefaults() *DescribeBlockedByIssueList200ResponseResponse`

NewDescribeBlockedByIssueList200ResponseResponseWithDefaults instantiates a new DescribeBlockedByIssueList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssues

`func (o *DescribeBlockedByIssueList200ResponseResponse) GetIssues() []IssueSimpleData`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *DescribeBlockedByIssueList200ResponseResponse) GetIssuesOk() (*[]IssueSimpleData, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *DescribeBlockedByIssueList200ResponseResponse) SetIssues(v []IssueSimpleData)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *DescribeBlockedByIssueList200ResponseResponse) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeBlockedByIssueList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeBlockedByIssueList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeBlockedByIssueList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeBlockedByIssueList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


