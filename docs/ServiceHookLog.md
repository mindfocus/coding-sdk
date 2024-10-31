# ServiceHookLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableInt64** | 创建时间 | [optional] 
**Event** | Pointer to **NullableString** | 事件名 | [optional] [default to ""]
**Id** | Pointer to **string** | 日志记录编号 | [optional] [default to ""]
**RequestContent** | Pointer to **NullableString** | 发送请求内容 | [optional] [default to ""]
**RequestHeaders** | Pointer to **NullableString** | 发送请求头 | [optional] [default to ""]
**RequestId** | Pointer to **NullableString** | 发送请求编号 | [optional] [default to ""]
**ResponseAt** | Pointer to **NullableInt64** | 响应事件 | [optional] 
**ResponseBody** | Pointer to **NullableString** | 响应内容 | [optional] [default to ""]
**ResponseHeaders** | Pointer to **NullableString** | 响应头 | [optional] [default to ""]
**ResponseStatus** | Pointer to **NullableInt64** | 响应状态码 | [optional] 
**SendAt** | Pointer to **NullableInt64** | 发送事件 | [optional] 
**ServiceHookId** | Pointer to **NullableString** | Service Hook 编号 | [optional] [default to ""]
**Status** | Pointer to **NullableInt64** | 发送状态 | [optional] 

## Methods

### NewServiceHookLog

`func NewServiceHookLog() *ServiceHookLog`

NewServiceHookLog instantiates a new ServiceHookLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceHookLogWithDefaults

`func NewServiceHookLogWithDefaults() *ServiceHookLog`

NewServiceHookLogWithDefaults instantiates a new ServiceHookLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServiceHookLog) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServiceHookLog) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServiceHookLog) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServiceHookLog) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ServiceHookLog) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ServiceHookLog) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetEvent

`func (o *ServiceHookLog) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ServiceHookLog) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ServiceHookLog) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *ServiceHookLog) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *ServiceHookLog) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *ServiceHookLog) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetId

`func (o *ServiceHookLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceHookLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceHookLog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceHookLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequestContent

`func (o *ServiceHookLog) GetRequestContent() string`

GetRequestContent returns the RequestContent field if non-nil, zero value otherwise.

### GetRequestContentOk

`func (o *ServiceHookLog) GetRequestContentOk() (*string, bool)`

GetRequestContentOk returns a tuple with the RequestContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestContent

`func (o *ServiceHookLog) SetRequestContent(v string)`

SetRequestContent sets RequestContent field to given value.

### HasRequestContent

`func (o *ServiceHookLog) HasRequestContent() bool`

HasRequestContent returns a boolean if a field has been set.

### SetRequestContentNil

`func (o *ServiceHookLog) SetRequestContentNil(b bool)`

 SetRequestContentNil sets the value for RequestContent to be an explicit nil

### UnsetRequestContent
`func (o *ServiceHookLog) UnsetRequestContent()`

UnsetRequestContent ensures that no value is present for RequestContent, not even an explicit nil
### GetRequestHeaders

`func (o *ServiceHookLog) GetRequestHeaders() string`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *ServiceHookLog) GetRequestHeadersOk() (*string, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *ServiceHookLog) SetRequestHeaders(v string)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *ServiceHookLog) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### SetRequestHeadersNil

`func (o *ServiceHookLog) SetRequestHeadersNil(b bool)`

 SetRequestHeadersNil sets the value for RequestHeaders to be an explicit nil

### UnsetRequestHeaders
`func (o *ServiceHookLog) UnsetRequestHeaders()`

UnsetRequestHeaders ensures that no value is present for RequestHeaders, not even an explicit nil
### GetRequestId

`func (o *ServiceHookLog) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ServiceHookLog) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ServiceHookLog) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ServiceHookLog) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### SetRequestIdNil

`func (o *ServiceHookLog) SetRequestIdNil(b bool)`

 SetRequestIdNil sets the value for RequestId to be an explicit nil

### UnsetRequestId
`func (o *ServiceHookLog) UnsetRequestId()`

UnsetRequestId ensures that no value is present for RequestId, not even an explicit nil
### GetResponseAt

`func (o *ServiceHookLog) GetResponseAt() int64`

GetResponseAt returns the ResponseAt field if non-nil, zero value otherwise.

### GetResponseAtOk

`func (o *ServiceHookLog) GetResponseAtOk() (*int64, bool)`

GetResponseAtOk returns a tuple with the ResponseAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseAt

`func (o *ServiceHookLog) SetResponseAt(v int64)`

SetResponseAt sets ResponseAt field to given value.

### HasResponseAt

`func (o *ServiceHookLog) HasResponseAt() bool`

HasResponseAt returns a boolean if a field has been set.

### SetResponseAtNil

`func (o *ServiceHookLog) SetResponseAtNil(b bool)`

 SetResponseAtNil sets the value for ResponseAt to be an explicit nil

### UnsetResponseAt
`func (o *ServiceHookLog) UnsetResponseAt()`

UnsetResponseAt ensures that no value is present for ResponseAt, not even an explicit nil
### GetResponseBody

`func (o *ServiceHookLog) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *ServiceHookLog) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *ServiceHookLog) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *ServiceHookLog) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### SetResponseBodyNil

`func (o *ServiceHookLog) SetResponseBodyNil(b bool)`

 SetResponseBodyNil sets the value for ResponseBody to be an explicit nil

### UnsetResponseBody
`func (o *ServiceHookLog) UnsetResponseBody()`

UnsetResponseBody ensures that no value is present for ResponseBody, not even an explicit nil
### GetResponseHeaders

`func (o *ServiceHookLog) GetResponseHeaders() string`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *ServiceHookLog) GetResponseHeadersOk() (*string, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *ServiceHookLog) SetResponseHeaders(v string)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *ServiceHookLog) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### SetResponseHeadersNil

`func (o *ServiceHookLog) SetResponseHeadersNil(b bool)`

 SetResponseHeadersNil sets the value for ResponseHeaders to be an explicit nil

### UnsetResponseHeaders
`func (o *ServiceHookLog) UnsetResponseHeaders()`

UnsetResponseHeaders ensures that no value is present for ResponseHeaders, not even an explicit nil
### GetResponseStatus

`func (o *ServiceHookLog) GetResponseStatus() int64`

GetResponseStatus returns the ResponseStatus field if non-nil, zero value otherwise.

### GetResponseStatusOk

`func (o *ServiceHookLog) GetResponseStatusOk() (*int64, bool)`

GetResponseStatusOk returns a tuple with the ResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatus

`func (o *ServiceHookLog) SetResponseStatus(v int64)`

SetResponseStatus sets ResponseStatus field to given value.

### HasResponseStatus

`func (o *ServiceHookLog) HasResponseStatus() bool`

HasResponseStatus returns a boolean if a field has been set.

### SetResponseStatusNil

`func (o *ServiceHookLog) SetResponseStatusNil(b bool)`

 SetResponseStatusNil sets the value for ResponseStatus to be an explicit nil

### UnsetResponseStatus
`func (o *ServiceHookLog) UnsetResponseStatus()`

UnsetResponseStatus ensures that no value is present for ResponseStatus, not even an explicit nil
### GetSendAt

`func (o *ServiceHookLog) GetSendAt() int64`

GetSendAt returns the SendAt field if non-nil, zero value otherwise.

### GetSendAtOk

`func (o *ServiceHookLog) GetSendAtOk() (*int64, bool)`

GetSendAtOk returns a tuple with the SendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendAt

`func (o *ServiceHookLog) SetSendAt(v int64)`

SetSendAt sets SendAt field to given value.

### HasSendAt

`func (o *ServiceHookLog) HasSendAt() bool`

HasSendAt returns a boolean if a field has been set.

### SetSendAtNil

`func (o *ServiceHookLog) SetSendAtNil(b bool)`

 SetSendAtNil sets the value for SendAt to be an explicit nil

### UnsetSendAt
`func (o *ServiceHookLog) UnsetSendAt()`

UnsetSendAt ensures that no value is present for SendAt, not even an explicit nil
### GetServiceHookId

`func (o *ServiceHookLog) GetServiceHookId() string`

GetServiceHookId returns the ServiceHookId field if non-nil, zero value otherwise.

### GetServiceHookIdOk

`func (o *ServiceHookLog) GetServiceHookIdOk() (*string, bool)`

GetServiceHookIdOk returns a tuple with the ServiceHookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHookId

`func (o *ServiceHookLog) SetServiceHookId(v string)`

SetServiceHookId sets ServiceHookId field to given value.

### HasServiceHookId

`func (o *ServiceHookLog) HasServiceHookId() bool`

HasServiceHookId returns a boolean if a field has been set.

### SetServiceHookIdNil

`func (o *ServiceHookLog) SetServiceHookIdNil(b bool)`

 SetServiceHookIdNil sets the value for ServiceHookId to be an explicit nil

### UnsetServiceHookId
`func (o *ServiceHookLog) UnsetServiceHookId()`

UnsetServiceHookId ensures that no value is present for ServiceHookId, not even an explicit nil
### GetStatus

`func (o *ServiceHookLog) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceHookLog) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceHookLog) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceHookLog) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ServiceHookLog) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ServiceHookLog) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


