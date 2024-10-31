# CreateCodingProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | 项目描述 | [optional] 
**DisplayName** | **string** | 项目名称 | 
**Icon** | Pointer to **string** | 项目图标 | [optional] 
**Name** | **string** | 项目标识 | 
**ProjectTemplate** | **string** | 项目模版 CODE_HOST 代码托管项目， PROJECT_MANAGE 项目管理项目， DEV_OPS DevOps项目， DEMO_BEGIN 范例项目 | 
**Shared** | **int32** | 0： 不公开 1：公开源代码 | 

## Methods

### NewCreateCodingProjectRequest

`func NewCreateCodingProjectRequest(displayName string, name string, projectTemplate string, shared int32, ) *CreateCodingProjectRequest`

NewCreateCodingProjectRequest instantiates a new CreateCodingProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCodingProjectRequestWithDefaults

`func NewCreateCodingProjectRequestWithDefaults() *CreateCodingProjectRequest`

NewCreateCodingProjectRequestWithDefaults instantiates a new CreateCodingProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateCodingProjectRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCodingProjectRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCodingProjectRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCodingProjectRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateCodingProjectRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateCodingProjectRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateCodingProjectRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIcon

`func (o *CreateCodingProjectRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CreateCodingProjectRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CreateCodingProjectRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CreateCodingProjectRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetName

`func (o *CreateCodingProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCodingProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCodingProjectRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProjectTemplate

`func (o *CreateCodingProjectRequest) GetProjectTemplate() string`

GetProjectTemplate returns the ProjectTemplate field if non-nil, zero value otherwise.

### GetProjectTemplateOk

`func (o *CreateCodingProjectRequest) GetProjectTemplateOk() (*string, bool)`

GetProjectTemplateOk returns a tuple with the ProjectTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTemplate

`func (o *CreateCodingProjectRequest) SetProjectTemplate(v string)`

SetProjectTemplate sets ProjectTemplate field to given value.


### GetShared

`func (o *CreateCodingProjectRequest) GetShared() int32`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *CreateCodingProjectRequest) GetSharedOk() (*int32, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *CreateCodingProjectRequest) SetShared(v int32)`

SetShared sets Shared field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


