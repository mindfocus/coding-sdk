# Program

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | 结束时间 | [optional] 
**Description** | Pointer to **string** | 描述信息 | [optional] 
**DisplayName** | Pointer to **string** | 展示名称 | [optional] 
**EndDate** | Pointer to **int64** | 结束时间 | [optional] 
**Icon** | Pointer to **string** | 图标 | [optional] 
**Id** | Pointer to **int32** | ID | [optional] 
**Name** | Pointer to **string** | 标识 | [optional] 
**NamePinyin** | Pointer to **string** | 名称拼音 | [optional] 
**StartDate** | Pointer to **int64** | 开始时间 | [optional] 
**UpdatedAt** | Pointer to **int64** | 修改时间 | [optional] 

## Methods

### NewProgram

`func NewProgram() *Program`

NewProgram instantiates a new Program object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgramWithDefaults

`func NewProgramWithDefaults() *Program`

NewProgramWithDefaults instantiates a new Program object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Program) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Program) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Program) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Program) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Program) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Program) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Program) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Program) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *Program) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Program) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Program) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Program) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEndDate

`func (o *Program) GetEndDate() int64`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Program) GetEndDateOk() (*int64, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Program) SetEndDate(v int64)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Program) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetIcon

`func (o *Program) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Program) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Program) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Program) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *Program) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Program) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Program) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Program) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Program) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Program) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Program) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Program) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamePinyin

`func (o *Program) GetNamePinyin() string`

GetNamePinyin returns the NamePinyin field if non-nil, zero value otherwise.

### GetNamePinyinOk

`func (o *Program) GetNamePinyinOk() (*string, bool)`

GetNamePinyinOk returns a tuple with the NamePinyin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePinyin

`func (o *Program) SetNamePinyin(v string)`

SetNamePinyin sets NamePinyin field to given value.

### HasNamePinyin

`func (o *Program) HasNamePinyin() bool`

HasNamePinyin returns a boolean if a field has been set.

### GetStartDate

`func (o *Program) GetStartDate() int64`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Program) GetStartDateOk() (*int64, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Program) SetStartDate(v int64)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Program) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Program) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Program) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Program) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Program) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


