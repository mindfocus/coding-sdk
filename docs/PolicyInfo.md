# PolicyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | 显示名称 | [optional] [default to ""]
**CurrentVersion** | Pointer to **string** | 版本号 | [optional] [default to ""]
**CurrentVersionId** | Pointer to **int64** | 版本记录 ID | [optional] 
**Description** | Pointer to **NullableString** | 描述 | [optional] [default to ""]
**Name** | Pointer to **string** | 名称 | [optional] [default to ""]
**PolicyId** | Pointer to **int64** | 权限组 ID | [optional] 
**PolicyType** | Pointer to **string** | 类型：IDENTITY | RESOURCE | [optional] [default to ""]
**Scope** | Pointer to **string** | 作用范围：SYSTEM | CUSTOM | [optional] [default to ""]

## Methods

### NewPolicyInfo

`func NewPolicyInfo() *PolicyInfo`

NewPolicyInfo instantiates a new PolicyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyInfoWithDefaults

`func NewPolicyInfoWithDefaults() *PolicyInfo`

NewPolicyInfoWithDefaults instantiates a new PolicyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *PolicyInfo) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *PolicyInfo) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *PolicyInfo) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *PolicyInfo) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *PolicyInfo) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *PolicyInfo) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *PolicyInfo) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *PolicyInfo) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetCurrentVersionId

`func (o *PolicyInfo) GetCurrentVersionId() int64`

GetCurrentVersionId returns the CurrentVersionId field if non-nil, zero value otherwise.

### GetCurrentVersionIdOk

`func (o *PolicyInfo) GetCurrentVersionIdOk() (*int64, bool)`

GetCurrentVersionIdOk returns a tuple with the CurrentVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersionId

`func (o *PolicyInfo) SetCurrentVersionId(v int64)`

SetCurrentVersionId sets CurrentVersionId field to given value.

### HasCurrentVersionId

`func (o *PolicyInfo) HasCurrentVersionId() bool`

HasCurrentVersionId returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PolicyInfo) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PolicyInfo) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *PolicyInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyId

`func (o *PolicyInfo) GetPolicyId() int64`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PolicyInfo) GetPolicyIdOk() (*int64, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PolicyInfo) SetPolicyId(v int64)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *PolicyInfo) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyType

`func (o *PolicyInfo) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *PolicyInfo) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *PolicyInfo) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *PolicyInfo) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetScope

`func (o *PolicyInfo) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PolicyInfo) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PolicyInfo) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *PolicyInfo) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


