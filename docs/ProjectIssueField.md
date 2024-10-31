# ProjectIssueField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | 创建时间戳 | [optional] 
**IssueField** | Pointer to [**IssueField**](IssueField.md) |  | [optional] 
**IssueFieldId** | Pointer to **int64** | 关联属性 ID | [optional] 
**IssueType** | Pointer to **string** | 事项类型 | [optional] [default to ""]
**NeedDefault** | Pointer to **bool** | 是否有默认值 | [optional] [default to false]
**Required** | Pointer to **bool** | 是否必填 | [optional] [default to false]
**UpdatedAt** | Pointer to **int64** | 修改时间戳 | [optional] 
**ValueString** | Pointer to **NullableString** | 默认值， JSON 字符串。例如：{\&quot;type\&quot;:\&quot;VARIABLE\&quot;,\&quot;value\&quot;:\&quot;CREATOR\&quot;}  type：默认值类型，VARIABLE（变量）、CONSTANT（常量）；  value：默认值，根据 IssueField.ComponentType，可能为不同类型的值（数值、字符串、数组）；常量值 CREATOR 表示创建者 | [optional] [default to ""]

## Methods

### NewProjectIssueField

`func NewProjectIssueField() *ProjectIssueField`

NewProjectIssueField instantiates a new ProjectIssueField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectIssueFieldWithDefaults

`func NewProjectIssueFieldWithDefaults() *ProjectIssueField`

NewProjectIssueFieldWithDefaults instantiates a new ProjectIssueField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ProjectIssueField) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectIssueField) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectIssueField) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectIssueField) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIssueField

`func (o *ProjectIssueField) GetIssueField() IssueField`

GetIssueField returns the IssueField field if non-nil, zero value otherwise.

### GetIssueFieldOk

`func (o *ProjectIssueField) GetIssueFieldOk() (*IssueField, bool)`

GetIssueFieldOk returns a tuple with the IssueField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueField

`func (o *ProjectIssueField) SetIssueField(v IssueField)`

SetIssueField sets IssueField field to given value.

### HasIssueField

`func (o *ProjectIssueField) HasIssueField() bool`

HasIssueField returns a boolean if a field has been set.

### GetIssueFieldId

`func (o *ProjectIssueField) GetIssueFieldId() int64`

GetIssueFieldId returns the IssueFieldId field if non-nil, zero value otherwise.

### GetIssueFieldIdOk

`func (o *ProjectIssueField) GetIssueFieldIdOk() (*int64, bool)`

GetIssueFieldIdOk returns a tuple with the IssueFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueFieldId

`func (o *ProjectIssueField) SetIssueFieldId(v int64)`

SetIssueFieldId sets IssueFieldId field to given value.

### HasIssueFieldId

`func (o *ProjectIssueField) HasIssueFieldId() bool`

HasIssueFieldId returns a boolean if a field has been set.

### GetIssueType

`func (o *ProjectIssueField) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *ProjectIssueField) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *ProjectIssueField) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *ProjectIssueField) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetNeedDefault

`func (o *ProjectIssueField) GetNeedDefault() bool`

GetNeedDefault returns the NeedDefault field if non-nil, zero value otherwise.

### GetNeedDefaultOk

`func (o *ProjectIssueField) GetNeedDefaultOk() (*bool, bool)`

GetNeedDefaultOk returns a tuple with the NeedDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedDefault

`func (o *ProjectIssueField) SetNeedDefault(v bool)`

SetNeedDefault sets NeedDefault field to given value.

### HasNeedDefault

`func (o *ProjectIssueField) HasNeedDefault() bool`

HasNeedDefault returns a boolean if a field has been set.

### GetRequired

`func (o *ProjectIssueField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ProjectIssueField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ProjectIssueField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ProjectIssueField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectIssueField) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectIssueField) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectIssueField) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectIssueField) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValueString

`func (o *ProjectIssueField) GetValueString() string`

GetValueString returns the ValueString field if non-nil, zero value otherwise.

### GetValueStringOk

`func (o *ProjectIssueField) GetValueStringOk() (*string, bool)`

GetValueStringOk returns a tuple with the ValueString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueString

`func (o *ProjectIssueField) SetValueString(v string)`

SetValueString sets ValueString field to given value.

### HasValueString

`func (o *ProjectIssueField) HasValueString() bool`

HasValueString returns a boolean if a field has been set.

### SetValueStringNil

`func (o *ProjectIssueField) SetValueStringNil(b bool)`

 SetValueStringNil sets the value for ValueString to be an explicit nil

### UnsetValueString
`func (o *ProjectIssueField) UnsetValueString()`

UnsetValueString ensures that no value is present for ValueString, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


