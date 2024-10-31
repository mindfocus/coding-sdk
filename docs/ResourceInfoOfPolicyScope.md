# ResourceInfoOfPolicyScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | 资源 ID | [default to ""]
**ResourceType** | **string** | 资源类型：例如 project | [default to ""]

## Methods

### NewResourceInfoOfPolicyScope

`func NewResourceInfoOfPolicyScope(resourceId string, resourceType string, ) *ResourceInfoOfPolicyScope`

NewResourceInfoOfPolicyScope instantiates a new ResourceInfoOfPolicyScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceInfoOfPolicyScopeWithDefaults

`func NewResourceInfoOfPolicyScopeWithDefaults() *ResourceInfoOfPolicyScope`

NewResourceInfoOfPolicyScopeWithDefaults instantiates a new ResourceInfoOfPolicyScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *ResourceInfoOfPolicyScope) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceInfoOfPolicyScope) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceInfoOfPolicyScope) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *ResourceInfoOfPolicyScope) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceInfoOfPolicyScope) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceInfoOfPolicyScope) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


