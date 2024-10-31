# BranchProtectionMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPush** | Pointer to **bool** | 是否允许直接推送 | [optional] [default to false]
**GlobalKey** | Pointer to **string** | 用户globalkey | [optional] [default to ""]
**Name** | Pointer to **string** | 名称 | [optional] [default to ""]

## Methods

### NewBranchProtectionMember

`func NewBranchProtectionMember() *BranchProtectionMember`

NewBranchProtectionMember instantiates a new BranchProtectionMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchProtectionMemberWithDefaults

`func NewBranchProtectionMemberWithDefaults() *BranchProtectionMember`

NewBranchProtectionMemberWithDefaults instantiates a new BranchProtectionMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPush

`func (o *BranchProtectionMember) GetAllowPush() bool`

GetAllowPush returns the AllowPush field if non-nil, zero value otherwise.

### GetAllowPushOk

`func (o *BranchProtectionMember) GetAllowPushOk() (*bool, bool)`

GetAllowPushOk returns a tuple with the AllowPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPush

`func (o *BranchProtectionMember) SetAllowPush(v bool)`

SetAllowPush sets AllowPush field to given value.

### HasAllowPush

`func (o *BranchProtectionMember) HasAllowPush() bool`

HasAllowPush returns a boolean if a field has been set.

### GetGlobalKey

`func (o *BranchProtectionMember) GetGlobalKey() string`

GetGlobalKey returns the GlobalKey field if non-nil, zero value otherwise.

### GetGlobalKeyOk

`func (o *BranchProtectionMember) GetGlobalKeyOk() (*string, bool)`

GetGlobalKeyOk returns a tuple with the GlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalKey

`func (o *BranchProtectionMember) SetGlobalKey(v string)`

SetGlobalKey sets GlobalKey field to given value.

### HasGlobalKey

`func (o *BranchProtectionMember) HasGlobalKey() bool`

HasGlobalKey returns a boolean if a field has been set.

### GetName

`func (o *BranchProtectionMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BranchProtectionMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BranchProtectionMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BranchProtectionMember) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


