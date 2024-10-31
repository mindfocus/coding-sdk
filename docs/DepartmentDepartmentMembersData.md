# DepartmentDepartmentMembersData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **bool** | 是否授权 实际表示已加入或已被授权，如果要表示 已授权未加入(allowed &amp;&amp; userId&gt;0) | [default to false]
**Avatar** | **string** | 头像 | [default to ""]
**Locked** | **bool** | 是否被锁定 | [default to false]
**Name** | **string** | 成员名 | [default to ""]
**Owner** | **bool** | 是否部门所有者 | [default to false]
**RefId** | **int64** | 部门用户的refId | [default to 0]
**Refs** | [**[]DepartmentDepartmentMemberRef**](DepartmentDepartmentMemberRef.md) | 关联的 部门等信息 | 
**ThirdPartyAvatar** | **string** | 三方头像 | [default to ""]
**ThirdPartyName** | **string** | 三方名 | [default to ""]
**UserGlobalKey** | **string** | 用户GK | [default to ""]
**UserId** | **int64** | 用户ID | [default to 0]
**UserStatus** | **int64** | 用户状态 0不活跃，1活跃，-1被锁定，-2锁定登录，-3退出团队 | [default to 0]

## Methods

### NewDepartmentDepartmentMembersData

`func NewDepartmentDepartmentMembersData(allowed bool, avatar string, locked bool, name string, owner bool, refId int64, refs []DepartmentDepartmentMemberRef, thirdPartyAvatar string, thirdPartyName string, userGlobalKey string, userId int64, userStatus int64, ) *DepartmentDepartmentMembersData`

NewDepartmentDepartmentMembersData instantiates a new DepartmentDepartmentMembersData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepartmentDepartmentMembersDataWithDefaults

`func NewDepartmentDepartmentMembersDataWithDefaults() *DepartmentDepartmentMembersData`

NewDepartmentDepartmentMembersDataWithDefaults instantiates a new DepartmentDepartmentMembersData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *DepartmentDepartmentMembersData) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *DepartmentDepartmentMembersData) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *DepartmentDepartmentMembersData) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.


### GetAvatar

`func (o *DepartmentDepartmentMembersData) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *DepartmentDepartmentMembersData) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *DepartmentDepartmentMembersData) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.


### GetLocked

`func (o *DepartmentDepartmentMembersData) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *DepartmentDepartmentMembersData) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *DepartmentDepartmentMembersData) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetName

`func (o *DepartmentDepartmentMembersData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DepartmentDepartmentMembersData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DepartmentDepartmentMembersData) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *DepartmentDepartmentMembersData) GetOwner() bool`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DepartmentDepartmentMembersData) GetOwnerOk() (*bool, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DepartmentDepartmentMembersData) SetOwner(v bool)`

SetOwner sets Owner field to given value.


### GetRefId

`func (o *DepartmentDepartmentMembersData) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *DepartmentDepartmentMembersData) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *DepartmentDepartmentMembersData) SetRefId(v int64)`

SetRefId sets RefId field to given value.


### GetRefs

`func (o *DepartmentDepartmentMembersData) GetRefs() []DepartmentDepartmentMemberRef`

GetRefs returns the Refs field if non-nil, zero value otherwise.

### GetRefsOk

`func (o *DepartmentDepartmentMembersData) GetRefsOk() (*[]DepartmentDepartmentMemberRef, bool)`

GetRefsOk returns a tuple with the Refs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefs

`func (o *DepartmentDepartmentMembersData) SetRefs(v []DepartmentDepartmentMemberRef)`

SetRefs sets Refs field to given value.


### GetThirdPartyAvatar

`func (o *DepartmentDepartmentMembersData) GetThirdPartyAvatar() string`

GetThirdPartyAvatar returns the ThirdPartyAvatar field if non-nil, zero value otherwise.

### GetThirdPartyAvatarOk

`func (o *DepartmentDepartmentMembersData) GetThirdPartyAvatarOk() (*string, bool)`

GetThirdPartyAvatarOk returns a tuple with the ThirdPartyAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyAvatar

`func (o *DepartmentDepartmentMembersData) SetThirdPartyAvatar(v string)`

SetThirdPartyAvatar sets ThirdPartyAvatar field to given value.


### GetThirdPartyName

`func (o *DepartmentDepartmentMembersData) GetThirdPartyName() string`

GetThirdPartyName returns the ThirdPartyName field if non-nil, zero value otherwise.

### GetThirdPartyNameOk

`func (o *DepartmentDepartmentMembersData) GetThirdPartyNameOk() (*string, bool)`

GetThirdPartyNameOk returns a tuple with the ThirdPartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyName

`func (o *DepartmentDepartmentMembersData) SetThirdPartyName(v string)`

SetThirdPartyName sets ThirdPartyName field to given value.


### GetUserGlobalKey

`func (o *DepartmentDepartmentMembersData) GetUserGlobalKey() string`

GetUserGlobalKey returns the UserGlobalKey field if non-nil, zero value otherwise.

### GetUserGlobalKeyOk

`func (o *DepartmentDepartmentMembersData) GetUserGlobalKeyOk() (*string, bool)`

GetUserGlobalKeyOk returns a tuple with the UserGlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGlobalKey

`func (o *DepartmentDepartmentMembersData) SetUserGlobalKey(v string)`

SetUserGlobalKey sets UserGlobalKey field to given value.


### GetUserId

`func (o *DepartmentDepartmentMembersData) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DepartmentDepartmentMembersData) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DepartmentDepartmentMembersData) SetUserId(v int64)`

SetUserId sets UserId field to given value.


### GetUserStatus

`func (o *DepartmentDepartmentMembersData) GetUserStatus() int64`

GetUserStatus returns the UserStatus field if non-nil, zero value otherwise.

### GetUserStatusOk

`func (o *DepartmentDepartmentMembersData) GetUserStatusOk() (*int64, bool)`

GetUserStatusOk returns a tuple with the UserStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatus

`func (o *DepartmentDepartmentMembersData) SetUserStatus(v int64)`

SetUserStatus sets UserStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


