# HostServerGroupDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | 主机组云账号 | [optional] [default to ""]
**AgentMachine** | Pointer to [**AgentMachine**](AgentMachine.md) |  | [optional] 
**AuthMethod** | Pointer to **string** | SSH 认证方式 | [optional] [default to ""]
**DisplayName** | Pointer to **string** | 主机组名称 | [optional] [default to ""]
**Id** | Pointer to **int64** | 主机组 id | [optional] 
**Ips** | Pointer to **[]string** | 实例 IP 列表 | [optional] 
**Labels** | Pointer to [**[]HostServerGroupLabel**](HostServerGroupLabel.md) | 主机组标签 | [optional] 
**Port** | Pointer to **int64** | SSH 端口 | [optional] 
**UserName** | Pointer to **string** | SSH 用户名 | [optional] [default to ""]

## Methods

### NewHostServerGroupDetail

`func NewHostServerGroupDetail() *HostServerGroupDetail`

NewHostServerGroupDetail instantiates a new HostServerGroupDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostServerGroupDetailWithDefaults

`func NewHostServerGroupDetailWithDefaults() *HostServerGroupDetail`

NewHostServerGroupDetailWithDefaults instantiates a new HostServerGroupDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *HostServerGroupDetail) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *HostServerGroupDetail) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *HostServerGroupDetail) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *HostServerGroupDetail) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAgentMachine

`func (o *HostServerGroupDetail) GetAgentMachine() AgentMachine`

GetAgentMachine returns the AgentMachine field if non-nil, zero value otherwise.

### GetAgentMachineOk

`func (o *HostServerGroupDetail) GetAgentMachineOk() (*AgentMachine, bool)`

GetAgentMachineOk returns a tuple with the AgentMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentMachine

`func (o *HostServerGroupDetail) SetAgentMachine(v AgentMachine)`

SetAgentMachine sets AgentMachine field to given value.

### HasAgentMachine

`func (o *HostServerGroupDetail) HasAgentMachine() bool`

HasAgentMachine returns a boolean if a field has been set.

### GetAuthMethod

`func (o *HostServerGroupDetail) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *HostServerGroupDetail) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *HostServerGroupDetail) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *HostServerGroupDetail) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetDisplayName

`func (o *HostServerGroupDetail) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *HostServerGroupDetail) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *HostServerGroupDetail) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *HostServerGroupDetail) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *HostServerGroupDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostServerGroupDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostServerGroupDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *HostServerGroupDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIps

`func (o *HostServerGroupDetail) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *HostServerGroupDetail) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *HostServerGroupDetail) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *HostServerGroupDetail) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetLabels

`func (o *HostServerGroupDetail) GetLabels() []HostServerGroupLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *HostServerGroupDetail) GetLabelsOk() (*[]HostServerGroupLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *HostServerGroupDetail) SetLabels(v []HostServerGroupLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *HostServerGroupDetail) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPort

`func (o *HostServerGroupDetail) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HostServerGroupDetail) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HostServerGroupDetail) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *HostServerGroupDetail) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUserName

`func (o *HostServerGroupDetail) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *HostServerGroupDetail) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *HostServerGroupDetail) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *HostServerGroupDetail) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


