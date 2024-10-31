# GrantInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantObjectId** | **string** | 授权对象 ID | [default to ""]
**GrantScope** | **string** | 授权对象类型：USER,USER_GROUP,DEPARTMENT | [default to ""]
**PolicyId** | **int64** | 权限组 ID | 
**PrincipalKey** | Pointer to **string** | 身份 key，后期扩展场景使用，暂时留空即可 | [optional] [default to ""]
**PrincipalValue** | Pointer to **string** | 身份 value，后期扩展场景使用，暂时留空即可 | [optional] [default to ""]

## Methods

### NewGrantInfo

`func NewGrantInfo(grantObjectId string, grantScope string, policyId int64, ) *GrantInfo`

NewGrantInfo instantiates a new GrantInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantInfoWithDefaults

`func NewGrantInfoWithDefaults() *GrantInfo`

NewGrantInfoWithDefaults instantiates a new GrantInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantObjectId

`func (o *GrantInfo) GetGrantObjectId() string`

GetGrantObjectId returns the GrantObjectId field if non-nil, zero value otherwise.

### GetGrantObjectIdOk

`func (o *GrantInfo) GetGrantObjectIdOk() (*string, bool)`

GetGrantObjectIdOk returns a tuple with the GrantObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantObjectId

`func (o *GrantInfo) SetGrantObjectId(v string)`

SetGrantObjectId sets GrantObjectId field to given value.


### GetGrantScope

`func (o *GrantInfo) GetGrantScope() string`

GetGrantScope returns the GrantScope field if non-nil, zero value otherwise.

### GetGrantScopeOk

`func (o *GrantInfo) GetGrantScopeOk() (*string, bool)`

GetGrantScopeOk returns a tuple with the GrantScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantScope

`func (o *GrantInfo) SetGrantScope(v string)`

SetGrantScope sets GrantScope field to given value.


### GetPolicyId

`func (o *GrantInfo) GetPolicyId() int64`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *GrantInfo) GetPolicyIdOk() (*int64, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *GrantInfo) SetPolicyId(v int64)`

SetPolicyId sets PolicyId field to given value.


### GetPrincipalKey

`func (o *GrantInfo) GetPrincipalKey() string`

GetPrincipalKey returns the PrincipalKey field if non-nil, zero value otherwise.

### GetPrincipalKeyOk

`func (o *GrantInfo) GetPrincipalKeyOk() (*string, bool)`

GetPrincipalKeyOk returns a tuple with the PrincipalKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalKey

`func (o *GrantInfo) SetPrincipalKey(v string)`

SetPrincipalKey sets PrincipalKey field to given value.

### HasPrincipalKey

`func (o *GrantInfo) HasPrincipalKey() bool`

HasPrincipalKey returns a boolean if a field has been set.

### GetPrincipalValue

`func (o *GrantInfo) GetPrincipalValue() string`

GetPrincipalValue returns the PrincipalValue field if non-nil, zero value otherwise.

### GetPrincipalValueOk

`func (o *GrantInfo) GetPrincipalValueOk() (*string, bool)`

GetPrincipalValueOk returns a tuple with the PrincipalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalValue

`func (o *GrantInfo) SetPrincipalValue(v string)`

SetPrincipalValue sets PrincipalValue field to given value.

### HasPrincipalValue

`func (o *GrantInfo) HasPrincipalValue() bool`

HasPrincipalValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


