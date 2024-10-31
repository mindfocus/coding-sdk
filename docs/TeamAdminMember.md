# TeamAdminMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** | 头像 | [optional] [default to ""]
**GlobalKey** | Pointer to **string** | 用户 GK | [optional] [default to ""]
**Id** | Pointer to **int32** | 用户Id | [optional] 
**Name** | Pointer to **string** | 用户名 | [optional] [default to ""]
**NamePinYin** | Pointer to **string** | 用户名拼音 | [optional] [default to ""]
**TeamId** | Pointer to **int32** | 团队Id | [optional] 

## Methods

### NewTeamAdminMember

`func NewTeamAdminMember() *TeamAdminMember`

NewTeamAdminMember instantiates a new TeamAdminMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamAdminMemberWithDefaults

`func NewTeamAdminMemberWithDefaults() *TeamAdminMember`

NewTeamAdminMemberWithDefaults instantiates a new TeamAdminMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *TeamAdminMember) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *TeamAdminMember) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *TeamAdminMember) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *TeamAdminMember) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetGlobalKey

`func (o *TeamAdminMember) GetGlobalKey() string`

GetGlobalKey returns the GlobalKey field if non-nil, zero value otherwise.

### GetGlobalKeyOk

`func (o *TeamAdminMember) GetGlobalKeyOk() (*string, bool)`

GetGlobalKeyOk returns a tuple with the GlobalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalKey

`func (o *TeamAdminMember) SetGlobalKey(v string)`

SetGlobalKey sets GlobalKey field to given value.

### HasGlobalKey

`func (o *TeamAdminMember) HasGlobalKey() bool`

HasGlobalKey returns a boolean if a field has been set.

### GetId

`func (o *TeamAdminMember) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamAdminMember) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamAdminMember) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TeamAdminMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TeamAdminMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamAdminMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamAdminMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamAdminMember) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamePinYin

`func (o *TeamAdminMember) GetNamePinYin() string`

GetNamePinYin returns the NamePinYin field if non-nil, zero value otherwise.

### GetNamePinYinOk

`func (o *TeamAdminMember) GetNamePinYinOk() (*string, bool)`

GetNamePinYinOk returns a tuple with the NamePinYin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePinYin

`func (o *TeamAdminMember) SetNamePinYin(v string)`

SetNamePinYin sets NamePinYin field to given value.

### HasNamePinYin

`func (o *TeamAdminMember) HasNamePinYin() bool`

HasNamePinYin returns a boolean if a field has been set.

### GetTeamId

`func (o *TeamAdminMember) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *TeamAdminMember) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *TeamAdminMember) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *TeamAdminMember) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


