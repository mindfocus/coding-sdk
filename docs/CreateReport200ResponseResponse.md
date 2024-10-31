# CreateReport200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ReportLittleData**](ReportLittleData.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewCreateReport200ResponseResponse

`func NewCreateReport200ResponseResponse() *CreateReport200ResponseResponse`

NewCreateReport200ResponseResponse instantiates a new CreateReport200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReport200ResponseResponseWithDefaults

`func NewCreateReport200ResponseResponseWithDefaults() *CreateReport200ResponseResponse`

NewCreateReport200ResponseResponseWithDefaults instantiates a new CreateReport200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateReport200ResponseResponse) GetData() ReportLittleData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateReport200ResponseResponse) GetDataOk() (*ReportLittleData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateReport200ResponseResponse) SetData(v ReportLittleData)`

SetData sets Data field to given value.

### HasData

`func (o *CreateReport200ResponseResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *CreateReport200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateReport200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateReport200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateReport200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

