# ResendServiceHookLog200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ServiceHookResendServiceHookResult**](ServiceHookResendServiceHookResult.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewResendServiceHookLog200ResponseResponse

`func NewResendServiceHookLog200ResponseResponse() *ResendServiceHookLog200ResponseResponse`

NewResendServiceHookLog200ResponseResponse instantiates a new ResendServiceHookLog200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResendServiceHookLog200ResponseResponseWithDefaults

`func NewResendServiceHookLog200ResponseResponseWithDefaults() *ResendServiceHookLog200ResponseResponse`

NewResendServiceHookLog200ResponseResponseWithDefaults instantiates a new ResendServiceHookLog200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ResendServiceHookLog200ResponseResponse) GetData() ServiceHookResendServiceHookResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResendServiceHookLog200ResponseResponse) GetDataOk() (*ServiceHookResendServiceHookResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResendServiceHookLog200ResponseResponse) SetData(v ServiceHookResendServiceHookResult)`

SetData sets Data field to given value.

### HasData

`func (o *ResendServiceHookLog200ResponseResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *ResendServiceHookLog200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ResendServiceHookLog200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ResendServiceHookLog200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ResendServiceHookLog200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


