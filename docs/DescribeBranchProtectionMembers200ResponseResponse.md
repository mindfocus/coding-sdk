# DescribeBranchProtectionMembers200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]BranchProtectionMember**](BranchProtectionMember.md) | 保护分支规则管理员列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeBranchProtectionMembers200ResponseResponse

`func NewDescribeBranchProtectionMembers200ResponseResponse() *DescribeBranchProtectionMembers200ResponseResponse`

NewDescribeBranchProtectionMembers200ResponseResponse instantiates a new DescribeBranchProtectionMembers200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeBranchProtectionMembers200ResponseResponseWithDefaults

`func NewDescribeBranchProtectionMembers200ResponseResponseWithDefaults() *DescribeBranchProtectionMembers200ResponseResponse`

NewDescribeBranchProtectionMembers200ResponseResponseWithDefaults instantiates a new DescribeBranchProtectionMembers200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *DescribeBranchProtectionMembers200ResponseResponse) GetMembers() []BranchProtectionMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *DescribeBranchProtectionMembers200ResponseResponse) GetMembersOk() (*[]BranchProtectionMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *DescribeBranchProtectionMembers200ResponseResponse) SetMembers(v []BranchProtectionMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *DescribeBranchProtectionMembers200ResponseResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeBranchProtectionMembers200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeBranchProtectionMembers200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeBranchProtectionMembers200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeBranchProtectionMembers200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


