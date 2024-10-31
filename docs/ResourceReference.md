# ResourceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **NullableInt64** | 项目ID | [optional] 
**ProjectName** | Pointer to **NullableString** | 项目名称 | [optional] [default to ""]
**ResourceCode** | Pointer to **NullableInt64** | 资源code | [optional] 
**ResourceId** | Pointer to **NullableInt64** | 资源ID | [optional] 
**ResourceName** | Pointer to **NullableString** | 资源名称 | [optional] [default to ""]
**ResourceStatus** | Pointer to **NullableString** | 资源状态 | [optional] [default to ""]
**ResourceType** | Pointer to **NullableString** | 资源类型 | [optional] [default to ""]

## Methods

### NewResourceReference

`func NewResourceReference() *ResourceReference`

NewResourceReference instantiates a new ResourceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceReferenceWithDefaults

`func NewResourceReferenceWithDefaults() *ResourceReference`

NewResourceReferenceWithDefaults instantiates a new ResourceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ResourceReference) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ResourceReference) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ResourceReference) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ResourceReference) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *ResourceReference) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *ResourceReference) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *ResourceReference) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ResourceReference) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ResourceReference) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *ResourceReference) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *ResourceReference) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *ResourceReference) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetResourceCode

`func (o *ResourceReference) GetResourceCode() int64`

GetResourceCode returns the ResourceCode field if non-nil, zero value otherwise.

### GetResourceCodeOk

`func (o *ResourceReference) GetResourceCodeOk() (*int64, bool)`

GetResourceCodeOk returns a tuple with the ResourceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCode

`func (o *ResourceReference) SetResourceCode(v int64)`

SetResourceCode sets ResourceCode field to given value.

### HasResourceCode

`func (o *ResourceReference) HasResourceCode() bool`

HasResourceCode returns a boolean if a field has been set.

### SetResourceCodeNil

`func (o *ResourceReference) SetResourceCodeNil(b bool)`

 SetResourceCodeNil sets the value for ResourceCode to be an explicit nil

### UnsetResourceCode
`func (o *ResourceReference) UnsetResourceCode()`

UnsetResourceCode ensures that no value is present for ResourceCode, not even an explicit nil
### GetResourceId

`func (o *ResourceReference) GetResourceId() int64`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceReference) GetResourceIdOk() (*int64, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceReference) SetResourceId(v int64)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ResourceReference) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *ResourceReference) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *ResourceReference) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetResourceName

`func (o *ResourceReference) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceReference) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ResourceReference) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ResourceReference) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *ResourceReference) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *ResourceReference) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetResourceStatus

`func (o *ResourceReference) GetResourceStatus() string`

GetResourceStatus returns the ResourceStatus field if non-nil, zero value otherwise.

### GetResourceStatusOk

`func (o *ResourceReference) GetResourceStatusOk() (*string, bool)`

GetResourceStatusOk returns a tuple with the ResourceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceStatus

`func (o *ResourceReference) SetResourceStatus(v string)`

SetResourceStatus sets ResourceStatus field to given value.

### HasResourceStatus

`func (o *ResourceReference) HasResourceStatus() bool`

HasResourceStatus returns a boolean if a field has been set.

### SetResourceStatusNil

`func (o *ResourceReference) SetResourceStatusNil(b bool)`

 SetResourceStatusNil sets the value for ResourceStatus to be an explicit nil

### UnsetResourceStatus
`func (o *ResourceReference) UnsetResourceStatus()`

UnsetResourceStatus ensures that no value is present for ResourceStatus, not even an explicit nil
### GetResourceType

`func (o *ResourceReference) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceReference) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceReference) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceReference) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### SetResourceTypeNil

`func (o *ResourceReference) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *ResourceReference) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


