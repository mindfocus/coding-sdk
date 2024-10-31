# DescribeServiceHookLogs200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ServiceHookLogPage**](ServiceHookLogPage.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeServiceHookLogs200ResponseResponse

`func NewDescribeServiceHookLogs200ResponseResponse() *DescribeServiceHookLogs200ResponseResponse`

NewDescribeServiceHookLogs200ResponseResponse instantiates a new DescribeServiceHookLogs200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeServiceHookLogs200ResponseResponseWithDefaults

`func NewDescribeServiceHookLogs200ResponseResponseWithDefaults() *DescribeServiceHookLogs200ResponseResponse`

NewDescribeServiceHookLogs200ResponseResponseWithDefaults instantiates a new DescribeServiceHookLogs200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DescribeServiceHookLogs200ResponseResponse) GetData() ServiceHookLogPage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DescribeServiceHookLogs200ResponseResponse) GetDataOk() (*ServiceHookLogPage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DescribeServiceHookLogs200ResponseResponse) SetData(v ServiceHookLogPage)`

SetData sets Data field to given value.

### HasData

`func (o *DescribeServiceHookLogs200ResponseResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeServiceHookLogs200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeServiceHookLogs200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeServiceHookLogs200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeServiceHookLogs200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


