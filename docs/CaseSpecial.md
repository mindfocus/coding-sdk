# CaseSpecial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableString** | 创建时间 | [optional] [default to ""]
**CreatedBy** | Pointer to **NullableString** | 创建人ID | [optional] [default to ""]
**CustomSteps** | Pointer to [**[]CustomStep**](CustomStep.md) | 自定义步骤 | [optional] 
**Expected** | Pointer to **NullableString** | 预期结果 （适用于文本用例） | [optional] [default to ""]
**Id** | Pointer to **NullableInt32** | ID 主键 | [optional] 
**Preconds** | Pointer to **NullableString** | 前置步骤 | [optional] [default to ""]
**Priority** | Pointer to **NullableInt32** | 优先级 | [optional] 
**SectionId** | Pointer to **NullableInt32** | 分组 ID | [optional] 
**Sort** | Pointer to **NullableInt32** | 排序值 | [optional] 
**Steps** | Pointer to **NullableString** | 文本描述（适用于文本用例） | [optional] [default to ""]
**TemplateType** | Pointer to **NullableString** | 用例类型，可选值：STEPS，TEXT | [optional] [default to ""]
**Title** | Pointer to **NullableString** | 标题 | [optional] [default to ""]
**UpdatedAt** | Pointer to **NullableString** | 更新时间 | [optional] [default to ""]

## Methods

### NewCaseSpecial

`func NewCaseSpecial() *CaseSpecial`

NewCaseSpecial instantiates a new CaseSpecial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaseSpecialWithDefaults

`func NewCaseSpecialWithDefaults() *CaseSpecial`

NewCaseSpecialWithDefaults instantiates a new CaseSpecial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *CaseSpecial) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CaseSpecial) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CaseSpecial) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CaseSpecial) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *CaseSpecial) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CaseSpecial) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *CaseSpecial) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CaseSpecial) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CaseSpecial) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CaseSpecial) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *CaseSpecial) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *CaseSpecial) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCustomSteps

`func (o *CaseSpecial) GetCustomSteps() []CustomStep`

GetCustomSteps returns the CustomSteps field if non-nil, zero value otherwise.

### GetCustomStepsOk

`func (o *CaseSpecial) GetCustomStepsOk() (*[]CustomStep, bool)`

GetCustomStepsOk returns a tuple with the CustomSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSteps

`func (o *CaseSpecial) SetCustomSteps(v []CustomStep)`

SetCustomSteps sets CustomSteps field to given value.

### HasCustomSteps

`func (o *CaseSpecial) HasCustomSteps() bool`

HasCustomSteps returns a boolean if a field has been set.

### SetCustomStepsNil

`func (o *CaseSpecial) SetCustomStepsNil(b bool)`

 SetCustomStepsNil sets the value for CustomSteps to be an explicit nil

### UnsetCustomSteps
`func (o *CaseSpecial) UnsetCustomSteps()`

UnsetCustomSteps ensures that no value is present for CustomSteps, not even an explicit nil
### GetExpected

`func (o *CaseSpecial) GetExpected() string`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *CaseSpecial) GetExpectedOk() (*string, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *CaseSpecial) SetExpected(v string)`

SetExpected sets Expected field to given value.

### HasExpected

`func (o *CaseSpecial) HasExpected() bool`

HasExpected returns a boolean if a field has been set.

### SetExpectedNil

`func (o *CaseSpecial) SetExpectedNil(b bool)`

 SetExpectedNil sets the value for Expected to be an explicit nil

### UnsetExpected
`func (o *CaseSpecial) UnsetExpected()`

UnsetExpected ensures that no value is present for Expected, not even an explicit nil
### GetId

`func (o *CaseSpecial) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaseSpecial) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaseSpecial) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CaseSpecial) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CaseSpecial) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CaseSpecial) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPreconds

`func (o *CaseSpecial) GetPreconds() string`

GetPreconds returns the Preconds field if non-nil, zero value otherwise.

### GetPrecondsOk

`func (o *CaseSpecial) GetPrecondsOk() (*string, bool)`

GetPrecondsOk returns a tuple with the Preconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconds

`func (o *CaseSpecial) SetPreconds(v string)`

SetPreconds sets Preconds field to given value.

### HasPreconds

`func (o *CaseSpecial) HasPreconds() bool`

HasPreconds returns a boolean if a field has been set.

### SetPrecondsNil

`func (o *CaseSpecial) SetPrecondsNil(b bool)`

 SetPrecondsNil sets the value for Preconds to be an explicit nil

### UnsetPreconds
`func (o *CaseSpecial) UnsetPreconds()`

UnsetPreconds ensures that no value is present for Preconds, not even an explicit nil
### GetPriority

`func (o *CaseSpecial) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CaseSpecial) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CaseSpecial) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CaseSpecial) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *CaseSpecial) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *CaseSpecial) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetSectionId

`func (o *CaseSpecial) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *CaseSpecial) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *CaseSpecial) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.

### HasSectionId

`func (o *CaseSpecial) HasSectionId() bool`

HasSectionId returns a boolean if a field has been set.

### SetSectionIdNil

`func (o *CaseSpecial) SetSectionIdNil(b bool)`

 SetSectionIdNil sets the value for SectionId to be an explicit nil

### UnsetSectionId
`func (o *CaseSpecial) UnsetSectionId()`

UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
### GetSort

`func (o *CaseSpecial) GetSort() int32`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *CaseSpecial) GetSortOk() (*int32, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *CaseSpecial) SetSort(v int32)`

SetSort sets Sort field to given value.

### HasSort

`func (o *CaseSpecial) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *CaseSpecial) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *CaseSpecial) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetSteps

`func (o *CaseSpecial) GetSteps() string`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *CaseSpecial) GetStepsOk() (*string, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *CaseSpecial) SetSteps(v string)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *CaseSpecial) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *CaseSpecial) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *CaseSpecial) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetTemplateType

`func (o *CaseSpecial) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *CaseSpecial) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *CaseSpecial) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.

### HasTemplateType

`func (o *CaseSpecial) HasTemplateType() bool`

HasTemplateType returns a boolean if a field has been set.

### SetTemplateTypeNil

`func (o *CaseSpecial) SetTemplateTypeNil(b bool)`

 SetTemplateTypeNil sets the value for TemplateType to be an explicit nil

### UnsetTemplateType
`func (o *CaseSpecial) UnsetTemplateType()`

UnsetTemplateType ensures that no value is present for TemplateType, not even an explicit nil
### GetTitle

`func (o *CaseSpecial) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CaseSpecial) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CaseSpecial) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CaseSpecial) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CaseSpecial) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CaseSpecial) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUpdatedAt

`func (o *CaseSpecial) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CaseSpecial) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CaseSpecial) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CaseSpecial) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *CaseSpecial) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CaseSpecial) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


