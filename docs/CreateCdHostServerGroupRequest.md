# CreateCdHostServerGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentMachineId** | **int64** | 堡垒机 ID | 
**AuthMethod** | **string** | SSH 认证方式（可选值：PUBLIC_KEY、PASSWORD） | 
**DisplayName** | **string** | 主机组名称 | 
**Ips** | **[]string** | 实例 IP 列表 | 
**Labels** | Pointer to [**[]HostServerGroupLabel**](HostServerGroupLabel.md) | 主机组标签 | [optional] 
**Password** | Pointer to **string** | SSH 密码 | [optional] 
**Port** | **int64** | SSH 端口 | 
**UserName** | **string** | SSH 用户名 | 

## Methods

### NewCreateCdHostServerGroupRequest

`func NewCreateCdHostServerGroupRequest(agentMachineId int64, authMethod string, displayName string, ips []string, port int64, userName string, ) *CreateCdHostServerGroupRequest`

NewCreateCdHostServerGroupRequest instantiates a new CreateCdHostServerGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCdHostServerGroupRequestWithDefaults

`func NewCreateCdHostServerGroupRequestWithDefaults() *CreateCdHostServerGroupRequest`

NewCreateCdHostServerGroupRequestWithDefaults instantiates a new CreateCdHostServerGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentMachineId

`func (o *CreateCdHostServerGroupRequest) GetAgentMachineId() int64`

GetAgentMachineId returns the AgentMachineId field if non-nil, zero value otherwise.

### GetAgentMachineIdOk

`func (o *CreateCdHostServerGroupRequest) GetAgentMachineIdOk() (*int64, bool)`

GetAgentMachineIdOk returns a tuple with the AgentMachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentMachineId

`func (o *CreateCdHostServerGroupRequest) SetAgentMachineId(v int64)`

SetAgentMachineId sets AgentMachineId field to given value.


### GetAuthMethod

`func (o *CreateCdHostServerGroupRequest) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *CreateCdHostServerGroupRequest) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *CreateCdHostServerGroupRequest) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.


### GetDisplayName

`func (o *CreateCdHostServerGroupRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateCdHostServerGroupRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateCdHostServerGroupRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetIps

`func (o *CreateCdHostServerGroupRequest) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *CreateCdHostServerGroupRequest) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *CreateCdHostServerGroupRequest) SetIps(v []string)`

SetIps sets Ips field to given value.


### GetLabels

`func (o *CreateCdHostServerGroupRequest) GetLabels() []HostServerGroupLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateCdHostServerGroupRequest) GetLabelsOk() (*[]HostServerGroupLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateCdHostServerGroupRequest) SetLabels(v []HostServerGroupLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateCdHostServerGroupRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPassword

`func (o *CreateCdHostServerGroupRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateCdHostServerGroupRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateCdHostServerGroupRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateCdHostServerGroupRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *CreateCdHostServerGroupRequest) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateCdHostServerGroupRequest) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateCdHostServerGroupRequest) SetPort(v int64)`

SetPort sets Port field to given value.


### GetUserName

`func (o *CreateCdHostServerGroupRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *CreateCdHostServerGroupRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *CreateCdHostServerGroupRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


