# PrincipalDel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalId** | **string** | 主体ID | [default to ""]
**PrincipalType** | **string** | 主体类型，支持USER_GROUP、DEPARTMENT、USER | [default to ""]

## Methods

### NewPrincipalDel

`func NewPrincipalDel(principalId string, principalType string, ) *PrincipalDel`

NewPrincipalDel instantiates a new PrincipalDel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalDelWithDefaults

`func NewPrincipalDelWithDefaults() *PrincipalDel`

NewPrincipalDelWithDefaults instantiates a new PrincipalDel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalId

`func (o *PrincipalDel) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *PrincipalDel) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *PrincipalDel) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetPrincipalType

`func (o *PrincipalDel) GetPrincipalType() string`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *PrincipalDel) GetPrincipalTypeOk() (*string, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *PrincipalDel) SetPrincipalType(v string)`

SetPrincipalType sets PrincipalType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


