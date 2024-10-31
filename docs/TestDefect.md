# TestDefect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedTo** | Pointer to **NullableString** | 处理人 | [optional] [default to ""]
**Author** | Pointer to **NullableString** | 创建人 | [optional] [default to ""]
**CreatedAt** | Pointer to **NullableString** | 创建时间 | [optional] [default to ""]
**Description** | Pointer to **NullableString** | 描述 | [optional] [default to ""]
**Id** | Pointer to **NullableInt32** | 缺陷 ID | [optional] 
**Name** | Pointer to **NullableString** | 标题 | [optional] [default to ""]
**Status** | Pointer to **NullableString** | 状态 | [optional] [default to ""]
**StatusName** | Pointer to **NullableString** | 状态名称 | [optional] [default to ""]

## Methods

### NewTestDefect

`func NewTestDefect() *TestDefect`

NewTestDefect instantiates a new TestDefect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestDefectWithDefaults

`func NewTestDefectWithDefaults() *TestDefect`

NewTestDefectWithDefaults instantiates a new TestDefect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedTo

`func (o *TestDefect) GetAssignedTo() string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *TestDefect) GetAssignedToOk() (*string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *TestDefect) SetAssignedTo(v string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *TestDefect) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### SetAssignedToNil

`func (o *TestDefect) SetAssignedToNil(b bool)`

 SetAssignedToNil sets the value for AssignedTo to be an explicit nil

### UnsetAssignedTo
`func (o *TestDefect) UnsetAssignedTo()`

UnsetAssignedTo ensures that no value is present for AssignedTo, not even an explicit nil
### GetAuthor

`func (o *TestDefect) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *TestDefect) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *TestDefect) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *TestDefect) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *TestDefect) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *TestDefect) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetCreatedAt

`func (o *TestDefect) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TestDefect) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TestDefect) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TestDefect) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *TestDefect) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *TestDefect) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDescription

`func (o *TestDefect) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TestDefect) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TestDefect) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TestDefect) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TestDefect) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TestDefect) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *TestDefect) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TestDefect) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TestDefect) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TestDefect) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TestDefect) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TestDefect) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *TestDefect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestDefect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestDefect) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestDefect) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TestDefect) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TestDefect) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *TestDefect) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestDefect) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestDefect) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestDefect) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TestDefect) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TestDefect) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusName

`func (o *TestDefect) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *TestDefect) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *TestDefect) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *TestDefect) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### SetStatusNameNil

`func (o *TestDefect) SetStatusNameNil(b bool)`

 SetStatusNameNil sets the value for StatusName to be an explicit nil

### UnsetStatusName
`func (o *TestDefect) UnsetStatusName()`

UnsetStatusName ensures that no value is present for StatusName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


