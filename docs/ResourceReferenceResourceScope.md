# ResourceReferenceResourceScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScopeType** | **int64** | 所属主体类型： 1项目、2团队 | [default to 0]
**ScopeName** | **string** | 所属主体名 | [default to ""]
**ScopeDisplayName** | **string** | 所属主体展示名 | [default to ""]
**ScopeId** | **int64** | 所属主体ID | [default to 0]

## Methods

### NewResourceReferenceResourceScope

`func NewResourceReferenceResourceScope(scopeType int64, scopeName string, scopeDisplayName string, scopeId int64, ) *ResourceReferenceResourceScope`

NewResourceReferenceResourceScope instantiates a new ResourceReferenceResourceScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceReferenceResourceScopeWithDefaults

`func NewResourceReferenceResourceScopeWithDefaults() *ResourceReferenceResourceScope`

NewResourceReferenceResourceScopeWithDefaults instantiates a new ResourceReferenceResourceScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopeType

`func (o *ResourceReferenceResourceScope) GetScopeType() int64`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *ResourceReferenceResourceScope) GetScopeTypeOk() (*int64, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *ResourceReferenceResourceScope) SetScopeType(v int64)`

SetScopeType sets ScopeType field to given value.


### GetScopeName

`func (o *ResourceReferenceResourceScope) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *ResourceReferenceResourceScope) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *ResourceReferenceResourceScope) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.


### GetScopeDisplayName

`func (o *ResourceReferenceResourceScope) GetScopeDisplayName() string`

GetScopeDisplayName returns the ScopeDisplayName field if non-nil, zero value otherwise.

### GetScopeDisplayNameOk

`func (o *ResourceReferenceResourceScope) GetScopeDisplayNameOk() (*string, bool)`

GetScopeDisplayNameOk returns a tuple with the ScopeDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeDisplayName

`func (o *ResourceReferenceResourceScope) SetScopeDisplayName(v string)`

SetScopeDisplayName sets ScopeDisplayName field to given value.


### GetScopeId

`func (o *ResourceReferenceResourceScope) GetScopeId() int64`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *ResourceReferenceResourceScope) GetScopeIdOk() (*int64, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *ResourceReferenceResourceScope) SetScopeId(v int64)`

SetScopeId sets ScopeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


