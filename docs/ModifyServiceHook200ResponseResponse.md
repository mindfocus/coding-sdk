# ModifyServiceHook200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceHook** | Pointer to [**ServiceHook**](ServiceHook.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewModifyServiceHook200ResponseResponse

`func NewModifyServiceHook200ResponseResponse() *ModifyServiceHook200ResponseResponse`

NewModifyServiceHook200ResponseResponse instantiates a new ModifyServiceHook200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyServiceHook200ResponseResponseWithDefaults

`func NewModifyServiceHook200ResponseResponseWithDefaults() *ModifyServiceHook200ResponseResponse`

NewModifyServiceHook200ResponseResponseWithDefaults instantiates a new ModifyServiceHook200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceHook

`func (o *ModifyServiceHook200ResponseResponse) GetServiceHook() ServiceHook`

GetServiceHook returns the ServiceHook field if non-nil, zero value otherwise.

### GetServiceHookOk

`func (o *ModifyServiceHook200ResponseResponse) GetServiceHookOk() (*ServiceHook, bool)`

GetServiceHookOk returns a tuple with the ServiceHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHook

`func (o *ModifyServiceHook200ResponseResponse) SetServiceHook(v ServiceHook)`

SetServiceHook sets ServiceHook field to given value.

### HasServiceHook

`func (o *ModifyServiceHook200ResponseResponse) HasServiceHook() bool`

HasServiceHook returns a boolean if a field has been set.

### GetRequestId

`func (o *ModifyServiceHook200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ModifyServiceHook200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ModifyServiceHook200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ModifyServiceHook200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


