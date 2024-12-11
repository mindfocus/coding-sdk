# ModifyTestRun200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**RunData**](RunData.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewModifyTestRun200ResponseResponse

`func NewModifyTestRun200ResponseResponse() *ModifyTestRun200ResponseResponse`

NewModifyTestRun200ResponseResponse instantiates a new ModifyTestRun200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyTestRun200ResponseResponseWithDefaults

`func NewModifyTestRun200ResponseResponseWithDefaults() *ModifyTestRun200ResponseResponse`

NewModifyTestRun200ResponseResponseWithDefaults instantiates a new ModifyTestRun200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ModifyTestRun200ResponseResponse) GetData() RunData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ModifyTestRun200ResponseResponse) GetDataOk() (*RunData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ModifyTestRun200ResponseResponse) SetData(v RunData)`

SetData sets Data field to given value.

### HasData

`func (o *ModifyTestRun200ResponseResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *ModifyTestRun200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ModifyTestRun200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ModifyTestRun200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ModifyTestRun200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


