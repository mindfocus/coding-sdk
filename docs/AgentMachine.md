# AgentMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 堡垒机 id | [optional] 
**Name** | Pointer to **string** | 堡垒机名称 | [optional] [default to ""]
**Status** | Pointer to **string** | 堡垒机状态 | [optional] [default to ""]

## Methods

### NewAgentMachine

`func NewAgentMachine() *AgentMachine`

NewAgentMachine instantiates a new AgentMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentMachineWithDefaults

`func NewAgentMachineWithDefaults() *AgentMachine`

NewAgentMachineWithDefaults instantiates a new AgentMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentMachine) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentMachine) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentMachine) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AgentMachine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AgentMachine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentMachine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentMachine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentMachine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *AgentMachine) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AgentMachine) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AgentMachine) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AgentMachine) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


