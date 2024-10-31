# IssueField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentType** | Pointer to **string** | 字段空间类型  TEXT_SINGLE_LINE - 单行文本  TEXT_MULTI_LINE - 多行文本  SELECT_SINGLE - 单选列表  SELECT_MULTI - 多选列表  RADIO - 单选框  CHECKBOX - 多选框  SELECT_MEMBER_SINGLE - 用户单选列表  SELECT_MEMBER_MULTI - 用户多选列表  TEXT_DATETIME - 日期时间选择框  TEXT_DATE - 日期选择框  TEXT_INTEGER - 整数输入框  TEXT_DECIMAL - 小数输入框 | [optional] [default to ""]
**CreatedAt** | Pointer to **int64** | 创建时间戳 | [optional] 
**CreatedBy** | Pointer to **int64** | 创建者 ID | [optional] 
**Deletable** | Pointer to **bool** | 是否可删除 | [optional] [default to false]
**Description** | Pointer to **NullableString** | 描述 | [optional] [default to ""]
**Editable** | Pointer to **bool** | 是否可修改 | [optional] [default to false]
**IconUrl** | Pointer to **NullableString** | 图标地址 | [optional] [default to ""]
**Id** | Pointer to **int64** | 属性 ID | [optional] 
**Name** | Pointer to **string** | 名称 | [optional] [default to ""]
**Options** | Pointer to [**[]IssueFieldOption**](IssueFieldOption.md) | 选项列表 | [optional] 
**Required** | Pointer to **bool** | 是否必填 | [optional] [default to false]
**Selectable** | Pointer to **bool** | 项目中属性可选 | [optional] [default to false]
**Sortable** | Pointer to **bool** | 是否可排序 | [optional] [default to false]
**Type** | Pointer to **string** | 字段类型 | [optional] [default to ""]
**Unit** | Pointer to **NullableString** | 单位 | [optional] [default to ""]
**UpdatedAt** | Pointer to **int64** | 修改时间戳 | [optional] 

## Methods

### NewIssueField

`func NewIssueField() *IssueField`

NewIssueField instantiates a new IssueField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueFieldWithDefaults

`func NewIssueFieldWithDefaults() *IssueField`

NewIssueFieldWithDefaults instantiates a new IssueField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentType

`func (o *IssueField) GetComponentType() string`

GetComponentType returns the ComponentType field if non-nil, zero value otherwise.

### GetComponentTypeOk

`func (o *IssueField) GetComponentTypeOk() (*string, bool)`

GetComponentTypeOk returns a tuple with the ComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentType

`func (o *IssueField) SetComponentType(v string)`

SetComponentType sets ComponentType field to given value.

### HasComponentType

`func (o *IssueField) HasComponentType() bool`

HasComponentType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IssueField) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueField) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueField) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssueField) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *IssueField) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *IssueField) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *IssueField) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *IssueField) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDeletable

`func (o *IssueField) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *IssueField) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *IssueField) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *IssueField) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDescription

`func (o *IssueField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IssueField) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IssueField) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEditable

`func (o *IssueField) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *IssueField) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *IssueField) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *IssueField) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetIconUrl

`func (o *IssueField) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *IssueField) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *IssueField) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *IssueField) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *IssueField) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *IssueField) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetId

`func (o *IssueField) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueField) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueField) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IssueField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IssueField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *IssueField) GetOptions() []IssueFieldOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *IssueField) GetOptionsOk() (*[]IssueFieldOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *IssueField) SetOptions(v []IssueFieldOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *IssueField) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *IssueField) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *IssueField) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetRequired

`func (o *IssueField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *IssueField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *IssueField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *IssueField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSelectable

`func (o *IssueField) GetSelectable() bool`

GetSelectable returns the Selectable field if non-nil, zero value otherwise.

### GetSelectableOk

`func (o *IssueField) GetSelectableOk() (*bool, bool)`

GetSelectableOk returns a tuple with the Selectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectable

`func (o *IssueField) SetSelectable(v bool)`

SetSelectable sets Selectable field to given value.

### HasSelectable

`func (o *IssueField) HasSelectable() bool`

HasSelectable returns a boolean if a field has been set.

### GetSortable

`func (o *IssueField) GetSortable() bool`

GetSortable returns the Sortable field if non-nil, zero value otherwise.

### GetSortableOk

`func (o *IssueField) GetSortableOk() (*bool, bool)`

GetSortableOk returns a tuple with the Sortable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortable

`func (o *IssueField) SetSortable(v bool)`

SetSortable sets Sortable field to given value.

### HasSortable

`func (o *IssueField) HasSortable() bool`

HasSortable returns a boolean if a field has been set.

### GetType

`func (o *IssueField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssueField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssueField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IssueField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnit

`func (o *IssueField) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *IssueField) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *IssueField) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *IssueField) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *IssueField) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *IssueField) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetUpdatedAt

`func (o *IssueField) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssueField) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssueField) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IssueField) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


