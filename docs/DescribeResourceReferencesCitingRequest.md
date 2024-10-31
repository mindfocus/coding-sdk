# DescribeResourceReferencesCitingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScopeType** | **int64** | 所属主体类型： 1项目、2团队 | [default to 0]
**ResourceCode** | **string** | 资源 ID | [default to ""]
**ScopeId** | **int64** | 所属主体 ID | [default to 0]

## Methods

### NewDescribeResourceReferencesCitingRequest

`func NewDescribeResourceReferencesCitingRequest(scopeType int64, resourceCode string, scopeId int64, ) *DescribeResourceReferencesCitingRequest`

NewDescribeResourceReferencesCitingRequest instantiates a new DescribeResourceReferencesCitingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeResourceReferencesCitingRequestWithDefaults

`func NewDescribeResourceReferencesCitingRequestWithDefaults() *DescribeResourceReferencesCitingRequest`

NewDescribeResourceReferencesCitingRequestWithDefaults instantiates a new DescribeResourceReferencesCitingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopeType

`func (o *DescribeResourceReferencesCitingRequest) GetScopeType() int64`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *DescribeResourceReferencesCitingRequest) GetScopeTypeOk() (*int64, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *DescribeResourceReferencesCitingRequest) SetScopeType(v int64)`

SetScopeType sets ScopeType field to given value.


### GetResourceCode

`func (o *DescribeResourceReferencesCitingRequest) GetResourceCode() string`

GetResourceCode returns the ResourceCode field if non-nil, zero value otherwise.

### GetResourceCodeOk

`func (o *DescribeResourceReferencesCitingRequest) GetResourceCodeOk() (*string, bool)`

GetResourceCodeOk returns a tuple with the ResourceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCode

`func (o *DescribeResourceReferencesCitingRequest) SetResourceCode(v string)`

SetResourceCode sets ResourceCode field to given value.


### GetScopeId

`func (o *DescribeResourceReferencesCitingRequest) GetScopeId() int64`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *DescribeResourceReferencesCitingRequest) GetScopeIdOk() (*int64, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *DescribeResourceReferencesCitingRequest) SetScopeId(v int64)`

SetScopeId sets ScopeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


