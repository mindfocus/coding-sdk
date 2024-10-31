# IssueCustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 自定义属性 Id | [optional] 
**Name** | Pointer to **string** | 自定义属性名称 | [optional] [default to ""]
**ValueString** | Pointer to **string** | 自定义属性值  根据自定义属性的 ComponentType，返回不同类型的值：  单个成员选择（SELECT_MEMBER_SINGLE）为 User 对象 JSON，  多个成员选择（SELECT_MEMBER_MULTI）为 User 数组 JSON，  多选菜单（SELECT_MULTI）为 String 数组 JSON，  其余类型为普通字符串 | [optional] [default to ""]

## Methods

### NewIssueCustomField

`func NewIssueCustomField() *IssueCustomField`

NewIssueCustomField instantiates a new IssueCustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueCustomFieldWithDefaults

`func NewIssueCustomFieldWithDefaults() *IssueCustomField`

NewIssueCustomFieldWithDefaults instantiates a new IssueCustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssueCustomField) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueCustomField) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueCustomField) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IssueCustomField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IssueCustomField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueCustomField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueCustomField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueCustomField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValueString

`func (o *IssueCustomField) GetValueString() string`

GetValueString returns the ValueString field if non-nil, zero value otherwise.

### GetValueStringOk

`func (o *IssueCustomField) GetValueStringOk() (*string, bool)`

GetValueStringOk returns a tuple with the ValueString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueString

`func (o *IssueCustomField) SetValueString(v string)`

SetValueString sets ValueString field to given value.

### HasValueString

`func (o *IssueCustomField) HasValueString() bool`

HasValueString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


