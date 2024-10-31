# IssueTypeDetailWithSplit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | 描述 | [optional] [default to ""]
**Id** | Pointer to **int64** | 事项类型 ID | [optional] 
**IsSystem** | Pointer to **bool** | 是否是系统类型 | [optional] [default to false]
**IssueType** | Pointer to **string** | 事项类型大类 | [optional] [default to ""]
**Name** | Pointer to **string** | 事项类型名称 | [optional] [default to ""]
**SplitTargetIssueTypeId** | Pointer to **[]int64** | 可分解类型 ID，SplitType &#x3D; SPECIFIC_TYPE 时需指定 | [optional] 
**SplitType** | Pointer to **string** | 需求分解类型，SPECIFIC_TYPE - 可分解为制定需求类型，UNSPLITTABLE - 不可分解需求，ALL_REQUIREMENT - 可分解为全部需求类型 | [optional] [default to ""]

## Methods

### NewIssueTypeDetailWithSplit

`func NewIssueTypeDetailWithSplit() *IssueTypeDetailWithSplit`

NewIssueTypeDetailWithSplit instantiates a new IssueTypeDetailWithSplit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeDetailWithSplitWithDefaults

`func NewIssueTypeDetailWithSplitWithDefaults() *IssueTypeDetailWithSplit`

NewIssueTypeDetailWithSplitWithDefaults instantiates a new IssueTypeDetailWithSplit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IssueTypeDetailWithSplit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueTypeDetailWithSplit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueTypeDetailWithSplit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueTypeDetailWithSplit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *IssueTypeDetailWithSplit) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueTypeDetailWithSplit) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueTypeDetailWithSplit) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IssueTypeDetailWithSplit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsSystem

`func (o *IssueTypeDetailWithSplit) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *IssueTypeDetailWithSplit) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *IssueTypeDetailWithSplit) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *IssueTypeDetailWithSplit) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.

### GetIssueType

`func (o *IssueTypeDetailWithSplit) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *IssueTypeDetailWithSplit) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *IssueTypeDetailWithSplit) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *IssueTypeDetailWithSplit) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetName

`func (o *IssueTypeDetailWithSplit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueTypeDetailWithSplit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueTypeDetailWithSplit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueTypeDetailWithSplit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSplitTargetIssueTypeId

`func (o *IssueTypeDetailWithSplit) GetSplitTargetIssueTypeId() []int64`

GetSplitTargetIssueTypeId returns the SplitTargetIssueTypeId field if non-nil, zero value otherwise.

### GetSplitTargetIssueTypeIdOk

`func (o *IssueTypeDetailWithSplit) GetSplitTargetIssueTypeIdOk() (*[]int64, bool)`

GetSplitTargetIssueTypeIdOk returns a tuple with the SplitTargetIssueTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitTargetIssueTypeId

`func (o *IssueTypeDetailWithSplit) SetSplitTargetIssueTypeId(v []int64)`

SetSplitTargetIssueTypeId sets SplitTargetIssueTypeId field to given value.

### HasSplitTargetIssueTypeId

`func (o *IssueTypeDetailWithSplit) HasSplitTargetIssueTypeId() bool`

HasSplitTargetIssueTypeId returns a boolean if a field has been set.

### GetSplitType

`func (o *IssueTypeDetailWithSplit) GetSplitType() string`

GetSplitType returns the SplitType field if non-nil, zero value otherwise.

### GetSplitTypeOk

`func (o *IssueTypeDetailWithSplit) GetSplitTypeOk() (*string, bool)`

GetSplitTypeOk returns a tuple with the SplitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitType

`func (o *IssueTypeDetailWithSplit) SetSplitType(v string)`

SetSplitType sets SplitType field to given value.

### HasSplitType

`func (o *IssueTypeDetailWithSplit) HasSplitType() bool`

HasSplitType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


