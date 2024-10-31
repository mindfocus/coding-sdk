# DescribeIssueStatusChangeLogList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**IssueStatusChangeLogList**](IssueStatusChangeLogList.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeIssueStatusChangeLogList200ResponseResponse

`func NewDescribeIssueStatusChangeLogList200ResponseResponse() *DescribeIssueStatusChangeLogList200ResponseResponse`

NewDescribeIssueStatusChangeLogList200ResponseResponse instantiates a new DescribeIssueStatusChangeLogList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIssueStatusChangeLogList200ResponseResponseWithDefaults

`func NewDescribeIssueStatusChangeLogList200ResponseResponseWithDefaults() *DescribeIssueStatusChangeLogList200ResponseResponse`

NewDescribeIssueStatusChangeLogList200ResponseResponseWithDefaults instantiates a new DescribeIssueStatusChangeLogList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *DescribeIssueStatusChangeLogList200ResponseResponse) GetLogs() IssueStatusChangeLogList`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *DescribeIssueStatusChangeLogList200ResponseResponse) GetLogsOk() (*IssueStatusChangeLogList, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *DescribeIssueStatusChangeLogList200ResponseResponse) SetLogs(v IssueStatusChangeLogList)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *DescribeIssueStatusChangeLogList200ResponseResponse) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeIssueStatusChangeLogList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeIssueStatusChangeLogList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeIssueStatusChangeLogList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeIssueStatusChangeLogList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


