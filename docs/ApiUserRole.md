# ApiUserRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleId** | Pointer to **int32** | 用户组Id | [optional] [default to 0]
**RoleType** | Pointer to **string** | 用户组类型： UserDefined 用户自定义的角色， EnterpriseOwner 企业所有者，EnterpriseAdmin 企业管理员， EnterpriseMember 企业普通成员， ProjectAdmin 项目管理员， ProjectMember 项目成员-&gt; 新的权限系统里面叫\&quot;开发\&quot;，ProjectGuest 项目受限成员 -&gt; 新的权限系统里面叫\&quot;测试\&quot;，ProjectManager 项目经理，ProductManager 产品经理，ProjectOperation 运维 ProgramOwner 项目集负责人，ProgramAdmin 项目集管理员，ProgramMember 项目集成员， ProgramProjectMember 项目集-项目成员 | [optional] [default to ""]
**RoleTypeName** | Pointer to **string** | 用户组类型名称 | [optional] [default to ""]

## Methods

### NewApiUserRole

`func NewApiUserRole() *ApiUserRole`

NewApiUserRole instantiates a new ApiUserRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserRoleWithDefaults

`func NewApiUserRoleWithDefaults() *ApiUserRole`

NewApiUserRoleWithDefaults instantiates a new ApiUserRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleId

`func (o *ApiUserRole) GetRoleId() int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ApiUserRole) GetRoleIdOk() (*int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ApiUserRole) SetRoleId(v int32)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *ApiUserRole) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### GetRoleType

`func (o *ApiUserRole) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *ApiUserRole) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *ApiUserRole) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *ApiUserRole) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetRoleTypeName

`func (o *ApiUserRole) GetRoleTypeName() string`

GetRoleTypeName returns the RoleTypeName field if non-nil, zero value otherwise.

### GetRoleTypeNameOk

`func (o *ApiUserRole) GetRoleTypeNameOk() (*string, bool)`

GetRoleTypeNameOk returns a tuple with the RoleTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTypeName

`func (o *ApiUserRole) SetRoleTypeName(v string)`

SetRoleTypeName sets RoleTypeName field to given value.

### HasRoleTypeName

`func (o *ApiUserRole) HasRoleTypeName() bool`

HasRoleTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


