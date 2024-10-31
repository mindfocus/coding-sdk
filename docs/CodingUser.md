# CodingUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **NullableString** | 用户头像地址 | [optional] [default to ""]
**Email** | Pointer to **string** | 用户电子邮箱地址 | [optional] [default to ""]
**GlobalKey** | Pointer to **NullableString** | 用户全局Key | [optional] [default to ""]
**Id** | Pointer to **int64** | 用户 ID | [optional] 
**Name** | Pointer to **NullableString** | 用户名字 | [optional] [default to ""]
**Status** | Pointer to **NullableString** | 用户状态 | [optional] [default to ""]
**TeamId** | Pointer to **NullableInt64** | 团队 ID | [optional] 

## Methods

### NewCodingUser

`func NewCodingUser() *CodingUser`

NewCodingUser instantiates a new CodingUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodingUserWithDefaults

`func NewCodingUserWithDefaults() *CodingUser`

NewCodingUserWithDefaults instantiates a new CodingUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *CodingUser) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *CodingUser) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *CodingUser) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *CodingUser) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *CodingUser) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *CodingUser) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetEmail

`func (o *CodingUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CodingUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CodingUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CodingUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGlobalKey

`func (o *CodingUser) GetGlobalKey() string`

GetGlobalKey returns the GlobalKey field if non-nil, zero value otherwise.

### GetGlobalKeyOk

`func (o *CodingUser) GetGlobalKeyOk() (*string, bool)`

GetGlobalKeyOk returns a tuple with the GlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalKey

`func (o *CodingUser) SetGlobalKey(v string)`

SetGlobalKey sets GlobalKey field to given value.

### HasGlobalKey

`func (o *CodingUser) HasGlobalKey() bool`

HasGlobalKey returns a boolean if a field has been set.

### SetGlobalKeyNil

`func (o *CodingUser) SetGlobalKeyNil(b bool)`

 SetGlobalKeyNil sets the value for GlobalKey to be an explicit nil

### UnsetGlobalKey
`func (o *CodingUser) UnsetGlobalKey()`

UnsetGlobalKey ensures that no value is present for GlobalKey, not even an explicit nil
### GetId

`func (o *CodingUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CodingUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CodingUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CodingUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CodingUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodingUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodingUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CodingUser) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CodingUser) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CodingUser) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *CodingUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CodingUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CodingUser) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CodingUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CodingUser) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CodingUser) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTeamId

`func (o *CodingUser) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *CodingUser) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *CodingUser) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *CodingUser) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *CodingUser) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *CodingUser) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


