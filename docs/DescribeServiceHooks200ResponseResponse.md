# DescribeServiceHooks200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ServiceHookPage**](ServiceHookPage.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeServiceHooks200ResponseResponse

`func NewDescribeServiceHooks200ResponseResponse() *DescribeServiceHooks200ResponseResponse`

NewDescribeServiceHooks200ResponseResponse instantiates a new DescribeServiceHooks200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeServiceHooks200ResponseResponseWithDefaults

`func NewDescribeServiceHooks200ResponseResponseWithDefaults() *DescribeServiceHooks200ResponseResponse`

NewDescribeServiceHooks200ResponseResponseWithDefaults instantiates a new DescribeServiceHooks200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DescribeServiceHooks200ResponseResponse) GetData() ServiceHookPage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DescribeServiceHooks200ResponseResponse) GetDataOk() (*ServiceHookPage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DescribeServiceHooks200ResponseResponse) SetData(v ServiceHookPage)`

SetData sets Data field to given value.

### HasData

`func (o *DescribeServiceHooks200ResponseResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeServiceHooks200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeServiceHooks200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeServiceHooks200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeServiceHooks200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


