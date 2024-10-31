# ModifyCdHostServerGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentMachineId** | **int64** | 堡垒机 ID | 
**AuthMethod** | **string** | SSH 认证方式（可选值：PUBLIC_KEY、PASSWORD） | 
**DisplayName** | **string** | 主机组名称 | 
**Id** | **int64** | 主机组 ID | 
**Ips** | **[]string** | 实例 IP 列表 | 
**Labels** | Pointer to [**[]HostServerGroupLabel**](HostServerGroupLabel.md) | 主机组标签 | [optional] 
**Password** | Pointer to **string** | SSH 密码 | [optional] 
**Port** | **int64** | SSH 端口 | 
**UserName** | **string** | SSH 用户名 | 

## Methods

### NewModifyCdHostServerGroupRequest

`func NewModifyCdHostServerGroupRequest(agentMachineId int64, authMethod string, displayName string, id int64, ips []string, port int64, userName string, ) *ModifyCdHostServerGroupRequest`

NewModifyCdHostServerGroupRequest instantiates a new ModifyCdHostServerGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyCdHostServerGroupRequestWithDefaults

`func NewModifyCdHostServerGroupRequestWithDefaults() *ModifyCdHostServerGroupRequest`

NewModifyCdHostServerGroupRequestWithDefaults instantiates a new ModifyCdHostServerGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentMachineId

`func (o *ModifyCdHostServerGroupRequest) GetAgentMachineId() int64`

GetAgentMachineId returns the AgentMachineId field if non-nil, zero value otherwise.

### GetAgentMachineIdOk

`func (o *ModifyCdHostServerGroupRequest) GetAgentMachineIdOk() (*int64, bool)`

GetAgentMachineIdOk returns a tuple with the AgentMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentMachineId

`func (o *ModifyCdHostServerGroupRequest) SetAgentMachineId(v int64)`

SetAgentMachineId sets AgentMachineId field to given value.


### GetAuthMethod

`func (o *ModifyCdHostServerGroupRequest) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *ModifyCdHostServerGroupRequest) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *ModifyCdHostServerGroupRequest) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.


### GetDisplayName

`func (o *ModifyCdHostServerGroupRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModifyCdHostServerGroupRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModifyCdHostServerGroupRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetId

`func (o *ModifyCdHostServerGroupRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModifyCdHostServerGroupRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModifyCdHostServerGroupRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetIps

`func (o *ModifyCdHostServerGroupRequest) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *ModifyCdHostServerGroupRequest) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *ModifyCdHostServerGroupRequest) SetIps(v []string)`

SetIps sets Ips field to given value.


### GetLabels

`func (o *ModifyCdHostServerGroupRequest) GetLabels() []HostServerGroupLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ModifyCdHostServerGroupRequest) GetLabelsOk() (*[]HostServerGroupLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ModifyCdHostServerGroupRequest) SetLabels(v []HostServerGroupLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ModifyCdHostServerGroupRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPassword

`func (o *ModifyCdHostServerGroupRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModifyCdHostServerGroupRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModifyCdHostServerGroupRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModifyCdHostServerGroupRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *ModifyCdHostServerGroupRequest) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ModifyCdHostServerGroupRequest) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ModifyCdHostServerGroupRequest) SetPort(v int64)`

SetPort sets Port field to given value.


### GetUserName

`func (o *ModifyCdHostServerGroupRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ModifyCdHostServerGroupRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ModifyCdHostServerGroupRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


