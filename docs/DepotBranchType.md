# DepotBranchType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | 分支类型的描述 | [optional] [default to ""]
**Name** | Pointer to **NullableString** | 分支类型名字 | [optional] [default to ""]
**Rule** | Pointer to **NullableString** | 分支类型的规则 | [optional] [default to ""]

## Methods

### NewDepotBranchType

`func NewDepotBranchType() *DepotBranchType`

NewDepotBranchType instantiates a new DepotBranchType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepotBranchTypeWithDefaults

`func NewDepotBranchTypeWithDefaults() *DepotBranchType`

NewDepotBranchTypeWithDefaults instantiates a new DepotBranchType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DepotBranchType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DepotBranchType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DepotBranchType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DepotBranchType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DepotBranchType) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DepotBranchType) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *DepotBranchType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DepotBranchType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DepotBranchType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DepotBranchType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DepotBranchType) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DepotBranchType) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetRule

`func (o *DepotBranchType) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *DepotBranchType) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *DepotBranchType) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *DepotBranchType) HasRule() bool`

HasRule returns a boolean if a field has been set.

### SetRuleNil

`func (o *DepotBranchType) SetRuleNil(b bool)`

 SetRuleNil sets the value for Rule to be an explicit nil

### UnsetRule
`func (o *DepotBranchType) UnsetRule()`

UnsetRule ensures that no value is present for Rule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


