# DescribePoliciesOnResourceTypeRequestFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyAlias** | Pointer to **string** | 权限组显示名称，模糊匹配 | [optional] [default to ""]
**ResourceType** | **string** | 资源类型，精确匹配 | [default to ""]
**Visible** | Pointer to **string** | 可见范围（默认 true），all：全部；true：用户可见；false：界面不可见（逻辑权限组） | [optional] [default to ""]

## Methods

### NewDescribePoliciesOnResourceTypeRequestFilter

`func NewDescribePoliciesOnResourceTypeRequestFilter(resourceType string, ) *DescribePoliciesOnResourceTypeRequestFilter`

NewDescribePoliciesOnResourceTypeRequestFilter instantiates a new DescribePoliciesOnResourceTypeRequestFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribePoliciesOnResourceTypeRequestFilterWithDefaults

`func NewDescribePoliciesOnResourceTypeRequestFilterWithDefaults() *DescribePoliciesOnResourceTypeRequestFilter`

NewDescribePoliciesOnResourceTypeRequestFilterWithDefaults instantiates a new DescribePoliciesOnResourceTypeRequestFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyAlias

`func (o *DescribePoliciesOnResourceTypeRequestFilter) GetPolicyAlias() string`

GetPolicyAlias returns the PolicyAlias field if non-nil, zero value otherwise.

### GetPolicyAliasOk

`func (o *DescribePoliciesOnResourceTypeRequestFilter) GetPolicyAliasOk() (*string, bool)`

GetPolicyAliasOk returns a tuple with the PolicyAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyAlias

`func (o *DescribePoliciesOnResourceTypeRequestFilter) SetPolicyAlias(v string)`

SetPolicyAlias sets PolicyAlias field to given value.

### HasPolicyAlias

`func (o *DescribePoliciesOnResourceTypeRequestFilter) HasPolicyAlias() bool`

HasPolicyAlias returns a boolean if a field has been set.

### GetResourceType

`func (o *DescribePoliciesOnResourceTypeRequestFilter) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DescribePoliciesOnResourceTypeRequestFilter) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DescribePoliciesOnResourceTypeRequestFilter) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetVisible

`func (o *DescribePoliciesOnResourceTypeRequestFilter) GetVisible() string`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *DescribePoliciesOnResourceTypeRequestFilter) GetVisibleOk() (*string, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *DescribePoliciesOnResourceTypeRequestFilter) SetVisible(v string)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *DescribePoliciesOnResourceTypeRequestFilter) HasVisible() bool`

HasVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

