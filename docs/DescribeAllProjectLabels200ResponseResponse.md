# DescribeAllProjectLabels200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to [**[]ProjectLabelLabels**](ProjectLabelLabels.md) | 标签列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeAllProjectLabels200ResponseResponse

`func NewDescribeAllProjectLabels200ResponseResponse() *DescribeAllProjectLabels200ResponseResponse`

NewDescribeAllProjectLabels200ResponseResponse instantiates a new DescribeAllProjectLabels200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAllProjectLabels200ResponseResponseWithDefaults

`func NewDescribeAllProjectLabels200ResponseResponseWithDefaults() *DescribeAllProjectLabels200ResponseResponse`

NewDescribeAllProjectLabels200ResponseResponseWithDefaults instantiates a new DescribeAllProjectLabels200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *DescribeAllProjectLabels200ResponseResponse) GetLabels() []ProjectLabelLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *DescribeAllProjectLabels200ResponseResponse) GetLabelsOk() (*[]ProjectLabelLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *DescribeAllProjectLabels200ResponseResponse) SetLabels(v []ProjectLabelLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *DescribeAllProjectLabels200ResponseResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeAllProjectLabels200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeAllProjectLabels200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeAllProjectLabels200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeAllProjectLabels200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


