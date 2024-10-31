# DescribeIssueCustomFieldLogList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldChangeLogList** | Pointer to [**[]CustomFieldChangeLog**](CustomFieldChangeLog.md) | 事项的自定义属性变更日志 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeIssueCustomFieldLogList200ResponseResponse

`func NewDescribeIssueCustomFieldLogList200ResponseResponse() *DescribeIssueCustomFieldLogList200ResponseResponse`

NewDescribeIssueCustomFieldLogList200ResponseResponse instantiates a new DescribeIssueCustomFieldLogList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIssueCustomFieldLogList200ResponseResponseWithDefaults

`func NewDescribeIssueCustomFieldLogList200ResponseResponseWithDefaults() *DescribeIssueCustomFieldLogList200ResponseResponse`

NewDescribeIssueCustomFieldLogList200ResponseResponseWithDefaults instantiates a new DescribeIssueCustomFieldLogList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldChangeLogList

`func (o *DescribeIssueCustomFieldLogList200ResponseResponse) GetFieldChangeLogList() []CustomFieldChangeLog`

GetFieldChangeLogList returns the FieldChangeLogList field if non-nil, zero value otherwise.

### GetFieldChangeLogListOk

`func (o *DescribeIssueCustomFieldLogList200ResponseResponse) GetFieldChangeLogListOk() (*[]CustomFieldChangeLog, bool)`

GetFieldChangeLogListOk returns a tuple with the FieldChangeLogList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldChangeLogList

`func (o *DescribeIssueCustomFieldLogList200ResponseResponse) SetFieldChangeLogList(v []CustomFieldChangeLog)`

SetFieldChangeLogList sets FieldChangeLogList field to given value.

### HasFieldChangeLogList

`func (o *DescribeIssueCustomFieldLogList200ResponseResponse) HasFieldChangeLogList() bool`

HasFieldChangeLogList returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeIssueCustomFieldLogList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeIssueCustomFieldLogList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeIssueCustomFieldLogList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeIssueCustomFieldLogList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


