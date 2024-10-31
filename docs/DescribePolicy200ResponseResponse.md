# DescribePolicy200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyInfo** | Pointer to [**PolicyDetail**](PolicyDetail.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribePolicy200ResponseResponse

`func NewDescribePolicy200ResponseResponse() *DescribePolicy200ResponseResponse`

NewDescribePolicy200ResponseResponse instantiates a new DescribePolicy200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribePolicy200ResponseResponseWithDefaults

`func NewDescribePolicy200ResponseResponseWithDefaults() *DescribePolicy200ResponseResponse`

NewDescribePolicy200ResponseResponseWithDefaults instantiates a new DescribePolicy200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyInfo

`func (o *DescribePolicy200ResponseResponse) GetPolicyInfo() PolicyDetail`

GetPolicyInfo returns the PolicyInfo field if non-nil, zero value otherwise.

### GetPolicyInfoOk

`func (o *DescribePolicy200ResponseResponse) GetPolicyInfoOk() (*PolicyDetail, bool)`

GetPolicyInfoOk returns a tuple with the PolicyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyInfo

`func (o *DescribePolicy200ResponseResponse) SetPolicyInfo(v PolicyDetail)`

SetPolicyInfo sets PolicyInfo field to given value.

### HasPolicyInfo

`func (o *DescribePolicy200ResponseResponse) HasPolicyInfo() bool`

HasPolicyInfo returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribePolicy200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribePolicy200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribePolicy200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribePolicy200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


