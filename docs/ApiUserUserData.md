# ApiUserUserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | **string** | 头像 | [default to ""]
**DepartmentMember** | [**ApiUserUserDepartmentMember**](ApiUserUserDepartmentMember.md) |  | 
**Email** | **string** | 邮件 | [default to ""]
**EmailValidation** | **int64** | 邮件校验 | [default to 0]
**GlobalKey** | **string** | gk | [default to ""]
**Id** | **int64** | 用户ID | [default to 0]
**Name** | **string** | 用户名 | [default to ""]
**NamePinYin** | **string** | 用户拼音名 | [default to ""]
**Phone** | **string** | 手机号 | [default to ""]
**PhoneValidation** | **int64** | 手机校验 | [default to 0]
**Roles** | [**[]ApiUserRole**](ApiUserRole.md) | 角色 | 
**Status** | **int64** | 用户状态 0不活跃，1活跃，-1被锁定，-2锁定登录，-3退出团队 | [default to 0]
**TeamId** | **int64** | 团队ID | [default to 0]
**UniqueExtField** | **string** | 团队用户扩展字段、唯一。非必填 | [default to ""]

## Methods

### NewApiUserUserData

`func NewApiUserUserData(avatar string, departmentMember ApiUserUserDepartmentMember, email string, emailValidation int64, globalKey string, id int64, name string, namePinYin string, phone string, phoneValidation int64, roles []ApiUserRole, status int64, teamId int64, uniqueExtField string, ) *ApiUserUserData`

NewApiUserUserData instantiates a new ApiUserUserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserUserDataWithDefaults

`func NewApiUserUserDataWithDefaults() *ApiUserUserData`

NewApiUserUserDataWithDefaults instantiates a new ApiUserUserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *ApiUserUserData) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ApiUserUserData) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ApiUserUserData) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetDepartmentMember

`func (o *ApiUserUserData) GetDepartmentMember() ApiUserUserDepartmentMember`

GetDepartmentMember returns the DepartmentMember field if non-nil, zero value otherwise.

### GetDepartmentMemberOk

`func (o *ApiUserUserData) GetDepartmentMemberOk() (*ApiUserUserDepartmentMember, bool)`

GetDepartmentMemberOk returns a tuple with the DepartmentMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentMember

`func (o *ApiUserUserData) SetDepartmentMember(v ApiUserUserDepartmentMember)`

SetDepartmentMember sets DepartmentMember field to given value.


### GetEmail

`func (o *ApiUserUserData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApiUserUserData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApiUserUserData) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEmailValidation

`func (o *ApiUserUserData) GetEmailValidation() int64`

GetEmailValidation returns the EmailValidation field if non-nil, zero value otherwise.

### GetEmailValidationOk

`func (o *ApiUserUserData) GetEmailValidationOk() (*int64, bool)`

GetEmailValidationOk returns a tuple with the EmailValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailValidation

`func (o *ApiUserUserData) SetEmailValidation(v int64)`

SetEmailValidation sets EmailValidation field to given value.


### GetGlobalKey

`func (o *ApiUserUserData) GetGlobalKey() string`

GetGlobalKey returns the GlobalKey field if non-nil, zero value otherwise.

### GetGlobalKeyOk

`func (o *ApiUserUserData) GetGlobalKeyOk() (*string, bool)`

GetGlobalKeyOk returns a tuple with the GlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalKey

`func (o *ApiUserUserData) SetGlobalKey(v string)`

SetGlobalKey sets GlobalKey field to given value.


### GetId

`func (o *ApiUserUserData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiUserUserData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiUserUserData) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ApiUserUserData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiUserUserData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiUserUserData) SetName(v string)`

SetName sets Name field to given value.


### GetNamePinYin

`func (o *ApiUserUserData) GetNamePinYin() string`

GetNamePinYin returns the NamePinYin field if non-nil, zero value otherwise.

### GetNamePinYinOk

`func (o *ApiUserUserData) GetNamePinYinOk() (*string, bool)`

GetNamePinYinOk returns a tuple with the NamePinYin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePinYin

`func (o *ApiUserUserData) SetNamePinYin(v string)`

SetNamePinYin sets NamePinYin field to given value.


### GetPhone

`func (o *ApiUserUserData) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ApiUserUserData) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ApiUserUserData) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetPhoneValidation

`func (o *ApiUserUserData) GetPhoneValidation() int64`

GetPhoneValidation returns the PhoneValidation field if non-nil, zero value otherwise.

### GetPhoneValidationOk

`func (o *ApiUserUserData) GetPhoneValidationOk() (*int64, bool)`

GetPhoneValidationOk returns a tuple with the PhoneValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneValidation

`func (o *ApiUserUserData) SetPhoneValidation(v int64)`

SetPhoneValidation sets PhoneValidation field to given value.


### GetRoles

`func (o *ApiUserUserData) GetRoles() []ApiUserRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiUserUserData) GetRolesOk() (*[]ApiUserRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiUserUserData) SetRoles(v []ApiUserRole)`

SetRoles sets Roles field to given value.


### GetStatus

`func (o *ApiUserUserData) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiUserUserData) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiUserUserData) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetTeamId

`func (o *ApiUserUserData) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *ApiUserUserData) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *ApiUserUserData) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.


### GetUniqueExtField

`func (o *ApiUserUserData) GetUniqueExtField() string`

GetUniqueExtField returns the UniqueExtField field if non-nil, zero value otherwise.

### GetUniqueExtFieldOk

`func (o *ApiUserUserData) GetUniqueExtFieldOk() (*string, bool)`

GetUniqueExtFieldOk returns a tuple with the UniqueExtField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueExtField

`func (o *ApiUserUserData) SetUniqueExtField(v string)`

SetUniqueExtField sets UniqueExtField field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


