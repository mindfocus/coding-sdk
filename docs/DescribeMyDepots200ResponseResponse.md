# DescribeMyDepots200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | Pointer to [**DepotData**](DepotData.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeMyDepots200ResponseResponse

`func NewDescribeMyDepots200ResponseResponse() *DescribeMyDepots200ResponseResponse`

NewDescribeMyDepots200ResponseResponse instantiates a new DescribeMyDepots200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeMyDepots200ResponseResponseWithDefaults

`func NewDescribeMyDepots200ResponseResponseWithDefaults() *DescribeMyDepots200ResponseResponse`

NewDescribeMyDepots200ResponseResponseWithDefaults instantiates a new DescribeMyDepots200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *DescribeMyDepots200ResponseResponse) GetPayload() DepotData`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *DescribeMyDepots200ResponseResponse) GetPayloadOk() (*DepotData, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *DescribeMyDepots200ResponseResponse) SetPayload(v DepotData)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *DescribeMyDepots200ResponseResponse) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeMyDepots200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeMyDepots200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeMyDepots200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeMyDepots200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


