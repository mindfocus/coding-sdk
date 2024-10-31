# ModifyProjectPermissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionFlag** | **bool** | 权限创建true、删除false | 
**ProjectId** | **int32** | 项目ID | 
**RoleId** | **int32** | 权限ID | 
**UserGK** | **string** | user gk | 

## Methods

### NewModifyProjectPermissionRequest

`func NewModifyProjectPermissionRequest(actionFlag bool, projectId int32, roleId int32, userGK string, ) *ModifyProjectPermissionRequest`

NewModifyProjectPermissionRequest instantiates a new ModifyProjectPermissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyProjectPermissionRequestWithDefaults

`func NewModifyProjectPermissionRequestWithDefaults() *ModifyProjectPermissionRequest`

NewModifyProjectPermissionRequestWithDefaults instantiates a new ModifyProjectPermissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionFlag

`func (o *ModifyProjectPermissionRequest) GetActionFlag() bool`

GetActionFlag returns the ActionFlag field if non-nil, zero value otherwise.

### GetActionFlagOk

`func (o *ModifyProjectPermissionRequest) GetActionFlagOk() (*bool, bool)`

GetActionFlagOk returns a tuple with the ActionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionFlag

`func (o *ModifyProjectPermissionRequest) SetActionFlag(v bool)`

SetActionFlag sets ActionFlag field to given value.


### GetProjectId

`func (o *ModifyProjectPermissionRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModifyProjectPermissionRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModifyProjectPermissionRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetRoleId

`func (o *ModifyProjectPermissionRequest) GetRoleId() int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ModifyProjectPermissionRequest) GetRoleIdOk() (*int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ModifyProjectPermissionRequest) SetRoleId(v int32)`

SetRoleId sets RoleId field to given value.


### GetUserGK

`func (o *ModifyProjectPermissionRequest) GetUserGK() string`

GetUserGK returns the UserGK field if non-nil, zero value otherwise.

### GetUserGKOk

`func (o *ModifyProjectPermissionRequest) GetUserGKOk() (*string, bool)`

GetUserGKOk returns a tuple with the UserGK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGK

`func (o *ModifyProjectPermissionRequest) SetUserGK(v string)`

SetUserGK sets UserGK field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


