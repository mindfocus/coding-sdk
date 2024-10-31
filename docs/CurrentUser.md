# CurrentUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **NullableString** | 头像地址 | [optional] [default to ""]
**Email** | Pointer to **NullableString** | 邮箱 | [optional] [default to ""]
**EmailValidation** | Pointer to **NullableInt64** | 邮箱是否已验证 | [optional] 
**GlobalKey** | Pointer to **NullableString** | 用户唯一标志 | [optional] [default to ""]
**Gravatar** | Pointer to **NullableString** | 无用头像 | [optional] [default to ""]
**Id** | Pointer to **int64** | 用户 ID | [optional] 
**Name** | Pointer to **NullableString** | 姓名 | [optional] [default to ""]
**NamePinYin** | Pointer to **NullableString** | 姓名拼音 | [optional] [default to ""]
**Phone** | Pointer to **NullableString** | 联系电话 | [optional] [default to ""]
**PhoneRegionCode** | Pointer to **NullableString** | 手机号地区 | [optional] [default to ""]
**PhoneValidation** | Pointer to **NullableInt64** | 手机是否已验证 | [optional] 
**Region** | Pointer to **NullableString** | 区号 | [optional] [default to ""]
**Status** | Pointer to **NullableInt64** | 状态(新用户/已激活) | [optional] 
**TeamId** | Pointer to **NullableInt64** | 团队 ID | [optional] 
**WeComId** | Pointer to **NullableString** | 企微id | [optional] [default to ""]

## Methods

### NewCurrentUser

`func NewCurrentUser() *CurrentUser`

NewCurrentUser instantiates a new CurrentUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentUserWithDefaults

`func NewCurrentUserWithDefaults() *CurrentUser`

NewCurrentUserWithDefaults instantiates a new CurrentUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *CurrentUser) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *CurrentUser) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *CurrentUser) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *CurrentUser) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *CurrentUser) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *CurrentUser) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetEmail

`func (o *CurrentUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CurrentUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CurrentUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CurrentUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CurrentUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CurrentUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailValidation

`func (o *CurrentUser) GetEmailValidation() int64`

GetEmailValidation returns the EmailValidation field if non-nil, zero value otherwise.

### GetEmailValidationOk

`func (o *CurrentUser) GetEmailValidationOk() (*int64, bool)`

GetEmailValidationOk returns a tuple with the EmailValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailValidation

`func (o *CurrentUser) SetEmailValidation(v int64)`

SetEmailValidation sets EmailValidation field to given value.

### HasEmailValidation

`func (o *CurrentUser) HasEmailValidation() bool`

HasEmailValidation returns a boolean if a field has been set.

### SetEmailValidationNil

`func (o *CurrentUser) SetEmailValidationNil(b bool)`

 SetEmailValidationNil sets the value for EmailValidation to be an explicit nil

### UnsetEmailValidation
`func (o *CurrentUser) UnsetEmailValidation()`

UnsetEmailValidation ensures that no value is present for EmailValidation, not even an explicit nil
### GetGlobalKey

`func (o *CurrentUser) GetGlobalKey() string`

GetGlobalKey returns the GlobalKey field if non-nil, zero value otherwise.

### GetGlobalKeyOk

`func (o *CurrentUser) GetGlobalKeyOk() (*string, bool)`

GetGlobalKeyOk returns a tuple with the GlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalKey

`func (o *CurrentUser) SetGlobalKey(v string)`

SetGlobalKey sets GlobalKey field to given value.

### HasGlobalKey

`func (o *CurrentUser) HasGlobalKey() bool`

HasGlobalKey returns a boolean if a field has been set.

### SetGlobalKeyNil

`func (o *CurrentUser) SetGlobalKeyNil(b bool)`

 SetGlobalKeyNil sets the value for GlobalKey to be an explicit nil

### UnsetGlobalKey
`func (o *CurrentUser) UnsetGlobalKey()`

UnsetGlobalKey ensures that no value is present for GlobalKey, not even an explicit nil
### GetGravatar

`func (o *CurrentUser) GetGravatar() string`

GetGravatar returns the Gravatar field if non-nil, zero value otherwise.

### GetGravatarOk

`func (o *CurrentUser) GetGravatarOk() (*string, bool)`

GetGravatarOk returns a tuple with the Gravatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatar

`func (o *CurrentUser) SetGravatar(v string)`

SetGravatar sets Gravatar field to given value.

### HasGravatar

`func (o *CurrentUser) HasGravatar() bool`

HasGravatar returns a boolean if a field has been set.

### SetGravatarNil

`func (o *CurrentUser) SetGravatarNil(b bool)`

 SetGravatarNil sets the value for Gravatar to be an explicit nil

### UnsetGravatar
`func (o *CurrentUser) UnsetGravatar()`

UnsetGravatar ensures that no value is present for Gravatar, not even an explicit nil
### GetId

`func (o *CurrentUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrentUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrentUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CurrentUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CurrentUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CurrentUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CurrentUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CurrentUser) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CurrentUser) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CurrentUser) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNamePinYin

`func (o *CurrentUser) GetNamePinYin() string`

GetNamePinYin returns the NamePinYin field if non-nil, zero value otherwise.

### GetNamePinYinOk

`func (o *CurrentUser) GetNamePinYinOk() (*string, bool)`

GetNamePinYinOk returns a tuple with the NamePinYin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePinYin

`func (o *CurrentUser) SetNamePinYin(v string)`

SetNamePinYin sets NamePinYin field to given value.

### HasNamePinYin

`func (o *CurrentUser) HasNamePinYin() bool`

HasNamePinYin returns a boolean if a field has been set.

### SetNamePinYinNil

`func (o *CurrentUser) SetNamePinYinNil(b bool)`

 SetNamePinYinNil sets the value for NamePinYin to be an explicit nil

### UnsetNamePinYin
`func (o *CurrentUser) UnsetNamePinYin()`

UnsetNamePinYin ensures that no value is present for NamePinYin, not even an explicit nil
### GetPhone

`func (o *CurrentUser) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CurrentUser) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CurrentUser) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CurrentUser) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *CurrentUser) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *CurrentUser) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetPhoneRegionCode

`func (o *CurrentUser) GetPhoneRegionCode() string`

GetPhoneRegionCode returns the PhoneRegionCode field if non-nil, zero value otherwise.

### GetPhoneRegionCodeOk

`func (o *CurrentUser) GetPhoneRegionCodeOk() (*string, bool)`

GetPhoneRegionCodeOk returns a tuple with the PhoneRegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneRegionCode

`func (o *CurrentUser) SetPhoneRegionCode(v string)`

SetPhoneRegionCode sets PhoneRegionCode field to given value.

### HasPhoneRegionCode

`func (o *CurrentUser) HasPhoneRegionCode() bool`

HasPhoneRegionCode returns a boolean if a field has been set.

### SetPhoneRegionCodeNil

`func (o *CurrentUser) SetPhoneRegionCodeNil(b bool)`

 SetPhoneRegionCodeNil sets the value for PhoneRegionCode to be an explicit nil

### UnsetPhoneRegionCode
`func (o *CurrentUser) UnsetPhoneRegionCode()`

UnsetPhoneRegionCode ensures that no value is present for PhoneRegionCode, not even an explicit nil
### GetPhoneValidation

`func (o *CurrentUser) GetPhoneValidation() int64`

GetPhoneValidation returns the PhoneValidation field if non-nil, zero value otherwise.

### GetPhoneValidationOk

`func (o *CurrentUser) GetPhoneValidationOk() (*int64, bool)`

GetPhoneValidationOk returns a tuple with the PhoneValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneValidation

`func (o *CurrentUser) SetPhoneValidation(v int64)`

SetPhoneValidation sets PhoneValidation field to given value.

### HasPhoneValidation

`func (o *CurrentUser) HasPhoneValidation() bool`

HasPhoneValidation returns a boolean if a field has been set.

### SetPhoneValidationNil

`func (o *CurrentUser) SetPhoneValidationNil(b bool)`

 SetPhoneValidationNil sets the value for PhoneValidation to be an explicit nil

### UnsetPhoneValidation
`func (o *CurrentUser) UnsetPhoneValidation()`

UnsetPhoneValidation ensures that no value is present for PhoneValidation, not even an explicit nil
### GetRegion

`func (o *CurrentUser) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CurrentUser) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CurrentUser) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CurrentUser) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *CurrentUser) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *CurrentUser) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetStatus

`func (o *CurrentUser) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CurrentUser) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CurrentUser) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CurrentUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CurrentUser) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CurrentUser) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTeamId

`func (o *CurrentUser) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *CurrentUser) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *CurrentUser) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *CurrentUser) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *CurrentUser) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *CurrentUser) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil
### GetWeComId

`func (o *CurrentUser) GetWeComId() string`

GetWeComId returns the WeComId field if non-nil, zero value otherwise.

### GetWeComIdOk

`func (o *CurrentUser) GetWeComIdOk() (*string, bool)`

GetWeComIdOk returns a tuple with the WeComId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeComId

`func (o *CurrentUser) SetWeComId(v string)`

SetWeComId sets WeComId field to given value.

### HasWeComId

`func (o *CurrentUser) HasWeComId() bool`

HasWeComId returns a boolean if a field has been set.

### SetWeComIdNil

`func (o *CurrentUser) SetWeComIdNil(b bool)`

 SetWeComIdNil sets the value for WeComId to be an explicit nil

### UnsetWeComId
`func (o *CurrentUser) UnsetWeComId()`

UnsetWeComId ensures that no value is present for WeComId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


