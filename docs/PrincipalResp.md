# PrincipalResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | 添加时间 | [optional] 
**Policies** | Pointer to [**[]Policy**](Policy.md) | 权限组Id | [optional] 
**PrincipalId** | Pointer to **string** | 成员主体Id | [optional] [default to ""]
**PrincipalName** | Pointer to **string** | 成员主体名称 | [optional] [default to ""]
**PrincipalType** | Pointer to **string** | 成员主体类型 | [optional] [default to ""]

## Methods

### NewPrincipalResp

`func NewPrincipalResp() *PrincipalResp`

NewPrincipalResp instantiates a new PrincipalResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalRespWithDefaults

`func NewPrincipalRespWithDefaults() *PrincipalResp`

NewPrincipalRespWithDefaults instantiates a new PrincipalResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PrincipalResp) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PrincipalResp) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PrincipalResp) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PrincipalResp) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPolicies

`func (o *PrincipalResp) GetPolicies() []Policy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *PrincipalResp) GetPoliciesOk() (*[]Policy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *PrincipalResp) SetPolicies(v []Policy)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *PrincipalResp) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### SetPoliciesNil

`func (o *PrincipalResp) SetPoliciesNil(b bool)`

 SetPoliciesNil sets the value for Policies to be an explicit nil

### UnsetPolicies
`func (o *PrincipalResp) UnsetPolicies()`

UnsetPolicies ensures that no value is present for Policies, not even an explicit nil
### GetPrincipalId

`func (o *PrincipalResp) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *PrincipalResp) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *PrincipalResp) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.

### HasPrincipalId

`func (o *PrincipalResp) HasPrincipalId() bool`

HasPrincipalId returns a boolean if a field has been set.

### GetPrincipalName

`func (o *PrincipalResp) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *PrincipalResp) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *PrincipalResp) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *PrincipalResp) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### GetPrincipalType

`func (o *PrincipalResp) GetPrincipalType() string`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *PrincipalResp) GetPrincipalTypeOk() (*string, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *PrincipalResp) SetPrincipalType(v string)`

SetPrincipalType sets PrincipalType field to given value.

### HasPrincipalType

`func (o *PrincipalResp) HasPrincipalType() bool`

HasPrincipalType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


