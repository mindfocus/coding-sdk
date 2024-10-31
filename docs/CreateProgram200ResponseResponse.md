# CreateProgram200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ProgramProgram**](ProgramProgram.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewCreateProgram200ResponseResponse

`func NewCreateProgram200ResponseResponse() *CreateProgram200ResponseResponse`

NewCreateProgram200ResponseResponse instantiates a new CreateProgram200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProgram200ResponseResponseWithDefaults

`func NewCreateProgram200ResponseResponseWithDefaults() *CreateProgram200ResponseResponse`

NewCreateProgram200ResponseResponseWithDefaults instantiates a new CreateProgram200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateProgram200ResponseResponse) GetData() ProgramProgram`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateProgram200ResponseResponse) GetDataOk() (*ProgramProgram, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateProgram200ResponseResponse) SetData(v ProgramProgram)`

SetData sets Data field to given value.

### HasData

`func (o *CreateProgram200ResponseResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *CreateProgram200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateProgram200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateProgram200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateProgram200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


