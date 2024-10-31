# GrantObjectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **int32** | 创建时间 | 
**GrantObjectId** | **string** | 授权对象 ID | [default to ""]
**GrantObjectName** | **string** | 授权对象名称 | [default to ""]
**GrantScope** | **string** | 授权对象类型：USER,USER_GROUP,DEPARTMENT | [default to ""]

## Methods

### NewGrantObjectInfo

`func NewGrantObjectInfo(createdAt int32, grantObjectId string, grantObjectName string, grantScope string, ) *GrantObjectInfo`

NewGrantObjectInfo instantiates a new GrantObjectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantObjectInfoWithDefaults

`func NewGrantObjectInfoWithDefaults() *GrantObjectInfo`

NewGrantObjectInfoWithDefaults instantiates a new GrantObjectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *GrantObjectInfo) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GrantObjectInfo) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GrantObjectInfo) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetGrantObjectId

`func (o *GrantObjectInfo) GetGrantObjectId() string`

GetGrantObjectId returns the GrantObjectId field if non-nil, zero value otherwise.

### GetGrantObjectIdOk

`func (o *GrantObjectInfo) GetGrantObjectIdOk() (*string, bool)`

GetGrantObjectIdOk returns a tuple with the GrantObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantObjectId

`func (o *GrantObjectInfo) SetGrantObjectId(v string)`

SetGrantObjectId sets GrantObjectId field to given value.


### GetGrantObjectName

`func (o *GrantObjectInfo) GetGrantObjectName() string`

GetGrantObjectName returns the GrantObjectName field if non-nil, zero value otherwise.

### GetGrantObjectNameOk

`func (o *GrantObjectInfo) GetGrantObjectNameOk() (*string, bool)`

GetGrantObjectNameOk returns a tuple with the GrantObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantObjectName

`func (o *GrantObjectInfo) SetGrantObjectName(v string)`

SetGrantObjectName sets GrantObjectName field to given value.


### GetGrantScope

`func (o *GrantObjectInfo) GetGrantScope() string`

GetGrantScope returns the GrantScope field if non-nil, zero value otherwise.

### GetGrantScopeOk

`func (o *GrantObjectInfo) GetGrantScopeOk() (*string, bool)`

GetGrantScopeOk returns a tuple with the GrantScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantScope

`func (o *GrantObjectInfo) SetGrantScope(v string)`

SetGrantScope sets GrantScope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


