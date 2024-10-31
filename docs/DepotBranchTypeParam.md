# DepotBranchTypeParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | 分支类型描述信息 | [optional] [default to ""]
**Name** | **string** | 分支类型名字 | [default to ""]
**Rule** | **string** | 分支类型规则 | [default to ""]

## Methods

### NewDepotBranchTypeParam

`func NewDepotBranchTypeParam(name string, rule string, ) *DepotBranchTypeParam`

NewDepotBranchTypeParam instantiates a new DepotBranchTypeParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepotBranchTypeParamWithDefaults

`func NewDepotBranchTypeParamWithDefaults() *DepotBranchTypeParam`

NewDepotBranchTypeParamWithDefaults instantiates a new DepotBranchTypeParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DepotBranchTypeParam) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DepotBranchTypeParam) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DepotBranchTypeParam) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DepotBranchTypeParam) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *DepotBranchTypeParam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DepotBranchTypeParam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DepotBranchTypeParam) SetName(v string)`

SetName sets Name field to given value.


### GetRule

`func (o *DepotBranchTypeParam) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *DepotBranchTypeParam) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *DepotBranchTypeParam) SetRule(v string)`

SetRule sets Rule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


