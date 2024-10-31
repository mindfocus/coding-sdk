# ModifyMergeMR200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MergeRequestInfo** | Pointer to [**MergeRequestInfo**](MergeRequestInfo.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewModifyMergeMR200ResponseResponse

`func NewModifyMergeMR200ResponseResponse() *ModifyMergeMR200ResponseResponse`

NewModifyMergeMR200ResponseResponse instantiates a new ModifyMergeMR200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyMergeMR200ResponseResponseWithDefaults

`func NewModifyMergeMR200ResponseResponseWithDefaults() *ModifyMergeMR200ResponseResponse`

NewModifyMergeMR200ResponseResponseWithDefaults instantiates a new ModifyMergeMR200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMergeRequestInfo

`func (o *ModifyMergeMR200ResponseResponse) GetMergeRequestInfo() MergeRequestInfo`

GetMergeRequestInfo returns the MergeRequestInfo field if non-nil, zero value otherwise.

### GetMergeRequestInfoOk

`func (o *ModifyMergeMR200ResponseResponse) GetMergeRequestInfoOk() (*MergeRequestInfo, bool)`

GetMergeRequestInfoOk returns a tuple with the MergeRequestInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeRequestInfo

`func (o *ModifyMergeMR200ResponseResponse) SetMergeRequestInfo(v MergeRequestInfo)`

SetMergeRequestInfo sets MergeRequestInfo field to given value.

### HasMergeRequestInfo

`func (o *ModifyMergeMR200ResponseResponse) HasMergeRequestInfo() bool`

HasMergeRequestInfo returns a boolean if a field has been set.

### GetRequestId

`func (o *ModifyMergeMR200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ModifyMergeMR200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ModifyMergeMR200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ModifyMergeMR200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


