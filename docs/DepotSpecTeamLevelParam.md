# DepotSpecTeamLevelParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPushWildRef** | **bool** | 允许创建规定分支类型以外的分支 | [default to false]
**DepotBranchTypeList** | Pointer to [**[]DepotBranchTypeParam**](DepotBranchTypeParam.md) | 分支类型列表 | [optional] 
**DepotMergeRequestRuleList** | Pointer to [**[]DepotMergeRequestRuleParam**](DepotMergeRequestRuleParam.md) | 合并方向规则列表 | [optional] 
**Description** | Pointer to **string** | 仓库规范描述 | [optional] [default to ""]
**Name** | **string** | 规范的名字唯一，当名字是已有规范的名字时，为修改；当名字不是已有规范名字时为新增 | [default to ""]
**ReName** | Pointer to **string** | 当需要修改已有规范的名字时，需要填写新名字 | [optional] [default to ""]

## Methods

### NewDepotSpecTeamLevelParam

`func NewDepotSpecTeamLevelParam(allowPushWildRef bool, name string, ) *DepotSpecTeamLevelParam`

NewDepotSpecTeamLevelParam instantiates a new DepotSpecTeamLevelParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepotSpecTeamLevelParamWithDefaults

`func NewDepotSpecTeamLevelParamWithDefaults() *DepotSpecTeamLevelParam`

NewDepotSpecTeamLevelParamWithDefaults instantiates a new DepotSpecTeamLevelParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPushWildRef

`func (o *DepotSpecTeamLevelParam) GetAllowPushWildRef() bool`

GetAllowPushWildRef returns the AllowPushWildRef field if non-nil, zero value otherwise.

### GetAllowPushWildRefOk

`func (o *DepotSpecTeamLevelParam) GetAllowPushWildRefOk() (*bool, bool)`

GetAllowPushWildRefOk returns a tuple with the AllowPushWildRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPushWildRef

`func (o *DepotSpecTeamLevelParam) SetAllowPushWildRef(v bool)`

SetAllowPushWildRef sets AllowPushWildRef field to given value.


### GetDepotBranchTypeList

`func (o *DepotSpecTeamLevelParam) GetDepotBranchTypeList() []DepotBranchTypeParam`

GetDepotBranchTypeList returns the DepotBranchTypeList field if non-nil, zero value otherwise.

### GetDepotBranchTypeListOk

`func (o *DepotSpecTeamLevelParam) GetDepotBranchTypeListOk() (*[]DepotBranchTypeParam, bool)`

GetDepotBranchTypeListOk returns a tuple with the DepotBranchTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotBranchTypeList

`func (o *DepotSpecTeamLevelParam) SetDepotBranchTypeList(v []DepotBranchTypeParam)`

SetDepotBranchTypeList sets DepotBranchTypeList field to given value.

### HasDepotBranchTypeList

`func (o *DepotSpecTeamLevelParam) HasDepotBranchTypeList() bool`

HasDepotBranchTypeList returns a boolean if a field has been set.

### GetDepotMergeRequestRuleList

`func (o *DepotSpecTeamLevelParam) GetDepotMergeRequestRuleList() []DepotMergeRequestRuleParam`

GetDepotMergeRequestRuleList returns the DepotMergeRequestRuleList field if non-nil, zero value otherwise.

### GetDepotMergeRequestRuleListOk

`func (o *DepotSpecTeamLevelParam) GetDepotMergeRequestRuleListOk() (*[]DepotMergeRequestRuleParam, bool)`

GetDepotMergeRequestRuleListOk returns a tuple with the DepotMergeRequestRuleList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotMergeRequestRuleList

`func (o *DepotSpecTeamLevelParam) SetDepotMergeRequestRuleList(v []DepotMergeRequestRuleParam)`

SetDepotMergeRequestRuleList sets DepotMergeRequestRuleList field to given value.

### HasDepotMergeRequestRuleList

`func (o *DepotSpecTeamLevelParam) HasDepotMergeRequestRuleList() bool`

HasDepotMergeRequestRuleList returns a boolean if a field has been set.

### GetDescription

`func (o *DepotSpecTeamLevelParam) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DepotSpecTeamLevelParam) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DepotSpecTeamLevelParam) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DepotSpecTeamLevelParam) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *DepotSpecTeamLevelParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DepotSpecTeamLevelParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DepotSpecTeamLevelParam) SetName(v string)`

SetName sets Name field to given value.


### GetReName

`func (o *DepotSpecTeamLevelParam) GetReName() string`

GetReName returns the ReName field if non-nil, zero value otherwise.

### GetReNameOk

`func (o *DepotSpecTeamLevelParam) GetReNameOk() (*string, bool)`

GetReNameOk returns a tuple with the ReName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReName

`func (o *DepotSpecTeamLevelParam) SetReName(v string)`

SetReName sets ReName field to given value.

### HasReName

`func (o *DepotSpecTeamLevelParam) HasReName() bool`

HasReName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


