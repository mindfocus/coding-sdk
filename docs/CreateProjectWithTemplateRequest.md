# CreateProjectWithTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | 项目描述 | [optional] 
**DisplayName** | **string** | 项目名称 | 
**Icon** | Pointer to **string** | 项目图标 | [optional] 
**Invisible** | **bool** | 隐藏项目在 CODING 入口不可见  true 不可见|false 可见 | 
**Label** | **string** | 标签(TKE、TCB、SLS、QUICKSTART、APIGW) | 
**Name** | **string** | 项目标识 | 
**ProjectTemplate** | **string** | 项目模版 CODE_HOST 代码托管项目， PROJECT_MANAGE 项目管理项目， DEV_OPS DevOps项目， DEMO_BEGIN 范例项目 | 
**Shared** | **int32** | 0： 不公开 1：公开源代码 | 

## Methods

### NewCreateProjectWithTemplateRequest

`func NewCreateProjectWithTemplateRequest(displayName string, invisible bool, label string, name string, projectTemplate string, shared int32, ) *CreateProjectWithTemplateRequest`

NewCreateProjectWithTemplateRequest instantiates a new CreateProjectWithTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectWithTemplateRequestWithDefaults

`func NewCreateProjectWithTemplateRequestWithDefaults() *CreateProjectWithTemplateRequest`

NewCreateProjectWithTemplateRequestWithDefaults instantiates a new CreateProjectWithTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateProjectWithTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProjectWithTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProjectWithTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateProjectWithTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *CreateProjectWithTemplateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateProjectWithTemplateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateProjectWithTemplateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIcon

`func (o *CreateProjectWithTemplateRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *CreateProjectWithTemplateRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *CreateProjectWithTemplateRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *CreateProjectWithTemplateRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetInvisible

`func (o *CreateProjectWithTemplateRequest) GetInvisible() bool`

GetInvisible returns the Invisible field if non-nil, zero value otherwise.

### GetInvisibleOk

`func (o *CreateProjectWithTemplateRequest) GetInvisibleOk() (*bool, bool)`

GetInvisibleOk returns a tuple with the Invisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvisible

`func (o *CreateProjectWithTemplateRequest) SetInvisible(v bool)`

SetInvisible sets Invisible field to given value.


### GetLabel

`func (o *CreateProjectWithTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateProjectWithTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateProjectWithTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetName

`func (o *CreateProjectWithTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProjectWithTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProjectWithTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProjectTemplate

`func (o *CreateProjectWithTemplateRequest) GetProjectTemplate() string`

GetProjectTemplate returns the ProjectTemplate field if non-nil, zero value otherwise.

### GetProjectTemplateOk

`func (o *CreateProjectWithTemplateRequest) GetProjectTemplateOk() (*string, bool)`

GetProjectTemplateOk returns a tuple with the ProjectTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTemplate

`func (o *CreateProjectWithTemplateRequest) SetProjectTemplate(v string)`

SetProjectTemplate sets ProjectTemplate field to given value.


### GetShared

`func (o *CreateProjectWithTemplateRequest) GetShared() int32`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *CreateProjectWithTemplateRequest) GetSharedOk() (*int32, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *CreateProjectWithTemplateRequest) SetShared(v int32)`

SetShared sets Shared field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


