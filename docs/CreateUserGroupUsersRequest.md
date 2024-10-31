# CreateUserGroupUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserGroupUserInfos** | [**[]UserGroupUserInfos**](UserGroupUserInfos.md) | 用户用户组id列表 | 

## Methods

### NewCreateUserGroupUsersRequest

`func NewCreateUserGroupUsersRequest(userGroupUserInfos []UserGroupUserInfos, ) *CreateUserGroupUsersRequest`

NewCreateUserGroupUsersRequest instantiates a new CreateUserGroupUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserGroupUsersRequestWithDefaults

`func NewCreateUserGroupUsersRequestWithDefaults() *CreateUserGroupUsersRequest`

NewCreateUserGroupUsersRequestWithDefaults instantiates a new CreateUserGroupUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserGroupUserInfos

`func (o *CreateUserGroupUsersRequest) GetUserGroupUserInfos() []UserGroupUserInfos`

GetUserGroupUserInfos returns the UserGroupUserInfos field if non-nil, zero value otherwise.

### GetUserGroupUserInfosOk

`func (o *CreateUserGroupUsersRequest) GetUserGroupUserInfosOk() (*[]UserGroupUserInfos, bool)`

GetUserGroupUserInfosOk returns a tuple with the UserGroupUserInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupUserInfos

`func (o *CreateUserGroupUsersRequest) SetUserGroupUserInfos(v []UserGroupUserInfos)`

SetUserGroupUserInfos sets UserGroupUserInfos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


