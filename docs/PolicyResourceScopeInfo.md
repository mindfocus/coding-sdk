# PolicyResourceScopeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | 创建时间 | [optional] 
**PolicyId** | Pointer to **int64** | 权限组 ID | [optional] 
**Resource** | Pointer to [**ResourceInfoOfPolicyScope**](ResourceInfoOfPolicyScope.md) |  | [optional] 

## Methods

### NewPolicyResourceScopeInfo

`func NewPolicyResourceScopeInfo() *PolicyResourceScopeInfo`

NewPolicyResourceScopeInfo instantiates a new PolicyResourceScopeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyResourceScopeInfoWithDefaults

`func NewPolicyResourceScopeInfoWithDefaults() *PolicyResourceScopeInfo`

NewPolicyResourceScopeInfoWithDefaults instantiates a new PolicyResourceScopeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PolicyResourceScopeInfo) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PolicyResourceScopeInfo) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PolicyResourceScopeInfo) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PolicyResourceScopeInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPolicyId

`func (o *PolicyResourceScopeInfo) GetPolicyId() int64`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PolicyResourceScopeInfo) GetPolicyIdOk() (*int64, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PolicyResourceScopeInfo) SetPolicyId(v int64)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *PolicyResourceScopeInfo) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetResource

`func (o *PolicyResourceScopeInfo) GetResource() ResourceInfoOfPolicyScope`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PolicyResourceScopeInfo) GetResourceOk() (*ResourceInfoOfPolicyScope, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PolicyResourceScopeInfo) SetResource(v ResourceInfoOfPolicyScope)`

SetResource sets Resource field to given value.

### HasResource

`func (o *PolicyResourceScopeInfo) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


