# UserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** | 头像 | [optional] [default to ""]
**Email** | Pointer to **string** | 邮箱 | [optional] [default to ""]
**EmailValidation** | Pointer to **int32** | 邮箱是否验证 0 否 /1 是 | [optional] 
**GlobalKey** | Pointer to **string** | 用户 GK | [optional] [default to ""]
**Id** | Pointer to **int32** | 用户Id | [optional] 
**Name** | Pointer to **string** | 用户名 | [optional] [default to ""]
**NamePinYin** | Pointer to **string** | 用户名拼音 | [optional] [default to ""]
**Phone** | Pointer to **string** | 手机号 | [optional] [default to ""]
**PhoneValidation** | Pointer to **int32** | 手机是否验证 0 否 /1 是 | [optional] 
**Roles** | Pointer to [**[]Role**](Role.md) | 用户组 | [optional] 
**Status** | Pointer to **int64** | 用户状态 0不活跃，1活跃，-1被锁定，-2锁定登录，-3退出团队 | [optional] 
**TeamId** | Pointer to **int32** | 团队Id | [optional] [default to 0]
**DepartmentMember** | Pointer to [**ProjectMemberDepartmentMember**](ProjectMemberDepartmentMember.md) |  | [optional] 
**UniqueExtField** | Pointer to **string** | 团队用户扩展字段、唯一。非必填 | [optional] [default to ""]

## Methods

### NewUserData

`func NewUserData() *UserData`

NewUserData instantiates a new UserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataWithDefaults

`func NewUserDataWithDefaults() *UserData`

NewUserDataWithDefaults instantiates a new UserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *UserData) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserData) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserData) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserData) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetEmail

`func (o *UserData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailValidation

`func (o *UserData) GetEmailValidation() int32`

GetEmailValidation returns the EmailValidation field if non-nil, zero value otherwise.

### GetEmailValidationOk

`func (o *UserData) GetEmailValidationOk() (*int32, bool)`

GetEmailValidationOk returns a tuple with the EmailValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailValidation

`func (o *UserData) SetEmailValidation(v int32)`

SetEmailValidation sets EmailValidation field to given value.

### HasEmailValidation

`func (o *UserData) HasEmailValidation() bool`

HasEmailValidation returns a boolean if a field has been set.

### GetGlobalKey

`func (o *UserData) GetGlobalKey() string`

GetGlobalKey returns the GlobalKey field if non-nil, zero value otherwise.

### GetGlobalKeyOk

`func (o *UserData) GetGlobalKeyOk() (*string, bool)`

GetGlobalKeyOk returns a tuple with the GlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalKey

`func (o *UserData) SetGlobalKey(v string)`

SetGlobalKey sets GlobalKey field to given value.

### HasGlobalKey

`func (o *UserData) HasGlobalKey() bool`

HasGlobalKey returns a boolean if a field has been set.

### GetId

`func (o *UserData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserData) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UserData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamePinYin

`func (o *UserData) GetNamePinYin() string`

GetNamePinYin returns the NamePinYin field if non-nil, zero value otherwise.

### GetNamePinYinOk

`func (o *UserData) GetNamePinYinOk() (*string, bool)`

GetNamePinYinOk returns a tuple with the NamePinYin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePinYin

`func (o *UserData) SetNamePinYin(v string)`

SetNamePinYin sets NamePinYin field to given value.

### HasNamePinYin

`func (o *UserData) HasNamePinYin() bool`

HasNamePinYin returns a boolean if a field has been set.

### GetPhone

`func (o *UserData) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserData) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserData) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UserData) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPhoneValidation

`func (o *UserData) GetPhoneValidation() int32`

GetPhoneValidation returns the PhoneValidation field if non-nil, zero value otherwise.

### GetPhoneValidationOk

`func (o *UserData) GetPhoneValidationOk() (*int32, bool)`

GetPhoneValidationOk returns a tuple with the PhoneValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneValidation

`func (o *UserData) SetPhoneValidation(v int32)`

SetPhoneValidation sets PhoneValidation field to given value.

### HasPhoneValidation

`func (o *UserData) HasPhoneValidation() bool`

HasPhoneValidation returns a boolean if a field has been set.

### GetRoles

`func (o *UserData) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserData) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserData) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserData) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetStatus

`func (o *UserData) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserData) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserData) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTeamId

`func (o *UserData) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *UserData) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *UserData) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *UserData) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetDepartmentMember

`func (o *UserData) GetDepartmentMember() ProjectMemberDepartmentMember`

GetDepartmentMember returns the DepartmentMember field if non-nil, zero value otherwise.

### GetDepartmentMemberOk

`func (o *UserData) GetDepartmentMemberOk() (*ProjectMemberDepartmentMember, bool)`

GetDepartmentMemberOk returns a tuple with the DepartmentMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentMember

`func (o *UserData) SetDepartmentMember(v ProjectMemberDepartmentMember)`

SetDepartmentMember sets DepartmentMember field to given value.

### HasDepartmentMember

`func (o *UserData) HasDepartmentMember() bool`

HasDepartmentMember returns a boolean if a field has been set.

### GetUniqueExtField

`func (o *UserData) GetUniqueExtField() string`

GetUniqueExtField returns the UniqueExtField field if non-nil, zero value otherwise.

### GetUniqueExtFieldOk

`func (o *UserData) GetUniqueExtFieldOk() (*string, bool)`

GetUniqueExtFieldOk returns a tuple with the UniqueExtField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueExtField

`func (o *UserData) SetUniqueExtField(v string)`

SetUniqueExtField sets UniqueExtField field to given value.

### HasUniqueExtField

`func (o *UserData) HasUniqueExtField() bool`

HasUniqueExtField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


