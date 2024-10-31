# IssueStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | 创建时间戳 | [optional] 
**Description** | Pointer to **NullableString** | 描述 | [optional] [default to ""]
**Id** | Pointer to **int64** | 事项状态 ID | [optional] 
**Index** | Pointer to **int64** | 状态序号 | [optional] 
**IsSystem** | Pointer to **bool** | 是否是系统内置 | [optional] [default to false]
**Name** | Pointer to **string** | 名称 | [optional] [default to ""]
**Type** | Pointer to **string** | 类型：TODO、PROCESSING、COMPLETED | [optional] [default to ""]
**UpdatedAt** | Pointer to **int64** | 修改时间戳 | [optional] 

## Methods

### NewIssueStatus

`func NewIssueStatus() *IssueStatus`

NewIssueStatus instantiates a new IssueStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueStatusWithDefaults

`func NewIssueStatusWithDefaults() *IssueStatus`

NewIssueStatusWithDefaults instantiates a new IssueStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *IssueStatus) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueStatus) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueStatus) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssueStatus) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *IssueStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IssueStatus) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IssueStatus) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *IssueStatus) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueStatus) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueStatus) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IssueStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *IssueStatus) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *IssueStatus) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *IssueStatus) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *IssueStatus) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetIsSystem

`func (o *IssueStatus) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *IssueStatus) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *IssueStatus) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *IssueStatus) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.

### GetName

`func (o *IssueStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IssueStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssueStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssueStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IssueStatus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IssueStatus) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssueStatus) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssueStatus) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IssueStatus) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


