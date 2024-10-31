# DescribeMergeRequestLog200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]MergeRequestLog**](MergeRequestLog.md) | 操作记录列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeMergeRequestLog200ResponseResponse

`func NewDescribeMergeRequestLog200ResponseResponse() *DescribeMergeRequestLog200ResponseResponse`

NewDescribeMergeRequestLog200ResponseResponse instantiates a new DescribeMergeRequestLog200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeMergeRequestLog200ResponseResponseWithDefaults

`func NewDescribeMergeRequestLog200ResponseResponseWithDefaults() *DescribeMergeRequestLog200ResponseResponse`

NewDescribeMergeRequestLog200ResponseResponseWithDefaults instantiates a new DescribeMergeRequestLog200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *DescribeMergeRequestLog200ResponseResponse) GetLogs() []MergeRequestLog`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *DescribeMergeRequestLog200ResponseResponse) GetLogsOk() (*[]MergeRequestLog, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *DescribeMergeRequestLog200ResponseResponse) SetLogs(v []MergeRequestLog)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *DescribeMergeRequestLog200ResponseResponse) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeMergeRequestLog200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeMergeRequestLog200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeMergeRequestLog200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeMergeRequestLog200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


