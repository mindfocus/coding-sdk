# DepotUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **NullableString** | 用户头像 | [optional] [default to ""]
**Email** | Pointer to **string** | 邮箱 | [optional] [default to ""]
**GlobalKey** | Pointer to **NullableString** | 用户GlobalKey | [optional] [default to ""]
**Id** | Pointer to **NullableInt64** | 用户id | [optional] 
**Name** | Pointer to **string** | 姓名 | [optional] [default to ""]
**Status** | Pointer to **NullableString** | 用户状态 | [optional] [default to ""]
**TeamId** | Pointer to **NullableInt64** | 团队id | [optional] 

## Methods

### NewDepotUser

`func NewDepotUser() *DepotUser`

NewDepotUser instantiates a new DepotUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepotUserWithDefaults

`func NewDepotUserWithDefaults() *DepotUser`

NewDepotUserWithDefaults instantiates a new DepotUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *DepotUser) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *DepotUser) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *DepotUser) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *DepotUser) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *DepotUser) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *DepotUser) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetEmail

`func (o *DepotUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DepotUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DepotUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DepotUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGlobalKey

`func (o *DepotUser) GetGlobalKey() string`

GetGlobalKey returns the GlobalKey field if non-nil, zero value otherwise.

### GetGlobalKeyOk

`func (o *DepotUser) GetGlobalKeyOk() (*string, bool)`

GetGlobalKeyOk returns a tuple with the GlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalKey

`func (o *DepotUser) SetGlobalKey(v string)`

SetGlobalKey sets GlobalKey field to given value.

### HasGlobalKey

`func (o *DepotUser) HasGlobalKey() bool`

HasGlobalKey returns a boolean if a field has been set.

### SetGlobalKeyNil

`func (o *DepotUser) SetGlobalKeyNil(b bool)`

 SetGlobalKeyNil sets the value for GlobalKey to be an explicit nil

### UnsetGlobalKey
`func (o *DepotUser) UnsetGlobalKey()`

UnsetGlobalKey ensures that no value is present for GlobalKey, not even an explicit nil
### GetId

`func (o *DepotUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DepotUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DepotUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DepotUser) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DepotUser) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DepotUser) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *DepotUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DepotUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DepotUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DepotUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DepotUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DepotUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DepotUser) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DepotUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DepotUser) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DepotUser) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTeamId

`func (o *DepotUser) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *DepotUser) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *DepotUser) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *DepotUser) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *DepotUser) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *DepotUser) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


