# HostServerGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | 主机组云账号 | [optional] [default to ""]
**AgentMachine** | Pointer to [**AgentMachine**](AgentMachine.md) |  | [optional] 
**DisplayName** | Pointer to **string** | 主机组名称 | [optional] [default to ""]
**Id** | Pointer to **int64** | 主机组 id | [optional] 

## Methods

### NewHostServerGroup

`func NewHostServerGroup() *HostServerGroup`

NewHostServerGroup instantiates a new HostServerGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostServerGroupWithDefaults

`func NewHostServerGroupWithDefaults() *HostServerGroup`

NewHostServerGroupWithDefaults instantiates a new HostServerGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *HostServerGroup) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *HostServerGroup) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *HostServerGroup) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *HostServerGroup) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAgentMachine

`func (o *HostServerGroup) GetAgentMachine() AgentMachine`

GetAgentMachine returns the AgentMachine field if non-nil, zero value otherwise.

### GetAgentMachineOk

`func (o *HostServerGroup) GetAgentMachineOk() (*AgentMachine, bool)`

GetAgentMachineOk returns a tuple with the AgentMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentMachine

`func (o *HostServerGroup) SetAgentMachine(v AgentMachine)`

SetAgentMachine sets AgentMachine field to given value.

### HasAgentMachine

`func (o *HostServerGroup) HasAgentMachine() bool`

HasAgentMachine returns a boolean if a field has been set.

### GetDisplayName

`func (o *HostServerGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *HostServerGroup) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *HostServerGroup) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *HostServerGroup) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *HostServerGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostServerGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostServerGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *HostServerGroup) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


