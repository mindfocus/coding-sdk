# IssueSimpleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to [**User**](User.md) |  | [optional] 
**Code** | Pointer to **int64** | 事项 Code | [optional] 
**IssueStatusId** | Pointer to **int64** | 事项状态 Id | [optional] 
**IssueStatusName** | Pointer to **string** | 事项状态名称 | [optional] [default to ""]
**IssueStatusType** | Pointer to **string** | 事项状态类型：  TODO｜PROCESSING｜ COMPLETED | [optional] [default to ""]
**IssueTypeDetail** | Pointer to [**IssueTypeDetail**](IssueTypeDetail.md) |  | [optional] 
**Name** | Pointer to **string** | 名称 | [optional] [default to ""]
**Priority** | Pointer to **NullableString** | 优先级：  \&quot;0\&quot; - 低，  \&quot;1\&quot; - 中，  \&quot;2\&quot; - 高，  \&quot;3\&quot; - 紧急，  \&quot;\&quot; - 未指定 | [optional] [default to ""]
**Type** | Pointer to **string** | 事项类型：  DEFECT - 缺陷，  REQUIREMENT - 需求，  MISSION - 任务，  EPIC - 史诗，  SUB_TASK - 子工作项 | [optional] [default to ""]

## Methods

### NewIssueSimpleData

`func NewIssueSimpleData() *IssueSimpleData`

NewIssueSimpleData instantiates a new IssueSimpleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueSimpleDataWithDefaults

`func NewIssueSimpleDataWithDefaults() *IssueSimpleData`

NewIssueSimpleDataWithDefaults instantiates a new IssueSimpleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *IssueSimpleData) GetAssignee() User`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *IssueSimpleData) GetAssigneeOk() (*User, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *IssueSimpleData) SetAssignee(v User)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *IssueSimpleData) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetCode

`func (o *IssueSimpleData) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IssueSimpleData) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IssueSimpleData) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *IssueSimpleData) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetIssueStatusId

`func (o *IssueSimpleData) GetIssueStatusId() int64`

GetIssueStatusId returns the IssueStatusId field if non-nil, zero value otherwise.

### GetIssueStatusIdOk

`func (o *IssueSimpleData) GetIssueStatusIdOk() (*int64, bool)`

GetIssueStatusIdOk returns a tuple with the IssueStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueStatusId

`func (o *IssueSimpleData) SetIssueStatusId(v int64)`

SetIssueStatusId sets IssueStatusId field to given value.

### HasIssueStatusId

`func (o *IssueSimpleData) HasIssueStatusId() bool`

HasIssueStatusId returns a boolean if a field has been set.

### GetIssueStatusName

`func (o *IssueSimpleData) GetIssueStatusName() string`

GetIssueStatusName returns the IssueStatusName field if non-nil, zero value otherwise.

### GetIssueStatusNameOk

`func (o *IssueSimpleData) GetIssueStatusNameOk() (*string, bool)`

GetIssueStatusNameOk returns a tuple with the IssueStatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueStatusName

`func (o *IssueSimpleData) SetIssueStatusName(v string)`

SetIssueStatusName sets IssueStatusName field to given value.

### HasIssueStatusName

`func (o *IssueSimpleData) HasIssueStatusName() bool`

HasIssueStatusName returns a boolean if a field has been set.

### GetIssueStatusType

`func (o *IssueSimpleData) GetIssueStatusType() string`

GetIssueStatusType returns the IssueStatusType field if non-nil, zero value otherwise.

### GetIssueStatusTypeOk

`func (o *IssueSimpleData) GetIssueStatusTypeOk() (*string, bool)`

GetIssueStatusTypeOk returns a tuple with the IssueStatusType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueStatusType

`func (o *IssueSimpleData) SetIssueStatusType(v string)`

SetIssueStatusType sets IssueStatusType field to given value.

### HasIssueStatusType

`func (o *IssueSimpleData) HasIssueStatusType() bool`

HasIssueStatusType returns a boolean if a field has been set.

### GetIssueTypeDetail

`func (o *IssueSimpleData) GetIssueTypeDetail() IssueTypeDetail`

GetIssueTypeDetail returns the IssueTypeDetail field if non-nil, zero value otherwise.

### GetIssueTypeDetailOk

`func (o *IssueSimpleData) GetIssueTypeDetailOk() (*IssueTypeDetail, bool)`

GetIssueTypeDetailOk returns a tuple with the IssueTypeDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeDetail

`func (o *IssueSimpleData) SetIssueTypeDetail(v IssueTypeDetail)`

SetIssueTypeDetail sets IssueTypeDetail field to given value.

### HasIssueTypeDetail

`func (o *IssueSimpleData) HasIssueTypeDetail() bool`

HasIssueTypeDetail returns a boolean if a field has been set.

### GetName

`func (o *IssueSimpleData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueSimpleData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueSimpleData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueSimpleData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *IssueSimpleData) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *IssueSimpleData) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *IssueSimpleData) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *IssueSimpleData) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *IssueSimpleData) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *IssueSimpleData) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetType

`func (o *IssueSimpleData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssueSimpleData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssueSimpleData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IssueSimpleData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


