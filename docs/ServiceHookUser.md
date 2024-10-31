# ServiceHookUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** | 用户头像 | [optional] [default to ""]
**Id** | Pointer to **int64** | 用户编号 | [optional] 
**Name** | Pointer to **string** | 用户名 | [optional] [default to ""]
**NamePinyin** | Pointer to **string** | 用户名拼音 | [optional] [default to ""]

## Methods

### NewServiceHookUser

`func NewServiceHookUser() *ServiceHookUser`

NewServiceHookUser instantiates a new ServiceHookUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceHookUserWithDefaults

`func NewServiceHookUserWithDefaults() *ServiceHookUser`

NewServiceHookUserWithDefaults instantiates a new ServiceHookUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *ServiceHookUser) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ServiceHookUser) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ServiceHookUser) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ServiceHookUser) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetId

`func (o *ServiceHookUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceHookUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceHookUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceHookUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceHookUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceHookUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceHookUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceHookUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamePinyin

`func (o *ServiceHookUser) GetNamePinyin() string`

GetNamePinyin returns the NamePinyin field if non-nil, zero value otherwise.

### GetNamePinyinOk

`func (o *ServiceHookUser) GetNamePinyinOk() (*string, bool)`

GetNamePinyinOk returns a tuple with the NamePinyin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePinyin

`func (o *ServiceHookUser) SetNamePinyin(v string)`

SetNamePinyin sets NamePinyin field to given value.

### HasNamePinyin

`func (o *ServiceHookUser) HasNamePinyin() bool`

HasNamePinyin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


