# ProtectedBranchMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalKey** | Pointer to **string** | 成员global_key | [optional] [default to ""]
**HasPushAccess** | Pointer to **bool** | 是否允许强制推送 | [optional] [default to false]
**HasUpdateAccess** | Pointer to **bool** | 是否允许合并修改分支 | [optional] [default to false]
**Name** | Pointer to **string** | 成员姓名 | [optional] [default to ""]

## Methods

### NewProtectedBranchMember

`func NewProtectedBranchMember() *ProtectedBranchMember`

NewProtectedBranchMember instantiates a new ProtectedBranchMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedBranchMemberWithDefaults

`func NewProtectedBranchMemberWithDefaults() *ProtectedBranchMember`

NewProtectedBranchMemberWithDefaults instantiates a new ProtectedBranchMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalKey

`func (o *ProtectedBranchMember) GetGlobalKey() string`

GetGlobalKey returns the GlobalKey field if non-nil, zero value otherwise.

### GetGlobalKeyOk

`func (o *ProtectedBranchMember) GetGlobalKeyOk() (*string, bool)`

GetGlobalKeyOk returns a tuple with the GlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalKey

`func (o *ProtectedBranchMember) SetGlobalKey(v string)`

SetGlobalKey sets GlobalKey field to given value.

### HasGlobalKey

`func (o *ProtectedBranchMember) HasGlobalKey() bool`

HasGlobalKey returns a boolean if a field has been set.

### GetHasPushAccess

`func (o *ProtectedBranchMember) GetHasPushAccess() bool`

GetHasPushAccess returns the HasPushAccess field if non-nil, zero value otherwise.

### GetHasPushAccessOk

`func (o *ProtectedBranchMember) GetHasPushAccessOk() (*bool, bool)`

GetHasPushAccessOk returns a tuple with the HasPushAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPushAccess

`func (o *ProtectedBranchMember) SetHasPushAccess(v bool)`

SetHasPushAccess sets HasPushAccess field to given value.

### HasHasPushAccess

`func (o *ProtectedBranchMember) HasHasPushAccess() bool`

HasHasPushAccess returns a boolean if a field has been set.

### GetHasUpdateAccess

`func (o *ProtectedBranchMember) GetHasUpdateAccess() bool`

GetHasUpdateAccess returns the HasUpdateAccess field if non-nil, zero value otherwise.

### GetHasUpdateAccessOk

`func (o *ProtectedBranchMember) GetHasUpdateAccessOk() (*bool, bool)`

GetHasUpdateAccessOk returns a tuple with the HasUpdateAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdateAccess

`func (o *ProtectedBranchMember) SetHasUpdateAccess(v bool)`

SetHasUpdateAccess sets HasUpdateAccess field to given value.

### HasHasUpdateAccess

`func (o *ProtectedBranchMember) HasHasUpdateAccess() bool`

HasHasUpdateAccess returns a boolean if a field has been set.

### GetName

`func (o *ProtectedBranchMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtectedBranchMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtectedBranchMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProtectedBranchMember) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


