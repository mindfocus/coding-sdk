# DescribeResourceReferencesCited200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | Pointer to [**[]ResourceReferenceResourceInfo**](ResourceReferenceResourceInfo.md) | 资源列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeResourceReferencesCited200ResponseResponse

`func NewDescribeResourceReferencesCited200ResponseResponse() *DescribeResourceReferencesCited200ResponseResponse`

NewDescribeResourceReferencesCited200ResponseResponse instantiates a new DescribeResourceReferencesCited200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeResourceReferencesCited200ResponseResponseWithDefaults

`func NewDescribeResourceReferencesCited200ResponseResponseWithDefaults() *DescribeResourceReferencesCited200ResponseResponse`

NewDescribeResourceReferencesCited200ResponseResponseWithDefaults instantiates a new DescribeResourceReferencesCited200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *DescribeResourceReferencesCited200ResponseResponse) GetResource() []ResourceReferenceResourceInfo`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DescribeResourceReferencesCited200ResponseResponse) GetResourceOk() (*[]ResourceReferenceResourceInfo, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DescribeResourceReferencesCited200ResponseResponse) SetResource(v []ResourceReferenceResourceInfo)`

SetResource sets Resource field to given value.

### HasResource

`func (o *DescribeResourceReferencesCited200ResponseResponse) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeResourceReferencesCited200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeResourceReferencesCited200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeResourceReferencesCited200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeResourceReferencesCited200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


