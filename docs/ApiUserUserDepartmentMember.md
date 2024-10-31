# ApiUserUserDepartmentMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefId** | **int64** | 部门成员refId | [default to 0]
**Refs** | [**[]ApiUserMemberRef**](ApiUserMemberRef.md) | 关联信息 | 
**ThirdPartyAvatar** | **string** | 三方头像 | [default to ""]
**ThirdPartyName** | **string** | 三方名 | [default to ""]
**ThirdPartyId** | Pointer to **string** | 三方ID，目前仅支持ldap的用户id信息 | [optional] [default to ""]

## Methods

### NewApiUserUserDepartmentMember

`func NewApiUserUserDepartmentMember(refId int64, refs []ApiUserMemberRef, thirdPartyAvatar string, thirdPartyName string, ) *ApiUserUserDepartmentMember`

NewApiUserUserDepartmentMember instantiates a new ApiUserUserDepartmentMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserUserDepartmentMemberWithDefaults

`func NewApiUserUserDepartmentMemberWithDefaults() *ApiUserUserDepartmentMember`

NewApiUserUserDepartmentMemberWithDefaults instantiates a new ApiUserUserDepartmentMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefId

`func (o *ApiUserUserDepartmentMember) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ApiUserUserDepartmentMember) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ApiUserUserDepartmentMember) SetRefId(v int64)`

SetRefId sets RefId field to given value.


### GetRefs

`func (o *ApiUserUserDepartmentMember) GetRefs() []ApiUserMemberRef`

GetRefs returns the Refs field if non-nil, zero value otherwise.

### GetRefsOk

`func (o *ApiUserUserDepartmentMember) GetRefsOk() (*[]ApiUserMemberRef, bool)`

GetRefsOk returns a tuple with the Refs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefs

`func (o *ApiUserUserDepartmentMember) SetRefs(v []ApiUserMemberRef)`

SetRefs sets Refs field to given value.


### GetThirdPartyAvatar

`func (o *ApiUserUserDepartmentMember) GetThirdPartyAvatar() string`

GetThirdPartyAvatar returns the ThirdPartyAvatar field if non-nil, zero value otherwise.

### GetThirdPartyAvatarOk

`func (o *ApiUserUserDepartmentMember) GetThirdPartyAvatarOk() (*string, bool)`

GetThirdPartyAvatarOk returns a tuple with the ThirdPartyAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyAvatar

`func (o *ApiUserUserDepartmentMember) SetThirdPartyAvatar(v string)`

SetThirdPartyAvatar sets ThirdPartyAvatar field to given value.


### GetThirdPartyName

`func (o *ApiUserUserDepartmentMember) GetThirdPartyName() string`

GetThirdPartyName returns the ThirdPartyName field if non-nil, zero value otherwise.

### GetThirdPartyNameOk

`func (o *ApiUserUserDepartmentMember) GetThirdPartyNameOk() (*string, bool)`

GetThirdPartyNameOk returns a tuple with the ThirdPartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyName

`func (o *ApiUserUserDepartmentMember) SetThirdPartyName(v string)`

SetThirdPartyName sets ThirdPartyName field to given value.


### GetThirdPartyId

`func (o *ApiUserUserDepartmentMember) GetThirdPartyId() string`

GetThirdPartyId returns the ThirdPartyId field if non-nil, zero value otherwise.

### GetThirdPartyIdOk

`func (o *ApiUserUserDepartmentMember) GetThirdPartyIdOk() (*string, bool)`

GetThirdPartyIdOk returns a tuple with the ThirdPartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartyId

`func (o *ApiUserUserDepartmentMember) SetThirdPartyId(v string)`

SetThirdPartyId sets ThirdPartyId field to given value.

### HasThirdPartyId

`func (o *ApiUserUserDepartmentMember) HasThirdPartyId() bool`

HasThirdPartyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


