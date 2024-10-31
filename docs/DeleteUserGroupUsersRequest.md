# DeleteUserGroupUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserGroupUserInfos** | [**[]UserGroupUserInfos**](UserGroupUserInfos.md) | 用户组用户信息 | 

## Methods

### NewDeleteUserGroupUsersRequest

`func NewDeleteUserGroupUsersRequest(userGroupUserInfos []UserGroupUserInfos, ) *DeleteUserGroupUsersRequest`

NewDeleteUserGroupUsersRequest instantiates a new DeleteUserGroupUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteUserGroupUsersRequestWithDefaults

`func NewDeleteUserGroupUsersRequestWithDefaults() *DeleteUserGroupUsersRequest`

NewDeleteUserGroupUsersRequestWithDefaults instantiates a new DeleteUserGroupUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserGroupUserInfos

`func (o *DeleteUserGroupUsersRequest) GetUserGroupUserInfos() []UserGroupUserInfos`

GetUserGroupUserInfos returns the UserGroupUserInfos field if non-nil, zero value otherwise.

### GetUserGroupUserInfosOk

`func (o *DeleteUserGroupUsersRequest) GetUserGroupUserInfosOk() (*[]UserGroupUserInfos, bool)`

GetUserGroupUserInfosOk returns a tuple with the UserGroupUserInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupUserInfos

`func (o *DeleteUserGroupUsersRequest) SetUserGroupUserInfos(v []UserGroupUserInfos)`

SetUserGroupUserInfos sets UserGroupUserInfos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


