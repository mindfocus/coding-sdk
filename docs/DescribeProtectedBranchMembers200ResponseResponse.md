# DescribeProtectedBranchMembers200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]ProtectedBranchMember**](ProtectedBranchMember.md) | 保护分支成员列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeProtectedBranchMembers200ResponseResponse

`func NewDescribeProtectedBranchMembers200ResponseResponse() *DescribeProtectedBranchMembers200ResponseResponse`

NewDescribeProtectedBranchMembers200ResponseResponse instantiates a new DescribeProtectedBranchMembers200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProtectedBranchMembers200ResponseResponseWithDefaults

`func NewDescribeProtectedBranchMembers200ResponseResponseWithDefaults() *DescribeProtectedBranchMembers200ResponseResponse`

NewDescribeProtectedBranchMembers200ResponseResponseWithDefaults instantiates a new DescribeProtectedBranchMembers200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *DescribeProtectedBranchMembers200ResponseResponse) GetMembers() []ProtectedBranchMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *DescribeProtectedBranchMembers200ResponseResponse) GetMembersOk() (*[]ProtectedBranchMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *DescribeProtectedBranchMembers200ResponseResponse) SetMembers(v []ProtectedBranchMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *DescribeProtectedBranchMembers200ResponseResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeProtectedBranchMembers200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeProtectedBranchMembers200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeProtectedBranchMembers200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeProtectedBranchMembers200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


