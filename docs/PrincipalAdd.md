# PrincipalAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyIds** | **[]int64** | 权限组ID | 
**PrincipalId** | **string** | 主体ID | [default to ""]
**PrincipalType** | **string** | 主体类型 | [default to ""]

## Methods

### NewPrincipalAdd

`func NewPrincipalAdd(policyIds []int64, principalId string, principalType string, ) *PrincipalAdd`

NewPrincipalAdd instantiates a new PrincipalAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalAddWithDefaults

`func NewPrincipalAddWithDefaults() *PrincipalAdd`

NewPrincipalAddWithDefaults instantiates a new PrincipalAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyIds

`func (o *PrincipalAdd) GetPolicyIds() []int64`

GetPolicyIds returns the PolicyIds field if non-nil, zero value otherwise.

### GetPolicyIdsOk

`func (o *PrincipalAdd) GetPolicyIdsOk() (*[]int64, bool)`

GetPolicyIdsOk returns a tuple with the PolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIds

`func (o *PrincipalAdd) SetPolicyIds(v []int64)`

SetPolicyIds sets PolicyIds field to given value.


### GetPrincipalId

`func (o *PrincipalAdd) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *PrincipalAdd) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *PrincipalAdd) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetPrincipalType

`func (o *PrincipalAdd) GetPrincipalType() string`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *PrincipalAdd) GetPrincipalTypeOk() (*string, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *PrincipalAdd) SetPrincipalType(v string)`

SetPrincipalType sets PrincipalType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


