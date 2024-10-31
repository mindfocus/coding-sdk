# CreateTestCase200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CaseDataSpecial**](CaseDataSpecial.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewCreateTestCase200ResponseResponse

`func NewCreateTestCase200ResponseResponse() *CreateTestCase200ResponseResponse`

NewCreateTestCase200ResponseResponse instantiates a new CreateTestCase200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestCase200ResponseResponseWithDefaults

`func NewCreateTestCase200ResponseResponseWithDefaults() *CreateTestCase200ResponseResponse`

NewCreateTestCase200ResponseResponseWithDefaults instantiates a new CreateTestCase200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateTestCase200ResponseResponse) GetData() CaseDataSpecial`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateTestCase200ResponseResponse) GetDataOk() (*CaseDataSpecial, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateTestCase200ResponseResponse) SetData(v CaseDataSpecial)`

SetData sets Data field to given value.

### HasData

`func (o *CreateTestCase200ResponseResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *CreateTestCase200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateTestCase200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateTestCase200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateTestCase200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


