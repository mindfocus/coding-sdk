# DescribeProtectedBranches200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProtectedBranches** | Pointer to [**[]ProtectedBranch**](ProtectedBranch.md) | 保护分支列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeProtectedBranches200ResponseResponse

`func NewDescribeProtectedBranches200ResponseResponse() *DescribeProtectedBranches200ResponseResponse`

NewDescribeProtectedBranches200ResponseResponse instantiates a new DescribeProtectedBranches200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProtectedBranches200ResponseResponseWithDefaults

`func NewDescribeProtectedBranches200ResponseResponseWithDefaults() *DescribeProtectedBranches200ResponseResponse`

NewDescribeProtectedBranches200ResponseResponseWithDefaults instantiates a new DescribeProtectedBranches200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtectedBranches

`func (o *DescribeProtectedBranches200ResponseResponse) GetProtectedBranches() []ProtectedBranch`

GetProtectedBranches returns the ProtectedBranches field if non-nil, zero value otherwise.

### GetProtectedBranchesOk

`func (o *DescribeProtectedBranches200ResponseResponse) GetProtectedBranchesOk() (*[]ProtectedBranch, bool)`

GetProtectedBranchesOk returns a tuple with the ProtectedBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedBranches

`func (o *DescribeProtectedBranches200ResponseResponse) SetProtectedBranches(v []ProtectedBranch)`

SetProtectedBranches sets ProtectedBranches field to given value.

### HasProtectedBranches

`func (o *DescribeProtectedBranches200ResponseResponse) HasProtectedBranches() bool`

HasProtectedBranches returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeProtectedBranches200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeProtectedBranches200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeProtectedBranches200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeProtectedBranches200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


