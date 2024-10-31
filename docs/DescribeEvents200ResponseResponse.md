# DescribeEvents200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to [**[]ServiceHookEvent**](ServiceHookEvent.md) | 事件列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeEvents200ResponseResponse

`func NewDescribeEvents200ResponseResponse() *DescribeEvents200ResponseResponse`

NewDescribeEvents200ResponseResponse instantiates a new DescribeEvents200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeEvents200ResponseResponseWithDefaults

`func NewDescribeEvents200ResponseResponseWithDefaults() *DescribeEvents200ResponseResponse`

NewDescribeEvents200ResponseResponseWithDefaults instantiates a new DescribeEvents200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *DescribeEvents200ResponseResponse) GetEvent() []ServiceHookEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *DescribeEvents200ResponseResponse) GetEventOk() (*[]ServiceHookEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *DescribeEvents200ResponseResponse) SetEvent(v []ServiceHookEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *DescribeEvents200ResponseResponse) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeEvents200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeEvents200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeEvents200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeEvents200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


