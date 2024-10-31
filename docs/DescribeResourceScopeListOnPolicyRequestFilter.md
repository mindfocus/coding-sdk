# DescribeResourceScopeListOnPolicyRequestFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | Pointer to **NullableString** | 资源类型（精确匹配） | [optional] [default to ""]

## Methods

### NewDescribeResourceScopeListOnPolicyRequestFilter

`func NewDescribeResourceScopeListOnPolicyRequestFilter() *DescribeResourceScopeListOnPolicyRequestFilter`

NewDescribeResourceScopeListOnPolicyRequestFilter instantiates a new DescribeResourceScopeListOnPolicyRequestFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeResourceScopeListOnPolicyRequestFilterWithDefaults

`func NewDescribeResourceScopeListOnPolicyRequestFilterWithDefaults() *DescribeResourceScopeListOnPolicyRequestFilter`

NewDescribeResourceScopeListOnPolicyRequestFilterWithDefaults instantiates a new DescribeResourceScopeListOnPolicyRequestFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *DescribeResourceScopeListOnPolicyRequestFilter) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DescribeResourceScopeListOnPolicyRequestFilter) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DescribeResourceScopeListOnPolicyRequestFilter) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *DescribeResourceScopeListOnPolicyRequestFilter) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *DescribeResourceScopeListOnPolicyRequestFilter) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *DescribeResourceScopeListOnPolicyRequestFilter) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


