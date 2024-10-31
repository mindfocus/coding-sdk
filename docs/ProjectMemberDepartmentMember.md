# ProjectMemberDepartmentMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefId** | **NullableInt64** | 部门成员refId | 
**Refs** | [**[]ProjectMemberMemberRef**](ProjectMemberMemberRef.md) | 关联信息 | 
**ThirdPartyAvatar** | **NullableString** | 三方头像 | [default to ""]
**ThirdPartyName** | **NullableString** | 三方名 | [default to ""]
**ThirdPartyId** | Pointer to **string** | 三方ID，目前仅支持ldap的用户id信息 | [optional] [default to ""]

## Methods

### NewProjectMemberDepartmentMember

`func NewProjectMemberDepartmentMember(refId NullableInt64, refs []ProjectMemberMemberRef, thirdPartyAvatar NullableString, thirdPartyName NullableString, ) *ProjectMemberDepartmentMember`

NewProjectMemberDepartmentMember instantiates a new ProjectMemberDepartmentMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectMemberDepartmentMemberWithDefaults

`func NewProjectMemberDepartmentMemberWithDefaults() *ProjectMemberDepartmentMember`

NewProjectMemberDepartmentMemberWithDefaults instantiates a new ProjectMemberDepartmentMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefId

`func (o *ProjectMemberDepartmentMember) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ProjectMemberDepartmentMember) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ProjectMemberDepartmentMember) SetRefId(v int64)`

SetRefId sets RefId field to given value.


### SetRefIdNil

`func (o *ProjectMemberDepartmentMember) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *ProjectMemberDepartmentMember) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetRefs

`func (o *ProjectMemberDepartmentMember) GetRefs() []ProjectMemberMemberRef`

GetRefs returns the Refs field if non-nil, zero value otherwise.

### GetRefsOk

`func (o *ProjectMemberDepartmentMember) GetRefsOk() (*[]ProjectMemberMemberRef, bool)`

GetRefsOk returns a tuple with the Refs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefs

`func (o *ProjectMemberDepartmentMember) SetRefs(v []ProjectMemberMemberRef)`

SetRefs sets Refs field to given value.


### SetRefsNil

`func (o *ProjectMemberDepartmentMember) SetRefsNil(b bool)`

 SetRefsNil sets the value for Refs to be an explicit nil

### UnsetRefs
`func (o *ProjectMemberDepartmentMember) UnsetRefs()`

UnsetRefs ensures that no value is present for Refs, not even an explicit nil
### GetThirdPartyAvatar

`func (o *ProjectMemberDepartmentMember) GetThirdPartyAvatar() string`

GetThirdPartyAvatar returns the ThirdPartyAvatar field if non-nil, zero value otherwise.

### GetThirdPartyAvatarOk

`func (o *ProjectMemberDepartmentMember) GetThirdPartyAvatarOk() (*string, bool)`

GetThirdPartyAvatarOk returns a tuple with the ThirdPartyAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyAvatar

`func (o *ProjectMemberDepartmentMember) SetThirdPartyAvatar(v string)`

SetThirdPartyAvatar sets ThirdPartyAvatar field to given value.


### SetThirdPartyAvatarNil

`func (o *ProjectMemberDepartmentMember) SetThirdPartyAvatarNil(b bool)`

 SetThirdPartyAvatarNil sets the value for ThirdPartyAvatar to be an explicit nil

### UnsetThirdPartyAvatar
`func (o *ProjectMemberDepartmentMember) UnsetThirdPartyAvatar()`

UnsetThirdPartyAvatar ensures that no value is present for ThirdPartyAvatar, not even an explicit nil
### GetThirdPartyName

`func (o *ProjectMemberDepartmentMember) GetThirdPartyName() string`

GetThirdPartyName returns the ThirdPartyName field if non-nil, zero value otherwise.

### GetThirdPartyNameOk

`func (o *ProjectMemberDepartmentMember) GetThirdPartyNameOk() (*string, bool)`

GetThirdPartyNameOk returns a tuple with the ThirdPartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyName

`func (o *ProjectMemberDepartmentMember) SetThirdPartyName(v string)`

SetThirdPartyName sets ThirdPartyName field to given value.


### SetThirdPartyNameNil

`func (o *ProjectMemberDepartmentMember) SetThirdPartyNameNil(b bool)`

 SetThirdPartyNameNil sets the value for ThirdPartyName to be an explicit nil

### UnsetThirdPartyName
`func (o *ProjectMemberDepartmentMember) UnsetThirdPartyName()`

UnsetThirdPartyName ensures that no value is present for ThirdPartyName, not even an explicit nil
### GetThirdPartyId

`func (o *ProjectMemberDepartmentMember) GetThirdPartyId() string`

GetThirdPartyId returns the ThirdPartyId field if non-nil, zero value otherwise.

### GetThirdPartyIdOk

`func (o *ProjectMemberDepartmentMember) GetThirdPartyIdOk() (*string, bool)`

GetThirdPartyIdOk returns a tuple with the ThirdPartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyId

`func (o *ProjectMemberDepartmentMember) SetThirdPartyId(v string)`

SetThirdPartyId sets ThirdPartyId field to given value.

### HasThirdPartyId

`func (o *ProjectMemberDepartmentMember) HasThirdPartyId() bool`

HasThirdPartyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


