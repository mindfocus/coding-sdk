# SubTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to [**User**](User.md) |  | [optional] 
**Code** | Pointer to **int64** | 子工作项 Code | [optional] 
**IssueStatusId** | Pointer to **NullableInt64** | 事项状态 Id | [optional] 
**IssueStatusName** | Pointer to **NullableString** | 事项状态名称 | [optional] [default to ""]
**Name** | Pointer to **string** | 名称 | [optional] [default to ""]
**Priority** | Pointer to **NullableString** | 优先级：  \&quot;0\&quot; - 低，  \&quot;1\&quot; - 中，  \&quot;2\&quot; - 高，  \&quot;3\&quot; - 紧急，  \&quot;\&quot; - 未指定 | [optional] [default to ""]
**Type** | Pointer to **NullableString** | 事项类型名称 | [optional] [default to ""]

## Methods

### NewSubTask

`func NewSubTask() *SubTask`

NewSubTask instantiates a new SubTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubTaskWithDefaults

`func NewSubTaskWithDefaults() *SubTask`

NewSubTaskWithDefaults instantiates a new SubTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *SubTask) GetAssignee() User`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *SubTask) GetAssigneeOk() (*User, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *SubTask) SetAssignee(v User)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *SubTask) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCode

`func (o *SubTask) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SubTask) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SubTask) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *SubTask) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIssueStatusId

`func (o *SubTask) GetIssueStatusId() int64`

GetIssueStatusId returns the IssueStatusId field if non-nil, zero value otherwise.

### GetIssueStatusIdOk

`func (o *SubTask) GetIssueStatusIdOk() (*int64, bool)`

GetIssueStatusIdOk returns a tuple with the IssueStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueStatusId

`func (o *SubTask) SetIssueStatusId(v int64)`

SetIssueStatusId sets IssueStatusId field to given value.

### HasIssueStatusId

`func (o *SubTask) HasIssueStatusId() bool`

HasIssueStatusId returns a boolean if a field has been set.

### SetIssueStatusIdNil

`func (o *SubTask) SetIssueStatusIdNil(b bool)`

 SetIssueStatusIdNil sets the value for IssueStatusId to be an explicit nil

### UnsetIssueStatusId
`func (o *SubTask) UnsetIssueStatusId()`

UnsetIssueStatusId ensures that no value is present for IssueStatusId, not even an explicit nil
### GetIssueStatusName

`func (o *SubTask) GetIssueStatusName() string`

GetIssueStatusName returns the IssueStatusName field if non-nil, zero value otherwise.

### GetIssueStatusNameOk

`func (o *SubTask) GetIssueStatusNameOk() (*string, bool)`

GetIssueStatusNameOk returns a tuple with the IssueStatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueStatusName

`func (o *SubTask) SetIssueStatusName(v string)`

SetIssueStatusName sets IssueStatusName field to given value.

### HasIssueStatusName

`func (o *SubTask) HasIssueStatusName() bool`

HasIssueStatusName returns a boolean if a field has been set.

### SetIssueStatusNameNil

`func (o *SubTask) SetIssueStatusNameNil(b bool)`

 SetIssueStatusNameNil sets the value for IssueStatusName to be an explicit nil

### UnsetIssueStatusName
`func (o *SubTask) UnsetIssueStatusName()`

UnsetIssueStatusName ensures that no value is present for IssueStatusName, not even an explicit nil
### GetName

`func (o *SubTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SubTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *SubTask) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SubTask) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SubTask) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SubTask) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *SubTask) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *SubTask) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetType

`func (o *SubTask) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubTask) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubTask) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubTask) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SubTask) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SubTask) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


