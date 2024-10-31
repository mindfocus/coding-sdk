# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyAlias** | Pointer to **string** | 项目成员 | [optional] [default to ""]
**PolicyId** | Pointer to **int64** | 1 | [optional] 
**PolicyName** | Pointer to **string** | ProjectMember类型 | [optional] [default to ""]

## Methods

### NewPolicy

`func NewPolicy() *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyAlias

`func (o *Policy) GetPolicyAlias() string`

GetPolicyAlias returns the PolicyAlias field if non-nil, zero value otherwise.

### GetPolicyAliasOk

`func (o *Policy) GetPolicyAliasOk() (*string, bool)`

GetPolicyAliasOk returns a tuple with the PolicyAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyAlias

`func (o *Policy) SetPolicyAlias(v string)`

SetPolicyAlias sets PolicyAlias field to given value.

### HasPolicyAlias

`func (o *Policy) HasPolicyAlias() bool`

HasPolicyAlias returns a boolean if a field has been set.

### GetPolicyId

`func (o *Policy) GetPolicyId() int64`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *Policy) GetPolicyIdOk() (*int64, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *Policy) SetPolicyId(v int64)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *Policy) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *Policy) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *Policy) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *Policy) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *Policy) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


