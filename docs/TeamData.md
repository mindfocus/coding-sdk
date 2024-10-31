# TeamData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamHost** | Pointer to **string** | 团队域名 | [optional] [default to ""]
**NamePinYin** | Pointer to **string** | 团队名-拼音 | [optional] [default to ""]
**Id** | Pointer to **int64** | 团队ID | [optional] [default to 0]
**TeamOwner** | Pointer to [**UserData**](UserData.md) |  | [optional] 
**Avatar** | Pointer to **string** | 团队图标 | [optional] [default to ""]
**Name** | Pointer to **string** | 团队名 | [optional] [default to ""]

## Methods

### NewTeamData

`func NewTeamData() *TeamData`

NewTeamData instantiates a new TeamData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamDataWithDefaults

`func NewTeamDataWithDefaults() *TeamData`

NewTeamDataWithDefaults instantiates a new TeamData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamHost

`func (o *TeamData) GetTeamHost() string`

GetTeamHost returns the TeamHost field if non-nil, zero value otherwise.

### GetTeamHostOk

`func (o *TeamData) GetTeamHostOk() (*string, bool)`

GetTeamHostOk returns a tuple with the TeamHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamHost

`func (o *TeamData) SetTeamHost(v string)`

SetTeamHost sets TeamHost field to given value.

### HasTeamHost

`func (o *TeamData) HasTeamHost() bool`

HasTeamHost returns a boolean if a field has been set.

### GetNamePinYin

`func (o *TeamData) GetNamePinYin() string`

GetNamePinYin returns the NamePinYin field if non-nil, zero value otherwise.

### GetNamePinYinOk

`func (o *TeamData) GetNamePinYinOk() (*string, bool)`

GetNamePinYinOk returns a tuple with the NamePinYin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePinYin

`func (o *TeamData) SetNamePinYin(v string)`

SetNamePinYin sets NamePinYin field to given value.

### HasNamePinYin

`func (o *TeamData) HasNamePinYin() bool`

HasNamePinYin returns a boolean if a field has been set.

### GetId

`func (o *TeamData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamData) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TeamData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTeamOwner

`func (o *TeamData) GetTeamOwner() UserData`

GetTeamOwner returns the TeamOwner field if non-nil, zero value otherwise.

### GetTeamOwnerOk

`func (o *TeamData) GetTeamOwnerOk() (*UserData, bool)`

GetTeamOwnerOk returns a tuple with the TeamOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamOwner

`func (o *TeamData) SetTeamOwner(v UserData)`

SetTeamOwner sets TeamOwner field to given value.

### HasTeamOwner

`func (o *TeamData) HasTeamOwner() bool`

HasTeamOwner returns a boolean if a field has been set.

### GetAvatar

`func (o *TeamData) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *TeamData) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *TeamData) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *TeamData) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetName

`func (o *TeamData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamData) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


