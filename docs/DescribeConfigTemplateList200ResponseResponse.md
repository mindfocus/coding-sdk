# DescribeConfigTemplateList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigTemplateList** | Pointer to [**[]ConfigTemplate**](ConfigTemplate.md) | 配置方案列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeConfigTemplateList200ResponseResponse

`func NewDescribeConfigTemplateList200ResponseResponse() *DescribeConfigTemplateList200ResponseResponse`

NewDescribeConfigTemplateList200ResponseResponse instantiates a new DescribeConfigTemplateList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeConfigTemplateList200ResponseResponseWithDefaults

`func NewDescribeConfigTemplateList200ResponseResponseWithDefaults() *DescribeConfigTemplateList200ResponseResponse`

NewDescribeConfigTemplateList200ResponseResponseWithDefaults instantiates a new DescribeConfigTemplateList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigTemplateList

`func (o *DescribeConfigTemplateList200ResponseResponse) GetConfigTemplateList() []ConfigTemplate`

GetConfigTemplateList returns the ConfigTemplateList field if non-nil, zero value otherwise.

### GetConfigTemplateListOk

`func (o *DescribeConfigTemplateList200ResponseResponse) GetConfigTemplateListOk() (*[]ConfigTemplate, bool)`

GetConfigTemplateListOk returns a tuple with the ConfigTemplateList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTemplateList

`func (o *DescribeConfigTemplateList200ResponseResponse) SetConfigTemplateList(v []ConfigTemplate)`

SetConfigTemplateList sets ConfigTemplateList field to given value.

### HasConfigTemplateList

`func (o *DescribeConfigTemplateList200ResponseResponse) HasConfigTemplateList() bool`

HasConfigTemplateList returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeConfigTemplateList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeConfigTemplateList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeConfigTemplateList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeConfigTemplateList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


