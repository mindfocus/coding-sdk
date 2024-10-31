# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** | 头像地址 | [optional] [default to ""]
**Email** | Pointer to **string** | 邮箱 | [optional] [default to ""]
**GlobalKey** | Pointer to **string** | gk | [optional] [default to ""]
**Id** | Pointer to **int32** | userId | [optional] 
**Name** | Pointer to **string** | 名称 | [optional] [default to ""]
**Phone** | Pointer to **string** | 电话 | [optional] [default to ""]
**Status** | Pointer to **NullableInt32** | 状态 | [optional] 
**TeamGlobalKey** | Pointer to **string** | 团队gk | [optional] [default to ""]
**TeamId** | Pointer to **int32** | 团队Id | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *User) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *User) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *User) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *User) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGlobalKey

`func (o *User) GetGlobalKey() string`

GetGlobalKey returns the GlobalKey field if non-nil, zero value otherwise.

### GetGlobalKeyOk

`func (o *User) GetGlobalKeyOk() (*string, bool)`

GetGlobalKeyOk returns a tuple with the GlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalKey

`func (o *User) SetGlobalKey(v string)`

SetGlobalKey sets GlobalKey field to given value.

### HasGlobalKey

`func (o *User) HasGlobalKey() bool`

HasGlobalKey returns a boolean if a field has been set.

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *User) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *User) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *User) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *User) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStatus

`func (o *User) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *User) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *User) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *User) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *User) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *User) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTeamGlobalKey

`func (o *User) GetTeamGlobalKey() string`

GetTeamGlobalKey returns the TeamGlobalKey field if non-nil, zero value otherwise.

### GetTeamGlobalKeyOk

`func (o *User) GetTeamGlobalKeyOk() (*string, bool)`

GetTeamGlobalKeyOk returns a tuple with the TeamGlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamGlobalKey

`func (o *User) SetTeamGlobalKey(v string)`

SetTeamGlobalKey sets TeamGlobalKey field to given value.

### HasTeamGlobalKey

`func (o *User) HasTeamGlobalKey() bool`

HasTeamGlobalKey returns a boolean if a field has been set.

### GetTeamId

`func (o *User) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *User) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *User) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *User) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


