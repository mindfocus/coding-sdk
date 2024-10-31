# DepotSpecDepotLevelParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPushWildRef** | Pointer to **bool** | 允许创建规定分支类型以外的分支 | [optional] [default to false]
**DepotBranchTypeList** | Pointer to [**[]DepotBranchTypeParam**](DepotBranchTypeParam.md) | 分支类型列表 | [optional] 
**DepotMergeRequestRuleList** | Pointer to [**[]DepotMergeRequestRuleParam**](DepotMergeRequestRuleParam.md) | 合并方向规则列表 | [optional] 

## Methods

### NewDepotSpecDepotLevelParam

`func NewDepotSpecDepotLevelParam() *DepotSpecDepotLevelParam`

NewDepotSpecDepotLevelParam instantiates a new DepotSpecDepotLevelParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepotSpecDepotLevelParamWithDefaults

`func NewDepotSpecDepotLevelParamWithDefaults() *DepotSpecDepotLevelParam`

NewDepotSpecDepotLevelParamWithDefaults instantiates a new DepotSpecDepotLevelParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPushWildRef

`func (o *DepotSpecDepotLevelParam) GetAllowPushWildRef() bool`

GetAllowPushWildRef returns the AllowPushWildRef field if non-nil, zero value otherwise.

### GetAllowPushWildRefOk

`func (o *DepotSpecDepotLevelParam) GetAllowPushWildRefOk() (*bool, bool)`

GetAllowPushWildRefOk returns a tuple with the AllowPushWildRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPushWildRef

`func (o *DepotSpecDepotLevelParam) SetAllowPushWildRef(v bool)`

SetAllowPushWildRef sets AllowPushWildRef field to given value.

### HasAllowPushWildRef

`func (o *DepotSpecDepotLevelParam) HasAllowPushWildRef() bool`

HasAllowPushWildRef returns a boolean if a field has been set.

### GetDepotBranchTypeList

`func (o *DepotSpecDepotLevelParam) GetDepotBranchTypeList() []DepotBranchTypeParam`

GetDepotBranchTypeList returns the DepotBranchTypeList field if non-nil, zero value otherwise.

### GetDepotBranchTypeListOk

`func (o *DepotSpecDepotLevelParam) GetDepotBranchTypeListOk() (*[]DepotBranchTypeParam, bool)`

GetDepotBranchTypeListOk returns a tuple with the DepotBranchTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotBranchTypeList

`func (o *DepotSpecDepotLevelParam) SetDepotBranchTypeList(v []DepotBranchTypeParam)`

SetDepotBranchTypeList sets DepotBranchTypeList field to given value.

### HasDepotBranchTypeList

`func (o *DepotSpecDepotLevelParam) HasDepotBranchTypeList() bool`

HasDepotBranchTypeList returns a boolean if a field has been set.

### GetDepotMergeRequestRuleList

`func (o *DepotSpecDepotLevelParam) GetDepotMergeRequestRuleList() []DepotMergeRequestRuleParam`

GetDepotMergeRequestRuleList returns the DepotMergeRequestRuleList field if non-nil, zero value otherwise.

### GetDepotMergeRequestRuleListOk

`func (o *DepotSpecDepotLevelParam) GetDepotMergeRequestRuleListOk() (*[]DepotMergeRequestRuleParam, bool)`

GetDepotMergeRequestRuleListOk returns a tuple with the DepotMergeRequestRuleList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotMergeRequestRuleList

`func (o *DepotSpecDepotLevelParam) SetDepotMergeRequestRuleList(v []DepotMergeRequestRuleParam)`

SetDepotMergeRequestRuleList sets DepotMergeRequestRuleList field to given value.

### HasDepotMergeRequestRuleList

`func (o *DepotSpecDepotLevelParam) HasDepotMergeRequestRuleList() bool`

HasDepotMergeRequestRuleList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


