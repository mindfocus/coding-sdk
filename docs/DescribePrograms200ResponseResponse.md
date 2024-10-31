# DescribePrograms200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ProgramData**](ProgramData.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribePrograms200ResponseResponse

`func NewDescribePrograms200ResponseResponse() *DescribePrograms200ResponseResponse`

NewDescribePrograms200ResponseResponse instantiates a new DescribePrograms200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribePrograms200ResponseResponseWithDefaults

`func NewDescribePrograms200ResponseResponseWithDefaults() *DescribePrograms200ResponseResponse`

NewDescribePrograms200ResponseResponseWithDefaults instantiates a new DescribePrograms200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DescribePrograms200ResponseResponse) GetData() ProgramData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DescribePrograms200ResponseResponse) GetDataOk() (*ProgramData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DescribePrograms200ResponseResponse) SetData(v ProgramData)`

SetData sets Data field to given value.

### HasData

`func (o *DescribePrograms200ResponseResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribePrograms200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribePrograms200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribePrograms200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribePrograms200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


