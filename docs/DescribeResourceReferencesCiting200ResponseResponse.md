# DescribeResourceReferencesCiting200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to [**[]ResourceReferenceResourceInfo**](ResourceReferenceResourceInfo.md) | 资源引用列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeResourceReferencesCiting200ResponseResponse

`func NewDescribeResourceReferencesCiting200ResponseResponse() *DescribeResourceReferencesCiting200ResponseResponse`

NewDescribeResourceReferencesCiting200ResponseResponse instantiates a new DescribeResourceReferencesCiting200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeResourceReferencesCiting200ResponseResponseWithDefaults

`func NewDescribeResourceReferencesCiting200ResponseResponseWithDefaults() *DescribeResourceReferencesCiting200ResponseResponse`

NewDescribeResourceReferencesCiting200ResponseResponseWithDefaults instantiates a new DescribeResourceReferencesCiting200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *DescribeResourceReferencesCiting200ResponseResponse) GetResource() []ResourceReferenceResourceInfo`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DescribeResourceReferencesCiting200ResponseResponse) GetResourceOk() (*[]ResourceReferenceResourceInfo, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DescribeResourceReferencesCiting200ResponseResponse) SetResource(v []ResourceReferenceResourceInfo)`

SetResource sets Resource field to given value.

### HasResource

`func (o *DescribeResourceReferencesCiting200ResponseResponse) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeResourceReferencesCiting200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeResourceReferencesCiting200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeResourceReferencesCiting200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeResourceReferencesCiting200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


