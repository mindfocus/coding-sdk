# IssueCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConstValue** | Pointer to **string** | 常量值  \&quot;UNSPECIFIC\&quot; 表示未指定，例如：处理人、需求类型、标签等字段，使用此值来筛选未指定处理人、需求类型或标签的事项。 | [optional] [default to ""]
**CustomFieldId** | Pointer to **int64** | 自定义字段 Id，Key 为 \&quot;CUSTOM\&quot; 时需设置该值，常规字段无需设置。 | [optional] 
**CustomFieldName** | Pointer to **string** | 自定义字段名称 | [optional] [default to ""]
**Key** | **string** | 筛选字段 KEY，可选值如下。  DEFECT_TYPE：缺陷类型，多选  REQUIREMENT_TYPE：需求类型，多选  MISSION_TYPE：任务类型，多选  PRIORITY：优先级，多选  DUE_DATE：截止日期，日期范围  UPDATED_AT：更新时间，日期范围  CREATED_AT：创建时间，日期范围  START_DATE：开始日期，日期范围  ASSIGNEE：处理人ID，多选  CREATOR：创建者ID，多选  WATCHER：关注者ID，多选  MODULE：模块，多选  LABEL：标签，多选  STATUS：状态，多选  STATUS_TYPE：状态类型，多选  KEYWORD：事项名称、CODE 模糊搜索  ISSUE_TYPE：事项类型，多选  ISSUE_SUB_TYPE：事项的子项类型，多选  WORKING_HOURS：预估工时，数值范围  ITERATION：迭代，多选  PARENT：父需求，多选  CUSTOM：自定义字段，同时需指定 CustomFieldId | [default to ""]
**Value** | Pointer to **string** | 筛选值，多选值用逗号隔开。  日期格式: 2020-08-01，日期时间格式: 2020-08-01 12:00:00  整数、小数、日期、日期时间类型的字段值应为一个范围，前后闭区间，范围开始值与结束值之间使用\&quot;_\&quot;连接，例如：\&quot;0.1_5.0\&quot;、\&quot;2020-08-01_2020-08-31\&quot;。 | [optional] [default to ""]

## Methods

### NewIssueCondition

`func NewIssueCondition(key string, ) *IssueCondition`

NewIssueCondition instantiates a new IssueCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueConditionWithDefaults

`func NewIssueConditionWithDefaults() *IssueCondition`

NewIssueConditionWithDefaults instantiates a new IssueCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstValue

`func (o *IssueCondition) GetConstValue() string`

GetConstValue returns the ConstValue field if non-nil, zero value otherwise.

### GetConstValueOk

`func (o *IssueCondition) GetConstValueOk() (*string, bool)`

GetConstValueOk returns a tuple with the ConstValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstValue

`func (o *IssueCondition) SetConstValue(v string)`

SetConstValue sets ConstValue field to given value.

### HasConstValue

`func (o *IssueCondition) HasConstValue() bool`

HasConstValue returns a boolean if a field has been set.

### GetCustomFieldId

`func (o *IssueCondition) GetCustomFieldId() int64`

GetCustomFieldId returns the CustomFieldId field if non-nil, zero value otherwise.

### GetCustomFieldIdOk

`func (o *IssueCondition) GetCustomFieldIdOk() (*int64, bool)`

GetCustomFieldIdOk returns a tuple with the CustomFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldId

`func (o *IssueCondition) SetCustomFieldId(v int64)`

SetCustomFieldId sets CustomFieldId field to given value.

### HasCustomFieldId

`func (o *IssueCondition) HasCustomFieldId() bool`

HasCustomFieldId returns a boolean if a field has been set.

### GetCustomFieldName

`func (o *IssueCondition) GetCustomFieldName() string`

GetCustomFieldName returns the CustomFieldName field if non-nil, zero value otherwise.

### GetCustomFieldNameOk

`func (o *IssueCondition) GetCustomFieldNameOk() (*string, bool)`

GetCustomFieldNameOk returns a tuple with the CustomFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldName

`func (o *IssueCondition) SetCustomFieldName(v string)`

SetCustomFieldName sets CustomFieldName field to given value.

### HasCustomFieldName

`func (o *IssueCondition) HasCustomFieldName() bool`

HasCustomFieldName returns a boolean if a field has been set.

### GetKey

`func (o *IssueCondition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IssueCondition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IssueCondition) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *IssueCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IssueCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IssueCondition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IssueCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


