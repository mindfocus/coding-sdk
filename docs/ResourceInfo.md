# ResourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelativeResource** | Pointer to **string** | 资源三段式描述的资源定位描述，支持 path 形式及模糊表达匹配 | [optional] [default to ""]
**ResourceId** | **string** | 资源 ID | [default to ""]
**ResourceScope** | Pointer to **string** | 资源三段式描述的 scope，目前固定 coding | [optional] [default to ""]
**ResourceType** | **string** | 资源类型：例如 project | [default to ""]
**ServiceName** | Pointer to **string** | 资源三段式描述的 serviceName 部分，表示业务模块 | [optional] [default to ""]

## Methods

### NewResourceInfo

`func NewResourceInfo(resourceId string, resourceType string, ) *ResourceInfo`

NewResourceInfo instantiates a new ResourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceInfoWithDefaults

`func NewResourceInfoWithDefaults() *ResourceInfo`

NewResourceInfoWithDefaults instantiates a new ResourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelativeResource

`func (o *ResourceInfo) GetRelativeResource() string`

GetRelativeResource returns the RelativeResource field if non-nil, zero value otherwise.

### GetRelativeResourceOk

`func (o *ResourceInfo) GetRelativeResourceOk() (*string, bool)`

GetRelativeResourceOk returns a tuple with the RelativeResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeResource

`func (o *ResourceInfo) SetRelativeResource(v string)`

SetRelativeResource sets RelativeResource field to given value.

### HasRelativeResource

`func (o *ResourceInfo) HasRelativeResource() bool`

HasRelativeResource returns a boolean if a field has been set.

### GetResourceId

`func (o *ResourceInfo) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceInfo) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceInfo) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceScope

`func (o *ResourceInfo) GetResourceScope() string`

GetResourceScope returns the ResourceScope field if non-nil, zero value otherwise.

### GetResourceScopeOk

`func (o *ResourceInfo) GetResourceScopeOk() (*string, bool)`

GetResourceScopeOk returns a tuple with the ResourceScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceScope

`func (o *ResourceInfo) SetResourceScope(v string)`

SetResourceScope sets ResourceScope field to given value.

### HasResourceScope

`func (o *ResourceInfo) HasResourceScope() bool`

HasResourceScope returns a boolean if a field has been set.

### GetResourceType

`func (o *ResourceInfo) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceInfo) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceInfo) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetServiceName

`func (o *ResourceInfo) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ResourceInfo) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ResourceInfo) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ResourceInfo) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


