# CustomFieldChangeLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | Pointer to **NullableString** | 动作类型  | [optional] [default to ""]
**Creator** | Pointer to **NullableInt64** | 创建人ID  | [optional] 
**FieldId** | Pointer to **NullableInt64** | 自定义属性ID  | [optional] 
**FieldName** | Pointer to **NullableString** | 自定义属性名字  | [optional] [default to ""]
**FieldType** | Pointer to **NullableString** | 自定义属性类型  | [optional] [default to ""]
**FieldValue** | Pointer to **NullableString** | 自定义属性的值  | [optional] [default to ""]
**IssueId** | Pointer to **NullableInt64** | 事项id（不是页面上的ID，是数据库中的唯一ID） | [optional] 
**UpdatedAt** | Pointer to **NullableInt64** | 更新时间戳  | [optional] 

## Methods

### NewCustomFieldChangeLog

`func NewCustomFieldChangeLog() *CustomFieldChangeLog`

NewCustomFieldChangeLog instantiates a new CustomFieldChangeLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldChangeLogWithDefaults

`func NewCustomFieldChangeLogWithDefaults() *CustomFieldChangeLog`

NewCustomFieldChangeLogWithDefaults instantiates a new CustomFieldChangeLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *CustomFieldChangeLog) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *CustomFieldChangeLog) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *CustomFieldChangeLog) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *CustomFieldChangeLog) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### SetActionTypeNil

`func (o *CustomFieldChangeLog) SetActionTypeNil(b bool)`

 SetActionTypeNil sets the value for ActionType to be an explicit nil

### UnsetActionType
`func (o *CustomFieldChangeLog) UnsetActionType()`

UnsetActionType ensures that no value is present for ActionType, not even an explicit nil
### GetCreator

`func (o *CustomFieldChangeLog) GetCreator() int64`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *CustomFieldChangeLog) GetCreatorOk() (*int64, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *CustomFieldChangeLog) SetCreator(v int64)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *CustomFieldChangeLog) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### SetCreatorNil

`func (o *CustomFieldChangeLog) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *CustomFieldChangeLog) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetFieldId

`func (o *CustomFieldChangeLog) GetFieldId() int64`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *CustomFieldChangeLog) GetFieldIdOk() (*int64, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *CustomFieldChangeLog) SetFieldId(v int64)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *CustomFieldChangeLog) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### SetFieldIdNil

`func (o *CustomFieldChangeLog) SetFieldIdNil(b bool)`

 SetFieldIdNil sets the value for FieldId to be an explicit nil

### UnsetFieldId
`func (o *CustomFieldChangeLog) UnsetFieldId()`

UnsetFieldId ensures that no value is present for FieldId, not even an explicit nil
### GetFieldName

`func (o *CustomFieldChangeLog) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *CustomFieldChangeLog) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *CustomFieldChangeLog) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *CustomFieldChangeLog) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### SetFieldNameNil

`func (o *CustomFieldChangeLog) SetFieldNameNil(b bool)`

 SetFieldNameNil sets the value for FieldName to be an explicit nil

### UnsetFieldName
`func (o *CustomFieldChangeLog) UnsetFieldName()`

UnsetFieldName ensures that no value is present for FieldName, not even an explicit nil
### GetFieldType

`func (o *CustomFieldChangeLog) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *CustomFieldChangeLog) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *CustomFieldChangeLog) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *CustomFieldChangeLog) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### SetFieldTypeNil

`func (o *CustomFieldChangeLog) SetFieldTypeNil(b bool)`

 SetFieldTypeNil sets the value for FieldType to be an explicit nil

### UnsetFieldType
`func (o *CustomFieldChangeLog) UnsetFieldType()`

UnsetFieldType ensures that no value is present for FieldType, not even an explicit nil
### GetFieldValue

`func (o *CustomFieldChangeLog) GetFieldValue() string`

GetFieldValue returns the FieldValue field if non-nil, zero value otherwise.

### GetFieldValueOk

`func (o *CustomFieldChangeLog) GetFieldValueOk() (*string, bool)`

GetFieldValueOk returns a tuple with the FieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValue

`func (o *CustomFieldChangeLog) SetFieldValue(v string)`

SetFieldValue sets FieldValue field to given value.

### HasFieldValue

`func (o *CustomFieldChangeLog) HasFieldValue() bool`

HasFieldValue returns a boolean if a field has been set.

### SetFieldValueNil

`func (o *CustomFieldChangeLog) SetFieldValueNil(b bool)`

 SetFieldValueNil sets the value for FieldValue to be an explicit nil

### UnsetFieldValue
`func (o *CustomFieldChangeLog) UnsetFieldValue()`

UnsetFieldValue ensures that no value is present for FieldValue, not even an explicit nil
### GetIssueId

`func (o *CustomFieldChangeLog) GetIssueId() int64`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *CustomFieldChangeLog) GetIssueIdOk() (*int64, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *CustomFieldChangeLog) SetIssueId(v int64)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *CustomFieldChangeLog) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### SetIssueIdNil

`func (o *CustomFieldChangeLog) SetIssueIdNil(b bool)`

 SetIssueIdNil sets the value for IssueId to be an explicit nil

### UnsetIssueId
`func (o *CustomFieldChangeLog) UnsetIssueId()`

UnsetIssueId ensures that no value is present for IssueId, not even an explicit nil
### GetUpdatedAt

`func (o *CustomFieldChangeLog) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomFieldChangeLog) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomFieldChangeLog) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CustomFieldChangeLog) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *CustomFieldChangeLog) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CustomFieldChangeLog) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


