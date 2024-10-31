# DescribeUsersOnResourceAndGrantObjectGrantInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantObjectId** | **NullableString** | 授权对象 ID | [default to ""]
**GrantScope** | **NullableString** | 授权对象类型：USER,USER_GROUP,DEPARTMENT | [default to ""]

## Methods

### NewDescribeUsersOnResourceAndGrantObjectGrantInfo

`func NewDescribeUsersOnResourceAndGrantObjectGrantInfo(grantObjectId NullableString, grantScope NullableString, ) *DescribeUsersOnResourceAndGrantObjectGrantInfo`

NewDescribeUsersOnResourceAndGrantObjectGrantInfo instantiates a new DescribeUsersOnResourceAndGrantObjectGrantInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeUsersOnResourceAndGrantObjectGrantInfoWithDefaults

`func NewDescribeUsersOnResourceAndGrantObjectGrantInfoWithDefaults() *DescribeUsersOnResourceAndGrantObjectGrantInfo`

NewDescribeUsersOnResourceAndGrantObjectGrantInfoWithDefaults instantiates a new DescribeUsersOnResourceAndGrantObjectGrantInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantObjectId

`func (o *DescribeUsersOnResourceAndGrantObjectGrantInfo) GetGrantObjectId() string`

GetGrantObjectId returns the GrantObjectId field if non-nil, zero value otherwise.

### GetGrantObjectIdOk

`func (o *DescribeUsersOnResourceAndGrantObjectGrantInfo) GetGrantObjectIdOk() (*string, bool)`

GetGrantObjectIdOk returns a tuple with the GrantObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantObjectId

`func (o *DescribeUsersOnResourceAndGrantObjectGrantInfo) SetGrantObjectId(v string)`

SetGrantObjectId sets GrantObjectId field to given value.


### SetGrantObjectIdNil

`func (o *DescribeUsersOnResourceAndGrantObjectGrantInfo) SetGrantObjectIdNil(b bool)`

 SetGrantObjectIdNil sets the value for GrantObjectId to be an explicit nil

### UnsetGrantObjectId
`func (o *DescribeUsersOnResourceAndGrantObjectGrantInfo) UnsetGrantObjectId()`

UnsetGrantObjectId ensures that no value is present for GrantObjectId, not even an explicit nil
### GetGrantScope

`func (o *DescribeUsersOnResourceAndGrantObjectGrantInfo) GetGrantScope() string`

GetGrantScope returns the GrantScope field if non-nil, zero value otherwise.

### GetGrantScopeOk

`func (o *DescribeUsersOnResourceAndGrantObjectGrantInfo) GetGrantScopeOk() (*string, bool)`

GetGrantScopeOk returns a tuple with the GrantScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantScope

`func (o *DescribeUsersOnResourceAndGrantObjectGrantInfo) SetGrantScope(v string)`

SetGrantScope sets GrantScope field to given value.


### SetGrantScopeNil

`func (o *DescribeUsersOnResourceAndGrantObjectGrantInfo) SetGrantScopeNil(b bool)`

 SetGrantScopeNil sets the value for GrantScope to be an explicit nil

### UnsetGrantScope
`func (o *DescribeUsersOnResourceAndGrantObjectGrantInfo) UnsetGrantScope()`

UnsetGrantScope ensures that no value is present for GrantScope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


