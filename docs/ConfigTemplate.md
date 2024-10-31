# ConfigTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableInt64** | 配置方案Code  | [optional] 
**CooperateMode** | Pointer to **NullableString** | 配置方案协作类型，包括 SCRUM 和 CLASSIC  | [optional] [default to ""]
**CreatedAt** | Pointer to **NullableInt64** | 创建时间 | [optional] 
**Description** | Pointer to **NullableString** | 配置方案描述  | [optional] [default to ""]
**Id** | Pointer to **NullableInt64** | 配置方案ID | [optional] 
**IsDraft** | Pointer to **NullableBool** | 是否是草稿配置方案  | [optional] [default to false]
**IsSystem** | Pointer to **NullableBool** | 是否是系统配置方案  | [optional] [default to false]
**Name** | Pointer to **NullableString** | 配置方案名字  | [optional] [default to ""]
**NameType** | Pointer to **NullableString** | 配置方案名字，取值和CooperateMode字段一样  | [optional] [default to ""]
**RelatedProjects** | Pointer to **NullableInt64** | 关联项目id  | [optional] 
**Scope** | Pointer to **NullableString** | 配置方案类型，和入参的TemplateType取值一样  | [optional] [default to ""]
**TeamId** | Pointer to **NullableInt64** | 团队ID | [optional] 
**UpdatedAt** | Pointer to **NullableInt64** | 更新时间 | [optional] 

## Methods

### NewConfigTemplate

`func NewConfigTemplate() *ConfigTemplate`

NewConfigTemplate instantiates a new ConfigTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigTemplateWithDefaults

`func NewConfigTemplateWithDefaults() *ConfigTemplate`

NewConfigTemplateWithDefaults instantiates a new ConfigTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ConfigTemplate) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ConfigTemplate) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ConfigTemplate) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *ConfigTemplate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ConfigTemplate) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ConfigTemplate) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCooperateMode

`func (o *ConfigTemplate) GetCooperateMode() string`

GetCooperateMode returns the CooperateMode field if non-nil, zero value otherwise.

### GetCooperateModeOk

`func (o *ConfigTemplate) GetCooperateModeOk() (*string, bool)`

GetCooperateModeOk returns a tuple with the CooperateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooperateMode

`func (o *ConfigTemplate) SetCooperateMode(v string)`

SetCooperateMode sets CooperateMode field to given value.

### HasCooperateMode

`func (o *ConfigTemplate) HasCooperateMode() bool`

HasCooperateMode returns a boolean if a field has been set.

### SetCooperateModeNil

`func (o *ConfigTemplate) SetCooperateModeNil(b bool)`

 SetCooperateModeNil sets the value for CooperateMode to be an explicit nil

### UnsetCooperateMode
`func (o *ConfigTemplate) UnsetCooperateMode()`

UnsetCooperateMode ensures that no value is present for CooperateMode, not even an explicit nil
### GetCreatedAt

`func (o *ConfigTemplate) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConfigTemplate) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConfigTemplate) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConfigTemplate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ConfigTemplate) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ConfigTemplate) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDescription

`func (o *ConfigTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConfigTemplate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConfigTemplate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *ConfigTemplate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigTemplate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigTemplate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ConfigTemplate) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ConfigTemplate) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsDraft

`func (o *ConfigTemplate) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *ConfigTemplate) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *ConfigTemplate) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.

### HasIsDraft

`func (o *ConfigTemplate) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### SetIsDraftNil

`func (o *ConfigTemplate) SetIsDraftNil(b bool)`

 SetIsDraftNil sets the value for IsDraft to be an explicit nil

### UnsetIsDraft
`func (o *ConfigTemplate) UnsetIsDraft()`

UnsetIsDraft ensures that no value is present for IsDraft, not even an explicit nil
### GetIsSystem

`func (o *ConfigTemplate) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *ConfigTemplate) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *ConfigTemplate) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *ConfigTemplate) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.

### SetIsSystemNil

`func (o *ConfigTemplate) SetIsSystemNil(b bool)`

 SetIsSystemNil sets the value for IsSystem to be an explicit nil

### UnsetIsSystem
`func (o *ConfigTemplate) UnsetIsSystem()`

UnsetIsSystem ensures that no value is present for IsSystem, not even an explicit nil
### GetName

`func (o *ConfigTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConfigTemplate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConfigTemplate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNameType

`func (o *ConfigTemplate) GetNameType() string`

GetNameType returns the NameType field if non-nil, zero value otherwise.

### GetNameTypeOk

`func (o *ConfigTemplate) GetNameTypeOk() (*string, bool)`

GetNameTypeOk returns a tuple with the NameType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameType

`func (o *ConfigTemplate) SetNameType(v string)`

SetNameType sets NameType field to given value.

### HasNameType

`func (o *ConfigTemplate) HasNameType() bool`

HasNameType returns a boolean if a field has been set.

### SetNameTypeNil

`func (o *ConfigTemplate) SetNameTypeNil(b bool)`

 SetNameTypeNil sets the value for NameType to be an explicit nil

### UnsetNameType
`func (o *ConfigTemplate) UnsetNameType()`

UnsetNameType ensures that no value is present for NameType, not even an explicit nil
### GetRelatedProjects

`func (o *ConfigTemplate) GetRelatedProjects() int64`

GetRelatedProjects returns the RelatedProjects field if non-nil, zero value otherwise.

### GetRelatedProjectsOk

`func (o *ConfigTemplate) GetRelatedProjectsOk() (*int64, bool)`

GetRelatedProjectsOk returns a tuple with the RelatedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedProjects

`func (o *ConfigTemplate) SetRelatedProjects(v int64)`

SetRelatedProjects sets RelatedProjects field to given value.

### HasRelatedProjects

`func (o *ConfigTemplate) HasRelatedProjects() bool`

HasRelatedProjects returns a boolean if a field has been set.

### SetRelatedProjectsNil

`func (o *ConfigTemplate) SetRelatedProjectsNil(b bool)`

 SetRelatedProjectsNil sets the value for RelatedProjects to be an explicit nil

### UnsetRelatedProjects
`func (o *ConfigTemplate) UnsetRelatedProjects()`

UnsetRelatedProjects ensures that no value is present for RelatedProjects, not even an explicit nil
### GetScope

`func (o *ConfigTemplate) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ConfigTemplate) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ConfigTemplate) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ConfigTemplate) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *ConfigTemplate) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *ConfigTemplate) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetTeamId

`func (o *ConfigTemplate) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *ConfigTemplate) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *ConfigTemplate) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *ConfigTemplate) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *ConfigTemplate) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *ConfigTemplate) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil
### GetUpdatedAt

`func (o *ConfigTemplate) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConfigTemplate) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConfigTemplate) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ConfigTemplate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ConfigTemplate) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ConfigTemplate) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


