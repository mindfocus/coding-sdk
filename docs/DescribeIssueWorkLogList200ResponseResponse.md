# DescribeIssueWorkLogList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkLogs** | Pointer to [**[]IssueWorkLog**](IssueWorkLog.md) | 工时日志列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeIssueWorkLogList200ResponseResponse

`func NewDescribeIssueWorkLogList200ResponseResponse() *DescribeIssueWorkLogList200ResponseResponse`

NewDescribeIssueWorkLogList200ResponseResponse instantiates a new DescribeIssueWorkLogList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIssueWorkLogList200ResponseResponseWithDefaults

`func NewDescribeIssueWorkLogList200ResponseResponseWithDefaults() *DescribeIssueWorkLogList200ResponseResponse`

NewDescribeIssueWorkLogList200ResponseResponseWithDefaults instantiates a new DescribeIssueWorkLogList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkLogs

`func (o *DescribeIssueWorkLogList200ResponseResponse) GetWorkLogs() []IssueWorkLog`

GetWorkLogs returns the WorkLogs field if non-nil, zero value otherwise.

### GetWorkLogsOk

`func (o *DescribeIssueWorkLogList200ResponseResponse) GetWorkLogsOk() (*[]IssueWorkLog, bool)`

GetWorkLogsOk returns a tuple with the WorkLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkLogs

`func (o *DescribeIssueWorkLogList200ResponseResponse) SetWorkLogs(v []IssueWorkLog)`

SetWorkLogs sets WorkLogs field to given value.

### HasWorkLogs

`func (o *DescribeIssueWorkLogList200ResponseResponse) HasWorkLogs() bool`

HasWorkLogs returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeIssueWorkLogList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeIssueWorkLogList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeIssueWorkLogList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeIssueWorkLogList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


