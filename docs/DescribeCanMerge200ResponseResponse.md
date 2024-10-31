# DescribeCanMerge200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MergeStatus** | Pointer to **string** | NOT_MERGEABLE(不允许合并),ALREADY_MERGED(完全一样的两个分支),MERGEABLE(可合并),FAILED(比较失败) | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeCanMerge200ResponseResponse

`func NewDescribeCanMerge200ResponseResponse() *DescribeCanMerge200ResponseResponse`

NewDescribeCanMerge200ResponseResponse instantiates a new DescribeCanMerge200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCanMerge200ResponseResponseWithDefaults

`func NewDescribeCanMerge200ResponseResponseWithDefaults() *DescribeCanMerge200ResponseResponse`

NewDescribeCanMerge200ResponseResponseWithDefaults instantiates a new DescribeCanMerge200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMergeStatus

`func (o *DescribeCanMerge200ResponseResponse) GetMergeStatus() string`

GetMergeStatus returns the MergeStatus field if non-nil, zero value otherwise.

### GetMergeStatusOk

`func (o *DescribeCanMerge200ResponseResponse) GetMergeStatusOk() (*string, bool)`

GetMergeStatusOk returns a tuple with the MergeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeStatus

`func (o *DescribeCanMerge200ResponseResponse) SetMergeStatus(v string)`

SetMergeStatus sets MergeStatus field to given value.

### HasMergeStatus

`func (o *DescribeCanMerge200ResponseResponse) HasMergeStatus() bool`

HasMergeStatus returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeCanMerge200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeCanMerge200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeCanMerge200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeCanMerge200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


