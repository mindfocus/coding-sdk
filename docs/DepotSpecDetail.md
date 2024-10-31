# DepotSpecDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPushWildRef** | Pointer to **NullableBool** | 允许创建规定分支类型以外的分支 | [optional] [default to false]
**DepotBranchTypeList** | Pointer to [**[]DepotBranchType**](DepotBranchType.md) | 分支类型列表 | [optional] 
**DepotMergeRequestRuleList** | Pointer to [**[]DepotMergeRequestRule**](DepotMergeRequestRule.md) | 合并方向规则列表 | [optional] 
**Description** | Pointer to **NullableString** | 仓库规范描述信息 | [optional] [default to ""]
**Name** | Pointer to **NullableString** | 仓库规范名字 | [optional] [default to ""]
**Type** | Pointer to **NullableString** | system：系统级别的规范；team：团队级别的规范 | [optional] [default to ""]
**UseExistingSolution** | Pointer to **NullableBool** | 使用的是系统/团队级别的仓库规范（这个字段只在使用仓库路径查询时有用） | [optional] [default to false]

## Methods

### NewDepotSpecDetail

`func NewDepotSpecDetail() *DepotSpecDetail`

NewDepotSpecDetail instantiates a new DepotSpecDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepotSpecDetailWithDefaults

`func NewDepotSpecDetailWithDefaults() *DepotSpecDetail`

NewDepotSpecDetailWithDefaults instantiates a new DepotSpecDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPushWildRef

`func (o *DepotSpecDetail) GetAllowPushWildRef() bool`

GetAllowPushWildRef returns the AllowPushWildRef field if non-nil, zero value otherwise.

### GetAllowPushWildRefOk

`func (o *DepotSpecDetail) GetAllowPushWildRefOk() (*bool, bool)`

GetAllowPushWildRefOk returns a tuple with the AllowPushWildRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPushWildRef

`func (o *DepotSpecDetail) SetAllowPushWildRef(v bool)`

SetAllowPushWildRef sets AllowPushWildRef field to given value.

### HasAllowPushWildRef

`func (o *DepotSpecDetail) HasAllowPushWildRef() bool`

HasAllowPushWildRef returns a boolean if a field has been set.

### SetAllowPushWildRefNil

`func (o *DepotSpecDetail) SetAllowPushWildRefNil(b bool)`

 SetAllowPushWildRefNil sets the value for AllowPushWildRef to be an explicit nil

### UnsetAllowPushWildRef
`func (o *DepotSpecDetail) UnsetAllowPushWildRef()`

UnsetAllowPushWildRef ensures that no value is present for AllowPushWildRef, not even an explicit nil
### GetDepotBranchTypeList

`func (o *DepotSpecDetail) GetDepotBranchTypeList() []DepotBranchType`

GetDepotBranchTypeList returns the DepotBranchTypeList field if non-nil, zero value otherwise.

### GetDepotBranchTypeListOk

`func (o *DepotSpecDetail) GetDepotBranchTypeListOk() (*[]DepotBranchType, bool)`

GetDepotBranchTypeListOk returns a tuple with the DepotBranchTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotBranchTypeList

`func (o *DepotSpecDetail) SetDepotBranchTypeList(v []DepotBranchType)`

SetDepotBranchTypeList sets DepotBranchTypeList field to given value.

### HasDepotBranchTypeList

`func (o *DepotSpecDetail) HasDepotBranchTypeList() bool`

HasDepotBranchTypeList returns a boolean if a field has been set.

### SetDepotBranchTypeListNil

`func (o *DepotSpecDetail) SetDepotBranchTypeListNil(b bool)`

 SetDepotBranchTypeListNil sets the value for DepotBranchTypeList to be an explicit nil

### UnsetDepotBranchTypeList
`func (o *DepotSpecDetail) UnsetDepotBranchTypeList()`

UnsetDepotBranchTypeList ensures that no value is present for DepotBranchTypeList, not even an explicit nil
### GetDepotMergeRequestRuleList

`func (o *DepotSpecDetail) GetDepotMergeRequestRuleList() []DepotMergeRequestRule`

GetDepotMergeRequestRuleList returns the DepotMergeRequestRuleList field if non-nil, zero value otherwise.

### GetDepotMergeRequestRuleListOk

`func (o *DepotSpecDetail) GetDepotMergeRequestRuleListOk() (*[]DepotMergeRequestRule, bool)`

GetDepotMergeRequestRuleListOk returns a tuple with the DepotMergeRequestRuleList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotMergeRequestRuleList

`func (o *DepotSpecDetail) SetDepotMergeRequestRuleList(v []DepotMergeRequestRule)`

SetDepotMergeRequestRuleList sets DepotMergeRequestRuleList field to given value.

### HasDepotMergeRequestRuleList

`func (o *DepotSpecDetail) HasDepotMergeRequestRuleList() bool`

HasDepotMergeRequestRuleList returns a boolean if a field has been set.

### SetDepotMergeRequestRuleListNil

`func (o *DepotSpecDetail) SetDepotMergeRequestRuleListNil(b bool)`

 SetDepotMergeRequestRuleListNil sets the value for DepotMergeRequestRuleList to be an explicit nil

### UnsetDepotMergeRequestRuleList
`func (o *DepotSpecDetail) UnsetDepotMergeRequestRuleList()`

UnsetDepotMergeRequestRuleList ensures that no value is present for DepotMergeRequestRuleList, not even an explicit nil
### GetDescription

`func (o *DepotSpecDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DepotSpecDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DepotSpecDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DepotSpecDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DepotSpecDetail) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DepotSpecDetail) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *DepotSpecDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DepotSpecDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DepotSpecDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DepotSpecDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DepotSpecDetail) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DepotSpecDetail) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *DepotSpecDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DepotSpecDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DepotSpecDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DepotSpecDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *DepotSpecDetail) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *DepotSpecDetail) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUseExistingSolution

`func (o *DepotSpecDetail) GetUseExistingSolution() bool`

GetUseExistingSolution returns the UseExistingSolution field if non-nil, zero value otherwise.

### GetUseExistingSolutionOk

`func (o *DepotSpecDetail) GetUseExistingSolutionOk() (*bool, bool)`

GetUseExistingSolutionOk returns a tuple with the UseExistingSolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExistingSolution

`func (o *DepotSpecDetail) SetUseExistingSolution(v bool)`

SetUseExistingSolution sets UseExistingSolution field to given value.

### HasUseExistingSolution

`func (o *DepotSpecDetail) HasUseExistingSolution() bool`

HasUseExistingSolution returns a boolean if a field has been set.

### SetUseExistingSolutionNil

`func (o *DepotSpecDetail) SetUseExistingSolutionNil(b bool)`

 SetUseExistingSolutionNil sets the value for UseExistingSolution to be an explicit nil

### UnsetUseExistingSolution
`func (o *DepotSpecDetail) UnsetUseExistingSolution()`

UnsetUseExistingSolution ensures that no value is present for UseExistingSolution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


